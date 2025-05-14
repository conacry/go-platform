package mocking

import (
	"errors"
	"fmt"
	"regexp"
	"runtime"
)

const (
	NoError = "no_error_mock_result"
)

var (
	ErrAssertMockResult     = errors.New("an error occurred at asserting mock result")
	ErrCannotCastResult     = errors.New("cannot cast expected mock result")
	ErrImplementMe          = errors.New("unknown returnFor in mock. Need some result")
	ErrIllegalResultForMock = errors.New("illegal result for a mocked method")
)

type BaseMock struct {
	mode        Mode
	mockMethods map[string]*mockMethod
}

func NewBaseMock(mode Mode) *BaseMock {
	if mode == "" {
		mode = Modes.Strict()
	}

	return &BaseMock{
		mode:        mode,
		mockMethods: map[string]*mockMethod{},
	}
}

func (m *BaseMock) Reset() {
	m.mockMethods = make(map[string]*mockMethod)
}

func (m *BaseMock) ProcessMethod(args ...any) (any, error) {
	methodName := getCurrentFuncName()

	m.initMockIfNotExists(methodName)
	mockedMethod, ok := m.mockMethods[methodName]
	if !ok {
		panic(fmt.Sprintf("there is no mocked method with name: %q", methodName))
	}

	mockedMethod.setCalledArguments(args...)
	mockedMethod.incrementCallCount()

	value := mockedMethod.getMockValue()
	if value == nil {
		if m.mode == Modes.Strict() {
			panic(fmt.Sprintf("behaviour for mock method = %q didn't set", methodName))
		}

		return nil, nil
	}

	switch {
	case value.isInvokable():
		return invokeMockedValue(value, args...)
	case value.isReturnable():
		return returnMockValue(value, methodName)
	default:
		panic(fmt.Sprintf("there is no mocked value to return or invoke for mockMethod: %q", methodName))
	}
}

func invokeMockedValue(value *mockValue, args ...any) (any, error) {
	err := value.expectedInvokeFunc(args...)
	return nil, err
}

func returnMockValue(value *mockValue, methodName string) (any, error) {
	expectedResult := value.expectedResult
	if expectedResult == nil {
		panic(fmt.Sprintf("expected result not found for process mock method = %q", methodName))
	}

	if expectedResult == NoError {
		return nil, nil
	}

	expectedError, isError := expectedResult.(error)
	if isError {
		return nil, expectedError
	} else {
		return expectedResult, nil
	}
}

func getCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(2)
	re := regexp.MustCompile(`\w+$`)
	return re.FindString(runtime.FuncForPC(pc).Name())
}

func (m *BaseMock) GetCountCallsFor(method string) uint {
	mockedMethod, exists := m.mockMethods[method]
	if !exists {
		return 0
	} else {
		return mockedMethod.callCount
	}
}

func (m *BaseMock) IsMethodCalled(method string) bool {
	mockedMethod, exists := m.mockMethods[method]
	return exists && mockedMethod.isCalled()
}

func (m *BaseMock) IsNeverCalled() bool {
	for _, mockedMethod := range m.mockMethods {
		if mockedMethod.isCalled() {
			return false
		}
	}

	return true
}

func (m *BaseMock) SetReturnsFor(methodName string, mockValue any) {
	mockedMethod := newMockMethod(methodName)
	m.addMockValueToMethod(mockedMethod, mockValue)
	m.mockMethods[methodName] = mockedMethod
}

func (m *BaseMock) GetCalledArgs(method string) []any {
	return m.mockMethods[method].getCalledArguments()
}

func (m *BaseMock) addMockValueToMethod(mockedMethod *mockMethod, mockValue any) {
	switch value := mockValue.(type) {
	case InvokeFunc:
		mockedMethod.addExpectedInvoke(value)
	default:
		mockedMethod.addExpectedResult(value)
	}
}

func (m *BaseMock) initMockIfNotExists(method string) {
	if _, exists := m.mockMethods[method]; !exists {
		m.mockMethods[method] = newMockMethod(method)
	}
}
