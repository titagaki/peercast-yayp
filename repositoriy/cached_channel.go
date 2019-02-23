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

func (r *CachedChannelRepository) SetChannels(channels model.ChannelList) {
	r.Cache.Set("ChannelList", channels, gocache.DefaultExpiration)
}

func (r *CachedChannelRepository) GetChannels() (model.ChannelList, bool) {
	channels, ok := r.Cache.Get("ChannelList")
	if !ok {
		return nil, false
	}
	return channels.(model.ChannelList), true
}
