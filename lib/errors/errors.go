package errors

// Error ...
type Error struct {
	message string
}

func New(message string) Error {
	return Error{message: message}
}

func (e Error) Message() string {
	return e.message
}
