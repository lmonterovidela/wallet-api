package exceptions

import "fmt"

type ForbiddenException struct {
	Message error
}

func (e *ForbiddenException) Error() string {
	return e.Message.Error()
}

func NewForbiddenException(message string, args ...string) error {
	return &ForbiddenException{Message: fmt.Errorf(message, args)}
}
