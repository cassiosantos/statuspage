package prometheus

import (
	"fmt"
	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/models"
)

type prometheusService struct {
	incident  incident.Service
	component component.Service
}

func NewPrometheusService(incident incident.Service, component component.Service) *prometheusService {
	return &prometheusService{incident: incident, component: component}
}

func (svc *prometheusService) PrometheusIncoming(incoming models.PrometheusIncomingWebhook) error {
	for _, alerts := range incoming.Alerts {
		if ref, err := svc.component.CreateComponent(alerts.Component); err != nil {
			if err.Error() != fmt.Sprintf(errors.ErrAlreadyExists, ref) {
				return err
			}
		}
		if err := svc.incident.CreateIncidents(alerts.Component.Ref, alerts.Incident); err != nil {
			return err
		}
	}
	return nil
}