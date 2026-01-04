package customer_get_order_consumer

import (
	"context"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

type GGOrderInfoProcessor interface {
	HandleCreateOrder(ctx context.Context, ggorderInfos []*models.GGOrderInfo) error
	HandleGetOrder(ctx context.Context, ggorderInfos []*models.GGOrderInfo) error
}

type CustomerGetOrderConsumer struct {
	GGOrderInfoProcessor GGOrderInfoProcessor
	kafkaBroker          []string
	topicName            string
}

func NewCustomerGetOrderConsumer(GGOrderInfoProcessor GGOrderInfoProcessor, kafkaBroker []string, topicName string) *CustomerGetOrderConsumer {
	return &CustomerGetOrderConsumer{
		GGOrderInfoProcessor: GGOrderInfoProcessor,
		kafkaBroker:          kafkaBroker,
		topicName:            topicName,
	}
}
