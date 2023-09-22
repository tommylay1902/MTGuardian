
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

        public Guid CreatePrescription(PrescriptionDTO prescription)
        {
            if (_prescriptionRepository.PrescriptionExistsByMedication(prescription.Medication) == null)
            {
                var p = _mapper.Map<PrescriptionDTO, Prescription>(prescription);
                Guid id = _prescriptionRepository.Add(p);
                return id;
            }
            else
            {
                throw new ResourceConflictException("You already have this medication prescribed");
            }
        }

        public void DeletePrescription(Guid id)
        {
            Prescription p = _prescriptionRepository.GetPrescriptionById(id);
            if(p != null)
            {
                _prescriptionRepository.DeletePrescriptionByEntity(p);
            }
            else
            {
                throw new ResourceNotFoundException("Prescription was not found");
            }
        }

        public List<Prescription> GetAllPrescriptions()
        {
            return _prescriptionRepository.GetAllPrescriptions();
        }

        public Prescription GetPrescription(Guid id)
        {
            return _prescriptionRepository.GetPrescriptionById(id);
        }

        public void UpdatePrescription(Guid id, PrescriptionDTO p)
        {
            //will throw 404 from database if id not found
            Prescription pToUpdate = _prescriptionRepository.GetPrescriptionById(id);

            bool hasChanges = false;
            if(p.Medication != null && p.Medication != pToUpdate.Medication )
            {
                if(_prescriptionRepository.PrescriptionExistsByMedication(p.Medication) == null)
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

            _prescriptionRepository.UpdatePrescriptionById(pToUpdate);


        }
    }
}

