package app

import (
	"PersonalPlanner/internal/config"
	"PersonalPlanner/internal/services/core/telegram"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunTelegramBot(ctx context.Context, cfg *config.Config) {
	bot, err := telegram.New(cfg.Core.CoreAPIToken, cfg.Core.WeatherAPIToken)
	if err != nil {
		log.Fatalln(err)
	}

	ctxCancel, stop := context.WithCancel(ctx)
	go func() {
		err = bot.Run(ctxCancel)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	log.Println("Engine started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(ctxCancel, 30*time.Second)
		go func() {
			<-shutdownCtx.Done()

			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatalln("Graceful shutdown timed out.. forcing exit")
			}
		}()

		if err = bot.Stop(shutdownCtx); err != nil {
			log.Fatalln("Engine can't stop ", err)
		}

		stop()
		cancel()
	}()

	<-ctxCancel.Done()
	log.Println("Bot stopped")
}
