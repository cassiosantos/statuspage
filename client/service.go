package client

import "github.com/involvestecnologia/statuspage/models"

type ClientService struct {
	repo Repository
}

func NewService(r Repository) *ClientService {
	return &ClientService{repo: r}
}

func (s *ClientService) AddClient(client models.Client) error {
	return s.repo.AddClient(client)
}

func (s *ClientService) UpdateClient(clientID string, client models.Client) error {
	return s.repo.UpdateClient(clientID, client)
}

func (s *ClientService) FindClient(queryBy string, clientID string) (models.Client, error) {
	if queryBy == "name" {
		return s.FindById(clientID)
	}
	return s.repo.FindByName(clientID)
}

func (s *ClientService) FindById(clientID string) (models.Client, error) {
	return s.repo.FindById(clientID)
}

func (s *ClientService) FindByName(clientID string) (models.Client, error) {
	return s.repo.FindByName(clientID)
}

func (s *ClientService) DeleteClient(clientID string) error {
	return s.repo.DeleteClient(clientID)
}

func (s *ClientService) ListClients() ([]models.Client, error) {
	return s.repo.ListClients()
}
