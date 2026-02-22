package store

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"peercast-yayp/internal/domain"
)

// ChannelStore はチャンネルの永続化を担う。
type ChannelStore struct {
	db *gorm.DB
}

func NewChannelStore(db *gorm.DB) *ChannelStore {
	return &ChannelStore{db: db}
}

// FindPlaying は現在配信中のチャンネル一覧を返す。
func (s *ChannelStore) FindPlaying(ctx context.Context) ([]*domain.Channel, error) {
	var channels []*domain.Channel
	if err := s.db.WithContext(ctx).Where("is_playing = ?", true).Find(&channels).Error; err != nil {
		return nil, fmt.Errorf("find playing channels: %w", err)
	}
	return channels, nil
}

// Save は既存のチャンネルを更新する。
func (s *ChannelStore) Save(ctx context.Context, ch *domain.Channel) error {
	if err := s.db.WithContext(ctx).Save(ch).Error; err != nil {
		return fmt.Errorf("save channel: %w", err)
	}
	return nil
}

// Create は新しいチャンネルを作成する。
func (s *ChannelStore) Create(ctx context.Context, ch *domain.Channel) error {
	if err := s.db.WithContext(ctx).Create(ch).Error; err != nil {
		return fmt.Errorf("create channel: %w", err)
	}
	return nil
}

// Upsert はチャンネルが新規の場合は作成、既存の場合は更新する。
func (s *ChannelStore) Upsert(ctx context.Context, ch *domain.Channel) error {
	if ch.ID == 0 {
		return s.Create(ctx, ch)
	}
	return s.Save(ctx, ch)
}
