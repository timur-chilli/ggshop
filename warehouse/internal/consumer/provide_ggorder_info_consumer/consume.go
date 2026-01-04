package provide_ggorder_info_consumer

import (
	"context"
	"encoding/json"
	"log"
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
			slog.Error("GGOrderInfoConsumer.Consume error", "error", err.Error())
		}
		var info *models.GGOrderInfo
		err = json.Unmarshal(msg.Value, &info)
		if err != nil {
			slog.Error("ProvideGGOrderInfoConsumer.Consume parse", "error", err)
			continue
		}
		log.Printf("ProvideGGOrderInfoConsumer.Consume success: consumed message %+v", info)
		var infos []*models.GGOrderInfo
		infos = append(infos, info)
		err = c.GGOrderInfoProcessor.GetRemoteIDsHandle(ctx, infos)
		if err != nil {
			slog.Error("ProvideGGOrderInfoConsumer.Consume handle", "error", err)
		}

	}

}
