package job

import (
	"fmt"
	"time"

	"peercast-yayp/models"
	"peercast-yayp/peercast"
)

func SyncChannel() {
	t := time.Now()
	fmt.Println("Time's up! @", t.UTC())

	data, err := peercast.GetStatXML()
	if err != nil {
		// ToDo: log
		return
	}

	var channels []models.Channel

	for _, v := range data.ChannelsFound.Channel {

		// ジャンル書式をパース
		option, ok := peercast.ParseGenre(v.Genre)
		if !ok {
			continue
		}

		// trackerのホストを決定
		var tracker *peercast.ChanHit
		for _, host := range v.ChanHitStat.ChanHit {
			if host.Push == true {
				continue
			}
			if host.Tracker == true {
				tracker = &host
				break
			} else if tracker == nil {
				tracker = &host
			}
		}
		if tracker == nil {
			continue
		}

		channel := models.Channel{
			GnuID : v.ID,
			Name: v.Name,
			Tip: tracker.IP,
			Bitrate: v.Bitrate,
			ContentType: v.Type,
			Listeners: v.ChanHitStat.Listeners,
			Relays: v.ChanHitStat.Relays,
			Age: v.Age,
			Genre: option.Genre,
			Description: v.Desc,
			Url: v.Url,
			Comment: v.Comment,
			TrackArtist: v.ChannelTrack.Artist,
			TrackTitle: v.ChannelTrack.Title,
			TrackAlbum: v.ChannelTrack.Album,
			TrackGenre:  v.ChannelTrack.Genre,
			TrackContact: v.ChannelTrack.Contact,
			IsHostDirect: tracker.Direct,
			HiddenListeners: option.HiddenListeners,
			IsPlaying: true,
		}

		channels = append(channels, channel)
	}

	fmt.Printf("%+v", channels)

}