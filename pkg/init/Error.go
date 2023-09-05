package initapp

import "github.com/conacry/go-platform/pkg/errors"

var (
	ErrDependencyInitChainIsRequired = errors.NewError("SYS", "Chain with dependencies is required")
)
