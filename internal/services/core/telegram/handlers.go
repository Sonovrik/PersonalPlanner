package telegram

import (
	"PersonalPlanner/internal/services/weather"
	"PersonalPlanner/internal/services/weather/yandex"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var weatherManager weather.WApi //nolint

func mustInitWeather(token string) {
	wToken := token

	weatherManager = weather.WApi(yandex.New(wToken))
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
	w, err := weatherManager.WeatherInfo(ctx, 55.755864, 37.617698)
	if err != nil {
		ErrorHandler(ctx, b, update, err)

		return
	}

	currentW := w.Current()
	nextW := w.Next()

	if nextW == "" || currentW == "" {
		err = fmt.Errorf("can't get current or next weather")
		ErrorHandler(ctx, b, update, err)

		return
	}

	// TODO handle error
	weatherMsg := currentW + "\n\n" + nextW
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   weatherMsg,
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
