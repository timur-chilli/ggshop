package bootstrap

import (
	server "github.com/timur-chilli/ggshop/customer_side/internal/api/customer_side_service_api"
	customerSideService "github.com/timur-chilli/ggshop/customer_side/internal/services/customer_side_service"
)

func InitCustomerSideServiceAPI(customerSideService *customerSideService.CustomerSideService) *server.CustomerSideServiceAPI {
	return server.NewCustomerSideServiceAPI(customerSideService)
}
