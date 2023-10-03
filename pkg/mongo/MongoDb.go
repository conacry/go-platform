package mongo

import (
	"context"
	"errors"
	"fmt"
	log "github.com/conacry/go-platform/pkg/logger"
	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const (
	indexTimeoutSeconds = 10
)

type handleOperationFunc func(ctx context.Context) (interface{}, error)

type MongoDB struct {
	logger log.Logger
	config *mongoModel.Config
	db     *mongo.Database
	client *mongo.Client
}

func (m *MongoDB) Start(ctx context.Context) error {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(m.config.URL))
	if err != nil {
		msg := fmt.Sprintf("Error creating mongo client. Cause: %q", err.Error())
		err := errors.New(msg)
		m.logger.LogError(ctx, err)
		return err
	}

	err = mongoClient.Connect(context.Background())
	if err != nil {
		msg := fmt.Sprintf("Mongo connection error. Cause: %q", err.Error())
		err := errors.New(msg)
		m.logger.LogError(ctx, err)
		return err
	}

	pingTimeout := time.Now().Add(1 * time.Second)
	ctx, cancelFunc := context.WithDeadline(ctx, pingTimeout)
	defer cancelFunc()

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		msg := fmt.Sprintf("Mongo ping error. Cause: %q", err.Error())
		err := errors.New(msg)
		m.logger.LogError(ctx, err)
		return err
	}
	m.client = mongoClient

	db := mongoClient.Database(m.config.Database)
	m.db = db

	m.logger.LogInfo(ctx, "MongoDB connection is initialized")

	return nil
}

func (m *MongoDB) Stop(ctx context.Context) error {
	err := m.client.Disconnect(ctx)
	if err != nil {
		m.logger.LogError(ctx, err)
		return err
	}

	m.logger.LogInfo(ctx, "MongoDB connection is closed")

	return nil
}

func (m *MongoDB) Insert(ctx context.Context, collectionName mongoModel.Collection, data interface{}) (string, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := m.db.Collection(collectionName.String())

		res, err := coll.InsertOne(
			ctx,
			data)

		if err != nil {
			var mongoErr mongo.WriteException
			isMongoErr := errors.As(err, &mongoErr)
			if isMongoErr {
				for _, we := range mongoErr.WriteErrors {
					if m.isDuplicateError(we) {
						return "", ErrDuplicateUniqueConstraint(err)
					}
					return "", err
				}
			} else {
				return "", err
			}
		}

		var id string
		if res != nil {
			if resID, ok := res.InsertedID.(primitive.ObjectID); ok {
				id = resID.Hex()
			}
		}

		return id, nil
	}

	res, err := m.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return "", err
	}

	return res.(string), nil
}

func (m *MongoDB) InsertMany(ctx context.Context, collectionName mongoModel.Collection, data []interface{}) ([]string, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := m.db.Collection(collectionName.String())

		res, err := coll.InsertMany(
			ctx,
			data)

		if err != nil {
			var mongoErr mongo.WriteException
			isMongoErr := errors.As(err, &mongoErr)
			if isMongoErr {
				for _, we := range mongoErr.WriteErrors {
					if m.isDuplicateError(we) {
						return nil, ErrDuplicateUniqueConstraint(err)
					}
					return nil, err
				}
			} else {
				return nil, err
			}
		}

		if res == nil || len(res.InsertedIDs) == 0 {
			return nil, err
		}

		ids := make([]string, 0, len(res.InsertedIDs))
		for _, id := range res.InsertedIDs {
			if resID, ok := id.(primitive.ObjectID); ok {
				ids = append(ids, resID.Hex())
			}
		}

		return ids, nil
	}

	res, err := m.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return nil, err
	}

	return res.([]string), nil
}

func (m *MongoDB) FindOneAndUpdate(
	ctx context.Context,
	collectionName mongoModel.Collection,
	resultModel,
	filter,
	updateData interface{},
	opt *options.FindOneAndUpdateOptions,
) error {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := m.db.Collection(collectionName.String())
		result := coll.FindOneAndUpdate(
			ctx,
			filter,
			updateData,
			opt)

		err := result.Err()
		if err != nil {
			return nil, err
		}

		err = result.Decode(resultModel)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err := m.handleOperation(ctx, handleFunc)
	return err
}

func (m *MongoDB) ReplaceOne(
	ctx context.Context,
	collectionName mongoModel.Collection,
	filter interface{},
	data interface{},
) error {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := m.db.Collection(collectionName.String())
		_, err := coll.ReplaceOne(
			ctx,
			filter,
			data)
		return nil, err
	}

	_, err := m.handleOperation(ctx, handleFunc)
	return err
}

func (m *MongoDB) UpdateOne(
	ctx context.Context,
	collectionName mongoModel.Collection,
	filter,
	data interface{},
	opts ...*options.UpdateOptions,
) (int64, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := m.db.Collection(collectionName.String())
		res, err := coll.UpdateOne(
			ctx,
			filter,
			data,
			opts...)
		if err != nil {
			return 0, err
		}

		return res.ModifiedCount, nil
	}

	res, err := m.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return 0, err
	}

	return res.(int64), nil
}

func (m *MongoDB) UpdateMany(
	ctx context.Context,
	collectionName mongoModel.Collection,
	filter interface{},
	data interface{},
	opts ...*options.UpdateOptions,
) (int64, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := m.db.Collection(collectionName.String())
		res, err := coll.UpdateMany(
			ctx,
			filter,
			data,
			opts...,
		)
		if err != nil {
			return 0, err
		}

		return res.ModifiedCount, nil
	}

	res, err := m.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return 0, err
	}

	return res.(int64), nil
}

func (m *MongoDB) Find(
	ctx context.Context,
	collectionName mongoModel.Collection,
	results interface{},
	find interface{},
	opt *options.FindOptions,
) error {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := m.db.Collection(collectionName.String())
		cursor, err := collection.Find(ctx, find, opt)
		if err != nil {
			return nil, err
		}

		err = cursor.All(ctx, results)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err := m.handleOperation(ctx, handleFunc)
	return err
}

func (m *MongoDB) FindOne(
	ctx context.Context,
	collectionName mongoModel.Collection,
	resultModel,
	findQuery interface{},
	findOptions *options.FindOneOptions,
) error {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := m.db.Collection(collectionName.String())
		result := collection.FindOne(ctx, findQuery, findOptions)
		err := result.Err()
		if err != nil {
			return nil, err
		}

		err = result.Decode(resultModel)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err := m.handleOperation(ctx, handleFunc)
	return err
}

func (m *MongoDB) DeleteOne(
	ctx context.Context,
	collectionName mongoModel.Collection,
	filter interface{},
	opt *options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := m.db.Collection(collectionName.String())
		result, err := collection.DeleteOne(ctx, filter, opt)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	res, err := m.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return nil, err
	}

	return res.(*mongo.DeleteResult), nil
}

func (m *MongoDB) DeleteMany(
	ctx context.Context,
	collectionName mongoModel.Collection,
	filter interface{},
	opt *options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := m.db.Collection(collectionName.String())
		result, err := collection.DeleteMany(ctx, filter, opt)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	res, err := m.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return nil, err
	}

	return res.(*mongo.DeleteResult), nil
}

func (m *MongoDB) Count(
	ctx context.Context,
	collectionName mongoModel.Collection,
	find interface{},
	opt *options.CountOptions,
) (int64, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := m.db.Collection(collectionName.String())
		count, err := collection.CountDocuments(ctx, find, opt)
		if err != nil {
			return 0, err
		}

		return count, nil
	}

	res, err := m.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return 0, err
	}

	return res.(int64), nil
}

func (m *MongoDB) Aggregate(
	ctx context.Context,
	collectionName mongoModel.Collection,
	pipe mongo.Pipeline,
) (*mongo.Cursor, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := m.db.Collection(collectionName.String())
		cursor, err := collection.Aggregate(ctx, pipe)
		if err != nil {
			return nil, err
		}

		return cursor, nil
	}

	res, err := m.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return nil, err
	}

	return res.(*mongo.Cursor), nil
}

func (m *MongoDB) CreateIndex(ctx context.Context, index *mongoModel.DBIndex) (string, error) {
	c := m.db.Collection(index.Collection.String())
	opts := options.CreateIndexes().SetMaxTime(indexTimeoutSeconds * time.Second)

	keysName := make([]bsonx.Elem, 0)
	for _, k := range index.Keys {
		keysName = append(keysName, bsonx.Elem{
			Key:   k,
			Value: bsonx.Int32(int32(index.Type)),
		})
	}
	keys := bsonx.Doc(keysName)
	indexModel := mongo.IndexModel{}
	indexModel.Keys = keys
	indexModel.Options = options.Index().SetName(index.Name)
	if index.Uniq {
		indexModel.Options.SetUnique(true)
	}

	return c.Indexes().CreateOne(ctx, indexModel, opts)
}

func (m *MongoDB) CreateTextIndex(ctx context.Context, index *mongoModel.DBTextIndex) (string, error) {
	c := m.db.Collection(index.Collection)
	opts := options.CreateIndexes().SetMaxTime(indexTimeoutSeconds * time.Second)

	keysName := make([]bsonx.Elem, 0)

	for _, k := range index.Keys {
		keysName = append(keysName, bsonx.Elem{
			Key:   k,
			Value: bsonx.String("text"),
		})
	}

	keys := bsonx.Doc(keysName)
	indexModel := mongo.IndexModel{}
	indexModel.Keys = keys
	indexModel.Options = options.Index().SetName(index.Name)

	return c.Indexes().CreateOne(ctx, indexModel, opts)
}

func (m *MongoDB) CollectionIndexes(ctx context.Context, collection mongoModel.Collection) (map[string]*mongoModel.DBIndex, error) {
	res := make(map[string]*mongoModel.DBIndex)
	c := m.db.Collection(collection.String())
	duration := indexTimeoutSeconds * time.Second
	opts := &options.ListIndexesOptions{MaxTime: &duration}
	cur, err := c.Indexes().List(ctx, opts)
	if err != nil {
		return res, err
	}
	for cur.Next(ctx) {
		index := &mongoModel.DBIndex{}
		if err := cur.Decode(&index); err != nil {
			return res, err
		}
		res[index.Name] = index
	}
	return res, nil
}

func (m *MongoDB) TryCreateIndex(ctx context.Context, index *mongoModel.DBIndex) error {
	indexes, err := m.CollectionIndexes(ctx, index.Collection)
	if err != nil {
		return err
	}

	if m.isIndexExist(index, indexes) {
		return nil
	}

	_, err = m.CreateIndex(ctx, index)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDB) handleOperation(
	ctx context.Context,
	handleFunc handleOperationFunc,
) (interface{}, error) {
	res, err := handleFunc(ctx)

	return res, err
}

func (m *MongoDB) isDuplicateError(we mongo.WriteError) bool {
	return we.Code == 11000
}

func (m *MongoDB) isIndexExist(index *mongoModel.DBIndex, indexes map[string]*mongoModel.DBIndex) bool {
	_, ok := indexes[index.Name]
	return ok
}
