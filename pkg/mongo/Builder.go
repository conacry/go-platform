package mongo

import (
	"github.com/conacry/go-platform/pkg/errors"
	log "github.com/conacry/go-platform/pkg/logger"
	mongoModel "github.com/conacry/go-platform/pkg/mongo/model"
)

type Builder struct {
	logger log.Logger
	config *mongoModel.Config

	errors *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		errors: errors.NewErrors(),
	}
}

func (b *Builder) Logger(logger log.Logger) *Builder {
	b.logger = logger
	return b
}

func (b *Builder) Config(config *mongoModel.Config) *Builder {
	b.config = config
	return b
}

func (b *Builder) Build() (*MongoDB, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *Builder) checkRequiredFields() {
	if b.logger == nil {
		b.errors.AddError(ErrLoggerIsRequired)
	}
	if b.config == nil {
		b.errors.AddError(ErrConfigIsRequired)
	}
}

func (b *Builder) createFromBuilder() *MongoDB {
	return &MongoDB{
		logger: b.logger,
		config: b.config,
	}
}
