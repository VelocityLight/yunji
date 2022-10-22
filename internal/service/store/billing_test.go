package store

import (
	"context"
	"testing"

	"yunji/configs"

	"github.com/stretchr/testify/assert"
)

func TestSelect1000(t *testing.T) {
	configPath := "../../../config.yaml"
	configs.LoadConfig(configPath)

	service := NewStore(configs.Config).Billing
	res, err := service.Select1000(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, true, len(res) > 0)
}
