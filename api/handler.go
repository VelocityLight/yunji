package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	Gin *gin.Engine
}

func NewHTTPHandler() *HTTPHandler {
	g := NewGinRouter()
	h := &HTTPHandler{
		Gin: g,
	}

	v1 := g.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello api v1")
		})

		resources := v1.Group("/resources")
		resources.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello resources")
		})

		bills := v1.Group("/bills")
		bills.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello bills")
		})
	}

	return h
}
