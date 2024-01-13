package apperror

type BadRequestError struct {
	Message string
	Code    int
}

func (e *BadRequestError) Error() string {
	return e.Message
}

func (e *BadRequestError) Is(tgt error) bool {
	target, ok := tgt.(*BadRequestError)
	if !ok {
		return false
	}
	return e.Code == target.Code
}
