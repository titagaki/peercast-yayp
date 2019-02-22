package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"peercast-yayp/config"
	"peercast-yayp/infrastructure"
	"peercast-yayp/repositoriy"
)

func GetChannels() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := infrastructure.NewDB(config.GetConfig())
		if err != nil {
			return err
		}
		defer db.Close()

		channelRepo := repositoriy.NewChannelRepository(db)
		channels := channelRepo.FindPlayingChannels()

		return c.JSON(http.StatusOK, channels)
	}
}
