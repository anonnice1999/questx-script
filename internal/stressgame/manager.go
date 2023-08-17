package stressgame

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
	mapHeight       int
	mapWidth        int
	nClients        int
	apiEndpoint     string
	characterID     string
	communityHandle string
}

func NewManager(serverAddress, tokenSecret, apiEndpoint, characterID, communityHandle string, mapWidth, mapHeight, nClients int) *Manager {
	return &Manager{
		serverAddress:   serverAddress,
		tokenEngine:     token.NewEngine(tokenSecret),
		mapWidth:        mapWidth,
		mapHeight:       mapHeight,
		nClients:        nClients,
		apiEndpoint:     apiEndpoint,
		characterID:     characterID,
		communityHandle: communityHandle,
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
			token, err := m.tokenEngine.Generate(time.Minute, map[string]any{"id": userID})
			if err != nil {
				return err
			}

			log.Println("Connect to", m.serverAddress)
			client, err := NewClient(m.serverAddress, userID, token, m.apiEndpoint, m.characterID,
				m.communityHandle, m.mapWidth, m.mapHeight)
			if err != nil {
				log.Println("Failed to connect for user", userID, ":", err)
				continue
			}

			wait.Add(1)
			go func() {
				log.Println("Start room for", userID)
				client.Run()
				wait.Done()
			}()
			time.Sleep(50 * time.Millisecond)
		}
	}

	wait.Wait()
	return nil
}
