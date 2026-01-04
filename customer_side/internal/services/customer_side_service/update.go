package customerSideService

import (
	"context"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

func (s *CustomerSideService) UpdateGGOrderInfo(ctx context.Context, infos []*models.GGOrderInfo) error {
	return s.ggorderStorage.UpdateGGOrderInfo(ctx, infos)
}
