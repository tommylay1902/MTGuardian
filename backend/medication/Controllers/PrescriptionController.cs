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
        public string GetById(int id)
        {
            return "value";
        }

        // POST api/values
        /// <summary>
        /// Creates a new item.
        /// </summary>
        /// <param name="prescirption">The prescription to create.</param>
        /// <returns>The created prescription.</returns>
        /// <response code="201">Returns the created prescription.</response>
        /// <response code="400">If the request is invalid or the item creation fails.</response>
        [HttpPost]
        [ProducesResponseType(typeof(Prescription), 201)] // 201 Created
        [ProducesResponseType(400)] // 400 Bad Request
        public ActionResult<PrescriptionDTO> Post([FromBody] Prescription prescription)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }
            Guid id = _prescriptionService.CreatePrescription(prescription);
            return CreatedAtAction(nameof(GetById), new { id = id }, prescription );
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

