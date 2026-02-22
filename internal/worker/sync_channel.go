package worker

import (
	"context"
	"time"

	"peercast-yayp/internal/domain"
	"peercast-yayp/internal/peercast"
)

func (w *Worker) syncChannels(ctx context.Context) {
	w.deps.Logger.Info("sync channels started")

	// 現在配信中のチャンネルをDBから取得
	channels, err := w.deps.Channels.FindPlaying(ctx)
	if err != nil {
		w.deps.Logger.Error("failed to find playing channels", "error", err)
		return
	}

	// PeerCastからチャンネル情報を取得
	data, err := w.deps.Peercast.FetchChannels(ctx)
	if err != nil {
		w.deps.Logger.Error("failed to fetch peercast data", "error", err)
		return
	}

	channelMap := makeChannelMap(channels)
	var activeChannels []*domain.Channel

	for _, xmlCh := range data.ChannelsFound.Channel {
		opts, ok := peercast.ParseGenre(xmlCh.Genre, w.deps.YPPrefix)
		if !ok {
			continue
		}

		tracker := findTracker(xmlCh.Hits.Host)
		if tracker == nil {
			continue
		}

		ch, exists := channelMap[xmlCh.ID]
		if exists {
			delete(channelMap, xmlCh.ID)
		} else {
			ch = &domain.Channel{}
		}

		applyXMLToChannel(ch, xmlCh, opts, tracker)

		if err := w.deps.Channels.Upsert(ctx, ch); err != nil {
			w.deps.Logger.Error("failed to save channel", "error", err, "name", ch.Name)
			continue
		}

		activeChannels = append(activeChannels, ch)
	}

	// 配信終了したチャンネルを更新
	for _, ch := range channelMap {
		ch.IsPlaying = false
		if err := w.deps.Channels.Save(ctx, ch); err != nil {
			w.deps.Logger.Error("failed to update channel", "error", err, "name", ch.Name)
		}
	}

	// キャッシュを更新
	w.deps.ChannelCache.Set(activeChannels)

	// 10分間隔でログを記録
	logTime := time.Now().Truncate(10 * time.Minute)
	if logTime != w.lastLogTime {
		if err := w.deps.ChannelLogs.CreateFromChannels(ctx, logTime, activeChannels); err != nil {
			w.deps.Logger.Error("failed to create channel logs", "error", err)
		} else {
			w.lastLogTime = logTime
		}
	}
}

func makeChannelMap(channels []*domain.Channel) map[string]*domain.Channel {
	m := make(map[string]*domain.Channel, len(channels))
	for _, ch := range channels {
		m[ch.CID] = ch
	}
	return m
}

func findTracker(hosts []peercast.XMLHost) *peercast.XMLHost {
	var candidate *peercast.XMLHost
	for i := range hosts {
		if hosts[i].Push {
			continue
		}
		if hosts[i].Tracker {
			return &hosts[i]
		}
		if candidate == nil {
			candidate = &hosts[i]
		}
	}
	return candidate
}

func applyXMLToChannel(ch *domain.Channel, xml peercast.XMLChannel, opts peercast.StreamOptions, tracker *peercast.XMLHost) {
	ch.CID = xml.ID
	ch.Name = xml.Name
	ch.Bitrate = xml.Bitrate
	ch.ContentType = xml.Type
	ch.Listeners = xml.Hits.Listeners
	ch.Relays = xml.Hits.Relays
	ch.Age = xml.Age
	ch.Genre = opts.Genre
	ch.Description = xml.Desc
	ch.URL = xml.URL
	ch.Comment = xml.Comment
	ch.TrackArtist = xml.Track.Artist
	ch.TrackTitle = xml.Track.Title
	ch.TrackAlbum = xml.Track.Album
	ch.TrackGenre = xml.Track.Genre
	ch.TrackContact = xml.Track.Contact
	ch.HiddenListeners = opts.HiddenListeners
	ch.TrackerIP = tracker.IP
	ch.TrackerDirect = tracker.Direct
	ch.IsPlaying = true
}
