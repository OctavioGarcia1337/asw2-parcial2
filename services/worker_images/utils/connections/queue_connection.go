package connections

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
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
