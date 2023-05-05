package weather

import (
	"context"
	"flag"
	"log"
)

const flagTokenName = "weatherToken"

type Weather interface {
	Current() string
	Next() string
}

type WApi interface {
	GetWeather(ctx context.Context, lat, lon float32) (Weather, error)
}

func MustToken() string {
	token := flag.String(
		flagTokenName,
		"",
		"token for access to weather api",
	)

	flag.Parse()

	if token == nil || *token == "" {
		log.Fatalln("token is not specified")
	}

	return *token
}
