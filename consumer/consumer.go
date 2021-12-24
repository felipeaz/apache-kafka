package consumer

import "github.com/segmentio/kafka-go"

type Consumer struct {
}

func New() Consumer {
	return Consumer{}
}

func (c Consumer) Consume() {
	kafka.Conn{}
}
