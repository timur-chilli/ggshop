package GGOrderInfoProcessor

import (
	"context"

	"github.com/timur-chilli/ggshop/warehouse/internal/models"
)

type GGOrderService interface {
	InsertGGOrderInfo(ctx context.Context, GGOrdersInfos []*models.GGOrderInfo) error
	GetGGOrderInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.GGOrderInfo, error)
}

type GGOrderInfoProcessor struct {
	ggorderService             GGOrderService
}

func NewGGOrderInfoProcessor(ggorderService GGOrderService) *GGOrderInfoProcessor {
	return &GGOrderInfoProcessor{
		ggorderService:             ggorderService,
	}
}
