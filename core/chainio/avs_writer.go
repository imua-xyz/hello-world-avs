package chainio

import (
	"context"
	"crypto/ecdsa"
	"errors"
	avs "github.com/ExocoreNetwork/exocore-avs/contracts/bindings/avs"
	"github.com/ExocoreNetwork/exocore-avs/core/chainio/eth"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type EXOWriter interface {
	RegisterAVSToExocore(
		ctx context.Context,
		avsName string,
		minStakeAmount uint64,
		taskAddr gethcommon.Address,
		slashAddr gethcommon.Address,
		rewardAddr gethcommon.Address,
		avsOwnerAddress []string,
		assetIds []string,
		avsUnbondingPeriod uint64,
		minSelfDelegation uint64,
		epochIdentifier string,
		params []uint64,
	) (*gethtypes.Receipt, error)

	RegisterBLSPublicKey(
		ctx context.Context,
		name string,
		pubKey []byte,
		pubKeyRegistrationSignature []byte,
		pubKeyRegistrationMessageHash []byte,
	) (*gethtypes.Receipt, error)

	CreateNewTask(
		ctx context.Context,
		name string,
		taskResponsePeriod uint64,
		taskChallengePeriod uint64,
		thresholdPercentage uint64,
		taskStatisticalPeriod uint64,
	) (*gethtypes.Receipt, error)
}

type EXOChainWriter struct {
	avsManager     avs.Contractavsservice
	exoChainReader EXOReader
	ethClient      eth.EthClient
	logger         logging.Logger
	txMgr          txmgr.TxManager
}

var _ EXOWriter = (*EXOChainWriter)(nil)

func NewELChainWriter(
	avsManager avs.Contractavsservice,
	exoChainReader EXOReader,
	ethClient eth.EthClient,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) *EXOChainWriter {
	return &EXOChainWriter{
		avsManager:     avsManager,
		exoChainReader: exoChainReader,
		logger:         logger,
		ethClient:      ethClient,
		txMgr:          txMgr,
	}
}

func BuildELChainWriter(
	avsAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) (*EXOChainWriter, error) {
	exoContractBindings, err := NewExocoreContractBindings(
		avsAddr,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	elChainReader := NewExoChainReader(
		*exoContractBindings.AVSManager,
		logger,
		ethClient,
	)
	return NewELChainWriter(
		*exoContractBindings.AVSManager,
		elChainReader,
		ethClient,
		logger,
		txMgr,
	), nil
}

func (w *EXOChainWriter) RegisterAVSToExocore(
	ctx context.Context,
	avsName string,
	minStakeAmount uint64,
	taskAddr gethcommon.Address,
	slashAddr gethcommon.Address,
	rewardAddr gethcommon.Address,
	avsOwnerAddress []string,
	assetIds []string,
	avsUnbondingPeriod uint64,
	minSelfDelegation uint64,
	epochIdentifier string,
	params []uint64,
) (*gethtypes.Receipt, error) {

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.RegisterAVS(noSendTxOpts,
		avsName,
		minStakeAmount,
		taskAddr,
		slashAddr,
		rewardAddr,
		avsOwnerAddress,
		assetIds,
		avsUnbondingPeriod,
		minSelfDelegation,
		epochIdentifier,
		params)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())

	return receipt, nil
}
func (w *EXOChainWriter) RegisterBLSPublicKey(
	ctx context.Context,
	name string,
	pubKey []byte,
	pubKeyRegistrationSignature []byte,
	pubKeyRegistrationMessageHash []byte,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.RegisterBLSPublicKey(
		noSendTxOpts,
		name,
		pubKey,
		pubKeyRegistrationSignature,
		pubKeyRegistrationMessageHash)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())

	return receipt, nil
}
func (w *EXOChainWriter) CreateNewTask(
	ctx context.Context,
	name string,
	taskResponsePeriod uint64,
	taskChallengePeriod uint64,
	thresholdPercentage uint64,
	taskStatisticalPeriod uint64,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.CreateNewTask(
		noSendTxOpts,
		name,
		taskResponsePeriod,
		taskChallengePeriod,
		thresholdPercentage,
		taskStatisticalPeriod)
	if err != nil {
		return nil, err
	}
	receipt, err := w.txMgr.Send(ctx, tx)
	if err != nil {
		return nil, errors.New("failed to send tx with err: " + err.Error())
	}
	w.logger.Infof("tx hash: %s", tx.Hash().String())

	return receipt, nil
}

func DeployAVS(
	ethClient eth.EthClient,
	logger logging.Logger,
	key ecdsa.PrivateKey,
	chainID *big.Int,
) (gethcommon.Address, string, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(&key, chainID)
	if err != nil {
		logger.Fatalf("Failed to make transactor: %v", err)
	}

	address, tx, _, err := avs.DeployContractavsservice(auth, ethClient)
	if err != nil {
		logger.Infof("deploy err: %s", err.Error())
		return gethcommon.Address{}, "", errors.New("failed to deploy contract with err: " + err.Error())
	}
	logger.Infof("tx hash: %s", tx.Hash().String())
	logger.Infof("contract address: %s", address.String())

	return address, tx.Hash().String(), nil
}
