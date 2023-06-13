package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func Producer() {
	for {
		w := &kafka.Writer{
			Addr:                   kafka.TCP("localhost:29092", "localhost:39092"),
			Topic:                  "topic-A",
			Balancer:               &kafka.LeastBytes{},
			AllowAutoTopicCreation: true,
		}

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

		if err := w.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
		time.Sleep(1 * time.Second)
	}
}
