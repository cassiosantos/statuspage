package component

import (
	"github.com/gin-gonic/gin"
)

func ComponentRouter(componentService Service, router *gin.RouterGroup) {

	componentController := NewComponentController(componentService)

	componentRouter := router.Group("/component")
	{
		componentRouter.POST("", componentController.Create)
		componentRouter.PATCH("/:id", componentController.Update)
		componentRouter.GET("/:id", componentController.Find)
		componentRouter.DELETE("/:id", componentController.Delete)
	}

	componentsRouter := router.Group("/components")
	{
		componentsRouter.GET("", componentController.List)
	}
}
