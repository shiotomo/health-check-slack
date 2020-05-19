package internal

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/shiotomo/health-check-slack/pkg/services"
)

func CallRouter(c echo.Context) error {
	server := services.CallServer()
	return c.JSON(http.StatusOK, server)
}

func CheckRouter(c echo.Context) error {
	server := services.CheckServer()
	return c.JSON(http.StatusOK, server)
}

func HelpRouter(c echo.Context) error {
	help := services.Help()
	return c.JSON(http.StatusOK, help)
}
