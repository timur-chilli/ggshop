package customerSideService

import (
	"context"
	"log"
	"log/slog"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

func (s *CustomerSideService) GetGGOrderInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.GGOrderInfo, error) {
	infos, err := s.ggorderStorage.GetGGOrderInfoByIDs(ctx, IDs)
	if err != nil {
		slog.Error("CustomerGetOrderProducer:", "error", err.Error())
	}
	s.getOrderProducer.Send(ctx, infos)
	log.Printf("s.getOrderProducer.Send: %v", infos)
	return infos, err
}
