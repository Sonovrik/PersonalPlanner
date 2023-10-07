package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func Init(configPath *string) (*Config, error) {
	var err error

	var yamlCfg []byte

	if configPath != nil {
		yamlCfg, err = os.ReadFile(*configPath)
	}

	if configPath == nil || err != nil {
		if err != nil {
			log.Printf("Cannot read config %s, using default config\n", *configPath)
		}

		yamlCfg = []byte(defaultConfig())
	}

	cfg := new(Config)
	if err = yaml.Unmarshal(yamlCfg, cfg); err != nil {
		return nil, fmt.Errorf("can't parse yaml config file")
	}

	return cfg, nil
}
