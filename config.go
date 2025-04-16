package main

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port     		string `env:"PORT" default:"8080"`
	DBConnection	string `env:"DB_CONNECTION" default:"./posts.db"`
}

func  LoadConfig()(*Config, error){
	var cfg Config
	err := envconfig.Process("",&cfg)
	return &cfg, err
}