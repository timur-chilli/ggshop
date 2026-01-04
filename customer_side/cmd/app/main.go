package main

import (
	"fmt"
	"os"

	"github.com/timur-chilli/ggshop/customer_side/config"
	"github.com/timur-chilli/ggshop/customer_side/internal/bootstrap"
)

func main() {

	cfg, err := config.LoadConfig(os.Getenv("customerConfigPath"))
	if err != nil {
		panic(fmt.Sprintf("ошибка парсинга конфигурации, %v", err))
	}

	//Здесь хранятся все данные о заказах GGOrderInfo (название не самое удачное)
	storage := bootstrap.InitPGStorage(cfg)

	customerCreateOrderProducer := bootstrap.InitCustomerCreateOrderProducer(cfg)
	customerGetOrderProducer := bootstrap.InitCustomerGetOrderProducer(cfg)

	customerSideService := bootstrap.InitCustomerSideService(storage, customerGetOrderProducer, customerCreateOrderProducer, cfg)

	provideGGOrderInfoProducer := bootstrap.InitProvideGGOrderInfoProducer(cfg)

	ggorderInfoMessagesProcessor := bootstrap.InitGGOrderInfoMessagesProcessor(customerSideService, provideGGOrderInfoProducer)
	askForGGOrderInfoConsumer := bootstrap.InitAskForGGOrderInfoConsumer(cfg, ggorderInfoMessagesProcessor)
	askForGGOrderInfoEditConsumer := bootstrap.InitAskForGGOrderInfoEditConsumer(cfg, ggorderInfoMessagesProcessor)
	customerCreateOrderConsumer := bootstrap.InitCustomerCreateOrderConsumer(cfg, ggorderInfoMessagesProcessor)
	customerGetOrderConsumer := bootstrap.InitCustomerGetOrderConsumer(cfg, ggorderInfoMessagesProcessor)
	customerSideAPI := bootstrap.InitCustomerSideServiceAPI(customerSideService)

	bootstrap.AppRun(*customerSideAPI, askForGGOrderInfoEditConsumer, askForGGOrderInfoConsumer, customerCreateOrderConsumer, customerGetOrderConsumer)
}
