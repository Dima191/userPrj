package main

import (
	"flag"
	"log"
	"petProject/internal/app"
	"petProject/pkg/logger"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/config.env", "path to config file")
}

func main() {
	flag.Parse()

	cfg := app.NewConfig(configPath)

	l, err := logger.New(cfg.LogLevel)
	if err != nil {
		log.Fatalln(err)
	}

	s := app.NewAPIServer(cfg, l)
	l.Fatalln(s.Start())
}
