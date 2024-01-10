package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebsocketCon(c echo.Context) error {

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	defer conn.Close()

	if err != nil {
		log.Println(err)
		return nil
	}
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}

	return nil
}
