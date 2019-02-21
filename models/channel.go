package models

import (
	"time"
)

type Channel struct {
	ID              uint
	GnuID           string `gorm:"size:32"`
	Name            string `gorm:"index"`
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
	TrackerIP       string `gorm:"size:53"`
	TrackerDirect   bool
	IsPlaying       bool `gorm:"index"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (db *DB) FindPlayingChannels() []*Channel {

	channel := make([]*Channel, 0)
	db.Where("is_playing = ?", true).Find(&channel)

	return channel
}
