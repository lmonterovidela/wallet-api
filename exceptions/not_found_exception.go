package exceptions

import "fmt"

type NotFoundException struct {
	Message error
}

func (e *NotFoundException) Error() string {
	return e.Message.Error()
}

func NewNotFoundException(message string, args ...string) error {
	return &NotFoundException{Message: fmt.Errorf(message, args)}
}
