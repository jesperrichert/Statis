package dto

type EventDto struct {
	Identifier         string   `json:"identifier"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Type               string   `json:"type"` // EventType
	Date               string   `json:"date"`
	IncidentIdentifier string   `json:"incident_identifier"`
	Monitors           []string `json:"monitors"`
}
