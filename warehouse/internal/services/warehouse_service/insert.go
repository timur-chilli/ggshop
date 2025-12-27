package warehouse_service

import (
	"context"

	"github.com/timur-chilli/ggshop/warehouse/internal/models"
)

func (s *WarehouseService) InsertGGOrderInfo(ctx context.Context, GGOrdersInfos []*models.GGOrderInfo) error {

	if err := s.validateInfo(GGOrdersInfos); err != nil {
		return err
	}
	return s.storage.InsertGGOrderInfo(ctx, GGOrdersInfos)
}