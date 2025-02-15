package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (srv *RabbitmqServer) setupGetPrivateKeyConsumer(ch *amqp.Channel) (<-chan amqp.Delivery, error) {
	// Declare exchange name for private key
	exchangeName := "GET_PRIVATE_KEY"
	err := srv.rabbitmqSvc.DeclareExchange(ch, exchangeName)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
		return nil, err
	}

	// Declare a queue for key notifications.amqp
	queueName := "KEY_QUEUE"
	err = srv.rabbitmqSvc.DeclareQueue(ch, queueName)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return nil, err
	}

	// Bind the queue to the exchange with routing key "key.generated".
	routingKey := "key.get.private.key"
	err = ch.QueueBind(
		"",           // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to bind queue: %v", err)
		return nil, err
	}

	// Consume messages from the queue.
	msgs, err := ch.Consume(
		queueName, // queue name
		"",        // consumer tag
		true,      // auto-acknowledge
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
		return nil, err
	}

	return msgs, nil
}
