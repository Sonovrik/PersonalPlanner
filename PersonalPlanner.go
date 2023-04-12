package main

import (
	"context"
	"flag"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}

func main() {
	token := mustToken()

	ctx, stop := context.WithCancel(context.Background())

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		panic(err)
	}

	go b.Start(ctx)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		go func() {
			<-shutdownCtx.Done()

			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatalln("graceful shutdown timed out.. forcing exit")
			}
		}()

		if ok, err := b.Close(shutdownCtx); !ok || err != nil {
			log.Fatalln("Bot can't closed ", err)
		}

		stop()
		cancel()
	}()

	<-ctx.Done()
	log.Print("Bot stopped")
}
