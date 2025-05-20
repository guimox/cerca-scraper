package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func getRequiredEnvDuration(key string) (time.Duration, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, fmt.Errorf("%s environment variable is required", key)
	}

	seconds, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a valid number: %v", key, err)
	}

	return time.Duration(seconds) * time.Second, nil
}
