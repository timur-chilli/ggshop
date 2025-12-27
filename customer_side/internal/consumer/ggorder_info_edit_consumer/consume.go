package ggorder_info_edit_consumer

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

func (c *AskGGOrderInfoEditConsumer) Consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           c.kafkaBroker,
		GroupID:           "CustomerSideService_group",
		Topic:             c.topicName,
		HeartbeatInterval: 3 * time.Second,
		SessionTimeout:    30 * time.Second,
	})
	defer r.Close()

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			slog.Error("GGOrderInfoEditConsumer.consume error", "error", err.Error())
		}
		var GGOrderInfo *models.GGOrderInfo
		err = json.Unmarshal(msg.Value, &GGOrderInfo)
		if err != nil {
			slog.Error("parce", "error", err)
			continue
		}
		err = c.GGOrderInfoProcessor.Handle(ctx, GGOrderInfo)
		if err != nil {
			slog.Error("Handle", "error", err)
		}
	}

}
