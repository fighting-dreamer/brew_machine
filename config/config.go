package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	appPort   int
	appName   string
	logLevel  string
	outletCnt int
}

var appConfiguration *Config

func LoadTestConfig() {
	load("application_test")
}

func Load() {
	load("application")
}

func load(configFileName string) {
	viper.AutomaticEnv()
	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.ReadInConfig()

	appConfiguration = &Config{
		appPort:   getIntOrPanic("app_port"),
		appName:   getStringOrPanic("app_name"),
		logLevel:  getStringOrPanic("log_level"),
		outletCnt: getIntOrPanic("outlet_count"),
	}

	return
}

func AppPort() int {
	return appConfiguration.appPort
}

func AppName() string {
	return appConfiguration.appName
}

func LogLevel() string {
	return appConfiguration.logLevel
}

func OutletCnt() int {
	return appConfiguration.outletCnt
}
