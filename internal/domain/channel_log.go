package domain

import "time"

// ChannelLog はチャンネルの定期ログ（10分ごと）を表すエンティティ。
type ChannelLog struct {
	ID              uint      `gorm:"primaryKey" json:"ID"`
	LogTime         time.Time `gorm:"type:datetime;uniqueIndex:uix_channel_logs_log_time_name" json:"LogTime"`
	ChannelID       uint      `gorm:"index" json:"ChannelID"`
	CID             string    `gorm:"column:cid;size:32" json:"CID"`
	Name            string    `gorm:"uniqueIndex:uix_channel_logs_log_time_name" json:"Name"`
	Bitrate         int       `json:"Bitrate"`
	ContentType     string    `json:"ContentType"`
	Listeners       int       `json:"Listeners"`
	Relays          int       `json:"Relays"`
	Age             uint      `json:"Age"`
	Genre           string    `json:"Genre"`
	Description     string    `json:"Description"`
	URL             string    `gorm:"column:url" json:"Url"`
	Comment         string    `json:"Comment"`
	TrackArtist     string    `json:"TrackArtist"`
	TrackTitle      string    `json:"TrackTitle"`
	TrackAlbum      string    `json:"TrackAlbum"`
	TrackGenre      string    `json:"TrackGenre"`
	TrackContact    string    `json:"TrackContact"`
	HiddenListeners bool      `gorm:"column:hidden_listeners" json:"-"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

func (ChannelLog) TableName() string {
	return "channel_logs"
}
