package worker

import (
	"context"
	"log/slog"
	"time"

	"peercast-yayp/internal/domain"
	"peercast-yayp/internal/peercast"
)

// ChannelRepository はチャンネルの永続化インターフェース。
type ChannelRepository interface {
	FindPlaying(ctx context.Context) ([]*domain.Channel, error)
	Upsert(ctx context.Context, ch *domain.Channel) error
	Save(ctx context.Context, ch *domain.Channel) error
}

// ChannelLogRepository はチャンネルログの永続化インターフェース。
type ChannelLogRepository interface {
	CreateFromChannels(ctx context.Context, logTime time.Time, channels []*domain.Channel) error
	FindByDate(ctx context.Context, date time.Time) ([]*domain.ChannelLog, error)
}

// SummaryRepository は日別集計の永続化インターフェース。
type SummaryRepository interface {
	Create(ctx context.Context, summary *domain.ChannelDailySummary) error
}

// ChannelCacheWriter はチャンネルキャッシュの書き込みインターフェース。
type ChannelCacheWriter interface {
	Set(channels []*domain.Channel)
}

// PeercastFetcher はPeerCastからのデータ取得インターフェース。
type PeercastFetcher interface {
	FetchChannels(ctx context.Context) (*peercast.StatXML, error)
}

// Dependencies はワーカーの全依存関係を保持する。
type Dependencies struct {
	Logger       *slog.Logger
	Channels     ChannelRepository
	ChannelLogs  ChannelLogRepository
	Summaries    SummaryRepository
	ChannelCache ChannelCacheWriter
	Peercast     PeercastFetcher
	YPPrefix     string
}

// Worker はバックグラウンドジョブを実行する。
type Worker struct {
	deps        Dependencies
	lastLogTime time.Time
}

// New は依存関係を注入して新しいWorkerを生成する。
func New(deps Dependencies) *Worker {
	return &Worker{deps: deps}
}

// Start はワーカーのメインループを開始する。コンテキストがキャンセルされるまでブロックする。
func (w *Worker) Start(ctx context.Context) {
	syncTicker := time.NewTicker(30 * time.Second)
	defer syncTicker.Stop()

	// 起動直後に一度同期を実行
	w.syncChannels(ctx)

	// 日次集計のスケジュール
	go w.runDailySummarySchedule(ctx)

	for {
		select {
		case <-ctx.Done():
			return
		case <-syncTicker.C:
			w.syncChannels(ctx)
		}
	}
}

func (w *Worker) runDailySummarySchedule(ctx context.Context) {
	for {
		now := time.Now()
		// 次の 0:03 を計算
		next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 3, 0, 0, time.Local)
		if now.Hour() == 0 && now.Minute() < 3 {
			next = time.Date(now.Year(), now.Month(), now.Day(), 0, 3, 0, 0, time.Local)
		}

		timer := time.NewTimer(time.Until(next))
		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-timer.C:
			w.dailySummary(ctx)
		}
	}
}
