package component

import (
	"github.com/involvestecnologia/statuspage/models"
)

// Read implements the read action methods
type Read interface {
	Find(queryParam map[string]interface{}) (models.Component, error)
	List() ([]models.Component, error)
}

// Write implements the write action methods
type Write interface {
	Delete(componentRef string) error
	Insert(component models.Component) (string, error)
	Update(componentRef string, component models.Component) error
}

// Repository describes the repository where the data will be writen and read from
type Repository interface {
	Read
	Write
}

// Service describes the use case
type Service interface {
	componentExists(name string) (models.Component, bool)
	CreateComponent(component models.Component) (string, error)
	FindComponent(queryParam map[string]interface{}) (models.Component, error)
	ListComponents() ([]models.Component, error)
	RemoveComponent(componentRef string) error
	UpdateComponent(componentRef string, component models.Component) error
}
