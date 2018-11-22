package component

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mgo "github.com/globalsign/mgo"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type ComponentController struct {
	service *Service
}

func NewComponentController(service *Service) *ComponentController {
	return &ComponentController{service: service}
}

func (ctrl *ComponentController) Create(c *gin.Context) {
	var newComponent models.Component
	if err := c.ShouldBindJSON(&newComponent); err != nil {
		c.JSON(http.StatusBadRequest, "Missing required parameter")
		return
	}

	if exists := ctrl.service.ComponentExists(newComponent.Name); exists {
		c.JSON(http.StatusConflict, "Component "+newComponent.Name+" already exists")
		return
	}
	err := ctrl.service.AddComponent(newComponent)
	if err != nil {
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
	if exist := ctrl.service.ComponentExists(newComponent.Name); exist {
		c.JSON(http.StatusPreconditionFailed, "A Component named "+newComponent.Name+" already exists")
		return
	}
	err := ctrl.service.UpdateComponent(id, newComponent)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, "")
			return
		}
		if err.Error() == errors.ErrInvalidHexID {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, "")
}

func (ctrl *ComponentController) Find(c *gin.Context) {
	queryBy := c.DefaultQuery("search", "id")
	id := c.Param("id")
	component, err := ctrl.service.GetComponent(queryBy, id)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, "")
			return
		}
		if err.Error() == errors.ErrInvalidHexID {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, component)
}

func (ctrl *ComponentController) FindByGroup(c *gin.Context) {
	group := c.Param("group")
	components, err := ctrl.service.GetComponentsByGroup(group)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, "")
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, components)
}

func (ctrl *ComponentController) List(c *gin.Context) {
	components, err := ctrl.service.GetAllComponents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, components)
}

func (ctrl *ComponentController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.service.DeleteComponent(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, "")
			return
		}
		if err.Error() == errors.ErrInvalidHexID {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, "")
}
