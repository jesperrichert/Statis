package dto

type MonitorDto struct {
	Identifier  string          `json:"identifier"`
	IsParent    bool            `json:"is_parent"`
	Group       string          `json:"group"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Type        string          `json:"type"` // MonitorType
	ExternalUrl string          `json:"external_url"`
	CheckData   any             `json:"check_data"`
	State       string          `json:"state"` // MonitorState
	EffectedBy  []MonitorEffect `json:"effected_by"`
}

type MonitorEffect struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
