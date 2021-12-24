package consumer

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func New(groupId, topic string, brokers []string) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			GroupID: groupId,
			Topic:   topic,
		}),
	}
}

func (c *Consumer) Consume(ctx context.Context) {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			panic("failed to read message from broker")
		}
		fmt.Printf("received: %s\n", string(msg.Value))
	}
}
