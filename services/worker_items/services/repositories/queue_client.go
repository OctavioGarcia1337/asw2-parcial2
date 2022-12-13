package repositories

import (
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"worker/config"
	e "worker/utils/errors"
)

type QueueClient struct {
	Connection *amqp.Connection
}

func (qc *QueueClient) ProcessMessages(exchange string, topic string, process func(string)) e.ApiError {
	channel, err := qc.Connection.Channel()
	if err != nil {
		return e.NewInternalServerApiError("Error opening channel", err)
	}
	err = channel.ExchangeDeclare(
		exchange,
		"topic",
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		return e.NewBadRequestApiError("Failed to declare an exchange")
	}

	queue, err := channel.QueueDeclare(
		config.QUEUENAME,
		false,
		false,
		true,
		true,
		nil,
	)

	log.Printf("Binding queue %s to exchange %s with routing key %s",
		queue.Name,
		exchange,
		topic)
	err = channel.QueueBind(
		queue.Name,
		topic,
		exchange,
		true,
		nil,
	)

	message, err := channel.Consume(
		queue.Name,
		"worker_items",
		true,
		false,
		false,
		true,
		nil)

	var forever chan struct{}
	go func() {
		for d := range message {
			process(string(d.Body))
		}
	}()
	<-forever
	return nil
}
