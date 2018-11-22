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

func (s *IncidentService) GetAllIncidents(monthFilter string) ([]models.IncidentWithComponentID, error) {
	if monthFilter == "" {
		return s.repo.GetAllIncidents()
	}
	month, err := validateMonth(monthFilter)
	if err != nil {
		return []models.IncidentWithComponentID{}, err
	}
	return s.repo.GetIncidentsByMonth(month)

}

func validateMonth(monthArg string) (int, error) {
	month, err := strconv.Atoi(monthArg)
	if err != nil {
		return -1, errors.E(errors.ErrInvalidMonth)
	}
	valid := month > -1 && month < 12
	if !valid {
		return -1, errors.E(errors.ErrInvalidMonth)
	}
	return month, nil
}
