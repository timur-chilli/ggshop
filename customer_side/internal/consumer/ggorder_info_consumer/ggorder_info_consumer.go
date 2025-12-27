package ggorder_info_consumer

import (
	"context"
)

type GGOrderInfoProcessor interface {
	RemoteHandleIDs(ctx context.Context, ggorderInfoID []uint64) error
}

type AskGGOrderInfoConsumer struct {
	GGOrderInfoProcessor GGOrderInfoProcessor
	kafkaBroker          []string
	topicName            string
}

func NewGGOrderInfoConsumer(GGOrderInfoProcessor GGOrderInfoProcessor, kafkaBroker []string, topicName string) *AskGGOrderInfoConsumer {
	return &AskGGOrderInfoConsumer{
		GGOrderInfoProcessor: GGOrderInfoProcessor,
		kafkaBroker:          kafkaBroker,
		topicName:            topicName,
	}
}
