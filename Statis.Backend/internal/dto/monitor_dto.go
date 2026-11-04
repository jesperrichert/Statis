package dto

type MonitorDto struct {
	Id          string `json:"id"`
	IsParent    bool   `json:"is_parent"`
	Group       string `json:"group"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"` // MonitorType
	ExternalUrl string `json:"external_url"`
	CheckData   any    `json:"check_data"`
	// IsOnline     bool
	// IsEffected   bool
}
