package services

import "github.com/tommylay1902/prescriptionmicro/api/dataaccess"

type PrescriptionService struct {
	PrescriptionDAO *dataaccess.PrescriptionDAO
}

func InitalizePrescriptionService(prescriptionDAO *dataaccess.PrescriptionDAO) *PrescriptionService {
	return &PrescriptionService{PrescriptionDAO: prescriptionDAO}
}
