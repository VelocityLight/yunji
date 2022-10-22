package api

import "testing"

func TestRouteHtml(t *testing.T) {
	router := NewGinRouter()
	router = RouteWebsite(router, "../website/build/")

	router.Run(":8080")
}
