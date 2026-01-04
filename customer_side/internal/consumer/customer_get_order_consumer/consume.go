package customer_get_order_consumer

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

func (c *CustomerGetOrderConsumer) Consume(ctx context.Context) {
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
			slog.Error("CustomerGetOrderConsumer.consume error", "error", err.Error())
		}
		var infos []*models.GGOrderInfo
		err = json.Unmarshal(msg.Value, &infos)
		if err != nil {
			slog.Error("parse", "error", err)
			continue
		}
		err = c.GGOrderInfoProcessor.HandleGetOrder(ctx, infos)
		if err != nil {
			slog.Error("handle", "error", err)
		}

	}

}
