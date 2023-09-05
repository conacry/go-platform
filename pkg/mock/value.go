package mocking

type InvokeFunc func(args ...interface{}) error

type mockValue struct {
	expectedResult     interface{}
	expectedInvokeFunc InvokeFunc
}

func newMockDataWithResult(expectedResult interface{}) *mockValue {
	return &mockValue{
		expectedResult: expectedResult,
	}
}

func newMockDataWithInvoke(expectedInvoke InvokeFunc) *mockValue {
	return &mockValue{
		expectedInvokeFunc: expectedInvoke,
	}
}

func (m *mockValue) isInvokable() bool {
	return m.expectedInvokeFunc != nil
}

func (m *mockValue) isReturnable() bool {
	return m.expectedResult != nil
}
