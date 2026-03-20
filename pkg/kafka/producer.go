package kafka

type Producer interface {
	Publish(message []byte, key string) error
	Close() error
}
