package postgres

import (
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDBInstance() (*gorm.DB, error) {
	var err error
	once.Do(func() {

		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable",
		}), &gorm.Config{})
	})
	return db, err
}
