using System;
using prescription.Entities;
using Microsoft.EntityFrameworkCore;

namespace prescription.Data
{
	public class PrescriptionContext : DbContext

	{
		public DbSet<Prescription> Prescriptions { get; set; } = null!;

        public PrescriptionContext(DbContextOptions<PrescriptionContext> options) : base(options)
        {
        }


    }
}

