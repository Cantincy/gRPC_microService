package config

import (
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath("./grpc/user/config")
	viper.SetConfigFile("yaml")
	viper.SetConfigName("config")
	viper.ReadInConfig()
}
