package types

const (
	TaskResponsePeriod    = uint64(3)
	TaskChallengePeriod   = uint64(3)
	ThresholdPercentage   = uint64(100)
	TaskStatisticalPeriod = uint64(3)
)

type NodeConfig struct {
	// used to set the logger level (true = info, false = debug)
	Production                       bool   `yaml:"production"`
	AVSOwnerAddress                  string `yaml:"avs_owner_address"`
	OperatorAddress                  string `yaml:"operator_address"`
	AVSAddress                       string `yaml:"avs_address"`
	EthRpcUrl                        string `yaml:"eth_rpc_url"`
	EthWsUrl                         string `yaml:"eth_ws_url"`
	BlsPrivateKeyStorePath           string `yaml:"bls_private_key_store_path"`
	OperatorEcdsaPrivateKeyStorePath string `yaml:"operator_ecdsa_private_key_store_path"`
	AVSEcdsaPrivateKeyStorePath      string `yaml:"avs_ecdsa_private_key_store_path"`
	RegisterOperatorOnStartup        bool   `yaml:"register_operator_on_startup"`
	NodeApiIpPortAddress             string `yaml:"node_api_ip_port_address"`
	EnableNodeApi                    bool   `yaml:"enable_node_api"`
	// register avs parameters
	MinStakeAmount     uint64   `yaml:"min_stake_amount"`
	AvsOwnerAddresses  []string `yaml:"avs_owner_addresses"`
	AssetIds           []string `yaml:"asset_ids"`
	AvsUnbondingPeriod uint64   `yaml:"avs_unbonding_period"`
	MinSelfDelegation  uint64   `yaml:"min_self_delegation"`
	EpochIdentifier    string   `yaml:"epoch_identifier"`
	Params             []uint64 `yaml:"params"`
}
