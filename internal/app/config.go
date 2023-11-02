package app

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"petProject/internal/app/storage/db/postgres"
	"sync"
)

type Config struct {
	Host     string `env:"SERVER_HOST" env-required:"true"`
	Port     string `env:"SERVER_PORT" env-required:"true"`
	LogLevel string `env:"LOG_LEVEL" env-default:"debug"`
	Store    postgres.Config
}

var once sync.Once

func NewConfig(configPath string) *Config {
	var cfg Config
	once.Do(func() {
		err := cleanenv.ReadConfig(configPath, &cfg)
		if err != nil {
			log.Fatalln(err)
		}
	})
	return &cfg
}
