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
