package config

type Config struct {
	Title string `yaml:"title"`
	Core  Core   `yaml:"core"`
}

type Core struct {
	CoreAPIToken    string `yaml:"coreAPITelegramToken"`
	WeatherAPIToken string `yaml:"weatherAPIYandexToken"`
}
