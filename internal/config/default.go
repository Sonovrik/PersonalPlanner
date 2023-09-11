package config

func defaultConfig() string {
	config := `
		title: "Personal Planner config"
		core:
  			coreAPITelegramToken: "Your tg token"
  			weatherAPIYandexToken: "Your yandex weather api token"
		`

	return config
}
