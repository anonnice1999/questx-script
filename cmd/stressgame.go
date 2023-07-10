package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/anonnice1999/questx-script/internal/stressgame"
)

func startStressGame(cCtx *cli.Context) error {
	tokenSecret := os.Getenv("TOKEN_SECRET")
	gameServer := getEnv("GAME_SERVER", "ws://localhost:8082")
	mapWidth := parseInt(getEnv("GAME_MAP_WIDTH", "4800"))
	mapHeight := parseInt(getEnv("GAME_MAP_HEIGHT", "4800"))
	nClients := parseInt(getEnv("N_CLIENTS", "30"))
	apiEndpoint := getEnv("API_ENDPOINT", "http://localhost:8080")
	characterID := getEnv("INIT_CHARACTER_ID", "unknown-character-id")
	communityHandle := getEnv("COMMUNITY_HANDLE", "unknown-handle")

	manager := stressgame.NewManager(gameServer, tokenSecret, apiEndpoint, characterID,
		communityHandle, mapWidth, mapHeight, nClients)
	return manager.Run()
}
