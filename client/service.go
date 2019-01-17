package client

import (
	"fmt"

	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type ClientService struct {
	repo Repository
}

func NewService(r Repository) *ClientService {
	return &ClientService{repo: r}
}

func (s *ClientService) CreateClient(client models.Client) (string, error) {
	if client.Ref != "" && s.clientRefAlreadyExists(client.Ref) {
		return client.Ref, errors.E(fmt.Sprintf(errors.ErrInvalidRef, client.Ref))
	}

	return s.repo.Insert(client)
}

func (s *ClientService) UpdateClient(clientID string, client models.Client) error {
	return s.repo.Update(clientID, client)
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

func (s *ClientService) clientRefAlreadyExists(clientRef string) bool {
	_, err := s.FindClient(map[string]interface{}{"ref": clientRef})
	return err == nil
}
