package rabbitmq

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitmqServer struct {
	rabbitmqSvc RabbitmqService
}

func NewRabbitmqServer(rabbitmqSvc RabbitmqService) *RabbitmqServer {
	return &RabbitmqServer{rabbitmqSvc}
}

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

func (srv *RabbitmqServer) startRabbitmqConsumer() error {
	log.Println("In function startRabbitmqConsumer")

	amqpConnectionStr := fmt.Sprintf("amqp://%s:%s@message-broker:5672",
		os.Getenv("MESSAGE_BROKER_USER"),
		os.Getenv("MESSAGE_BROKER_PASSWORD"))
	conn, err := amqp.Dial(amqpConnectionStr)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return err
	}
	defer ch.Close()

	msgs, err := srv.setupGetPrivateKeyConsumer(ch)
	if err != nil {
		log.Fatalf("Setup get private key consumer failed %v\n", err)
		return err
	}

	forever := make(chan bool)
	go func() {
		log.Println("Inside the loop to process exchange topics")
		for d := range msgs {
			if d.Exchange == "GET_PRIVATE_KEY" {
				log.Printf("Received message from exchange GET_PRIVATE_KEY: %s\n", d.Body)
			} else {
				log.Printf("Received message from exchange %s\n", d.Exchange)
			}
		}
	}()

	<-forever

	return nil

}

func (s *RabbitmqServer) Run() {
	// Start rabbitmq consumer
	s.startRabbitmqConsumer()
}
