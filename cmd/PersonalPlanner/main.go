package main

import (
	"PersonalPlanner/internal/app"
	"PersonalPlanner/internal/config"
	"context"
	"flag"
	"log"
)

func mustConfigPath() string {
	var configPath string

	flag.StringVar(&configPath, "configPath", "", "path to config file")
	flag.Parse()

	if configPath == "" {
		configPath = "./cfg/config.yml"
		log.Println("Using default path to config: ", configPath)
	}

	return configPath
}

func main() {
	cfgPath := mustConfigPath()

	cfg, err := config.Init(&cfgPath)
	if err != nil {
		log.Fatalln("Can't init config: ", err.Error())
	}

	app.RunTelegramBot(context.Background(), cfg)
}
