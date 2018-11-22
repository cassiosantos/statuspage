package component

import (
	"github.com/gin-gonic/gin"
)

func ComponentRouter(componentRepo Repository, router *gin.Engine) {

	componentService := NewService(componentRepo)
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
		componentsRouter.GET("/group/:group", componentController.FindByGroup)
	}
}
