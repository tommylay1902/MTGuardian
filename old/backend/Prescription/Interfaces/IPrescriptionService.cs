using System;
using prescription.DTO;
using prescription.Entities;

namespace prescription.Interfaces
{
	public interface IPrescriptionService
	{
		public Task<Guid> CreatePrescriptionAsync(PrescriptionDTO prescription);
		public Task<Prescription> GetPrescriptionAsync(Guid id);
		public Task<List<Prescription>> GetAllPrescriptionsAsync();
		public Task UpdatePrescriptionAsync(Guid id, PrescriptionDTO p);
		public Task DeletePrescriptionAsync(Guid id);
	}
}

