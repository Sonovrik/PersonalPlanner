package yandex

import (
	"PersonalPlanner/utils"
	"context"
	"fmt"
	"net/http"
)

const apiURLTemplate = "https://api.weather.yandex.ru/v2/informers?lat=%f&lon=%f"
const weatherIconURLTemplate = "https://yastatic.net/weather/i/icons/funky/dark/%s.svg"

// GetCondition получение описания погоды на русском языке
func (f Fact) GetCondition() string {
	return conditions[f.Condition]
}

// GetMoon получение фазы луны
func (f Forecast) GetMoon() string {
	if f.MoonCode == 0 {
		return "полнолуние"
	}
	if f.MoonCode >= 1 && f.MoonCode <= 3 {
		return "убывающая Луна"
	}
	if f.MoonCode == 4 {
		return "последняя четверть"
	}
	if f.MoonCode >= 5 && f.MoonCode <= 7 {
		return "убывающая Луна"
	}
	if f.MoonCode == 8 {
		return "новолуние"
	}
	if f.MoonCode >= 9 && f.MoonCode <= 11 {
		return "растущая Луна"
	}
	if f.MoonCode == 12 {
		return "первая четверть"
	}
	if f.MoonCode >= 9 && f.MoonCode <= 11 {
		return "растущая Луна"
	}
	return ""
}

// GetPartName получение названия времени суток на русском языке
func (p Part) GetPartName() string {
	return partNames[p.PartName]
}

// GetCondition получение описания погоды на русском языке
func (p Part) GetCondition() string {
	return conditions[p.Condition]
}

// GetWeather Получение погоды из Яндекс API
func GetWeather(ctx context.Context, yandexWeatherApiKey string, lat float32, lon float32) (*Weather, error) {
	url := fmt.Sprintf(apiURLTemplate, lat, lon)
	w := &Weather{}
	err := utils.GetRequest(
		ctx,
		url,
		w,
		func(req *http.Request) {
			req.Header.Add("X-Yandex-API-Key", yandexWeatherApiKey)
		},
	)

	if err != nil {
		return nil, err
	}

	return w, nil
}
