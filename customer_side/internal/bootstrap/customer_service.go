package bootstrap

import (
	"context"

	"github.com/timur-chilli/ggshop/customer_side/config"
	customerCreateOrderProducer "github.com/timur-chilli/ggshop/customer_side/internal/producers/customer_create_order_producer"
	customerGetOrderProducer "github.com/timur-chilli/ggshop/customer_side/internal/producers/customer_get_order_producer"
	customerSideService "github.com/timur-chilli/ggshop/customer_side/internal/services/customer_side_service"
	"github.com/timur-chilli/ggshop/customer_side/internal/storage/pgstorage"
)

func InitCustomerSideService(storage *PGStorage.PGStorage, getOrderProducer *customerGetOrderProducer.CustomerGetOrderProducer, createOrderProducer *customerCreateOrderProducer.CustomerCreateOrderProducer, cfg *config.Config) *customerSideService.CustomerSideService {

	return customerSideService.NewCustomerSideService(context.Background(), storage, getOrderProducer, createOrderProducer, cfg.CustomerSideServiceSettings.MinNameLen, cfg.CustomerSideServiceSettings.MaxNameLen)
}
