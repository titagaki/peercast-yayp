package models

import (
	"time"
)

type Channel struct {
	ID               uint
	GnuID            string    `gorm:"size:32"`
	Name             string    `gorm:"index"`
	Tip              string    `gorm:"size:53"`
	Bitrate          int
	ContentType      string
	Listeners        int
	Relays           int
	Age              int
	Genre            string
	Description      string
	Url              string
	Comment          string
	TrackArtist      string
	TrackTitle       string
	TrackAlbum       string
	TrackGenre       string
	TrackContact     string
	IsHostDirect     bool
	HiddenListeners  bool
	IsPlaying        bool      `gorm:"index"`
	IsBanned         bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

