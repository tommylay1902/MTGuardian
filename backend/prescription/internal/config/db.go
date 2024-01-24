package config

import (
	"log"

	"github.com/tommylay1902/prescriptionmicro/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB(port string, host string) *gorm.DB {

	// dsn := "host=" + host + " user=postgres password=password dbname=prescription port=" + port + " sslmode=disable"
	// dsn := "host=" + host + " user=postgres password=password dbname=prescription port=" + port + " sslmode=disable"

	dsn := "postgresql://postgres:password@" + host + ":5432/prescription"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("error connecting to database", err)
	}

	db.AutoMigrate(&model.Prescription{})

	return db
}
