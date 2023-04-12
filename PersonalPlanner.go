package main

import (
	"flag"
	"fmt"
	"log"
)

func mustToken() string {
	token := flag.String(
		"token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if token == nil || *token == "" {
		log.Fatalln("token is not specified")
	}

	return *token

}

func main() {
	fmt.Println("Hello world", mustToken())
}
