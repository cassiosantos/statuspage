package client

import (
	"github.com/gin-gonic/gin"
)

func ClientRouter(clientRepo Repository, router *gin.RouterGroup) {

	clientService := NewService(clientRepo)
	clientController := NewClientController(clientService)

	clientRouter := router.Group("/client")
	{
		clientRouter.POST("", clientController.Create)
		clientRouter.PATCH("/:clientRef", clientController.Update)
		clientRouter.GET("/:clientRef", clientController.Find)
		clientRouter.DELETE("/:clientRef", clientController.Delete)
	}

	clientsRouter := router.Group("/clients")
	{
		clientsRouter.GET("", clientController.List)
	}
}
