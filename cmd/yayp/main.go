package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"peercast-yayp/internal/cache"
	"peercast-yayp/internal/config"
	"peercast-yayp/internal/handler"
	"peercast-yayp/internal/peercast"
	"peercast-yayp/internal/server"
	"peercast-yayp/internal/store"
	"peercast-yayp/internal/worker"
)

func main() {
	configPath := flag.String("config", "./yayp.toml", "path of the config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	logger := newLogger(cfg)

	db, err := store.NewDB(cfg.Database, cfg.Server.Debug, logger)
	if err != nil {
		logger.Error("failed to connect database", "error", err)
		os.Exit(1)
	}

	// Stores
	channelStore := store.NewChannelStore(db)
	channelLogStore := store.NewChannelLogStore(db)
	summaryStore := store.NewSummaryStore(db)
	infoStore := store.NewInformationStore(db)

	// Cache
	appCache := cache.New()
	channelCache := cache.NewChannelCache(appCache)
	infoCache := cache.NewInformationCache(appCache)

	// PeerCast client
	pcClient := peercast.NewClient(cfg.Peercast)

	// Worker (background jobs)
	w := worker.New(worker.Dependencies{
		Logger:       logger,
		Channels:     channelStore,
		ChannelLogs:  channelLogStore,
		Summaries:    summaryStore,
		ChannelCache: channelCache,
		Peercast:     pcClient,
		YPPrefix:     cfg.Server.YPPrefix,
	})

	// Handler (HTTP)
	h := handler.New(handler.Dependencies{
		Channels:     channelStore,
		ChannelLogs:  channelLogStore,
		Information:  infoStore,
		ChannelCache: channelCache,
		InfoCache:    infoCache,
	})

	// Start worker
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go w.Start(ctx)

	// Start server
	srv := server.New(cfg.Server, logger, h)
	go func() {
		if err := srv.Start(); err != nil {
			logger.Info("server stopped", "error", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down...")
	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("server forced to shutdown", "error", err)
	}
}

func newLogger(cfg *config.Config) *slog.Logger {
	opts := &slog.HandlerOptions{}
	if cfg.Server.Debug {
		opts.Level = slog.LevelDebug
	}

	writers := []io.Writer{os.Stdout}
	if cfg.Server.LogPath != "" {
		f, err := os.OpenFile(cfg.Server.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			writers = append(writers, f)
		}
	}

	w := io.MultiWriter(writers...)
	return slog.New(slog.NewTextHandler(w, opts))
}
