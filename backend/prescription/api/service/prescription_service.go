package service

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionmicro/api/dao"
	dto "github.com/tommylay1902/prescriptionmicro/internal/dto/prescription"
	"github.com/tommylay1902/prescriptionmicro/internal/error/apperror"
	"github.com/tommylay1902/prescriptionmicro/internal/model"
)

type PrescriptionService struct {
	DAO dao.IPrescriptionDao
}

func Initialize(dao dao.IPrescriptionDao) *PrescriptionService {
	return &PrescriptionService{DAO: dao}
}

func (ps *PrescriptionService) CreatePrescription(pDTO *dto.PrescriptionDTO) (*uuid.UUID, error) {
	create, dtoErr := dto.MapPrescriptionDTOToModel(pDTO)

	if dtoErr != nil {
		return nil, dtoErr
	}

	id, err := ps.DAO.CreatePrescription(create)

	if err != nil {
		return nil, err
	}

	return id, nil
}

func (ps *PrescriptionService) GetPrescriptionById(id uuid.UUID, email string) (*model.Prescription, error) {
	p, err := ps.DAO.GetPrescriptionById(id, email)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (ps *PrescriptionService) GetPrescriptions(searchQuery map[string]string, owner *string) ([]model.Prescription, error) {
	prescriptions, err := ps.DAO.GetAllPrescriptions(searchQuery, owner)

	if err != nil {
		return nil, err
	}

	return prescriptions, nil
}

func (ps *PrescriptionService) DeletePrescription(id uuid.UUID, email string) error {
	prescription, err := ps.DAO.GetPrescriptionById(id, email)
	if err != nil {
		return &apperror.ResourceNotFound{
			Message: err.Error(),
			Code:    404,
		}
	}
	daoError := ps.DAO.DeletePrescription(prescription, email)
	if daoError != nil {
		return daoError
	}
	return nil
}

func (ps *PrescriptionService) UpdatePrescription(pDTO *dto.PrescriptionDTO, id uuid.UUID, email string) error {
	pUpdate, err := ps.DAO.GetPrescriptionById(id, email)
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

	if pUpdate.Refills != nil && *pDTO.Refills != *pUpdate.Refills {
		hasUpdate = true
		*pUpdate.Refills = *pDTO.Refills
	}

	if hasUpdate {
		return ps.DAO.UpdatePrescription(pUpdate, email)
	}

	return &apperror.BadRequestError{Message: "No updates found for the prescription", Code: 400}
}
