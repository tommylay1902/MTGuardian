package rxdto

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionmicro/internal/model"
)

func MapPrescriptionDTOToModel(dto *PrescriptionDTO) (*model.Prescription, error) {
	var id, err = uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	model := &model.Prescription{
		ID:         id,
		Medication: dto.Medication,
		Dosage:     dto.Dosage,
		Notes:      dto.Notes,
		Started:    dto.Started,
		Ended:      dto.Ended,
		Refills:    dto.Refills,
		Owner:      dto.Owner,
	}
	return model, nil
}

func MapPrescriptionModelToDTO(p *model.Prescription) (*PrescriptionDTO, error) {
	dto := &PrescriptionDTO{
		Medication: p.Medication,
		Dosage:     p.Dosage,
		Notes:      p.Notes,
		Started:    p.Started,
		Ended:      p.Ended,
		Refills:    p.Refills,
		Owner:      p.Owner,
	}
	return dto, nil
}

func MapPrescriptionModelSliceToDTOSlice(prescriptions []model.Prescription) ([]PrescriptionDTO, error) {
	var resultMapping []PrescriptionDTO
	for _, p := range prescriptions {
		dto, err := MapPrescriptionModelToDTO(&p)
		if err != nil {
			return nil, err
		}
		resultMapping = append(resultMapping, *dto)
	}
	return resultMapping, nil
}

func MapPrescriptionDTOSliceToModelSlice(prescriptions []PrescriptionDTO) ([]model.Prescription, error) {
	var resultMapping []model.Prescription
	for _, dto := range prescriptions {
		p, err := MapPrescriptionDTOToModel(&dto)
		if err != nil {
			return nil, err
		}
		resultMapping = append(resultMapping, *p)
	}
	return resultMapping, nil
}
