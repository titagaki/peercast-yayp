package model

import (
	"time"
)

type Channel struct {
	ID              uint
	CID             string `gorm:"column:cid;size:32"`
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
	IsPlaying       bool      `gorm:"index" json:"-"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

type ChannelList []*Channel

func (list ChannelList) HideListeners() ChannelList {
	var newList ChannelList
	for _, c := range list {
		if c.HiddenListeners {
			tmp := *c
			tmp.Listeners = -1
			tmp.Relays = -1
			c = &tmp
		}
		newList = append(newList, c)
	}
	return newList
}
