package job

import (
	"github.com/jasonlvhit/gocron"
	gocache "github.com/patrickmn/go-cache"
)

func RunScheduler(cache *gocache.Cache) {
	s := gocron.NewScheduler()
	s.Every(30).Seconds().Do(SyncChannel, cache)
	s.Every(1).Day().At("0:03").Do(DailySummary)
	<-s.Start()
}
