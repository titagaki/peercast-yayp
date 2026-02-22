package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"peercast-yayp/internal/domain"
)

// GetChannels は現在配信中のチャンネル一覧を返すAPIハンドラー。
func (h *Handler) GetChannels(c echo.Context) error {
	channels, ok := h.deps.ChannelCache.Get()
	if !ok {
		var err error
		channels, err = h.deps.Channels.FindPlaying(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch channels")
		}
	}

	channels = domain.MaskListeners(channels)

	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
	return c.JSON(http.StatusOK, channels)
}
