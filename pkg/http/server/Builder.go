package httpServer

import (
	"github.com/conacry/go-platform/pkg/errors"
	httpServerModel "github.com/conacry/go-platform/pkg/http/server/model"
	httpResponse "github.com/conacry/go-platform/pkg/http/server/response"
	log "github.com/conacry/go-platform/pkg/logger"
	"github.com/go-chi/chi/v5"
)

type Builder struct {
	logger              log.Logger
	config              *httpServerModel.Config
	responseWriter      *httpResponse.Writer
	errorResponseWriter *httpResponse.ErrorWriter
	errors              *errors.Errors
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

func (b *Builder) Config(config *httpServerModel.Config) *Builder {
	b.config = config
	return b
}

func (b *Builder) ResponseWriter(responseWriter *httpResponse.Writer) *Builder {
	b.responseWriter = responseWriter
	return b
}

func (b *Builder) ErrorResponseWriter(errResponseWriter *httpResponse.ErrorWriter) *Builder {
	b.errorResponseWriter = errResponseWriter
	return b
}

func (b *Builder) Build() (*HttpServer, error) {
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
		b.errors.AddError(ErrHttpConfigIsRequired)
	}
}

func (b *Builder) createFromBuilder() *HttpServer {
	return &HttpServer{
		router:              chi.NewRouter(),
		logger:              b.logger,
		config:              b.config,
		responseWriter:      b.responseWriter,
		errorResponseWriter: b.errorResponseWriter,
	}
}
