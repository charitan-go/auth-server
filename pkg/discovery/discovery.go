package discovery

import (
	"log"
	"os"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
)

func DiscoverService(serviceName string) string {
	config := consulapi.DefaultConfig()
	config.Address = os.Getenv("SERVICE_REGISTRY_URI")
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}

	services, err := consul.Agent().Services()
	if err != nil {
		log.Fatalf("Failed to retrieve services from Consul: %v", err)
	}

	// Look for the gRPC service explicitly
	if service, ok := services[serviceName]; ok {
		log.Printf("Found %s service at: %s:%d", serviceName, service.Address, service.Port)
		return service.Address + ":" + strconv.Itoa(service.Port)
	}

	log.Fatalf("Service %s not found in Consul", serviceName)
	return ""
}

// func SetupServiceRegistry() {
// 	log.Println("Start for service discovery")
//
// 	config := consulapi.DefaultConfig()
// 	config.Address = os.Getenv("SERVICE_REGISTRY_URI")
// 	consul, err := consulapi.NewClient(config)
// 	if err != nil {
// 		log.Println("Cannot connect with service registry", err)
// 	}
//
// 	serviceId := os.Getenv("SERVICE_ID")
// 	port, _ := strconv.Atoi(os.Getenv("PORT"))
// 	address, _ := os.Hostname()
//
// 	registration := &consulapi.AgentServiceRegistration{
// 		ID:      serviceId,
// 		Name:    serviceId,
// 		Port:    port,
// 		Address: serviceId,
// 		// Address: address,
// 		Check: &consulapi.AgentServiceCheck{
// 			HTTP:     fmt.Sprintf("http://%s:%v/health", serviceId, port),
// 			Interval: "10s",
// 			Timeout:  "5s",
// 		},
// 	}
//
// 	err = consul.Agent().ServiceRegister(registration)
// 	if err != nil {
// 		log.Println(err)
// 		log.Printf("Failed to register service: %s:%v\n", address, port)
// 	} else {
// 		log.Printf("successfully register service: %s:%v\n", address, port)
// 	}
// }
