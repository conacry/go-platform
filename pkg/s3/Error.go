package s3Storage

import "github.com/conacry/go-platform/pkg/errors"

var (
	ErrS3ConfigIsRequired = errors.NewError("SYS", "S3 config is required")
	ErrLoggerIsRequired   = errors.NewError("SYS", "Logger is required")
)
