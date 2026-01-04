package customer_get_order_producer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/timur-chilli/ggshop/customer_side/config"
	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

type CustomerGetOrderProducer struct {
	w *kafka.Writer
}

func NewCustomerGetOrderProducer(cfg config.KafkaConfig) *CustomerGetOrderProducer {
	return &CustomerGetOrderProducer{
		w: &kafka.Writer{
			Addr:         kafka.TCP(fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)),
			Topic:        cfg.CustomerGetOrderTopicName,
			Balancer:     &kafka.Hash{},
			BatchTimeout: 20 * time.Millisecond,
			RequiredAcks: kafka.RequireOne,
		},
	}
}

func (p *CustomerGetOrderProducer) Send(ctx context.Context, ggorderInfos []*models.GGOrderInfo) error {
	for _, v := range ggorderInfos {
		encodedInfo, err := json.Marshal(*v)
		if err != nil {
			return err
		}
		if err := p.w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(v.CustomerName),
			Value: encodedInfo,
		}); err != nil {
			return fmt.Errorf("kafka write: %w", err)
		}
	}
	return nil
}

func (p *CustomerGetOrderProducer) Close() error {
	return p.w.Close()
}
