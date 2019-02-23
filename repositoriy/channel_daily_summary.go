package repositoriy

import (
	"github.com/jinzhu/gorm"

	"peercast-yayp/model"
)

type ChannelDailySummaryRepository struct {
	*gorm.DB
}

func NewChannelDailySummaryRepository(db *gorm.DB) *ChannelDailySummaryRepository {
	return &ChannelDailySummaryRepository{db}
}

func (r *ChannelDailySummaryRepository) Create(summary *model.ChannelDailySummary) {
	r.DB.Create(summary)
}

func (r *ChannelDailySummaryRepository) FindChannelDailySummaryByName(name string) []*model.ChannelDailySummary {
	summaries := make([]*model.ChannelDailySummary, 0)
	r.DB.Where("name = ?", name).Find(&summaries)
	return summaries
}
