package prometheus

import (
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

func (svc *prometheusService) ProcessIncomingWebhook(incoming models.PrometheusIncomingWebhook) error {
	for _, alerts := range incoming.Alerts {
		if alerts.Component.Name == "" {
			return errors.E(errors.ErrComponentNameIsEmpty)
		}
		component, exists := svc.component.ComponentExists(map[string]interface{}{"name": alerts.Component.Name})
		if ! exists {
			ref, err := svc.component.CreateComponent(alerts.Component)
			if err != nil {
				return err
			}
			alerts.Incident.ComponentRef = ref
		} else {
			alerts.Incident.ComponentRef = component.Ref
		}
		if err := svc.incident.CreateIncidents(alerts.Incident); err != nil {
			return err
		}
	}
	return nil
}
