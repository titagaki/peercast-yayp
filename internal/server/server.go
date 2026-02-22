package server

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"peercast-yayp/internal/config"
	"peercast-yayp/internal/handler"
)

// Server はHTTPサーバーの起動・停止を管理する。
type Server struct {
	echo   *echo.Echo
	port   string
	logger *slog.Logger
}

// New は新しいHTTPサーバーを構築する。
func New(cfg config.ServerConfig, logger *slog.Logger, h *handler.Handler) *Server {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				"method", v.Method,
				"uri", v.URI,
				"status", v.Status,
				"latency", v.Latency,
			)
			return nil
		},
	}))

	// API routes
	e.GET("/index.txt", h.IndexTxt)
	e.GET("/api/channels", h.GetChannels)
	e.GET("/api/channelLogs", h.GetChannelLogs)
	e.GET("/api/channelDailyLogs", h.GetChannelLogs)

	// Static files (SPA fallback)
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "public",
		HTML5: true,
	}))

	return &Server{
		echo:   e,
		port:   cfg.Port,
		logger: logger,
	}
}

// Start はHTTPサーバーを起動する。ブロッキング呼び出し。
func (s *Server) Start() error {
	return s.echo.Start(fmt.Sprintf(":%s", s.port))
}

// Shutdown はHTTPサーバーをGracefulに停止する。
func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}
