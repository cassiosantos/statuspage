package models

import "time"

type PrometheusIncomingWebhook struct {
	Alerts []PrometheusAlerts `json:"alerts" binding:"required"`
}

type PrometheusAlerts struct {
	Status       string    `json:"status"`
	Incident     Incident  `json:"labels,omitempty" binding:"required"`
	Component    Component `json:"annotations,omitempty" binding:"required"`
	StartsAt     time.Time `json:"startsAt,omitempty" time_format:"2006-01-02T15:04:05Z07:00"`
	EndsAt       time.Time `json:"endsAt" time_format:"2006-01-02T15:04:05Z07:00"`
	GeneratorURL string    `json:"generatorURL,omitempty"`
}
