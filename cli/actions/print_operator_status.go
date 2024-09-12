package actions

import (
	"encoding/json"
	"log"

	"github.com/ExocoreNetwork/exocore-avs/core/config"
	"github.com/ExocoreNetwork/exocore-avs/operator"
	"github.com/ExocoreNetwork/exocore-avs/types"
	sdkutils "github.com/ExocoreNetwork/exocore-sdk/utils"
	"github.com/urfave/cli"
)

func PrintOperatorStatus(ctx *cli.Context) error {

	configPath := ctx.GlobalString(config.FileFlag.Name)
	nodeConfig := types.NodeConfig{}
	err := sdkutils.ReadYamlConfig(configPath, &nodeConfig)
	if err != nil {
		return err
	}
	// need to make sure we don't register the operator on startup
	// when using the cli commands to register the operator.
	nodeConfig.RegisterOperatorOnStartup = false
	configJson, err := json.MarshalIndent(nodeConfig, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Config:", string(configJson))

	operator, err := operator.NewOperatorFromConfig(nodeConfig)
	if err != nil {
		return err
	}

	err = operator.PrintOperatorStatus()
	if err != nil {
		return err
	}

	return nil
}
