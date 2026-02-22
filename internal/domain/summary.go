package domain

import "time"

// ChannelDailySummary はチャンネルの日別集計を表すエンティティ。
type ChannelDailySummary struct {
	ID               uint      `gorm:"primaryKey"`
	LogDate          time.Time `gorm:"type:date;uniqueIndex:uix_channel_daily_logs_log_date_name"`
	Name             string    `gorm:"uniqueIndex:uix_channel_daily_logs_log_date_name"`
	NumLogs          int
	MaxListeners     int
	AverageListeners float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (ChannelDailySummary) TableName() string {
	return "channel_daily_summaries"
}
