package httpServer

import "github.com/conacry/go-platform/pkg/errors"

var (
	ErrLoggerIsRequired     = errors.NewError("SYS", "Logger is required")
	ErrHttpConfigIsRequired = errors.NewError("SYS", "HttpConfig is required")
)
