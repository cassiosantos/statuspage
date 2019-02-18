package prometheus

import (
	"strconv"
	"time"

	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/models"
)

type prometheusService struct {
	incident  incident.Service
	component component.Service
}

//NewPrometheusService creates implementation of the Service interface
func NewPrometheusService(incident incident.Service, component component.Service) Service {
	return &prometheusService{incident: incident, component: component}
}

func (svc *prometheusService) ProcessIncomingWebhook(incoming models.PrometheusIncomingWebhook) error {
	for _, alerts := range incoming.Alerts {
		ref, err := svc.component.CreateComponent(alerts.Component)
		alerts.Component.Ref = ref
		if svc.shouldFail(&alerts,err) {
			return err
		}
		incident, err := svc.LabelToIncident(alerts)
		if err != nil {
			return err
		}
		if err := svc.incident.CreateIncidents(incident); err != nil {
			if svc.shouldFail(&alerts, err) {
				return err
			}
		}
	}
	return nil
}

func (svc *prometheusService) shouldFail(alerts *models.PrometheusAlerts, err error) bool {
	switch err.(type) {
	case *errors.ErrComponentNameIsEmpty:
		return true
	case *errors.ErrComponentNameAlreadyExists:
		if err = svc.addExistingComponentRef(alerts); err != nil {
			return true
		}
		return false
	case *errors.ErrComponentRefAlreadyExists:
		return false
	case *errors.ErrIncidentStatusIgnored:
		return false
	case nil:
		return false
	default:
		return true
	}
}

func (svc *prometheusService) addExistingComponentRef(alerts *models.PrometheusAlerts) (error) {
	c, err := svc.component.FindComponent(map[string]interface{}{"name": alerts.Component.Name})
	if err != nil {
		return err
	}
	alerts.PrometheusLabel.ComponentRef = c.Ref
	return nil
}

func (svc *prometheusService) LabelToIncident(p models.PrometheusAlerts) (inc models.Incident, err error) {
	if p.PrometheusLabel.Date.IsZero() {
		p.PrometheusLabel.Date = time.Now()
	}
	if p.Status == "resolved" {
		p.PrometheusLabel.Status = "1"
	}
	status, err := strconv.Atoi(p.PrometheusLabel.Status)
	if err != nil {
		return inc, err
	}
	inc.Status = status
	inc.Date = p.PrometheusLabel.Date
	inc.ComponentRef = p.PrometheusLabel.ComponentRef
	return inc, nil
}
