package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/schedule"
	"jespersen.zip.statis/internal/util"
)

type NotificationsController struct {
	DB               *gorm.DB
	WebhookScheduler *schedule.WebhookScheduler
}

func NewNotificationsController(db *gorm.DB, webhookScheduler *schedule.WebhookScheduler) *NotificationsController {
	return &NotificationsController{
		DB:               db,
		WebhookScheduler: webhookScheduler,
	}
}

// Post
func (c *NotificationsController) WebhookPost(ctx *gin.Context) {

	var body dto.WebhookDto
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing body"})
		return
	}
	if util.CheckNull(body.Url) ||
		util.CheckNull(body.Identifier) ||
		util.CheckNull(body.JSONTemplate) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "dto not valid})"})
		return
	}

	var webhook model.Webhook
	c.DB.Where("identifier = ?", body.Identifier).First(&webhook)
	if len(webhook.Identifier) > 0 {
		ctx.JSON(http.StatusConflict, gin.H{"error": "this webhook already exists."})
		return
	}

	c.DB.Create(&model.Webhook{
		Identifier:   body.Identifier,
		Url:          body.Url,
		JSONTemplate: body.JSONTemplate,
	})
	c.WebhookScheduler.ScheduleWebhook(webhook)

	ctx.JSON(http.StatusOK, gin.H{})
}
