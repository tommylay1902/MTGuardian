
using AutoMapper;
using prescription.DTO;
using prescription.Entities;
using prescription.ErrorHandling.Exceptions;
using prescription.Interfaces;

namespace prescription.ServicesLayer
{
    public class PrescriptionService : IPrescriptionService
	{
        private readonly IPrescriptionRepository _prescriptionRepository;
        private readonly IMapper _mapper;
		public PrescriptionService(IPrescriptionRepository prescriptionRepository, IMapper mapper)
		{
            _prescriptionRepository = prescriptionRepository;
            _mapper = mapper;
            
		}

        public async Task<Guid> CreatePrescriptionAsync(PrescriptionDTO prescription)
        {
            
            if (await _prescriptionRepository.PrescriptionExistsByMedicationAsync(prescription.Medication) == null)
            {
                var p = _mapper.Map<PrescriptionDTO, Prescription>(prescription);
                Guid id = await _prescriptionRepository.AddAsync(p);
                return id;
            }
            else
            {
                throw new ResourceConflictException("You already have this medication prescribed");
            }
        }

        public async Task DeletePrescriptionAsync(Guid id)
        {
            Prescription p = await _prescriptionRepository.GetPrescriptionByIdAsync(id);
            if(p != null)
            {
                await _prescriptionRepository.DeletePrescriptionByEntityAsync(p);
            }
            else
            {
                throw new ResourceNotFoundException("Prescription was not found");
            }
        }

        public async Task<List<Prescription>> GetAllPrescriptionsAsync()
        {
            return await _prescriptionRepository.GetAllPrescriptionsAsync();
        }

        public async Task<Prescription> GetPrescriptionAsync(Guid id)
        {
            return await _prescriptionRepository.GetPrescriptionByIdAsync(id);
        }

        public async Task UpdatePrescriptionAsync(Guid id, PrescriptionDTO p)
        {
            //will throw 404 from database if id not found
            Prescription pToUpdate = await _prescriptionRepository.GetPrescriptionByIdAsync(id);

            bool hasChanges = false;
            if(p.Medication != null && p.Medication != pToUpdate.Medication )
            {
                if(await _prescriptionRepository.PrescriptionExistsByMedicationAsync(p.Medication) == null)
                {
                    pToUpdate.Medication = p.Medication;
                    hasChanges = true;
                }
                else
                {
                    throw new ResourceConflictException("Medication already exists");
                }
            }
            if(p.Doseage != null && p.Doseage != pToUpdate.Doseage)
            {
                pToUpdate.Doseage = p.Doseage;
                hasChanges = true;
            }

            if(p.Notes != pToUpdate.Notes)
            {
                pToUpdate.Notes = p.Notes;
                hasChanges = true;
            }

            if(p.PrescribedAt != pToUpdate.PrescribedAt)
            {
                pToUpdate.PrescribedAt = p.PrescribedAt;
                hasChanges = true;
            }

            if (!hasChanges)
            {
                throw new BadRequestException("no changes found");
            }

            await _prescriptionRepository.UpdatePrescriptionByIdAsync(pToUpdate);


        }
    }
}

