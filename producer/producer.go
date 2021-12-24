package producer

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	ctx     context.Context
	brokers []string
	writer  kafka.Writer
}

func New(ctx context.Context, topic string, brokers []string) *Producer {
	return &Producer{
		ctx:     ctx,
		brokers: brokers,
		writer: kafka.Writer{
			Addr:  kafka.TCP(brokers...),
			Topic: topic,
		},
	}
}

func (p *Producer) Produce(message []byte) error {
	err := p.writer.WriteMessages(p.ctx, kafka.Message{
		Key:   []byte(uuid.NewString()),
		Value: message,
	})
	if err != nil {
		panic("failed to write message to broker")
	}
	fmt.Printf("Broker received a message: %s\n", string(message))
	return nil
}
