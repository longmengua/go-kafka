package main

import (
	"fmt"
	"go-kafka/internal/kafka"
)

func main() {
	ch := make(chan string, 1)
	urls := []string{"localhost:29092", "localhost:39092"}
	topic := "topic"

	p := kafka.NewProducer(urls, topic)
	go func() {
		p.PublishMsg()
	}()

	groupId := "group"
	c := kafka.NewConsumer(urls, groupId, topic)
	go func() {
		c.SubscribeMsg()
	}()

	fmt.Printf("End %s", <-ch)
}
