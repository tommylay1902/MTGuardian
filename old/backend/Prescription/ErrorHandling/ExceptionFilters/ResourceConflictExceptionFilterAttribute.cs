using System.Net;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using prescription.ErrorHandling.Exceptions;

namespace prescription.ErrorHandling.ExceptionFilters
{
    public class ResourceConflictExceptionFilterAttribute : ExceptionFilterAttribute
	{
		public override void OnException(ExceptionContext context)
        {
            if (context.Exception is ResourceConflictException)
            {
                context.Result = new ObjectResult(new { error = context.Exception.Message })
                {
                    StatusCode = (int)HttpStatusCode.Conflict
                };
                context.ExceptionHandled = true;
            }
        }
    }
}

