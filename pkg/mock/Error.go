package mocking

import (
	"fmt"
	"github.com/conacry/go-platform/pkg/errors"
)

var (
	UnsupportedModeErrorCode = errors.ErrorCode("af6cbd01-001")
)

func ErrUnsupportedMode(modeStr string) error {
	errMsg := fmt.Sprintf("Unsupported Mode = %q", modeStr)
	return errors.NewError(UnsupportedModeErrorCode, errMsg)
}
