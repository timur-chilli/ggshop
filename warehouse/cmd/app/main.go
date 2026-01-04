package main

import (
	"fmt"
	"os"

	"github.com/timur-chilli/ggshop/warehouse/config"
	"github.com/timur-chilli/ggshop/warehouse/internal/bootstrap"
)

func main() {

	cfg, err := config.LoadConfig(os.Getenv("warehouseConfigPath"))
	if err != nil {
		panic(fmt.Sprintf("ошибка парсинга конфигурации, %v", err))
	}

	//Хотел использовать Redis, но решил вырезать, так как не разобрался до конца, как с ним работать
	//Ну и в исходной архитектуре его не было
	//Зато обновление записей всё равно работает
	// cache, err := bootstrap.InitGGOrderInfoCache(cfg)
	// if err != nil {
	// 	panic(fmt.Sprintf("ошибка инициализации кэша записей GGOrderInfo, %v", err))
	// }
	//GGOrderInfo - не очень удачное название для объекта, представляющего собой заказ в магазине
	//(чтобы не было в случае чего конфликта с SQL Order и прочей возни)

	//Отправляют сообщения в Kafka, когда нужно получить запись из БД или изменить её
	askForGGOrderInfoProducer := bootstrap.InitAskForGGOrderInfoProducer(cfg)
	askForGGOrderInfoEditProducer := bootstrap.InitAskForGGOrderInfoEditProducer(cfg)

	//Интерфейс для отправки сообщений и получения ответа из кэша, который пополняется процессором в связке с консьюмером
	storage := bootstrap.InitRemoteStorage(askForGGOrderInfoEditProducer, askForGGOrderInfoProducer)

	//Здесь, думаю, всё понятно
	warehouseService := bootstrap.InitWarehouseService(storage, cfg)
	warehouseServiceAPI := bootstrap.InitWarehouseServiceAPI(warehouseService)

	//Ожидает ответа от customer_side
	provideGGOrderInfoProcessor := bootstrap.InitGGOrderInfoProcessor(warehouseService)

	//Ждёт сообщения от customer_side с ответом, который будет записан в кэш
	provideGGOrderInfoConsumer := bootstrap.InitProvideGGOrderInfoConsumer(provideGGOrderInfoProcessor, cfg)

	bootstrap.AppRun(*warehouseServiceAPI, provideGGOrderInfoConsumer)
}
