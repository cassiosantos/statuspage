package client

import "github.com/involvestecnologia/statuspage/models"

// Read implements the read action methods
type Read interface {
	Find(queryParam map[string]interface{}) (models.Client, error)
	List() ([]models.Client, error)
}

// Write implements the write action methods
type Write interface {
	Delete(clientRef string) error
	Insert(client models.Client) (string, error)
	Update(clientRef string, client models.Client) error
}

// Repository describes the repository where the data will be written and read from
type Repository interface {
	Read
	Write
}

// Service describes the use case
type Service interface {
	CreateClient(client models.Client) (string, error)
	FindClient(queryParam map[string]interface{}) (models.Client, error)
	ListClients() ([]models.Client, error)
	RemoveClient(clientRef string) error
	UpdateClient(clientRef string, client models.Client) error
}
