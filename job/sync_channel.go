package job

import (
	"fmt"
	"time"

	"peercast-yayp/config"
	"peercast-yayp/infrastructure"
	"peercast-yayp/model"
	"peercast-yayp/peercast"
	"peercast-yayp/repositoriy"

	gocache "github.com/patrickmn/go-cache"
)

func SyncChannel(cache *gocache.Cache) {
	t := time.Now()
	fmt.Println("Time's up! @", t.UTC())

	db, err := infrastructure.NewDB(config.GetConfig())
	if err != nil {
		panic(err)

		// ToDo: log
		return
	}
	defer db.Close()

	channelRepo := repositoriy.NewChannelRepository(db)
	channels := channelRepo.FindPlayingChannels()

	data, err := peercast.GetStatXML()
	if err != nil {
		panic(err)
		// ToDo: log
		return
	}

	// チャンネルIDをキーとしたchannelsのmapを作成
	channelsMap := makeChannelsMap(channels)

	var newChannels model.ChannelList

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
			channel = new(model.Channel)
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

		channelRepo.SaveOrCreate(channel)

		newChannels = append(newChannels, channel)
	}

	for _, c := range channelsMap {
		c.IsPlaying = false
		db.Save(c)
	}

	// キャッシュに保持する
	repositoriy.NewCachedChannelRepository(cache).SetChannels(newChannels)

	// 10分置きにログを追加
	logTime := time.Now().Truncate(10 * time.Minute)
	lastLogTime, ok := cache.Get("LastLogTime")
	if !ok || logTime != lastLogTime.(time.Time) {
		channelLogRepo := repositoriy.NewChannelLogRepository(db)
		channelLogRepo.CreateChannelLogs(logTime, newChannels)

		cache.Set("LastLogTime", logTime, gocache.DefaultExpiration)
	}
}

func makeChannelsMap(channels model.ChannelList) map[string]*model.Channel {
	cMap := map[string]*model.Channel{}

	for _, c := range channels {
		cMap[c.GnuID] = c
	}
	return cMap
}
