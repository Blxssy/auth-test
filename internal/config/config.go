package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string `yaml:"env" env-default:"local"`
	Server   `yaml:"server"`
	Database `yaml:"database"`
}

type Server struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
}

type Database struct {
	Dialect  string `yaml:"dialect" default:"postgres"`
	Host     string `yaml:"host" default:"localhost"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"name"`
	Username string `yaml:"username"`
	Password string
}

func InitConfig() *Config {
	configPath := os.Getenv("ConfigPath")

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	cfg.Database.Password = os.Getenv("DB_PASSWORD")

	return &cfg
}
