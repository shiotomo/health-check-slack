package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shiotomo/health-check-slack/internal"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/call", internal.CallRouter)
	e.GET("/check", internal.CheckRouter)
	e.GET("/help", internal.HelpRouter)

	// サーバー起動
	e.Logger.Fatal(e.Start(":1323"))
}
