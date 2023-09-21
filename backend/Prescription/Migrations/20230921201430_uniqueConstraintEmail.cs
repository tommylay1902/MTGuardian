using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace prescription.Migrations
{
    /// <inheritdoc />
    public partial class uniqueConstraintEmail : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateIndex(
                name: "IX_Prescriptions_Medication",
                table: "Prescriptions",
                column: "Medication",
                unique: true);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropIndex(
                name: "IX_Prescriptions_Medication",
                table: "Prescriptions");
        }
    }
}
