package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	fmt.Println("Hello World")

	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))

	if err != nil {
		// can not read .env
		fmt.Println("can not read .env")
		panic(err)
	}

	token := os.Getenv("SLACK_BOT_TOKEN")

	api := slack.New(
		token,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			rtm.SendMessage(rtm.NewOutgoingMessage("hello world", ev.Channel))
		}
	}
}
