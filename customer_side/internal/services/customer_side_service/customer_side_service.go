package customerSideService

import (
	"context"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

type GGOrderStorage interface {
	GetGGOrderInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.GGOrderInfo, error)
	InsertGGOrderInfo(ctx context.Context, GGOrderInfos []*models.GGOrderInfo) error
}

type CustomerSideService struct {
	ggorderStorage         GGOrderStorage
	minNameLen             int
	maxNameLen             int
}

func NewCustomerSideService(ctx context.Context, ggorderStorage GGOrderStorage, minNameLen, maxNameLen int) *CustomerSideService {
	return &CustomerSideService{
		ggorderStorage:         ggorderStorage,
		minNameLen:             minNameLen,
		maxNameLen:             maxNameLen,
	}
}
