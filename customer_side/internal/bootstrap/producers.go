package bootstrap

import (
	"github.com/timur-chilli/ggshop/customer_side/config"
	customerCreateOrderProducer "github.com/timur-chilli/ggshop/customer_side/internal/producers/customer_create_order_producer"
	customerGetOrderProducer "github.com/timur-chilli/ggshop/customer_side/internal/producers/customer_get_order_producer"
	provideGGOrderInfoProducer "github.com/timur-chilli/ggshop/customer_side/internal/producers/provide_ggorder_info_producer"
)

func InitProvideGGOrderInfoProducer(config *config.Config) *provideGGOrderInfoProducer.ProvideGGOrderInfoProducer {
	return provideGGOrderInfoProducer.NewProvideGGOrderInfoProducer(config.Kafka)
}

func InitCustomerCreateOrderProducer(config *config.Config) *customerCreateOrderProducer.CustomerCreateOrderProducer {
	return customerCreateOrderProducer.NewCustomerCreateOrderProducer(config.Kafka)
}

func InitCustomerGetOrderProducer(config *config.Config) *customerGetOrderProducer.CustomerGetOrderProducer {
	return customerGetOrderProducer.NewCustomerGetOrderProducer(config.Kafka)
}
