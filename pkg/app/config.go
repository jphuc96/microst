package app

import (
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("microst.yaml")
	return viper.ReadInConfig()
}
