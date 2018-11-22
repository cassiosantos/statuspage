package component

import (
	"github.com/involvestecnologia/statuspage/models"
)

// Read implements the read action methods
type Read interface {
	GetAllComponents() ([]models.Component, error)
	GetComponentsByGroup(groupName string) ([]models.Component, error)
	GetComponentById(id string) (models.Component, error)
	GetComponentByName(name string) (models.Component, error)
}

// Write implements the write action methods
type Write interface {
	AddComponent(component models.Component) error
	UpdateComponent(id string, component models.Component) error
	DeleteComponent(id string) error
}

// Repository describes the repository where the data will be writen and read from
type Repository interface {
	Read
	Write
}

// Service describes the use case
type Service interface {
	GetComponent(queryBy string, id string) (models.Component, error)
	ComponentExists(name string) bool
	Read
	Write
}
