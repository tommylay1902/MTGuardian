package apperror

type ResourceConflictError struct {
	Message string
	Code    int `default:"409"`
}

func (e *ResourceConflictError) Error() string {
	return e.Message
}

func (e *ResourceConflictError) Is(tgt error) bool {
	target, ok := tgt.(*ResourceConflictError)
	if !ok {
		return false
	}
	return e.Code == target.Code
}
