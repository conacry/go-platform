package commonTestcontainer

import (
	"context"
	"fmt"
	"github.com/conacry/go-platform/pkg/mongo"
	mongoTestUtil "github.com/conacry/go-platform/pkg/mongo/test/util"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"time"
)

const (
	defaultMongoImage = "mongo:5.0.6"
	innerMongoPort    = "27017"
)

type MongoContainer struct {
	container   testcontainers.Container
	mongoConfig *mongoTestUtil.MongoTestConfig
	ctx         context.Context
}

func NewMongoContainer() *MongoContainer {
	return &MongoContainer{}
}

func (c *MongoContainer) SetupContainer() {
	c.SetupContainerByParams(defaultMongoImage)
}

func (c *MongoContainer) SetupContainerByParams(mongoImage string) {
	c.ctx = context.Background()

	c.container = startMongoContainer(c.ctx, mongoImage)

	mongoHost, err := c.container.Host(c.ctx)
	if err != nil {
		panic(err)
	}
	mappedPort, err := c.container.MappedPort(c.ctx, innerMongoPort)
	if err != nil {
		panic(err)
	}

	c.mongoConfig = &mongoTestUtil.MongoTestConfig{
		Host: mongoHost,
		Port: mappedPort.Port(),
	}
}

func (c *MongoContainer) TerminateContainer() {
	err := c.container.Terminate(c.ctx)
	if err != nil {
		panic(err)
	}
}

func (c *MongoContainer) MongoConfig() *mongoTestUtil.MongoTestConfig {
	return c.mongoConfig
}

func (c *MongoContainer) NewMongoRepository(dbName string) *mongo.Repository {
	return mongoTestUtil.NewMongoRepository(c.mongoConfig, dbName)
}

func startMongoContainer(ctx context.Context, mongoImage string) testcontainers.Container {
	exposedPort := fmt.Sprintf("%s/tcp", innerMongoPort)

	containerRequest := testcontainers.ContainerRequest{
		Image:        mongoImage,
		ExposedPorts: []string{exposedPort},
		WaitingFor:   wait.ForLog("Waiting for connections").WithStartupTimeout(30 * time.Second),
	}

	mongoContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: containerRequest,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}

	return mongoContainer
}
