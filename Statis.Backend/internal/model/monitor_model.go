package model

type Monitor struct {
	ID           int `gorm:"primaryKey"`
	Identifier   string
	IsParent     bool
	Group        string
	Name         string
	Description  string
	Type         string // MonitorType
	ScheduleTime int
	TypeData     any               `gorm:"serializer:json"`
	State        string            // MonitorState
	EffectedBy   map[string]string `gorm:"serializer:json"`
	ExternalUrl  string
}
