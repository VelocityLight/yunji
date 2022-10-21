package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create gin-routers
func NewGinRouter() (router *gin.Engine) {
	router = gin.Default()

	// Error
	routeError(router)

	// Cors
	routeCors(router)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}

func routeError(router *gin.Engine) {
	router.Use(APIErrorJSONReporter())
}

func routeCors(router *gin.Engine) {
	router.Use(Cors())
}
