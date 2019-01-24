package component

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ref, err := ctrl.service.CreateComponent(newComponent)
	if err != nil {
		switch err.Error() {
		case errors.ErrComponentNameAlreadyExists:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		case errors.ErrComponentRefAlreadyExists:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusCreated, ref)
}

func (ctrl *ComponentController) Update(c *gin.Context) {
	id := c.Param("id")
	var newComponent models.Component
	if err := c.ShouldBindJSON(&newComponent); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := ctrl.service.UpdateComponent(id, newComponent)
	if err != nil {
		switch err.Error() {
		case errors.ErrNotFound:
			c.AbortWithError(http.StatusNotFound, err)
			return
		case errors.ErrComponentNameAlreadyExists:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusOK, "")
}

func (ctrl *ComponentController) Find(c *gin.Context) {
	queryBy := c.DefaultQuery("search", "ref")
	id := c.Param("id")
	component, err := ctrl.service.FindComponent(map[string]interface{}{queryBy: id})
	if err != nil {
		if err.Error() == errors.ErrNotFound {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, component)
}

func (ctrl *ComponentController) List(c *gin.Context) {
	components, err := ctrl.service.ListComponents()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, components)
}

func (ctrl *ComponentController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.service.RemoveComponent(id)
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
