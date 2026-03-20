package kafka

type Message struct {
	Key   []byte
	Value []byte
}

type Consumer interface {
	ReadMessage() (Message, error)
	Close() error
}
