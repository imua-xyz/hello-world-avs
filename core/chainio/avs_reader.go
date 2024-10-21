package chainio

import (
	sdkmath "cosmossdk.io/math"
	avs "github.com/ExocoreNetwork/exocore-avs/contracts/bindings/avs"
	"github.com/ExocoreNetwork/exocore-avs/core/chainio/eth"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
)

type EXOReader interface {
	GetOptInOperators(
		opts *bind.CallOpts,
		avsAddress string,
	) ([]string, error)

	GetRegisteredPubkey(
		opts *bind.CallOpts,
		operator string,
	) ([]byte, error)
	GtAVSUSDValue(
		opts *bind.CallOpts,
		avsAddress string,
	) (sdkmath.LegacyDec, error)

	GetOperatorOptedUSDValue(
		opts *bind.CallOpts,
		avsAddress string,
		operatorAddr string,
	) (sdkmath.LegacyDec, error)
	GetAVSInfo(
		opts *bind.CallOpts,
		avsAddress string,
	) (string, error)
	GetTaskInfo(
		opts *bind.CallOpts,
		avsAddress string,
		taskID uint64,
	) ([]uint64, error)
	IsOperator(
		opts *bind.CallOpts,
		operator string,
	) (bool, error)
}

type EXOChainReader struct {
	logger     logging.Logger
	avsManager avs.Contractavsservice
	ethClient  eth.EthClient
}

// forces EthReader to implement the chainio.Reader interface
var _ EXOReader = (*EXOChainReader)(nil)

func NewExoChainReader(
	avsManager avs.Contractavsservice,
	logger logging.Logger,
	ethClient eth.EthClient,
) *EXOChainReader {
	return &EXOChainReader{
		avsManager: avsManager,
		logger:     logger,
		ethClient:  ethClient,
	}
}

func BuildExoChainReader(
	avsAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
) (*EXOChainReader, error) {
	exoContractBindings, err := NewExocoreContractBindings(
		avsAddr,
		ethClient,
		logger,
	)
	if err != nil {
		return nil, err
	}
	return NewExoChainReader(
		*exoContractBindings.AVSManager,
		logger,
		ethClient,
	), nil
}

func (r *EXOChainReader) GetOptInOperators(
	opts *bind.CallOpts,
	avsAddress string,
) ([]string, error) {
	operators, err := r.avsManager.GetOptInOperators(
		opts,
		gethcommon.HexToAddress(avsAddress))
	if err != nil {
		r.logger.Error("Failed to GetOptInOperators ", "err", err)
		return nil, err
	}
	return operators, nil
}

func (r *EXOChainReader) GetRegisteredPubkey(opts *bind.CallOpts, operator string) ([]byte, error) {
	pukKey, err := r.avsManager.GetRegisteredPubkey(
		opts,
		operator)
	if err != nil {
		r.logger.Error("Failed to GetRegisteredPubkey ", "err", err)
		return nil, err
	}
	return pukKey, nil
}

func (r *EXOChainReader) GtAVSUSDValue(opts *bind.CallOpts, avsAddress string) (sdkmath.LegacyDec, error) {
	amount, err := r.avsManager.GetAVSUSDValue(
		opts,
		gethcommon.HexToAddress(avsAddress))
	if err != nil {
		r.logger.Error("Failed to GtAVSUSDValue ", "err", err)
		return sdkmath.LegacyDec{}, err
	}
	return sdkmath.LegacyNewDecFromBigInt(amount), nil
}

func (r *EXOChainReader) GetOperatorOptedUSDValue(opts *bind.CallOpts, avsAddress string, operatorAddr string) (sdkmath.LegacyDec, error) {
	amount, err := r.avsManager.GetOperatorOptedUSDValue(
		opts,
		gethcommon.HexToAddress(avsAddress), operatorAddr)
	if err != nil {
		r.logger.Error("Failed to GetOperatorOptedUSDValue ", "err", err)
		return sdkmath.LegacyDec{}, err
	}
	return sdkmath.LegacyNewDecFromBigInt(amount), nil
}

func (r *EXOChainReader) GetAVSInfo(opts *bind.CallOpts, avsAddress string) (string, error) {
	epochIdentifier, err := r.avsManager.GetAVSInfo(
		opts,
		gethcommon.HexToAddress(avsAddress))
	if err != nil {
		r.logger.Error("Failed to GetAVSInfo ", "err", err)
		return "", err
	}
	return epochIdentifier, nil
}
func (r *EXOChainReader) GetTaskInfo(opts *bind.CallOpts, avsAddress string, taskID uint64) ([]uint64, error) {
	info, err := r.avsManager.GetTaskInfo(
		opts,
		gethcommon.HexToAddress(avsAddress), taskID)
	if err != nil {
		r.logger.Error("Failed to GetTaskInfo ", "err", err)
		return nil, err
	}
	return info, nil
}

func (r *EXOChainReader) IsOperator(opts *bind.CallOpts, operator string) (bool, error) {
	flag, err := r.avsManager.IsOperator(
		opts,
		operator)
	if err != nil {
		r.logger.Error("Failed to exec IsOperator ", "err", err)
		return false, err
	}
	return flag, nil
}
