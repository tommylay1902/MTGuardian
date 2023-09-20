using System;

namespace prescription.DTO
{
	public class PrescriptionDTO
	{


        public String Medication { get; set; }

        public String Doseage { get; set; }

        public String? Notes { get; set; }

        public DateTime PrescribedAt { get; set; }

        public PrescriptionDTO()
		{
		}
		
	}
}

