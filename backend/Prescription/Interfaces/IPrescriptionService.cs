using System;
using prescription.Entities;

namespace prescription.Interfaces
{
	public interface IPrescriptionService
	{
		public Guid CreatePrescription(Prescription prescription);
		public Prescription GetPrescription(Guid id);
		public List<Prescription> GetAllPrescriptions();

	}
}

