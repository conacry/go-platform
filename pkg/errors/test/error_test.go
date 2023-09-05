package errorsTest

import (
	goerr "errors"
	"github.com/conacry/go-platform/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ErrorShould struct {
	suite.Suite
}

func TestErrorShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ErrorShould))
}

func (c *ErrorShould) TestNewError_AllParams_ReturnError() {
	expectedCode := errors.ErrorCode("CODE333")
	expectedMsg := "message"

	err := errors.NewError(expectedCode, expectedMsg)

	require.NotNil(c.T(), err)
	assert.Equal(c.T(), expectedCode, err.Code())
	assert.Equal(c.T(), expectedMsg, err.Message())
}

func (c *ErrorShould) TestNewErrorFrom_FromCommonError_ReturnError() {
	expectedError := errors.NewError("CODE333", "message")
	err := errors.ErrorFrom(expectedError)
	require.NotNil(c.T(), err)
	assert.Equal(c.T(), expectedError, err)
}

func (c *ErrorShould) TestNewErrorFrom_FromError_ReturnError() {
	originalError := goerr.New("123")
	err := errors.ErrorFrom(originalError)
	require.NotNil(c.T(), err)

	assert.NotEqual(c.T(), originalError, err)
	assert.Equal(c.T(), errors.UnknownErrorCode, err.Code())
	assert.Equal(c.T(), originalError.Error(), err.Message())
}

func (c *ErrorShould) TestNewError_Valid_ReturnEmptyErrors() {
	errs := errors.NewErrors()
	assert.NotNil(c.T(), errs)
	assert.True(c.T(), errs.IsEmpty())
}

func (c *ErrorShould) TestNewErrorsFrom_WithoutParams_ReturnErrors() {
	errs := errors.ErrorsFrom()

	require.NotNil(c.T(), errs)
	assert.True(c.T(), errs.IsEmpty())
}

func (c *ErrorShould) TestNewErrorsFrom_SingleError_ReturnErrors() {
	expectedError := errors.NewError("C", "msg")
	errs := errors.ErrorsFrom(expectedError)

	require.NotNil(c.T(), errs)
	assert.Equal(c.T(), 1, errs.Size())
	assert.True(c.T(), errs.Contains(expectedError))
	assert.False(c.T(), errs.ContainsByCode(errors.UnknownErrorCode))
}

func (c *ErrorShould) TestNewErrorsFrom_SeveralErrors_ReturnErrors() {
	expectedError1 := errors.NewError("C", "msg")
	expectedError2 := goerr.New("123")
	errs := errors.ErrorsFrom(expectedError1, expectedError2)

	require.NotNil(c.T(), errs)
	assert.Equal(c.T(), 2, errs.Size())
	assert.True(c.T(), errs.Contains(expectedError1))
	assert.False(c.T(), errs.Contains(expectedError2))
	assert.True(c.T(), errs.ContainsByCode(errors.UnknownErrorCode))
}

func (c *ErrorShould) TestCanAddErrors() {
	errs := errors.NewErrors()

	errs2 := errors.NewErrors()
	errs2.AddError(errors.NewError("C", "msg1"))

	errs.AddError(errs2)
	errs.CreateAndAddError("B", "msg2")

	assert.NotNil(c.T(), errs)
	assert.True(c.T(), !errs.IsEmpty())
	assert.Equal(c.T(), 2, errs.Size())
	assert.Equal(c.T(), 1, errs2.Size())
}

func (c *ErrorShould) TestEqualByCode() {
	expectedErrorCode := errors.ErrorCode("CODE333")
	var expectedError error = errors.NewError(expectedErrorCode, "message")

	assert.True(c.T(), errors.EqualByCode(expectedError, expectedErrorCode))
	assert.False(c.T(), errors.EqualByCode(expectedError, "CODE444"))
}

func (c *ErrorShould) TestContainByCode() {
	var expectedError1 = errors.NewError("CODE3", "message1")
	var expectedError2 = errors.NewError("CODE4", "message2")
	var expectedError3 = errors.NewError("CODE5", "message3")
	errs := errors.NewErrors()
	errs.AddError(expectedError1)
	errs.AddError(expectedError2)
	errs.AddError(expectedError3)

	assert.True(c.T(), errors.ContainByCode(errs, expectedError1.Code()))
	assert.False(c.T(), errors.ContainByCode(errs, "45354"))

	assert.True(c.T(), errors.ContainByCode(expectedError1, expectedError1.Code()))
	assert.False(c.T(), errors.ContainByCode(expectedError1, expectedError2.Code()))
}

func (c *ErrorShould) TestCastOrWrap() {
	var expectedError1 error = errors.NewError("CODE1", "message1")
	var wrapErrorCode errors.ErrorCode = "CODE2"

	actualError := errors.CastOrWrap(expectedError1, wrapErrorCode)
	assert.Equal(c.T(), expectedError1, actualError)

	var expectedError2 = goerr.New("unknown error")

	actualError = errors.CastOrWrap(expectedError2, wrapErrorCode)
	assert.True(c.T(), errors.EqualByCode(actualError, wrapErrorCode))
	assert.Equal(c.T(), expectedError2.Error(), actualError.Message())
}
