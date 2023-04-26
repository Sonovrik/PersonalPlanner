package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"PersonalPlanner/core/telegram"
)

func mustToken() string {
	token := flag.String(
		"token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if token == nil || *token == "" {
		log.Fatalln("token is not specified")
	}

	return *token
}

func main() {
	token := mustToken()

	ctx, stop := context.WithCancel(context.Background())

	engine, err := telegram.New(token)
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
