package chainio

import (
	"context"
	"errors"
	avs "github.com/ExocoreNetwork/exocore-avs/contracts/bindings/avs"
	"github.com/ExocoreNetwork/exocore-avs/core/chainio/eth"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

type ExoWriter interface {
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

	OperatorSubmitTask(
		ctx context.Context,
		taskID uint64,
		taskResponse []byte,
		blsSignature []byte,
		taskContractAddress string,
		stage string,
	) (*gethtypes.Receipt, error)

	RegisterOperatorToExocore(
		ctx context.Context,
		metaInfo string,
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
func (w *ExoChainWriter) RegisterBLSPublicKey(
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
func (w *ExoChainWriter) CreateNewTask(
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

func (w *ExoChainWriter) OperatorSubmitTask(
	ctx context.Context,
	taskID uint64,
	taskResponse []byte,
	blsSignature []byte,
	taskContractAddress string,
	stage string,
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
		stage)
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

func (w *ExoChainWriter) RegisterOperatorToExocore(
	ctx context.Context,
	metaInfo string,
) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.RegisterOperatorToExocore(
		noSendTxOpts,
		metaInfo)
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
