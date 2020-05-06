package models

import "strconv"

type Server struct {
	Host     string `json:"host"`
	Os       string `json:"os"`
	Platform string `json:"platform"`
	Uptime   uint64 `json:"uptime"`
}

func ServerToString(server Server) string {
	host := "host: " + server.Host + "\n"
	os := "os: " + server.Os + "\n"
	platform := "platform: " + server.Platform + "\n"
	uptime := "uptime: " + strconv.FormatUint(server.Uptime, 10) + "\n"

	return host + os + platform + uptime
}
