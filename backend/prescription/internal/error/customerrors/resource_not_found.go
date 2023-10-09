package customerrors

type ResourceNotFound struct {
	Message string
	Code    int
}

func (e *ResourceNotFound) Error() string {
	return e.Message
}

func (e *ResourceNotFound) Is(tgt error) bool {
	target, ok := tgt.(*ResourceNotFound)
	if !ok {
		return false
	}
	return e.Code == target.Code
}
