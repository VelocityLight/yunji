package api

import (
	"net/http"

	"github.com/gin-contrib/static"
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

func RouteWebsite(g *gin.Engine, file string) *gin.Engine {
	routeHtml(g, file)
	return g
}

func routeHtml(router *gin.Engine, file string) {
	router.Use(
		static.Serve("/", static.LocalFile(file, true)),
	)
	router.Use(
		static.Serve("/static", static.LocalFile(file, true)),
	)
	homePages := router.Group("/home")
	{
		homePages.GET("/*any", func(c *gin.Context) {
			c.FileFromFS("/", http.Dir(file))
		})
	}
}

func routeError(router *gin.Engine) {
	router.Use(APIErrorJSONReporter())
}

func routeCors(router *gin.Engine) {
	router.Use(Cors())
}
