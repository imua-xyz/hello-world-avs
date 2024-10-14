package chainio

import (
	avs "github.com/ExocoreNetwork/exocore-avs/contracts/bindings/avs"
	"github.com/ExocoreNetwork/exocore-avs/core/chainio/eth"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"math/big"
)

type EXOReader interface {
	//	PublicKey          []byte
	//Name               string
	//TotalRewardsEarned *big.Int
	//IsRegistered       bool
}
type Operator struct {
	PublicKey          []byte
	Name               string
	TotalRewardsEarned *big.Int
	IsRegistered       bool
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
