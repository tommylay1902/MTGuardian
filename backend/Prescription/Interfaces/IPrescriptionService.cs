using System;
using prescription.DTO;
using prescription.Entities;

namespace prescription.Interfaces
{
	public interface IPrescriptionService
	{
		public Guid CreatePrescription(PrescriptionDTO prescription);
		public Prescription GetPrescription(Guid id);
		public List<Prescription> GetAllPrescriptions();
		public void UpdatePrescription(Guid id, PrescriptionDTO p);
		public void DeletePrescription(Guid id);
	}
}

