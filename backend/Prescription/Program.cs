using prescription.Data;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using prescription.Interfaces;
using prescription.ServicesLayer;
using prescription.Repositories;
using prescription.ErrorHandling.ExceptionsFilters;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Services.AddDbContext<PrescriptionContext>(options =>
    options.UseNpgsql(builder.Configuration.GetConnectionString("DefaultConnection")
));

// Add automapper
builder.Services.AddAutoMapper(typeof(Program).Assembly);

// Add exception filters
builder.Services.AddSingleton<ResourceNotFoundExceptionFilterAttribute>();

// Add repos and services
builder.Services.AddScoped<IPrescriptionRepository, PrescriptionRepository>();
builder.Services.AddScoped<IPrescriptionService, PrescriptionService>();


var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
