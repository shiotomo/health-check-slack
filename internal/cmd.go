package internal

import (
	"fmt"
	"strings"

	"github.com/shiotomo/health-check-slack/pkg/config"
	"github.com/shiotomo/health-check-slack/pkg/models"
	"github.com/shiotomo/health-check-slack/pkg/services"
	"github.com/shiotomo/health-check-slack/pkg/utils"
	"github.com/shirou/gopsutil/host"
	"github.com/slack-go/slack"
)

// コマンドの実行を行う関数
func RunCmd(ev *slack.MessageEvent, rtm *slack.RTM, api *slack.Client) {
	text := ev.Text
	cmd := strings.Split(text, " ")
	switch cmd[config.CmdNum] {
	case "check":
		if len(cmd) > config.MinCmdArgc {
			info, _ := host.Info()
			if utils.JudgeHost(cmd[config.ServerName], info.Hostname) {
				rtm.SendMessage(rtm.NewOutgoingMessage(models.CheckToString(services.CheckServer()), ev.Channel))
			}
		}
	case "call":
		rtm.SendMessage(rtm.NewOutgoingMessage(models.CallToString(services.CallServer()), ev.Channel))
	default:
		fmt.Println("no cmd")
	}
}

// 管理コマンドの実行を行う関数
func RunMasterCmd(ev *slack.MessageEvent, rtm *slack.RTM, api *slack.Client) {
	text := ev.Text
	cmd := strings.Split(text, " ")
	switch cmd[config.CmdNum] {
	case "help":
		rtm.SendMessage(rtm.NewOutgoingMessage(services.Help().Message, ev.Channel))
	default:
		fmt.Println("no master cmd")
	}
}
