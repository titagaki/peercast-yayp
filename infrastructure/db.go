package infrastructure

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"

	"peercast-yayp/config"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local&timeout=15s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DB,
	)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect database")
	}

	if cfg.Server.Debug {
		db.LogMode(true)
	}

	return db, nil
}
