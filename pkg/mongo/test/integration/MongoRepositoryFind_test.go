package integration

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (s *MongoContainerShould) TestFind_WithNoFilterAndAscendingOrder_ReturnSuccess() {
	opt := options.Find().SetSort(bson.M{"order": 1})

	var mongoModels []*MongoModel
	err := s.MongoRepo.Find(s.ctx, mongoModelCollection, &mongoModels, bson.M{}, opt)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.savedMongoModels, mongoModels)
}

func (s *MongoContainerShould) TestFind_WithNoFilterAndDescendingOrder_ReturnSuccess() {
	opt := options.Find().SetSort(bson.M{"order": -1})

	var mongoModels []*MongoModel
	err := s.MongoRepo.Find(s.ctx, mongoModelCollection, &mongoModels, bson.M{}, opt)

	expectedMongoModels := make([]*MongoModel, mongoModelCount)
	copy(expectedMongoModels, s.savedMongoModels)
	reverseOrder(expectedMongoModels)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedMongoModels, mongoModels)
}

func (s *MongoContainerShould) TestFind_WithLimit_ReturnSuccess() {
	limit := int64(10)
	opt := options.Find().SetSort(bson.M{"order": 1}).SetLimit(limit)

	var mongoModels []*MongoModel
	err := s.MongoRepo.Find(s.ctx, mongoModelCollection, &mongoModels, bson.M{}, opt)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.savedMongoModels[:limit], mongoModels)
}

func (s *MongoContainerShould) TestFind_WithSkipping_ReturnSuccess() {
	skipping := int64(10)
	opt := options.Find().SetSort(bson.M{"order": 1}).SetSkip(skipping)

	var mongoModels []*MongoModel
	err := s.MongoRepo.Find(s.ctx, mongoModelCollection, &mongoModels, bson.M{}, opt)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.savedMongoModels[skipping:], mongoModels)
}

func (s *MongoContainerShould) TestFind_WithFilter_ReturnSuccess() {
	find := bson.M{"tags": bson.A{"multiple_of_three", "multiple_of_five"}}
	opt := options.Find().SetSort(bson.M{"order": 1})

	var mongoModels []*MongoModel
	err := s.MongoRepo.Find(s.ctx, mongoModelCollection, &mongoModels, find, opt)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.savedMongoModelsByTags["multiple_of_three,multiple_of_five"], mongoModels)
}

func (s *MongoContainerShould) TestFind_WithNoAppropriateMongoModels_ReturnSuccess() {
	find := bson.M{"tags": bson.A{"nonexistent"}}
	opt := options.Find().SetSort(bson.M{"order": 1})

	var mongoModels []*MongoModel
	err := s.MongoRepo.Find(s.ctx, mongoModelCollection, &mongoModels, find, opt)

	assert.NoError(s.T(), err)
	assert.Nil(s.T(), mongoModels)
}

func reverseOrder(mongoModels []*MongoModel) {
	for left, right := 0, len(mongoModels)-1; left < right; left, right = left+1, right-1 {
		mongoModels[left], mongoModels[right] = mongoModels[right], mongoModels[left]
	}
}
