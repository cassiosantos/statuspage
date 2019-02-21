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

//ProcessIncomingWebhook
func (svc *prometheusService) ProcessIncomingWebhook(incoming models.PrometheusIncomingWebhook) (err error) {
	for _, alert := range incoming.Alerts {
		alert.Component, err = svc.getComponent(alert)
		if err != nil {
			return err
		}
		incident, err := svc.LabelToIncident(alert)
		if err != nil {
			return err
		}
		if err := svc.incident.CreateIncidents(incident); err != nil {
			if svc.shouldFail(err) {
				return err
			}
		}
	}
	return nil
}

func (svc *prometheusService) getComponent(alert models.PrometheusAlerts) (models.Component, error) {
	var ref string
	component, err := svc.component.FindComponent(map[string]interface{}{"name": alert.Component.Name})
	if err != nil {
		if _, ok := err.(*errors.ErrNotFound); ok {
			alert.Component.Labels = []string{"Unknown"}
			ref, err = svc.component.CreateComponent(alert.Component)
			if svc.shouldFail(err) {
				return component, err
			}
			alert.Component.Ref = ref
			return alert.Component, err
		}
		return component, err
	}
	return component, err
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
	inc.ComponentRef = p.Component.Ref
	inc.Description = p.PrometheusLabel.Description

	return inc, nil
}

func (svc *prometheusService) shouldFail(err error) bool {
	switch err.(type) {
	case *errors.ErrComponentNameIsEmpty:
		return true
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
