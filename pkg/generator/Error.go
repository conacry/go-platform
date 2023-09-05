package generator

import (
	"fmt"
	"github.com/conacry/go-platform/pkg/errors"
)

var (
	ParseUUIDErrorCode = errors.ErrorCode("12a8df9e-001")
)

func ErrParseUUID(invalidUUIDStr string, cause error) error {
	errMsg := fmt.Sprintf("Fail create uuid from string = %q. Cause: %q", invalidUUIDStr, cause.Error())
	return errors.NewError(ParseUUIDErrorCode, errMsg)
}
