package main

import (
	"encoding/json"

	"distributed-event-platform/pkg/config"
	kafkawrapper "distributed-event-platform/pkg/kafka"
	"distributed-event-platform/pkg/logger"
	"distributed-event-platform/pkg/model"
)

func main() {
	cfg := config.Load()

	consumer := kafkawrapper.NewSegmentioConsumer(
		cfg.KafkaBroker,
		cfg.KafkaTopicRaw,
		"ingestion-group",
	)
	defer consumer.Close()

	producer := kafkawrapper.NewSegmentioProducer(
		cfg.KafkaBroker,
		cfg.KafkaTopicValidated,
	)
	defer producer.Close()

	logger.Info("ingestion-service started")
	logger.Info("consuming from topic: " + cfg.KafkaTopicRaw)
	logger.Info("publishing to topic: " + cfg.KafkaTopicValidated)

	for {
		msg, err := consumer.ReadMessage()
		if err != nil {
			logger.Error("failed to read message: " + err.Error())
			continue
		}

		var event model.Event
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			logger.Error("failed to unmarshal event: " + err.Error())
			continue
		}

		// enrich nhẹ
		if event.Metadata == nil {
			event.Metadata = map[string]interface{}{}
		}
		event.Metadata["ingested_by"] = "ingestion-service"

		validatedMessage, err := json.Marshal(event)
		if err != nil {
			logger.Error("failed to marshal validated event: " + err.Error())
			continue
		}

		if err := producer.Publish(validatedMessage, event.PartitionKey); err != nil {
			logger.Error("failed to publish validated event: " + err.Error())
			continue
		}

		logger.Info("validated and forwarded event: " + event.EventID)
	}
}
