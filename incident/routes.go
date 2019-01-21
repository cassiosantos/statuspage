package incident

import (
	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/component"
)

func IncidentRouter(incidentRepo Repository, component component.Service, router *gin.RouterGroup) {

	incidentService := NewService(incidentRepo, component)
	incidentController := NewIncidentController(incidentService)

	incidentRouter := router.Group("/incident")
	{
		incidentRouter.POST("", incidentController.Create)
		incidentRouter.GET("/:componentRef", incidentController.Find)
	}

	incidentsRouter := router.Group("/incidents")
	{
		incidentsRouter.GET("", incidentController.List)
	}
}
