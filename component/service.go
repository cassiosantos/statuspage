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
	if valid, err, failureRef := s.IsValidComponent(component); !valid {
		return failureRef, err
	}
	return s.repo.Insert(component)
}

func (s *ComponentService) UpdateComponent(ref string, component models.Component) error {
	component.Ref = ref
	if component.Name == "" {
		return errors.E(errors.ErrComponentNameIsEmpty)
	}
	c, exist := s.ComponentExists(map[string]interface{}{"name": component.Name})
	if exist && c.Ref != ref {
		return errors.E(fmt.Sprintf(errors.ErrAlreadyExists, c.Name))
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

func (s *ComponentService) IsValidComponent(c models.Component) (bool, error, string) {

	if c.Name == "" {
		return false, errors.E(errors.ErrComponentNameIsEmpty), c.Name
	}
	if _, exist := s.ComponentExists(map[string]interface{}{"name": c.Name}); exist {
		return false, errors.E(fmt.Sprintf(errors.ErrAlreadyExists, c.Name)), c.Name
	}
	if _, exist := s.ComponentExists(map[string]interface{}{"ref": c.Ref}); exist {
		return false, errors.E(fmt.Sprintf(errors.ErrAlreadyExists, c.Ref)), c.Ref
	}
	return true, nil, ""
}
