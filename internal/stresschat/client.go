package stresschat

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/anonnice1999/questx-script/pkg/api"
	"github.com/gorilla/websocket"
)

type Client struct {
	userID    string
	wsClient  *websocket.Conn
	channelID int64
}

func NewClient(serverAddr, userID, token, apiEndpoint, communityHandle string, channelID int64) (*Client, error) {
	apiGenerator := api.NewGenerator()
	resp, err := apiGenerator.New(apiEndpoint, "/follow").
		Body(api.JSON{
			"community_handle": communityHandle,
		}).
		POST(context.Background(), api.OAuth2("Bearer", token))
	if err != nil {
		return nil, err
	}
	body, ok := resp.Body.(api.JSON)
	if !ok {
		return nil, fmt.Errorf("invalid body: %v", resp.Body)
	}

	code, err := body.GetInt("code")
	if err != nil {
		return nil, err
	}

	if code != 0 && code != 100000 {
		return nil, fmt.Errorf("invalid response: %v", body)
	}

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

	return &Client{userID: userID, wsClient: wsClient, channelID: channelID}, nil
}

func (c *Client) Close() error {
	return c.wsClient.Close()
}

func (c *Client) Run(apiEndpoint, communityHandle, token string, channelID int64) {
	done := make(chan struct{})
	apiGenerator := api.NewGenerator()

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

			if t == websocket.TextMessage {
				decodedMsg := make([]byte, base64.StdEncoding.DecodedLen(len(msg)))
				if _, err := base64.StdEncoding.Decode(decodedMsg, msg); err != nil {
					log.Println("Cannot decode:", err)
					continue
				}
			}
		}
	}()

	ticker := time.NewTicker(50 * time.Millisecond)
	numberMsgTicker := time.NewTicker(100 * time.Second)
	numberMsg := 0
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			resp, err := apiGenerator.New(apiEndpoint, "/createMessage").
				Body(api.JSON{
					"community_handle": communityHandle,
					"content":          "something",
					"channel_id":       channelID,
				}).
				POST(context.Background(), api.OAuth2("Bearer", token))
			if err != nil {
				log.Printf("unable to create message: %v\n", err)
			}
			body, ok := resp.Body.(api.JSON)
			if !ok {
				log.Printf("invalid body: %v\n", err)
			}

			code, err := body.GetInt("code")
			if err != nil {
				log.Printf("code not exists in body: %v\n", err)
			}

			if code != 0 && code != 100006 {
				log.Printf("invalid response: %v\n", err)
			}
			numberMsg++

		case <-numberMsgTicker.C:
			fmt.Println(numberMsg)
			numberMsg = 0

		case <-interrupt:
			log.Println("interrupt")
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

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
