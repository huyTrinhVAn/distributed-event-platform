package kafka

import (
	"context"

	kafkago "github.com/segmentio/kafka-go"
)

type SegmentioConsumer struct {
	reader *kafkago.Reader
}

func NewSegmentioConsumer(broker, topic, groupID string) Consumer {
	reader := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers:  []string{broker},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 1,
		MaxBytes: 10e6,
	})

	return &SegmentioConsumer{
		reader: reader,
	}
}

func (c *SegmentioConsumer) ReadMessage() (Message, error) {
	msg, err := c.reader.ReadMessage(context.Background())
	if err != nil {
		return Message{}, err
	}

	return Message{
		Key:   msg.Key,
		Value: msg.Value,
	}, nil
}

func (c *SegmentioConsumer) Close() error {
	return c.reader.Close()
}
