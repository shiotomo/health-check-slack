package services

import (
	"fmt"
	"strings"

	"github.com/shiotomo/health-check-slack/pkg/models"
	"github.com/shirou/gopsutil/host"
	"github.com/slack-go/slack"
)

const (
	CMD_NUM      = 0
	MIN_CMD_ARGC = 1
)

func checkServer() string {
	var server models.Server
	info, _ := host.Info()

	server.Host = info.Hostname
	server.Uptime, _ = host.Uptime()
	server.Os = info.OS
	server.Platform = info.Platform

	return models.ServerToString(server)
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
	fmt.Println(cmd)
	switch cmd[CMD_NUM] {
	case "check":
		if len(cmd) > MIN_CMD_ARGC {
			info, _ := host.Info()
			if judgeHost(cmd[1], info.Hostname) {
				rtm.SendMessage(rtm.NewOutgoingMessage(checkServer(), ev.Channel))
			}
		}
	}
}
