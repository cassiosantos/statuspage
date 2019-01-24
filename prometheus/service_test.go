package prometheus

import (
	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewServicesMock() *prometheusService {
	componentDAO := mock.NewMockComponentDAO()
	incidentDAO := mock.NewMockIncidentDAO()
	componentService := component.NewService(componentDAO)
	incidentService := incident.NewService(incidentDAO, componentService)
	return NewPrometheusService(incidentService, componentService)
}

func NewServiceIncidentFailureMock() *prometheusService {
	componentDAO := mock.NewMockComponentDAO()
	incidentDAO := mock.NewMockFailureIncidentDAO()
	componentService := component.NewService(componentDAO)
	incidentService := incident.NewService(incidentDAO, componentService)
	return NewPrometheusService(incidentService, componentService)
}

func NewServiceComponentFailureMock() *prometheusService {
	componentDAO := mock.NewMockFailureComponentDAO()
	incidentDAO := mock.NewMockIncidentDAO()
	componentService := component.NewService(componentDAO)
	incidentService := incident.NewService(incidentDAO, componentService)
	return NewPrometheusService(incidentService, componentService)
}

func TestPrometheusService_Service(t *testing.T) {
	assert.Implements(t, (*Service)(nil), NewServicesMock())
}

func TestPrometheusService_ProcessIncomingWebhookReturnNil(t *testing.T) {
	newPrometheusMock := mock.PrometheusModel()
	err := NewServicesMock().ProcessIncomingWebhook(newPrometheusMock["ModelComplete"])
	assert.Nil(t, err)
}

func TestPrometheusService_ProcessIncomingWebhookReturnIncidentErr(t *testing.T) {
	newPrometheusMock := mock.PrometheusModel()
	err := NewServiceIncidentFailureMock().ProcessIncomingWebhook(newPrometheusMock["ModelWithoutRef"])
	assert.NotNil(t, err)

	err = NewServiceIncidentFailureMock().ProcessIncomingWebhook(newPrometheusMock["ModelWithoutName"])
	assert.NotNil(t, err)
}

func TestPrometheusService_ProcessIncomingWebhookReturnComponentErr(t *testing.T) {
	newPrometheusMock := mock.PrometheusModel()
	err := NewServiceComponentFailureMock().ProcessIncomingWebhook(newPrometheusMock["ModelWithoutRef"])
	assert.NotNil(t, err)
}
