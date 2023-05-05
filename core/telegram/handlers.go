package telegram

import (
	"PersonalPlanner/services/weather"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"PersonalPlanner/services/weather/yandex"
)

const (
	startCMD   = "/start"
	helpCMD    = "/help"
	weatherCMD = "/weather"
)

func HandlerOptions() []bot.Option {
	opts := []bot.Option{
		bot.WithMessageTextHandler(startCMD, bot.MatchTypeExact, StartHandler),
		bot.WithMessageTextHandler(helpCMD, bot.MatchTypeExact, HelpHandler),
		bot.WithMessageTextHandler(weatherCMD, bot.MatchTypeExact, WeatherHandler),
		bot.WithDefaultHandler(UnknownHandler),
	}

	return opts
}

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
	wToken := weather.MustToken()

	wApi := weather.WApi(yandex.New(wToken))
	w, err := wApi.GetWeather(ctx, 55.755864, 37.617698)
	if err != nil {
		ErrorHandler(ctx, b, update, err)

		return
	}

	currentW := w.Current()
	if len(currentW) == 0 {
		ErrorHandler(ctx, b, update, err)

		return
	}

	// TODO handle error
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   currentW,
	})

	nextW := w.Next()
	if len(nextW) == 0 {
		ErrorHandler(ctx, b, update, err)

		return
	}

	// TODO handle error
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   nextW,
	})
}

func UnknownHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// TODO handle error
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   msgUnknownCommand,
	})
}

func ErrorHandler(ctx context.Context, b *bot.Bot, update *models.Update, err error) {
	// TODO handle error
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   msgError + err.Error(),
	})
}
