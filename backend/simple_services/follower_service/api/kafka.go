package api

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func StartKafkaProducer(broker string, quit chan struct{}) *KafkaProducer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  broker,
		"acks":               "all",
		"enable.idempotence": true,
	})

	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	log.Default().Printf("Follower service connected to Kafka broker at %s\n", broker)

	kp := &KafkaProducer{
		producer: p,
	}

	go kp.listenEvents(quit)

	return kp
}

func (kp *KafkaProducer) Produce(topic string, key []byte, value []byte) {
	kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          value,
	}, nil)
}

func (kp *KafkaProducer) listenEvents(quit chan struct{}) {
	for {
		select {
		case <-quit:
			kp.producer.Close()
			log.Default().Println("Producer closed.")
			return
		case e := <-kp.producer.Events():
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}
}
