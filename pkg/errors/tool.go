package errors

import (
	"errors"
)

func CastOrWrap(err error, orErrorCode ErrorCode) *Error {
	if err == nil {
		return nil
	}

	var targetErr *Error
	if errors.As(err, &targetErr) {
		return targetErr
	}

	return NewError(orErrorCode, err.Error())
}

func ContainByCode(err error, errorCode ErrorCode) bool {
	if err == nil {
		return false
	}

	var targetErr *Error
	var targetErrs *Errors
	switch {
	case errors.As(err, &targetErr):
		return EqualByCode(targetErr, errorCode)
	case errors.As(err, &targetErrs):
		return targetErrs.ContainsByCode(errorCode)
	default:
		return false
	}
}

func EqualByCode(err error, errorCode ErrorCode) bool {
	if err == nil {
		return false
	}

	var targetErr *Error
	ok := errors.As(err, &targetErr)
	if !ok {
		return false
	}

	return targetErr.Code() == errorCode
}
