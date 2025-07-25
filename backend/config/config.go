package config

import (
	"github.com/tkanos/gonfig"
)

type DBConfigurations struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

type AppConfig struct {
	Mode string `validate:"required", oneof=dev prod test debug`
	TZ string `validate:"required"`
	HTTPPort string `validate:"required"`
	HTTPRoot string `validate:"required"`
	DBConfigurations
	SSLConfig
}

type SSLConfig struct {
	CertPath string `validate:"omitempty"`
	KeyPath  string `validate:"omitempty"`
}

func GetConfig() DBConfigurations {
	conf := DBConfigurations{}
	gonfig.GetConf("config/config.json", &conf)
	return conf
}