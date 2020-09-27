
package object_models

import (
	"fmt"
)

type Error struct {
	Service ServiceName
	Err     error
	Message string
}

func NewError(service ServiceName, err error, message string) *Error {
	return &Error{Service: service, Err: err, Message: message}
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%s]%s: %s", e.Service, e.Message, e.Err)
}
