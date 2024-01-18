package rxhistorydto

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/prescriptionhistory/internal/model"
)

// import (
// 	"github.com/google/uuid"
// 	"github.com/tommylay1902/prescriptionhistory/internal/models"
// )

func MapDTOToModel(dto *PrescriptionHistoryDTO) (*model.PrescriptionHistory, error) {
	var id, err = uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	model := &model.PrescriptionHistory{
		Id:             &id,
		Owner:          dto.Owner,
		PrescriptionId: dto.PrescriptionId,
		Taken:          dto.Taken,
	}

	return model, nil
}

func MapModelToDTO(p *model.PrescriptionHistory) (*PrescriptionHistoryDTO, error) {

	dto := &PrescriptionHistoryDTO{
		PrescriptionId: p.PrescriptionId,
		Owner:          p.Owner,
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
