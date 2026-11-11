package dto

type WebhookDto struct {
	Identifier   string `json:"identifier"`
	Url          string `json:"url"`
	JSONTemplate string `json:"json_template"`
}
