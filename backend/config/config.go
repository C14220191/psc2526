package config

import (
	"github.com/tkanos/gonfig"
)

type Configurations struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func GetConfig() Configurations {
	conf := Configurations{}
	gonfig.GetConf("config/config.json", &conf)
	return conf
}