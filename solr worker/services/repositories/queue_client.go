package repositories

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
	e "wesolr/utils/errors"
)

type QueueClient struct {
	Connection *amqp.Connection
}

func NewQueueClient(user string, pass string, host string, port int) *QueueClient {
	Connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", user, pass, host, port))
	if err != nil {
		log.Panic("Failed to connect to RabbitMQ")
	}
	return &QueueClient{
		Connection: Connection,
	}
}

func (qc *QueueClient) SendMessage(qname string, message string) e.ApiError {
	channel, err := qc.Connection.Channel()
	queue, err := channel.QueueDeclare(
		qname, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return e.NewBadRequestApiError("Failed to declare a queue")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := message
	err = channel.PublishWithContext(ctx,
		"",
		queue.Name,
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

func (qc *QueueClient) ProcessMessages(qname string, process func(string)) e.ApiError {
	channel, err := qc.Connection.Channel()
	message, err := channel.Consume(qname,
		"items",
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		return e.NewBadRequestApiError("Failed to register a consumer")
	}

	go func() {
		for true {
			d := <-message
			process(string(d.Body))
		}
	}()
	return nil
}
