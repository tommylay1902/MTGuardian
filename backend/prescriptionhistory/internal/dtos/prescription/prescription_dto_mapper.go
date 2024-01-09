package prescriptiondto

// import (
// 	"github.com/google/uuid"
// 	"github.com/tommylay1902/prescriptionhistory/internal/models"
// )

// func MapPrescriptionDTOToModel(dto *PrescriptionDTO) (*models.Prescription, error) {
// 	var id, err = uuid.NewRandom()
// 	if err != nil {
// 		return nil, err
// 	}
// 	model := &models.Prescription{
// 		ID:         id,
// 		Medication: dto.Medication,
// 		Dosage:     dto.Dosage,
// 		Notes:      dto.Notes,
// 		Started:    dto.Started,
// 		Ended:      dto.Ended,
// 		Refills:    dto.Refills,
// 		Owner:      dto.Owner,
// 	}
// 	return model, nil
// }

// func MapPrescriptionModelToDTO(p *models.Prescription) (*PrescriptionDTO, error) {
// 	dto := &PrescriptionDTO{
// 		Medication: p.Medication,
// 		Dosage:     p.Dosage,
// 		Notes:      p.Notes,
// 		Started:    p.Started,
// 		Ended:      p.Ended,
// 		Refills:    p.Refills,
// 		Owner:      p.Owner,
// 	}
// 	return dto, nil
// }

// func MapPrescriptionModelSliceToDTOSlice(prescriptions []models.Prescription) ([]PrescriptionDTO, error) {
// 	var resultMapping []PrescriptionDTO
// 	for _, p := range prescriptions {
// 		dto, err := MapPrescriptionModelToDTO(&p)
// 		if err != nil {
// 			return nil, err
// 		}
// 		resultMapping = append(resultMapping, *dto)
// 	}
// 	return resultMapping, nil
// }

// func MapPrescriptionDTOSliceToModelSlice(prescriptions []PrescriptionDTO) ([]models.Prescription, error) {
// 	var resultMapping []models.Prescription
// 	for _, dto := range prescriptions {
// 		p, err := MapPrescriptionDTOToModel(&dto)
// 		if err != nil {
// 			return nil, err
// 		}
// 		resultMapping = append(resultMapping, *p)
// 	}
// 	return resultMapping, nil
// }
