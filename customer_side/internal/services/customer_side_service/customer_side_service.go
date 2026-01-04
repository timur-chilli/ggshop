package customerSideService

import (
	"context"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

type GGOrderStorage interface {
	GetGGOrderInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.GGOrderInfo, error)
	InsertGGOrderInfo(ctx context.Context, GGOrderInfos []*models.GGOrderInfo) error
	UpdateGGOrderInfo(ctx context.Context, originalInfos []*models.GGOrderInfo) error
}

type CustomerCreateOrderProducer interface {
	Send(ctx context.Context, ggorderInfos []*models.GGOrderInfo) error
	Close() error
}

type CustomerGetOrderProducer interface {
	Send(ctx context.Context, ggorderInfos []*models.GGOrderInfo) error
	Close() error
}

type CustomerSideService struct {
	ggorderStorage      GGOrderStorage
	getOrderProducer    CustomerGetOrderProducer
	createOrderProducer CustomerCreateOrderProducer
	minNameLen          int
	maxNameLen          int
}

func NewCustomerSideService(ctx context.Context, ggorderStorage GGOrderStorage, getProd CustomerGetOrderProducer, createProd CustomerCreateOrderProducer, minNameLen, maxNameLen int) *CustomerSideService {
	return &CustomerSideService{
		ggorderStorage: ggorderStorage,
		getOrderProducer: getProd,
		createOrderProducer: createProd,
		minNameLen:     minNameLen,
		maxNameLen:     maxNameLen,
	}
}
