package infrastructure

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"

	"peercast-yayp/config"
)

func NewDB() (*gorm.DB, error) {
	cfg := config.GetConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=15s",
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
		logW, err := NewLogger("db")
		if err != nil {
			return nil, err
		}
		db.SetLogger(logW)
	}

	return db, nil
}
