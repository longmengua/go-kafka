package main

import (
	"fmt"
	"go-kafka/internal/kafka"
	"go-kafka/tools"
)

func main() {
	ch := make(chan string, 1)
	urls := []string{"localhost:29092", "localhost:39092"}
	topic := "topic"

	p := kafka.NewProducer(urls, topic)
	tools.Async(p.PublishMsg)

	groupId := "group"
	c := kafka.NewConsumer(urls, groupId, topic)
	tools.Async(c.SubscribeMsg)

	fmt.Printf("End %s", <-ch)
}
