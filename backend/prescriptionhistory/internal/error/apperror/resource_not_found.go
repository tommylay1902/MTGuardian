package apperror

import "fmt"

type ResourceNotFound struct {
	Message string
	Code    int
}

func (e *ResourceNotFound) Error() string {
	return fmt.Sprintf("Resource not found: %s (Code: %d)", e.Message, e.Code)
}

func (e *ResourceNotFound) Is(tgt error) bool {

	target, ok := tgt.(*ResourceNotFound)
	if !ok {
		return false
	}
	return e.Code == target.Code
}
