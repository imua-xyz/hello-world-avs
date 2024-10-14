package chainio

import (
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	gethcommon "github.com/ethereum/go-ethereum/common"

	avs "github.com/ExocoreNetwork/exocore-avs/contracts/bindings/avs"
	"github.com/ExocoreNetwork/exocore-avs/core/chainio/eth"
)

type ExocoreContractBindings struct {
	AvsAddr    gethcommon.Address
	AVSManager *avs.Contractavsservice
}

func NewExocoreContractBindings(
	avsAddr gethcommon.Address,
	ethclient eth.EthClient,
	logger logging.Logger,
) (*ExocoreContractBindings, error) {
	contractAvsManager, err := avs.NewContractavsservice(avsAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch Avs contract", "err", err)
		return nil, err
	}

	return &ExocoreContractBindings{
		AvsAddr:    avsAddr,
		AVSManager: contractAvsManager,
	}, nil
}
