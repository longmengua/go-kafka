package main

import (
	"context"
	"go-kafka/internal/kafka"
)

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	go kafka.Producer(ctx)
	kafka.Consumer(ctx)
}
