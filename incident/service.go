package incident

import (
	"strconv"
	"time"

	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type IncidentService struct {
	component component.Service
	repo      Repository
}

func NewService(r Repository, component component.Service) *IncidentService {
	return &IncidentService{
		component: component,
		repo:      r}
}

func (s *IncidentService) CreateIncidents(incident models.Incident) error {
	_, err := s.component.FindComponent(map[string]interface{}{"ref": incident.ComponentRef})
	if err != nil {
		return err
	}
	return s.repo.Insert(incident)
}

func (s *IncidentService) FindIncidents(query map[string]interface{}) ([]models.Incident, error) {
	return s.repo.Find(query)
}

func (s *IncidentService) ListIncidents(year string, month string) ([]models.Incident, error) {
	var iComp []models.Incident
	var start, end time.Time
	if year == "" && month == "" {
		return s.repo.List(start, end)
	}

	m, err := s.ValidateMonth(month)
	if err != nil && month != "" {
		return iComp, err
	}

	y, err := s.ValidateYear(year)
	if err != nil && year != "" {
		return iComp, err
	}

	if y == -1 {
		y = time.Now().Year()
	}

	if m == -1 {
		start = time.Date(y, 1, 1, 0, 0, 0, 0, time.UTC)
		end = time.Date(y+1, 1, 0, 23, 59, 59, 0, time.UTC)
		return s.repo.List(start, end)
	}

	start = time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
	end = time.Date(y, time.Month(m+1), 0, 23, 59, 59, 0, time.UTC)

	return s.repo.List(start, end)
}

func (s *IncidentService) ValidateMonth(month string) (int, error) {
	m, err := strconv.Atoi(month)
	if err != nil {
		return -1, err
	}
	valid := m > 0 && m < 13
	if !valid {
		return -1, errors.E(errors.ErrInvalidMonth)
	}
	return m, nil
}

func (s *IncidentService) ValidateYear(year string) (int, error) {
	y, err := strconv.Atoi(year)
	if err != nil {
		return -1, err
	}
	valid := y > 0 && y <= time.Now().Year()
	if !valid {
		return -1, errors.E(errors.ErrInvalidYear)
	}
	return y, nil
}
