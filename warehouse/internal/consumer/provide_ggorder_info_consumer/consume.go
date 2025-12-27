package provide_ggorder_info_consumer

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"
	"github.com/segmentio/kafka-go"
	"github.com/timur-chilli/ggshop/warehouse/internal/models"
)

func (c *ProvideGGOrderInfoConsumer) Consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           c.kafkaBroker,
		GroupID:           "WarehouseService_group",
		Topic:             c.topicName,
		HeartbeatInterval: 3 * time.Second,
		SessionTimeout:    30 * time.Second,
	})
	defer r.Close()

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			slog.Error("GGOrderInfoConsumer.consume error", "error", err.Error())
		}
		var ggorderInfos []*models.GGOrderInfo
		err = json.Unmarshal(msg.Value, &ggorderInfos)
		if err != nil {
			slog.Error("parse", "error", err)
			continue
		}
		err = c.GGOrderInfoProcessor.GetRemoteIDsHandle(ctx, ggorderInfos)
		if err != nil {
			slog.Error("handle", "error", err)
		}

	}

}
