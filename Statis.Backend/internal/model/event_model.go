package model

type Event struct {
	ID          int `gorm:"primaryKey"`
	Identifier  string
	Title       string
	Description string
	Type        string // EventType
	Date        string
	IncidentID  int
	Incident    Incident
}
