package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New(
		"YOUR_TOKEN_HERE",
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			// do something
			rtm.SendMessage(rtm.NewOutgoingMessage("Hello", ev.Channel))
		default:
			// ignore other events
			// fmt.Printf("ignored: %v¥n", msg.Data)
		}
	}
}
