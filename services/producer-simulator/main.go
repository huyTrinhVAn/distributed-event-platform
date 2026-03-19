package main

import (
	"net/http"
	"time"

	"distributed-event-platform/pkg/config"
	"distributed-event-platform/pkg/logger"
	"distributed-event-platform/pkg/model"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "producer-simulator",
			"env":     cfg.AppEnv,
		})
	})
	r.POST("/publish", func(c *gin.Context) {
		event := model.Event{
			EventID:       "evt-001",
			CorrelationID: "corr-001",
			SourceService: "payment-service",
			Environment:   cfg.AppEnv,
			EventType:     "payment.failed",
			Severity:      "high",
			Timestamp:     time.Now().UTC().Format(time.RFC3339),
			RetryCount:    0,
			MaxRetries:    3,
			PartitionKey:  "order-1001",
			Payload: map[string]interface{}{
				"order_id": "order-1001",
				"reason":   "card_declined",
			},
			Metadata: map[string]interface{}{
				"version": "1.0.0",
			},
		}

		logger.Info("created test event in /publish")

		c.JSON(http.StatusOK, gin.H{
			"message": "event created successfully",
			"event":   event,
		})
	})
	logger.Info("starting producer-simulator on port " + cfg.AppPort)

	if err := r.Run(":" + cfg.AppPort); err != nil {
		logger.Error("failed to start server: " + err.Error())
	}

}
