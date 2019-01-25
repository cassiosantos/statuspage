package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/incident"
)

func PrometheusRouter(incident incident.Service, component component.Service, router *gin.RouterGroup) {
	prometheusService := NewPrometheusService(incident, component)
	prometheusController := NewPrometheusController(prometheusService)
	incidentRouter := router.Group("/prometheus")
	{
		incidentRouter.POST("/incoming", prometheusController.Incoming)
	}
}
