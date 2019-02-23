package model

import (
	"time"
)

type ChannelDailySummary struct {
	ID               uint
	LogDate          time.Time `gorm:"type:date;unique_index:uix_channel_daily_logs_log_date_name"`
	Name             string    `gorm:"unique_index:uix_channel_daily_logs_log_date_name"`
	NumLogs          int
	MaxListeners     int
	AverageListeners float32
}
