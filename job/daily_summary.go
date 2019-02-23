package job

import (
	"time"

	"github.com/labstack/gommon/log"

	"peercast-yayp/config"
	"peercast-yayp/infrastructure"
	"peercast-yayp/model"
	"peercast-yayp/repositoriy"
)

func DailySummary() {
	log.Info("DailySummary is started")

	db, err := infrastructure.NewDB(config.GetConfig())
	if err != nil {
		log.Error(err)
		return
	}
	defer db.Close()

	logTime := time.Now().Add(-24 * time.Hour)
	logs := repositoriy.NewChannelLogRepository(db).FindChannelLogsByLogTime(logTime)
	summaryRepo := repositoriy.NewChannelDailySummaryRepository(db)

	var name string
	var n, sumListeners, maxListeners int
	for _, l := range logs {
		if len(name) > 0 && l.Name != name {
			summary := model.ChannelDailySummary{
				LogDate:          logTime,
				Name:             name,
				NumLogs:          n,
				MaxListeners:     maxListeners,
				AverageListeners: average(n, sumListeners),
			}
			summaryRepo.Create(&summary)
		}
		name = l.Name
		n++
		sumListeners += l.Listeners
		if l.Listeners > maxListeners {
			maxListeners = l.Listeners
		}
	}
	if len(name) > 0 {
		summary := model.ChannelDailySummary{
			LogDate:          logTime,
			Name:             name,
			NumLogs:          n,
			MaxListeners:     maxListeners,
			AverageListeners: average(n, sumListeners),
		}
		summaryRepo.Create(&summary)
	}
}

func average(n, sum int) float32 {
	if n == 0 {
		return 0
	}
	return float32(sum) / float32(n)
}
