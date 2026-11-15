package model

type EventModel struct {
	ID          int `gorm:"primaryKey"`
	Identifier  string
	Title       string
	Description string
	Type        string // EventType
	Date        string
	Monitors    []*MonitorModel `gorm:"many2many:monitor_events;"`
	IncidentID  int
	Incident    IncidentModel
}

func (EventModel) TableName() string { return "events" }
