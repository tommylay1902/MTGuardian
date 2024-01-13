package prescriptionhistorydto

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionhistory/internal/models"
)

// import (
// 	"github.com/google/uuid"
// 	"github.com/tommylay1902/prescriptionhistory/internal/models"
// )

func MapDTOToModel(dto *PrescriptionHistoryDTO) (*models.PrescriptionHistory, error) {
	var id, err = uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	model := &models.PrescriptionHistory{
		Id:             id,
		OwnerId:        dto.OwnerId,
		PrescriptionId: dto.PrescriptionId,
		Taken:          dto.Taken,
	}

	return model, nil
}

func MapModelToDTO(p *models.PrescriptionHistory) (*PrescriptionHistoryDTO, error) {

	dto := &PrescriptionHistoryDTO{
		PrescriptionId: p.PrescriptionId,
		OwnerId:        p.OwnerId,
		Taken:          p.Taken,
	}

	return dto, nil
}

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
