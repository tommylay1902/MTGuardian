using System;
using prescription.Entities;

namespace prescription.Interfaces
{
	public interface IPrescriptionRepository
	{
		Guid Add(Prescription prescription);
	}
}

