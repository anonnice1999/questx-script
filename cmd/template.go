package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/anonnice1999/questx-script/internal/template"
)

func startTemplate(cCtx *cli.Context) error {
	apiEndpoint := getEnv("API_ENDPOINT", "http://localhost:8080")
	accessToken := os.Getenv("ACCESS_TOKEN")

	return template.Run(apiEndpoint, accessToken)
}
