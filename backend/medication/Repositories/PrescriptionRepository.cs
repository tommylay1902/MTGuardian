using System;
using prescription.Data;
using prescription.Entities;
using prescription.Interfaces;

namespace prescription.Repositories
{
	public class PrescriptionRespotiory:IPrescriptionRepository

	{
        private readonly PrescriptionContext _context;

		public PrescriptionRespotiory(PrescriptionContext context)
		{
            _context = context;
		}

        public void Add(Prescription prescription)
        {
            _context.Add(prescription);
            _context.SaveChanges();
        }
    }
}

