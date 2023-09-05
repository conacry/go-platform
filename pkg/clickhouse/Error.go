package clickhouse

import "github.com/conacry/go-platform/pkg/errors"

var (
	ErrClickHouseConfigIsRequired = errors.NewError("SYS", "ClickHouse config is required")
	ErrLoggerIsRequired           = errors.NewError("SYS", "Logger is required")
)
