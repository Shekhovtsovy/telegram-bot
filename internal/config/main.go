package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"sync"
)

type Config struct {
	Host string `envconfig:"APP_HOST" default:"localhost"`
	Port string `envconfig:"APP_PORT" default:"8100"`
	Telegram
	Db
}

var (
	config *Config
	mu     sync.RWMutex
)

func initAndGetConfig() (*Config, error) {
	var err error
	err = godotenv.Load()
	config = &Config{}
	err = envconfig.Process("", config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func GetConfig() Config {
	var err error
	mu.RLock()
	defer mu.RUnlock()
	if config == nil {
		config, err = initAndGetConfig()
		if err != nil {
			panic(err)
		}
	}
	return *config
}
