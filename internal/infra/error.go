package infra

import (
	"fmt"
)

type errorType string

type Error struct {
	Type  errorType
	Inner error
}

func (e Error) Error() string {
	return fmt.Sprintf("infra error (%v): %v", e.Type, e.Inner.Error())
}

func IsInfraError(err error) bool {
	_, ok := err.(Error)
	return ok
}

const (
	Http     = errorType("http failure")
	Encoding = errorType("encoding failed")
	Decoding = errorType("decoding failed")
)

func NewHttpError(inner error) Error {
	return Error{Type: Http, Inner: inner}
}

func NewEncodingError(inner error) Error {
	return Error{Type: Encoding, Inner: inner}
}

func NewDecodingError(inner error) Error {
	return Error{Type: Decoding, Inner: inner}
}
