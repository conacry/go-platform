package mongo

import (
	"fmt"
	"github.com/conacry/go-platform/pkg/errors"
)

var (
	ErrLoggerIsRequired = errors.NewError("SYS", "Logger is required")
	ErrConfigIsRequired = errors.NewError("SYS", "Config is required")
)

var (
	DuplicateUniqueConstraintErrorCode = errors.ErrorCode("a15da443-001")
)

func ErrDuplicateUniqueConstraint(cause error) error {
	detailMsg := fmt.Sprintf("Duplicate unique constraint. Cause: %s", cause)
	return errors.NewError(DuplicateUniqueConstraintErrorCode, detailMsg)
}
