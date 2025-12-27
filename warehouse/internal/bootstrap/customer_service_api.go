package bootstrap

import (
	server "github.com/timur-chilli/ggshop/warehouse/internal/api/warehouse_service_api"
	warehouseService "github.com/timur-chilli/ggshop/warehouse/internal/services/warehouse_service"
)

func InitWarehouseServiceAPI(warehouseService *warehouseService.WarehouseService) *server.WarehouseServiceAPI {
	return server.NewWarehouseServiceAPI(warehouseService)
}
