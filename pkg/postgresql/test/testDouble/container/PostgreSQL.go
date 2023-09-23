package postgresqlContainer

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"time"
)

const (
	defaultPostgreSQLImage = "postgres:15.4-alpine3.18"
	innerPostgreSQLPort    = "5432"
)

type ContainerInfo struct {
	Host string
	Port string
}

type PostgreSQL struct {
	ctx       context.Context
	container testcontainers.Container
	info      ContainerInfo
}

func NewContainer() *PostgreSQL {
	return &PostgreSQL{}
}

func (c *PostgreSQL) Setup() {
	c.SetupWithImage(defaultPostgreSQLImage)
}

func (c *PostgreSQL) SetupWithImage(imageName string) {
	c.ctx = context.Background()
	c.container = startContainer(c.ctx, imageName)

	host, err := c.container.Host(c.ctx)
	if err != nil {
		panic(err)
	}

	mappedPort, err := c.container.MappedPort(c.ctx, innerPostgreSQLPort)
	if err != nil {
		panic(err)
	}

	c.info = ContainerInfo{
		Host: host,
		Port: mappedPort.Port(),
	}
}

func startContainer(ctx context.Context, imageName string) testcontainers.Container {
	dbName := "test_db"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage(imageName),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		panic(err)
	}

	return postgresContainer
}
