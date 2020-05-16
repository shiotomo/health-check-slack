package utils

import (
	"testing"

	"github.com/shiotomo/health-check-slack/pkg/config"
)

func TestGetBotRole(t *testing.T) {
	roleA := "master"
	roleB := "Master"
	roleC := "MASTER"
	roleD := "node"
	roleE := "Node"
	roleF := "NODE"
	roleG := "hoge"

	// masterのRoleか判定
	// masterでなければエラー
	if GetBotRole(roleA) != config.Master {
		t.Fatal("Error: This role is not master")
	}
	if GetBotRole(roleB) != config.Master {
		t.Fatal("Error: This role is not master")
	}
	if GetBotRole(roleC) != config.Master {
		t.Fatal("Error: This role is not master")
	}

	// nodeのRoleか判定
	// masterであればエラー
	if GetBotRole(roleD) == config.Master {
		t.Fatal("Error: This role is master")
	}
	if GetBotRole(roleE) == config.Master {
		t.Fatal("Error: This role is master")
	}
	if GetBotRole(roleF) == config.Master {
		t.Fatal("Error: This role is master")
	}
	if GetBotRole(roleG) == config.Master {
		t.Fatal("Error: This role is master")
	}

	t.Log("Success!!")
}
