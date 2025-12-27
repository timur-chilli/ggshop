package warehouse_service

import (
	"context"

	"github.com/timur-chilli/ggshop/warehouse/internal/models"
)

type GGOrderStorage interface {
	GetGGOrderInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.GGOrderInfo, error)
	UpdateGGOrderInfo(ctx context.Context, GGOrderInfos []*models.GGOrderInfo) error
	InsertGGOrderInfo(ctx context.Context, GGOrderInfos []*models.GGOrderInfo) error
}

// type Producer interface {
// 	Send(ctx context.Context, ggorderInfos []*models.GGOrderInfo) error
// 	Close() error
// }

type WarehouseService struct {
	storage   GGOrderStorage
	minNameLen int
	maxNameLen int
}

func NewWarehouseService(ctx context.Context, storage GGOrderStorage, minNameLen, maxNameLen int) *WarehouseService {
	return &WarehouseService{
		storage: storage,
		minNameLen: minNameLen,
		maxNameLen: maxNameLen,
	}
}
