package handlers

import "github.com/tommylay1902/prescriptionmicro/api/services"

type PrescriptionHandler struct {
	PrescriptionService *services.PrescriptionService
}

func InitializePrescriptionHandler(prescriptionService *services.PrescriptionService) *PrescriptionHandler {
	return &PrescriptionHandler{PrescriptionService: prescriptionService}
}
