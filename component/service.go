package component

import (
	"github.com/involvestecnologia/statuspage/models"
)

type ComponentService struct {
	repo Repository
}

func NewService(r Repository) *ComponentService {
	return &ComponentService{repo: r}
}

func (s *ComponentService) ComponentExists(name string) bool {
	_, err := s.repo.GetComponentByName(name)
	return err == nil
}

func (s *ComponentService) AddComponent(component models.Component) error {
	return s.repo.AddComponent(component)
}

func (s *ComponentService) UpdateComponent(id string, component models.Component) error {
	return s.repo.UpdateComponent(id, component)
}

func (s *ComponentService) GetComponent(queryBy string, id string) (models.Component, error) {
	if queryBy == "name" {
		return s.GetComponentByName(id)
	}
	return s.GetComponentById(id)
}

func (s *ComponentService) GetComponentByName(name string) (models.Component, error) {
	return s.repo.GetComponentByName(name)
}

func (s *ComponentService) GetComponentById(id string) (models.Component, error) {
	return s.repo.GetComponentById(id)
}

func (s *ComponentService) GetAllComponents() ([]models.Component, error) {
	return s.repo.GetAllComponents()
}

func (s *ComponentService) GetComponentsByGroup(groupName string) ([]models.Component, error) {
	return s.repo.GetComponentsByGroup(groupName)
}

func (s *ComponentService) DeleteComponent(id string) error {
	return s.repo.DeleteComponent(id)
}
