package operator

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func (o *Operator) registerOperatorOnStartup(operatorEcdsaPrivateKey *ecdsa.PrivateKey) {
	// err := o.RegisterOperatorWithExocore()
	//if err != nil {
	//	// This error might only be that the operator was already registered with exocore, so we don't want to fatal
	//	o.logger.Error("Error registering operator with exocore", "err", err)
	//} else {
	//	o.logger.Infof("Registered operator with exocore")
	//}
	//
	//// err = o.RegisterOperatorWithAvs(operatorEcdsaPrivateKey)
	//if err != nil {
	//	o.logger.Fatal("Error registering operator with avs", "err", err)
	//}
	o.logger.Infof("Registered operator with avs")
}

//func (o *Operator) RegisterOperatorWithExocore() error {
//	op := exodkTypes.Operator{
//		Address:                 o.operatorAddr.String(),
//		EarningsReceiverAddress: o.operatorAddr.String(),
//	}
//	_, err := o.avsWriter.RegisterAsOperator(context.Background(), op)
//	if err != nil {
//		o.logger.Errorf("Error registering operator with exocore")
//		return err
//	}
//	return nil
//}

// RegisterOperatorWithAvs Registration specific functions
//func (o *Operator) RegisterOperatorWithAvs(
//	operatorEcdsaKeyPair *ecdsa.PrivateKey,
//) error {
//	// hardcode these things for now
//	quorumNumbers := []byte{0}
//	socket := "Not Needed"
//	operatorToAvsRegistrationSigSalt := [32]byte{123}
//	curBlockNum, err := o.ethClient.BlockNumber(context.Background())
//	if err != nil {
//		o.logger.Errorf("Unable to get current block number")
//		return err
//	}
//	curBlock, err := o.ethClient.BlockByNumber(context.Background(), big.NewInt(int64(curBlockNum)))
//	if err != nil {
//		o.logger.Errorf("Unable to get current block")
//		return err
//	}
//	sigValidForSeconds := int64(1_000_000)
//	operatorToAvsRegistrationSigExpiry := big.NewInt(int64(curBlock.Time()) + sigValidForSeconds)
//	_, err = o.avsWriter.RegisterOperatorInQuorumWithAVSRegistryCoordinator(
//		context.Background(),
//		operatorEcdsaKeyPair, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry,
//		o.blsKeypair, quorumNumbers, socket,
//	)
//	if err != nil {
//		o.logger.Errorf("Unable to register operator with avs registry coordinator")
//		return err
//	}
//	o.logger.Infof("Registered operator with avs registry coordinator.")
//
//	return nil
//}

// OperatorStatus PRINTING STATUS OF OPERATOR: 1
// operator address: 0xa0ee7a142d267c1f36714e4a8f75612f20a79720
// dummy token balance: 0
// delegated shares in dummyTokenStrat: 200
// operator pubkey hash in AVS pubkey compendium (0 if not registered): 0x4b7b8243d970ff1c90a7c775c008baad825893ec6e806dfa5d3663dc093ed17f
// operator is opted in to exocore: true
// operator is opted in to playgroundAVS (aka can be slashed): true
// operator status in AVS registry: REGISTERED
//
//	operatorId: 0x4b7b8243d970ff1c90a7c775c008baad825893ec6e806dfa5d3663dc093ed17f
//	middlewareTimesLen (# of stake updates): 0
//
// operator is frozen: false
type OperatorStatus struct {
	EcdsaAddress string
	// pubkey compendium related
	PubkeysRegistered bool
	G1Pubkey          string
	G2Pubkey          string
	// avs related
	RegisteredWithAvs bool
	OperatorId        string
}

func (o *Operator) PrintOperatorStatus() error {
	fmt.Println("Printing operator status")
	registeredWithAvs := o.operatorId != [32]byte{}
	operatorStatus := OperatorStatus{
		EcdsaAddress:      o.operatorAddr.String(),
		PubkeysRegistered: true,
		G1Pubkey:          o.blsKeypair.GetPubKeyG1().String(),
		G2Pubkey:          o.blsKeypair.GetPubKeyG2().String(),
		RegisteredWithAvs: registeredWithAvs,
		OperatorId:        hex.EncodeToString(o.operatorId[:]),
	}
	operatorStatusJson, err := json.MarshalIndent(operatorStatus, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(operatorStatusJson))
	return nil
}
