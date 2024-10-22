package operator

import (
	"context"
	"encoding/json"
	use "github.com/ExocoreNetwork/exocore-avs/avs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
)

func (o *Operator) registerOperatorOnStartup() {
	err := o.RegisterOperatorWithExocore()
	if err != nil {
		// This error might only be that the operator was already registered with exocore, so we don't want to fatal
		o.logger.Error("Error registering operator with exocore", "err", err)
	} else {
		o.logger.Infof("Registered operator with exocore")
	}

	err = o.RegisterOperatorWithAvs()
	if err != nil {
		o.logger.Fatal("Error registering operator with avs", "err", err)
	}
	o.logger.Infof("Registered operator with avs")
}
func (o *Operator) RegisterOperatorWithExocore() error {
	flag, err := o.avsReader.IsOperator(&bind.CallOpts{}, o.operatorAddr.String())
	if err != nil {
		o.logger.Error("Cannot exec IsOperator", "err", err)
		return err
	}
	if !flag {
		o.logger.Info("Operator is not registered.")
		_, err = o.avsWriter.RegisterOperatorToExocore(context.Background(), use.GenerateRandomName(8))
		if err != nil {
			o.logger.Error("Avs failed to RegisterOperatorToExocore", "err", err)
			return err
		}
	}
	if err != nil {
		o.logger.Errorf("Error registering operator with exocore")
		return err
	}
	return nil
}

// RegisterOperatorWithAvs Registration specific functions
func (o *Operator) RegisterOperatorWithAvs() error {
	operators, err := o.avsReader.GetOptInOperators(&bind.CallOpts{}, o.avsAddr.String())
	if err != nil {
		o.logger.Error("Cannot exec IsOperator", "err", err)
		return err
	}
	str := strings.Join(operators, "")
	found := strings.Contains(str, o.operatorAddr.String())
	if !found {
		o.logger.Info("Operator is not opt-in this avs.")
		_, err = o.avsWriter.RegisterOperatorToAVS(context.Background())
		if err != nil {
			o.logger.Error("Avs failed to RegisterOperatorToAVS", "err", err)
			return err
		}
	}

	o.logger.Infof("Registered operator with avs.")

	return nil
}

type OperatorStatus struct {
	EcdsaAddress string
	// pubkey compendium related
	PubkeysRegistered bool
	Pubkey            string
	// avs related
	RegisteredWithAvs bool
}

func (o *Operator) PrintOperatorStatus() error {
	o.logger.Info("Printing operator status")
	operatorStatus := OperatorStatus{
		EcdsaAddress:      o.operatorAddr.String(),
		PubkeysRegistered: true,
		Pubkey:            string(o.blsKeypair.PublicKey().Marshal()),
		RegisteredWithAvs: false,
	}
	operatorStatusJson, err := json.MarshalIndent(operatorStatus, "", " ")
	if err != nil {
		return err
	}
	o.logger.Info(string(operatorStatusJson))
	return nil
}
