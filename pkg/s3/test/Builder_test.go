package s3Test

import (
	logMock "github.com/conacry/go-platform/pkg/logger/test/testDouble/mock"
	"github.com/conacry/go-platform/pkg/postgresql"
	s3Storage "github.com/conacry/go-platform/pkg/s3"
	commonTesting "github.com/conacry/go-platform/pkg/testing/error"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type S3Builder struct {
	suite.Suite
}

func TestS3Builder(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(S3Builder))
}

func (s *S3Builder) TestNewBuilder_ReturnBuilder() {
	builder := s3Storage.NewBuilder()
	require.NotNil(s.T(), builder)
}

func (s *S3Builder) TestBuild_NoParamGiven_ReturnErr() {
	expectedErr := []error{
		s3Storage.ErrS3ConfigIsRequired,
		s3Storage.ErrLoggerIsRequired,
	}

	builder, err := s3Storage.NewBuilder().Build()
	require.Nil(s.T(), builder)
	commonTesting.AssertErrors(s.T(), err, expectedErr)
}

func (s *S3Builder) TestBuild_LoggerIsNil_ReturnErr() {
	expectedErr := []error{
		s3Storage.ErrLoggerIsRequired,
	}

	config := s3Storage.Config{}

	builder := s3Storage.NewBuilder()
	builder.Config(&config)
	s3, err := builder.Build()

	require.Nil(s.T(), s3)
	commonTesting.AssertErrors(s.T(), err, expectedErr)
}

func (s *S3Builder) TestBuild_ConfigIsNil_ReturnErr() {
	expectedErr := []error{
		postgresql.ErrConfigIsRequired,
	}

	logger := logMock.GetLogger()
	builder := postgresql.NewBuilder()
	builder.Logger(logger)
	s3, err := builder.Build()

	require.Nil(s.T(), s3)
	commonTesting.AssertErrors(s.T(), err, expectedErr)
}

func (s *S3Builder) TestBuild_AllParam_ReturnClickhouse() {
	logger := logMock.GetLogger()
	config := postgresql.Config{}

	builder := postgresql.NewBuilder()
	builder.Logger(logger)
	builder.Config(&config)
	s3, err := builder.Build()

	require.NoError(s.T(), err)
	require.NotNil(s.T(), s3)
}
