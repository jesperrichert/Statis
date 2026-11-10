package dto

type IncidentDto struct {
	Identifier string `json:"identifier"`
	EventId    string `json:"event_id"`
	Content    string `json:"content"`
}
