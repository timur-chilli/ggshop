package bootstrap

import (
	"fmt"
	"log"

	"github.com/timur-chilli/ggshop/customer_side/config"
	"github.com/timur-chilli/ggshop/customer_side/internal/storage/pgstorage"
)

func InitPGStorage(cfg *config.Config) *PGStorage.PGStorage {

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)
	storage, err := PGStorage.NewPGStorage(connectionString)
	if err != nil {
		log.Panic(fmt.Sprintf("ошибка инициализации БД, %v", err))
		panic(err)
	}
	return storage
}
