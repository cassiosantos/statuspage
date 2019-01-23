package prometheus

import (
	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/mock"
	"github.com/involvestecnologia/statuspage/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewPrometheusService(t *testing.T) {
	componentDAO := mock.NewMockComponentDAO()
	incidentDAO := mock.NewMockIncidentDAO()
	componentService := component.NewService(componentDAO)
	incidentService := incident.NewService(incidentDAO, componentService)

	p := NewPrometheusService(incidentService, componentService)
	assert.Implements(t, (*Service)(nil), p)
}

func TestPrometheusService_ProcessIncomingWebhookReturnErr(t *testing.T) {
	componentDAO := mock.NewMockComponentDAO()
	incidentDAO := mock.NewMockIncidentDAO()
	componentService := component.NewService(componentDAO)
	incidentService := incident.NewService(incidentDAO, componentService)

	p := NewPrometheusService(incidentService, componentService)

	m := models.PrometheusIncomingWebhook{
		Alerts: []models.PrometheusAlerts{
			models.PrometheusAlerts{
				Status: "RESOLVED",
				Incident: models.Incident{
					ComponentRef: mock.ZeroTimeHex,
					Description:  "status ok",
					Status:       0,
				},
				Component: models.Component{
					Ref:     mock.ZeroTimeHex,
					Name:    "first",
					Address: "",
				},
				StartsAt: time.Now(),
				EndsAt: time.Now(),
				GeneratorURL: "ur.com",
			},
		},

	}

	err := p.ProcessIncomingWebhook(m)
	assert.NotNil(t, err)
}
