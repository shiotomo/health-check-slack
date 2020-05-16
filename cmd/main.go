package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shiotomo/health-check-slack/internal"
	"github.com/shiotomo/health-check-slack/pkg/config"
	"github.com/shiotomo/health-check-slack/pkg/utils"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))

	if err != nil {
		// can not read .env
		fmt.Println("can not read .env")
		panic(err)
	}

	token := os.Getenv("SLACK_BOT_TOKEN")
	role := utils.GetBotRole(os.Getenv(("ROLE")))

	api := slack.New(
		token,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Println("Event Received")
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if role == config.Master {
				internal.RunMasterCmd(ev, rtm, api)
			}
			internal.RunCmd(ev, rtm, api)
		}
	}
}
