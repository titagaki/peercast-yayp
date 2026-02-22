package domain

import "time"

// Channel はPeerCastのチャンネル情報を表すエンティティ。
type Channel struct {
	ID              uint      `gorm:"primaryKey" json:"ID"`
	CID             string    `gorm:"column:cid;size:32" json:"CID"`
	Name            string    `gorm:"index" json:"Name"`
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
	TrackerIP       string    `gorm:"column:tracker_ip;size:53" json:"-"`
	TrackerDirect   bool      `gorm:"column:tracker_direct" json:"-"`
	IsPlaying       bool      `gorm:"index" json:"-"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

func (Channel) TableName() string {
	return "channels"
}

// MaskListeners はリスナー非表示設定のチャンネルについて、
// Listeners/Relays を -1 に置き換えたコピーを返す。
func MaskListeners(channels []*Channel) []*Channel {
	result := make([]*Channel, len(channels))
	for i, ch := range channels {
		if ch.HiddenListeners {
			c := *ch
			c.Listeners = -1
			c.Relays = -1
			result[i] = &c
		} else {
			result[i] = ch
		}
	}
	return result
}
