package worker

import (
	"context"
	"time"

	"peercast-yayp/internal/domain"
)

func (w *Worker) dailySummary(ctx context.Context) {
	w.deps.Logger.Info("daily summary started")

	yesterday := time.Now().Add(-24 * time.Hour)

	logs, err := w.deps.ChannelLogs.FindByDate(ctx, yesterday)
	if err != nil {
		w.deps.Logger.Error("failed to find channel logs for summary", "error", err)
		return
	}

	summaries := aggregateLogs(yesterday, logs)

	for _, s := range summaries {
		if err := w.deps.Summaries.Create(ctx, s); err != nil {
			w.deps.Logger.Error("failed to create daily summary", "error", err, "name", s.Name)
		}
	}
}

// aggregateLogs はログをチャンネル名ごとに集計して日別サマリーを生成する。
func aggregateLogs(date time.Time, logs []*domain.ChannelLog) []*domain.ChannelDailySummary {
	type accumulator struct {
		count        int
		sumListeners int
		maxListeners int
	}

	grouped := make(map[string]*accumulator)
	var order []string

	for _, l := range logs {
		acc, exists := grouped[l.Name]
		if !exists {
			acc = &accumulator{}
			grouped[l.Name] = acc
			order = append(order, l.Name)
		}
		acc.count++
		acc.sumListeners += l.Listeners
		if l.Listeners > acc.maxListeners {
			acc.maxListeners = l.Listeners
		}
	}

	summaries := make([]*domain.ChannelDailySummary, 0, len(grouped))
	for _, name := range order {
		acc := grouped[name]
		summaries = append(summaries, &domain.ChannelDailySummary{
			LogDate:          date,
			Name:             name,
			NumLogs:          acc.count,
			MaxListeners:     acc.maxListeners,
			AverageListeners: float64(acc.sumListeners) / float64(acc.count),
		})
	}

	return summaries
}
