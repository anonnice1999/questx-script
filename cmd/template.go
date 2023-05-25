package main

import (
	"github.com/urfave/cli/v2"

	"github.com/anonnice1999/questx-script/internal/template"
)

func startTemplate(cCtx *cli.Context) error {
	apiEndpoint := cCtx.String("endpoint")
	accessToken := cCtx.String("token")
	if apiEndpoint == "" {
		apiEndpoint = "http://localhost:8080"
	}

	return template.Run(apiEndpoint, accessToken)
}
