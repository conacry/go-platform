package httpController

import (
	"fmt"
	"github.com/conacry/go-platform/pkg/errors"
)

var (
	ErrLoggerIsRequired              = errors.NewError("SYS", "Logger is required")
	ErrResponseWriterIsRequired      = errors.NewError("SYS", "ResponseWriter is required")
	ErrErrorResponseWriterIsRequired = errors.NewError("SYS", "Error response writer is required")
)

var (
	UnmarshalRequestErrorCode errors.ErrorCode = "316ad077-001"
)

func ErrUnmarshalRequest(causeDescription string) error {
	errMsg := fmt.Sprintf("Malformed request. Cause - %s", causeDescription)
	return errors.NewError(UnmarshalRequestErrorCode, errMsg)
}
