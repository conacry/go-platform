package storage

import (
	"context"
	storageModel "github.com/conacry/go-platform/pkg/storage/model"
)

type Storage interface {
	UploadFile(ctx context.Context, file *storageModel.File) error
	GetFile(ctx context.Context, scope, path string) (*storageModel.File, error)
	RemoveFile(ctx context.Context, scope, path string) error
	GetFileMetaData(ctx context.Context, scope, path string) (*storageModel.FileMetaData, error)
}
