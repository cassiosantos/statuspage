package models

import (
	"time"
)

// Incident represents some event categorized by one three possible status
type Incident struct {
	Status       int       `json:"status" binding:"required"`
	Description  string    `json:"description,omitempty"`
	Date         time.Time `json:"occurrence_date" binding:"required"`
	ComponentRef string    `json:"component_ref" bson:"component_ref" binding:"required"`
}

const (
	// IncidentStatusOK means "This status is fully operational."
	IncidentStatusOK = 0
	// IncidentStatusUnstable means "You received reports of a problem."
	IncidentStatusUnstable = 1
	// IncidentStatusOutage means "A severe problem happened"
	IncidentStatusOutage = 2
)
