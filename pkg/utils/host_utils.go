package utils

import "strings"

func JudgeHost(hostA string, hostB string) bool {
	if hostA == hostB {
		return true
	}
	if strings.Replace(hostA, ".local", "", 1) == strings.Replace(hostB, ".local", "", 1) {
		return true
	}
	return false
}
