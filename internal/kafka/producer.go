package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Config kafka.WriterConfig
}

func (p *Producer) PublishMsg() {
	writer := kafka.NewWriter(p.Config)
	defer writer.Close()

	log.Println("Start Publish")

	for {
		msg := kafka.Message{
			// Partition: 0,
			Key:   []byte("Key"),
			Value: []byte("Hello World! " + time.Now().Format(time.RFC3339)),
		}
		// Write the message to Kafka
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Fatal("Failed to write message: ", err)
		}
		time.Sleep(1 * time.Second)
	}
}

func NewProducer(
	Addresses []string,
	Topic string,
) Producer {
	p := Producer{}
	p.Config = kafka.WriterConfig{
		Brokers: Addresses,
		Topic:   Topic,
	}
	return p
}
