package store

import (
	"fmt"
	"log/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"peercast-yayp/internal/config"
)

// NewDB はMySQL接続を確立して *gorm.DB を返す。
func NewDB(cfg config.DatabaseConfig, debug bool, log *slog.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=15s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB,
	)

	gormCfg := &gorm.Config{}
	if debug {
		gormCfg.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.Open(dsn), gormCfg)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	return db, nil
}
