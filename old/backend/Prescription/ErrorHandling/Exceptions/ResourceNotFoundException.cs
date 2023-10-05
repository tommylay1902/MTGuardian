
namespace prescription.ErrorHandling.Exceptions
{
    public class ResourceNotFoundException : Exception
    {
        public ResourceNotFoundException() : base("Resource not found.")
        {
        }
        public ResourceNotFoundException(string message) : base(message)
        {
        }
        public ResourceNotFoundException(string message, Exception innerException) : base(message, innerException)
        {
        }
    }
}
