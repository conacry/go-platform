package logMock

import (
	"context"
	mocking "github.com/conacry/go-platform/pkg/mock"
)

type LoggerMock struct {
	*mocking.BaseMock
}

func GetLogger() *LoggerMock {
	return &LoggerMock{
		BaseMock: mocking.NewBaseMock(),
	}
}

func (m *LoggerMock) LogError(ctx context.Context, errs ...error) {
	_, err := m.ProcessMethod(ctx, errs)
	if err != nil {
		panic(err)
	}
}

func (m *LoggerMock) LogWarn(ctx context.Context, messages ...string) {
	_, err := m.ProcessMethod(ctx, messages)
	if err != nil {
		panic(err)
	}
}

func (m *LoggerMock) LogInfo(ctx context.Context, messages ...string) {
	_, err := m.ProcessMethod(ctx, messages)
	if err != nil {
		panic(err)
	}
}

func (m *LoggerMock) LogDebug(ctx context.Context, messages ...string) {
	_, err := m.ProcessMethod(ctx, messages)
	if err != nil {
		panic(err)
	}
}

func (m *LoggerMock) GetErrorsFromCalledArgs() []error {
	if !m.IsMethodCalled("LogError") {
		panic("LogError not called")
	}

	logArgs := m.GetCalledArgs("LogError")
	if len(logArgs) != 2 {
		panic("length args for method 'LogError' is not equals 2")
	}

	logErrors, ok := logArgs[1].([]error)
	if !ok {
		panic("error in called arg not have is '[]error'")
	}

	return logErrors
}

func (m *LoggerMock) GetErrorFromCalledArgs() error {
	logErrors := m.GetErrorsFromCalledArgs()
	if len(logErrors) != 1 {
		panic("length errors in arg for method 'LogError' is not equals 1")
	}

	return logErrors[0]
}
