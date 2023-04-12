package telegram

import (
	"log"
	"strings"
)

const (
	StartCmd   = "/start"
	HelpCmd    = "/help"
	WeatherCmd = "/weather"
)

func (p *TgProcessor) doCmd(text string, chatID int, username string) error {
	messageText := strings.TrimSpace(text)

	log.Printf("New cmd from user %s to execute: %s", username, messageText)

	switch messageText {
	case StartCmd:

	case HelpCmd:

	case WeatherCmd:

	default:

	}

	return nil
}

func (p *TgProcessor) sendHelp() error {
	return nil
}

func (p *TgProcessor) sendStart() error {
	return nil
}

func (p *TgProcessor) sendWeather() error {
	return nil
}

func (p *TgProcessor) sendUnknownCmd() error {
	return nil
}
