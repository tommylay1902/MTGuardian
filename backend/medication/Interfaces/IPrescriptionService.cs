using System;
using prescription.Entities;

namespace prescription.Interfaces
{
	public interface IPrescriptionService
	{
		public void CreatePrescription(Prescription prescription);
	}
}

