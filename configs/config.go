package configs

import "github.com/jinzhu/configor"

// Database configuration
type ConfigYaml struct {
	Feishu struct {
		AppId     string `required:"false"`
		AppSecret string `required:"false"`
	}

	// secret config
	Secret
}

type Secret struct {
	DSN string `yaml:"-"`
}

var Config = &ConfigYaml{}

// Load config from file into 'Config' variable
func LoadConfig(file string) {
	configor.Load(Config, file)
}
