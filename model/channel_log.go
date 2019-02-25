package model

import (
	"time"
)

type ChannelLog struct {
	ID              uint
	LogTime         time.Time `gorm:"type:datetime;unique_index:uix_channel_logs_log_time_name"`
	ChannelID       uint      `gorm:"index"`
	CID             string    `gorm:"column:cid;size:32"`
	Name            string    `gorm:"unique_index:uix_channel_logs_log_time_name"`
	Bitrate         int
	ContentType     string
	Listeners       int
	Relays          int
	Age             uint
	Genre           string
	Description     string
	Url             string
	Comment         string
	TrackArtist     string
	TrackTitle      string
	TrackAlbum      string
	TrackGenre      string
	TrackContact    string
	HiddenListeners bool
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}
