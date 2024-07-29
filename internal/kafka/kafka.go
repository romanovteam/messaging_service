package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaProducer interface {
	SendMessage(message string) error
}

type kafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokerAddress, topic string) KafkaProducer {
	return &kafkaProducer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokerAddress),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (k *kafkaProducer) SendMessage(message string) error {
	err := k.writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(message),
		},
	)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v\n", err)
	}
	return err
}
