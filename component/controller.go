package component

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mgo "github.com/globalsign/mgo"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type ComponentController struct {
	service Service
}

func NewComponentController(service Service) *ComponentController {
	return &ComponentController{service: service}
}

func (ctrl *ComponentController) Create(c *gin.Context) {
	var newComponent models.Component
	if err := c.ShouldBindJSON(&newComponent); err != nil {
		c.JSON(http.StatusBadRequest, "Missing required parameter")
		return
	}

	_, err := ctrl.service.CreateComponent(newComponent)
	if err != nil {
		if err.Error() == errors.ErrInvalidRef {
			c.JSON(http.StatusConflict, "Component "+newComponent.Name+" already exists")
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusCreated, "")
}

func (ctrl *ComponentController) Update(c *gin.Context) {
	id := c.Param("id")
	var newComponent models.Component
	if err := c.ShouldBindJSON(&newComponent); err != nil {
		c.JSON(http.StatusBadRequest, "Missing required parameter")
		return
	}
	err := ctrl.service.UpdateComponent(id, newComponent)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, "")
			return
		}
		if err.Error() == errors.ErrInvalidRef {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, "")
}

func (ctrl *ComponentController) Find(c *gin.Context) {
	queryBy := c.DefaultQuery("search", "ref")
	id := c.Param("id")
	component, err := ctrl.service.FindComponent(map[string]interface{}{queryBy: id})
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, "")
			return
		}
		if err.Error() == errors.ErrInvalidRef {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, component)
}

func (ctrl *ComponentController) List(c *gin.Context) {
	components, err := ctrl.service.ListComponents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, components)
}

func (ctrl *ComponentController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.service.RemoveComponent(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, "")
			return
		}
		if err.Error() == errors.ErrInvalidRef {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, "")
}
