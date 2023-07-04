package stressgame

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/anonnice1999/questx-script/pkg/token"
)

type Manager struct {
	serverAddress string
	tokenEngine   token.Engine
	mapHeight     int
	mapWidth      int
	nClients      int
}

func NewManager(serverAddress, tokenSecret string, mapWidth, mapHeight, nClients int) *Manager {
	return &Manager{
		serverAddress: serverAddress,
		tokenEngine:   token.NewEngine(tokenSecret),
		mapWidth:      mapWidth,
		mapHeight:     mapHeight,
		nClients:      nClients,
	}
}

func (m *Manager) Run() error {
	wait := sync.WaitGroup{}
	wait.Add(m.nClients)

	for i := 0; i < m.nClients; i++ {
		user := fmt.Sprintf("testuser%d", i)
		token, err := m.tokenEngine.Generate(time.Minute, map[string]any{"id": user})
		if err != nil {
			return err
		}

		log.Println("Connect to", m.serverAddress)
		client, err := NewClient(m.serverAddress, token, m.mapWidth, m.mapHeight)
		if err != nil {
			return err
		}

		go func() {
			log.Println("Start room for", user)
			client.Run()
			wait.Done()
		}()
	}

	wait.Wait()
	return nil
}
