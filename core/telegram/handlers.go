package telegram

import (
	"PersonalPlanner/services/weather/yandex"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"time"
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
	yandexWeatherApiKey := "35e9673a-3d3b-4d4f-ba7f-956e1b1f6a10"
	w, err := yandex.GetWeather(ctx, yandexWeatherApiKey, 55.755864, 37.617698)
	if err != nil {
		ErrorHandler(ctx, b, update, err)
		return
	}

	msg := fmt.Sprintf("Погода на %s\n"+
		"Погода - %s\n"+
		"Температура - %d (°C)\n"+
		"Ощущается как - %d (°C)\n",
		time.Unix(w.Now, 0), w.Fact.GetCondition(), w.Fact.Temp, w.Fact.FeelsLike)

	// TODO handle error
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   msg,
	})

	next := ""
	for _, v := range w.Forecast.Parts {
		next += fmt.Sprintf("Прогноз на %s - %s\n"+
			"Средняя температура - %d (°C)\n"+
			"Ощущается как - %d (°C)\n"+
			"Скорость ветра - %.1f\n"+
			"Количество осадков - %d мм\n"+
			"Вероятность выпадения осадков - %d\n"+
			"Период осадков - %d мин\n\n",
			v.GetPartName(), v.GetCondition(), v.TempAvg,
			v.FeelsLike, v.WindSpeed, v.PrecMm, v.PrecProb, v.PrecPeriod)
	}

	// TODO handle error
	_, _ = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   next,
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
