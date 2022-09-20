package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Moderation struct {
	ID            primitive.ObjectID `json:"_id,omitempty"`
	Input         string             `json:"input"`
	Ouput         string             `json:"ouput,omitempty"`
	ApplicationID string             `json:"application_id"`
	CreatedDate   time.Time          `json:"created_date,omitempty"`
	LangProccesed []LangProccesed    `json:"lang_proccesed,omitempty"`
}
