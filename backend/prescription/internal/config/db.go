package config

import (
	"github.com/tommylay1902/prescriptionmicro/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB(port string, host string) *gorm.DB {

	// dsn := "host=" + host + " user=postgres password=password dbname=prescription port=" + port + " sslmode=disable"
	dsn := "host=" + host + " user=postgres password=password dbname=prescription port=" + port + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("error connecting to database")
	}

	db.AutoMigrate(&model.Prescription{})

	return db
}
