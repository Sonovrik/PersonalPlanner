package main

import (
	"PersonalPlanner/internal/app"
	"PersonalPlanner/internal/config"
	"context"
	"fmt"
	"log"
	"os"
)

func mustConfigPath() string {
	var configPath string

	dir := os.Getenv("CONFIGS_DIR")
	if dir == "" {
		configPath = "./configs/config.yml"
		log.Println("Using default path to config: ", configPath)
	} else {
		configPath = dir + "config.yml"
	}

	return configPath
}

func main() {
	cfgPath := mustConfigPath()

	cfg, err := config.Init(&cfgPath)
	if err != nil {
		log.Fatalln("Can't init config: ", err.Error())
	}

	fmt.Println(cfg.Core.CoreAPIToken)

	app.RunTelegramBot(context.Background(), cfg)
}
