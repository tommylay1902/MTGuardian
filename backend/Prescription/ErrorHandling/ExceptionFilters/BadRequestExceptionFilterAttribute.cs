
using System.Net;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using prescription.ErrorHandling.Exceptions;

namespace prescription.ErrorHandling.ExceptionFilters
{
    public class BadRequestExceptionFilterAttribute : ExceptionFilterAttribute
    {

        public override void OnException(ExceptionContext context)
        {
            if (context.Exception is BadRequestException)
            {
                context.Result = new ObjectResult(new { error = context.Exception.Message })
                {
                    StatusCode = (int)HttpStatusCode.BadRequest
                };
                context.ExceptionHandled = true;
            }
        }

    }
}

