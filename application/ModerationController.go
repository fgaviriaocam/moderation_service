package application

import (
	"moderation_service/domain/models"
	"moderation_service/domain/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ModerationController struct {
	moderationService service.ModerationService
}

func InitModerationController(router *gin.Engine) {

	customerController := ModerationController{
		moderationService: service.InitModerationServiceImpl(),
	}

	router.POST("/moderation", customerController.Moderation)
}

func (r *ModerationController) Moderation(c *gin.Context) {

	var moderation models.Moderation
	disregarded, _ := strconv.ParseBool(c.Request.Header.Get("disregarded"))

	if err := c.ShouldBindJSON(&moderation); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.Response{})
		return
	}

	result, response := r.moderationService.ModerationProccess(moderation, disregarded)

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, result)
}
