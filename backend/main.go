package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hellow World")
	})
	e.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "somethings")
	})

	e.GET("/ws", WebsocketCon)
	e.Logger.Fatal(e.Start(":4000"))
}
