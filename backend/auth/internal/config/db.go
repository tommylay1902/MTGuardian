package config

import (
	"github.com/tommylay1902/authmicro/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB(port string, host string) *gorm.DB {

	// dsn := "host=" + host + " user=postgres password=password dbname=auth port=" + port + " sslmode=disable"

	dsn := "postgresql://postgres:password@" + host + ":5432/auth"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("error connecting to database")
	}

	db.AutoMigrate(&model.Auth{})

	return db
}
