package incident

import (
	"github.com/involvestecnologia/statuspage/models"
)

// Read implements the read action methods
type Read interface {
	GetIncidentsByComponentID(id string) ([]models.Incident, error)
	GetAllIncidents() ([]models.IncidentWithComponentID, error)
	GetIncidentsByMonth(month int) ([]models.IncidentWithComponentID, error)
}

// Write implements the write action methods
type Write interface {
	AddIncidentToComponent(componentID string, incident models.Incident) error
}

// Repository describes the repository where the data will be writen and read from
type Repository interface {
	Read
	Write
}

// Service describes the use case
type Service interface {
	GetIncidents(query string) ([]models.IncidentWithComponentID, error)
	validateMonth(monthArg int) (int, error)
	Read
	Write
}
