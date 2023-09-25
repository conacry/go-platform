package mongo

import (
	"github.com/conacry/go-platform/pkg/errors"
	log "github.com/conacry/go-platform/pkg/logger"
	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"
)

type RepositoryBuilder struct {
	logger log.Logger
	config *mongoModel.Config

	errors *errors.Errors
}

func NewRepositoryBuilder() *RepositoryBuilder {
	return &RepositoryBuilder{
		errors: errors.NewErrors(),
	}
}

func (b *RepositoryBuilder) Logger(logger log.Logger) *RepositoryBuilder {
	b.logger = logger
	return b
}

func (b *RepositoryBuilder) Config(config *mongoModel.Config) *RepositoryBuilder {
	b.config = config
	return b
}

func (b *RepositoryBuilder) Build() (*MongoDB, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *RepositoryBuilder) checkRequiredFields() {
	if b.logger == nil {
		b.errors.AddError(ErrLoggerIsRequired)
	}
	if b.config == nil {
		b.errors.AddError(ErrConfigIsRequired)
	}
}

func (b *RepositoryBuilder) createFromBuilder() *MongoDB {
	return &MongoDB{
		logger: b.logger,
		config: b.config,
	}
}
