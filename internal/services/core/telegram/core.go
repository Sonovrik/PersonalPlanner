package telegram

import (
	"PersonalPlanner/internal/services/core"
	"context"
	"flag"
	"log"
	"time"

	"github.com/go-telegram/bot"
)

const flagCoreTokenName = "coreToken"

func MustToken() string {
	token := flag.String(
		flagCoreTokenName,
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if token == nil || *token == "" {
		log.Fatalln("Core token is not specified")
	}

	return *token
}

type Engine struct {
	server *bot.Bot
}

func New(engineToken, weatherAPIToken string) (core.Core, error) {
	handlers := HandlerOptions()

	botEngine, err := bot.New(engineToken, handlers...)
	if err != nil {
		return nil, err
	}

	server := &Engine{
		server: botEngine,
	}

	mustInitWeather(weatherAPIToken)

	return server, nil
}

func (e *Engine) Run(ctx context.Context) error {
	var tmp int64 = 717143592

	ticker := time.NewTicker(10 * time.Second)

	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				_, err := e.server.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: tmp,
					Text:   "weatherMsg",
				})
				if err != nil {
					return
				}
			case <-quit:
				ticker.Stop()

				return
			}
		}
	}()

	e.server.Start(ctx)

	return nil
}

func (e *Engine) Stop(_ context.Context) error {
	return nil
}
