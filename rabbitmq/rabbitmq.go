package rabbitmq

import (
	"fmt"
	"log"
	"net/http"
	"os"

	consulapi "github.com/hashicorp/consul/api"
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

func (s *RabbitmqServer) Run() {
	// Setup health server
	s.setupHealthServer()

	// Setup and connect to service registry
	s.setupServiceRegistry()

	// Setup health server
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("Health server failed: %v", err)
	} else {
		log.Println("Health server for RabbitMQ start at :9000")
	}
}
