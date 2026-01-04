package GGOrderInfoProcessor

import (
	"context"
	"log"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

func (p *GGOrderInfoProcessor) HandleRemoteEdit(ctx context.Context, infos *models.GGOrderInfo) error {
	log.Printf("Triggered edit Handle: %+v", infos)
	return p.ggorderService.UpdateGGOrderInfo(ctx, []*models.GGOrderInfo{infos})
}

func (p *GGOrderInfoProcessor) HandleRemoteGetByID(ctx context.Context, ggorderInfoID []uint64) error {
	log.Printf("Triggered remote get-by-id Handle: %+v", ggorderInfoID)
	ggorders, err := p.ggorderService.GetGGOrderInfoByIDs(ctx, ggorderInfoID)
	err = p.provideGGOrderInfoProducer.Send(ctx, ggorders)
	// defer p.provideGGOrderInfoProducer.Close()
	return err
}

func (p *GGOrderInfoProcessor) HandleCreateOrder(ctx context.Context, infos []*models.GGOrderInfo) error {
	log.Printf("Triggered create order Handle: %+v", infos)
	return nil
}

func (p *GGOrderInfoProcessor) HandleGetOrder(ctx context.Context, infos []*models.GGOrderInfo) error {
	log.Printf("Triggered get order Handle: %+v", infos)
	return nil
}