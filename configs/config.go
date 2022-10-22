package configs

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Database configuration
type ConfigYaml struct {
	Feishu struct {
		AppId     string `yaml:"appId"`
		AppSecret string `yaml:"appSecret"`
	} `yaml:"feishu"`

	// secret config
	Secret `yaml:"secret"`
}

type Secret struct {
	DSN string `yaml:"dsn"`
}

var Config = &ConfigYaml{}

// Load config from file into 'Config' variable
func LoadConfig(configPath string) {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(fmt.Sprintf("read config file failed, err= %v", err))
	}

	if err := yaml.Unmarshal(content, Config); err != nil {
		panic(fmt.Sprintf("parse config file failed, err= %v", err))
	}
	fmt.Print("dsn: ", Config.DSN)
}
