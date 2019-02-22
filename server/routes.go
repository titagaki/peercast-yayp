package server

import (
	"github.com/labstack/echo"

	"peercast-yayp/handler"
)

func Routes(e *echo.Echo) {
	e.GET("/index.txt", handler.IndexTxt())
	e.GET("/api/channels", handler.GetChannels())
	e.GET("/api/channelLogs", handler.GetChannelLogs())
	e.GET("/api/channelDailyLogs", handler.GetChannelLogs())
	e.Static("/*", "public")
}
