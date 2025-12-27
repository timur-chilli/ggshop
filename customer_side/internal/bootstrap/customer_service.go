package bootstrap

import (
	"context"

	"github.com/timur-chilli/ggshop/customer_side/config"
	customerSideService "github.com/timur-chilli/ggshop/customer_side/internal/services/customer_side_service"
	"github.com/timur-chilli/ggshop/customer_side/internal/storage/pgstorage"
)

func InitCustomerSideService(storage *PGStorage.PGStorage, cfg *config.Config) *customerSideService.CustomerSideService {

	return customerSideService.NewCustomerSideService(context.Background(), storage, cfg.CustomerSideServiceSettings.MinNameLen, cfg.CustomerSideServiceSettings.MaxNameLen)
}
