package incident

import (
	"time"

	"github.com/involvestecnologia/statuspage/models"
)

// Read implements the read action methods
type Read interface {
	Find(componentRef string) ([]models.Incident, error)
	List(startDt time.Time, endDt time.Time) ([]models.IncidentWithComponentName, error)
}

// Write implements the write action methods
type Write interface {
	Insert(componentRef string, incident models.Incident) error
}

// Repository describes the repository where the data will be writen and read from
type Repository interface {
	Read
	Write
}

// Service describes the use case
type Service interface {
	CreateIncidents(componentRef string, incident models.Incident) error
	FindIncidents(componentRef string) ([]models.Incident, error)
	ListIncidents(year string, month string) ([]models.IncidentWithComponentName, error)
	validateMonth(monthArg string) (int, error)
}
