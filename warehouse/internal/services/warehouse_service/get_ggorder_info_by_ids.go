package warehouse_service

import (
	"context"

	"github.com/timur-chilli/ggshop/warehouse/internal/models"
)

func (s *WarehouseService) GetGGOrderInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.GGOrderInfo, error) {
	return s.storage.GetGGOrderInfoByIDs(ctx, IDs)
}
