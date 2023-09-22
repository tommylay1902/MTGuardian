using Microsoft.EntityFrameworkCore;
using prescription.Data;
using prescription.DTO;
using prescription.Entities;
using prescription.ErrorHandling.Exceptions;
using prescription.Interfaces;

namespace prescription.Repositories
{
    public class PrescriptionRepository : IPrescriptionRepository

    {
        private readonly PrescriptionContext _context;

        public PrescriptionRepository(PrescriptionContext context)
        {
            _context = context;
        }

        public async Task<Guid> AddAsync(Prescription prescription)
        {
            
            await _context.AddAsync(prescription);
            await _context.SaveChangesAsync();
            return prescription.Id;
        }

        public async Task DeletePrescriptionByEntityAsync(Prescription p)
        {
            _context.Prescriptions.Remove(p);
            await _context.SaveChangesAsync();
        }

        public async Task<List<Prescription>> GetAllPrescriptionsAsync()
        {
            return await _context.Prescriptions.ToListAsync();
        }

        public async Task<Prescription> GetPrescriptionByIdAsync(Guid id)
        {
            var prescription = await _context.Prescriptions.FindAsync(id);
            if (prescription == null)
            {
                throw new ResourceNotFoundException("Prescription not found");
            }
            return prescription;
        }

        public async Task<Prescription?> PrescriptionExistsByMedicationAsync(string medication)
        {
            return await _context.Prescriptions
            .FirstOrDefaultAsync(p => p.Medication == medication);
        }

        public async Task UpdatePrescriptionByIdAsync(Prescription p)
        {

            // Attach the object to the context and mark it as modified.
            _context.Attach(p);
            _context.Entry(p).State = EntityState.Modified;

            // Save changes to persist the update.
            await _context.SaveChangesAsync();
        }

       
    }
}

