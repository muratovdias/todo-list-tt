package config

import "time"

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
