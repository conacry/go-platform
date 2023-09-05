package errors

import "fmt"

const UnknownErrorCode = ErrorCode("unknown_error")

type ErrorCode string

func (c ErrorCode) String() string {
	return string(c)
}

type Error struct {
	code ErrorCode
	msg  string
}

func NewError(code ErrorCode, msg string) *Error {
	return &Error{
		code: code,
		msg:  msg,
	}
}

func ErrorFrom(err error) *Error {
	return CastOrWrap(err, UnknownErrorCode)
}

func (e *Error) Code() ErrorCode {
	return e.code
}

func (e *Error) Message() string {
	return e.msg
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s.", e.code, e.msg)
}

func (e *Error) Equals(err *Error) bool {
	if err == nil {
		return false
	}

	return e.code == err.code && e.msg == err.msg
}
