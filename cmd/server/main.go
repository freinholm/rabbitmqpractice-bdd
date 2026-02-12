package main

import (
	"fmt"
	"log"

	"github.com/freinholm/rabbitmqpractice-bdd/internal/pubsub"
	"github.com/freinholm/rabbitmqpractice-bdd/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")

	const rabbitConnString = "amqp://guest:guest@localhost:5672/"

	rabbitMQ, err := amqp.Dial(rabbitConnString)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	fmt.Println("Connection to RabbitMQ is successful.")

	publishChan, err := rabbitMQ.Channel()
	if err != nil {
		log.Fatalf("Could not create channel in RabbitMQ: %v", err)
	}

	err = pubsub.PublishJSON(
		publishChan,
		routing.ExchangePerilDirect,
		routing.PauseKey,
		routing.PlayingState{
			IsPaused: true,
		},
	)
	if err != nil {
		log.Printf("could not publish time: %v", err)
	}
	fmt.Println("Pause message sent!")
}
