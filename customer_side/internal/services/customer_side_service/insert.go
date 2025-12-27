package customerSideService

import (
	"context"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

func (s *CustomerSideService) InsertGGOrderInfo(ctx context.Context, GGOrdersInfos []*models.GGOrderInfo) error {

	if err := s.validateInfo(GGOrdersInfos); err != nil {
		return err
	}
	return s.ggorderStorage.InsertGGOrderInfo(ctx, GGOrdersInfos)
}
