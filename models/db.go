package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type DB struct {
	*gorm.DB
}

func NewDB(dsn string) (*DB, error) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect database")
	}
	defer db.Close()

	return &DB{db}, nil
}
