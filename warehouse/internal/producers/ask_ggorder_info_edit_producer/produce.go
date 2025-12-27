package ask_ggorder_info_edit_producer

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/timur-chilli/ggshop/warehouse/config"
	"github.com/timur-chilli/ggshop/warehouse/internal/models"
)

type AskGGOrderInfoEditProducer struct {
	w *kafka.Writer
}

func NewAskGGOrderInfoEditProducer(cfg config.KafkaConfig) *AskGGOrderInfoEditProducer {
	return &AskGGOrderInfoEditProducer{
		w: &kafka.Writer{
			Addr:         kafka.TCP(fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)),
			Topic:        cfg.AskForGGOrderInfoEditTopicName,
			Balancer:     &kafka.Hash{},
			BatchTimeout: 20 * time.Millisecond,
			RequiredAcks: kafka.RequireOne,
		},
	}
}

func (p *AskGGOrderInfoEditProducer) Send(ctx context.Context, ggorderInfos []*models.GGOrderInfo) error {
	//не понял, как через lo.Map сделать обработку массива
	for _, v := range ggorderInfos {
		encodedInfo, err := json.Marshal(*v)
		if err != nil {
			return err
		}
		if err := p.w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.FormatUint(v.ID, 10)),
			Value: encodedInfo,
		}); err != nil {
			return fmt.Errorf("kafka write: %w", err)
		}
	}
	return nil
}

func (p *AskGGOrderInfoEditProducer) Close() error {
	return p.w.Close()
}
