package config

import (
	"fmt"
	"log"

	"github.com/tommylay1902/authmicro/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB(dbUsername string, dbHostName string, dbPort string, dbPassword string, dbName string) *gorm.DB {

	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v", dbUsername, dbPassword, dbHostName, dbPort, dbName)

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("error connecting to db", err)
	}

	db.AutoMigrate(&model.Auth{})

	return db
}
