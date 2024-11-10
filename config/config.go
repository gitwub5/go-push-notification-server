package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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

// LoadConfig reads config.yml and overwrites it with environment variables if available
func LoadConfig() (*Config, error) {
	// .env 파일 로드
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, continuing without it")
	}

	file, err := os.ReadFile("config/config.yml")
	if err != nil {
		log.Printf("Failed to read config file: %v", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Printf("Failed to parse config file: %v", err)
	}

	// Overwrite with environment variables if available
	if port := os.Getenv("SERVER_PORT"); port != "" {
		if parsedPort, err := strconv.Atoi(port); err == nil {
			cfg.Server.Port = parsedPort
		}
	}
	if redisHost := os.Getenv("REDIS_HOST"); redisHost != "" {
		cfg.Redis.Host = redisHost
	}
	if redisPort := os.Getenv("REDIS_PORT"); redisPort != "" {
		if parsedPort, err := strconv.Atoi(redisPort); err == nil {
			cfg.Redis.Port = parsedPort
		}
	}
	if redisPassword := os.Getenv("REDIS_PASSWORD"); redisPassword != "" {
		cfg.Redis.Password = redisPassword
	}
	if mysqlHost := os.Getenv("MYSQL_HOST"); mysqlHost != "" {
		cfg.MySQL.Host = mysqlHost
	}
	if mysqlPort := os.Getenv("MYSQL_PORT"); mysqlPort != "" {
		if parsedPort, err := strconv.Atoi(mysqlPort); err == nil {
			cfg.MySQL.Port = parsedPort
		}
	}
	if mysqlUser := os.Getenv("MYSQL_USER"); mysqlUser != "" {
		cfg.MySQL.User = mysqlUser
	}
	if mysqlPassword := os.Getenv("MYSQL_PASSWORD"); mysqlPassword != "" {
		cfg.MySQL.Password = mysqlPassword
	}
	if mysqlDatabase := os.Getenv("MYSQL_DATABASE"); mysqlDatabase != "" {
		cfg.MySQL.Database = mysqlDatabase
	}

	return &cfg, nil
}
