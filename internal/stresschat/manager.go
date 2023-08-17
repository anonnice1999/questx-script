package stresschat

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/anonnice1999/questx-script/pkg/token"
)

type Manager struct {
	serverAddress   string
	tokenEngine     token.Engine
	nClients        int
	apiEndpoint     string
	communityHandle string
	channelID       int64
}

func NewManager(serverAddress, tokenSecret, apiEndpoint, communityHandle string, channelID int64, nClients int) *Manager {
	return &Manager{
		serverAddress:   serverAddress,
		tokenEngine:     token.NewEngine(tokenSecret),
		nClients:        nClients,
		apiEndpoint:     apiEndpoint,
		communityHandle: communityHandle,
		channelID:       channelID,
	}
}

func (m *Manager) Run() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	wait := sync.WaitGroup{}

	for i := 0; i < m.nClients; i++ {
		select {
		case <-interrupt:
			i = m.nClients

		default:
			userID := fmt.Sprintf("testuser%d", i)
			token, err := m.tokenEngine.Generate(2000*time.Hour, map[string]any{"id": userID})
			if err != nil {
				return err
			}

			client, err := NewClient(m.serverAddress, userID, token, m.apiEndpoint,
				m.communityHandle, m.channelID)
			if err != nil {
				log.Println("Failed to connect for user", userID, ":", err)
				continue
			}

			wait.Add(1)
			go func() {
				log.Println("Start room for", userID)
				client.Run(m.apiEndpoint, m.communityHandle, token, m.channelID)
				wait.Done()
			}()
			time.Sleep(50 * time.Millisecond)
		}
	}

	wait.Wait()
	return nil
}

func (m *Manager) RunCycle(n int) error {
	for x := 0; x < n; x++ {
		clients := []*Client{}
		for i := 1000; i < 1000+m.nClients; i++ {
			userID := fmt.Sprintf("testuser%d", i)
			token, err := m.tokenEngine.Generate(2000*time.Hour, map[string]any{"id": userID})
			if err != nil {
				return err
			}

			client, err := NewClient(m.serverAddress, userID, token, m.apiEndpoint,
				m.communityHandle, m.channelID)
			if err != nil {
				log.Println("Failed to connect for user", userID, ":", err)
				continue
			}

			clients = append(clients, client)

			go func() {
				log.Println("Start room for", userID)
				client.Run(m.apiEndpoint, m.communityHandle, token, m.channelID)
			}()
			time.Sleep(50 * time.Millisecond)
		}

		time.Sleep(time.Minute)
		for _, c := range clients {
			c.Close()
		}

		time.Sleep(10 * time.Minute)
	}

	return nil
}
