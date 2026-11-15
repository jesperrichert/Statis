package model

type WebhookModel struct {
	ID           int `gorm:"primaryKey"`
	Identifier   string
	Url          string `json:"url"`
	JSONTemplate string `json:"json_template"`
}

func (WebhookModel) TableName() string { return "webhooks" }
