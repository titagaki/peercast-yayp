package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"peercast-yayp/config"
	"peercast-yayp/infrastructure"
	"peercast-yayp/job"
	"peercast-yayp/server"
)

func initConfig() (*config.Config, error) {
	configPath := flag.String("config", "./yayp.toml", "path of the config file")
	flag.Parse()

	cfg, err := config.FromFile(*configPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config")
	}
	return cfg, nil
}

func main() {
	if _, err := initConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cache := infrastructure.NewCache()

	go job.RunScheduler(cache)

	if err := server.Start(cache); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
