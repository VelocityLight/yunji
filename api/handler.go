package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	Gin *gin.Engine
}

func NewHTTPHandler(g *gin.Engine) *HTTPHandler {
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
		bills.GET("/trend", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": Response{
				Body: []TrendDTO{
					{
						Time: "2021-10",
						Cost: 1000,
						Tag:  "A",
					},
					{
						Time: "2022-11",
						Cost: 2000,
						Tag:  "A",
					},
					{
						Time: "2022-12",
						Cost: 1000,
						Tag:  "B",
					},
				},
			}})
		})

	}

	return h
}

type Response struct {
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"body"`
}

type TrendDTO struct {
	Time string `json:"time,omitempty"`
	Cost int    `json:"cost,omitempty"`
	Tag  string `json:"tag,omitempty"`
}
