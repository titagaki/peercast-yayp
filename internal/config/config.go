package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Config はアプリケーション全体の設定を保持する。
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Peercast PeercastConfig
}

// ServerConfig はHTTPサーバーの設定。
type ServerConfig struct {
	Port     string
	YPPrefix string
	LogPath  string
	Debug    bool
}

// DatabaseConfig はデータベース接続の設定。
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

// PeercastConfig はPeerCast接続の設定。
type PeercastConfig struct {
	Host         string
	Port         string
	AuthType     string
	AuthUser     string
	AuthPassword string
}

// Load は指定されたTOMLファイルから設定を読み込む。
func Load(path string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return nil, fmt.Errorf("decode config: %w", err)
	}
	return &cfg, nil
}
