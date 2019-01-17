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

func (s *ComponentService) componentExists(name string) (models.Component, bool) {
	c, err := s.repo.Find(map[string]interface{}{"name": name})
	return c, err == nil
}

func (s *ComponentService) CreateComponent(component models.Component) (string, error) {
	_, exist := s.componentExists(component.Name)
	if component.Ref != "" && exist {
		return component.Ref, errors.E(fmt.Sprintf(errors.ErrInvalidRef, component.Ref))
	}
	return s.repo.Insert(component)
}

func (s *ComponentService) UpdateComponent(ref string, component models.Component) error {
	c, exist := s.componentExists(component.Name)
	if exist && c.Ref != ref {
		return errors.E(fmt.Sprintf(errors.ErrAlreadyExists, component.Name))
	}
	return s.repo.Update(ref, component)
}

func (s *ComponentService) FindComponent(queryParam map[string]interface{}) (models.Component, error) {
	return s.repo.Find(queryParam)
}

func (s *ComponentService) ListComponents() ([]models.Component, error) {
	return s.repo.List()
}

func (s *ComponentService) RemoveComponent(id string) error {
	return s.repo.Delete(id)
}
