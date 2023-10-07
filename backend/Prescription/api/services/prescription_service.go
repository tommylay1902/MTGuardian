package services

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionmicro/api/dataaccess"
	"github.com/tommylay1902/prescriptionmicro/internal/dtos"
	"github.com/tommylay1902/prescriptionmicro/internal/models"
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
	err := ps.PrescriptionDAO.CreatePrescription(create)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PrescriptionService) GetPrescriptionById(id uuid.UUID) (*models.Prescription, error) {
	p, err := ps.PrescriptionDAO.GetPrescriptionById(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *PrescriptionService) GetPrescriptions() ([]models.Prescription, error) {
	prescriptions, err := ps.PrescriptionDAO.GetAllPrescriptions()

	if err != nil {
		return nil, err
	}

	return prescriptions, nil
}
