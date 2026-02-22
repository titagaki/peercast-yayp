package handler

import (
	"fmt"
	"html"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"

	"peercast-yayp/internal/domain"
)

// IndexTxt はPeerCast互換の index.txt 形式でチャンネル情報を返すハンドラー。
func (h *Handler) IndexTxt(c echo.Context) error {
	ctx := c.Request().Context()

	channels, chOK := h.deps.ChannelCache.Get()
	info, infoOK := h.deps.InfoCache.Get()

	if !chOK {
		var err error
		channels, err = h.deps.Channels.FindPlaying(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch channels")
		}
		h.deps.ChannelCache.Set(channels)
	}

	if !infoOK {
		var err error
		info, err = h.deps.Information.FindAll(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch information")
		}
		h.deps.InfoCache.Set(info)
	}

	channels = domain.MaskListeners(channels)

	var buf []byte
	for _, ch := range channels {
		buf = appendChannelLine(buf, ch)
	}
	for _, i := range info {
		buf = appendInformationLine(buf, i)
	}

	return c.String(http.StatusOK, string(buf))
}

func appendChannelLine(buf []byte, ch *domain.Channel) []byte {
	buf = append(buf, html.EscapeString(ch.Name)...)
	buf = append(buf, "<>"...)
	buf = append(buf, html.EscapeString(ch.CID)...)
	buf = append(buf, "<>"...)
	buf = append(buf, html.EscapeString(ch.TrackerIP)...)
	buf = append(buf, "<>"...)
	buf = append(buf, html.EscapeString(ch.URL)...)
	buf = append(buf, "<>"...)
	buf = append(buf, html.EscapeString(ch.Genre)...)
	buf = append(buf, "<>"...)
	buf = append(buf, html.EscapeString(ch.Description)...)
	buf = append(buf, "<>"...)
	buf = append(buf, strconv.Itoa(ch.Listeners)...)
	buf = append(buf, "<>"...)
	buf = append(buf, strconv.Itoa(ch.Relays)...)
	buf = append(buf, "<>"...)
	buf = append(buf, strconv.Itoa(ch.Bitrate)...)
	buf = append(buf, "<>"...)
	buf = append(buf, html.EscapeString(ch.ContentType)...)
	buf = append(buf, "<>"...)
	buf = append(buf, html.EscapeString(ch.TrackArtist)...)
	buf = append(buf, "<>"...)
	buf = append(buf, html.EscapeString(ch.TrackAlbum)...)
	buf = append(buf, "<>"...)
	buf = append(buf, html.EscapeString(ch.TrackTitle)...)
	buf = append(buf, "<>"...)
	buf = append(buf, html.EscapeString(ch.TrackContact)...)
	buf = append(buf, "<>"...)
	buf = append(buf, url.QueryEscape(ch.Name)...)
	buf = append(buf, "<>"...)
	buf = append(buf, formatDuration(ch.Age)...)
	buf = append(buf, "<>"...)
	buf = append(buf, "click"...)
	buf = append(buf, "<>"...)
	buf = append(buf, html.EscapeString(ch.Comment)...)
	buf = append(buf, "<>"...)
	buf = append(buf, boolToDigit(ch.TrackerDirect)...)
	buf = append(buf, '\n')
	return buf
}

func appendInformationLine(buf []byte, info *domain.Information) []byte {
	buf = append(buf, html.EscapeString(info.Name)...)
	buf = append(buf, "<>00000000000000000000000000000000<><><><>"...)
	buf = append(buf, html.EscapeString(info.Description)...)
	buf = append(buf, "<>"...)
	buf = append(buf, strconv.Itoa(info.Priority)...)
	buf = append(buf, "<>"...)
	buf = append(buf, strconv.Itoa(info.Priority)...)
	buf = append(buf, "<>0<>RAW<><><><><><>00:00<>click<><>0\n"...)
	return buf
}

func boolToDigit(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func formatDuration(seconds uint) string {
	s := seconds % 60
	total := (seconds - s) / 60
	m := total % 60
	h := (total - m) / 60
	return fmt.Sprintf("%d:%02d", h, m)
}
