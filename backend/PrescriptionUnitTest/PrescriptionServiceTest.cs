using Moq;
using prescription.Entities;
using prescription.Interfaces;
using prescription.ServicesLayer;

public class PrescriptionServiceTest
{
    [Fact]
    public void CreatePrescription_ValidInput_ReturnsGuid()
    {
        // Arrange
        var prescriptionRepositoryMock = new Mock<IPrescriptionRepository>();

        var prescriptionService = new PrescriptionService(prescriptionRepositoryMock.Object);

        var expectedGuid = Guid.NewGuid();
        prescriptionRepositoryMock.Setup(repo => repo.Add(It.IsAny<Prescription>()))
            .Returns(expectedGuid);

        var prescription = new Prescription
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

    //[Fact]
    //public void CreatePrescription_ValidInput_ReturnsGuid()
    //{
    //    // Arrange
    //    var prescriptionRepositoryMock = new Mock<IPrescriptionRepository>();

    //    var prescriptionService = new PrescriptionService(prescriptionRepositoryMock.Object);

    //    var expectedGuid = Guid.NewGuid();
    //    prescriptionRepositoryMock.Setup(repo => repo.Add(It.IsAny<Prescription>()))
    //        .Returns(expectedGuid);

    //    var prescription = new Prescription
    //    {
    //        Medication = "Dexamethasone",
    //        Doseage = "20 mg Daily",
    //        Notes = "A steroid used to help inflamed areas of the body",
    //        PrescribedAt = DateTime.Parse("2023-09-21T00:59:37.942Z").ToUniversalTime()
    //    };

    //    // Act
    //    var result = prescriptionService.CreatePrescription(prescription);

    //    // Assert
    //    Assert.Equal(expectedGuid, result);
    //}
}
