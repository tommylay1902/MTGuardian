using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using Newtonsoft.Json;

namespace prescription.Entities
{
	public class Prescription
	{
		[Key]
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
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

        [JsonConstructor]
        public Prescription(string medication, string doseage, DateTime prescribedAt)
        {
            Medication = medication;
            Doseage = doseage;
            PrescribedAt = prescribedAt;
        }

        public Prescription()
        {
        }
    }
}

