package GGOrderInfoProcessor

import (
	"context"

	"github.com/timur-chilli/ggshop/warehouse/internal/models"
)

func (p *GGOrderInfoProcessor) Handle(ctx context.Context, GGOrderInfos *models.GGOrderInfo) error {
	return p.ggorderService.InsertGGOrderInfo(ctx, []*models.GGOrderInfo{GGOrderInfos})
}

func (p *GGOrderInfoProcessor) GetRemoteIDsHandle(ctx context.Context, ggorderInfoID []*models.GGOrderInfo) error {
	return p.ggorderService.InsertGGOrderInfo(ctx, ggorderInfoID)
}
