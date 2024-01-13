package config

import (
	"github.com/tommylay1902/authmicro/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {

	dsn := "host=dbauth user=postgres password=password dbname=auth port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		dsnRetry := "host=localhost user=postgres password=password dbname=auth port=8003 sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsnRetry), &gorm.Config{})
		if err != nil {
			panic("error connecting to database")
		}
	}

	db.AutoMigrate(&model.Auth{})

	return db
}
