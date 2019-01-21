package prometheus

import "github.com/involvestecnologia/statuspage/models"

type Service interface {
	ProcessIncomingWebhook(models.PrometheusIncomingWebhook) error
}
