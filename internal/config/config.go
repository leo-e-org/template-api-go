package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"template-api-go/internal/logger"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Logger.Fatal("Error reading configuration file", zap.Error(err))
	}

	for _, k := range viper.AllKeys() {
		v := viper.GetString(k)
		viper.Set(k, os.ExpandEnv(v))
	}
}
