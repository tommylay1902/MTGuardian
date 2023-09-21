using Testcontainers.PostgreSql;
using prescription.Repositories;
using prescription.Data;
using Microsoft.EntityFrameworkCore;
using prescription.Entities;
using prescription.ErrorHandling.Exceptions;


namespace PrescriptionUnitTest;

public sealed class PrescriptionDAOTest : IAsyncLifetime
{
    private readonly PostgreSqlContainer _postgreSqlContainer = new PostgreSqlBuilder()
        .WithImage("postgres:15-alpine")
        .Build();

    private PrescriptionRepository? _prescriptionRepository;
    private PrescriptionContext? _context;

    public async Task InitializeAsync()
    {
        await _postgreSqlContainer.StartAsync();

        // Initialize the connection to the test database.
        var connectionString = _postgreSqlContainer.GetConnectionString();
        connectionString += ";Pooling=false";
        var options = new DbContextOptionsBuilder<PrescriptionContext>()
            .UseNpgsql(connectionString) // Use your database provider here
            .Options;

        // Initialize the PrescriptionContext with the options.
        _context = new PrescriptionContext(options);

        // Ensure the database is created if it doesn't exist.
        _context.Database.EnsureCreated();

        // Initialize the PrescriptionRepository with the context.
        _prescriptionRepository = new PrescriptionRepository(_context);
    }



    public async Task DisposeAsync()
    {
        if(_context != null)
        {
            await _context.DisposeAsync();
        }
        await _postgreSqlContainer.DisposeAsync().AsTask();
    }


    [Fact]
    public void CreatePrescription_ValidInput_ReturnsPrescriptionId()
    {
        // Arrange: Create a test prescription object.
        var prescription = new Prescription
        {
            Medication = "Dexamethasone",
            Doseage = "20 mg Daily",
            Notes = "A steroid used to help inflamed areas of the body",
            PrescribedAt = DateTime.Parse("2023-09-21T00:59:37.942Z").ToUniversalTime()
        };

        if(_prescriptionRepository != null)
        {
            // Act: Create the prescription using the DAO.
            var prescriptionId = _prescriptionRepository.Add(prescription);

            // Assert: Verify that the prescription was created successfully and has an ID.
            Assert.IsType<Guid>(prescriptionId);
        }
        else
        {
            Assert.Fail("_prescriptionRepository is null");
        }
    }

    [Theory]
    [InlineData(true, "Dexamethasone", "20 mg Daily", "A steroid used to help inflamed areas of the body", "2023-09-21T00:59:37.942Z")]
    [InlineData(false, "Nonexistent Medication", "10 mg Daily", "Description", "2023-09-21T01:00:00.000Z")]
    public void GetPrescriptionById_ValidOrInvalidId_ReturnsPrescriptionOrThrowsException(bool validId, string medication, string dosage, string notes, string prescribedAt)
    {
        if (_context != null && _prescriptionRepository != null)
        {
            // Arrange
            var prescriptionId = Guid.NewGuid(); // Use a valid or invalid ID
            var expectedPrescription = new Prescription
            {
                Id = prescriptionId,
                Medication = medication,
                Doseage = dosage,
                Notes = notes,
                PrescribedAt = DateTime.Parse(prescribedAt).ToUniversalTime()
            };

            _context.Prescriptions.Add(expectedPrescription);
            _context.SaveChanges();

            if (validId)
            {
                // Act
                var result = _prescriptionRepository.GetPrescriptionById(prescriptionId);

                // Assert
                Assert.NotNull(result);
                Assert.Equal(prescriptionId, result.Id);
            }
            else
            {
                var invalidId = Guid.NewGuid();
                // Act & Assert: Verify that an exception of type ResourceNotFoundException is thrown.
                var ex = Assert.Throws<ResourceNotFoundException>(() => _prescriptionRepository.GetPrescriptionById(invalidId));

                Assert.Equal("Prescription not found", ex.Message);
            }
        }
        else
        {
            Assert.Fail("_context and/or _prescriptionRepository is null");
        }
    }

    [Fact]
    public void GetAllPrescriptions_ReturnsAllPrescriptions()
    {
        if(_context != null && _prescriptionRepository != null)
        {
            // Arrange: Create a list of test prescriptions.
            var expectedPrescriptions = new List<Prescription>
            {
                new Prescription
                {
                    Id = new Guid(),
                    Medication = "Medication1",
                    Doseage = "10 mg Daily",
                    Notes = "Note1",
                    PrescribedAt = DateTime.UtcNow
                },
                new Prescription
                {
                    Id = new Guid(),
                    Medication = "Medication2",
                    Doseage = "20 mg Daily",
                    Notes = "Note2",
                    PrescribedAt = DateTime.UtcNow
                }
            };

            foreach (var prescription in expectedPrescriptions)
            {
                _context.Prescriptions.Add(prescription);
            }
            _context.SaveChanges();

            // Act: Retrieve all prescriptions using the DAO method.
            var actualPrescriptions = _prescriptionRepository.GetAllPrescriptions();

            // Assert: Verify that all expected prescriptions are retrieved.
            Assert.NotNull(actualPrescriptions);
            Assert.Equal(expectedPrescriptions.Count, actualPrescriptions.Count);


            foreach (var expectedPrescription in expectedPrescriptions)
            {
                var actualPrescription = actualPrescriptions.SingleOrDefault(p => p.Id == expectedPrescription.Id);
                Assert.NotNull(actualPrescription);
                Assert.Equal(expectedPrescription.Medication, actualPrescription.Medication);
                Assert.Equal(expectedPrescription.Doseage, actualPrescription.Doseage);
                Assert.Equal(expectedPrescription.Notes, actualPrescription.Notes);
                Assert.Equal(expectedPrescription.PrescribedAt, actualPrescription.PrescribedAt);
            }
        }
        else
        {
            Assert.Fail("_context or _prescriptionRepository is null");
        }
        
    }



}
