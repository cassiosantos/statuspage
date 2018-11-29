package incident

import (
	"github.com/gin-gonic/gin"
)

func IncidentRouter(incidentRepo Repository, router *gin.RouterGroup) {

	incidentService := NewService(incidentRepo)
	incidentController := NewIncidentController(incidentService)

	incidentRouter := router.Group("/incident")
	{
		incidentRouter.POST("/:componentId", incidentController.Create)
		incidentRouter.GET("/:componentId", incidentController.Find)
	}

	incidentsRouter := router.Group("/incidents")
	{
		incidentsRouter.GET("", incidentController.List)
	}
}
