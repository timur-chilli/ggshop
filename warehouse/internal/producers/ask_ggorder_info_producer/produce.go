package ask_ggorder_info_producer

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/timur-chilli/ggshop/warehouse/config"
)

type AskGGOrderInfoProducer struct {
	w *kafka.Writer
}

func NewAskGGOrderInfoProducer(cfg config.KafkaConfig) *AskGGOrderInfoProducer {
	return &AskGGOrderInfoProducer{
		w: &kafka.Writer{
			Addr:         kafka.TCP(fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)),
			Topic:        cfg.AskForGGOrderInfoTopicName,
			Balancer:     &kafka.Hash{},
			BatchTimeout: 20 * time.Millisecond,
			RequiredAcks: kafka.RequireOne,
		},
	}
}

func (p *AskGGOrderInfoProducer) Send(ctx context.Context, ggorderID uint64) error {
	encodedInfo, err := json.Marshal(ggorderID)
	if err != nil {
		return err
	}
	if err := p.w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(strconv.FormatUint(ggorderID, 10)),
		Value: encodedInfo,
	}); err != nil {
		return fmt.Errorf("kafka write: %w", err)
	}
	return nil
}

func (p *AskGGOrderInfoProducer) Close() error {
	return p.w.Close()
}
