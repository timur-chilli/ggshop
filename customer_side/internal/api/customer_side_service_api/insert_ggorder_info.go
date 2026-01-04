package customer_side_service_api

import (
	"context"
	"log"

	"github.com/samber/lo"
	"github.com/timur-chilli/ggshop/customer_side/internal/models"
	customerSideApi "github.com/timur-chilli/ggshop/customer_side/internal/pb/customer_side_api"
	pbModels "github.com/timur-chilli/ggshop/customer_side/internal/pb/models"
)

func (s *CustomerSideServiceAPI) InsertGGOrderInfos(ctx context.Context, req *customerSideApi.GetGGOrderInfoInsertRequest) (*customerSideApi.GetGGOrderInfoInsertResponse, error) {
	log.Printf("API: Called InsertGGOrderInfos, new Orders: %v", req.GGOrders)
	err := s.customerSideService.InsertGGOrderInfo(ctx, mapGGOrderInfo(req.GGOrders))
	if err != nil {
		return &customerSideApi.GetGGOrderInfoInsertResponse{}, err
	}
	return &customerSideApi.GetGGOrderInfoInsertResponse{}, nil
}

func mapGGOrderInfo(GGOrderInfo []*pbModels.GGOrderInsertModel) []*models.GGOrderInfo {
	return lo.Map(GGOrderInfo, func(s *pbModels.GGOrderInsertModel, _ int) *models.GGOrderInfo {
		return &models.GGOrderInfo{
			CustomerName: s.CustomerName,
			Email:        s.Email,
			Details:      s.Details,
		}
	})
}
