package repositoriy

import (
	gocache "github.com/patrickmn/go-cache"

	"peercast-yayp/model"
)

type CachedInformationRepository struct {
	*gocache.Cache
}

func NewCachedInformationRepository(cache *gocache.Cache) *CachedInformationRepository {
	return &CachedInformationRepository{cache}
}

func (r *CachedInformationRepository) SetInfo(info []*model.Information) {
	r.Cache.Set("Information", info, gocache.DefaultExpiration)
}

func (r *CachedInformationRepository) GetInfo() ([]*model.Information, bool) {
	info, ok := r.Cache.Get("Information")
	if !ok {
		return nil, false
	}
	return info.([]*model.Information), true
}
