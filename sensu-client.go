package main

import (
	"log"

	"github.com/gbolo/go-sensu-tls/sensu/transport/rabbitmq"
	"github.com/gbolo/go-sensu-client-tls/sensu"
)

func main() {
	sensuConfig, err := sensu.NewConfigFromFlagSet(sensu.ExtractFlags())

	if err != nil {
		log.Fatalf("Error reading Sensu config: %s", err)
	}

	haConfig, err := sensuConfig.RabbitMQHAConfig()
	if err != nil {
		log.Fatalf("Error reading RabbitMQ HA config: %s", err)
	}

	transport := rabbitmq.NewRabbitMQHATransport(haConfig)

	sensuClient := sensu.NewClient(transport, sensuConfig)

	err = sensuClient.Start()

	if err != nil {
		log.Fatalf("Failed to start the Sensu client: %s", err)
	}
}
