package gobot

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type slackbot struct {
	cli        *slack.Client
	socketMode *socketmode.Client
}

func NewSlackBot(slackBotToken string, slackAppToken string) *slackbot {
	cli := slack.New(
		slackBotToken,
		slack.OptionAppLevelToken(slackAppToken),
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "api: ", log.Lshortfile|log.LstdFlags)),
	)
	socketMode := socketmode.New(
		cli,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "sm: ", log.Lshortfile|log.LstdFlags)),
	)

	at, err := cli.AuthTest()
	if err != nil {
		return nil
	}
	fmt.Println("authTest User: ", at.User, "authTest UserID", at.UserID)
	return &slackbot{
		cli,
		socketMode,
	}
}

func (sb *slackbot) Listen() {
	for ev := range sb.socketMode.Events {
		payload, _ := ev.Data.(slackevents.EventsAPIEvent)
		switch payload.Type {
		case slackevents.CallbackEvent:
			event := payload.InnerEvent.Data.(*slackevents.MessageEvent)
			sb.cli.PostMessage(
				event.Channel,
				slack.MsgOptionText(
					fmt.Sprintf(":wave: こんにちは <@%v> さん！", event.User),
					false,
				),
			)

		default:
			log.Println("default")
		}
	}
}

func (sb *slackbot) RunSocketMode() {
	sb.socketMode.Run()
}
