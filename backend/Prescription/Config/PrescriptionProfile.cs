using System;
using AutoMapper;
using prescription.DTO;
using prescription.Entities;

namespace prescription.Config
{
	public class PrescriptionProfile : Profile
	{
		public PrescriptionProfile()
		{
			CreateMap<Prescription, PrescriptionDTO>();
            CreateMap<PrescriptionDTO, Prescription>();
        }
    }
}

