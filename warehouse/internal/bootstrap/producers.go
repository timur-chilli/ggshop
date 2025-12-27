package bootstrap

import (

	"github.com/timur-chilli/ggshop/warehouse/config"
	ggorderInfoEditProducer "github.com/timur-chilli/ggshop/warehouse/internal/producers/ask_ggorder_info_edit_producer"
	ggorderInfoProducer "github.com/timur-chilli/ggshop/warehouse/internal/producers/ask_ggorder_info_producer"
)

func InitAskForGGOrderInfoEditProducer(cfg *config.Config) *ggorderInfoEditProducer.AskGGOrderInfoEditProducer {
	// kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return ggorderInfoEditProducer.NewAskGGOrderInfoEditProducer(cfg.Kafka)
}

func InitAskForGGOrderInfoProducer(cfg *config.Config) *ggorderInfoProducer.AskGGOrderInfoProducer {
	// kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return ggorderInfoProducer.NewAskGGOrderInfoProducer(cfg.Kafka)
}
