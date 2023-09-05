package errors

import (
	"errors"
	"strings"
)

type Errors struct {
	errors []*Error
}

func NewErrors() *Errors {
	return &Errors{
		errors: make([]*Error, 0),
	}
}

func ErrorsFrom(errs ...error) *Errors {
	commonErrs := NewErrors()
	commonErrs.AddErrors(errs)

	return commonErrs
}

func (e *Errors) AddError(err error) {
	var targetErr *Error
	var targetErrs *Errors
	switch {
	case errors.As(err, &targetErr):
		e.errors = append(e.errors, targetErr)
	case errors.As(err, &targetErrs):
		e.errors = append(e.errors, targetErrs.errors...)
	default:
		commonErr := ErrorFrom(err)
		e.errors = append(e.errors, commonErr)
	}
}

func (e *Errors) CreateAndAddError(code ErrorCode, msg string) {
	e.AddError(NewError(code, msg))
}

func (e *Errors) AddErrors(errs []error) {
	for _, err := range errs {
		e.AddError(err)
	}
}

func (e *Errors) IsEmpty() bool {
	return len(e.errors) == 0
}

func (e *Errors) IsPresent() bool {
	return len(e.errors) > 0
}

func (e *Errors) Size() int {
	return len(e.errors)
}

func (e *Errors) Contains(err error) bool {
	if err == nil {
		return false
	}

	commonErr, isCommonErr := err.(*Error)
	if !isCommonErr {
		return false
	}

	for _, errorItem := range e.errors {
		if errorItem.Equals(commonErr) {
			return true
		}
	}

	return false
}

func (e *Errors) ContainsByCode(errorCode ErrorCode) bool {
	for _, v := range e.errors {
		if v.Code() == errorCode {
			return true
		}
	}

	return false
}

func (e *Errors) Error() string {
	errStrings := make([]string, 0)
	for _, e := range e.errors {
		errStrings = append(errStrings, e.Error())
	}
	return strings.Join(errStrings, "\n")
}

func (e *Errors) ToArray() []*Error {
	errsCopy := make([]*Error, len(e.errors))
	copy(errsCopy, e.errors)
	return errsCopy
}
