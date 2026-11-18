package services

import (
	"io"
	"net/http"
	"strings"

	"gorm.io/gorm"
	"jespersen.zip.statis/internal/dto"
	"jespersen.zip.statis/internal/model"
	"jespersen.zip.statis/internal/schedule"
	"jespersen.zip.statis/internal/util"
)

type WebhookService struct {
	DB               *gorm.DB
	WebhookScheduler *schedule.WebhookScheduler
}

func NewWebhookService(db *gorm.DB, scheduler *schedule.WebhookScheduler) *WebhookService {
	return &WebhookService{
		DB:               db,
		WebhookScheduler: scheduler,
	}
}

func (e *WebhookService) GetWebhook(identifier string) (dto.WebhookDto, bool) {
	var webhook model.WebhookModel
	e.DB.Where("identifier = ?", identifier).First(&webhook)

	if util.CheckNull(webhook.Identifier) {
		return dto.WebhookDto{}, false
	}

	return dto.WebhookDto{
		Identifier:   webhook.Identifier,
		Url:          webhook.Url,
		JSONTemplate: webhook.JSONTemplate,
	}, true
}

func (e *WebhookService) DeleteWebhook(identifier string) bool {
	var webhookModel model.WebhookModel
	e.DB.Where("identifier = ?", identifier).First(&webhookModel)
	if util.CheckNull(webhookModel.Identifier) {
		return false
	}
	e.DB.Delete(&webhookModel)
	return true
}

func (e *WebhookService) Validate(dto *dto.WebhookDto) (int, string) {

	if util.CheckNull(dto.Url) ||
		util.CheckNull(dto.Identifier) ||
		util.CheckNull(dto.JSONTemplate) {
		return http.StatusBadRequest, "dto not valid"
	}

	if strings.HasPrefix(dto.JSONTemplate, "http") &&
		strings.HasSuffix(dto.JSONTemplate, ".json") {
		req, err := http.Get(dto.JSONTemplate)
		if err != nil {
			return http.StatusBadRequest, "failed to fetch "
		}
		defer req.Body.Close()
		bytes, _ := io.ReadAll(req.Body)

		dto.JSONTemplate = string(bytes)
	}

	var webhook model.WebhookModel
	e.DB.Where("identifier = ?", dto.Identifier).First(&webhook)
	if len(webhook.Identifier) > 0 {
		return http.StatusConflict, "this webhook already exists."
	}

	return http.StatusOK, ""
}

func (e *WebhookService) Add(dto *dto.WebhookDto) int {
	var webhook model.WebhookModel
	e.DB.Create(&model.WebhookModel{
		Identifier:   dto.Identifier,
		Url:          dto.Url,
		JSONTemplate: dto.JSONTemplate,
	}).Where("identifier = ?", dto.Identifier).First(&webhook)

	e.WebhookScheduler.ScheduleWebhook(&webhook)

	return http.StatusOK
}

func (e *WebhookService) Update(identifier string, dto *dto.WebhookDto) int {

	var webhook model.WebhookModel
	e.DB.Where("identifier = ?", identifier).First(&webhook)
	if util.CheckNull(webhook.Identifier) {
		return http.StatusNotFound
	}

	if !util.CheckNull(dto.Url) {
		webhook.Url = dto.Url
	}
	if !util.CheckNull(dto.JSONTemplate) {
		webhook.JSONTemplate = dto.JSONTemplate
	}

	e.WebhookScheduler.ScheduleWebhook(&webhook)
	e.DB.Where("id = ?", webhook.ID).Updates(&webhook)

	return http.StatusOK
}
