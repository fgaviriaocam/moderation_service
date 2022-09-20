package persistence

import (
	"context"
	"moderation_service/domain/models"
	utils "moderation_service/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type ModerationRepositoryImpl struct {
	client *mongo.Client
}

func (r *ModerationRepositoryImpl) Moderation(o models.Moderation) (interface{}, error) {

	moderation := r.client.Database("fury_service").Collection("moderation")
	doc, _ := utils.ToDoc(o)
	result, err := moderation.InsertOne(context.TODO(), doc)

	return result.InsertedID, err
}
