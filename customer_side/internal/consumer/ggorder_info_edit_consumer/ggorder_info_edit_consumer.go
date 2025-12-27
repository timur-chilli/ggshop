package ggorder_info_edit_consumer

import (
	"context"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

type GGOrderInfoProcessor interface {
	Handle(ctx context.Context, GGOrdersInfo *models.GGOrderInfo) error
}

type AskGGOrderInfoEditConsumer struct {
	GGOrderInfoProcessor GGOrderInfoProcessor
	kafkaBroker          []string
	topicName            string
}

func NewGGOrderInfoEditConsumer(GGOrdersInfoProcessor GGOrderInfoProcessor, kafkaBroker []string, topicName string) *AskGGOrderInfoEditConsumer {
	return &AskGGOrderInfoEditConsumer{
		GGOrderInfoProcessor: GGOrdersInfoProcessor,
		kafkaBroker:          kafkaBroker,
		topicName:            topicName,
	}
}
