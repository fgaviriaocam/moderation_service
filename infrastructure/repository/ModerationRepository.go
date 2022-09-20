package repository

import "moderation_service/domain/models"

type ModerationRepository interface {
	Moderation(o models.Moderation) (interface{}, error)
}
