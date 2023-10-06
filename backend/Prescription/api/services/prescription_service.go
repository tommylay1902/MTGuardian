package services

import (
	"fmt"

	"github.com/tommylay1902/prescriptionmicro/api/dataaccess"
	"github.com/tommylay1902/prescriptionmicro/internal/dtos"
)

type PrescriptionService struct {
	PrescriptionDAO *dataaccess.PrescriptionDAO
}

func InitalizePrescriptionService(prescriptionDAO *dataaccess.PrescriptionDAO) *PrescriptionService {
	return &PrescriptionService{PrescriptionDAO: prescriptionDAO}
}

func (ps *PrescriptionService) CreatePrescription(prescription *dtos.PrescriptionDTO) error {
	create, dtoErr := dtos.MapPrescriptionDTOToModel(prescription)
	if dtoErr != nil {
		return dtoErr
	}
	fmt.Println("id", create.ID)
	err := ps.PrescriptionDAO.CreatePrescription(create)
	if err != nil {
		return err
	}
	return nil
}
