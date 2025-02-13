package operator

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
}
func (o *Operator) RegisterOperatorWithExocore() error {

	for _, avsReader := range o.avsReaders {
		flag, err := avsReader.IsOperator(&bind.CallOpts{}, o.operatorAddr.String())
		if err != nil {
			o.logger.Error("Cannot exec IsOperator", "err", err)
			return err
		}
		if !flag {
			o.logger.Info("Operator is not registered.")
			panic(fmt.Sprintf("Operator is not registered: %s", o.operatorAddr.String()))

		}
		if err != nil {
			o.logger.Errorf("Error registering operator with exocore")
			return err
		}
	}

	return nil
}

// RegisterOperatorWithAvs Registration specific functions
func (o *Operator) RegisterOperatorWithAvs() error {
	for _, avsWriter := range o.avsWriters {
		_, err := avsWriter.RegisterOperatorToAVS(context.Background())
		if err != nil {
			o.logger.Error("Avs failed to RegisterOperatorToAVS", "err", err)
			return err
		}
	}

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
