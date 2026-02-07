package parco

import (
	"errors"
	"fmt"
)

var (
	ErrNotIntegerType    = errors.New("not an integer type")
	ErrOverflow          = errors.New("bytes overflow")
	ErrCannotRead        = errors.New("unsufficient bytes read")
	ErrCannotWrite       = errors.New("unsufficient bytes written")
	ErrAlreadyRegistered = errors.New("builder is registered already")
	ErrUnknownType       = errors.New("unknown type")
	ErrInvalidLength     = errors.New("invalid length")
)

type ErrUnSufficientBytes struct {
	want int
	have int
}

func (e ErrUnSufficientBytes) Error() string {
	return fmt.Sprintf("unsufficient bytes. want %d have %d",
		e.want, e.have)
}

func NewErrUnSufficientBytesError(want, have int) ErrUnSufficientBytes {
	return ErrUnSufficientBytes{
		want: want,
		have: have,
	}
}

type ErrFieldNotFound struct {
	field string
}

func (e ErrFieldNotFound) Error() string {
	return "field not found " + e.field
}

func NewErrFieldNotFoundError(field string) ErrFieldNotFound {
	return ErrFieldNotFound{field: field}
}

type ErrTypeAssertion struct {
	expectedType string
	actualType   string
}

func (e ErrTypeAssertion) Error() string {
	return fmt.Sprintf("unexpected type, want %s but have %s",
		e.expectedType, e.actualType)
}

func NewErrTypeAssertionError(want, have string) ErrTypeAssertion {
	return ErrTypeAssertion{expectedType: want, actualType: have}
}

type ErrCompile struct {
	reason string
}

func (e ErrCompile) Error() string {
	return e.reason
}

func NewErrCompile(reason string) ErrCompile {
	return ErrCompile{reason: reason}
}
