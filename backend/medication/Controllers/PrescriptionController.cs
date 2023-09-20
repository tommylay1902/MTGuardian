using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using AutoMapper;
using Microsoft.AspNetCore.Mvc;
using prescription.DTO;
using prescription.Entities;
using prescription.Interfaces;

// For more information on enabling Web API for empty projects, visit https://go.microsoft.com/fwlink/?LinkID=397860

namespace prescription.Controllers
{
    [Route("api/v1/[controller]")]
    public class PrescriptionController : Controller
    {
        private readonly IMapper _mapper;
        private readonly IPrescriptionService _prescriptionService;

        public PrescriptionController(IMapper mapper, IPrescriptionService prescriptionService)
        {
            _mapper = mapper;
            _prescriptionService = prescriptionService;
        }


        // GET: api/values
        [HttpGet]
        public IEnumerable<string> Get()
        {
            return new string[] { "value1", "value2" };
        }

        // GET api/values/5
        [HttpGet("{id}")]
        public string Get(int id)
        {
            return "value";
        }

        // POST api/values
        [HttpPost]
        public ActionResult<PrescriptionDTO> Post([FromBody]Prescription prescription)
        {
            _prescriptionService.CreatePrescription(prescription);
            return Ok();
        }

        // PUT api/values/5
        [HttpPut("{id}")]
        public void Put(int id, [FromBody]string value)
        {
        }

        // DELETE api/values/5
        [HttpDelete("{id}")]
        public void Delete(int id)
        {
        }
    }
}

