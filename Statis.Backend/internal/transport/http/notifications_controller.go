package http

import (
	"io"
	"net/http"
	"strings"

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

	var webhook model.WebhookModel
	c.DB.Where("identifier = ?", body.Identifier).First(&webhook)
	if len(webhook.Identifier) > 0 {
		ctx.JSON(http.StatusConflict, gin.H{"error": "this webhook already exists."})
		return
	}

	if strings.HasPrefix(body.JSONTemplate, "http") &&
		strings.HasSuffix(body.JSONTemplate, ".json") {
		req, err := http.Get(body.JSONTemplate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		defer req.Body.Close()
		bytes, _ := io.ReadAll(req.Body)
		body.JSONTemplate = string(bytes)
	}

	c.DB.Create(&model.WebhookModel{
		Identifier:   body.Identifier,
		Url:          body.Url,
		JSONTemplate: body.JSONTemplate,
	})

	c.WebhookScheduler.ScheduleWebhook(&webhook)
	ctx.JSON(http.StatusOK, gin.H{})
}

// Post
func (c *NotificationsController) UpdateWebhook(ctx *gin.Context) {
	id := ctx.Param("id")
	var webhook model.WebhookModel
	c.DB.Where("identifier = ?", id).First(&webhook)
	if util.CheckNull(webhook.Identifier) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "webhook not found"})
		return
	}

	var body dto.WebhookDto
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid dto"})
		return
	}

	if !util.CheckNull(body.JSONTemplate) {
		if strings.HasPrefix(body.JSONTemplate, "http") &&
			strings.HasSuffix(body.JSONTemplate, ".json") {
			req, err := http.Get(body.JSONTemplate)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{})
				return
			}
			defer req.Body.Close()
			bytes, _ := io.ReadAll(req.Body)
			webhook.JSONTemplate = string(bytes)
		} else {
			webhook.JSONTemplate = body.JSONTemplate
		}
	}
	if !util.CheckNull(body.Url) {
		webhook.Url = body.Url
	}

	for k, v := range c.WebhookScheduler.Jobs {
		if strings.Split(k, ":")[0] == string(rune(webhook.ID))+webhook.Identifier {
			_ = c.WebhookScheduler.Scheduler.RemoveJob(v.Job.ID())
		}
	}
	c.WebhookScheduler.ScheduleWebhook(&webhook)

	c.DB.Where("id = ?", webhook.ID).Updates(&webhook)
}
