package dto

type MonitorDto[T any] struct {
	Identifier   string          `json:"identifier"`
	IsParent     bool            `json:"is_parent"`
	Group        string          `json:"group"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Type         string          `json:"type"` // MonitorType
	ScheduleTime int             `json:"schedule_time"`
	TypeData     T               `json:"type_data"`
	State        string          `json:"state"` // MonitorState
	Events       []string        `json:"events"`
	AffectedBy   []MonitorEffect `json:"affected_by"`
	ExternalUrl  string          `json:"external_url"`
}

type MonitorEffect struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
