package mocking

type mockMethod struct {
	methodName      string
	calledArguments []interface{}
	mockValue       *mockValue
	callCount       uint
}

func newMockMethod(method string) *mockMethod {
	return &mockMethod{
		methodName:      method,
		calledArguments: make([]interface{}, 0),
	}
}

func (m *mockMethod) addExpectedResult(expectedResult interface{}) {
	m.mockValue = newMockDataWithResult(expectedResult)
}

func (m *mockMethod) addExpectedInvoke(expectedInvoke InvokeFunc) {
	m.mockValue = newMockDataWithInvoke(expectedInvoke)
}

func (m *mockMethod) isCalled() bool {
	return m.callCount > 0
}

func (m *mockMethod) incrementCallCount() {
	m.callCount++
}

func (m *mockMethod) getMockValue() *mockValue {
	return m.mockValue
}

func (m *mockMethod) setCalledArguments(args ...interface{}) {
	m.calledArguments = args
}

func (m *mockMethod) getCalledArguments() []interface{} {
	return m.calledArguments
}
