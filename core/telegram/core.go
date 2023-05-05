package telegram

import (
	"context"
	"flag"
	"log"

	"github.com/go-telegram/bot"
)

func MustToken() string {
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

type Engine struct {
	server *bot.Bot
}

func New(engineToken, weatherAPIToken string) (*Engine, error) {
	handlers := HandlerOptions()

	botEngine, err := bot.New(engineToken, handlers...)
	if err != nil {
		return nil, err
	}

	server := &Engine{
		server: botEngine,
	}

	return server, nil
}

func (e *Engine) Run(ctx context.Context) error {
	e.server.Start(ctx)

	return nil
}

func (e *Engine) Stop(_ context.Context) error {
	return nil
}
