package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"websocket-tester/pkg/log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var config Config

func init() {
	flag.StringVar(&config.Listen, "l", ":8000", "listening address")
	flag.StringVar(&config.Path, "p", "/ws", "path")
	flag.Parse()
}

func main() {
	r := gin.Default()
	r.GET(config.Path, func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Logger().Errorf("conn upgrade error, %v", err)
			return
		}
		var msg []byte
		for {
			_, msg, err = conn.ReadMessage()
			if err != nil {
				log.Logger().Errorf("read message error, %v", err)
				return
			}
			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Logger().Errorf("write message error, %v", err)
				return
			}
		}
	})
	log.Logger().Panic(r.Run(config.Listen))
}
