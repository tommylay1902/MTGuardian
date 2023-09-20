using System;
using medication.Models;
using Microsoft.EntityFrameworkCore;

namespace medication.Data
{
	public class PrescriptionContext : DbContext

	{
		public DbSet<Prescription> Prescriptions { get; set; } = null!;

        public PrescriptionContext(DbContextOptions<PrescriptionContext> options) : base(options)
        {
        }


    }
}

