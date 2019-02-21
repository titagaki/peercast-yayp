package job

import (
	"fmt"
	"time"

	"peercast-yayp/config"
	"peercast-yayp/models"
	"peercast-yayp/peercast"
)

func SyncChannel() {
	t := time.Now()
	fmt.Println("Time's up! @", t.UTC())

	db, err := models.NewDB(config.GetConfig())
	if err != nil {
		panic(err)

		// ToDo: log
		return
	}
	channels := db.FindPlayingChannels()

	data, err := peercast.GetStatXML()
	if err != nil {
		panic(err)
		// ToDo: log
		return
	}

	// チャンネルIDをキーとしたchannelsのmapを作成
	channelsMap := makeChannelsMap(channels)

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

		channel, ok := channelsMap[v.ID]
		if ok {
			delete(channelsMap, v.ID)
		} else {
			channel = new(models.Channel)
		}

		channel.GnuID = v.ID
		channel.Name = v.Name
		channel.Bitrate = v.Bitrate
		channel.ContentType = v.Type
		channel.Listeners = v.ChanHitStat.Listeners
		channel.Relays = v.ChanHitStat.Relays
		channel.Age = v.Age
		channel.Genre = option.Genre
		channel.Description = v.Desc
		channel.Url = v.Url
		channel.Comment = v.Comment
		channel.TrackArtist = v.ChannelTrack.Artist
		channel.TrackTitle = v.ChannelTrack.Title
		channel.TrackAlbum = v.ChannelTrack.Album
		channel.TrackGenre = v.ChannelTrack.Genre
		channel.TrackContact = v.ChannelTrack.Contact
		channel.HiddenListeners = option.HiddenListeners
		channel.TrackerIP = tracker.IP
		channel.TrackerDirect = tracker.Direct
		channel.IsPlaying = true

		if db.NewRecord(channel) {
			db.Create(channel)
		} else {
			db.Save(channel)
		}
	}

	for _, c := range channelsMap {
		c.IsPlaying = false
		db.Save(c)
	}

	//fmt.Printf("%+v", channels)
}

func makeChannelsMap(channels []*models.Channel) map[string]*models.Channel {
	cMap := map[string]*models.Channel{}

	for _, c := range channels {
		cMap[c.GnuID] = c
	}
	return cMap
}
