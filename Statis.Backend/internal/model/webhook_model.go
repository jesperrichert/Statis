package model

type WebhookModel struct {
	ID           int    `gorm:"primaryKey"`
	Identifier   string `gorm:"unique"`
	Url          string
	JSONTemplate string
}

func (WebhookModel) TableName() string { return "webhooks" }
