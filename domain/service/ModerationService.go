package service

import "moderation_service/domain/models"

type ModerationService interface {
	ModerationProccess(moderation models.Moderation, disregarded bool) (interface{}, models.Response)
}
