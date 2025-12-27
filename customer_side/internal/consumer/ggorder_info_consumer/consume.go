package ggorder_info_consumer

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
)

func (c *AskGGOrderInfoConsumer) Consume(ctx context.Context) {
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
		var ids []uint64
		var tempId uint64
		err = json.Unmarshal(msg.Value, &tempId)
		if err != nil {
			slog.Error("parse", "error", err)
			continue
		}

		ids = append(ids, tempId)
		err = c.GGOrderInfoProcessor.RemoteHandleIDs(ctx, ids)
		if err != nil {
			slog.Error("handle", "error", err)
		}

	}

}
