package producer

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	brokers []string
	writer  *kafka.Writer
}

func New(topic string, brokers []string) *Producer {
	return &Producer{
		brokers: brokers,
		writer: &kafka.Writer{
			Addr:  kafka.TCP(brokers...),
			Topic: topic,
		},
	}
}

func (p *Producer) Produce(ctx context.Context, message []byte) {
	err := p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(uuid.NewString()),
		Value: message,
	})
	if err != nil {
		panic("failed to write message to broker")
	}
	fmt.Printf("sent: %s\n", string(message))
}
