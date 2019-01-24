package mock

import (
	"github.com/involvestecnologia/statuspage/models"
	"time"
)

func PrometheusModel() map[string]models.PrometheusIncomingWebhook {
	return map[string]models.PrometheusIncomingWebhook{
		"ModelComplete": {
			Alerts: []models.PrometheusAlerts{
				{
					Status: "RESOLVED",
					Incident: models.Incident{
						ComponentRef: ZeroTimeHex,
						Description:  "status ok",
						Status:       0,
					},
					Component: models.Component{
						Ref:     ZeroTimeHex,
						Name:    "first",
						Address: "",
					},
					StartsAt:     time.Now(),
					EndsAt:       time.Now(),
					GeneratorURL: "ur.com",
				},
			},
		},
		"ModelWithoutName": {
			Alerts: []models.PrometheusAlerts{
				{
					Status: "RESOLVED",
					Incident: models.Incident{
						ComponentRef: ZeroTimeHex,
						Description:  "status ok",
						Status:       0,
					},
					Component: models.Component{
						Ref:     ZeroTimeHex,
						Address: "",
					},
					StartsAt:     time.Now(),
					EndsAt:       time.Now(),
					GeneratorURL: "ur.com",
				},
			},
		},
		"ModelWithoutRef": {
			Alerts: []models.PrometheusAlerts{
				{
					Status: "RESOLVED",
					Component: models.Component{
						Name:    "newComponent",
						Address: "",
					},
					StartsAt:     time.Now(),
					EndsAt:       time.Now(),
					GeneratorURL: "ur.com",
				},
			},
		},
		"ModelWithoutComponent": {
			Alerts: []models.PrometheusAlerts{
				{
					Status: "RESOLVED",
					Incident: models.Incident{
						ComponentRef: ZeroTimeHex,
						Description:  "status ok",
						Status:       0,
					},
					StartsAt:     time.Now(),
					EndsAt:       time.Now(),
					GeneratorURL: "ur.com",
				},
			},
		},
		"ModelBlank": {},
	}
}
