package customer_create_order_consumer

import (
	"context"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

type GGOrderInfoProcessor interface {
	HandleCreateOrder(ctx context.Context, ggorderInfos []*models.GGOrderInfo) error
}

type CustomerCreateOrderConsumer struct {
	GGOrderInfoProcessor GGOrderInfoProcessor
	kafkaBroker          []string
	topicName            string
}

func NewCustomerCreateOrderConsumer(GGOrderInfoProcessor GGOrderInfoProcessor, kafkaBroker []string, topicName string) *CustomerCreateOrderConsumer {
	return &CustomerCreateOrderConsumer{
		GGOrderInfoProcessor: GGOrderInfoProcessor,
		kafkaBroker:          kafkaBroker,
		topicName:            topicName,
	}
}
