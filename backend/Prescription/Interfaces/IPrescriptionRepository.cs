using System;
using prescription.Entities;

namespace prescription.Interfaces
{
	public interface IPrescriptionRepository
	{
		Guid Add(Prescription prescription);
		Prescription GetPrescriptionById(Guid id);
		List<Prescription> GetAllPrescriptions();
		Prescription? PrescriptionExistsByMedication(String medication);
	}
}

