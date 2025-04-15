package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"template-api-go/internal/logger"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Fatal("Error reading configuration file", zap.Error(err))
	}
}
