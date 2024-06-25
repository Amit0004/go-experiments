package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

/* A simple web server */
// func main() {
// 	fmt.Println("Hello, World!")

// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run()
// }

var upgrader2 = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ws(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade", err)
		return
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Print("read", err)
			break
		}
		log.Printf("recv: %s", message)

		//write websocket data
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("write", err)
			break
		}
	}
}

func server() {
	fmt.Println("Websocket Server!")

	bindAddress := "localhost:8448"
	r := gin.Default()
	r.GET("/ws", ws)
	r.Run(bindAddress)
}
