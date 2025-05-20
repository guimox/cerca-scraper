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

func LoadConfig() (Config, error) {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		return Config{}, fmt.Errorf("SERVER_PORT environment variable is required")
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		return Config{}, fmt.Errorf("BASE_URL environment variable is required")
	}

	readTimeout, err := getRequiredEnvDuration("SERVER_READ_TIMEOUT")
	if err != nil {
		return Config{}, err
	}

	writeTimeout, err := getRequiredEnvDuration("SERVER_WRITE_TIMEOUT")
	if err != nil {
		return Config{}, err
	}

	idleTimeout, err := getRequiredEnvDuration("SERVER_IDLE_TIMEOUT")
	if err != nil {
		return Config{}, err
	}

	config := Config{
		ServerPort: port,
		BaseURL:    baseURL,
		Server: ServerConfig{
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			IdleTimeout:  idleTimeout,
		},
	}

	return config, nil
}

func (c Config) GetServerAddress() string {
	return fmt.Sprintf(":%s", c.ServerPort)
}
