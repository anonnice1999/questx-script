package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cliapp := cli.NewApp()
	cliapp.Action = cli.ShowAppHelp
	cliapp.Name = "Script"
	cliapp.Usage = ""
	cliapp.Commands = []*cli.Command{
		{
			Action:    startTemplate,
			Name:      "generate-template",
			Usage:     "Generate template",
			ArgsUsage: "<genesisPath>",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "endpoint",
					DefaultText: "http://localhost:8080",
					Aliases:     []string{"e"},
				},
				&cli.StringFlag{
					Name:     "token",
					Required: true,
					Aliases:  []string{"t"},
				},
			},
		},
	}

	if err := cliapp.Run(os.Args); err != nil {
		panic(err)
	}
}
