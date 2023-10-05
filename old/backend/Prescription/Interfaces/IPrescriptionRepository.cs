using System;
using prescription.DTO;
using prescription.Entities;

namespace prescription.Interfaces
{
	public interface IPrescriptionRepository
	{
		Task<Guid> AddAsync(Prescription prescription);
		Task<Prescription> GetPrescriptionByIdAsync(Guid id);
		Task<List<Prescription>> GetAllPrescriptionsAsync();
		Task<Prescription?> PrescriptionExistsByMedicationAsync(String medication);
		Task UpdatePrescriptionByIdAsync(Prescription p);
		Task DeletePrescriptionByEntityAsync(Prescription p);
	}
}

