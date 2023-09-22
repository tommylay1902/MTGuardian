using System;
using prescription.DTO;
using prescription.Entities;

namespace prescription.Interfaces
{
	public interface IPrescriptionService
	{
		public Guid CreatePrescription(Prescription prescription);
		public Prescription GetPrescription(Guid id);
		public List<Prescription> GetAllPrescriptions();
		public void UpdatePrescription(Guid id, PrescriptionDTO p);
	}
}

