package handler

import (
	"context"
	"time"

	"peercast-yayp/internal/domain"
)

// ChannelReader はチャンネルの読み取りインターフェース。
type ChannelReader interface {
	FindPlaying(ctx context.Context) ([]*domain.Channel, error)
}

// ChannelLogReader はチャンネルログの読み取りインターフェース。
type ChannelLogReader interface {
	FindByNameAndDate(ctx context.Context, name string, date time.Time) ([]*domain.ChannelLog, error)
}

// InformationReader はお知らせ情報の読み取りインターフェース。
type InformationReader interface {
	FindAll(ctx context.Context) ([]*domain.Information, error)
}

// ChannelCacheAccessor はチャンネルキャッシュの読み書きインターフェース。
type ChannelCacheAccessor interface {
	Get() ([]*domain.Channel, bool)
	Set(channels []*domain.Channel)
}

// InformationCacheAccessor はお知らせキャッシュの読み書きインターフェース。
type InformationCacheAccessor interface {
	Get() ([]*domain.Information, bool)
	Set(info []*domain.Information)
}

// Dependencies はハンドラーの全依存関係を保持する。
type Dependencies struct {
	Channels     ChannelReader
	ChannelLogs  ChannelLogReader
	Information  InformationReader
	ChannelCache ChannelCacheAccessor
	InfoCache    InformationCacheAccessor
}

// Handler はHTTPハンドラーメソッドを提供する。
type Handler struct {
	deps Dependencies
}

// New は依存関係を注入して新しいHandlerを生成する。
func New(deps Dependencies) *Handler {
	return &Handler{deps: deps}
}
