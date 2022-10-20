package dal

import (
	"yunji/configs"
)

func InitDBConnection() {
	configs.LoadConfig("../../../config.yaml")
	config := configs.Config
	Connect(config)
}
