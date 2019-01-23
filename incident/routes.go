package incident

import (
	"github.com/gin-gonic/gin"
)

func IncidentRouter(incidentService Service, router *gin.RouterGroup) {

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
