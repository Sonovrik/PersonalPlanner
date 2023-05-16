package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Title string `yaml:"title"`
	Core  Core   `yaml:"core"`
}

type Core struct {
	CoreAPIToken    string `yaml:"coreAPITelegramToken"`
	WeatherAPIToken string `yaml:"weatherAPIYandexToken"`
}

func Init(confPath string) (c *Config, err error) {
	yamlFile, err := os.ReadFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("using default config cause err reading config-file")
	}

	c = new(Config)
	if err := yaml.Unmarshal(yamlFile, c); err != nil {
		return nil, fmt.Errorf("can't parse yaml file %s", confPath)
	}

	return c, nil
}
