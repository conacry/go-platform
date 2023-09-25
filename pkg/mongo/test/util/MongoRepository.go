package mongoTestUtil

import (
	"context"
	"fmt"
	logMock "github.com/conacry/go-platform/pkg/logger/test/testDouble/mock"
	"github.com/conacry/go-platform/pkg/mongo"
	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"
)

type MongoTestConfig struct {
	Host string
	Port string
}

func (c *MongoTestConfig) MongoURI() string {
	return fmt.Sprintf("mongodb://%s:%s", c.Host, c.Port)
}

func NewMongoRepository(testConfig *MongoTestConfig, dbName string) *mongo.MongoDB {
	mongoConfig := &mongoModel.Config{
		URL:      testConfig.MongoURI(),
		Database: dbName,
	}

	mongoDB, err := mongo.NewRepositoryBuilder().
		Logger(logMock.GetLogger()).
		Config(mongoConfig).
		Build()
	if err != nil {
		panic(err)
	}

	err = mongoDB.Start(context.Background())
	if err != nil {
		panic(err)
	}

	return mongoDB
}
