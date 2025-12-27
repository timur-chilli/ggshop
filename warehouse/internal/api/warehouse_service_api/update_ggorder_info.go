package warehouse_service_api

import (
	"context"

	"github.com/samber/lo"
	"github.com/timur-chilli/ggshop/warehouse/internal/models"
	pbModels "github.com/timur-chilli/ggshop/warehouse/internal/pb/models"
	warehouseApi "github.com/timur-chilli/ggshop/warehouse/internal/pb/warehouse_api"
)

func (s *WarehouseServiceAPI) UpdateGGOrderInfos(ctx context.Context, req *warehouseApi.GGOrderInfoUpdateRequest) (*warehouseApi.GGOrderInfoUpdateResponse, error) {
	err := s.warehouseService.UpdateGGOrderInfo(ctx, mapGGOrderInfo(req.GGOrders))
	if err != nil {
		return &warehouseApi.GGOrderInfoUpdateResponse{}, err
	}
	return &warehouseApi.GGOrderInfoUpdateResponse{}, nil
}

func mapGGOrderInfo(GGOrderInfo []*pbModels.GGOrderUpdateModel) []*models.GGOrderInfo {
	return lo.Map(GGOrderInfo, func(s *pbModels.GGOrderUpdateModel, _ int) *models.GGOrderInfo {
		return &models.GGOrderInfo{
			ID:           s.Id,
			CustomerName: s.CustomerName,
			Email:        s.Email,
			Details:      s.Details,
		}
	})
}
