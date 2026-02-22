package cache

import (
	gocache "github.com/patrickmn/go-cache"

	"peercast-yayp/internal/domain"
)

const channelListKey = "channels"

// ChannelCache はチャンネル一覧のキャッシュを提供する。
type ChannelCache struct {
	cache *gocache.Cache
}

func NewChannelCache(c *gocache.Cache) *ChannelCache {
	return &ChannelCache{cache: c}
}

// Get はキャッシュからチャンネル一覧を取得する。
func (c *ChannelCache) Get() ([]*domain.Channel, bool) {
	v, ok := c.cache.Get(channelListKey)
	if !ok {
		return nil, false
	}
	channels, ok := v.([]*domain.Channel)
	return channels, ok
}

// Set はチャンネル一覧をキャッシュに保存する。
func (c *ChannelCache) Set(channels []*domain.Channel) {
	c.cache.Set(channelListKey, channels, gocache.DefaultExpiration)
}
