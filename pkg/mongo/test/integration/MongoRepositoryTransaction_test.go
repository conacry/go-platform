package integration

import (
	"context"
	errorsStub "github.com/conacry/go-platform/pkg/errors/test/testDouble/stub"
	mongoInterface "github.com/conacry/go-platform/pkg/mongo/interface"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 1 запись создана, измнения, в другом потоке не видно
// 2 записи нет, создаем, в другом потоке не видно
// 3 ошибка внутри замыкания
//

func (s *MongoContainerShould) TestTransaction_CreateDocumentNotSeenInAnotherGoroutine_ReturnSuccess() {

}

func (s *MongoContainerShould) TestTransaction_ErrorFromCallbackFunc_ReturnError() {

}

func (s *MongoContainerShould) TestTransaction_AbortingInsert_ReturnError() {
	s.T().Skip("No way to run without mongo cluster")
	model := GetRandomMongoModel()
	expectedError := errorsStub.GetError()
	callback := func(ctx context.Context, repository mongoInterface.MongoRepository) (interface{}, error) {
		_, err := s.MongoRepo.Insert(ctx, mongoModelCollection, model)
		require.NoError(s.T(), err)

		return nil, expectedError
	}

	res, err := s.MongoRepo.Transaction(s.ctx, callback)
	require.Nil(s.T(), res)
	require.ErrorIs(s.T(), err, expectedError)

	var mongoModels []*MongoModel
	err = s.MongoRepo.Find(s.ctx, mongoModelCollection, &mongoModels, bson.M{}, options.Find())
	require.NoError(s.T(), err)

	require.Empty(s.T(), mongoModels)
}

func (s *MongoContainerShould) TestTransaction_ValidInsert_ReturnSuccess() {

}
