package models

import (
	"strconv"
	"strings"
)

type Server struct {
	Host     string   `json:"host"`
	Os       string   `json:"os"`
	Platform string   `json:"platform"`
	Uptime   uint64   `json:"uptime"`
	Ip       []string `json:ip`
}

func CheckToString(server Server) string {
	host := "host: " + server.Host + "\n"
	os := "os: " + server.Os + "\n"
	platform := "platform: " + server.Platform + "\n"
	uptime := "uptime: " + strconv.FormatUint(server.Uptime, 10) + "\n"
	ip := "ip: " + strings.Join(server.Ip, ", ") + "\n"

	return host + os + platform + uptime + ip
}

func CallToString(server Server) string {
	host := "host: " + server.Host + "\n"
	ip := "ip: " + strings.Join(server.Ip, ", ") + "\n"

	return host + ip
}
