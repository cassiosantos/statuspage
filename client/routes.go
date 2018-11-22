package client

import (
	"github.com/gin-gonic/gin"
)

func ClientRouter(clientRepo Repository, router *gin.Engine) {

	clientService := NewService(clientRepo)
	clientController := NewClientController(clientService)

	clientRouter := router.Group("/client")
	{
		clientRouter.POST("", clientController.Create)
		clientRouter.PATCH("/:clientId", clientController.Update)
		clientRouter.GET("/:clientId", clientController.Find)
		clientRouter.DELETE("/:clientId", clientController.Delete)
	}

	clientsRouter := router.Group("/clients")
	{
		clientsRouter.GET("", clientController.List)
	}
}
