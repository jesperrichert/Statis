package dto

type IncidentDto struct {
	Identifier      string `json:"identifier"`
	EventIdentifier string `json:"event_identifier"`
	Content         string `json:"content"`
}
