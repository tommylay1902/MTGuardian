﻿
using prescription.DTO;
using prescription.Entities;
using prescription.ErrorHandling.Exceptions;
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
            if (_prescriptionRepository.PrescriptionExistsByMedication(prescription.Medication) == null)
            {
                Guid id = _prescriptionRepository.Add(prescription);
                return id;
            }
            else
            {
                throw new ResourceConflictException("You already have this medication prescribed");
            }
        }

        public List<Prescription> GetAllPrescriptions()
        {
            return _prescriptionRepository.GetAllPrescriptions();
        }

        public Prescription GetPrescription(Guid id)
        {
            return _prescriptionRepository.GetPrescriptionById(id);
        }

        public void UpdatePrescription(Guid id, PrescriptionDTO p)
        {
            //will throw 404 from database if id not found
            Prescription pToUpdate = _prescriptionRepository.GetPrescriptionById(id);

            Boolean hasChanges = false;
        //            public String Medication { get; set; }

        //public String Doseage { get; set; }

        //public String? Notes { get; set; }

        //public DateTime PrescribedAt { get; set; }
            if(p.Medication != null && p.Medication != pToUpdate.Medication )
            {
                if(_prescriptionRepository.PrescriptionExistsByMedication(p.Medication) == null)
                {
                    pToUpdate.Medication = p.Medication;
                    hasChanges = true;
                }
                else
                {
                    throw new ResourceConflictException("Medication already exists");
                }
            }
            if(p.Doseage != null && p.Doseage != pToUpdate.Doseage)
            {
                pToUpdate.Doseage = p.Doseage;
                hasChanges = true;
            }

            if(p.Notes != pToUpdate.Notes)
            {
                pToUpdate.Notes = p.Notes;
                hasChanges = true;
            }

            if(p.PrescribedAt != pToUpdate.PrescribedAt)
            {
                pToUpdate.PrescribedAt = p.PrescribedAt;
                hasChanges = true;
            }

            if (!hasChanges)
            {
                throw new BadRequestException("no changes found");
            }

            _prescriptionRepository.UpdatePrescriptionById(pToUpdate);


        }
    }
}

