package main

import (
	"log"

	"github.com/gbolo/go-sensu-tls/sensu/transport/rabbitmq"
	"github.com/gbolo/go-sensu-client-tls/sensu"
)

func main() {
	cfg, err := sensu.NewConfigFromFlagSet(sensu.ExtractFlags())

	if err != nil {
		log.Fatal(err.Error())
	}

	t, err := rabbitmq.NewRabbitMQTransport(cfg.RabbitMQURI())

	if err != nil {
		log.Fatal(err.Error())
	}

	client := sensu.NewClient(t, cfg)

	client.Start()
}
