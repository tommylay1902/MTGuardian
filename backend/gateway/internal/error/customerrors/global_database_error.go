package customerrors

type GlobalDatabaseError struct {
	Message string
	Code    int `default:"500"`
}

func (e *GlobalDatabaseError) Error() string {
	return e.Message
}

func (e *GlobalDatabaseError) Is(tgt error) bool {
	target, ok := tgt.(*GlobalDatabaseError)
	if !ok {
		return false
	}
	return e.Code == target.Code
}
