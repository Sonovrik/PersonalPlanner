package telegram

import (
	"PersonalPlanner/services/weather"
	"PersonalPlanner/services/weather/yandex"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var weatherAPI weather.WApi //nolint

func mustInitWeather(token string) {
	wToken := token

	weatherAPI = weather.WApi(yandex.New(wToken))
}

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
	w, err := weatherAPI.GetWeather(ctx, 55.755864, 37.617698)
	if err != nil {
		ErrorHandler(ctx, b, update, err)

		return
	}

	currentW := w.Current()
	if currentW == "" {
		ErrorHandler(ctx, b, update, err)

		return
	}

	// TODO handle error
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   currentW,
	})

	nextW := w.Next()
	if nextW == "" {
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
