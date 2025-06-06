package mongoMock

import (
	"context"

	mocking "github.com/conacry/go-platform/pkg/mock"
	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (m *RepositoryMock) Insert(
	ctx context.Context,
	collection mongoModel.Collection,
	data any,
) (string, error) {
	result, err := m.ProcessMethod(ctx, collection, data)
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

func (m *RepositoryMock) InsertMany(
	ctx context.Context,
	collection mongoModel.Collection,
	data []any,
) ([]string, error) {
	result, err := m.ProcessMethod(ctx, collection, data)
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

func (m *RepositoryMock) FindOneAndUpdate(
	ctx context.Context,
	collection mongoModel.Collection,
	resultModel, filter,
	updateData any,
	opt *options.FindOneAndUpdateOptions,
) error {
	_, err := m.ProcessMethod(ctx, collection, resultModel, filter, updateData, opt)
	return err
}

func (m *RepositoryMock) ReplaceOne(
	ctx context.Context,
	collection mongoModel.Collection,
	filter, data any,
) error {
	_, err := m.ProcessMethod(ctx, collection, filter, data)
	return err
}

func (m *RepositoryMock) UpdateOne(
	ctx context.Context,
	collection mongoModel.Collection,
	filter, data any,
	opts ...options.Lister[options.UpdateOneOptions],
) (int64, error) {
	result, err := m.ProcessMethod(ctx, collection, filter, data, opts)
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
	collection mongoModel.Collection,
	filter, data any,
	opts ...options.Lister[options.UpdateManyOptions],
) (int64, error) {
	result, err := m.ProcessMethod(ctx, collection, filter, data, opts)
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

func (m *RepositoryMock) Find(
	ctx context.Context,
	collection mongoModel.Collection,
	results, find any,
	opt *options.FindOptions,
) error {
	_, err := m.ProcessMethod(ctx, collection, results, find, opt)
	return err
}

func (m *RepositoryMock) FindOne(
	ctx context.Context,
	collection mongoModel.Collection,
	resultModel, findQuery any,
	opts ...options.Lister[options.FindOneOptions],
) error {
	args := []any{ctx, collection, resultModel, findQuery}
	args = append(args, opts)
	_, err := m.ProcessMethod(args...)
	return err
}

func (m *RepositoryMock) DeleteOne(
	ctx context.Context,
	collection mongoModel.Collection,
	filter any,
	opts ...options.Lister[options.DeleteOneOptions],
) (*mongo.DeleteResult, error) {
	args := []any{ctx, collection, filter}
	args = append(args, opts)
	result, err := m.ProcessMethod(args...)
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

func (m *RepositoryMock) DeleteMany(
	ctx context.Context,
	collection mongoModel.Collection,
	filter any,
	opts ...options.Lister[options.DeleteManyOptions],
) (*mongo.DeleteResult, error) {
	args := []any{ctx, collection, filter}
	args = append(args, opts)
	result, err := m.ProcessMethod(args...)
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

func (m *RepositoryMock) Count(
	ctx context.Context,
	collection mongoModel.Collection,
	find any,
	opt *options.CountOptions,
) (int64, error) {
	result, err := m.ProcessMethod(ctx, collection, find, opt)
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

func (m *RepositoryMock) Aggregate(
	ctx context.Context,
	collection mongoModel.Collection,
	pipe mongo.Pipeline,
) (*mongo.Cursor, error) {
	result, err := m.ProcessMethod(ctx, collection, pipe)
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

func (m *RepositoryMock) CollectionIndexes(
	ctx context.Context,
	collection mongoModel.Collection,
) (map[string]*mongoModel.DBIndex, error) {
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
