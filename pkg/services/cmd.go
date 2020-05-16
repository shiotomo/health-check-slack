package services

import (
	"io/ioutil"
	"os"

	"github.com/shiotomo/health-check-slack/pkg/models"
	"github.com/shiotomo/health-check-slack/pkg/utils"
	"github.com/shirou/gopsutil/host"
)

// checkコマンド
func CheckServer() string {
	var server models.Server
	info, _ := host.Info()

	server.Host = info.Hostname
	server.Uptime, _ = host.Uptime()
	server.Os = info.OS
	server.Platform = info.Platform
	server.Ip = utils.GetHostIpAddr()

	return models.CheckToString(server)
}

// callコマンド
func CallServer() string {
	var server models.Server
	info, _ := host.Info()

	server.Host = info.Hostname
	server.Ip = utils.GetHostIpAddr()

	return models.CallToString(server)
}

// helpコマンド
func Help() string {
	f, err := os.Open("i18n/help/help_ja.txt")
	if err != nil {
		return "Error help command."
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)

	if err != nil {
		return "Error help command."
	}

	return string(buf)
}
