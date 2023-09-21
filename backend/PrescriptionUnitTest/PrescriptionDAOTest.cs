using System.Data.Common;
using Npgsql;
using prescription.Interfaces;
using Testcontainers.PostgreSql;
using prescription.Repositories;
using prescription.Data;
using Microsoft.EntityFrameworkCore;
using prescription.Entities;
using Xunit.Abstractions;

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


    }

}
