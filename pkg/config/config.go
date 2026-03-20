package config

import "os"

type Config struct {
	AppPort             string
	AppEnv              string
	KafkaBroker         string
	KafkaTopicRaw       string
	KafkaTopicValidated string
}

func Load() Config {
	return Config{
		AppPort:             getEnv("APP_PORT", "8080"),
		AppEnv:              getEnv("APP_ENV", "dev"),
		KafkaBroker:         getEnv("KAFKA_BROKER", "localhost:9092"),
		KafkaTopicRaw:       getEnv("KAFKA_TOPIC_RAW", "events.raw"),
		KafkaTopicValidated: getEnv("KAFKA_TOPIC_VALIDATED", "events.validated"),
	}
}
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
