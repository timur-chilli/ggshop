package redis_storage 

import (
	"context"

	"github.com/redis/go-redis/v9"
	cache "github.com/timur-chilli/ggshop/warehouse/internal/cache/ggorder_info_cache"
	"github.com/timur-chilli/ggshop/warehouse/config"

)

func NewRedisCache(cfg *config.Config) (*cache.GGOrderCache, error) {
	redisAddr := cfg.RedisAddr()
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   cfg.Redis.DB,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return cache.NewGGOrderCache(client, cfg.Redis.TTL), nil
}
