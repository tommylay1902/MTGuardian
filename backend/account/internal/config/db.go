package config

import (
	"github.com/tommylay1902/accountmicro/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {

	dsn := "host=dbprescription user=postgres password=password dbname=prescription port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		dsnRetry := "host=localhost user=postgres password=password dbname=prescription port=8001 sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsnRetry), &gorm.Config{})
		if err != nil {
			panic("error connecting to database")
		}
	}

	db.AutoMigrate(&models.Account{})

	return db
}
