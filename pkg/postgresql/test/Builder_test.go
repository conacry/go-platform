package posgresqlTest

import (
	logMock "github.com/conacry/go-platform/pkg/logger/test/testDouble/mock"
	"github.com/conacry/go-platform/pkg/postgresql"
	commonTesting "github.com/conacry/go-platform/pkg/testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PostgreSQLBuilder struct {
	suite.Suite
}

func TestPostgreSQLBuilder(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(PostgreSQLBuilder))
}

func (p *PostgreSQLBuilder) TestNewBuilder_ReturnBuilder() {
	builder := postgresql.NewBuilder()
	require.NotNil(p.T(), builder)
}

func (p *PostgreSQLBuilder) TestBuild_NoParamGiven_ReturnErr() {
	expectedErr := []error{
		postgresql.ErrConfigIsRequired,
		postgresql.ErrLoggerIsRequired,
	}

	builder, err := postgresql.NewBuilder().Build()
	require.Nil(p.T(), builder)
	commonTesting.AssertErrors(p.T(), err, expectedErr)
}

func (p *PostgreSQLBuilder) TestBuild_AllParam_ReturnPostgreSQL() {
	logger := logMock.GetLogger()
	config := postgresql.Config{}

	builder := postgresql.NewBuilder()
	builder.Logger(logger)
	builder.Config(&config)
	postgreSQL, err := builder.Build()

	require.NoError(p.T(), err)
	require.NotNil(p.T(), postgreSQL)
}
