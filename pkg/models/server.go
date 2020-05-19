package models

import (
	"strconv"
	"strings"
)

// Server Model
type Server struct {
	Host     string   `json:"host"`
	Os       string   `json:"os"`
	Platform string   `json:"platform"`
	Uptime   uint64   `json:"uptime"`
	Ip       []string `json:ip`
}

// Checkコマンド用toString
func CheckToString(server Server) string {
	host := "host: " + server.Host + "\n"
	os := "os: " + server.Os + "\n"
	platform := "platform: " + server.Platform + "\n"
	uptime := "uptime: " + strconv.FormatUint(server.Uptime, 10) + "\n"
	ip := "ip: " + strings.Join(server.Ip, ", ") + "\n"

	return host + os + platform + uptime + ip
}

// Callコマンド用toString
func CallToString(server Server) string {
	host := "host: " + server.Host + "\n"
	ip := "ip: " + strings.Join(server.Ip, ", ") + "\n"

	return host + ip
}
