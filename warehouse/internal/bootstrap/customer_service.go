package bootstrap

import (
	"context"

	"github.com/timur-chilli/ggshop/warehouse/config"
	warehouseService "github.com/timur-chilli/ggshop/warehouse/internal/services/warehouse_service"
)

func InitWarehouseService(storage warehouseService.GGOrderStorage, cfg *config.Config) *warehouseService.WarehouseService {
	return warehouseService.NewWarehouseService(context.Background(), storage, cfg.WarehouseServiceSettings.MinNameLen, cfg.WarehouseServiceSettings.MaxNameLen)
}
