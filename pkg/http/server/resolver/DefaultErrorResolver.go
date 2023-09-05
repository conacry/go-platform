package httpErrorResolver

import (
	goErrors "errors"
	"github.com/conacry/go-platform/pkg/errors"
	httpRequest "github.com/conacry/go-platform/pkg/http/server/request"
	"net/http"
)

const (
	UnknownErrorCode = "UNKNOWN_CODE"
)

type DefaultErrorResolver struct{}

func NewDefaultErrorResolver() *DefaultErrorResolver {
	return &DefaultErrorResolver{}
}

func (es *DefaultErrorResolver) GetErrorCode(err error) string {
	switch errT := err.(type) {
	case *errors.Error:
		return errT.Code().String()
	default:
		return UnknownErrorCode
	}
}

func (es *DefaultErrorResolver) GetErrorText(err error) string {
	switch errT := err.(type) {
	case *errors.Error:
		return errT.Message()
	default:
		return err.Error()
	}
}

func (es *DefaultErrorResolver) GetHttpCode(err error) int {
	var customErr *errors.Error
	ok := goErrors.As(err, &customErr)
	if !ok {
		return http.StatusInternalServerError
	}

	switch customErr.Code() {
	case httpRequest.UnmarshalRequestErrorCode:
		return http.StatusBadRequest
	default:
		return http.StatusUnprocessableEntity
	}
}
