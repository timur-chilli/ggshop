package bootstrap

import (
	"fmt"

	"github.com/timur-chilli/ggshop/customer_side/config"
	customerCreateOrderConsumer "github.com/timur-chilli/ggshop/customer_side/internal/consumer/customer_create_order_consumer"
	customerGetOrderConsumer "github.com/timur-chilli/ggshop/customer_side/internal/consumer/customer_get_order_consumer"
	askGGOrderInfoConsumer "github.com/timur-chilli/ggshop/customer_side/internal/consumer/ggorder_info_consumer"
	askGGOrderInfoEditConsumer "github.com/timur-chilli/ggshop/customer_side/internal/consumer/ggorder_info_edit_consumer"
	ggorderInfoMessagesProcessor "github.com/timur-chilli/ggshop/customer_side/internal/services/processors/ggorder_info_processor"
)

func InitAskForGGOrderInfoEditConsumer(cfg *config.Config, processor *ggorderInfoMessagesProcessor.GGOrderInfoProcessor) *askGGOrderInfoEditConsumer.AskGGOrderInfoEditConsumer {
	kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return askGGOrderInfoEditConsumer.NewGGOrderInfoEditConsumer(processor, kafkaBrockers, cfg.Kafka.AskForGGOrderInfoEditTopicName)
}

func InitAskForGGOrderInfoConsumer(cfg *config.Config, processor *ggorderInfoMessagesProcessor.GGOrderInfoProcessor) *askGGOrderInfoConsumer.AskGGOrderInfoConsumer {
	kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return askGGOrderInfoConsumer.NewGGOrderInfoConsumer(processor, kafkaBrockers, cfg.Kafka.AskForGGOrderInfoTopicName)
}

func InitCustomerCreateOrderConsumer(cfg *config.Config, processor *ggorderInfoMessagesProcessor.GGOrderInfoProcessor) *customerCreateOrderConsumer.CustomerCreateOrderConsumer {
	kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return customerCreateOrderConsumer.NewCustomerCreateOrderConsumer(processor, kafkaBrockers, cfg.Kafka.CustomerCreateOrderTopicName)
}

func InitCustomerGetOrderConsumer(cfg *config.Config, processor *ggorderInfoMessagesProcessor.GGOrderInfoProcessor) *customerGetOrderConsumer.CustomerGetOrderConsumer {
	kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return customerGetOrderConsumer.NewCustomerGetOrderConsumer(processor, kafkaBrockers, cfg.Kafka.CustomerGetOrderTopicName)
}
