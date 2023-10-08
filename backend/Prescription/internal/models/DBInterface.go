package models

import "gorm.io/gorm"

type DBInterface interface {
	Create(interface{}) *gorm.DB
	// Add other methods you use in your DAO as needed
}
