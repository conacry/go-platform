package integration

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/conacry/go-platform/pkg/generator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const randomByteCount = 100

type MongoModel struct {
	ID      primitive.ObjectID `bson:"_id"`
	Order   int                `bson:"order"`
	Tags    []string           `bson:"tags"`
	Payload string             `bson:"payload"`
}

func GetMongoModelWithRandomPayload(order int, tags []string) *MongoModel {
	randomBytes := make([]byte, randomByteCount)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	return &MongoModel{
		ID:      primitive.NewObjectID(),
		Order:   order,
		Tags:    tags,
		Payload: hex.EncodeToString(randomBytes),
	}
}

func GetRandomMongoModel() *MongoModel {
	return GetMongoModelWithRandomPayload(
		generator.RandomNumber(1000000, 9999999),
		[]string{generator.RandomDefaultStr()},
	)
}
