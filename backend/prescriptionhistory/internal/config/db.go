package config

import (
	"github.com/tommylay1902/prescriptionhistory/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {

	dsn := "host=dbprescriptionhistory user=postgres password=password dbname=prescriptionhistory port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		dsnRetry := "host=localhost user=postgres password=password dbname=prescriptionhistory port=8005 sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsnRetry), &gorm.Config{})
		if err != nil {
			panic("error connecting to database")
		}
	}

	db.AutoMigrate(&models.PrescriptionHistory{})

	return db
}
