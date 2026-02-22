package store

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"peercast-yayp/internal/domain"
)

// InformationStore はお知らせ情報の永続化を担う。
type InformationStore struct {
	db *gorm.DB
}

func NewInformationStore(db *gorm.DB) *InformationStore {
	return &InformationStore{db: db}
}

// FindAll は全てのお知らせ情報を返す（論理削除済みを除く）。
func (s *InformationStore) FindAll(ctx context.Context) ([]*domain.Information, error) {
	var info []*domain.Information
	if err := s.db.WithContext(ctx).Find(&info).Error; err != nil {
		return nil, fmt.Errorf("find information: %w", err)
	}
	return info, nil
}
