using System;
using prescription.Data;
using prescription.Entities;
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
    }
}

