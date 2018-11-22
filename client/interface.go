package client

import "github.com/involvestecnologia/statuspage/models"

// Read implements the read action methods
type Read interface {
	FindById(clientID string) (models.Client, error)
	FindByName(clientID string) (models.Client, error)
	ListClients() ([]models.Client, error)
}

// Write implements the write action methods
type Write interface {
	AddClient(client models.Client) error
	UpdateClient(clientID string, client models.Client) error
	DeleteClient(clientID string) error
}

// Repository describes the repository where the data will be writen and read from
type Repository interface {
	Read
	Write
}

// Service describes the use case
type Service interface {
	Read
	Write
	FindClient(queryBy string, clientID string) (models.Client, error)
}
