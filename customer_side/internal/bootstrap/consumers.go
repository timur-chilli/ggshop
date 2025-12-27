package bootstrap

import (
	"fmt"

	"github.com/timur-chilli/ggshop/customer_side/config"
	askGGOrderInfoConsumer "github.com/timur-chilli/ggshop/customer_side/internal/consumer/ggorder_info_consumer"
	askGGOrderInfoEditConsumer "github.com/timur-chilli/ggshop/customer_side/internal/consumer/ggorder_info_edit_consumer"
	ggorderInfoMessagesProcessor "github.com/timur-chilli/ggshop/customer_side/internal/services/processors/ggorder_info_processor"
)

func InitAskForGGOrderInfoEditConsumer(cfg *config.Config, GGOrderInfoMessagesProcessor *ggorderInfoMessagesProcessor.GGOrderInfoProcessor) *askGGOrderInfoEditConsumer.AskGGOrderInfoEditConsumer {
	kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return askGGOrderInfoEditConsumer.NewGGOrderInfoEditConsumer(GGOrderInfoMessagesProcessor, kafkaBrockers, cfg.Kafka.AskForGGOrderInfoEditTopicName)
}

func InitAskForGGOrderInfoConsumer(cfg *config.Config, GGOrderInfoMessagesProcessor *ggorderInfoMessagesProcessor.GGOrderInfoProcessor) *askGGOrderInfoConsumer.AskGGOrderInfoConsumer {
	kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return askGGOrderInfoConsumer.NewGGOrderInfoConsumer(GGOrderInfoMessagesProcessor, kafkaBrockers, cfg.Kafka.AskForGGOrderInfoTopicName)
}
