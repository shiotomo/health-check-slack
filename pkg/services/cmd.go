package services

import (
	"io/ioutil"
	"os"

	"github.com/shiotomo/health-check-slack/pkg/models"
	"github.com/shiotomo/health-check-slack/pkg/utils"
	"github.com/shirou/gopsutil/host"
)

// checkコマンド
func CheckServer() models.Server {
	var server models.Server
	info, _ := host.Info()

	server.Host = info.Hostname
	server.Uptime, _ = host.Uptime()
	server.Os = info.OS
	server.Platform = info.Platform
	server.Ip = utils.GetHostIpAddr()

	return server
}

// callコマンド
func CallServer() models.Server {
	var server models.Server
	info, _ := host.Info()

	server.Host = info.Hostname
	server.Ip = utils.GetHostIpAddr()

	return server
}

// helpコマンド
func Help() models.Help {
	var help models.Help
	f, err := os.Open("i18n/help/help_ja.txt")
	if err != nil {
		help.Message = "Error help command"
		return help
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)

	if err != nil {
		help.Message = "Error help command"
		return help
	}

	help.Message = string(buf)

	return help
}
