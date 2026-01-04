package customer_side_service_api

import (
	"context"
	"log"

	"github.com/samber/lo"
	"github.com/timur-chilli/ggshop/customer_side/internal/models"
	customerSideAPI "github.com/timur-chilli/ggshop/customer_side/internal/pb/customer_side_api"
	pbModels "github.com/timur-chilli/ggshop/customer_side/internal/pb/models"
)

func (s *CustomerSideServiceAPI) GetGGOrderInfoByIDs(ctx context.Context, req *customerSideAPI.GetGGOrderInfoByIDsRequest) (*customerSideAPI.GetGGOrderInfoByIDsResponse, error) {
	log.Printf("API: Received request with IDs: %v", req.Ids)

	response, err := s.customerSideService.GetGGOrderInfoByIDs(ctx, req.Ids)
	if err != nil {
		return &customerSideAPI.GetGGOrderInfoByIDsResponse{}, err
	}
	return &customerSideAPI.GetGGOrderInfoByIDsResponse{GGOrderInfos: mapGGOrderInfoByResponse(response)}, nil
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
