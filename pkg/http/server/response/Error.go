package httpResponse

import "github.com/conacry/go-platform/pkg/errors"

var (
	ErrLoggerIsRequired        = errors.NewError("SYS", "Logger is required")
	ErrErrorResolverIsRequired = errors.NewError("SYS", "Error resolver is required")
)

var (
	ErrWriteResponse = errors.NewError("546e893f-001", "An error occurred at write http response")
)
