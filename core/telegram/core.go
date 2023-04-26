package telegram

import (
	"context"

	"github.com/go-telegram/bot"
)

type Engine struct {
	server *bot.Bot
}

func New(token string) (*Engine, error) {
	handlers := HandlerOptions()

	botEngine, err := bot.New(token, handlers...)
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
