package config

import (
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func NewSQLITE() (*lib.Database, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &lib.Database{DB: db}, nil
}
