using System;
using prescription.Entities;
using prescription.Interfaces;

namespace prescription.ServicesLayer
{
	public class PrescriptionService:IPrescriptionService
	{
        private readonly IPrescriptionRepository _prescriptionRepository;
		public PrescriptionService(IPrescriptionRepository prescriptionRepository)
		{
            _prescriptionRepository = prescriptionRepository;
		}

        public void CreatePrescription(Prescription prescription)
        {
            _prescriptionRepository.Add(prescription);
        }
    }
}

