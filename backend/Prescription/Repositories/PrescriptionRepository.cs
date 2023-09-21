using System;
using Microsoft.EntityFrameworkCore;
using prescription.Data;
using prescription.Entities;
using prescription.ErrorHandling.Exceptions;
using prescription.Interfaces;

namespace prescription.Repositories
{
	public class PrescriptionRepository:IPrescriptionRepository

	{
        private readonly PrescriptionContext _context;

		public PrescriptionRepository(PrescriptionContext context)
		{
            _context = context;
		}

        public Guid Add(Prescription prescription)
        {
            _context.Add(prescription);
            _context.SaveChanges();
            return prescription.Id;
        }

        public List<Prescription> GetAllPrescriptions()
        {
            return _context.Prescriptions.ToList();
        }

        public Prescription GetPrescriptionById(Guid id)
        {
            var prescription = _context.Prescriptions.Find(id);
            if(prescription == null)
            {
                throw new ResourceNotFoundException("Prescription not found");
            }
            return prescription;
        }

        public Prescription PrescriptionExistsByMedication(string medication)
        {
            return _context.Prescriptions
            .FirstOrDefault(p => p.Medication == medication);
        }
    }
}

