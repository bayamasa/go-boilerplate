package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Env      string
	Database struct {
		Name string
		Host struct {
			Read  string
			Write string
		}
		Port     string
		User     string
		Password string
	}
}

func NewConfig() (*Config, error) {
	var cfg Config

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	viper.SetConfigName(env + ".yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/environment")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
