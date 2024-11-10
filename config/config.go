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

// LoadConfig reads config.yml and unmarshals the contents into a Config struct
func LoadConfig() (*Config, error) {
	env := os.Getenv("ENV")
	cfg := &Config{}

	if env == "production" {
		// production 환경일 때 config.yml 파일 불러오기
		file, err := os.ReadFile("config/config.yml")
		if err != nil {
			log.Fatalf("Failed to read config file: %v", err)
			return nil, err
		}

		err = yaml.Unmarshal(file, cfg)
		if err != nil {
			log.Fatalf("Failed to parse config file: %v", err)
			return nil, err
		}
	} else {
		// 개발 환경일 때 .env 파일 로드
		err := godotenv.Load(".env")
		if err != nil {
			log.Printf("No .env file found, using system environment variables")
		}

		// .env 파일이나 환경 변수에서 설정 값 가져오기
		cfg.Server.Port, _ = strconv.Atoi(getEnv("SERVER_PORT", "8080"))
		cfg.Redis.Host = getEnv("REDIS_HOST", "localhost")
		cfg.Redis.Port, _ = strconv.Atoi(getEnv("REDIS_PORT", "6379"))
		cfg.Redis.Password = getEnv("REDIS_PASSWORD", "")
		cfg.MySQL.Host = getEnv("MYSQL_HOST", "localhost")
		cfg.MySQL.Port, _ = strconv.Atoi(getEnv("MYSQL_PORT", "3306"))
		cfg.MySQL.User = getEnv("MYSQL_USER", "root")
		cfg.MySQL.Password = getEnv("MYSQL_PASSWORD", "password")
		cfg.MySQL.Database = getEnv("MYSQL_DATABASE", "push_notification_db")
	}

	return cfg, nil
}

// getEnv는 환경 변수를 가져오거나 기본값을 반환하는 헬퍼 함수입니다.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
