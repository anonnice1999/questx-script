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
		},
		{
			Action:    startStressGame,
			Name:      "stress-game",
			Usage:     "Stress test game",
			ArgsUsage: "<genesisPath>",
		},
		{
			Action:    startStressChat,
			Name:      "stress-chat",
			Usage:     "Stress test chat",
			ArgsUsage: "<genesisPath>",
		},
		{
			Action:    startWallet,
			Name:      "wallet",
			Usage:     "generate wallet",
			ArgsUsage: "<genesisPath>",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "nonce",
					DefaultText: "",
					Aliases:     []string{"n"},
				},
			},
		},
	}

	if err := cliapp.Run(os.Args); err != nil {
		panic(err)
	}
}
