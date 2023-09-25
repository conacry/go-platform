package integration

import (
	"context"
	"github.com/conacry/go-platform/pkg/mongo"
	mongoContainer "github.com/conacry/go-platform/pkg/mongo/test/testDouble/container"
	commonTesting "github.com/conacry/go-platform/pkg/testing"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName = "go_common_test"

	mongoModelCollection = "mongo_models"
	mongoModelCount      = 100
)

type MongoContainerShould struct {
	suite.Suite

	MongoRepo *mongo.MongoDB

	mongoContainer         *mongoContainer.MongoContainer
	savedMongoModels       []*MongoModel
	savedMongoModelsByTags map[string][]*MongoModel
	ctx                    context.Context
}

func TestMongoContainerShould(t *testing.T) {
	suite.Run(t, &MongoContainerShould{
		mongoContainer:         mongoContainer.NewMongoContainer(),
		savedMongoModelsByTags: make(map[string][]*MongoModel),
	})
}

func (s *MongoContainerShould) SetupSuite() {
	commonTesting.SkipIntegrationTestIfNeed(s.T())
	s.mongoContainer.SetupContainer()
}

func (s *MongoContainerShould) TearDownSuite() {
	_ = s.MongoRepo.Stop(context.Background())
	s.mongoContainer.TerminateContainer()
}

func (s *MongoContainerShould) SetupTest() {
	s.MongoRepo = s.mongoContainer.NewMongoRepository(dbName)

	for index := 0; index < mongoModelCount; index++ {
		var tags []string
		if isMultipleOf(index, 3) {
			tags = append(tags, "multiple_of_three")
		}
		if isMultipleOf(index, 5) {
			tags = append(tags, "multiple_of_five")
		}

		savedMongoModel := GetMongoModelWithRandomPayload(index, tags)
		_, err := s.MongoRepo.Insert(context.Background(), mongoModelCollection, savedMongoModel)
		require.Nil(s.T(), err)

		mergedTags := strings.Join(tags, ",")
		s.savedMongoModels = append(s.savedMongoModels, savedMongoModel)
		s.savedMongoModelsByTags[mergedTags] = append(s.savedMongoModelsByTags[mergedTags], savedMongoModel)

		s.ctx = context.Background()
	}
}

func (s *MongoContainerShould) TearDownTest() {
	_, err := s.MongoRepo.DeleteMany(s.ctx, mongoModelCollection, bson.D{}, options.Delete())
	require.NoError(s.T(), err)

	s.savedMongoModels = nil
	s.savedMongoModelsByTags = make(map[string][]*MongoModel)
}

func isMultipleOf(value int, factor int) bool {
	return value%factor == 0
}
