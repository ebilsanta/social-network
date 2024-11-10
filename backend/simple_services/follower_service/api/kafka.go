package api

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaClient struct {
	producer *kafka.Producer
	consumer *kafka.Consumer
}

func InitKafka(broker string, quit chan struct{}) *KafkaClient {
	p := startKafkaProducer(broker)
	c := startKafkaConsumer(broker)
	k := &KafkaClient{
		producer: p,
		consumer: c,
	}

	return k
}

func startKafkaProducer(broker string) *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  broker,
		"acks":               "all",
		"enable.idempotence": true,
	})

	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	log.Default().Printf("Follower service connected to Kafka broker at %s\n", broker)

	return p
}

func startKafkaConsumer(broker string) *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    broker,
		"group.id":             "follower-service",
		"max.poll.interval.ms": 60000,
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}
	topics := []string{"new-user.add-graph-user"}
	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s\n", err)
	}
	log.Default().Printf("Follower service connected to Kafka broker %s and subscribed to topics %v\n", broker, topics)

	return c
}

func (k *KafkaClient) Produce(topic string, key []byte, value []byte) {
	k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          value,
	}, nil)
}
