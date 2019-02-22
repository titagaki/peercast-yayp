package model

import (
	"time"
)

type Channel struct {
	ID              uint
	GnuID           string `gorm:"size:32"`
	Name            string `gorm:"size:100;index"`
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

type ChannelList []*Channel
