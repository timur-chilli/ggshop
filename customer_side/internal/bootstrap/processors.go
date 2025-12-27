package bootstrap

import (
	customerSideService "github.com/timur-chilli/ggshop/customer_side/internal/services/customer_side_service"
	ggorderInfoProcessor "github.com/timur-chilli/ggshop/customer_side/internal/services/processors/ggorder_info_processor"
	provideGGOrderInfoProducer "github.com/timur-chilli/ggshop/customer_side/internal/producers/provide_ggorder_info_producer"
)

func InitGGOrderInfoMessagesProcessor(customerSideService *customerSideService.CustomerSideService, ggorderInfoProvider *provideGGOrderInfoProducer.ProvideGGOrderInfoProducer) *ggorderInfoProcessor.GGOrderInfoProcessor {
	return ggorderInfoProcessor.NewGGOrderInfoProcessor(customerSideService, ggorderInfoProvider)
}
