using System;
using prescription.Entities;

namespace prescription.Interfaces
{
	public interface IPrescriptionService
	{
		public Guid CreatePrescription(Prescription prescription);
	}
}

