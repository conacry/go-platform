package mongoDummy

import (
	"context"
	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepositoryDummy struct{}

func GetRepository() *RepositoryDummy {
	return &RepositoryDummy{}
}

func (m *RepositoryDummy) Insert(ctx context.Context, collection mongoModel.Collection, data interface{}) (string, error) {
	return "", nil
}

func (m *RepositoryDummy) InsertMany(ctx context.Context, collection mongoModel.Collection, data []interface{}) ([]string, error) {
	return nil, nil
}

func (m *RepositoryDummy) FindOneAndUpdate(ctx context.Context, collection mongoModel.Collection, resultModel, filter, updateData interface{}, opt *options.FindOneAndUpdateOptions) error {
	return nil
}

func (m *RepositoryDummy) ReplaceOne(ctx context.Context, collection mongoModel.Collection, filter, data interface{}) error {
	return nil
}

func (m *RepositoryDummy) UpdateOne(ctx context.Context, collection mongoModel.Collection, filter, data interface{}, opts ...*options.UpdateOptions) (int64, error) {
	return 0, nil
}

func (m *RepositoryDummy) UpdateMany(
	ctx context.Context,
	collection mongoModel.Collection,
	filter interface{},
	data interface{},
	opts ...*options.UpdateOptions,
) (int64, error) {
	return 0, nil
}

func (m *RepositoryDummy) Find(ctx context.Context, collection mongoModel.Collection, results, find interface{}, opt *options.FindOptions) error {
	return nil
}

func (m *RepositoryDummy) FindOne(ctx context.Context, collection mongoModel.Collection, resultModel, findQuery interface{}, findOptions *options.FindOneOptions) error {
	return nil
}

func (m *RepositoryDummy) DeleteOne(ctx context.Context, collection mongoModel.Collection, filter interface{}, opt *options.DeleteOptions) (*mongo.DeleteResult, error) {
	return nil, nil
}

func (m *RepositoryDummy) DeleteMany(ctx context.Context, collection mongoModel.Collection, filter interface{}, opt *options.DeleteOptions) (*mongo.DeleteResult, error) {
	return nil, nil
}

func (m *RepositoryDummy) Count(ctx context.Context, collection mongoModel.Collection, find interface{}, opt *options.CountOptions) (int64, error) {
	return 0, nil
}

func (m *RepositoryDummy) Aggregate(ctx context.Context, collection mongoModel.Collection, pipe mongo.Pipeline) (*mongo.Cursor, error) {
	return nil, nil
}

func (m *RepositoryDummy) CreateIndex(ctx context.Context, index *mongoModel.DBIndex) (string, error) {
	return "", nil
}

func (m *RepositoryDummy) CreateTextIndex(ctx context.Context, index *mongoModel.DBTextIndex) (string, error) {
	return "", nil
}

func (m *RepositoryDummy) CollectionIndexes(ctx context.Context, collection mongoModel.Collection) (map[string]*mongoModel.DBIndex, error) {
	return nil, nil
}

func (m *RepositoryDummy) TryCreateIndex(ctx context.Context, index *mongoModel.DBIndex) error {
	return nil
}
