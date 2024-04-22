package kafka

import (
	"context"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

func PublishMessage(broker string, topic string, message string) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})

	log.Print("publishing to kafka")
	defer writer.Close()

	msg := kafka.Message{
		Value: []byte(message),
	}

	return writer.WriteMessages(context.Background(), msg)
}
