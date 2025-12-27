package customer_side_service_api

import (
	"context"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
	customerSideAPI "github.com/timur-chilli/ggshop/customer_side/internal/pb/customer_side_api"
)

type customerSideService interface {
	GetGGOrderInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.GGOrderInfo, error)
	InsertGGOrderInfo(ctx context.Context, GGOrderInfos []*models.GGOrderInfo) error
}


type CustomerSideServiceAPI struct {
	customerSideAPI.UnimplementedCustomerSideServiceServer
	customerSideService customerSideService
}

func NewCustomerSideServiceAPI(CustomerSideService customerSideService) *CustomerSideServiceAPI {
	return &CustomerSideServiceAPI{
		customerSideService: CustomerSideService,
	}
}
