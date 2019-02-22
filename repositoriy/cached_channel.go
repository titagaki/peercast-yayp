package repositoriy

import (
	gocache "github.com/patrickmn/go-cache"

	"peercast-yayp/model"
)

type CachedChannelRepository struct {
	*gocache.Cache
}

func NewCachedChannelRepository(cache *gocache.Cache) *CachedChannelRepository {
	return &CachedChannelRepository{cache}
}

func (cache *CachedChannelRepository) SetChannels(channels model.ChannelList) {
	cache.Set("ChannelList", channels, gocache.DefaultExpiration)
}

func (cache *CachedChannelRepository) GetChannels() (model.ChannelList, bool) {
	channels, ok := cache.Get("ChannelList")
	if !ok {
		return nil, false
	}
	return channels.(model.ChannelList), true
}
