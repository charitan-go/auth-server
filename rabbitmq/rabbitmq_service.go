package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitmqService interface {
	DeclareExchange(ch *amqp.Channel, exchangeName string) error
	DeclareQueue(ch *amqp.Channel, queueName string) error
}

type rabbitmqServiceImpl struct {
}

func NewRabbitmqService() RabbitmqService {
	return &rabbitmqServiceImpl{}
}

func (*rabbitmqServiceImpl) DeclareExchange(ch *amqp.Channel, exchangeName string) error {
	// exchangeName := "GET_PRIVATE_KEY"
	err := ch.ExchangeDeclare(
		exchangeName, // exchange name
		"topic",      // type: topic
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	if err != nil {
		log.Fatalf("Declare exchange failed")
	}

	return err

}

func (*rabbitmqServiceImpl) DeclareQueue(ch *amqp.Channel, queueName string) error {
	_, err := ch.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	return err
}
