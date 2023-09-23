package mongo

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDatabase(url, mongoDBName string) (*mongo.Database, error) {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil, errors.Wrap(err, "error creating mongo client")
	}

	err = mongoClient.Connect(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "mongo connection error")
	}

	pingTimeout := time.Now().Add(1 * time.Second)
	ctx, cancelFunc := context.WithDeadline(context.Background(), pingTimeout)
	defer cancelFunc()

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "mongo ping error")
	}

	mongoBD := mongoClient.Database(mongoDBName)

	return mongoBD, nil
}
