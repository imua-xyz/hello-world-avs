package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/ExocoreNetwork/exocore-avs/challenge"
	"github.com/ExocoreNetwork/exocore-avs/core/config"
	"github.com/ExocoreNetwork/exocore-avs/types"

	sdkutils "github.com/ExocoreNetwork/exocore-sdk/utils"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{config.FileFlag}
	app.Name = "hello-world-demo-challenge"
	app.Usage = "hello-world-demo Challenge"
	app.Description = "Service that challenger listens to AVS contract events, Initiate challenges and validate the tasks already submitted by the operator."

	app.Action = challengeMain
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln("Application failed. Message:", err)
	}
}

func challengeMain(ctx *cli.Context) error {
	log.Println("Initializing challenge")
	configPath := ctx.GlobalString(config.FileFlag.Name)
	nodeConfig := types.NodeConfig{}
	err := sdkutils.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return err
	}
	configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Config:", string(configJson))

	log.Println("initializing challenge")
	operator, err := challenge.NewChallengeFromConfig(nodeConfig)
	if err != nil {
		return err
	}
	log.Println("initialized challenge")

	log.Println("starting challenge")
	err = operator.Exec(context.Background())
	if err != nil {
		return err
	}
	log.Println("started challenge")

	return nil
}
