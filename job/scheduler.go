package job

import (
	"github.com/jasonlvhit/gocron"
)

func RunScheduler() {
	s := gocron.NewScheduler()
	s.Every(30).Seconds().Do(SyncChannel)
	<-s.Start()
}
