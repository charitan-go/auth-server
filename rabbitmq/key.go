package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	KEY_EXCHANGE_NAME = "key.exchange"
	KEY_QUEUE_NAME    = "auth.private_key.queue"
	KEY_ROUTING_KEY   = "key.get_private_key"
)

func (srv *RabbitmqServer) setupGetPrivateKeyConsumer(ch *amqp.Channel) (<-chan amqp.Delivery, error) {
	// Declare exchange name for private key
	err := srv.rabbitmqSvc.DeclareExchange(ch, KEY_EXCHANGE_NAME)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
		return nil, err
	}

	// Declare a queue for key notifications.amqp
	err = srv.rabbitmqSvc.DeclareQueue(ch, KEY_QUEUE_NAME)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return nil, err
	}

	// Bind the queue to the exchange with routing key "key.generated".
	// srv.rabbitmqSvc.
	err = srv.rabbitmqSvc.QueueBind(ch, KEY_QUEUE_NAME, KEY_ROUTING_KEY, KEY_EXCHANGE_NAME)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return nil, err
	}

	// Consume messages from the queue.
	msgs, err := srv.rabbitmqSvc.Consume(ch, KEY_QUEUE_NAME)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
		return nil, err
	}

	return msgs, nil
}
