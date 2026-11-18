package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/services"
	"jespersen.zip.statis/internal/util"
)

type NotificationsController struct {
	DB             *gorm.DB
	WebhookService *services.WebhookService
}

func NewNotificationsController(db *gorm.DB, service *services.WebhookService) *NotificationsController {
	return &NotificationsController{
		DB: db,

		WebhookService: service,
	}
}

// GET
func (c *NotificationsController) GetWebhook(ctx *gin.Context) {

	webhookDto, ok := c.WebhookService.GetWebhook(ctx.Param("id"))
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, webhookDto)
}

// Post
func (c *NotificationsController) WebhookPost(ctx *gin.Context) {

	var body dto.WebhookDto
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error parsing body"})
		return
	}

	status, message := c.WebhookService.Validate(&body)
	if status != http.StatusOK {
		ctx.JSON(status, util.JsonResult(message, map[string]any{}))
		return
	}
	status = c.WebhookService.Add(&body)
	if status != http.StatusCreated {
		ctx.JSON(status, util.JsonResult("failed to create", map[string]any{}))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// Post
func (c *NotificationsController) UpdateWebhook(ctx *gin.Context) {

	var body dto.WebhookDto
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid dto"})
		return
	}

	status, message := c.WebhookService.Validate(&body)
	if status != http.StatusOK {
		ctx.JSON(status, util.JsonResult(message, map[string]any{}))
		return
	}
	status = c.WebhookService.Update(ctx.Param("id"), &body)
	if status != http.StatusCreated {
		ctx.JSON(status, util.JsonResult("failed to create", map[string]any{}))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

// GET
func (c *NotificationsController) DeleteWebhook(ctx *gin.Context) {

	ok := c.WebhookService.DeleteWebhook(ctx.Param("id"))
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
