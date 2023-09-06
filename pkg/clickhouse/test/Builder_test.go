package clickhouseTest

import (
	"github.com/conacry/go-platform/pkg/clickhouse"
	logMock "github.com/conacry/go-platform/pkg/logger/test/testDouble/mock"
	commonTesting "github.com/conacry/go-platform/pkg/testing/error"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClickhouseBuilder struct {
	suite.Suite
}

func TestClickhouseBuilder(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ClickhouseBuilder))
}

func (c *ClickhouseBuilder) TestNewBuilder_ReturnBuilder() {
	builder := clickhouse.NewBuilder()
	require.NotNil(c.T(), builder)
}

func (c *ClickhouseBuilder) TestBuild_NoParamGiven_ReturnErr() {
	expectedErr := []error{
		clickhouse.ErrClickHouseConfigIsRequired,
		clickhouse.ErrLoggerIsRequired,
	}

	builder, err := clickhouse.NewBuilder().Build()
	require.Nil(c.T(), builder)
	commonTesting.AssertErrors(c.T(), err, expectedErr)
}

func (c *ClickhouseBuilder) TestBuild_LoggerIsNil_ReturnErr() {
	expectedErr := []error{
		clickhouse.ErrLoggerIsRequired,
	}

	config := clickhouse.Config{}

	builder := clickhouse.NewBuilder()
	builder.Config(&config)
	ch, err := builder.Build()

	require.Nil(c.T(), ch)
	commonTesting.AssertErrors(c.T(), err, expectedErr)
}

func (c *ClickhouseBuilder) TestBuild_ConfigIsNil_ReturnErr() {
	expectedErr := []error{
		clickhouse.ErrClickHouseConfigIsRequired,
	}

	logger := logMock.GetLogger()
	builder := clickhouse.NewBuilder()
	builder.Logger(logger)
	ch, err := builder.Build()

	require.Nil(c.T(), ch)
	commonTesting.AssertErrors(c.T(), err, expectedErr)
}

func (c *ClickhouseBuilder) TestBuild_AllParam_ReturnClickhouse() {
	logger := logMock.GetLogger()
	config := clickhouse.Config{}

	builder := clickhouse.NewBuilder()
	builder.Logger(logger)
	builder.Config(&config)
	ch, err := builder.Build()

	require.NoError(c.T(), err)
	require.NotNil(c.T(), ch)
}
