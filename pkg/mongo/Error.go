package mongo

import (
	"fmt"
	"github.com/conacry/go-platform/pkg/errors"
)

var (
	ErrMongoDBIsRequired = errors.NewError("SYS", "MongoDB is required")
	ErrLoggerIsRequired  = errors.NewError("SYS", "Logger is required")

	StartSessionErrorCode              = errors.ErrorCode("f69256a6-001")
	StartTransactionErrorCode          = errors.ErrorCode("f69256a6-002")
	AbortTransactionErrorCode          = errors.ErrorCode("f69256a6-003")
	CommitTransactionErrorCode         = errors.ErrorCode("f69256a6-004")
	DuplicateUniqueConstraintErrorCode = errors.ErrorCode("f69256a6-005")
)

func ErrStartSession(cause error) error {
	detailMsg := fmt.Sprintf("Unable start session: %s", cause)
	return errors.NewError(StartSessionErrorCode, detailMsg)
}

func ErrStartTransaction(cause error) error {
	detailMsg := fmt.Sprintf("Unable start transaction: %s", cause)
	return errors.NewError(StartTransactionErrorCode, detailMsg)
}

func ErrAbortTransaction(cause error, abortCause error) error {
	detailMsg := fmt.Sprintf("Unable abort transaction: %s, cause to abort: %s", cause, abortCause)
	return errors.NewError(AbortTransactionErrorCode, detailMsg)
}

func ErrCommitTransaction(cause error) error {
	detailMsg := fmt.Sprintf("Unable commit transaction: %s", cause)
	return errors.NewError(CommitTransactionErrorCode, detailMsg)
}

func ErrDuplicateUniqueConstraint(cause error) error {
	detailMsg := fmt.Sprintf("Duplicate unique constraint. Cause: %s", cause)
	return errors.NewError(DuplicateUniqueConstraintErrorCode, detailMsg)
}
