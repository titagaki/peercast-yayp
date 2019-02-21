package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"peercast-yayp/config"
	"peercast-yayp/models"
)

func GetChannels() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := models.NewDB(config.GetConfig())
		if err != nil {
			return err
		}

		channels := db.FindPlayingChannels()

		return c.JSON(http.StatusOK, channels)
	}
}
