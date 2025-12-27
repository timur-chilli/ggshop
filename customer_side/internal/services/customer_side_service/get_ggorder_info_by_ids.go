package customerSideService

import (
	"context"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

func (s *CustomerSideService) GetGGOrderInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.GGOrderInfo, error) {
	return s.ggorderStorage.GetGGOrderInfoByIDs(ctx, IDs)
}
