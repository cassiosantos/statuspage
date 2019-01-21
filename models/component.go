package models

// Component represents a host, database or any resource with an address
type Component struct {
	Ref       string     `json:"ref,omitempty"`
	Name      string     `json:"name" binding:"required"`
	Address   string     `json:"address"`
	Incidents []Incident `json:"incidents_history"`
}
