﻿using AutoMapper;
using Microsoft.EntityFrameworkCore;
using Moq;
using prescription.DTO;
using prescription.Entities;
using prescription.ErrorHandling.Exceptions;
using prescription.Interfaces;
using prescription.ServicesLayer;

public class PrescriptionServiceTest
{
    [Theory]
    [InlineData(true)] // Test for a valid prescription
    [InlineData(false)] // Test for an invalid prescription
    public void CreatePrescription_ValidInput_ReturnsGuidOrThrowsException(bool isValidPrescription)
    {
        // Arrange
        var prescriptionRepositoryMock = new Mock<IPrescriptionRepository>();
        var mapperMock = new Mock<IMapper>();
        var prescriptionService = new PrescriptionService(prescriptionRepositoryMock.Object, mapperMock.Object);

        if (isValidPrescription)
        {
            var expectedGuid = Guid.NewGuid();
            prescriptionRepositoryMock.Setup(repo => repo.Add(It.IsAny<Prescription>()))
                .Returns(expectedGuid);

            var prescription = new PrescriptionDTO
            {
                Medication = "Dexamethasone",
                Doseage = "20 mg Daily",
                Notes = "A steroid used to help inflamed areas of the body",
                PrescribedAt = DateTime.Parse("2023-09-21T00:59:37.942Z").ToUniversalTime()
            };

            // Act
            var result = prescriptionService.CreatePrescription(prescription);

            // Assert
            Assert.Equal(expectedGuid, result);
        }
        else
        {
            var prescriptionWithDuplicateMedicationDto = new PrescriptionDTO
            {
                Medication = "Dexamethasone", // Create a prescription with the same medication value
                Doseage = "20 mg Daily",
                Notes = "A steroid used to help inflamed areas of the body",
                PrescribedAt = DateTime.Parse("2023-09-21T00:59:37.942Z").ToUniversalTime()
            };

            // Mock AutoMapper behavior to simulate mapping
            mapperMock.Setup(mapper => mapper.Map<PrescriptionDTO, Prescription>(It.IsAny<PrescriptionDTO>()))
                .Returns(new Prescription());


            // Set up the mock to throw an exception when Add is called with a duplicate Medication value
            prescriptionRepositoryMock.Setup(repo => repo.Add(It.IsAny<Prescription>()))
                .Throws<ResourceConflictException>(); // Replace with the appropriate exception type for a unique constraint violation

            // Act and Assert (for exception)
            Assert.Throws<ResourceConflictException>(() => prescriptionService.CreatePrescription(prescriptionWithDuplicateMedicationDto));
        }
    }


    [Fact]
    public void GetPrescription_Returns_Prescription_When_IdExists()
    {
        // Arrange
        Guid existingPrescriptionId = Guid.NewGuid();
        Prescription expectedPrescription = new Prescription
        {
            Id = existingPrescriptionId,
            Medication = "Dexamethasone",
            Doseage = "20 mg Daily",
            Notes = "A steroid used to help inflamed areas of the body",
            PrescribedAt = DateTime.Parse("2023-09-21T00:59:37.942Z").ToUniversalTime()
        };

        var mockPrescriptionRepository = new Mock<IPrescriptionRepository>();
        var mockMapper = new Mock<IMapper>();
        mockPrescriptionRepository
            .Setup(repo => repo.GetPrescriptionById(existingPrescriptionId))
            .Returns(expectedPrescription);

        var prescriptionService = new PrescriptionService(mockPrescriptionRepository.Object, mockMapper.Object);

        // Act
        Prescription result = prescriptionService.GetPrescription(existingPrescriptionId);

        // Assert
        Assert.NotNull(result);
        Assert.Equal(existingPrescriptionId, result.Id);
        Assert.Equal(expectedPrescription.Medication, result.Medication);
        Assert.Equal(expectedPrescription.Doseage, result.Doseage);
        Assert.Equal(expectedPrescription.Notes, result.Notes);
        Assert.Equal(expectedPrescription.PrescribedAt, result.PrescribedAt);
    }

    [Fact]
    public void GetPrescription_Returns_Null_When_IdDoesNotExist()
    {
        // Arrange
        Guid nonExistentPrescriptionId = Guid.NewGuid();

        var mockPrescriptionRepository = new Mock<IPrescriptionRepository>();
        var mockMapper = new Mock<IMapper>();

        mockPrescriptionRepository
            .Setup(repo => repo.GetPrescriptionById(nonExistentPrescriptionId))
            .Returns( (Prescription)null);

        var prescriptionService = new PrescriptionService(mockPrescriptionRepository.Object, mockMapper.Object);

        // Act
        Prescription result = prescriptionService.GetPrescription(nonExistentPrescriptionId);

        // Assert
        Assert.Null(result);
    }

    [Fact]
    public void GetPrescription_Throws_Exception_When_DaoLayerFails()
    {
        // Arrange
        Guid prescriptionId = Guid.NewGuid();
        var mockPrescriptionRepository = new Mock<IPrescriptionRepository>();
        var mockMapper = new Mock<IMapper>();

        mockPrescriptionRepository
            .Setup(repo => repo.GetPrescriptionById(prescriptionId))
            .Throws(new ResourceNotFoundException("Prescription not found")); // Customize the exception as per your actual implementation

        var prescriptionService = new PrescriptionService(mockPrescriptionRepository.Object, mockMapper.Object);

        // Act & Assert
        var exception = Assert.Throws<ResourceNotFoundException>(() => prescriptionService.GetPrescription(prescriptionId));
        Assert.Equal("Prescription not found", exception.Message);
    }


    [Fact]
    public void GetAllPrescriptions_ReturnsAllPrescriptions()
    {
        // Arrange: Create a list of test prescriptions.
        var expectedPrescriptions = new List<Prescription>
        {
            new Prescription
            {
          
                Medication = "Medication1",
                Doseage = "10 mg Daily",
                Notes = "Note1",
                PrescribedAt = System.DateTime.UtcNow
            },
            new Prescription
            {
            
                Medication = "Medication2",
                Doseage = "20 mg Daily",
                Notes = "Note2",
                PrescribedAt = System.DateTime.UtcNow
            }
        };

        var mockPrescriptionRepository = new Mock<IPrescriptionRepository>();
        var mockMapper = new Mock<IMapper>();

        mockPrescriptionRepository.Setup(repo => repo.GetAllPrescriptions())
            .Returns(expectedPrescriptions);

        var prescriptionService = new PrescriptionService(mockPrescriptionRepository.Object, mockMapper.Object);

        // Act: Retrieve all prescriptions using the service method.
        var actualPrescriptions = prescriptionService.GetAllPrescriptions();

        // Assert: Verify that all expected prescriptions are retrieved.
        Assert.NotNull(actualPrescriptions);

        // Compare the count first to quickly catch mismatches.
        Assert.Equal(expectedPrescriptions.Count, actualPrescriptions.Count);

        foreach (var expectedPrescription in expectedPrescriptions)
        {
            // Find the actual prescription by Medication, as IDs may not be predictable.
            var actualPrescription = actualPrescriptions.First(p =>
                (p.Medication == expectedPrescription.Medication
                && p.PrescribedAt == expectedPrescription.PrescribedAt));

            Assert.NotNull(actualPrescription);
            Assert.Equal(expectedPrescription.Medication, actualPrescription.Medication);
            Assert.Equal(expectedPrescription.Doseage, actualPrescription.Doseage);
            Assert.Equal(expectedPrescription.Notes, actualPrescription.Notes);
            Assert.Equal(expectedPrescription.PrescribedAt, actualPrescription.PrescribedAt);
        }
    }

    [Fact]
    public void UpdatePrescription_UpdatesPrescription_WhenValidInputProvided()
    {
        // Arrange
        Guid prescriptionId = Guid.NewGuid();
        var prescriptionToUpdate = new Prescription
        {
            Id = prescriptionId,
            Medication = "Original Medication",
            Doseage = "10 mg Daily",
            Notes = "Note1",
            PrescribedAt = DateTime.UtcNow
        };

        var updatedPrescriptionDTO = new PrescriptionDTO
        {
            Medication = "Updated Medication",
            Doseage = "20 mg Daily",
            Notes = "Updated Note",
            PrescribedAt = DateTime.UtcNow.AddDays(1)
        };

        var mockPrescriptionRepository = new Mock<IPrescriptionRepository>();
        var mockMapper = new Mock<IMapper>();

        mockPrescriptionRepository.Setup(repo => repo.GetPrescriptionById(prescriptionId))
            .Returns(prescriptionToUpdate);
        mockPrescriptionRepository.Setup(repo => repo.PrescriptionExistsByMedication(updatedPrescriptionDTO.Medication))
            .Returns((Prescription)null); // Medication doesn't exist

        var prescriptionService = new PrescriptionService(mockPrescriptionRepository.Object, mockMapper.Object);

        // Act
        prescriptionService.UpdatePrescription(prescriptionId, updatedPrescriptionDTO);

        // Assert
        Assert.Equal(updatedPrescriptionDTO.Medication, prescriptionToUpdate.Medication);
        Assert.Equal(updatedPrescriptionDTO.Doseage, prescriptionToUpdate.Doseage);
        Assert.Equal(updatedPrescriptionDTO.Notes, prescriptionToUpdate.Notes);
        Assert.Equal(updatedPrescriptionDTO.PrescribedAt, prescriptionToUpdate.PrescribedAt);
    }

    [Fact]
    public void UpdatePrescription_ThrowsConflictException_WhenMedicationExists()
    {
        // Arrange
        Guid prescriptionId = Guid.NewGuid();
        var prescriptionToUpdate = new Prescription
        {
            Id = prescriptionId,
            Medication = "Original Medication",
            Doseage = "10 mg Daily",
            Notes = "Note1",
            PrescribedAt = DateTime.UtcNow
        };

        var updatedPrescriptionDTO = new PrescriptionDTO
        {
            Medication = "Updated Medication",
            Doseage = "20 mg Daily",
            Notes = "Updated Note",
            PrescribedAt = DateTime.UtcNow.AddDays(1)
        };

        var mockPrescriptionRepository = new Mock<IPrescriptionRepository>();
        var mockMapper = new Mock<IMapper>();

        mockPrescriptionRepository.Setup(repo => repo.GetPrescriptionById(prescriptionId))
            .Returns(prescriptionToUpdate);
        mockPrescriptionRepository.Setup(repo => repo.PrescriptionExistsByMedication(updatedPrescriptionDTO.Medication))
            .Returns(new Prescription()); // Medication already exists

        var prescriptionService = new PrescriptionService(mockPrescriptionRepository.Object, mockMapper.Object);

        // Act and Assert
        var exception = Assert.Throws<ResourceConflictException>(() => prescriptionService.UpdatePrescription(prescriptionId, updatedPrescriptionDTO));
        Assert.Equal("Medication already exists", exception.Message);
    }

    [Fact]
    public void UpdatePrescription_ThrowsBadRequestException_WhenNoChangesFound()
    {
        // Arrange
        Guid prescriptionId = Guid.NewGuid();
        var prescriptionToUpdate = new Prescription
        {
            Id = prescriptionId,
            Medication = "Original Medication",
            Doseage = "10 mg Daily",
            Notes = "Note1",
            PrescribedAt = DateTime.UtcNow
        };

        var updatedPrescriptionDTO = new PrescriptionDTO
        {
            Medication = "Original Medication", // No changes
            Doseage = "10 mg Daily", // No changes
            Notes = "Note1", // No changes
            PrescribedAt = prescriptionToUpdate.PrescribedAt // No changes
        };

        var mockPrescriptionRepository = new Mock<IPrescriptionRepository>();
        var mockMapper = new Mock<IMapper>();

        mockPrescriptionRepository.Setup(repo => repo.GetPrescriptionById(prescriptionId))
            .Returns(prescriptionToUpdate);

        var prescriptionService = new PrescriptionService(mockPrescriptionRepository.Object, mockMapper.Object);

        // Act and Assert
        var exception = Assert.Throws<BadRequestException>(() => prescriptionService.UpdatePrescription(prescriptionId, updatedPrescriptionDTO));
        Assert.Equal("no changes found", exception.Message);
    }


}
