package mongoDummy

import (
	"context"

	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type RepositoryDummy struct{}

func GetRepository() *RepositoryDummy {
	return &RepositoryDummy{}
}

func (m *RepositoryDummy) Insert(
	ctx context.Context,
	collection mongoModel.Collection,
	data any,
) (string, error) {
	return "", nil
}

func (m *RepositoryDummy) InsertMany(
	ctx context.Context,
	collection mongoModel.Collection,
	data []any,
) ([]string, error) {
	return nil, nil
}

func (m *RepositoryDummy) FindOneAndUpdate(
	ctx context.Context,
	collection mongoModel.Collection,
	resultModel, filter,
	updateData any,
	opt *options.FindOneAndUpdateOptions,
) error {
	return nil
}

func (m *RepositoryDummy) ReplaceOne(
	ctx context.Context,
	collection mongoModel.Collection,
	filter, data any,
) error {
	return nil
}

func (m *RepositoryDummy) UpdateOne(
	ctx context.Context,
	collection mongoModel.Collection,
	filter, data any,
	opts ...options.Lister[options.UpdateOneOptions],
) (int64, error) {
	return 0, nil
}

func (m *RepositoryDummy) UpdateMany(
	ctx context.Context,
	collection mongoModel.Collection,
	filter any,
	data any,
	opts ...options.Lister[options.UpdateManyOptions],
) (int64, error) {
	return 0, nil
}

func (m *RepositoryDummy) Find(
	ctx context.Context,
	collection mongoModel.Collection,
	results, find any,
	opt *options.FindOptions,
) error {
	return nil
}

func (m *RepositoryDummy) FindOne(
	ctx context.Context,
	collection mongoModel.Collection,
	resultModel, findQuery any,
	findOptions *options.FindOneOptions,
) error {
	return nil
}

func (m *RepositoryDummy) DeleteOne(
	ctx context.Context,
	collection mongoModel.Collection,
	filter any,
	opts ...options.Lister[options.DeleteOneOptions],
) (*mongo.DeleteResult, error) {
	return nil, nil
}

func (m *RepositoryDummy) DeleteMany(
	ctx context.Context,
	collection mongoModel.Collection,
	filter any,
	opts ...options.Lister[options.DeleteManyOptions],
) (*mongo.DeleteResult, error) {
	return nil, nil
}

func (m *RepositoryDummy) Count(
	ctx context.Context,
	collection mongoModel.Collection,
	find any,
	opt *options.CountOptions,
) (int64, error) {
	return 0, nil
}

func (m *RepositoryDummy) Aggregate(
	ctx context.Context,
	collection mongoModel.Collection,
	pipe mongo.Pipeline,
) (*mongo.Cursor, error) {
	return nil, nil
}

func (m *RepositoryDummy) CreateIndex(
	ctx context.Context,
	index *mongoModel.DBIndex,
) (string, error) {
	return "", nil
}

func (m *RepositoryDummy) CreateTextIndex(
	ctx context.Context,
	index *mongoModel.DBTextIndex,
) (string, error) {
	return "", nil
}

func (m *RepositoryDummy) CollectionIndexes(
	ctx context.Context,
	collection mongoModel.Collection,
) (map[string]*mongoModel.DBIndex, error) {
	return nil, nil
}

func (m *RepositoryDummy) TryCreateIndex(
	ctx context.Context,
	index *mongoModel.DBIndex,
) error {
	return nil
}
