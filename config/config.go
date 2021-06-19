package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	DB    Database
	Kafka Kafka
}

type Database struct {
	Host     string
	Port     string
	DbName   string
	Username string
	Password string
}

type Kafka struct {
	TopicName string
	Bootstrap Bootstrap
	Consumer  Consumer
}

type Consumer struct {
	GroupId string
}

type Bootstrap struct {
	Servers string
}

var appConfig AppConfig

func Init() {
	viper.SetConfigName("app_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	_ = viper.Unmarshal(&appConfig)
}

func GetAppConfig() AppConfig {
	return appConfig
}
