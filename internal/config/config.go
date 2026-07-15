package config

import "os"

type Config struct {
	DBUrl    string
	Port     string
	RedisUrl string
}

func NewConfig() *Config {
	return &Config{
		DBUrl:    os.Getenv("DB_URL"),
		Port:     os.Getenv("PORT"),
		RedisUrl: os.Getenv("REDIS_URL"),
	}
}
