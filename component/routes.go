package component

import (
	"github.com/gin-gonic/gin"
)

func ComponentRouter(componentService Service, router *gin.RouterGroup) {

	componentController := NewComponentController(componentService)

	componentRouter := router.Group("/component")
	{
		componentRouter.POST("", componentController.Create)
		componentRouter.PATCH("/:ref", componentController.Update)
		componentRouter.GET("/:ref", componentController.Find)
		componentRouter.DELETE("/:ref", componentController.Delete)
	}

	componentsRouter := router.Group("/components")
	{
		componentsRouter.POST("", componentController.List)
	}
}
