package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mgo "github.com/globalsign/mgo"
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
		c.JSON(http.StatusBadRequest, "missing required parameter")
		return
	}
	_, err := ctrl.service.CreateClient(newClient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	}
	c.JSON(http.StatusCreated, "")
}

func (ctrl *ClientController) Update(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, "missing required parameter")
		return
	}
	clientID := c.Param("clientId")
	err := ctrl.service.UpdateClient(clientID, client)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, "")
			return
		}
		if err.Error() == errors.ErrInvalidRef {
			c.JSON(http.StatusBadRequest, "")
			return
		}
		c.JSON(http.StatusInternalServerError, "")
	}
	c.JSON(http.StatusOK, "")
}

func (ctrl *ClientController) Find(c *gin.Context) {
	queryBy := c.DefaultQuery("search", "ref")
	qValue := c.Param("clientId")
	client, err := ctrl.service.FindClient(map[string]interface{}{queryBy: qValue})
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, "")
			return
		}
		if err.Error() == errors.ErrInvalidRef {
			c.JSON(http.StatusBadRequest, "")
			return
		}
		c.JSON(http.StatusInternalServerError, "")
	}
	c.JSON(http.StatusOK, client)
}

func (ctrl *ClientController) Delete(c *gin.Context) {
	clientID := c.Param("clientId")
	err := ctrl.service.RemoveClient(clientID)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, "")
			return
		}
		if err.Error() == errors.ErrInvalidRef {
			c.JSON(http.StatusBadRequest, "")
			return
		}
		c.JSON(http.StatusInternalServerError, "")
	}
	c.JSON(http.StatusNoContent, "")
}

func (ctrl *ClientController) List(c *gin.Context) {
	clients, err := ctrl.service.ListClients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	}
	c.JSON(http.StatusOK, clients)
}
