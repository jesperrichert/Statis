package model

type IncidentModel struct {
	ID         int `gorm:"primaryKey"`
	Identifier string
	Content    string
	EventID    int
	Event      *EventModel
}

func (IncidentModel) TableName() string { return "incidents" }
