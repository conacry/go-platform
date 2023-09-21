package mockingTest

import (
	"errors"
	mocking "github.com/conacry/go-platform/pkg/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type Result struct {
	str string
}

type MethodArg struct {
	count int
	str   string
}

type MockedStruct struct {
	*mocking.BaseMock
}

func (m *MockedStruct) MockedVoidMethod(arg interface{}) error {
	res, err := m.ProcessMethod(arg)
	if res != nil {
		panic(mocking.ErrIllegalResultForMock)
	}

	return err
}

func (m *MockedStruct) MockedReturnedMethod(arg interface{}) (*Result, error) {
	result, err := m.ProcessMethod(arg)
	if err != nil {
		return nil, err
	}

	if result != nil {
		res, ok := result.(*Result)
		if !ok {
			panic(mocking.ErrCannotCastResult)
		}

		return res, nil
	}

	panic(mocking.ErrImplementMe)
}

type BaseMockShould struct {
	suite.Suite
	mockedStruct *MockedStruct
}

func TestBaseMockShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(BaseMockShould))
}

func (s *BaseMockShould) SetupTest() {
	s.mockedStruct = &MockedStruct{
		BaseMock: mocking.NewBaseMock(mocking.Modes.Strict()),
	}
}

func (s *BaseMockShould) TestMock_NoMethodsAreMocked_Panic() {
	assert.Panics(s.T(), func() {
		_, _ = s.mockedStruct.MockedReturnedMethod(struct{}{})
	})
}

func (s *BaseMockShould) TestMock_MockValueIsNil_Panic() {
	s.mockedStruct.SetReturnsFor("MockedReturnedMethod", nil)

	assert.Panics(s.T(), func() {
		_, _ = s.mockedStruct.MockedReturnedMethod(struct{}{})
	})
}

func (s *BaseMockShould) TestSetReturnsFor_MethodReturnErrButSetAnotherType_Panic() {
	s.mockedStruct.SetReturnsFor("MockedVoidMethod", struct{}{})

	assert.Panics(s.T(), func() {
		_ = s.mockedStruct.MockedVoidMethod(struct{}{})
	})
}

func (s *BaseMockShould) TestSetReturnsFor_MockVoidMethodWithNoError_MockReturnNil() {
	s.mockedStruct.SetReturnsFor("MockedVoidMethod", mocking.NoError)

	err := s.mockedStruct.MockedVoidMethod(struct{}{})
	assert.NoError(s.T(), err)
}

func (s *BaseMockShould) TestSetReturnsFor_MockVoidMethodWithError_MockReturnErr() {
	expectedErr := errors.New("expected error")
	s.mockedStruct.SetReturnsFor("MockedVoidMethod", expectedErr)

	err := s.mockedStruct.MockedVoidMethod(struct{}{})
	assert.ErrorIs(s.T(), err, expectedErr)
}

func (s *BaseMockShould) TestSetReturnsFor_MockReturnedMethodWithValue_MockReturnValue() {
	expectedRes := &Result{str: "qwe"}
	s.mockedStruct.SetReturnsFor("MockedReturnedMethod", expectedRes)

	actualRes, err := s.mockedStruct.MockedReturnedMethod(struct{}{})
	assert.NoError(s.T(), err)
	assert.EqualValues(s.T(), expectedRes, actualRes)
}

func (s *BaseMockShould) TestSetReturnsFor_MockReturnedMethodWithErr_MockReturnErr() {
	expectedErr := errors.New("expected error")
	s.mockedStruct.SetReturnsFor("MockedReturnedMethod", expectedErr)

	actualRes, err := s.mockedStruct.MockedReturnedMethod(struct{}{})
	assert.Nil(s.T(), actualRes)
	assert.ErrorIs(s.T(), err, expectedErr)
}

func (s *BaseMockShould) TestSetReturnsFor_MockReturnedMethodWithFunc_ArgIsFilled() {
	expectedArg := &MethodArg{
		count: 10,
		str:   "asd",
	}

	var funcToInvoke mocking.InvokeFunc = func(args ...interface{}) error {
		require.Len(s.T(), args, 1)

		arg, ok := args[0].(*MethodArg)
		require.True(s.T(), ok)

		arg.count = expectedArg.count
		arg.str = expectedArg.str

		return nil
	}

	s.mockedStruct.SetReturnsFor("MockedVoidMethod", funcToInvoke)

	actualArg := &MethodArg{}
	err := s.mockedStruct.MockedVoidMethod(actualArg)
	require.NoError(s.T(), err)
	assert.EqualValues(s.T(), expectedArg, actualArg)
}
