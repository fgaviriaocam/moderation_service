package persistence

import (
	"moderation_service/infrastructure/repository"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbHelper struct {
	ModerationRepository repository.ModerationRepository
	db                   *mongo.Client
}

func InitDbHelper() (*DbHelper, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	return &DbHelper{
		ModerationRepository: &ModerationRepositoryImpl{client},
		db:                   client,
	}, nil
}
