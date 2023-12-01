package services

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionmicro/api/dataaccess"
	dto "github.com/tommylay1902/prescriptionmicro/internal/dtos/prescription"
	"github.com/tommylay1902/prescriptionmicro/internal/error/customerrors"
	"github.com/tommylay1902/prescriptionmicro/internal/models"
)

type PrescriptionService struct {
	dao dataaccess.IPrescriptionDao
}

func InitalizePrescriptionService(prescriptionDAO dataaccess.IPrescriptionDao) *PrescriptionService {
	return &PrescriptionService{dao: prescriptionDAO}
}

func (ps *PrescriptionService) CreatePrescription(prescription *dto.PrescriptionDTO) (*uuid.UUID, error) {
	create, dtoErr := dto.MapPrescriptionDTOToModel(prescription)

	if dtoErr != nil {
		return nil, dtoErr
	}
	id, err := ps.dao.CreatePrescription(create)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (ps *PrescriptionService) GetPrescriptionById(id uuid.UUID, email string) (*models.Prescription, error) {
	p, err := ps.dao.GetPrescriptionById(id, email)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (ps *PrescriptionService) GetPrescriptions(searchQuery map[string]string, owner *string) ([]models.Prescription, error) {

	prescriptions, err := ps.dao.GetAllPrescriptions(searchQuery, owner)

	if err != nil {
		return nil, err
	}

	return prescriptions, nil
}

func (ps *PrescriptionService) DeletePrescription(id uuid.UUID, email string) error {
	p, err := ps.dao.GetPrescriptionById(id, email)
	if err != nil {
		return &customerrors.ResourceNotFound{
			Message: err.Error(),
			Code:    404,
		}
	}
	daoError := ps.dao.DeletePrescription(p, email)
	if daoError != nil {
		return daoError
	}
	return nil
}

// test
func (ps *PrescriptionService) UpdatePrescription(pDTO *dto.PrescriptionDTO, id uuid.UUID, email string) error {
	pUpdate, err := ps.dao.GetPrescriptionById(id, email)
	if err != nil {
		return err
	}
	hasUpdate := false
	if pDTO.Dosage != nil && *pDTO.Dosage != *pUpdate.Dosage {
		hasUpdate = true
		*pUpdate.Dosage = *pDTO.Dosage
	}

	if pDTO.Medication != nil && *pDTO.Medication != *pUpdate.Medication {
		hasUpdate = true
		*pUpdate.Medication = *pDTO.Medication
	}

	if pDTO.Notes != nil && *pDTO.Notes != *pUpdate.Notes {
		hasUpdate = true
		*pUpdate.Notes = *pDTO.Notes
	}

	if pDTO.Started != nil && *pDTO.Started != *pUpdate.Started {
		hasUpdate = true
		*pUpdate.Started = *pDTO.Started
	}

	if pUpdate.Ended == nil && pDTO.Ended != nil || pDTO.Ended == nil && pUpdate.Ended != nil {
		hasUpdate = true
		pUpdate.Ended = pDTO.Ended
	} else if pUpdate.Ended != nil && pDTO.Ended != nil && *pUpdate.Ended != *pDTO.Ended {
		hasUpdate = true
		*pUpdate.Ended = *pDTO.Ended
	}

	if hasUpdate {
		return ps.dao.UpdatePrescription(pUpdate, email)
	}

	return &customerrors.BadRequestError{Message: "No updates found for the prescription", Code: 400}
}
