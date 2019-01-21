package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
	"net/http"
)

type prometheusController struct {
	service Service
}

func NewPrometheusController(service Service) *prometheusController {
	return &prometheusController{service: service}
}

func (prom *prometheusController) Incoming(c *gin.Context) {
	var incoming models.PrometheusIncomingWebhook
	if err := c.ShouldBindJSON(&incoming); err != nil {
		c.JSON(http.StatusBadRequest, "missing required parameter "+err.Error())
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}
	err := prom.service.PrometheusIncoming(incoming)
	if err != nil {
		if err.Error() == errors.ErrInvalidRef {
			c.JSON(http.StatusBadRequest, err.Error())
			c.AbortWithError(http.StatusBadGateway, err)
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, "")
}
