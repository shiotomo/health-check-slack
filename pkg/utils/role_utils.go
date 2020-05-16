package utils

import (
	"strings"

	"github.com/shiotomo/health-check-slack/pkg/config"
)

// Roleを返却する
func GetBotRole(role string) string {
	if strings.ToLower(role) == config.Master {
		return config.Master
	}
	return config.Node
}
