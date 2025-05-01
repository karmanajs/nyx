package config

import "time"

type ServerConfig struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func DefaultConfig() *ServerConfig {
	return &ServerConfig{
		Port:    "8080",
		Timeout: 2 * time.Second,
	}
}
