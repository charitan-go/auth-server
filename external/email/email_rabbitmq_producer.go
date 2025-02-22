package email

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	rabbitmqservice "github.com/charitan-go/auth-server/rabbitmq/service"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	EMAIL_EXCHANGE_NAME   = "email.exchange"
	EMAIL_AUTH_QUEUE_NAME = "email.auth.queue"

	REGISTER_DONOR_ACCOUNT_ROUTING_KEY   = "email.auth.donor.register_account"
	REGISTER_CHARITY_ACCOUNT_ROUTING_KEY = "email.auth.charity.register_account"
)

type EmailRabbitmqProducer interface {
	NotiSendRegisterDonorAccountEmail(reqDto *SendRegisterDonorAccountEmailRequestDto) error
}

type emailRabbitmqProducerImpl struct {
	rabbitmqSvc rabbitmqservice.RabbitmqService
}

func NewEmailRabbitmqProducer(rabbitmqSvc rabbitmqservice.RabbitmqService) EmailRabbitmqProducer {
	return &emailRabbitmqProducerImpl{rabbitmqSvc}
}

func (p *emailRabbitmqProducerImpl) NotiSendRegisterDonorAccountEmail(reqDto *SendRegisterDonorAccountEmailRequestDto) error {
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

	// Declare a topic exchange (will create it if it doesn't exist)
	err = p.rabbitmqSvc.DeclareExchange(ch, EMAIL_EXCHANGE_NAME)
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
		return err
	}

	// Publish a message indicating the email generation is complete
	body, err := json.Marshal(reqDto)
	if err != nil {
		log.Fatalf("Failed to convert object to json: %v", err)
		return err
	}

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}
	err = p.rabbitmqSvc.Publish(ch, EMAIL_EXCHANGE_NAME, REGISTER_DONOR_ACCOUNT_ROUTING_KEY, msg)
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	} else {
		log.Printf("Published message: %s", body)
	}

	return nil
}
