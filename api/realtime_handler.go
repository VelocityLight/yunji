package api

import (
	"net/http"
	provider "yunji/internal/app/data_fetcher/data_provider"
	"yunji/internal/pkg"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RealtimeHackerOps struct {
	IsAttack       bool
	AttackResource pkg.AWSResourceType
}

func (h *HTTPHandler) PostHackerTrigger(c *gin.Context) {
	opts := RealtimeHackerOps{}

	if err := c.ShouldBindWith(&opts, binding.JSON); err != nil {
		c.Error(err)
		return
	}

	provider.Config.IsAttack = opts.IsAttack
	provider.Config.AttackResource = opts.AttackResource

	c.JSON(http.StatusOK, opts)
}
