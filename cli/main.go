package main

import (
	"log"
	"os"

	"github.com/ExocoreNetwork/exocore-avs/cli/actions"
	"github.com/ExocoreNetwork/exocore-avs/core/config"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{config.FileFlag}
	app.Commands = []cli.Command{
		{
			Name:    "register-operator-with-exocore",
			Aliases: []string{"rel"},
			Usage:   "registers operator with exocore",
			Action:  actions.RegisterOperatorWithExocore,
		},
		{
			Name:    "register-operator-with-avs",
			Aliases: []string{"r"},
			Usage:   "operator opt-in avs ",
			Action:  actions.RegisterOperatorWithAvs,
		},
		{
			Name:    "deregister-operator-with-avs",
			Aliases: []string{"d"},
			Action: func(ctx *cli.Context) error {
				log.Fatal("Command not implemented.")
				return nil
			},
		},
		{
			Name:    "print-operator-status",
			Aliases: []string{"s"},
			Usage:   "prints operator status as viewed from incredible squaring contracts",
			Action:  actions.PrintOperatorStatus,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
