package GGOrderInfoProcessor

import (
	"context"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

type GGOrderService interface {
	InsertGGOrderInfo(ctx context.Context, GGOrdersInfos []*models.GGOrderInfo) error
	GetGGOrderInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.GGOrderInfo, error)
	UpdateGGOrderInfo(ctx context.Context, originalInfos []*models.GGOrderInfo) error
}

type ProvideGGOrderInfoProducer interface {
	Send(ctx context.Context, ggorderInfos []*models.GGOrderInfo) error
	Close() error
}

type GGOrderInfoProcessor struct {
	ggorderService             GGOrderService
	provideGGOrderInfoProducer ProvideGGOrderInfoProducer
}

func NewGGOrderInfoProcessor(ggorderService GGOrderService, provideGGOrderInfoProducer ProvideGGOrderInfoProducer) *GGOrderInfoProcessor {
	return &GGOrderInfoProcessor{
		ggorderService: ggorderService,
		provideGGOrderInfoProducer : provideGGOrderInfoProducer,
	}
}
