package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

// New はデフォルト有効期限5分のインメモリキャッシュを生成する。
func New() *gocache.Cache {
	return gocache.New(5*time.Minute, 10*time.Minute)
}
