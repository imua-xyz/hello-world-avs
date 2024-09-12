package config

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli"

	"github.com/ExocoreNetwork/exocore-sdk/chainio/clients/eth"
	"github.com/ExocoreNetwork/exocore-sdk/chainio/txmgr"
	"github.com/ExocoreNetwork/exocore-sdk/crypto/bls"
	sdklogging "github.com/ExocoreNetwork/exocore-sdk/logging"
	"github.com/ExocoreNetwork/exocore-sdk/signerv2"

	sdkutils "github.com/ExocoreNetwork/exocore-sdk/utils"
)

type Config struct {
	EcdsaPrivateKey *ecdsa.PrivateKey
	BlsPrivateKey   *bls.PrivateKey
	Logger          sdklogging.Logger
	// we need the url for the exocore-sdk currently... eventually standardize api to
	// only take an ethclient or an rpcUrl (and build the ethclient at each constructor site)
	EthHttpRpcUrl              string
	EthWsRpcUrl                string
	EthHttpClient              eth.EthClient
	EthWsClient                eth.EthClient
	OperatorStateRetrieverAddr common.Address
	AvsRegistryCoordinatorAddr common.Address
	AggregatorServerIpPortAddr string
	RegisterOperatorOnStartup  bool
	// json:"-" skips this field when marshaling (only used for logging to stdout), since SignerFn doesnt implement marshalJson
	SignerFn          signerv2.SignerFn `json:"-"`
	TxMgr             txmgr.TxManager
	AggregatorAddress common.Address
}

// ConfigRaw These are read from ConfigFileFlag
type ConfigRaw struct {
	Environment                sdklogging.LogLevel `yaml:"environment"`
	EthRpcUrl                  string              `yaml:"eth_rpc_url"`
	EthWsUrl                   string              `yaml:"eth_ws_url"`
	AggregatorServerIpPortAddr string              `yaml:"aggregator_server_ip_port_address"`
	RegisterOperatorOnStartup  bool                `yaml:"register_operator_on_startup"`
}

// AvsDeploymentRaw These are read from CredibleSquaringDeploymentFileFlag
type AvsDeploymentRaw struct {
	Addresses AvsContractsRaw `json:"addresses"`
}
type AvsContractsRaw struct {
	RegistryCoordinatorAddr    string `json:"registryCoordinator"`
	OperatorStateRetrieverAddr string `json:"operatorStateRetriever"`
}

// NewConfig parses config file to read from flags or environment variables
// Note: This config is shared by challenger and aggregator, and so we put in the core.
// Operator has a different config and is meant to be used by the operator CLI.
func NewConfig(ctx *cli.Context) (*Config, error) {

	var configRaw ConfigRaw
	configFilePath := ctx.GlobalString(FileFlag.Name)
	if configFilePath != "" {
		sdkutils.ReadYamlConfig(configFilePath, &configRaw)
	}

	var credibleSquaringDeploymentRaw AvsDeploymentRaw
	credibleSquaringDeploymentFilePath := ctx.GlobalString(CredibleSquaringDeploymentFileFlag.Name)
	if _, err := os.Stat(credibleSquaringDeploymentFilePath); errors.Is(err, os.ErrNotExist) {
		panic("Path " + credibleSquaringDeploymentFilePath + " does not exist")
	}
	sdkutils.ReadJsonConfig(credibleSquaringDeploymentFilePath, &credibleSquaringDeploymentRaw)

	logger, err := sdklogging.NewZapLogger(configRaw.Environment)
	if err != nil {
		return nil, err
	}

	ethRpcClient, err := eth.NewClient(configRaw.EthRpcUrl)
	if err != nil {
		logger.Errorf("Cannot create http ethclient", "err", err)
		return nil, err
	}

	ethWsClient, err := eth.NewClient(configRaw.EthWsUrl)
	if err != nil {
		logger.Errorf("Cannot create ws ethclient", "err", err)
		return nil, err
	}

	ecdsaPrivateKeyString := ctx.GlobalString(EcdsaPrivateKeyFlag.Name)
	if ecdsaPrivateKeyString[:2] == "0x" {
		ecdsaPrivateKeyString = ecdsaPrivateKeyString[2:]
	}
	ecdsaPrivateKey, err := crypto.HexToECDSA(ecdsaPrivateKeyString)
	if err != nil {
		logger.Errorf("Cannot parse ecdsa private key", "err", err)
		return nil, err
	}

	aggregatorAddr, err := sdkutils.EcdsaPrivateKeyToAddress(ecdsaPrivateKey)
	if err != nil {
		logger.Error("Cannot get operator address", "err", err)
		return nil, err
	}

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainId)
	if err != nil {
		panic(err)
	}
	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, logger, signerV2, aggregatorAddr)

	config := &Config{
		EcdsaPrivateKey:            ecdsaPrivateKey,
		Logger:                     logger,
		EthWsRpcUrl:                configRaw.EthWsUrl,
		EthHttpRpcUrl:              configRaw.EthRpcUrl,
		EthHttpClient:              ethRpcClient,
		EthWsClient:                ethWsClient,
		OperatorStateRetrieverAddr: common.HexToAddress(credibleSquaringDeploymentRaw.Addresses.OperatorStateRetrieverAddr),
		AvsRegistryCoordinatorAddr: common.HexToAddress(credibleSquaringDeploymentRaw.Addresses.RegistryCoordinatorAddr),
		AggregatorServerIpPortAddr: configRaw.AggregatorServerIpPortAddr,
		RegisterOperatorOnStartup:  configRaw.RegisterOperatorOnStartup,
		SignerFn:                   signerV2,
		TxMgr:                      txMgr,
		AggregatorAddress:          aggregatorAddr,
	}
	config.validate()
	return config, nil
}

func (c *Config) validate() {
	// TODO: make sure every pointer is non-nil
	if c.OperatorStateRetrieverAddr == common.HexToAddress("") {
		panic("Config: BLSOperatorStateRetrieverAddr is required")
	}
	if c.AvsRegistryCoordinatorAddr == common.HexToAddress("") {
		panic("Config: AvsRegistryCoordinatorAddr is required")
	}
}

var (
	FileFlag = cli.StringFlag{
		Name:     "config",
		Required: true,
		Usage:    "Load configuration from `FILE`",
	}
	CredibleSquaringDeploymentFileFlag = cli.StringFlag{
		Name:     "credible-squaring-deployment",
		Required: true,
		Usage:    "Load credible squaring contract addresses from `FILE`",
	}
	EcdsaPrivateKeyFlag = cli.StringFlag{
		Name:     "ecdsa-private-key",
		Usage:    "Ethereum private key",
		Required: true,
		EnvVar:   "ECDSA_PRIVATE_KEY",
	}
	/* Optional Flags */
)

var requiredFlags = []cli.Flag{
	FileFlag,
	CredibleSquaringDeploymentFileFlag,
	EcdsaPrivateKeyFlag,
}

var optionalFlags = []cli.Flag{}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag
