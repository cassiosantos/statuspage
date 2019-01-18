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
		c.JSON(http.StatusBadRequest, "missing required parameter "+err.Error())
		return
	}
	componentID := c.Param("componentName")
	err := ctrl.service.CreateIncidents(componentID, newIncident)
	if err != nil {
		if err.Error() == errors.ErrNotFound {
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusCreated, "")
}

func (ctrl *IncidentController) Find(c *gin.Context) {
	componentID := c.Param("componentName")
	incidents, err := ctrl.service.FindIncidents(componentID)
	if err != nil {
		if err.Error() == errors.ErrNotFound {
			c.JSON(http.StatusNotFound, "")
			return
		}
		c.JSON(http.StatusInternalServerError, "")
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
			c.JSON(http.StatusBadRequest, "")
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, incidents)
}
