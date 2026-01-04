package customerSideService

import (
	"context"
	"log"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

func (s *CustomerSideService) InsertGGOrderInfo(ctx context.Context, infos []*models.GGOrderInfo) error {
	if err := s.validateInfo(infos); err != nil {
		return err
	}
	s.createOrderProducer.Send(ctx, infos)
	log.Printf("s.createOrderProducer.Send: %v", infos)
	return s.ggorderStorage.InsertGGOrderInfo(ctx, infos)
}
