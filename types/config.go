package types

const (
	TaskResponsePeriod    = uint64(3)
	TaskChallengePeriod   = uint64(3)
	ThresholdPercentage   = uint64(100)
	TaskStatisticalPeriod = uint64(3)
)

type NodeConfig struct {
	// used to set the logger level (true = info, false = debug)
	Production                bool   `yaml:"production"`
	AVSOwnerAddress           string `yaml:"avs_owner_address"`
	OperatorAddress           string `yaml:"operator_address"`
	AVSAddress                string `yaml:"avs_address"`
	EthRpcUrl                 string `yaml:"eth_rpc_url"`
	EthWsUrl                  string `yaml:"eth_ws_url"`
	BlsPrivateKeyStorePath    string `yaml:"bls_private_key_store_path"`
	EcdsaPrivateKeyStorePath  string `yaml:"ecdsa_private_key_store_path"`
	RegisterOperatorOnStartup bool   `yaml:"register_operator_on_startup"`
	NodeApiIpPortAddress      string `yaml:"node_api_ip_port_address"`
	EnableNodeApi             bool   `yaml:"enable_node_api"`
}
