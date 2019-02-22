package model

import (
	"time"
)

type ChannelDailyLog struct {
	ID      uint
	LogDate time.Time `gorm:"type:date;unique_index:uix_channel_daily_logs_log_date_name"`
	Name    string    `gorm:"size:100;unique_index:uix_channel_daily_logs_log_date_name"`
}
