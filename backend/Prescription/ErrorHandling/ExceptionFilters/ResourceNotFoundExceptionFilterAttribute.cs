using System;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using prescription.ErrorHandling.Exceptions;

namespace prescription.ErrorHandling.ExceptionFilters
{
    public class ResourceNotFoundExceptionFilterAttribute : ExceptionFilterAttribute
    {
		
		public override void OnException(ExceptionContext context)
        {
            if (context.Exception is ResourceConflictException)
            {
                context.Result = new NotFoundObjectResult(new { error = context.Exception.Message });
                context.ExceptionHandled = true;
            }
        }
    }

}

