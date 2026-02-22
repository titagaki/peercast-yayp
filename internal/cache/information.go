package cache

import (
	gocache "github.com/patrickmn/go-cache"

	"peercast-yayp/internal/domain"
)

const informationKey = "information"

// InformationCache はお知らせ情報のキャッシュを提供する。
type InformationCache struct {
	cache *gocache.Cache
}

func NewInformationCache(c *gocache.Cache) *InformationCache {
	return &InformationCache{cache: c}
}

// Get はキャッシュからお知らせ情報を取得する。
func (c *InformationCache) Get() ([]*domain.Information, bool) {
	v, ok := c.cache.Get(informationKey)
	if !ok {
		return nil, false
	}
	info, ok := v.([]*domain.Information)
	return info, ok
}

// Set はお知らせ情報をキャッシュに保存する。
func (c *InformationCache) Set(info []*domain.Information) {
	c.cache.Set(informationKey, info, gocache.DefaultExpiration)
}
