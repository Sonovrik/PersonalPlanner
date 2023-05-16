package telegram

import (
	"PersonalPlanner/core"
	"context"
	"flag"
	"log"

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
	e.server.Start(ctx)

	return nil
}

func (e *Engine) Stop(_ context.Context) error {
	return nil
}
