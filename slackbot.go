package gobot

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type slackbot struct {
	cli        *slack.Client
	socketMode *socketmode.Client
	// userID     string
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

	// at, err := cli.AuthTest()
	// if err != nil {
	// 	return nil
	// }
	// userID := at.UserID
	return &slackbot{
		cli,
		socketMode,
		// userID,
	}
}

func (sb *slackbot) CreateObject(objText string) slack.MsgOption {
	text := slack.NewTextBlockObject(slack.MarkdownType, objText, false, false)
	textSection := slack.NewSectionBlock(text, nil, nil)

	btnText := slack.NewTextBlockObject(slack.PlainTextType, "YES", false, false)
	btn := slack.NewButtonBlockElement("", "test", btnText)
	btn.WithStyle(slack.StylePrimary)
	actionBlock := slack.NewActionBlock("confirm-development", btn)

	blocks := slack.MsgOptionBlocks(textSection, actionBlock)
	return blocks
}

func (sb *slackbot) Listen() {
	for ev := range sb.socketMode.Events {
		switch ev.Type {
		case socketmode.EventTypeConnecting:
			fmt.Println("Connecting to Slack with Socket Mode...")
		case socketmode.EventTypeConnectionError:
			fmt.Println("Connection failed. Retrying later...")
		case socketmode.EventTypeConnected:
			fmt.Println("Connected to Slack with Socket Mode.")
		case socketmode.EventTypeEventsAPI:
			sb.socketMode.Ack(*ev.Request)
			payload, _ := ev.Data.(slackevents.EventsAPIEvent)
			switch payload.Type {
			case slackevents.CallbackEvent:
				event := payload.InnerEvent.Data.(*slackevents.MessageEvent)
				if strings.Contains(event.Text, "Hello") {
					msgOption := sb.CreateObject("Hello, World")
					sb.cli.PostMessage(event.Channel, msgOption)
				}
				if !strings.Contains(event.Text, "こんにちは") {
					return
				}
				sb.cli.PostMessage(
					event.Channel,
					slack.MsgOptionText(
						fmt.Sprintf(":wave: <@%v> さん！", event.User),
						false,
					),
				)

			default:
				sb.socketMode.Debugf("Skipped: %v", ev.Type)
			}
		default:
			sb.socketMode.Debugf("TypeError")
		}
	}
}

func (sb *slackbot) RunSocketMode() {
	sb.socketMode.Run()
}
