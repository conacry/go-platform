package mongoModel

import "github.com/conacry/go-platform/pkg/errors"

var (
	ErrCollectionValueIsRequired = errors.NewError("135bc16c-001", "Value for mongo collection is required")
)
