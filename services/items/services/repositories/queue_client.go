package repositories

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	e "items/utils/errors"
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type QueueClient struct {
	Connection *amqp.Connection
}

func NewQueueClient(user string, pass string, host string, port int) *QueueClient {
	Connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", user, pass, host, port))
	failOnError(err, "Failed to connect to RabbitMQ")
	return &QueueClient{
		Connection: Connection,
	}
}

func (qc *QueueClient) SendMessage(itemid string, action string, message string) e.ApiError {
	channel, err := qc.Connection.Channel()

	err = channel.ExchangeDeclare(
		"items", // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return e.NewBadRequestApiError("Failed to declare an exchange")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := message
	err = channel.PublishWithContext(ctx,
		"items",
		fmt.Sprintf("%s.%s", itemid, action),
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return e.NewBadRequestApiError("Failed to publish a message")
	}
	log.Printf("[x] Sent %s\n", body)
	return nil
}
