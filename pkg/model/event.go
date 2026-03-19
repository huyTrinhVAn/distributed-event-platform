package model

type Event struct {
	EventID       string                 `json:"event_id"`
	CorrelationID string                 `json:"correlation_id"`
	SourceService string                 `json:"source_service"`
	Environment   string                 `json:"environment"`
	EventType     string                 `json:"event_type"`
	Severity      string                 `json:"severity"`
	Timestamp     string                 `json:"timestamp"`
	RetryCount    int                    `json:"retry_count"`
	MaxRetries    int                    `json:"max_retries"`
	PartitionKey  string                 `json:"partition_key"`
	Payload       map[string]interface{} `json:"payload"`
	Metadata      map[string]interface{} `json:"metadata"`
}
