package postgresqlContainerTest

import (
	postgresqlContainer "github.com/conacry/go-platform/pkg/postgresql/test/testDouble/container"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PostgreSQLShould struct {
	suite.Suite
}

func TestPostgreSQLShould(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(PostgreSQLShould))
}

func (s *PostgreSQLShould) Test() {
	pgContainer := postgresqlContainer.NewContainer()

	assert.NotPanics(s.T(), func() {
		pgContainer.Setup()
	})
}
