package main

import (
	"bytes"
	"context"
	"flag"
	"github.com/gorilla/websocket"
	"net"
	"net/url"
	"os"
	"time"
	"websocket-tester/pkg/log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var testData = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a}

var config Config

func init() {
	flag.StringVar(&config.WebsocketURL, "s", "wss://stream.nndroid.com/ws", "url")
	flag.StringVar(&config.IPAddress, "a", "1.1.1.1", "ip address")
	flag.StringVar(&config.Timeout, "t", "10s", "timeout")
	flag.Parse()
}

func main() {
	u, err := url.Parse(config.WebsocketURL)
	if err != nil {
		log.Logger().Fatalf("url parse error: %v", err)
	}
	duration, err := time.ParseDuration(config.Timeout)
	if err != nil {
		log.Logger().Fatalf("duration parse error: %v", err)
	}
	netDialer := &net.Dialer{}
	wsDialer := websocket.Dialer{
		HandshakeTimeout: duration,
		NetDialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			_, port, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, err
			}
			return netDialer.DialContext(ctx, network, net.JoinHostPort(config.IPAddress, port))
		},
	}
	var start time.Time
	var end time.Time
	var cost time.Duration
	start = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	c, _, err := wsDialer.DialContext(ctx, u.String(), nil)
	if err != nil {
		log.Logger().Fatalf("websocket dial error: %v", err)
	}
	end = time.Now()
	cost = end.Sub(start)
	log.Logger().Infof("connected to %s, cost: %s", u.String(), cost)
	defer c.Close()
	start = time.Now()
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
	end = time.Now()
	cost = end.Sub(start)
	log.Logger().Infof("read & write cost: %s", cost)
	if bytes.Compare(testData, msg) == 0 {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
