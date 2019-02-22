package handler

import (
	"fmt"
	"html"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo"
	gocache "github.com/patrickmn/go-cache"

	"peercast-yayp/config"
	"peercast-yayp/infrastructure"
	"peercast-yayp/model"
	"peercast-yayp/repositoriy"
)

func IndexTxt() echo.HandlerFunc {
	return func(c echo.Context) error {

		channels, ok := getCachedCannels(c)
		if !ok {
			db, err := infrastructure.NewDB(config.GetConfig())
			if err != nil {
				return err
			}

			channelRepo := repositoriy.NewChannelRepository(db)
			channels = channelRepo.FindPlayingChannels()
		}

		s := make([]byte, 0, 100)

		for _, c := range channels {
			listeners := c.Listeners
			relays := c.Relays
			if c.HiddenListeners {
				listeners = -1
				relays = -1
			}
			s = append(s, html.EscapeString(c.Name)...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(c.GnuID)...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(c.TrackerIP)...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(c.Url)...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(c.Genre)...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(c.Description)...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(strconv.Itoa(listeners))...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(strconv.Itoa(relays))...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(strconv.Itoa(c.Bitrate))...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(c.ContentType)...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(c.TrackArtist)...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(c.TrackAlbum)...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(c.TrackTitle)...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(c.TrackContact)...)
			s = append(s, "<>"...)
			s = append(s, url.QueryEscape(c.Name)...)
			s = append(s, "<>"...)
			s = append(s, formatTime(c.Age)...)
			s = append(s, "<>"...)
			s = append(s, "click"...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(c.Comment)...)
			s = append(s, "<>"...)
			s = append(s, html.EscapeString(btos(c.TrackerDirect))...)
			s = append(s, "\n"...)
		}
		return c.String(http.StatusOK, string(s))
	}
}

func getCachedCannels(c echo.Context) (model.ChannelList, bool) {
	cache, ok := c.Get("cache").(*gocache.Cache)
	if !ok {
		return nil, false
	}

	channels, ok := cache.Get("ChannelList")
	if !ok {
		return nil, false
	}

	return channels.(model.ChannelList), true
}

func btos(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func formatTime(t uint) string {
	s := t % 60
	t = (t - s) / 60
	m := t % 60
	h := (t - m) / 60
	return fmt.Sprintf("%d:%02d", h, m)
}
