package app

import (
	"fmt"
)

type errorType string

type Error struct {
	Type  errorType
	Inner error
}

func (e Error) Error() string {
	return fmt.Sprintf("app error (%v): %v", e.Type, e.Inner.Error())
}

func IsAppError(err error) bool {
	_, ok := err.(Error)
	return ok
}

const (
	Internal         = errorType("internal")
	NotFound         = errorType("not found")
	Aborted          = errorType("aborted")
	ValidationFailed = errorType("validation failed")
	InvalidInput     = errorType("invalid input")
)

func NewInternalError(inner error) Error {
	return Error{Type: Internal, Inner: inner}
}

func NewNotFoundError(inner error) Error {
	return Error{Type: NotFound, Inner: inner}
}

func NewAbortedError(inner error) Error {
	return Error{Type: Aborted, Inner: inner}
}

func NewValidationError(inner error) Error {
	return Error{Type: ValidationFailed, Inner: inner}
}

func NewInvalidInputError(inner error) Error {
	return Error{Type: InvalidInput, Inner: inner}
}
