package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	ServerPort string
	BaseURL    string
	Server     ServerConfig
}

type ServerConfig struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

const (
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 10 * time.Second
	defaultIdleTimeout  = 120 * time.Second
)

func LoadConfig() (Config, error) {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		return Config{}, fmt.Errorf("SERVER_PORT environment variable is required")
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		return Config{}, fmt.Errorf("BASE_URL environment variable is required")
	}

	config := Config{
		ServerPort: port,
		BaseURL:    baseURL,
		Server: ServerConfig{
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
			IdleTimeout:  defaultIdleTimeout,
		},
	}

	return config, nil
}

func (c Config) GetServerAddress() string {
	return fmt.Sprintf(":%s", c.ServerPort)
}
