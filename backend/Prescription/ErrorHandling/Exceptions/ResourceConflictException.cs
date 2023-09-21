using System;
namespace prescription.ErrorHandling.Exceptions
{
	public class ResourceConflictException: Exception
	{
        public ResourceConflictException() : base("Resource not found.")
        {
        }

        public ResourceConflictException(string message) : base(message)
        {
        }

        public ResourceConflictException(string message, Exception innerException) : base(message, innerException)
        {
        }
    }
}

