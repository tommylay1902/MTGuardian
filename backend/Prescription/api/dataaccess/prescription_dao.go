package dataaccess

import "gorm.io/gorm"

type PrescriptionDAO struct {
	DB *gorm.DB
}

func InitalizePrescriptionService(db *gorm.DB) *PrescriptionDAO {
	return &PrescriptionDAO{DB: db}
}
