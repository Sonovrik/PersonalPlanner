package telegram

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	startCMD   = "/start"
	helpCMD    = "/help"
	weatherCMD = "/weather"
)

func HelpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

}

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

}

func WeatherHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

}

func HandlerOptions() []bot.Option {
	opts := []bot.Option{
		bot.WithMessageTextHandler(startCMD, bot.MatchTypeExact, StartHandler),
		bot.WithMessageTextHandler(helpCMD, bot.MatchTypeExact, HelpHandler),
		bot.WithMessageTextHandler(weatherCMD, bot.MatchTypeExact, WeatherHandler),
	}

	return opts
}
