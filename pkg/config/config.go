package config

import (
	"os"
	"strings"
)

type Config struct {
	data map[string]string
}

func LoadConfig(envPrefix string) (*Config, error) {
	config := &Config{
		data: make(map[string]string),
	}

	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if strings.HasPrefix(parts[0], envPrefix) {
			key := strings.TrimPrefix(parts[0], envPrefix)
			config.data[strings.ToLower(key)] = parts[1]
		}
	}

	return config, nil
}

func (c *Config) Get(key string) string {
	return c.data[strings.ToLower(key)]
}

func (c *Config) Set(key, value string) {
	c.data[strings.ToLower(key)] = value
}
