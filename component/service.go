package component

import (
	"fmt"

	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type ComponentService struct {
	repo Repository
}

func NewService(r Repository) *ComponentService {
	return &ComponentService{repo: r}
}

func (s *ComponentService) ComponentExists(componentFields map[string]interface{}) (models.Component, bool) {
	c, err := s.repo.Find(componentFields)
	return c, err == nil
}

func (s *ComponentService) CreateComponent(component models.Component) (string, error) {
	if _, exist := s.ComponentExists(map[string]interface{}{"name": component.Name}); exist {
		return component.Name, errors.E(fmt.Sprintf(errors.ErrAlreadyExists, component.Name))
	}
	if component.Ref != "" {
		if _, exist := s.ComponentExists(map[string]interface{}{"ref": component.Ref}); exist {
			return component.Ref, errors.E(fmt.Sprintf(errors.ErrAlreadyExists, component.Ref))
		}
	}
	return s.repo.Insert(component)
}

func (s *ComponentService) UpdateComponent(ref string, component models.Component) error {
	c, exist := s.ComponentExists(map[string]interface{}{"name": component.Name})
	if exist && c.Ref != ref {
		return errors.E(fmt.Sprintf(errors.ErrAlreadyExists, component.Name))
	}
	return s.repo.Update(ref, component)
}

func (s *ComponentService) FindComponent(queryParam map[string]interface{}) (models.Component, error) {
	if len(queryParam) == 0 {
		return models.Component{}, errors.E(errors.ErrInvalidQuery)
	}
	return s.repo.Find(queryParam)
}

func (s *ComponentService) ListComponents() ([]models.Component, error) {
	return s.repo.List()
}

func (s *ComponentService) RemoveComponent(id string) error {
	return s.repo.Delete(id)
}
