package mock

import (
	"fmt"
	"time"

	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/models"
)

type MockIncidentDAO struct {
	incidents []models.Incident
}

func NewMockIncidentDAO() incident.Repository {
	return &MockIncidentDAO{
		incidents: []models.Incident{
			{
				ComponentRef: ZeroTimeHex,
				Description:  "status ok",
				Status:       0,
			},
			{
				ComponentRef: OneSecTimeHex,
				Description:  "status outage",
				Status:       2,
			},
		},
	}
}

func (m *MockIncidentDAO) Find(query map[string]interface{}) ([]models.Incident, error) {
	var incidents []models.Incident

	for _, i := range m.incidents {
		if i.ComponentRef == query["component_ref"] {
			incidents = append(incidents, i)
		}
	}
	if len(incidents) > 0 {
		return incidents, nil
	}
	return incidents, errors.E(errors.ErrNotFound)
}
func (m *MockIncidentDAO) List(start time.Time, end time.Time) ([]models.Incident, error) {
	var inc []models.Incident
	for _, i := range m.incidents {
		if (i.Date.Before(end) && i.Date.Before(end)) || (start.IsZero() && end.IsZero()) {
			inc = append(inc, i)
		}
	}
	return inc, nil
}
func (m *MockIncidentDAO) Insert(incident models.Incident) error {
	m.incidents = append(m.incidents, incident)
	return nil
}

type MockFailureIncidentDAO struct {
}

func NewMockFailureIncidentDAO() incident.Repository {
	return &MockFailureIncidentDAO{}
}

func (m *MockFailureIncidentDAO) Find(query map[string]interface{}) ([]models.Incident, error) {
	return []models.Incident{}, fmt.Errorf("Failure DAO")
}
func (m *MockFailureIncidentDAO) List(start time.Time, end time.Time) ([]models.Incident, error) {
	return []models.Incident{}, fmt.Errorf("Failure DAO")
}
func (m *MockFailureIncidentDAO) Insert(incident models.Incident) error {
	return fmt.Errorf("Failure DAO")
}
