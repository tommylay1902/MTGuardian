using System;
using System.ComponentModel.DataAnnotations;

namespace medication.Models
{
	public class Prescription
	{
		[Key]
		public Guid Id { get; set; }

		[Required]
		public String Medication { get; set; }

		[Required]
		public String Doseage { get; set; }

		public String? Notes { get; set; }

		[Required]
		public DateTime PrescribedAt {get; set;}

		

        public Prescription(Guid id, string medication, string doseage, string notes, DateTime prescribedAt)
        {
            Id = id;
            Medication = medication;
            Doseage = doseage;
            Notes = notes;
            PrescribedAt = prescribedAt;
        }


        public Prescription(Guid id, string medication, string doseage,DateTime prescribedAt)
        {
            Id = id;
            Medication = medication;
            Doseage = doseage;
            PrescribedAt = prescribedAt;
        }
    }
}

