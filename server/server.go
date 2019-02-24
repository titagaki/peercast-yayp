package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/patrickmn/go-cache"

	"peercast-yayp/config"
	"peercast-yayp/infrastructure"
)

func Start(cache *cache.Cache) error {
	e := echo.New()

	logW, err := infrastructure.NewLogWriter()
	if err != nil {
		return err
	}

	e.Logger.SetOutput(logW)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logW,
	}))
	e.Use(middleware.Recover())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("cache", cache)
			return next(c)
		}
	})

	Routes(e)

	// Start server with GraceShutdown
	go func() {
		cfg := config.GetConfig()
		if err := e.Start(fmt.Sprintf(":%s", cfg.Server.Port)); err != nil {
			e.Logger.Info("shutting down the server.")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	return nil
}
