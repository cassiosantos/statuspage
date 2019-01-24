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
		if ref, err := svc.component.CreateComponent(alerts.Component); err != nil {
			if v := ! svc.isValidErr(err); v {
				return err
			}
			alerts.Incident.ComponentRef = ref
		}
		if err := svc.incident.CreateIncidents(alerts.Incident); err != nil {
			return err
		}
	}
	return nil
}

func (svc *prometheusService) isValidErr(err error) bool {
	switch err.Error() {
	case errors.ErrComponentNameIsEmpty:
		return true
	case errors.ErrComponentNameAlreadyExists:
		return true
	case errors.ErrComponentRefAlreadyExists:
		return true
	default :
		return false
	}
}