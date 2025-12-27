package bootstrap

import (
	"fmt"

	"github.com/timur-chilli/ggshop/warehouse/config"
	provideGGOrderInfoConsumer "github.com/timur-chilli/ggshop/warehouse/internal/consumer/provide_ggorder_info_consumer"
	ggorderInfoProcessor "github.com/timur-chilli/ggshop/warehouse/internal/services/processors/ggorder_info_processor"
)


func InitProvideGGOrderInfoConsumer(proc *ggorderInfoProcessor.GGOrderInfoProcessor, cfg *config.Config) *provideGGOrderInfoConsumer.ProvideGGOrderInfoConsumer {
	kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return provideGGOrderInfoConsumer.NewProvideGGOrderInfoConsumer(proc, kafkaBrockers, cfg.Kafka.ProvideGGOrderInfoTopicName)
}
