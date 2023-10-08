package services

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionmicro/api/dataaccess"
	dto "github.com/tommylay1902/prescriptionmicro/internal/dtos/prescription"
	"github.com/tommylay1902/prescriptionmicro/internal/error/customerrors"
)

type PrescriptionService struct {
	PrescriptionDAO *dataaccess.PrescriptionDAO
}

func InitalizePrescriptionService(prescriptionDAO *dataaccess.PrescriptionDAO) *PrescriptionService {
	return &PrescriptionService{PrescriptionDAO: prescriptionDAO}
}

func (ps *PrescriptionService) CreatePrescription(prescription *dto.PrescriptionDTO) error {
	create, dtoErr := dto.MapPrescriptionDTOToModel(prescription)
	if dtoErr != nil {
		return dtoErr
	}
	err := ps.PrescriptionDAO.CreatePrescription(create)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PrescriptionService) GetPrescriptionById(id uuid.UUID) (*dto.PrescriptionDTO, error) {
	p, err := ps.PrescriptionDAO.GetPrescriptionById(id)
	if err != nil {
		return nil, err
	}
	resultMapping, mappingErr := dto.MapPrescriptionModelToDTO(p)
	if mappingErr != nil {
		return nil, mappingErr
	}
	return resultMapping, nil
}

func (ps *PrescriptionService) GetPrescriptions() ([]dto.PrescriptionDTO, error) {
	prescriptions, err := ps.PrescriptionDAO.GetAllPrescriptions()

	if err != nil {
		return nil, err
	}
	resultMapping, mappingErr := dto.MapPrescriptionModelSliceToDTOSlice(prescriptions)
	if mappingErr != nil {
		return nil, mappingErr
	}
	return resultMapping, nil
}

func (ps *PrescriptionService) DeletePrescription(id uuid.UUID) error {
	p, err := ps.PrescriptionDAO.GetPrescriptionById(id)
	if err != nil {
		return &customerrors.ResourceNotFound{
			Message: err.Error(),
			Code:    404,
		}
	}
	daoError := ps.PrescriptionDAO.DeletePrescription(p)
	if daoError != nil {
		return daoError
	}

	return nil
}

func (ps *PrescriptionService) UpdatePrescription(pDTO *dto.PrescriptionDTO, id uuid.UUID) error {
	pUpdate, err := ps.PrescriptionDAO.GetPrescriptionById(id)
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

	if hasUpdate {
		return ps.PrescriptionDAO.UpdatePrescription(pUpdate)
	}
	return &customerrors.BadRequestError{Message: "No updates found for the prescription", Code: 400}
}
