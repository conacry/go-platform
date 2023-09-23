package mongoTestUtil

import (
	"fmt"
	logMock "github.com/conacry/go-platform/pkg/logger/test/testDouble/mock"
	"github.com/conacry/go-platform/pkg/mongo"
)

type MongoTestConfig struct {
	Host string
	Port string
}

func (c *MongoTestConfig) MongoURI() string {
	return fmt.Sprintf("mongodb://%s:%s", c.Host, c.Port)
}

func NewMongoRepository(config *MongoTestConfig, dbName string) *mongo.Repository {
	mongoDB, err := mongo.InitMongoDatabase(config.MongoURI(), dbName)
	if err != nil {
		panic(err)
	}

	mongoRepository, err := mongo.NewMongoRepositoryBuilder().
		MongoDb(mongoDB).
		Logger(logMock.GetLogger()).
		Build()
	if err != nil {
		panic(err)
	}

	return mongoRepository
}
