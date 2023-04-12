package telegram

import "github.com/go-telegram/bot"

type Engine struct {
	server *bot.Bot
}

func New(token string) (*Engine, error) {
	opts := HandlerOptions()
	botEngine, err := bot.New(token, opts...)
	if err != nil {
		return nil, err
	}

	server := &Engine{
		server: botEngine,
	}

	return server, nil
}
