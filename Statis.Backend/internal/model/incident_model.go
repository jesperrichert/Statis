package model

type Incident struct {
	ID         int `gorm:"primaryKey"`
	Identifier string
	Content    string
	EventID    int
	Event      *Event
}
