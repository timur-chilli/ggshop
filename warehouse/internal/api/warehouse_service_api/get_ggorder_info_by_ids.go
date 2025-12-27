package warehouse_service_api

import (
	"context"
	"log"

	"github.com/samber/lo"
	"github.com/timur-chilli/ggshop/warehouse/internal/models"
	pbModels "github.com/timur-chilli/ggshop/warehouse/internal/pb/models"
	warehouseAPI "github.com/timur-chilli/ggshop/warehouse/internal/pb/warehouse_api"
)

func (s *WarehouseServiceAPI) GetGGOrderInfoByIDs(ctx context.Context, req *warehouseAPI.GetGGOrderInfoByIDsRequest) (*warehouseAPI.GetGGOrderInfoByIDsResponse, error) {
	log.Printf("Received request with IDs: %v", req.Ids)

	response, err := s.warehouseService.GetGGOrderInfoByIDs(ctx, req.Ids)
	if err != nil {
		return &warehouseAPI.GetGGOrderInfoByIDsResponse{}, err
	}
	return &warehouseAPI.GetGGOrderInfoByIDsResponse{GGOrderInfos: mapGGOrderInfoByResponse(response)}, nil
}

func mapGGOrderInfoByResponse(ggorderInfos []*models.GGOrderInfo) []*pbModels.GGOrderInfoModel {
	return lo.Map(ggorderInfos, func(s *models.GGOrderInfo, _ int) *pbModels.GGOrderInfoModel {
		return &pbModels.GGOrderInfoModel{
			Id:           s.ID,
			CustomerName: s.CustomerName,
			Email:        s.Email,
			Details:      s.Details,
		}
	})
}
