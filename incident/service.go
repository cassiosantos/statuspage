package incident

import (
	"strconv"
	"time"

	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type IncidentService struct {
	repo Repository
}

func NewService(r Repository) *IncidentService {
	return &IncidentService{repo: r}
}

func (s *IncidentService) CreateIncidents(componentID string, incident models.Incident) error {
	return s.repo.Insert(componentID, incident)
}

func (s *IncidentService) FindIncidents(componentID string) ([]models.Incident, error) {
	return s.repo.Find(componentID)
}

func (s *IncidentService) ListIncidents(year string, month string) ([]models.IncidentWithComponentName, error) {
	var iComp []models.IncidentWithComponentName
	var start, end time.Time
	if year == "" && month == "" {
		return s.repo.List(start, end)
	}

	m, err := s.validateMonth(month)
	if err != nil && month != "" {
		return iComp, err
	}

	y, err := s.validateYear(year)
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

func (s *IncidentService) validateMonth(month string) (int, error) {
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

func (s *IncidentService) validateYear(year string) (int, error) {
	y, err := strconv.Atoi(year)
	if err != nil {
		return -1, err
	}
	valid := y > 0 && y < time.Now().Year()
	if !valid {
		return -1, errors.E(errors.ErrInvalidYear)
	}
	return y, nil
}
