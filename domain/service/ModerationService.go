package service

import "moderation_service/domain/models"

type ModerationService interface {
	ModerationProccess(order models.Moderation) (interface{}, models.Response)
}
