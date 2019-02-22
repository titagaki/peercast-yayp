package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	gocache "github.com/patrickmn/go-cache"

	"peercast-yayp/config"
	"peercast-yayp/infrastructure"
	"peercast-yayp/repositoriy"
)

func GetChannels() echo.HandlerFunc {
	return func(c echo.Context) error {
		cache, ok := c.Get("cache").(*gocache.Cache)
		if !ok {
			return errors.New("can't available cache")
		}

		channels, ok := repositoriy.NewCachedChannelRepository(cache).GetChannels()
		if !ok {
			db, err := infrastructure.NewDB(config.GetConfig())
			if err != nil {
				return err
			}
			defer db.Close()

			channels = repositoriy.NewChannelRepository(db).FindPlayingChannels()
		}

		return c.JSON(http.StatusOK, channels)
	}
}
