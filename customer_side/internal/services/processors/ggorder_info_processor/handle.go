package GGOrderInfoProcessor

import (
	"context"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

func (p *GGOrderInfoProcessor) Handle(ctx context.Context, GGOrderInfos *models.GGOrderInfo) error {
	return p.ggorderService.InsertGGOrderInfo(ctx, []*models.GGOrderInfo{GGOrderInfos})
}

func (p *GGOrderInfoProcessor) RemoteHandleIDs(ctx context.Context, ggorderInfoID []uint64) (error) {
	ggorders, err := p.ggorderService.GetGGOrderInfoByIDs(ctx, ggorderInfoID)
	err = p.provideGGOrderInfoProducer.Send(ctx, ggorders)
	defer p.provideGGOrderInfoProducer.Close()
	return err
}
