package chainio

import (
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	"github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/signerv2"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/imua-xyz/imua-avs/core/chainio/eth"
)

type BuildAllConfig struct {
	EthHttpUrl string
	EthWsUrl   string
	AvsAddr    string
	AvsName    string
}

type Clients struct {
	AvsRegistryChainSubscriber *AvsRegistryChainSubscriber
	EXOChainReader             *ExoChainReader
	EXOChainWriter             *ExoChainWriter
	EthHttpClient              *eth.Client
	EthWsClient                *eth.Client
}

func BuildAll(
	config BuildAllConfig,
	signerAddr gethcommon.Address,
	signerFn signerv2.SignerFn,
	logger logging.Logger,
) (*Clients, error) {
	config.validate(logger)

	// creating two types of Eth clients: HTTP and WS
	ethHttpClient, err := eth.NewClient(config.EthHttpUrl)
	if err != nil {
		logger.Error("Failed to create Eth Http client", "err", err)
		return nil, err
	}

	ethWsClient, err := eth.NewClient(config.EthWsUrl)
	if err != nil {
		logger.Error("Failed to create Eth WS client", "err", err)
		return nil, err
	}

	txMgr := txmgr.NewSimpleTxManager(ethHttpClient, logger, signerFn, signerAddr)
	// creating Exo clients: Reader, Writer and Subscriber
	exoChainReader, exoChainWriter, avsRegistrySubscriber, err := config.buildExoClients(
		ethHttpClient,
		txMgr,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create Exo Reader, Writer and Subscriber", "err", err)
		return nil, err
	}

	return &Clients{
		EXOChainReader:             exoChainReader,
		EXOChainWriter:             exoChainWriter,
		AvsRegistryChainSubscriber: avsRegistrySubscriber,
		EthHttpClient:              ethHttpClient,
		EthWsClient:                ethWsClient,
	}, nil

}

func (config *BuildAllConfig) buildExoClients(
	ethHttpClient eth.EthClient,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ExoChainReader, *ExoChainWriter, *AvsRegistryChainSubscriber, error) {
	exoContractBindings, err := NewExocoreContractBindings(
		gethcommon.HexToAddress(config.AvsAddr),
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create ExocoreContractBindings", "err", err)
		return nil, nil, nil, err
	}

	// get the Reader for the Exo contracts
	exoChainReader := NewExoChainReader(
		*exoContractBindings.AVSManager,
		logger,
		ethHttpClient,
	)

	exoChainWriter := NewExoChainWriter(
		*exoContractBindings.AVSManager,
		exoChainReader,
		ethHttpClient,
		logger,
		txMgr,
	)
	if err != nil {
		logger.Error("Failed to create ExoChainWriter", "err", err)
		return nil, nil, nil, err
	}

	avsRegistrySubscriber, err := BuildAvsRegistryChainSubscriber(
		exoContractBindings.AvsAddr,
		ethHttpClient,
		logger,
	)
	if err != nil {
		logger.Error("Failed to create ExoChainSubscriber", "err", err)
		return nil, nil, nil, err
	}
	return exoChainReader, exoChainWriter, avsRegistrySubscriber, err
}

// Very basic validation that makes sure all fields are nonempty
// we might eventually want more sophisticated validation, based on regexp,
// or use something like https://json-schema.org/ (?)
func (config *BuildAllConfig) validate(logger logging.Logger) {
	if config.EthHttpUrl == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing eth http url")
	}
	if config.EthWsUrl == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing eth ws url")
	}
	if config.AvsAddr == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing bls registry coordinator address")
	}
	if config.AvsName == "" {
		logger.Fatalf("BuildAllConfig.validate: Missing avs name")
	}
}
