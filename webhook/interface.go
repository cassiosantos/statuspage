package webhook

import (
	"github.com/involvestecnologia/statuspage/models"
)

// Read implements the read action methods
type Read interface {
	FindWebhook(id string) (models.Webhook, error)
	FindWebhookByNameAndType(name string, webhookType string) (models.Webhook, error)
	ListWebhookByType(webhookType string) ([]models.Webhook, error)
}

// Write implements the write action methods
type Write interface {
	CreateWebhook(webhook models.Webhook) error
	UpdateWebhook(id string, webhook models.Webhook) error
	DeleteWebhook(id string) error
}

// Repository describes the repository where the data will be written and read from
type Repository interface {
	Read
	Write
}

// Service describes the use case
type Service interface {
	Create(webhook models.Webhook) error
	Update(id string, webhook models.Webhook) error
	Delete(id string) error
	List(webhookType string) ([]models.Webhook, error)
	TriggerWebhook(element interface{}, action string) error
	WebhookExists(webhook models.Webhook) bool
	Read
}
