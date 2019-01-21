package incident

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type IncidentController struct {
	service Service
}

func NewIncidentController(service Service) *IncidentController {
	return &IncidentController{service: service}
}

func (ctrl *IncidentController) Create(c *gin.Context) {
	var newIncident models.Incident
	if err := c.ShouldBindJSON(&newIncident); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := ctrl.service.CreateIncidents(newIncident)
	if err != nil {
		if err.Error() == errors.ErrNotFound {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, "")
}

func (ctrl *IncidentController) Find(c *gin.Context) {
	queryBy := c.DefaultQuery("search", "component_ref")
	queryValue := c.Param("componentRef")
	incidents, err := ctrl.service.FindIncidents(map[string]interface{}{queryBy: queryValue})
	if err != nil {
		if err.Error() == errors.ErrNotFound {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, incidents)
}

func (ctrl *IncidentController) List(c *gin.Context) {
	mQ := c.Query("month")
	yQ := c.Query("year")

	incidents, err := ctrl.service.ListIncidents(yQ, mQ)
	if err != nil {
		if err.Error() == errors.ErrInvalidYear || err.Error() == errors.ErrInvalidMonth {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, incidents)
}
