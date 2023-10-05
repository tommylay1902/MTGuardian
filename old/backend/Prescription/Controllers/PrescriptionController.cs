using AutoMapper;
using Microsoft.AspNetCore.Mvc;
using prescription.DTO;
using prescription.Entities;
using prescription.ErrorHandling.ExceptionFilters;
using prescription.ErrorHandling.Exceptions;
using prescription.Interfaces;

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
        /// <summary>
        /// Gets all prescriptions
        /// </summary>
        /// <returns>The corresponding prescription.</returns>
        /// <response code="200">an array of all prescriptions</response>
        /// <response code="400">If the request is invalid.</response>
        [HttpGet]
        [ProducesResponseType(typeof(PrescriptionDTO), 200)] // 200 Created
        [ProducesResponseType(400)] // 400 Bad Request
        public async  Task<List<PrescriptionDTO>> GetAllPrescriptions()
        {
            List<Prescription> prescriptions = await _prescriptionService.GetAllPrescriptionsAsync();
            return _mapper.Map<List<Prescription>, List<PrescriptionDTO>>(prescriptions);
        }

        // GET api/values/5
        /// <summary>
        /// Gets a prescription by its id
        /// </summary>
        /// <param name="id">The id of the prescription to get.</param>
        /// <returns>The corresponding prescription.</returns>
        /// <response code="200">Returns the prescription information of the corresponding id.</response>
        /// <response code="400">If the request is invalid.</response>
        [HttpGet("{id}")]
        [ProducesResponseType(typeof(PrescriptionDTO), 200)] // 200 Created
        [ProducesResponseType(400)] // 400 Bad Request
        [ResourceNotFoundExceptionFilter]
        public async Task<PrescriptionDTO> GetById(Guid id)
        {
            return _mapper.Map<PrescriptionDTO>(await _prescriptionService.GetPrescriptionAsync(id));
        }

        // POST api/values
        /// <summary>
        /// Creates a new prescription.
        /// </summary>
        /// <param name="prescription">The prescription to create.</param>
        /// <returns>The created prescription.</returns>
        /// <response code="201">Returns the created prescription.</response>
        /// <response code="400">If the request is invalid or the prescription creation fails.</response>
        [HttpPost]
        [ProducesResponseType(typeof(Guid), 201)] // 201 Created
        [ProducesResponseType(400)] // 400 Bad Request
        [ProducesResponseType(409)]
        [ResourceConflictExceptionFilter]
        [AnotherExceptionFilters]
        public async Task<IActionResult> CreatePrescription([FromBody] PrescriptionDTO prescription)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }
            Guid id = await _prescriptionService.CreatePrescriptionAsync(prescription);
            return CreatedAtAction(nameof(GetById), new { id = id }, id );
        }

        // PUT api/values/5
        [HttpPut("{id}")]
        [ProducesResponseType(200)] // 200 Created
        [ProducesResponseType(400)] // 400 Bad Request
        [BadRequestExceptionFilter]
        [ResourceNoFoundExceptionFilter]
        public async Task<IActionResult> Put(Guid id, [FromBody]PrescriptionDTO p)
        {
            await _prescriptionService.UpdatePrescriptionAsync(id, p);
            return NoContent();
        }

        // DELETE api/values/5
        [HttpDelete("{id}")]
        [ProducesResponseType(200)] // 200 Created
        [ResourceNotFoundExceptionFilter]
        public async Task<IActionResult> Delete(Guid id)
        {
            await _prescriptionService.DeletePrescriptionAsync(id);
            return NoContent();
        }
    }
}

