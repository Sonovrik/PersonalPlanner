package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"PersonalPlanner/core/telegram"
)

// Engine part
const (
	TelegramEngine = iota
)

// Weather part
const (
	Yandex = iota
)

func mustEngineToken(engineType int) string {
	var token string

	switch engineType {
	case TelegramEngine:
		token = telegram.MustToken()
	default:
		log.Fatalln("Wrong engine")
	}

	return token
}

func main() {
	engineToken := mustEngineToken(TelegramEngine)

	ctx, stop := context.WithCancel(context.Background())

	engine, err := telegram.New(engineToken)
	if err != nil {
		log.Fatalln(err)
	}

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
