package remotestorage

import (
	"context"
	"log"
	"strconv"

	ggorderInfoCache "github.com/timur-chilli/ggshop/warehouse/internal/cache/ggorder_info_cache"
	"github.com/timur-chilli/ggshop/warehouse/internal/models"
)

type GGOrderInfoProducer interface {
	Send(context.Context, uint64) error
	Close() error
}

type GGOrderInfoEditProducer interface {
	Send(context.Context, []*models.GGOrderInfo) error
	Close() error
}

type RemoteStorage struct {
	cache                      *ggorderInfoCache.GGOrderCache
	askGGOrderInfoProducer     GGOrderInfoProducer
	askGGOrderInfoEditProducer GGOrderInfoEditProducer
}

func NewRemoteStorage(cache *ggorderInfoCache.GGOrderCache, iprod GGOrderInfoProducer,
	edprod GGOrderInfoEditProducer) *RemoteStorage {
	return &RemoteStorage{
		cache:                      cache,
		askGGOrderInfoProducer:     iprod,
		askGGOrderInfoEditProducer: edprod,
	}
}

func (s *RemoteStorage) GetGGOrderInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.GGOrderInfo, error) {
	log.Printf("GetGGOrderInfoByIDs called")
	var ggorders []*models.GGOrderInfo
	for _, id := range IDs {
		ggorder, succ := s.cache.Get(context.Background(), strconv.FormatUint(id, 10))
		if succ {
			ggorders = append(ggorders, ggorder)
		} else {
			s.askGGOrderInfoProducer.Send(ctx, id)
			defer s.askGGOrderInfoProducer.Close()
		}
	}
	if ggorders != nil {
		return ggorders, nil
	}
	return nil, nil
}

func (s *RemoteStorage) UpdateGGOrderInfo(ctx context.Context, ggorderInfos []*models.GGOrderInfo) error {
	s.askGGOrderInfoEditProducer.Send(ctx, ggorderInfos)
	defer s.askGGOrderInfoEditProducer.Close()
	return nil
}

// вставка пришедших из другого сервиса значений в кэш
func (s *RemoteStorage) InsertGGOrderInfo(ctx context.Context, ggorderInfos []*models.GGOrderInfo) error {
	log.Printf("Trying to insert %v", ggorderInfos)
	for _, info := range ggorderInfos {
		err := s.cache.Set(context.Background(), strconv.FormatUint(info.ID, 10), info)
		if err != nil {
			s.askGGOrderInfoProducer.Send(ctx, info.ID)
			defer s.askGGOrderInfoProducer.Close()
		}
	}
	return nil
}
