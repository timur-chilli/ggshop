package ggorder_info_cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/timur-chilli/ggshop/warehouse/internal/models"
)

type cachedGGOrderInfo struct {
	ggorderInfo *models.GGOrderInfo
}

type GGOrderCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewGGOrderCache(client *redis.Client, ttlSeconds int) *GGOrderCache {
	return &GGOrderCache{client: client, ttl: time.Duration(ttlSeconds) * time.Second}
}

func (c *GGOrderCache) key(ggorderID string) string {
	return fmt.Sprintf("ggorderInfo:%s", ggorderID)
}

func (c *GGOrderCache) Get(ctx context.Context, ggorderID string) (*models.GGOrderInfo, bool) {
	if c == nil || c.client == nil {
		return nil, false
	}

	cacheKey := c.key(ggorderID)
	data, err := c.client.Get(ctx, cacheKey).Bytes()
	if err != nil {
		return nil, false
	}

	var cached cachedGGOrderInfo
	if err := json.Unmarshal(data, &cached.ggorderInfo); err != nil {
		return nil, false
	}

	return cached.ggorderInfo, true
}

func (c *GGOrderCache) Set(ctx context.Context, ggorderID string, ggorderInfo *models.GGOrderInfo) error {
	if c == nil || c.client == nil {
		return fmt.Errorf("нет ссылки на redis")
	}

	cacheKey := c.key(ggorderID)
	value, err := json.Marshal(cachedGGOrderInfo{ggorderInfo: ggorderInfo})
	if err != nil {
		return err
	}

	return c.client.Set(ctx, cacheKey, value, c.ttl).Err()
}
