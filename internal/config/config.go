package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Address string
	}
}

func New() *Config {

	return &Config{}
}

func (c *Config) Load(filePath string) error {

	c.SetDefaults()

	viper.SetConfigFile(filePath)

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return nil
}
