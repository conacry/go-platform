package mongo

import (
	"context"
	log "github.com/conacry/go-platform/pkg/logger"
	mongoInterface "github.com/conacry/go-platform/pkg/mongo/interface"
	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const (
	indexTimeoutSeconds = 10
)

type handleOperationFunc func(ctx context.Context) (interface{}, error)

type Repository struct {
	mongoDb *mongo.Database
	logger  log.Logger
}

func (r *Repository) Insert(ctx context.Context, collectionName string, data interface{}) (string, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := r.mongoDb.Collection(collectionName)

		res, err := coll.InsertOne(
			ctx,
			data)

		if err != nil {
			mongoErr, isMongoErr := err.(mongo.WriteException)
			if isMongoErr {
				for _, we := range mongoErr.WriteErrors {
					if r.isDuplicateError(we) {
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

	res, err := r.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return "", err
	}

	return res.(string), nil
}

func (r *Repository) InsertMany(ctx context.Context, collectionName string, data []interface{}) ([]string, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := r.mongoDb.Collection(collectionName)

		res, err := coll.InsertMany(
			ctx,
			data)

		if err != nil {
			mongoErr, isMongoErr := err.(mongo.WriteException)
			if isMongoErr {
				for _, we := range mongoErr.WriteErrors {
					if r.isDuplicateError(we) {
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

	res, err := r.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return nil, err
	}

	return res.([]string), nil
}

func (r *Repository) FindOneAndUpdate(
	ctx context.Context,
	collectionName string,
	resultModel,
	filter,
	updateData interface{},
	opt *options.FindOneAndUpdateOptions,
) error {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := r.mongoDb.Collection(collectionName)
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

	_, err := r.handleOperation(ctx, handleFunc)
	return err
}

func (r *Repository) ReplaceOne(
	ctx context.Context,
	collectionName string,
	filter interface{},
	data interface{},
) error {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := r.mongoDb.Collection(collectionName)
		_, err := coll.ReplaceOne(
			ctx,
			filter,
			data)
		return nil, err
	}

	_, err := r.handleOperation(ctx, handleFunc)
	return err
}

func (r *Repository) UpdateOne(
	ctx context.Context,
	collectionName string,
	filter,
	data interface{},
	opts ...*options.UpdateOptions,
) (int64, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := r.mongoDb.Collection(collectionName)
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

	res, err := r.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return 0, err
	}

	return res.(int64), nil
}

func (r *Repository) UpdateMany(
	ctx context.Context,
	collectionName string,
	filter interface{},
	data interface{},
	opts ...*options.UpdateOptions,
) (int64, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		coll := r.mongoDb.Collection(collectionName)
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

	res, err := r.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return 0, err
	}

	return res.(int64), nil
}

func (r *Repository) Find(
	ctx context.Context,
	collectionName string,
	results interface{},
	find interface{},
	opt *options.FindOptions,
) error {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := r.mongoDb.Collection(collectionName)
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

	_, err := r.handleOperation(ctx, handleFunc)
	return err
}

func (r *Repository) FindOne(
	ctx context.Context,
	collectionName string,
	resultModel,
	findQuery interface{},
	findOptions *options.FindOneOptions,
) error {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := r.mongoDb.Collection(collectionName)
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

	_, err := r.handleOperation(ctx, handleFunc)
	return err
}

func (r *Repository) DeleteOne(
	ctx context.Context,
	collectionName string,
	filter interface{},
	opt *options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := r.mongoDb.Collection(collectionName)
		result, err := collection.DeleteOne(ctx, filter, opt)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	res, err := r.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return nil, err
	}

	return res.(*mongo.DeleteResult), nil
}

func (r *Repository) DeleteMany(
	ctx context.Context,
	collectionName string,
	filter interface{},
	opt *options.DeleteOptions,
) (*mongo.DeleteResult, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := r.mongoDb.Collection(collectionName)
		result, err := collection.DeleteMany(ctx, filter, opt)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	res, err := r.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return nil, err
	}

	return res.(*mongo.DeleteResult), nil
}

func (r *Repository) Count(
	ctx context.Context,
	collectionName string,
	find interface{},
	opt *options.CountOptions,
) (int64, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := r.mongoDb.Collection(collectionName)
		count, err := collection.CountDocuments(ctx, find, opt)
		if err != nil {
			return 0, err
		}

		return count, nil
	}

	res, err := r.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return 0, err
	}

	return res.(int64), nil
}

func (r *Repository) Aggregate(
	ctx context.Context,
	collectionName string,
	pipe mongo.Pipeline,
) (*mongo.Cursor, error) {
	handleFunc := func(ctx context.Context) (interface{}, error) {
		collection := r.mongoDb.Collection(collectionName)
		cursor, err := collection.Aggregate(ctx, pipe)
		if err != nil {
			return nil, err
		}

		return cursor, nil
	}

	res, err := r.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return nil, err
	}

	return res.(*mongo.Cursor), nil
}

func (r *Repository) CreateIndex(ctx context.Context, index *mongoModel.DBIndex) (string, error) {
	c := r.mongoDb.Collection(index.Collection)
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

func (r *Repository) CreateTextIndex(ctx context.Context, index *mongoModel.DBTextIndex) (string, error) {
	c := r.mongoDb.Collection(index.Collection)
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

func (r *Repository) CollectionIndexes(ctx context.Context, collection string) (map[string]*mongoModel.DBIndex, error) {
	res := make(map[string]*mongoModel.DBIndex)
	c := r.mongoDb.Collection(collection)
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

func (r *Repository) Transaction(
	ctx context.Context,
	transactionFunc mongoInterface.TransactionCallbackFunc,
) (interface{}, error) {
	handleFunc := r.getTransactionWrapper(transactionFunc)

	res, err := r.handleOperation(ctx, handleFunc)
	if res == nil || err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Repository) TryCreateIndex(ctx context.Context, index *mongoModel.DBIndex) error {
	indexes, err := r.CollectionIndexes(ctx, index.Collection)
	if err != nil {
		return err
	}

	if r.isIndexExist(index, indexes) {
		return nil
	}

	_, err = r.CreateIndex(ctx, index)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) handleOperation(
	ctx context.Context,
	handleFunc handleOperationFunc,
) (interface{}, error) {
	res, err := handleFunc(ctx)

	return res, err
}

func (r *Repository) isDuplicateError(we mongo.WriteError) bool {
	return we.Code == 11000
}

func (r *Repository) isIndexExist(index *mongoModel.DBIndex, indexes map[string]*mongoModel.DBIndex) bool {
	_, ok := indexes[index.Name]
	return ok
}

func (r *Repository) getTransactionWrapper(
	transactionFunc mongoInterface.TransactionCallbackFunc,
) handleOperationFunc {
	return r.getRegularModeTransactionWrapper(transactionFunc)
}

func (r *Repository) getRegularModeTransactionWrapper(
	transactionFunc mongoInterface.TransactionCallbackFunc,
) handleOperationFunc {

	handleFunc := func(ctx context.Context) (interface{}, error) {
		wc := writeconcern.New(writeconcern.WMajority())
		rc := readconcern.Snapshot()
		txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

		session, err := r.mongoDb.Client().StartSession()
		if err != nil {
			return nil, ErrStartSession(err)
		}
		defer session.EndSession(ctx)

		if err = session.StartTransaction(txnOpts); err != nil {
			return nil, ErrStartTransaction(err)
		}

		var transactionResult interface{}
		err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
			var errTx error
			transactionResult, errTx = transactionFunc(sc, r)
			if errTx != nil {
				if errAborting := session.AbortTransaction(sc); errAborting != nil {
					return ErrAbortTransaction(errAborting, errTx)
				}

				return errTx
			}

			if errCommiting := session.CommitTransaction(sc); errCommiting != nil {
				domainCommitError := ErrCommitTransaction(errCommiting)
				if errAborting := session.AbortTransaction(sc); errAborting != nil {
					return ErrAbortTransaction(errAborting, domainCommitError)
				}

				return domainCommitError
			}

			return nil
		})

		if err != nil {
			return nil, err
		}

		return transactionResult, nil
	}

	return handleFunc
}

func (r *Repository) getStandAloneModeTransactionWrapper(
	transactionFunc mongoInterface.TransactionCallbackFunc,
) handleOperationFunc {

	handleFunc := func(ctx context.Context) (interface{}, error) {
		session, err := r.mongoDb.Client().StartSession()
		if err != nil {
			return nil, ErrStartSession(err)
		}
		defer session.EndSession(ctx)

		var transactionResult interface{}
		err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
			var errTx error
			transactionResult, errTx = transactionFunc(sc, r)
			if errTx != nil {
				return errTx
			}

			return nil
		})

		if err != nil {
			return nil, err
		}

		return transactionResult, nil
	}

	return handleFunc
}
