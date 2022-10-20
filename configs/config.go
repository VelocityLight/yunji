package configs

import "github.com/jinzhu/configor"

// Database configuration
type ConfigYaml struct {
	TiDB struct {
		UserName string `default:"root"`
		PassWord string `required:"true"`
		Host     string `required:"true"`
		Port     string `default:"3306"`
		DataBase string `required:"true"`
		CharSet  string `default:"utf8"`
		TimeZone string `default:"Asia%2FShanghai"`
	}
}

var Config = &ConfigYaml{}

// Load config from file into 'Config' variable
func LoadConfig(file string) {
	configor.Load(Config, file)
}
