package main

import (
	"bytes"
	"flag"
	"github.com/gorilla/websocket"
	"net"
	"net/url"
	"os"
	"websocket-tester/pkg/log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var testData = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a}

var config Config

func init() {
	flag.StringVar(&config.WebsocketURL, "s", "wss://stream.nndroid.com/ws", "websocket url")
	flag.StringVar(&config.IPAddress, "a", "1.1.1.1", "ip address")
	flag.Parse()
}

func main() {
	u, err := url.Parse(config.WebsocketURL)
	if err != nil {
		log.Logger().Fatalf("url parse error: %v", err)
	}
	dialer := websocket.Dialer{
		NetDial: func(network, addr string) (net.Conn, error) {
			_, port, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, err
			}
			return net.Dial(network, net.JoinHostPort(config.IPAddress, port))
		},
	}
	c, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		log.Logger().Fatalf("websocket dial error: %v", err)
	}
	log.Logger().Infof("connected to %s", u.String())
	defer c.Close()
	err = c.WriteMessage(websocket.TextMessage, testData)
	if err != nil {
		log.Logger().Fatalf("write data error: %v", err)
	}
	log.Logger().Infof("sent: %v", testData)
	_, msg, err := c.ReadMessage()
	if err != nil {
		log.Logger().Fatalf("read data error: %v", err)
	}
	log.Logger().Infof("received: %v", msg)
	if bytes.Compare(testData, msg) == 0 {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
