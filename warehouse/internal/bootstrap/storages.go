package bootstrap

import (
	"github.com/timur-chilli/ggshop/warehouse/config"
	ggorderInfoCache "github.com/timur-chilli/ggshop/warehouse/internal/cache/ggorder_info_cache"
	askGGOrderInfoEditProducer "github.com/timur-chilli/ggshop/warehouse/internal/producers/ask_ggorder_info_edit_producer"
	askGGOrderInfoProducer "github.com/timur-chilli/ggshop/warehouse/internal/producers/ask_ggorder_info_producer"
	redisStorage "github.com/timur-chilli/ggshop/warehouse/internal/storage/redis_storage"
	"github.com/timur-chilli/ggshop/warehouse/internal/storage/remote_storage"
)

func InitGGOrderInfoCache(cfg *config.Config) (*ggorderInfoCache.GGOrderCache, error) {
	return redisStorage.NewRedisCache(cfg)
}

func InitRemoteStorage(cache *ggorderInfoCache.GGOrderCache, 
	editGGOrderProducer *askGGOrderInfoEditProducer.AskGGOrderInfoEditProducer,
	ggorderProducer *askGGOrderInfoProducer.AskGGOrderInfoProducer) *remotestorage.RemoteStorage {
		return remotestorage.NewRemoteStorage(cache, ggorderProducer, editGGOrderProducer)
}