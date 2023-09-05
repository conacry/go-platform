package clickhouse

import "github.com/conacry/go-platform/pkg/errors"

type Builder struct {
	config *Config
	logger Logger

	errors *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		errors: errors.NewErrors(),
	}
}

func (b *Builder) Config(config *Config) *Builder {
	b.config = config
	return b
}

func (b *Builder) Logger(logger Logger) *Builder {
	b.logger = logger
	return b
}

func (b *Builder) Build() (*ClickHouse, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.createFromBuilder(), nil
}

func (b *Builder) checkRequiredFields() {
	if b.config == nil {
		b.errors.AddError(ErrClickHouseConfigIsRequired)
	}
	if b.logger == nil {
		b.errors.AddError(ErrLoggerIsRequired)
	}
}

func (b *Builder) createFromBuilder() *ClickHouse {
	return &ClickHouse{
		config: b.config,
		logger: b.logger,
	}
}
