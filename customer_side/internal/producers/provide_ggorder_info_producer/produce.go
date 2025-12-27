package provide_ggorder_info_producer

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/timur-chilli/ggshop/customer_side/config"
	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

type ProvideGGOrderInfoProducer struct {
	w *kafka.Writer
}

func NewProvideGGOrderInfoProducer(cfg config.KafkaConfig) *ProvideGGOrderInfoProducer {
	return &ProvideGGOrderInfoProducer{
		w: &kafka.Writer{
			Addr:         kafka.TCP(fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)),
			Topic:        cfg.ProvideGGOrderInfoTopicName,
			Balancer:     &kafka.Hash{},
			BatchTimeout: 20 * time.Millisecond,
			RequiredAcks: kafka.RequireOne,
		},
	}
}

func (p *ProvideGGOrderInfoProducer) Send(ctx context.Context, ggorderInfos []*models.GGOrderInfo) error {
	//не понял, как через lo.Map сделать обработку массива
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

func (p *ProvideGGOrderInfoProducer) Close() error {
	return p.w.Close()
}
