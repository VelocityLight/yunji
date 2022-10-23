package api

import (
	"net/http"

	"yunji/configs"
	"yunji/internal/service/store"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	Gin   *gin.Engine
	store *store.Store
}

func NewHTTPHandler(g *gin.Engine, config *configs.ConfigYaml) *HTTPHandler {
	h := &HTTPHandler{
		Gin:   g,
		store: store.NewStore(config),
	}

	v1 := g.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "hello api v1")
		})

		resources := v1.Group("/costs")
		resources.GET("/", h.GetCostByTeam)

		bills := v1.Group("/bills")

		bills.GET("/used-by-tags", h.GetUsedByTags)
		bills.GET("/component-tags", h.GetTags)
		bills.GET("/services", h.GetServices)
		bills.GET("/trend", h.GetTrending)

		realtime := v1.Group("/hacker")
		realtime.POST("/trigger", h.PostHackerTrigger)
		realtime.GET("/realtime", h.GetRealTime)
	}

	return h
}

func (h *HTTPHandler) Shutdown() error {
	return h.store.Shutdown()
}

func (h *HTTPHandler) GetStore() *store.Store {
	return h.store
}
