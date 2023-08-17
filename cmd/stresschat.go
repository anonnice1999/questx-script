package main

import (
	"os"

	"github.com/anonnice1999/questx-script/internal/stresschat"
	"github.com/urfave/cli/v2"
)

func startStressChat(cCtx *cli.Context) error {
	tokenSecret := os.Getenv("TOKEN_SECRET")
	nClients := parseInt(getEnv("N_CLIENTS", "30"))
	apiEndpoint := getEnv("API_ENDPOINT", "http://localhost:8080")
	communityHandle := getEnv("COMMUNITY_HANDLE", "unknown-handle")
	channelID := parseInt64(getEnv("CHANNEL_ID", "30"))
	chatServer := getEnv("CHAT_SERVER", "ws://localhost:8082")

	manager := stresschat.NewManager(chatServer, tokenSecret, apiEndpoint,
		communityHandle, channelID, nClients)
	return manager.RunCycle(40)
}
