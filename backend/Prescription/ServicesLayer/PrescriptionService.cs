
using prescription.Entities;
using prescription.ErrorHandling.Exceptions;
using prescription.Interfaces;

namespace prescription.ServicesLayer
{
    public class PrescriptionService : IPrescriptionService
	{
        private readonly IPrescriptionRepository _prescriptionRepository;
		public PrescriptionService(IPrescriptionRepository prescriptionRepository)
		{
            _prescriptionRepository = prescriptionRepository;
		}

        public Guid CreatePrescription(Prescription prescription)
        {
            if (_prescriptionRepository.PrescriptionExistsByMedication(prescription.Medication) == null)
            {
                Guid id = _prescriptionRepository.Add(prescription);
                return id;
            }
            else
            {
                throw new ResourceConflictException("You already have this medication prescribed");
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
    }
}

