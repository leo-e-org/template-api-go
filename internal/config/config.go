package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
	"template-api-go/internal/logger"
)

func LoadConfig() {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Fatal("Error reading configuration file", zap.Error(err))
	}
}
