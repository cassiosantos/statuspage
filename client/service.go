package client

import (
	"fmt"

	"github.com/involvestecnologia/statuspage/component"

	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type ClientService struct {
	repo      Repository
	component component.Service
}

func NewService(r Repository, component component.Service) *ClientService {
	return &ClientService{
		repo:      r,
		component: component,
	}
}

func (s *ClientService) CreateClient(client models.Client) (string, error) {
	if s.clientAlreadyExists(map[string]interface{}{"name": client.Name}) {
		return client.Name, errors.E(fmt.Sprintf(errors.ErrAlreadyExists, client.Name))
	}
	if client.Ref != "" {
		if s.clientAlreadyExists(map[string]interface{}{"ref": client.Ref}) {
			return client.Ref, errors.E(fmt.Sprintf(errors.ErrAlreadyExists, client.Ref))
		}
	}
	for _, compRef := range client.Resources {
		if _, exists := s.component.ComponentExists(map[string]interface{}{"ref": compRef}); !exists {
			return client.Ref, errors.E(errors.ErrInvalidRef)
		}
	}

	return s.repo.Insert(client)
}

func (s *ClientService) UpdateClient(clientRef string, client models.Client) error {
	for _, compRef := range client.Resources {
		if _, exists := s.component.ComponentExists(map[string]interface{}{"ref": compRef}); !exists {
			return errors.E(errors.ErrInvalidRef)
		}
	}
	return s.repo.Update(clientRef, client)
}

func (s *ClientService) FindClient(queryParam map[string]interface{}) (models.Client, error) {
	if len(queryParam) == 0 {
		return models.Client{}, errors.E(errors.ErrInvalidQuery)
	}
	return s.repo.Find(queryParam)
}

func (s *ClientService) RemoveClient(clientID string) error {
	return s.repo.Delete(clientID)
}

func (s *ClientService) ListClients() ([]models.Client, error) {
	return s.repo.List()
}

func (s *ClientService) clientAlreadyExists(clientFields map[string]interface{}) bool {
	_, err := s.FindClient(clientFields)
	return err == nil
}
