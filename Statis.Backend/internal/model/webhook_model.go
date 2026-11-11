package model

type Webhook struct {
	ID           int `gorm:"primaryKey"`
	Identifier   string
	Url          string `json:"url"`
	JSONTemplate string `json:"json_template"`
}
