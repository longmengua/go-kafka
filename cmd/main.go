package main

import (
	"fmt"
	"go-kafka/internal/kafka"
)

func main() {
	ch := make(chan string, 1)
	urls := []string{"localhost:29092", "localhost:39092"}
	groupId := "group"
	topic := "topic"

	p := kafka.NewProducer(urls, topic)
	c := kafka.NewConsumer(urls, groupId, topic)

	go func() {
		p.PublishMsg()
	}()

	go func() {
		c.SubscribeMsg()
	}()

	fmt.Printf("End %s", <-ch)
}
