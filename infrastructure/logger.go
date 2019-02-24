package infrastructure

import (
	"github.com/pkg/errors"
	"io"
	"os"

	"github.com/labstack/gommon/log"

	"peercast-yayp/config"
)

func NewLogger(prefix string) (*log.Logger, error) {
	logger := log.New(prefix)

	logW, err := NewLogWriter()
	if err != nil {
		return nil, err
	}
	logger.SetOutput(logW)

	return logger, nil
}

func NewLogWriter() (io.Writer, error) {
	cfg := config.GetConfig()

	fp, err := os.OpenFile(cfg.Server.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup logger")
	}

	if cfg.Server.Debug {
		return io.MultiWriter(fp, os.Stdout), nil
	}

	return toIoWriter(fp), nil
}

func toIoWriter(w io.Writer) io.Writer {
	return w
}
