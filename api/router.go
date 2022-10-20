package api

import (
	"github.com/gin-gonic/gin"
)

// Create gin-routers
func Routers() (router *gin.Engine) {
	router = gin.Default()

	// Error
	routeError(router)

	// Cors
	routeCors(router)

	// Real REST-API registry
	routeRestAPI(router)

	return router
}

func routeError(router *gin.Engine) {
	router.Use(APIErrorJSONReporter())
}

func routeCors(router *gin.Engine) {
	router.Use(Cors())
}

// REST-API Function
func routeRestAPI(router *gin.Engine) {

}
