package main

import (
	"os"

	"github.com/gba-3/gobot"
)

func main() {
	slackBotToken := os.Getenv("SLACK_BOT_TOKEN")
	slackAppToken := os.Getenv("SLACK_APP_TOKEN")
	bot := gobot.NewSlackBot(slackBotToken, slackAppToken)
	go bot.Listen()
	bot.RunSocketMode()
}
