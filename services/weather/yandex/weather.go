package yandex

import (
	"PersonalPlanner/services/weather"
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	apiURLTemplate         = "https://api.weather.yandex.ru/v2/informers?lat=%f&lon=%f"
	weatherIconURLTemplate = "https://yastatic.net/weather/i/icons/funky/dark/%s.svg"
)

// GetCondition получение описания погоды на русском языке
func (f *Fact) GetCondition() string {
	return conditions[f.Condition]
}

// GetMoon получение фазы луны
func (f *Forecast) GetMoon() string {
	switch {
	case f.MoonCode == 0:
		return "полнолуние"
	case (f.MoonCode >= 1 && f.MoonCode <= 3) || (f.MoonCode >= 5 && f.MoonCode <= 7):
		return "убывающая Луна"
	case f.MoonCode == 4:
		return "последняя четверть"
	case f.MoonCode == 8:
		return "новолуние"
	case f.MoonCode >= 9 && f.MoonCode <= 11:
		return "растущая Луна"
	case f.MoonCode == 12:
		return "первая четверть"
	default:
		return ""
	}
}

// GetPartName получение названия времени суток на русском языке
func (p *Part) GetPartName() string {
	return partNames[p.PartName]
}

// GetCondition получение описания погоды на русском языке
func (p *Part) GetCondition() string {
	return conditions[p.Condition]
}

func (w *Weather) Current() string {
	return fmt.Sprintf("Погода на %s\n"+
		"Погода - %s\n"+
		"Температура - %d (°C)\n"+
		"Ощущается как - %d (°C)\n",
		time.Unix(w.Now, 0), w.Fact.GetCondition(), w.Fact.Temp, w.Fact.FeelsLike)
}

func (w *Weather) Next() string {
	next := ""

	for i := range w.Forecast.Parts {
		v := &w.Forecast.Parts[i]
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

	return next
}

// GetWeather Получение погоды из Яндекс API
func (w *WApi) GetWeather(ctx context.Context, lat, lon float32) (weather.Weather, error) {
	url := fmt.Sprintf(apiURLTemplate, lat, lon)

	wr := &Weather{}

	err := getRequest(
		ctx,
		url,
		wr,
		func(req *http.Request) {
			req.Header.Add("X-Yandex-API-Key", w.token)
		},
	)
	if err != nil {
		return nil, err
	}

	return weather.Weather(wr), nil
}
