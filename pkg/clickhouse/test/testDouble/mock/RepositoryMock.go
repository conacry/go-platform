package clickhouseMock

import (
	"context"
	mocking "github.com/conacry/go-platform/pkg/mock"
)

func GetRepository() *RepositoryMock {
	return &RepositoryMock{
		BaseMock: mocking.NewBaseMock(mocking.Modes.Strict()),
	}
}

type RepositoryMock struct {
	*mocking.BaseMock
}

func (m *RepositoryMock) Select(ctx context.Context, result any, query string, queryArgs map[string]any) error {
	_, err := m.ProcessMethod(ctx, result, query, queryArgs)
	return err
}
