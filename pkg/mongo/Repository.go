package mongo

import (
	"context"

	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository interface {
	Insert(ctx context.Context, collection mongoModel.Collection, data interface{}) (string, error)

	InsertMany(ctx context.Context, collection mongoModel.Collection, data []interface{}) ([]string, error)

	FindOneAndUpdate(
		ctx context.Context,
		collection mongoModel.Collection,
		resultModel,
		filter,
		updateData interface{},
		opt *options.FindOneAndUpdateOptions,
	) error

	ReplaceOne(ctx context.Context, collection mongoModel.Collection, filter, data interface{}) error

	UpdateOne(
		ctx context.Context,
		collection mongoModel.Collection,
		filter,
		data interface{},
		opts ...options.Lister[options.UpdateOneOptions],
	) (int64, error)

	UpdateMany(
		ctx context.Context,
		collection mongoModel.Collection,
		filter interface{},
		data interface{},
		opts ...options.Lister[options.UpdateManyOptions],
	) (int64, error)

	Find(ctx context.Context, collection mongoModel.Collection, results, find interface{}, opt *options.FindOptions) error

	FindOne(
		ctx context.Context,
		collection mongoModel.Collection,
		resultModel,
		findQuery interface{},
		findOptions *options.FindOneOptions,
	) error

	DeleteOne(ctx context.Context,
		collection mongoModel.Collection,
		filter interface{},
		opts ...options.Lister[options.DeleteOneOptions],
	) (*mongo.DeleteResult, error)

	DeleteMany(ctx context.Context,
		collection mongoModel.Collection,
		filter interface{},
		opts ...options.Lister[options.DeleteManyOptions],
	) (*mongo.DeleteResult, error)

	Count(ctx context.Context, collection mongoModel.Collection, find interface{}, opt *options.CountOptions) (int64, error)

	Aggregate(ctx context.Context, collection mongoModel.Collection, pipe mongo.Pipeline) (*mongo.Cursor, error)

	CreateTextIndex(ctx context.Context, index *mongoModel.DBTextIndex) (string, error)

	CreateIndex(ctx context.Context, index *mongoModel.DBIndex) (string, error)

	TryCreateIndex(ctx context.Context, index *mongoModel.DBIndex) error

	CollectionIndexes(ctx context.Context, collection mongoModel.Collection) (map[string]*mongoModel.DBIndex, error)
}
