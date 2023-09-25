package mongoMock

import (
	"context"
	mocking "github.com/conacry/go-platform/pkg/mock"
	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetRepository() *RepositoryMock {
	return &RepositoryMock{
		BaseMock: mocking.NewBaseMock(
			mocking.Modes.Strict(),
		),
	}
}

type RepositoryMock struct {
	*mocking.BaseMock
}

func (m *RepositoryMock) Insert(ctx context.Context, collectionName string, data interface{}) (string, error) {
	result, err := m.ProcessMethod(ctx, collectionName, data)
	if err != nil {
		return "", err
	}

	if result != nil {
		id, ok := result.(string)
		if !ok {
			return "", mocking.ErrAssertMockResult
		}

		return id, nil
	}

	return "", err
}

func (m *RepositoryMock) InsertMany(ctx context.Context, collectionName string, data []interface{}) ([]string, error) {
	result, err := m.ProcessMethod(ctx, collectionName, data)
	if err != nil {
		return nil, err
	}

	if result != nil {
		ids, ok := result.([]string)
		if !ok {
			return nil, mocking.ErrAssertMockResult
		}

		return ids, nil
	}

	return nil, err
}

func (m *RepositoryMock) FindOneAndUpdate(ctx context.Context, collectionName string, resultModel, filter, updateData interface{}, opt *options.FindOneAndUpdateOptions) error {
	_, err := m.ProcessMethod(ctx, collectionName, resultModel, filter, updateData, opt)
	return err
}

func (m *RepositoryMock) ReplaceOne(ctx context.Context, collectionName string, filter, data interface{}) error {
	_, err := m.ProcessMethod(ctx, collectionName, filter, data)
	return err
}

func (m *RepositoryMock) UpdateOne(ctx context.Context, collectionName string, filter, data interface{}, opts ...*options.UpdateOptions) (int64, error) {
	result, err := m.ProcessMethod(ctx, collectionName, filter, data, opts)
	if err != nil {
		return 0, err
	}

	if result != nil {
		count, ok := result.(int64)
		if !ok {
			return 0, mocking.ErrAssertMockResult
		}

		return count, nil
	}

	return 0, err
}

func (m *RepositoryMock) UpdateMany(
	ctx context.Context,
	collectionName string,
	filter interface{},
	data interface{},
	opts ...*options.UpdateOptions,
) (int64, error) {
	result, err := m.ProcessMethod(ctx, collectionName, filter, data, opts)
	if err != nil {
		return 0, err
	}

	if result != nil {
		count, ok := result.(int64)
		if !ok {
			return 0, mocking.ErrAssertMockResult
		}

		return count, nil
	}

	return 0, err
}

func (m *RepositoryMock) Find(ctx context.Context, collectionName string, results, find interface{}, opt *options.FindOptions) error {
	_, err := m.ProcessMethod(ctx, collectionName, results, find, opt)
	return err
}

func (m *RepositoryMock) FindOne(ctx context.Context, collectionName string, resultModel, findQuery interface{}, findOptions *options.FindOneOptions) error {
	_, err := m.ProcessMethod(ctx, collectionName, resultModel, findQuery, findOptions)
	return err
}

func (m *RepositoryMock) DeleteOne(ctx context.Context, collectionName string, filter interface{}, opt *options.DeleteOptions) (*mongo.DeleteResult, error) {
	result, err := m.ProcessMethod(ctx, collectionName, filter, opt)
	if err != nil {
		return nil, err
	}

	if result != nil {
		deleteResult, ok := result.(*mongo.DeleteResult)
		if !ok {
			return nil, mocking.ErrAssertMockResult
		}

		return deleteResult, nil
	}

	return nil, err
}

func (m *RepositoryMock) DeleteMany(ctx context.Context, collectionName string, filter interface{}, opt *options.DeleteOptions) (*mongo.DeleteResult, error) {
	result, err := m.ProcessMethod(ctx, collectionName, filter, opt)
	if err != nil {
		return nil, err
	}

	if result != nil {
		deleteResult, ok := result.(*mongo.DeleteResult)
		if !ok {
			return nil, mocking.ErrAssertMockResult
		}

		return deleteResult, nil
	}

	return nil, err
}

func (m *RepositoryMock) Count(ctx context.Context, collectionName string, find interface{}, opt *options.CountOptions) (int64, error) {
	result, err := m.ProcessMethod(ctx, collectionName, find, opt)
	if err != nil {
		return 0, err
	}

	if result != nil {
		count, ok := result.(int64)
		if !ok {
			return 0, mocking.ErrAssertMockResult
		}

		return count, nil
	}

	return 0, err
}

func (m *RepositoryMock) Aggregate(ctx context.Context, collectionName string, pipe mongo.Pipeline) (*mongo.Cursor, error) {
	result, err := m.ProcessMethod(ctx, collectionName, pipe)
	if err != nil {
		return nil, err
	}

	if result != nil {
		cursor, ok := result.(*mongo.Cursor)
		if !ok {
			return nil, mocking.ErrAssertMockResult
		}

		return cursor, nil
	}

	return nil, err
}

func (m *RepositoryMock) CreateTextIndex(ctx context.Context, index *mongoModel.DBTextIndex) (string, error) {
	result, err := m.ProcessMethod(ctx, index)
	if err != nil {
		return "", err
	}

	if result != nil {
		indexName, ok := result.(string)
		if !ok {
			return "", mocking.ErrAssertMockResult
		}

		return indexName, nil
	}

	return "", err
}

func (m *RepositoryMock) CreateIndex(ctx context.Context, index *mongoModel.DBIndex) (string, error) {
	result, err := m.ProcessMethod(ctx, index)
	if err != nil {
		return "", err
	}

	if result != nil {
		indexName, ok := result.(string)
		if !ok {
			return "", mocking.ErrAssertMockResult
		}

		return indexName, nil
	}

	return "", err
}

func (m *RepositoryMock) TryCreateIndex(ctx context.Context, index *mongoModel.DBIndex) error {
	_, err := m.ProcessMethod(ctx, index)
	return err
}

func (m *RepositoryMock) CollectionIndexes(ctx context.Context, collection string) (map[string]*mongoModel.DBIndex, error) {
	result, err := m.ProcessMethod(ctx, collection)
	if err != nil {
		return nil, err
	}

	if result != nil {
		indexes, ok := result.(map[string]*mongoModel.DBIndex)
		if !ok {
			return nil, mocking.ErrAssertMockResult
		}

		return indexes, nil
	}

	return nil, err
}
