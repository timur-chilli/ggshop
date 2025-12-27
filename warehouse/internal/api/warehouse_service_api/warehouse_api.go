package warehouse_service_api

import (
	"context"

	"github.com/timur-chilli/ggshop/warehouse/internal/models"
	warehouseAPI "github.com/timur-chilli/ggshop/warehouse/internal/pb/warehouse_api"
)

type warehouseService interface {
	GetGGOrderInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.GGOrderInfo, error)
	UpdateGGOrderInfo(ctx context.Context, GGOrderInfos []*models.GGOrderInfo) error
}

type WarehouseServiceAPI struct {
	warehouseAPI.UnimplementedWarehouseServiceServer
	warehouseService warehouseService
}

func NewWarehouseServiceAPI(WarehouseService warehouseService) *WarehouseServiceAPI {
	return &WarehouseServiceAPI{
		warehouseService: WarehouseService,
	}
}
