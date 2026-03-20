package kafka

import (
	"context"
	"time"

	kafkago "github.com/segmentio/kafka-go"
)

type SegmentioProducer struct {
	writer *kafkago.Writer
}

func NewSegmentioProducer(broker, topic string) Producer {
	writer := &kafkago.Writer{
		Addr:         kafkago.TCP(broker),
		Topic:        topic,
		Balancer:     &kafkago.LeastBytes{},
		RequiredAcks: kafkago.RequireOne,
		Async:        false,
	}

	return &SegmentioProducer{
		writer: writer,
	}
}

func (p *SegmentioProducer) Publish(message []byte, key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return p.writer.WriteMessages(ctx, kafkago.Message{
		Key:   []byte(key),
		Value: message,
	})
}

func (p *SegmentioProducer) Close() error {
	return p.writer.Close()
}
