package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"time"
)

const configPath = "./config/config.yaml"

type Config struct {
	HttpServer `yaml:"http_server"`
	DB
}

type HttpServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DB struct {
	URI  string `env:"MONGODB_URI"`
	Name string `env:"MONGODB_NAME"`
}

func LoadConfig() *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("can not read config %s", err)
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfgDB DB
	err = cleanenv.ReadEnv(&cfgDB)
	if err != nil {
		log.Fatalf("can not read db config %s", err)
	}

	cfg.DB = cfgDB

	return &cfg
}
