// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractavsservice

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractavsserviceMetaData contains all meta data concerning the Contractavsservice contract.
var ContractavsserviceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DataLogged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"name\":\"createNewTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregisterOperatorFromAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddress\",\"type\":\"address\"}],\"name\":\"getOptInOperators\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"}],\"name\":\"getRegisteredPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"minStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"avsOwnerAddress\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"assetIds\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSelfDelegation\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64[]\",\"name\":\"params\",\"type\":\"uint64[]\"}],\"name\":\"registerAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyRegistrationSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyRegistrationMessageHash\",\"type\":\"bytes\"}],\"name\":\"registerBLSPublicKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerOperatorToAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"minStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"avsOwnerAddress\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"assetIds\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSelfDelegation\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64[]\",\"name\":\"params\",\"type\":\"uint64[]\"}],\"name\":\"updateAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506119e88061005f6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638da5cb5b116100665780638da5cb5b14610158578063af1991b314610176578063c208dd99146101a6578063c49bb521146101c4578063de16bf46146101f457610093565b80631d4c80071461009857806354c77f71146100c857806355b42cbe146100f85780636f48e1a214610128575b600080fd5b6100b260048036038101906100ad9190610933565b610212565b6040516100bf9190610ab2565b60405180910390f35b6100e260048036038101906100dd9190610caa565b6102a4565b6040516100ef9190610d9c565b60405180910390f35b610112600480360381019061010d9190610db7565b61033b565b60405161011f9190610e55565b60405180910390f35b610142600480360381019061013d9190611060565b610404565b60405161014f9190610d9c565b60405180910390f35b6101606104b0565b60405161016d91906111ef565b60405180910390f35b610190600480360381019061018b919061120a565b6104d4565b60405161019d9190610d9c565b60405180910390f35b6101ae6106a3565b6040516101bb9190610d9c565b60405180910390f35b6101de60048036038101906101d99190611060565b61072d565b6040516101eb9190610d9c565b60405180910390f35b6101fc6107d9565b6040516102099190610d9c565b60405180910390f35b6060600061090173ffffffffffffffffffffffffffffffffffffffff16631d4c8007846040518263ffffffff1660e01b815260040161025191906111ef565b6000604051808303816000875af1158015610270573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061029991906113c6565b905080915050919050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663aa57f77c33888888886040518663ffffffff1660e01b81526004016102ea959493929190611459565b6020604051808303816000875af1158015610309573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061032d91906114f4565b905080915050949350505050565b6060600061090173ffffffffffffffffffffffffffffffffffffffff166355b42cbe846040518263ffffffff1660e01b815260040161037a9190611521565b6000604051808303816000875af1158015610399573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906103c291906115b3565b90507ff8a35bbe260cb3c99c0febbe1d21b55e78ca43bdc6ed00e30007a9fadfe67e43816040516103f39190610e55565b60405180910390a180915050919050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663d9e5daa0338f8f8f8f8f8f8f8f8f8f8f6040518d63ffffffff1660e01b81526004016104589c9b9a999897969594939291906116c9565b6020604051808303816000875af1158015610477573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049b91906114f4565b9050809150509b9a5050505050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006104de610863565b86816000018190525085816040019067ffffffffffffffff16908167ffffffffffffffff168152505084816060019067ffffffffffffffff16908167ffffffffffffffff1681525050828160a0019067ffffffffffffffff16908167ffffffffffffffff168152505083816080019067ffffffffffffffff16908167ffffffffffffffff1681525050600061090173ffffffffffffffffffffffffffffffffffffffff16638bf30a69338a8560405160200161059a9190611848565b604051602081830303815290604052805190602001206040516020016105c09190611895565b6040516020818303038152906040528b8b8b8b6040518863ffffffff1660e01b81526004016105f597969594939291906118b0565b6020604051808303816000875af1158015610614573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061063891906114f4565b90507fc3d2056aaa49a6e50ff7a353c7777b5cda982f2bf1e2214af811056cef07232d82602001513384600001518560400151866060015187608001518860a0015160405161068d979695949392919061193c565b60405180910390a1809250505095945050505050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663d7a2398b336040518263ffffffff1660e01b81526004016106e191906111ef565b6020604051808303816000875af1158015610700573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072491906114f4565b90508091505090565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663cde09950338f8f8f8f8f8f8f8f8f8f8f6040518d63ffffffff1660e01b81526004016107819c9b9a999897969594939291906116c9565b6020604051808303816000875af11580156107a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c491906114f4565b9050809150509b9a5050505050505050505050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663a364f4da336040518263ffffffff1660e01b815260040161081791906111ef565b6020604051808303816000875af1158015610836573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085a91906114f4565b90508091505090565b6040518060c001604052806060815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610900826108d5565b9050919050565b610910816108f5565b811461091b57600080fd5b50565b60008135905061092d81610907565b92915050565b600060208284031215610949576109486108cb565b5b60006109578482850161091e565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156109c65780820151818401526020810190506109ab565b60008484015250505050565b6000601f19601f8301169050919050565b60006109ee8261098c565b6109f88185610997565b9350610a088185602086016109a8565b610a11816109d2565b840191505092915050565b6000610a2883836109e3565b905092915050565b6000602082019050919050565b6000610a4882610960565b610a52818561096b565b935083602082028501610a648561097c565b8060005b85811015610aa05784840389528151610a818582610a1c565b9450610a8c83610a30565b925060208a01995050600181019050610a68565b50829750879550505050505092915050565b60006020820190508181036000830152610acc8184610a3d565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610b16826109d2565b810181811067ffffffffffffffff82111715610b3557610b34610ade565b5b80604052505050565b6000610b486108c1565b9050610b548282610b0d565b919050565b600067ffffffffffffffff821115610b7457610b73610ade565b5b610b7d826109d2565b9050602081019050919050565b82818337600083830152505050565b6000610bac610ba784610b59565b610b3e565b905082815260208101848484011115610bc857610bc7610ad9565b5b610bd3848285610b8a565b509392505050565b600082601f830112610bf057610bef610ad4565b5b8135610c00848260208601610b99565b91505092915050565b600067ffffffffffffffff821115610c2457610c23610ade565b5b610c2d826109d2565b9050602081019050919050565b6000610c4d610c4884610c09565b610b3e565b905082815260208101848484011115610c6957610c68610ad9565b5b610c74848285610b8a565b509392505050565b600082601f830112610c9157610c90610ad4565b5b8135610ca1848260208601610c3a565b91505092915050565b60008060008060808587031215610cc457610cc36108cb565b5b600085013567ffffffffffffffff811115610ce257610ce16108d0565b5b610cee87828801610bdb565b945050602085013567ffffffffffffffff811115610d0f57610d0e6108d0565b5b610d1b87828801610c7c565b935050604085013567ffffffffffffffff811115610d3c57610d3b6108d0565b5b610d4887828801610c7c565b925050606085013567ffffffffffffffff811115610d6957610d686108d0565b5b610d7587828801610c7c565b91505092959194509250565b60008115159050919050565b610d9681610d81565b82525050565b6000602082019050610db16000830184610d8d565b92915050565b600060208284031215610dcd57610dcc6108cb565b5b600082013567ffffffffffffffff811115610deb57610dea6108d0565b5b610df784828501610bdb565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000610e2782610e00565b610e318185610e0b565b9350610e418185602086016109a8565b610e4a816109d2565b840191505092915050565b60006020820190508181036000830152610e6f8184610e1c565b905092915050565b600067ffffffffffffffff82169050919050565b610e9481610e77565b8114610e9f57600080fd5b50565b600081359050610eb181610e8b565b92915050565b600067ffffffffffffffff821115610ed257610ed1610ade565b5b602082029050602081019050919050565b600080fd5b6000610efb610ef684610eb7565b610b3e565b90508083825260208201905060208402830185811115610f1e57610f1d610ee3565b5b835b81811015610f6557803567ffffffffffffffff811115610f4357610f42610ad4565b5b808601610f508982610bdb565b85526020850194505050602081019050610f20565b5050509392505050565b600082601f830112610f8457610f83610ad4565b5b8135610f94848260208601610ee8565b91505092915050565b600067ffffffffffffffff821115610fb857610fb7610ade565b5b602082029050602081019050919050565b6000610fdc610fd784610f9d565b610b3e565b90508083825260208201905060208402830185811115610fff57610ffe610ee3565b5b835b8181101561102857806110148882610ea2565b845260208401935050602081019050611001565b5050509392505050565b600082601f83011261104757611046610ad4565b5b8135611057848260208601610fc9565b91505092915050565b60008060008060008060008060008060006101608c8e031215611086576110856108cb565b5b60008c013567ffffffffffffffff8111156110a4576110a36108d0565b5b6110b08e828f01610bdb565b9b505060206110c18e828f01610ea2565b9a505060406110d28e828f0161091e565b99505060606110e38e828f0161091e565b98505060806110f48e828f0161091e565b97505060a08c013567ffffffffffffffff811115611115576111146108d0565b5b6111218e828f01610f6f565b96505060c08c013567ffffffffffffffff811115611142576111416108d0565b5b61114e8e828f01610f6f565b95505060e061115f8e828f01610ea2565b9450506101006111718e828f01610ea2565b9350506101208c013567ffffffffffffffff811115611193576111926108d0565b5b61119f8e828f01610bdb565b9250506101408c013567ffffffffffffffff8111156111c1576111c06108d0565b5b6111cd8e828f01611032565b9150509295989b509295989b9093969950565b6111e9816108f5565b82525050565b600060208201905061120460008301846111e0565b92915050565b600080600080600060a08688031215611226576112256108cb565b5b600086013567ffffffffffffffff811115611244576112436108d0565b5b61125088828901610bdb565b955050602061126188828901610ea2565b945050604061127288828901610ea2565b935050606061128388828901610ea2565b925050608061129488828901610ea2565b9150509295509295909350565b60006112b46112af84610b59565b610b3e565b9050828152602081018484840111156112d0576112cf610ad9565b5b6112db8482856109a8565b509392505050565b600082601f8301126112f8576112f7610ad4565b5b81516113088482602086016112a1565b91505092915050565b600061132461131f84610eb7565b610b3e565b9050808382526020820190506020840283018581111561134757611346610ee3565b5b835b8181101561138e57805167ffffffffffffffff81111561136c5761136b610ad4565b5b80860161137989826112e3565b85526020850194505050602081019050611349565b5050509392505050565b600082601f8301126113ad576113ac610ad4565b5b81516113bd848260208601611311565b91505092915050565b6000602082840312156113dc576113db6108cb565b5b600082015167ffffffffffffffff8111156113fa576113f96108d0565b5b61140684828501611398565b91505092915050565b600082825260208201905092915050565b600061142b8261098c565b611435818561140f565b93506114458185602086016109a8565b61144e816109d2565b840191505092915050565b600060a08201905061146e60008301886111e0565b81810360208301526114808187611420565b905081810360408301526114948186610e1c565b905081810360608301526114a88185610e1c565b905081810360808301526114bc8184610e1c565b90509695505050505050565b6114d181610d81565b81146114dc57600080fd5b50565b6000815190506114ee816114c8565b92915050565b60006020828403121561150a576115096108cb565b5b6000611518848285016114df565b91505092915050565b6000602082019050818103600083015261153b8184611420565b905092915050565b600061155661155184610c09565b610b3e565b90508281526020810184848401111561157257611571610ad9565b5b61157d8482856109a8565b509392505050565b600082601f83011261159a57611599610ad4565b5b81516115aa848260208601611543565b91505092915050565b6000602082840312156115c9576115c86108cb565b5b600082015167ffffffffffffffff8111156115e7576115e66108d0565b5b6115f384828501611585565b91505092915050565b61160581610e77565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61164081610e77565b82525050565b60006116528383611637565b60208301905092915050565b6000602082019050919050565b60006116768261160b565b6116808185611616565b935061168b83611627565b8060005b838110156116bc5781516116a38882611646565b97506116ae8361165e565b92505060018101905061168f565b5085935050505092915050565b6000610180820190506116df600083018f6111e0565b81810360208301526116f1818e611420565b9050611700604083018d6115fc565b61170d606083018c6111e0565b61171a608083018b6111e0565b61172760a083018a6111e0565b81810360c08301526117398189610a3d565b905081810360e083015261174d8188610a3d565b905061175d6101008301876115fc565b61176b6101208301866115fc565b81810361014083015261177e8185611420565b9050818103610160830152611793818461166b565b90509d9c50505050505050505050505050565b6000819050919050565b6117b9816117a6565b82525050565b600060c08301600083015184820360008601526117dc82826109e3565b91505060208301516117f160208601826117b0565b5060408301516118046040860182611637565b5060608301516118176060860182611637565b50608083015161182a6080860182611637565b5060a083015161183d60a0860182611637565b508091505092915050565b6000602082019050818103600083015261186281846117bf565b905092915050565b6000819050919050565b6000819050919050565b61188f61188a8261186a565b611874565b82525050565b60006118a1828461187e565b60208201915081905092915050565b600060e0820190506118c5600083018a6111e0565b81810360208301526118d78189611420565b905081810360408301526118eb8188610e1c565b90506118fa60608301876115fc565b61190760808301866115fc565b61191460a08301856115fc565b61192160c08301846115fc565b98975050505050505050565b611936816117a6565b82525050565b600060e082019050611951600083018a61192d565b61195e60208301896111e0565b81810360408301526119708188611420565b905061197f60608301876115fc565b61198c60808301866115fc565b61199960a08301856115fc565b6119a660c08301846115fc565b9897505050505050505056fea2646970667358221220da2a2937e6def3cea6af6cb4f72ae027f3396d8897fc4c18a46b492dd706fdd464736f6c63430008190033",
}

// ContractavsserviceABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractavsserviceMetaData.ABI instead.
var ContractavsserviceABI = ContractavsserviceMetaData.ABI

// ContractavsserviceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractavsserviceMetaData.Bin instead.
var ContractavsserviceBin = ContractavsserviceMetaData.Bin

// DeployContractavsservice deploys a new Ethereum contract, binding an instance of Contractavsservice to it.
func DeployContractavsservice(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contractavsservice, error) {
	parsed, err := ContractavsserviceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractavsserviceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contractavsservice{ContractavsserviceCaller: ContractavsserviceCaller{contract: contract}, ContractavsserviceTransactor: ContractavsserviceTransactor{contract: contract}, ContractavsserviceFilterer: ContractavsserviceFilterer{contract: contract}}, nil
}

// Contractavsservice is an auto generated Go binding around an Ethereum contract.
type Contractavsservice struct {
	ContractavsserviceCaller     // Read-only binding to the contract
	ContractavsserviceTransactor // Write-only binding to the contract
	ContractavsserviceFilterer   // Log filterer for contract events
}

// ContractavsserviceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractavsserviceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractavsserviceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractavsserviceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractavsserviceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractavsserviceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractavsserviceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractavsserviceSession struct {
	Contract     *Contractavsservice // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractavsserviceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractavsserviceCallerSession struct {
	Contract *ContractavsserviceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContractavsserviceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractavsserviceTransactorSession struct {
	Contract     *ContractavsserviceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContractavsserviceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractavsserviceRaw struct {
	Contract *Contractavsservice // Generic contract binding to access the raw methods on
}

// ContractavsserviceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractavsserviceCallerRaw struct {
	Contract *ContractavsserviceCaller // Generic read-only contract binding to access the raw methods on
}

// ContractavsserviceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractavsserviceTransactorRaw struct {
	Contract *ContractavsserviceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractavsservice creates a new instance of Contractavsservice, bound to a specific deployed contract.
func NewContractavsservice(address common.Address, backend bind.ContractBackend) (*Contractavsservice, error) {
	contract, err := bindContractavsservice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractavsservice{ContractavsserviceCaller: ContractavsserviceCaller{contract: contract}, ContractavsserviceTransactor: ContractavsserviceTransactor{contract: contract}, ContractavsserviceFilterer: ContractavsserviceFilterer{contract: contract}}, nil
}

// NewContractavsserviceCaller creates a new read-only instance of Contractavsservice, bound to a specific deployed contract.
func NewContractavsserviceCaller(address common.Address, caller bind.ContractCaller) (*ContractavsserviceCaller, error) {
	contract, err := bindContractavsservice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractavsserviceCaller{contract: contract}, nil
}

// NewContractavsserviceTransactor creates a new write-only instance of Contractavsservice, bound to a specific deployed contract.
func NewContractavsserviceTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractavsserviceTransactor, error) {
	contract, err := bindContractavsservice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractavsserviceTransactor{contract: contract}, nil
}

// NewContractavsserviceFilterer creates a new log filterer instance of Contractavsservice, bound to a specific deployed contract.
func NewContractavsserviceFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractavsserviceFilterer, error) {
	contract, err := bindContractavsservice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractavsserviceFilterer{contract: contract}, nil
}

// bindContractavsservice binds a generic wrapper to an already deployed contract.
func bindContractavsservice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractavsserviceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractavsservice *ContractavsserviceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractavsservice.Contract.ContractavsserviceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractavsservice *ContractavsserviceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractavsservice.Contract.ContractavsserviceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractavsservice *ContractavsserviceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractavsservice.Contract.ContractavsserviceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractavsservice *ContractavsserviceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractavsservice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractavsservice *ContractavsserviceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractavsservice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractavsservice *ContractavsserviceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractavsservice.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractavsservice *ContractavsserviceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractavsservice *ContractavsserviceSession) Owner() (common.Address, error) {
	return _Contractavsservice.Contract.Owner(&_Contractavsservice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contractavsservice *ContractavsserviceCallerSession) Owner() (common.Address, error) {
	return _Contractavsservice.Contract.Owner(&_Contractavsservice.CallOpts)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xaf1991b3.
//
// Solidity: function createNewTask(string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) CreateNewTask(opts *bind.TransactOpts, name string, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint64, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "createNewTask", name, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xaf1991b3.
//
// Solidity: function createNewTask(string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(bool)
func (_Contractavsservice *ContractavsserviceSession) CreateNewTask(name string, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint64, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _Contractavsservice.Contract.CreateNewTask(&_Contractavsservice.TransactOpts, name, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xaf1991b3.
//
// Solidity: function createNewTask(string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) CreateNewTask(name string, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint64, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _Contractavsservice.Contract.CreateNewTask(&_Contractavsservice.TransactOpts, name, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "deregisterOperatorFromAVS")
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns(bool)
func (_Contractavsservice *ContractavsserviceSession) DeregisterOperatorFromAVS() (*types.Transaction, error) {
	return _Contractavsservice.Contract.DeregisterOperatorFromAVS(&_Contractavsservice.TransactOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) DeregisterOperatorFromAVS() (*types.Transaction, error) {
	return _Contractavsservice.Contract.DeregisterOperatorFromAVS(&_Contractavsservice.TransactOpts)
}

// GetOptInOperators is a paid mutator transaction binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) returns(string[])
func (_Contractavsservice *ContractavsserviceTransactor) GetOptInOperators(opts *bind.TransactOpts, avsAddress common.Address) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "getOptInOperators", avsAddress)
}

// GetOptInOperators is a paid mutator transaction binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) returns(string[])
func (_Contractavsservice *ContractavsserviceSession) GetOptInOperators(avsAddress common.Address) (*types.Transaction, error) {
	return _Contractavsservice.Contract.GetOptInOperators(&_Contractavsservice.TransactOpts, avsAddress)
}

// GetOptInOperators is a paid mutator transaction binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) returns(string[])
func (_Contractavsservice *ContractavsserviceTransactorSession) GetOptInOperators(avsAddress common.Address) (*types.Transaction, error) {
	return _Contractavsservice.Contract.GetOptInOperators(&_Contractavsservice.TransactOpts, avsAddress)
}

// GetRegisteredPubkey is a paid mutator transaction binding the contract method 0x55b42cbe.
//
// Solidity: function getRegisteredPubkey(string operator) returns(bytes)
func (_Contractavsservice *ContractavsserviceTransactor) GetRegisteredPubkey(opts *bind.TransactOpts, operator string) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "getRegisteredPubkey", operator)
}

// GetRegisteredPubkey is a paid mutator transaction binding the contract method 0x55b42cbe.
//
// Solidity: function getRegisteredPubkey(string operator) returns(bytes)
func (_Contractavsservice *ContractavsserviceSession) GetRegisteredPubkey(operator string) (*types.Transaction, error) {
	return _Contractavsservice.Contract.GetRegisteredPubkey(&_Contractavsservice.TransactOpts, operator)
}

// GetRegisteredPubkey is a paid mutator transaction binding the contract method 0x55b42cbe.
//
// Solidity: function getRegisteredPubkey(string operator) returns(bytes)
func (_Contractavsservice *ContractavsserviceTransactorSession) GetRegisteredPubkey(operator string) (*types.Transaction, error) {
	return _Contractavsservice.Contract.GetRegisteredPubkey(&_Contractavsservice.TransactOpts, operator)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0xc49bb521.
//
// Solidity: function registerAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) RegisterAVS(opts *bind.TransactOpts, avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "registerAVS", avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0xc49bb521.
//
// Solidity: function registerAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_Contractavsservice *ContractavsserviceSession) RegisterAVS(avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterAVS(&_Contractavsservice.TransactOpts, avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0xc49bb521.
//
// Solidity: function registerAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) RegisterAVS(avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterAVS(&_Contractavsservice.TransactOpts, avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x54c77f71.
//
// Solidity: function registerBLSPublicKey(string name, bytes pubKey, bytes pubKeyRegistrationSignature, bytes pubKeyRegistrationMessageHash) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, name string, pubKey []byte, pubKeyRegistrationSignature []byte, pubKeyRegistrationMessageHash []byte) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "registerBLSPublicKey", name, pubKey, pubKeyRegistrationSignature, pubKeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x54c77f71.
//
// Solidity: function registerBLSPublicKey(string name, bytes pubKey, bytes pubKeyRegistrationSignature, bytes pubKeyRegistrationMessageHash) returns(bool)
func (_Contractavsservice *ContractavsserviceSession) RegisterBLSPublicKey(name string, pubKey []byte, pubKeyRegistrationSignature []byte, pubKeyRegistrationMessageHash []byte) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterBLSPublicKey(&_Contractavsservice.TransactOpts, name, pubKey, pubKeyRegistrationSignature, pubKeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x54c77f71.
//
// Solidity: function registerBLSPublicKey(string name, bytes pubKey, bytes pubKeyRegistrationSignature, bytes pubKeyRegistrationMessageHash) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) RegisterBLSPublicKey(name string, pubKey []byte, pubKeyRegistrationSignature []byte, pubKeyRegistrationMessageHash []byte) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterBLSPublicKey(&_Contractavsservice.TransactOpts, name, pubKey, pubKeyRegistrationSignature, pubKeyRegistrationMessageHash)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "registerOperatorToAVS")
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns(bool)
func (_Contractavsservice *ContractavsserviceSession) RegisterOperatorToAVS() (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterOperatorToAVS(&_Contractavsservice.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) RegisterOperatorToAVS() (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterOperatorToAVS(&_Contractavsservice.TransactOpts)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x6f48e1a2.
//
// Solidity: function updateAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) UpdateAVS(opts *bind.TransactOpts, avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "updateAVS", avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x6f48e1a2.
//
// Solidity: function updateAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_Contractavsservice *ContractavsserviceSession) UpdateAVS(avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _Contractavsservice.Contract.UpdateAVS(&_Contractavsservice.TransactOpts, avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x6f48e1a2.
//
// Solidity: function updateAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) UpdateAVS(avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _Contractavsservice.Contract.UpdateAVS(&_Contractavsservice.TransactOpts, avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// ContractavsserviceDataLoggedIterator is returned from FilterDataLogged and is used to iterate over the raw logs and unpacked data for DataLogged events raised by the Contractavsservice contract.
type ContractavsserviceDataLoggedIterator struct {
	Event *ContractavsserviceDataLogged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractavsserviceDataLoggedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractavsserviceDataLogged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractavsserviceDataLogged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractavsserviceDataLoggedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractavsserviceDataLoggedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractavsserviceDataLogged represents a DataLogged event raised by the Contractavsservice contract.
type ContractavsserviceDataLogged struct {
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDataLogged is a free log retrieval operation binding the contract event 0xf8a35bbe260cb3c99c0febbe1d21b55e78ca43bdc6ed00e30007a9fadfe67e43.
//
// Solidity: event DataLogged(bytes data)
func (_Contractavsservice *ContractavsserviceFilterer) FilterDataLogged(opts *bind.FilterOpts) (*ContractavsserviceDataLoggedIterator, error) {

	logs, sub, err := _Contractavsservice.contract.FilterLogs(opts, "DataLogged")
	if err != nil {
		return nil, err
	}
	return &ContractavsserviceDataLoggedIterator{contract: _Contractavsservice.contract, event: "DataLogged", logs: logs, sub: sub}, nil
}

// WatchDataLogged is a free log subscription operation binding the contract event 0xf8a35bbe260cb3c99c0febbe1d21b55e78ca43bdc6ed00e30007a9fadfe67e43.
//
// Solidity: event DataLogged(bytes data)
func (_Contractavsservice *ContractavsserviceFilterer) WatchDataLogged(opts *bind.WatchOpts, sink chan<- *ContractavsserviceDataLogged) (event.Subscription, error) {

	logs, sub, err := _Contractavsservice.contract.WatchLogs(opts, "DataLogged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractavsserviceDataLogged)
				if err := _Contractavsservice.contract.UnpackLog(event, "DataLogged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDataLogged is a log parse operation binding the contract event 0xf8a35bbe260cb3c99c0febbe1d21b55e78ca43bdc6ed00e30007a9fadfe67e43.
//
// Solidity: event DataLogged(bytes data)
func (_Contractavsservice *ContractavsserviceFilterer) ParseDataLogged(log types.Log) (*ContractavsserviceDataLogged, error) {
	event := new(ContractavsserviceDataLogged)
	if err := _Contractavsservice.contract.UnpackLog(event, "DataLogged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractavsserviceTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the Contractavsservice contract.
type ContractavsserviceTaskCreatedIterator struct {
	Event *ContractavsserviceTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractavsserviceTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractavsserviceTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractavsserviceTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractavsserviceTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractavsserviceTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractavsserviceTaskCreated represents a TaskCreated event raised by the Contractavsservice contract.
type ContractavsserviceTaskCreated struct {
	TaskId                *big.Int
	Issuer                common.Address
	Name                  string
	TaskResponsePeriod    uint64
	TaskChallengePeriod   uint64
	ThresholdPercentage   uint64
	TaskStatisticalPeriod uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0xc3d2056aaa49a6e50ff7a353c7777b5cda982f2bf1e2214af811056cef07232d.
//
// Solidity: event TaskCreated(uint256 taskId, address issuer, string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod)
func (_Contractavsservice *ContractavsserviceFilterer) FilterTaskCreated(opts *bind.FilterOpts) (*ContractavsserviceTaskCreatedIterator, error) {

	logs, sub, err := _Contractavsservice.contract.FilterLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return &ContractavsserviceTaskCreatedIterator{contract: _Contractavsservice.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0xc3d2056aaa49a6e50ff7a353c7777b5cda982f2bf1e2214af811056cef07232d.
//
// Solidity: event TaskCreated(uint256 taskId, address issuer, string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod)
func (_Contractavsservice *ContractavsserviceFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractavsserviceTaskCreated) (event.Subscription, error) {

	logs, sub, err := _Contractavsservice.contract.WatchLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractavsserviceTaskCreated)
				if err := _Contractavsservice.contract.UnpackLog(event, "TaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCreated is a log parse operation binding the contract event 0xc3d2056aaa49a6e50ff7a353c7777b5cda982f2bf1e2214af811056cef07232d.
//
// Solidity: event TaskCreated(uint256 taskId, address issuer, string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod)
func (_Contractavsservice *ContractavsserviceFilterer) ParseTaskCreated(log types.Log) (*ContractavsserviceTaskCreated, error) {
	event := new(ContractavsserviceTaskCreated)
	if err := _Contractavsservice.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
