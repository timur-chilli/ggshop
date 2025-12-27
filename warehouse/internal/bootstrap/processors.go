package bootstrap

import (
	ggorderInfoProcessor "github.com/timur-chilli/ggshop/warehouse/internal/services/processors/ggorder_info_processor"
	warehouseService "github.com/timur-chilli/ggshop/warehouse/internal/services/warehouse_service"
)

func InitGGOrderInfoProcessor(warehouseService *warehouseService.WarehouseService) *ggorderInfoProcessor.GGOrderInfoProcessor {
	return ggorderInfoProcessor.NewGGOrderInfoProcessor(warehouseService)
}
