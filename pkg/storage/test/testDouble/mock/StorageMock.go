package storageMock

import (
	"context"
	mocking "github.com/conacry/go-platform/pkg/mock"
	storageModel "github.com/conacry/go-platform/pkg/storage/model"
)

func GetStorage() *StorageMock {
	return &StorageMock{
		BaseMock: mocking.NewBaseMock(),
	}
}

type StorageMock struct {
	*mocking.BaseMock
}

func (m *StorageMock) UploadFile(ctx context.Context, file *storageModel.File) error {
	_, err := m.ProcessMethod(ctx, file)
	return err
}

func (m *StorageMock) GetFile(ctx context.Context, scope, path string) (*storageModel.File, error) {
	result, err := m.ProcessMethod(ctx, scope, path)
	if err != nil {
		return nil, err
	}

	if result != nil {
		file, ok := result.(*storageModel.File)
		if !ok {
			panic(mocking.ErrCannotCastResult)
		}

		return file, nil
	}

	panic(mocking.ErrImplementMe)
}

func (m *StorageMock) RemoveFile(ctx context.Context, scope, path string) error {
	_, err := m.ProcessMethod(ctx, scope, path)
	return err
}

func (m *StorageMock) GetFileMetaData(ctx context.Context, scope, path string) (*storageModel.FileMetaData, error) {
	result, err := m.ProcessMethod(ctx, scope, path)
	if err != nil {
		return nil, err
	}

	if result != nil {
		metaData, ok := result.(*storageModel.FileMetaData)
		if !ok {
			panic(mocking.ErrCannotCastResult)
		}

		return metaData, nil
	}

	panic(mocking.ErrImplementMe)
}
