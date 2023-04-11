package main

import (
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
)

func mustToken() string {
	token := flag.String(
		"token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if token == nil || *token == "" {
		logrus.Fatalln("token is not specified")
	}

	return *token
}

func main() {
	fmt.Println("Hello world", mustToken())
}
