package errorsStub

import (
	"github.com/conacry/go-platform/pkg/errors"
	"github.com/conacry/go-platform/pkg/generator"
)

func GetError() *errors.Error {
	randomErrorCode := errors.ErrorCode(generator.RandomDefaultStr())
	randomErrorText := generator.RandomDefaultStr()
	return errors.NewError(randomErrorCode, randomErrorText)
}
