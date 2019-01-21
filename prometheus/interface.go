package prometheus

import "github.com/involvestecnologia/statuspage/models"

type Service interface {
	PrometheusIncoming(models.PrometheusIncomingWebhook) error
}