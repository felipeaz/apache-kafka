package main

import (
	"apache-kafka/consumer"
	"apache-kafka/producer"
	"context"
	"fmt"
	"log"
	"os"
)

const (
	numberOfMessages = 10
)

func main() {
	topic, err := os.LookupEnv("MESSAGE_TOPIC")
	if !err {
		log.Fatal(err)
	}
	broker1, err := os.LookupEnv("BROKER_1_HOST")
	if !err {
		log.Fatal(err)
	}

	brokers := []string{broker1}
	p := producer.New(topic, brokers)
	c := consumer.New("my-group", topic, brokers)

	ctx := context.TODO()
	go c.Consume(ctx)

	var i int
	for {
		msg := fmt.Sprintf("Msg ID: %d", i)
		p.Produce(ctx, []byte(msg))
		i++
	}
}
