package infrastructure

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

var cache *gocache.Cache

func NewCache() *gocache.Cache {
	cache = gocache.New(5*time.Minute, 0)

	return cache
}

func GetCache() *gocache.Cache {
	return cache
}
