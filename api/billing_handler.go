package api

import (
	"net/http"
	"time"

	"yunji/common"
	"yunji/utils"

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

func (h *HTTPHandler) GetTrending(c *gin.Context) {
	var StartedAt *time.Time
	var EndedAt *time.Time
	if c.Query("started_at") != "" {
		s, err := utils.ParseTime(c.Query("started_at"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		StartedAt = &s
	}
	if c.Query("ended_at") != "" {
		e, err := utils.ParseTime(c.Query("ended_at"))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		EndedAt = &e
	}

	tags, err := utils.ParseCSV(c.Query("tags"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	svc, err := utils.ParseCSV(c.Query("service"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	resp, err := h.store.Billing.GetTrends(c.Request.Context(), common.GetTrendOpts{
		Tags:      tags,
		Service:   svc,
		StartedAt: StartedAt,
		EndedAt:   EndedAt,
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, resp)
}
