package bootstrap

import (
	"github.com/timur-chilli/ggshop/customer_side/config"
	provideGGOrderInfoProducer "github.com/timur-chilli/ggshop/customer_side/internal/producers/provide_ggorder_info_producer"
)

func InitProvideGGOrderInfoProducer(config *config.Config) *provideGGOrderInfoProducer.ProvideGGOrderInfoProducer {
	return provideGGOrderInfoProducer.NewProvideGGOrderInfoProducer(config.Kafka)
}
