package customer_create_order_producer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/timur-chilli/ggshop/customer_side/config"
	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

type CustomerCreateOrderProducer struct {
	w *kafka.Writer
}

func NewCustomerCreateOrderProducer(cfg config.KafkaConfig) *CustomerCreateOrderProducer {
	return &CustomerCreateOrderProducer{
		w: &kafka.Writer{
			Addr:         kafka.TCP(fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)),
			Topic:        cfg.CustomerCreateOrderTopicName,
			Balancer:     &kafka.Hash{},
			BatchTimeout: 20 * time.Millisecond,
			RequiredAcks: kafka.RequireOne,
		},
	}
}

func (p *CustomerCreateOrderProducer) Send(ctx context.Context, ggorderInfos []*models.GGOrderInfo) error {
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

func (p *CustomerCreateOrderProducer) Close() error {
	return p.w.Close()
}
