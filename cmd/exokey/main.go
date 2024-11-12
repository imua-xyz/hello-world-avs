package main

import (
	"fmt"
	"os"

	"github.com/ExocoreNetwork/exocore-avs/cmd/exokey/import"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "exokey"
	app.Description = "Exocore key manager"
	app.Commands = []*cli.Command{
		_import.Command,
	}

	app.Usage = "Import keys for testing purpose"

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

}
