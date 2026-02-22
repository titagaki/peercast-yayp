package store

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"peercast-yayp/internal/domain"
)

// SummaryStore はチャンネル日別集計の永続化を担う。
type SummaryStore struct {
	db *gorm.DB
}

func NewSummaryStore(db *gorm.DB) *SummaryStore {
	return &SummaryStore{db: db}
}

// Create は新しい日別集計レコードを作成する。
func (s *SummaryStore) Create(ctx context.Context, summary *domain.ChannelDailySummary) error {
	if err := s.db.WithContext(ctx).Create(summary).Error; err != nil {
		return fmt.Errorf("create daily summary: %w", err)
	}
	return nil
}
