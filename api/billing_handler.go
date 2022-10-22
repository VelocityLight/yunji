package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) GetCostByTeam(c *gin.Context) {
	resp, err := h.store.Billing.GetCostByTeam(c.Request.Context())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, resp)
}
