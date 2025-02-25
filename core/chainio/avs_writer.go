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

type ExoWriter interface {
	RegisterAVSToExocore(
		ctx context.Context,
		params avs.AVSParams,
	) (*gethtypes.Receipt, error)

	RegisterBLSPublicKey(
		ctx context.Context,
		avsAddr string,
		pubKey []byte,
		pubKeyRegistrationSignature []byte,
	) (*gethtypes.Receipt, error)

	CreateNewTask(
		ctx context.Context,
		name string,
		numberToBeSquared uint64,
		taskResponsePeriod uint64,
		taskChallengePeriod uint64,
		thresholdPercentage uint8,
		taskStatisticalPeriod uint64,
	) (*gethtypes.Receipt, error)

	OperatorSubmitTask(
		ctx context.Context,
		taskID uint64,
		taskResponse []byte,
		blsSignature []byte,
		taskContractAddress string,
		phase uint8,
	) (*gethtypes.Receipt, error)

	Challenge(
		ctx context.Context,
		req avs.AvsServiceContractChallengeReq,
	) (*gethtypes.Receipt, error)

	RegisterOperatorToAVS(
		ctx context.Context,
	) (*gethtypes.Receipt, error)
}

type ExoChainWriter struct {
	avsManager     avs.ContracthelloWorld
	exoChainReader ExoReader
	ethClient      eth.EthClient
	logger         logging.Logger
	txMgr          txmgr.TxManager
}

var _ ExoWriter = (*ExoChainWriter)(nil)

func NewExoChainWriter(
	avsManager avs.ContracthelloWorld,
	exoChainReader ExoReader,
	ethClient eth.EthClient,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) *ExoChainWriter {
	return &ExoChainWriter{
		avsManager:     avsManager,
		exoChainReader: exoChainReader,
		logger:         logger,
		ethClient:      ethClient,
		txMgr:          txMgr,
	}
}

func BuildExoChainWriter(
	avsAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) (*ExoChainWriter, error) {
	exoContractBindings, err := NewExocoreContractBindings(
		avsAddr,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	exoChainReader := NewExoChainReader(
		*exoContractBindings.AVSManager,
		logger,
		ethClient,
	)
	return NewExoChainWriter(
		*exoContractBindings.AVSManager,
		exoChainReader,
		ethClient,
		logger,
		txMgr,
	), nil
}

func (w *ExoChainWriter) RegisterAVSToExocore(
	ctx context.Context,
	params avs.AVSParams,
) (*gethtypes.Receipt, error) {

	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.RegisterAVS(
		noSendTxOpts,
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
func (w *ExoChainWriter) RegisterBLSPublicKey(
	ctx context.Context,
	avsAddr string,
	pubKey []byte,
	pubKeyRegistrationSignature []byte,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.RegisterBLSPublicKey(
		noSendTxOpts,
		gethcommon.HexToAddress(avsAddr),
		pubKey,
		pubKeyRegistrationSignature)
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
func (w *ExoChainWriter) CreateNewTask(
	ctx context.Context,
	name string,
	numberToBeSquared uint64,
	taskResponsePeriod uint64,
	taskChallengePeriod uint64,
	thresholdPercentage uint8,
	taskStatisticalPeriod uint64,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.CreateNewTask(
		noSendTxOpts,
		name,
		numberToBeSquared,
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

func (w *ExoChainWriter) OperatorSubmitTask(
	ctx context.Context,
	taskID uint64,
	taskResponse []byte,
	blsSignature []byte,
	taskContractAddress string,
	phase uint8,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.OperatorSubmitTask(
		noSendTxOpts,
		taskID,
		taskResponse,
		blsSignature,
		gethcommon.HexToAddress(taskContractAddress),
		phase)
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
func (w *ExoChainWriter) Challenge(ctx context.Context, req avs.AvsServiceContractChallengeReq) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.RaiseAndResolveChallenge(
		noSendTxOpts,
		req)
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

func (w *ExoChainWriter) RegisterOperatorToAVS(
	ctx context.Context,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.RegisterOperatorToAVS(
		noSendTxOpts)
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

	address, tx, _, err := avs.DeployContracthelloWorld(auth, ethClient)
	if err != nil {
		logger.Infof("deploy err: %s", err.Error())
		return gethcommon.Address{}, "", errors.New("failed to deploy contract with err: " + err.Error())
	}
	logger.Infof("tx hash: %s", tx.Hash().String())
	logger.Infof("contract address: %s", address.String())

	return address, tx.Hash().String(), nil
}
