package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	config kafka.ReaderConfig
}

func NewConsumer(
	brokers []string,
	groupID string,
	topic string,
) Consumer {
	c := Consumer{}
	c.config = kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: groupID,
		Topic:   topic,
	}

	return c
}

func (c *Consumer) SubscribeMsg() {
	reader := kafka.NewReader(c.config)
	defer reader.Close()

	for {
		// Read a single message from Kafka
		message, err := reader.FetchMessage(context.Background())
		if err != nil {
			log.Fatal("Failed to fetch message: ", err)
		}

		// Print the received message
		fmt.Printf("Received message: Key = %s, Value = %s\n", string(message.Key), string(message.Value))

		// Process the message
		// ...

		// Mark the message as processed
		err = reader.CommitMessages(context.Background(), message)
		if err != nil {
			log.Fatal("Failed to commit message: ", err)
		}
	}
}
