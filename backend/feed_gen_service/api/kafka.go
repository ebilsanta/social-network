package api

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func StartKafkaConsumer(broker string, topic string, quit chan struct{}) *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "feed-gen-service",
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s\n", err)
	}

	log.Default().Printf("Feed generation service connected to Kafka broker %s and subscribed to topic %s\n", broker, topic)

	return c
}
