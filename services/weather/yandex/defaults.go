package yandex

// conditions - коды погодного описания и значения на русском языке, получаемые из Fact.Condition
var conditions = map[string]string{
	"partly-cloudy":          "малооблачно",
	"overcast":               "пасмурно",
	"drizzle":                "морось",
	"cloudy":                 "облачно с прояснениями",
	"clear":                  "ясно",
	"light-rain":             "небольшой дождь",
	"rain":                   "дождь",
	"moderate-rain":          "умеренно сильный дождь",
	"heavy-rain":             "сильный дождь",
	"continuous-heavy-rain":  "длительный сильный дождь",
	"showers":                "ливень",
	"wet-snow":               "дождь со снегом",
	"light-snow":             "небольшой снег",
	"snow":                   "снег",
	"snow-showers":           "снегопад",
	"hail":                   "град",
	"thunderstorm":           "гроза",
	"thunderstorm-with-rain": "дождь с грозой",
	"thunderstorm-with-hail": "гроза с градом",
}

// conditions - коды погодного описания и значения на русском языке, получаемые из Fact.Condition
var partNames = map[string]string{
	"night":   "ночь",
	"morning": "утро",
	"day":     "день",
	"evening": "вечер",
}
