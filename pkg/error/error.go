package error

type Error struct {
	Code    Code
	Message string
	Cause   error
}

func (e *Error) Error() string {
	return e.Message
}
