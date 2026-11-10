package model

type Monitor struct {
	ID          int `gorm:"primaryKey"`
	IsParent    bool
	Group       string
	Name        string
	Description string
	Type        string // MonitorType
	ExternalUrl string
	CheckData   any
	State       string   // MonitorState
	EffectedBy  []string `gorm:"serializer:json"`
}
