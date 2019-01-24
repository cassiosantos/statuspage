package client

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type ClientController struct {
	service Service
}

func NewClientController(service Service) *ClientController {
	return &ClientController{service: service}
}

func (ctrl *ClientController) Create(c *gin.Context) {
	var newClient models.Client
	if err := c.ShouldBindJSON(&newClient); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ref, err := ctrl.service.CreateClient(newClient)
	if err != nil {
		switch err.Error() {
		case errors.ErrClientNameAlreadyExists:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		case errors.ErrClientRefAlreadyExists:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		case errors.ErrInvalidRef:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusCreated, ref)
}

func (ctrl *ClientController) Update(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	clientRef := c.Param("clientRef")
	err := ctrl.service.UpdateClient(clientRef, client)
	if err != nil {
		if err.Error() == errors.ErrNotFound {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		if err.Error() == errors.ErrInvalidRef {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

func (ctrl *ClientController) Find(c *gin.Context) {
	queryBy := c.DefaultQuery("search", "ref")
	qValue := c.Param("clientRef")
	fmt.Printf("%s: %s\n", queryBy, qValue)
	client, err := ctrl.service.FindClient(map[string]interface{}{queryBy: qValue})
	if err != nil {
		if err.Error() == errors.ErrNotFound {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, client)
}

func (ctrl *ClientController) Delete(c *gin.Context) {
	clientID := c.Param("clientRef")
	err := ctrl.service.RemoveClient(clientID)
	if err != nil {
		if err.Error() == errors.ErrNotFound {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusNoContent, "")
}

func (ctrl *ClientController) List(c *gin.Context) {
	clients, err := ctrl.service.ListClients()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, clients)
}
