package api

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func StartKafkaConsumer(broker string, quit chan struct{}) *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    broker,
		"group.id":             "feed-gen-service",
		"max.poll.interval.ms": 60000,
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}
	topics := []string{"new-post.update-feed"}
	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s\n", err)
	}

	log.Default().Printf("Feed generation service connected to Kafka broker %s and subscribed to topics %v\n", broker, topics)

	return c
}
