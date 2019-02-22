package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"

	"peercast-yayp/config"
)

func NewDB(conf *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.DB,
	)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect database")
	}

	// ログ出力
	db.LogMode(true)

	return db, nil
}
