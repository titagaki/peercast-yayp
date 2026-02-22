package store

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"peercast-yayp/internal/domain"
)

// ChannelLogStore はチャンネルログの永続化を担う。
type ChannelLogStore struct {
	db *gorm.DB
}

func NewChannelLogStore(db *gorm.DB) *ChannelLogStore {
	return &ChannelLogStore{db: db}
}

// CreateFromChannels はチャンネル一覧からログレコードを一括作成する。
func (s *ChannelLogStore) CreateFromChannels(ctx context.Context, logTime time.Time, channels []*domain.Channel) error {
	for _, ch := range channels {
		log := domain.ChannelLog{
			LogTime:         logTime,
			ChannelID:       ch.ID,
			CID:             ch.CID,
			Name:            ch.Name,
			Bitrate:         ch.Bitrate,
			ContentType:     ch.ContentType,
			Listeners:       ch.Listeners,
			Relays:          ch.Relays,
			Age:             ch.Age,
			Genre:           ch.Genre,
			Description:     ch.Description,
			URL:             ch.URL,
			Comment:         ch.Comment,
			TrackArtist:     ch.TrackArtist,
			TrackTitle:      ch.TrackTitle,
			TrackAlbum:      ch.TrackAlbum,
			TrackGenre:      ch.TrackGenre,
			TrackContact:    ch.TrackContact,
			HiddenListeners: ch.HiddenListeners,
		}
		if err := s.db.WithContext(ctx).Create(&log).Error; err != nil {
			return fmt.Errorf("create channel log for %q: %w", ch.Name, err)
		}
	}
	return nil
}

// FindByNameAndDate は指定チャンネル名・日付のログを返す。
func (s *ChannelLogStore) FindByNameAndDate(ctx context.Context, name string, date time.Time) ([]*domain.ChannelLog, error) {
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
	end := start.Add(24 * time.Hour)

	var logs []*domain.ChannelLog
	err := s.db.WithContext(ctx).
		Where("name = ? AND log_time >= ? AND log_time < ?", name, start, end).
		Find(&logs).Error
	if err != nil {
		return nil, fmt.Errorf("find channel logs by name and date: %w", err)
	}
	return logs, nil
}

// FindByDate は指定日のすべてのログを名前・時刻順で返す。
func (s *ChannelLogStore) FindByDate(ctx context.Context, date time.Time) ([]*domain.ChannelLog, error) {
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
	end := start.Add(24 * time.Hour)

	var logs []*domain.ChannelLog
	err := s.db.WithContext(ctx).
		Where("log_time >= ? AND log_time < ?", start, end).
		Order("name ASC, log_time ASC").
		Find(&logs).Error
	if err != nil {
		return nil, fmt.Errorf("find channel logs by date: %w", err)
	}
	return logs, nil
}
