package rabbitmq

import (
	"fmt"
	"log"
	"net/http"
	"os"

	consulapi "github.com/hashicorp/consul/api"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitmqServer struct {
}

func NewRabbitmqServer() *RabbitmqServer {
	return &RabbitmqServer{}
}

func (*RabbitmqServer) setupHealthServer() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	log.Println("Health server listening on :9000")
}

func (*RabbitmqServer) startHealthServer() {
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("Health server failed: %v", err)
	} else {
		log.Println("Health server for RabbitMQ start at :9000")
	}
}

func (*RabbitmqServer) setupServiceRegistry() {
	log.Println("Start for grpc service registry")

	config := consulapi.DefaultConfig()
	config.Address = os.Getenv("SERVICE_REGISTRY_URI")
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalf("Cannot connect with service registry %v\n", err)
	}

	address := os.Getenv("ADDRESS")
	rabbitmqServiceId := fmt.Sprintf("%s-rabbitmq", address)
	rabbitmqRegistration := &consulapi.AgentServiceRegistration{
		ID:      rabbitmqServiceId,
		Name:    rabbitmqServiceId,
		Address: address,
		Port:    9000,
		Tags:    []string{"rabbitmq"},
		Check: &consulapi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:9000/health", address), // Health check URL.
			Interval: "10s",                                         // Check every 10 seconds.
			Timeout:  "5s",                                          // Timeout after 5 second.
		},
	}

	err = consul.Agent().ServiceRegister(rabbitmqRegistration)
	if err != nil {
		log.Fatalf("Failed to register RabbitMQ service with Consul: %v", err)
	} else {
		log.Println("Register grpc service successfully")
	}
}

func (*RabbitmqServer) startRabbitmqConsumer() error {
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

	exchangeName := "AUTH_GET_PRIVATE_KEY"
	err = ch.ExchangeDeclare(
		exchangeName, // exchange name
		"topic",      // type: topic
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
		return err
	}

	// Declare a queue for key notifications.
	queueName := "KEY_QUEUE"
	q, err := ch.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return err
	}

	// Bind the queue to the exchange with routing key "key.generated".
	err = ch.QueueBind(
		q.Name,                // queue name
		"key.get.private.key", // routing key
		exchangeName,          // exchange
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to bind queue: %v", err)
		return err
	}

	// Consume messages from the queue.
	msgs, err := ch.Consume(
		q.Name, // queue name
		"",     // consumer tag
		true,   // auto-acknowledge
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
		return err
	}

	forever := make(chan bool)

	go func() {
		log.Println("Inside the loop to process exchange topics")
		for d := range msgs {
			fmt.Printf("Received message: %s\n", d.Body)
			// Here you might trigger a gRPC call to fetch the new key.
		}
	}()

	fmt.Println("Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil

}

func (s *RabbitmqServer) Run() {
	// Setup health server
	s.setupHealthServer()

	// Setup and connect to service registry
	s.setupServiceRegistry()

	// Start health server
	s.startHealthServer()

	// Start rabbitmq consumer
	s.startRabbitmqConsumer()
}
