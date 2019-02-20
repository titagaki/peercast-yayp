package handler

import (
	"net/http"
	"peercast-yayp/models"

	"github.com/labstack/echo"
)

func GetChannels() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, _ := models.NewDB("root:password@/yayp?charset=utf8&parseTime=True&loc=Local")
		channels := db.FindPlayingChannels()

		return c.JSON(http.StatusOK, channels)
	}
}
