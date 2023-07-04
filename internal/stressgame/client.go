package stressgame

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	wsClient  *websocket.Conn
	mapWidth  int
	mapHeight int
}

func NewClient(serverAddr, token string, mapWidth, mapHeight int) (*Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(serverAddr)
	if err != nil {
		return nil, err
	}

	if serverURL.Scheme == "ws" {
		serverURL.Scheme = "http"
	} else {
		serverURL.Scheme = "https"
	}

	jar.SetCookies(serverURL, []*http.Cookie{
		{
			Name:  "access_token",
			Value: token,
		},
	})

	dialer := websocket.Dialer{
		HandshakeTimeout: 45 * time.Second,
		Jar:              jar,
	}

	wsClient, _, err := dialer.Dial(serverAddr, nil)
	if err != nil {
		return nil, err
	}

	return &Client{wsClient: wsClient, mapWidth: mapWidth, mapHeight: mapHeight}, nil
}

func (c *Client) Close() error {
	return c.wsClient.Close()
}

func (c *Client) Run() {
	done := make(chan struct{})

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		defer close(done)
		for {
			t, msg, err := c.wsClient.ReadMessage()
			if err != nil {
				log.Println(c.wsClient.LocalAddr(), "cannot read anymore:", err)
				return
			}

			if t == websocket.CloseMessage {
				log.Println("Close reason message:", msg)
			}
		}
	}()

	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			move := map[string]any{
				"type": "move",
				"value": map[string]any{
					"x":         rand.Intn(c.mapWidth),
					"y":         rand.Intn(c.mapHeight),
					"direction": "down",
				},
			}

			b, err := json.Marshal(move)
			if err != nil {
				log.Println("Cannot marshal move action:", err)
				return
			}

			err = c.wsClient.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				log.Println("Cannot write message anymore:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.wsClient.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Cannot write a close message:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
