package services

import (
	"fmt"
	"strings"

	"github.com/shiotomo/health-check-slack/pkg/config"
	"github.com/shiotomo/health-check-slack/pkg/models"
	"github.com/shiotomo/health-check-slack/pkg/utils"
	"github.com/shirou/gopsutil/host"
	"github.com/slack-go/slack"
)

func checkServer() string {
	var server models.Server
	info, _ := host.Info()

	server.Host = info.Hostname
	server.Uptime, _ = host.Uptime()
	server.Os = info.OS
	server.Platform = info.Platform
	server.Ip = utils.GetHostIpAddr()

	return models.CheckToString(server)
}

func callServer() string {
	var server models.Server
	info, _ := host.Info()

	server.Host = info.Hostname
	server.Ip = utils.GetHostIpAddr()

	return models.CallToString(server)
}

func RunCmd(ev *slack.MessageEvent, rtm *slack.RTM, api *slack.Client) {
	text := ev.Text
	cmd := strings.Split(text, " ")
	switch cmd[config.CMD_NUM] {
	case "check":
		if len(cmd) > config.MIN_CMD_ARGC {
			info, _ := host.Info()
			if utils.JudgeHost(cmd[config.SERVER_NAME], info.Hostname) {
				rtm.SendMessage(rtm.NewOutgoingMessage(checkServer(), ev.Channel))
			}
		}
	case "call":
		rtm.SendMessage(rtm.NewOutgoingMessage(callServer(), ev.Channel))
	default:
		fmt.Println("no cmd")
	}
}
