package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"peercast-yayp/config"
	"peercast-yayp/infrastructure"
	"peercast-yayp/job"
	"peercast-yayp/server"
)

func initConfig() *config.Config {
	configPath := flag.String("config", "config/config.toml", "path of the config file")
	flag.Parse()

	cfg, err := config.FromFile(*configPath)
	if err != nil {
		panic(err)
	}
	return cfg
}

func getLogWriter(cfg *config.Config) io.Writer {
	log, err := os.OpenFile(cfg.Server.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return io.MultiWriter(log, os.Stdout)
}

func main() {
	cfg := initConfig()
	log := getLogWriter(cfg)

	cache := infrastructure.NewCache()

	go job.RunScheduler(cache)

	e := echo.New()

	e.Logger.SetOutput(log)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: log,
	}))
	e.Use(middleware.Recover())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("cache", cache)
			return next(c)
		}
	})

	server.Routes(e)

	// Start server with GraceShutdown
	go func() {
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
}
