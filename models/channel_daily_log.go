package models

import (
	"time"
)

type ChannelDailyLog struct {
	ID      uint
	LogDate time.Time `gorm:"type:date;unique_index:uix_channel_daily_logs_log_date_name"`
	Name    string    `gorm:"unique_index:uix_channel_daily_logs_log_date_name"`
}

func (db *DB) FindChannelDailyLogByName(name string) []*ChannelDailyLog {

	logs := make([]*ChannelDailyLog, 0)
	db.Where("name = ?", name).Find(&logs)

	return logs
}
