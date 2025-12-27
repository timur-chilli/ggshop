package provide_ggorder_info_consumer

import (
	"context"

	"github.com/timur-chilli/ggshop/warehouse/internal/models"
)

type GGOrderInfoProcessor interface {
	GetRemoteIDsHandle(ctx context.Context, ggorderInfoID []*models.GGOrderInfo) error
}

type ProvideGGOrderInfoConsumer struct {
	GGOrderInfoProcessor GGOrderInfoProcessor
	kafkaBroker          []string
	topicName            string
}

func NewProvideGGOrderInfoConsumer(GGOrderInfoProcessor GGOrderInfoProcessor, kafkaBroker []string, topicName string) *ProvideGGOrderInfoConsumer {
	return &ProvideGGOrderInfoConsumer{
		GGOrderInfoProcessor: GGOrderInfoProcessor,
		kafkaBroker:          kafkaBroker,
		topicName:            topicName,
	}
}
