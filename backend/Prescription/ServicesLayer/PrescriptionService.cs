using System;
using System.Net;
using prescription.Entities;
using prescription.Interfaces;

namespace prescription.ServicesLayer
{
    public class PrescriptionService : IPrescriptionService
	{
        private readonly IPrescriptionRepository _prescriptionRepository;
		public PrescriptionService(IPrescriptionRepository prescriptionRepository)
		{
            _prescriptionRepository = prescriptionRepository;
		}

        public Guid CreatePrescription(Prescription prescription)
        {
            Guid id = _prescriptionRepository.Add(prescription);
            return id;
        }

        public List<Prescription> GetAllPrescriptions()
        {
            return _prescriptionRepository.GetAllPrescriptions();
        }

        public Prescription GetPrescription(Guid id)
        {
            return _prescriptionRepository.GetPrescriptionById(id);
        }
    }
}

