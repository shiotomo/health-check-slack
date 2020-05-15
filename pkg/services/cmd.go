package services

import (
	"fmt"
	"strings"

	"github.com/shiotomo/health-check-slack/pkg/models"
	"github.com/shiotomo/health-check-slack/pkg/utils"
	"github.com/shirou/gopsutil/host"
	"github.com/slack-go/slack"
)

const (
	CMD_NUM      = 0
	MIN_CMD_ARGC = 1
	SERVER_NAME  = 1
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

func judgeHost(hostA string, hostB string) bool {
	if hostA == hostB {
		return true
	}
	if strings.Replace(hostA, ".local", "", 1) == strings.Replace(hostB, ".local", "", 1) {
		return true
	}
	return false
}

func RunCmd(ev *slack.MessageEvent, rtm *slack.RTM, api *slack.Client) {
	text := ev.Text
	cmd := strings.Split(text, " ")
	switch cmd[CMD_NUM] {
	case "check":
		if len(cmd) > MIN_CMD_ARGC {
			info, _ := host.Info()
			if judgeHost(cmd[SERVER_NAME], info.Hostname) {
				rtm.SendMessage(rtm.NewOutgoingMessage(checkServer(), ev.Channel))
			}
		}
	case "call":
		rtm.SendMessage(rtm.NewOutgoingMessage(callServer(), ev.Channel))
	default:
		fmt.Println("no cmd")
	}
}
