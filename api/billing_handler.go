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

func (h *HTTPHandler) GetUsedByTags(c *gin.Context) {
	resp, err := h.store.Billing.GetUsedByTags(c.Request.Context())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, resp)
}

func (h *HTTPHandler) GetTags(c *gin.Context) {
	resp, err := h.store.Billing.GetTags(c.Request.Context())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, resp)
}
