package chainio

import (
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	gethcommon "github.com/ethereum/go-ethereum/common"

	avs "github.com/imua-xyz/imua-avs/contracts/bindings/avs"
	"github.com/imua-xyz/imua-avs/core/chainio/eth"
)

type ExocoreContractBindings struct {
	AvsAddr    gethcommon.Address
	AVSManager *avs.ContracthelloWorld
}

func NewExocoreContractBindings(
	avsAddr gethcommon.Address,
	ethclient eth.EthClient,
	logger logging.Logger,
) (*ExocoreContractBindings, error) {
	contractAvsManager, err := avs.NewContracthelloWorld(avsAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch Avs contract", "err", err)
		return nil, err
	}

	return &ExocoreContractBindings{
		AvsAddr:    avsAddr,
		AVSManager: contractAvsManager,
	}, nil
}
