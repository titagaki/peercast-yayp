package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"peercast-yayp/handler"
)

func Routes(e *echo.Echo) {
	e.GET("/index.txt", handler.IndexTxt())
	e.GET("/api/channels", handler.GetChannels())
	e.GET("/api/channelLogs", handler.GetChannelLogs())
	e.GET("/api/channelDailyLogs", handler.GetChannelLogs())

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "public",
		HTML5: true,
	}))
}
