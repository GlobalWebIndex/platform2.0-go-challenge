package common

import "github.com/tkanos/gonfig"

type Configuration struct {
	Secret          string
	IsDemo          bool
	DbPath          string
	SqlDropPath     string
	SqlCreatePath   string
	SqlPopulatePath string
}

var configuration *Configuration

func InitConfig() {
	config := Configuration{}
	err := gonfig.GetConf("./config.json", &config)
	if err != nil {
		panic(err)
	}
	configuration = &config
}

func GetConfig() *Configuration {
	return configuration
}

func IsDemo() bool {
	return configuration.IsDemo
}
