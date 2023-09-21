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

        protected override void OnModelCreating(ModelBuilder builder)
        {
            builder.Entity<Prescription>()
                .HasIndex(p => p.Medication)
                .IsUnique();
        }

    }
}

