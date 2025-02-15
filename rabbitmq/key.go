package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (srv *RabbitmqServer) setupGetPrivateKeyConsumer(ch *amqp.Channel) (<-chan amqp.Delivery, error) {
	// Declare exchange name for private key
	exchangeName := "key.exchange"
	err := srv.rabbitmqSvc.DeclareExchange(ch, exchangeName)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
		return nil, err
	}

	// Declare a queue for key notifications.amqp
	queueName := "auth.private_key.queue"
	err = srv.rabbitmqSvc.DeclareQueue(ch, queueName)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return nil, err
	}

	// Bind the queue to the exchange with routing key "key.generated".
	// srv.rabbitmqSvc.
	routingKey := "key.get_private_key"
	err = srv.rabbitmqSvc.QueueBind(ch, queueName, routingKey, exchangeName)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return nil, err
	}

	// Consume messages from the queue.
	msgs, err := srv.rabbitmqSvc.Consume(ch, queueName)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
		return nil, err
	}

	return msgs, nil
}
