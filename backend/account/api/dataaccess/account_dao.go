package dataaccess

import "gorm.io/gorm"

type AccountDAO struct {
	DB *gorm.DB
}

func InitializeAccountDAO(db *gorm.DB) *AccountDAO {
	return &AccountDAO{
		DB: db,
	}
}
