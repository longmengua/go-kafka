package main

import (
	"go-kafka/internal/kafka"
	"go-kafka/tools"
)

func main() {
	tools.Async(kafka.Producer)
	kafka.Consumer()
}
