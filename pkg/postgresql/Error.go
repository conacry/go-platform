package postgresql

import "github.com/conacry/go-platform/pkg/errors"

var (
	ErrConfigIsRequired = errors.NewError("SYS", "PostgreSQL config is required")
	ErrLoggerIsRequired = errors.NewError("SYS", "Logger is required")
)
