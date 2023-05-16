package main

import (
	"PersonalPlanner/utils"
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"PersonalPlanner/core/telegram"
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

	cfg, err := utils.Init(cfgPath)
	if err != nil {
		log.Fatalln("Can't init config: ", err.Error())
	}

	engine, err := telegram.New(cfg.Core.CoreAPIToken, cfg.Core.WeatherAPIToken)
	if err != nil {
		log.Fatalln(err)
	}

	ctx, stop := context.WithCancel(context.Background())
	go func() {
		err = engine.Run(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	log.Println("Engine started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		go func() {
			<-shutdownCtx.Done()

			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatalln("Graceful shutdown timed out.. forcing exit")
			}
		}()

		if err = engine.Stop(shutdownCtx); err != nil {
			log.Fatalln("Engine can't stop ", err)
		}

		stop()
		cancel()
	}()

	<-ctx.Done()
	log.Println("Bot stopped")
}
