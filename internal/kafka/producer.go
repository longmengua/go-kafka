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
	w := &kafka.Writer{
		Addr:                   kafka.TCP(p.Config.Brokers...),
		Topic:                  p.Config.Topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	defer func() {
		if err := w.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
	}()

	for {
		err := w.WriteMessages(
			context.Background(),
			kafka.Message{
				// Partition: 0,
				Key:   []byte("Key"),
				Value: []byte("Hello World! " + time.Now().Format(time.RFC3339)),
			},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
		log.Println("write messages ok")
		time.Sleep(500 * time.Millisecond)
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
