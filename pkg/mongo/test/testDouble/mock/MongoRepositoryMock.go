package mongoMock

import (
	"context"
	mocking "github.com/conacry/go-platform/pkg/mock"
	mongoInterface "github.com/conacry/go-platform/pkg/mongo/interface"
	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepositoryMock struct {
	*mocking.BaseMock
}

func GetMongoRepository() *MongoRepositoryMock {
	return &MongoRepositoryMock{
		BaseMock: mocking.NewBaseMock(mocking.Modes.Base()),
	}
}

func (m *MongoRepositoryMock) Insert(ctx context.Context, collectionName string, data interface{}) (string, error) {
	result, err := m.ProcessMethod(collectionName, data)
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

func (m *MongoRepositoryMock) InsertMany(ctx context.Context, collectionName string, data []interface{}) ([]string, error) {
	result, err := m.ProcessMethod(collectionName, data)
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

func (m *MongoRepositoryMock) FindOneAndUpdate(ctx context.Context, collectionName string, resultModel, filter, updateData interface{}, opt *options.FindOneAndUpdateOptions) error {
	_, err := m.ProcessMethod(collectionName, resultModel, filter, updateData, opt)
	return err
}

func (m *MongoRepositoryMock) ReplaceOne(ctx context.Context, collectionName string, filter, data interface{}) error {
	_, err := m.ProcessMethod(collectionName, filter, data)
	return err
}

func (m *MongoRepositoryMock) UpdateOne(ctx context.Context, collectionName string, filter, data interface{}, opts ...*options.UpdateOptions) (int64, error) {
	result, err := m.ProcessMethod(collectionName, filter, data, opts)
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

func (m *MongoRepositoryMock) UpdateMany(
	ctx context.Context,
	collectionName string,
	filter interface{},
	data interface{},
	opts ...*options.UpdateOptions,
) (int64, error) {
	result, err := m.ProcessMethod(collectionName, filter, data, opts)
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

func (m *MongoRepositoryMock) Find(ctx context.Context, collectionName string, results, find interface{}, opt *options.FindOptions) error {
	_, err := m.ProcessMethod(collectionName, results, find, opt)
	return err
}

func (m *MongoRepositoryMock) FindOne(ctx context.Context, collectionName string, resultModel, findQuery interface{}, findOptions *options.FindOneOptions) error {
	_, err := m.ProcessMethod(collectionName, resultModel, findQuery, findOptions)
	return err
}

func (m *MongoRepositoryMock) DeleteOne(ctx context.Context, collectionName string, filter interface{}, opt *options.DeleteOptions) (*mongo.DeleteResult, error) {
	result, err := m.ProcessMethod(collectionName, filter, opt)
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

func (m *MongoRepositoryMock) DeleteMany(ctx context.Context, collectionName string, filter interface{}, opt *options.DeleteOptions) (*mongo.DeleteResult, error) {
	result, err := m.ProcessMethod(collectionName, filter, opt)
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

func (m *MongoRepositoryMock) Count(ctx context.Context, collectionName string, find interface{}, opt *options.CountOptions) (int64, error) {
	result, err := m.ProcessMethod(collectionName, find, opt)
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

func (m *MongoRepositoryMock) Aggregate(ctx context.Context, collectionName string, pipe mongo.Pipeline) (*mongo.Cursor, error) {
	result, err := m.ProcessMethod(collectionName, pipe)
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

func (m *MongoRepositoryMock) CreateTextIndex(ctx context.Context, index *mongoModel.DBTextIndex) (string, error) {
	result, err := m.ProcessMethod(index)
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

func (m *MongoRepositoryMock) CreateIndex(ctx context.Context, index *mongoModel.DBIndex) (string, error) {
	result, err := m.ProcessMethod(index)
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

func (m *MongoRepositoryMock) TryCreateIndex(ctx context.Context, index *mongoModel.DBIndex) error {
	_, err := m.ProcessMethod(index)
	return err
}

func (m *MongoRepositoryMock) CollectionIndexes(ctx context.Context, collection string) (map[string]*mongoModel.DBIndex, error) {
	result, err := m.ProcessMethod(collection)
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

func (m *MongoRepositoryMock) Transaction(ctx context.Context, transactionFunc mongoInterface.TransactionCallbackFunc) (interface{}, error) {
	return m.ProcessMethod(ctx, transactionFunc)
}
