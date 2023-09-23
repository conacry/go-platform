package mongo

import (
	"github.com/conacry/go-platform/pkg/errors"
	log "github.com/conacry/go-platform/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryBuilder struct {
	mongoDb *mongo.Database
	logger  log.Logger

	errors *errors.Errors
}

func NewMongoRepositoryBuilder() *RepositoryBuilder {
	return &RepositoryBuilder{
		errors: errors.NewErrors(),
	}
}

func (b *RepositoryBuilder) MongoDb(mongoDb *mongo.Database) *RepositoryBuilder {
	b.mongoDb = mongoDb
	return b
}

func (b *RepositoryBuilder) Logger(logger log.Logger) *RepositoryBuilder {
	b.logger = logger
	return b
}

func (b *RepositoryBuilder) Build() (*Repository, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *RepositoryBuilder) checkRequiredFields() {
	if b.mongoDb == nil {
		b.errors.AddError(ErrMongoDBIsRequired)
	}
	if b.logger == nil {
		b.errors.AddError(ErrLoggerIsRequired)
	}
}

func (b *RepositoryBuilder) createFromBuilder() *Repository {
	return &Repository{
		mongoDb: b.mongoDb,
		logger:  b.logger,
	}
}
