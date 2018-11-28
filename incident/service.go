package incident

import (
	"strconv"

	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type IncidentService struct {
	repo Repository
}

func NewService(r Repository) *IncidentService {
	return &IncidentService{repo: r}
}

func (s *IncidentService) AddIncidentToComponent(componentID string, incident models.Incident) error {
	return s.repo.AddIncidentToComponent(componentID, incident)
}

func (s *IncidentService) GetIncidentsByComponentID(id string) ([]models.Incident, error) {
	return s.repo.GetIncidentsByComponentID(id)
}

func (s *IncidentService) GetIncidents(query string) ([]models.IncidentWithComponentID, error) {
	var incidents []models.IncidentWithComponentID
	if query == "" {
		return s.GetAllIncidents()
	}
	month, err := strconv.Atoi(query)
	if err != nil {
		return incidents, err
	}
	return s.GetIncidentsByMonth(month)
}

func (s *IncidentService) GetAllIncidents() ([]models.IncidentWithComponentID, error) {
	return s.repo.GetAllIncidents()
}

func (s *IncidentService) GetIncidentsByMonth(monthFilter int) ([]models.IncidentWithComponentID, error) {
	month, err := s.validateMonth(monthFilter)
	if err != nil {
		return []models.IncidentWithComponentID{}, err
	}
	return s.repo.GetIncidentsByMonth(month)
}

func (s *IncidentService) validateMonth(month int) (int, error) {
	valid := month > -1 && month < 12
	if !valid {
		return -1, errors.E(errors.ErrInvalidMonth)
	}
	return month, nil
}
