package httpTest

import (
	httpServer "github.com/conacry/go-platform/pkg/http/server"
	httpServerModel "github.com/conacry/go-platform/pkg/http/server/model"
	logMock "github.com/conacry/go-platform/pkg/logger/test/testDouble/mock"
	commonTesting "github.com/conacry/go-platform/pkg/testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type HttpBuilder struct {
	suite.Suite
}

func TestHttpBuilder(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(HttpBuilder))
}

func (h *HttpBuilder) TestNewBuilder_ReturnBuilder() {
	builder := httpServer.NewBuilder()
	require.NotNil(h.T(), builder)
}

func (h *HttpBuilder) TestBuild_NoParamGiven_ReturnErr() {
	expectedErr := []error{
		httpServer.ErrHttpConfigIsRequired,
		httpServer.ErrLoggerIsRequired,
	}

	builder, err := httpServer.NewBuilder().Build()
	require.Nil(h.T(), builder)
	commonTesting.AssertErrors(h.T(), err, expectedErr)
}

func (h *HttpBuilder) TestBuild_AllParam_ReturnHttpServer() {
	logger := logMock.GetLogger()
	config := httpServerModel.Config{}

	builder := httpServer.NewBuilder()
	builder.Logger(logger)
	builder.Config(&config)
	server, err := builder.Build()

	require.NoError(h.T(), err)
	require.NotNil(h.T(), server)
}
