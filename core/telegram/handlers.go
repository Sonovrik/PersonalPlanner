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
	// TODO handle error
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   msgHelp,
	})
}

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// TODO handle error
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   msgHello,
	})
}

func WeatherHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

}

func UnknownHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// TODO handle error
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   msgUnknownCommand,
	})
}

func HandlerOptions() []bot.Option {
	opts := []bot.Option{
		bot.WithMessageTextHandler(startCMD, bot.MatchTypeExact, StartHandler),
		bot.WithMessageTextHandler(helpCMD, bot.MatchTypeExact, HelpHandler),
		bot.WithMessageTextHandler(weatherCMD, bot.MatchTypeExact, WeatherHandler),
		bot.WithDefaultHandler(UnknownHandler),
	}

	return opts
}
