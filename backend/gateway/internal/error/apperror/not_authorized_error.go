package apperror

type NotAuthorizedError struct {
	Message string
	Code    int `default:"401"`
}

func (e *NotAuthorizedError) Error() string {
	return e.Message
}

func (e *NotAuthorizedError) Is(tgt error) bool {
	target, ok := tgt.(*NotAuthorizedError)
	if !ok {
		return false
	}
	return e.Code == target.Code
}
