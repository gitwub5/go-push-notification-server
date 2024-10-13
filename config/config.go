package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
	MySQL struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`
}

// LoadConfig reads config.yml and unmarshals the contents into a Config struct
func LoadConfig() (*Config, error) {
	file, err := os.ReadFile("config/config.yml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
		return nil, err
	}

	return &cfg, nil
}
