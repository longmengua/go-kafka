package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Config kafka.WriterConfig
	Msg    kafka.Message
}

func (p *Producer) ProduceMsg() {
	p.Msg = kafka.Message{
		// Partition: 0,
		Key:   []byte("Key"),
		Value: []byte("Hello World! " + time.Now().Format(time.RFC3339)),
	}
}

func (p *Producer) PublishMsg() {
	writer := kafka.NewWriter(p.Config)
	defer writer.Close()
	for {
		// Write the message to Kafka
		err := writer.WriteMessages(context.Background(), p.Msg)
		if err != nil {
			log.Fatal("Failed to write message: ", err)
		}
	}
}

func NewProducer(
	Addresses []string,
	Topic string,
) Producer {
	p := Producer{}
	p.Config = kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
	}
	return p
}
