package services

import (
	"fmt"
	"strings"

	"github.com/slack-go/slack"
)

const (
	CMD_NUMBER = 0
)

func RunCmd(ev *slack.MessageEvent, rtm *slack.RTM, api *slack.Client) {
	text := ev.Text
	cmd := strings.Split(text, " ")
	fmt.Println(cmd)
	if text == "aaaa" {
		rtm.SendMessage(rtm.NewOutgoingMessage(text, ev.Channel))
	}
}
