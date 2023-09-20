using System;
using prescription.Entities;

namespace prescription.Interfaces
{
	public interface IPrescriptionRepository
	{
		void Add(Prescription prescription);
	}
}

