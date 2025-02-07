// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracthelloWorld

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

// AVSParams is an auto generated low-level Go binding around an user-defined struct.
type AVSParams struct {
	Sender              common.Address
	AvsName             string
	MinStakeAmount      uint64
	TaskAddr            common.Address
	SlashAddr           common.Address
	RewardAddr          common.Address
	AvsOwnerAddress     []common.Address
	WhitelistAddress    []common.Address
	AssetIds            []string
	AvsUnbondingPeriod  uint64
	MinSelfDelegation   uint64
	EpochIdentifier     string
	MiniOptInOperators  uint64
	MinTotalStakeAmount uint64
	AvsRewardProportion uint64
	AvsSlashProportion  uint64
}

// ContracthelloWorldMetaData contains all meta data concerning the ContracthelloWorld contract.
var ContracthelloWorldMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"name\":\"createNewTask\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregisterOperatorFromAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"}],\"name\":\"getAVSEpochIdentifier\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"}],\"name\":\"getAVSUSDValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"}],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operatorAddr\",\"type\":\"address\"}],\"name\":\"getOperatorOptedUSDValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddress\",\"type\":\"address\"}],\"name\":\"getOptInOperators\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"}],\"name\":\"getRegisteredPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"taskID\",\"type\":\"uint64\"}],\"name\":\"getTaskInfo\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"taskID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"taskResponse\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"}],\"name\":\"operatorSubmitTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"minStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"avsOwnerAddress\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"whitelistAddress\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"assetIds\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSelfDelegation\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"miniOptInOperators\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minTotalStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"avsRewardProportion\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"avsSlashProportion\",\"type\":\"uint64\"}],\"internalType\":\"structAVSParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"registerAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyRegistrationSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyRegistrationMessageHash\",\"type\":\"bytes\"}],\"name\":\"registerBLSPublicKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerOperatorToAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"minStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"avsOwnerAddress\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"whitelistAddress\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"assetIds\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSelfDelegation\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"miniOptInOperators\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minTotalStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"avsRewardProportion\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"avsSlashProportion\",\"type\":\"uint64\"}],\"internalType\":\"structAVSParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"updateAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50600080546001600160a01b03191633179055611780806100316000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80639943aa2711610097578063de16bf4611610066578063de16bf461461025e578063e093841414610266578063e2906f3d14610279578063fd6edb4e1461029957600080fd5b80639943aa27146101f8578063af1991b314610218578063c208dd9914610243578063dcf61b2c1461024b57600080fd5b80634d568f24116100d35780634d568f24146101735780636d70f7ae146101945780638da5cb5b146101a7578063992907fb146101d257600080fd5b80630b70f322146101055780631d4c80071461012d57806326135e9d1461014d5780633a72b90014610160575b600080fd5b6101186101133660046109d6565b6102ac565b60405190151581526020015b60405180910390f35b61014061013b366004610a2d565b610320565b6040516101249190610a98565b61011861015b366004610b68565b610394565b61011861016e3660046109d6565b61041c565b610186610181366004610c19565b610445565b604051908152602001610124565b6101186101a2366004610a2d565b6104c4565b6000546101ba906001600160a01b031681565b6040516001600160a01b039091168152602001610124565b6101e56101e0366004610d12565b61050f565b60405160079190910b8152602001610124565b61020b610206366004610c19565b61057a565b6040516101249190610d46565b61022b610226366004610d59565b6105f3565b6040516001600160401b039091168152602001610124565b610118610758565b610186610259366004610a2d565b6107c7565b610118610836565b61020b610274366004610a2d565b610860565b61028c610287366004610de2565b6108d4565b6040516101249190610e19565b6101186102a7366004610e65565b610957565b6040516305b8799160e11b8152600090819061090190630b70f322906102d6908690600401611068565b6020604051808303816000875af11580156102f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103199190611282565b9392505050565b604051631d4c800760e01b81526001600160a01b038216600482015260609060009061090190631d4c800790602401600060405180830381865afa15801561036c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103199190810190611317565b60405163046d13b160e11b81526000908190610901906308da2762906103cc9033908d908d908d908d908d908d908d906004016113cc565b6020604051808303816000875af11580156103eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061040f9190611282565b9998505050505050505050565b604051623a72b960e81b8152600090819061090190633a72b900906102d6908690600401611068565b604051631355a3c960e21b81526001600160a01b03808416600483015282166024820152600090819061090190634d568f2490604401602060405180830381865afa158015610498573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104bc919061143c565b949350505050565b6040516336b87bd760e11b81526001600160a01b0382166004820152600090819061090190636d70f7ae90602401602060405180830381865afa1580156102f5573d6000803e3d6000fd5b60405163992907fb60e01b815260009081906109019063992907fb90610539908690600401610d46565b602060405180830381865afa158015610556573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103199190611455565b604051639943aa2760e01b81526001600160a01b0380841660048301528216602482015260609061090190639943aa2790604401600060405180830381865afa1580156105cb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103199190810190611478565b6040805160c081018252600060208083018290528883526001600160401b0388811684860152878116606085015285811660a08501528616608084015292519092839161090191638bf30a699133918c91610650918891016114c0565b60408051601f19818403018152828252805160209182012090830152016040516020818303038152906040528b8b8b8b6040518863ffffffff1660e01b81526004016106a2979695949392919061153e565b6020604051808303816000875af11580156106c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e591906115a9565b6001600160401b0381166020840181905283516040808601516060870151608088015160a089015193519697507fc3d2056aaa49a6e50ff7a353c7777b5cda982f2bf1e2214af811056cef07232d966107459695339590949392916115c6565b60405180910390a1979650505050505050565b60405163d7a2398b60e01b815233600482015260009081906109019063d7a2398b906024015b6020604051808303816000875af115801561079d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c19190611282565b92915050565b60405163373d86cb60e21b81526001600160a01b038216600482015260009081906109019063dcf61b2c90602401602060405180830381865afa158015610812573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610319919061143c565b6040516351b27a6d60e11b815233600482015260009081906109019063a364f4da9060240161077e565b604051633824e10560e21b81526001600160a01b03821660048201526060906000906109019063e093841490602401600060405180830381865afa1580156108ac573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103199190810190611624565b60405163e2906f3d60e01b81526001600160a01b03831660048201526001600160401b03821660248201526060906000906109019063e2906f3d90604401600060405180830381865afa15801561092f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104bc9190810190611658565b6040516395af9dc760e01b81526000908190610901906395af9dc7906109899033908a908a908a908a906004016116ec565b6020604051808303816000875af11580156109a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109cc9190611282565b9695505050505050565b6000602082840312156109e857600080fd5b81356001600160401b038111156109fe57600080fd5b8201610200818503121561031957600080fd5b80356001600160a01b0381168114610a2857600080fd5b919050565b600060208284031215610a3f57600080fd5b61031982610a11565b60005b83811015610a63578181015183820152602001610a4b565b50506000910152565b60008151808452610a84816020860160208601610a48565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015610af157603f19878603018452610adc858351610a6c565b94506020938401939190910190600101610ac0565b50929695505050505050565b6001600160401b0381168114610b1257600080fd5b50565b8035610a2881610afd565b60008083601f840112610b3257600080fd5b5081356001600160401b03811115610b4957600080fd5b602083019150836020828501011115610b6157600080fd5b9250929050565b600080600080600080600060a0888a031215610b8357600080fd5b8735610b8e81610afd565b965060208801356001600160401b03811115610ba957600080fd5b610bb58a828b01610b20565b90975095505060408801356001600160401b03811115610bd457600080fd5b610be08a828b01610b20565b9095509350610bf3905060608901610a11565b9150608088013560ff81168114610c0957600080fd5b8091505092959891949750929550565b60008060408385031215610c2c57600080fd5b610c3583610a11565b9150610c4360208401610a11565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715610c8a57610c8a610c4c565b604052919050565b60006001600160401b03821115610cab57610cab610c4c565b50601f01601f191660200190565b600082601f830112610cca57600080fd5b8135602083016000610ce3610cde84610c92565b610c62565b9050828152858383011115610cf757600080fd5b82826020830137600092810160200192909252509392505050565b600060208284031215610d2457600080fd5b81356001600160401b03811115610d3a57600080fd5b6104bc84828501610cb9565b6020815260006103196020830184610a6c565b600080600080600060a08688031215610d7157600080fd5b85356001600160401b03811115610d8757600080fd5b610d9388828901610cb9565b9550506020860135610da481610afd565b93506040860135610db481610afd565b92506060860135610dc481610afd565b91506080860135610dd481610afd565b809150509295509295909350565b60008060408385031215610df557600080fd5b610dfe83610a11565b91506020830135610e0e81610afd565b809150509250929050565b602080825282518282018190526000918401906040840190835b81811015610e5a5783516001600160401b0316835260209384019390920191600101610e33565b509095945050505050565b60008060008060808587031215610e7b57600080fd5b610e8485610a11565b935060208501356001600160401b03811115610e9f57600080fd5b610eab87828801610cb9565b93505060408501356001600160401b03811115610ec757600080fd5b610ed387828801610cb9565b92505060608501356001600160401b03811115610eef57600080fd5b610efb87828801610cb9565b91505092959194509250565b6000808335601e19843603018112610f1e57600080fd5b83016020810192503590506001600160401b03811115610f3d57600080fd5b803603821315610b6157600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e19843603018112610f8c57600080fd5b83016020810192503590506001600160401b03811115610fab57600080fd5b8060051b3603821315610b6157600080fd5b81835260208301925060008160005b84811015610ffb576001600160a01b03610fe583610a11565b1686526020958601959190910190600101610fcc565b5093949350505050565b60008383855260208501945060208460051b8201018360005b8681101561105c57838303601f190188526110398287610f07565b611044858284610f4c565b60209a8b019a9095509390930192505060010161101e565b50909695505050505050565b602081526110896020820161107c84610a11565b6001600160a01b03169052565b60006110986020840184610f07565b61020060408501526110af61022085018284610f4c565b9150506110be60408501610b15565b6001600160401b0381166060850152506110da60608501610a11565b6001600160a01b0381166080850152506110f660808501610a11565b6001600160a01b03811660a08501525061111260a08501610a11565b6001600160a01b03811660c08501525061112f60c0850185610f75565b848303601f190160e0860152611146838284610fbd565b9250505061115760e0850185610f75565b848303601f190161010086015261116f838284610fbd565b92505050611181610100850185610f75565b848303601f1901610120860152611199838284611005565b925050506111aa6101208501610b15565b6001600160401b038116610140850152506111c86101408501610b15565b6001600160401b038116610160850152506111e7610160850185610f07565b848303601f19016101808601526111ff838284610f4c565b925050506112106101808501610b15565b6001600160401b0381166101a08501525061122e6101a08501610b15565b6001600160401b0381166101c08501525061124c6101c08501610b15565b6001600160401b0381166101e08501525061126a6101e08501610b15565b6001600160401b038116610200850152509392505050565b60006020828403121561129457600080fd5b8151801515811461031957600080fd5b60006001600160401b038211156112bd576112bd610c4c565b5060051b60200190565b60006112d5610cde84610c92565b90508281528383830111156112e957600080fd5b610319836020830184610a48565b600082601f83011261130857600080fd5b610319838351602085016112c7565b60006020828403121561132957600080fd5b81516001600160401b0381111561133f57600080fd5b8201601f8101841361135057600080fd5b805161135e610cde826112a4565b8082825260208201915060208360051b85010192508683111561138057600080fd5b602084015b838110156113c15780516001600160401b038111156113a357600080fd5b6113b2896020838901016112f7565b84525060209283019201611385565b509695505050505050565b6001600160a01b03891681526001600160401b038816602082015260c060408201819052600090611400908301888a610f4c565b8281036060840152611413818789610f4c565b6001600160a01b03959095166080840152505060ff9190911660a0909101529695505050505050565b60006020828403121561144e57600080fd5b5051919050565b60006020828403121561146757600080fd5b81518060070b811461031957600080fd5b60006020828403121561148a57600080fd5b81516001600160401b038111156114a057600080fd5b8201601f810184136114b157600080fd5b6104bc848251602084016112c7565b602081526000825160c060208401526114dc60e0840182610a6c565b9050602084015160408401526001600160401b0360408501511660608401526001600160401b0360608501511660808401526001600160401b0360808501511660a08401526001600160401b0360a08501511660c08401528091505092915050565b6001600160a01b038816815260e06020820181905260009061156290830189610a6c565b82810360408401526115748189610a6c565b6001600160401b0397881660608501529587166080840152505091841660a083015290921660c0909201919091529392505050565b6000602082840312156115bb57600080fd5b815161031981610afd565b8781526001600160a01b038716602082015260e0604082018190526000906115f090830188610a6c565b6001600160401b03968716606084015294861660808301525091841660a083015290921660c0909201919091529392505050565b60006020828403121561163657600080fd5b81516001600160401b0381111561164c57600080fd5b6104bc848285016112f7565b60006020828403121561166a57600080fd5b81516001600160401b0381111561168057600080fd5b8201601f8101841361169157600080fd5b805161169f610cde826112a4565b8082825260208201915060208360051b8501019250868311156116c157600080fd5b6020840193505b828410156109cc5783516116db81610afd565b8252602093840193909101906116c8565b6001600160a01b0386811682528516602082015260a06040820181905260009061171890830186610a6c565b828103606084015261172a8186610a6c565b9050828103608084015261173e8185610a6c565b9897505050505050505056fea26469706673582212201b0b5bc78ccd5f0ea3c0082b87735a3e9f617a50281cd9362eba8344736b5e7e64736f6c634300081c0033",
}

// ContracthelloWorldABI is the input ABI used to generate the binding from.
// Deprecated: Use ContracthelloWorldMetaData.ABI instead.
var ContracthelloWorldABI = ContracthelloWorldMetaData.ABI

// ContracthelloWorldBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContracthelloWorldMetaData.Bin instead.
var ContracthelloWorldBin = ContracthelloWorldMetaData.Bin

// DeployContracthelloWorld deploys a new Ethereum contract, binding an instance of ContracthelloWorld to it.
func DeployContracthelloWorld(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContracthelloWorld, error) {
	parsed, err := ContracthelloWorldMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContracthelloWorldBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContracthelloWorld{ContracthelloWorldCaller: ContracthelloWorldCaller{contract: contract}, ContracthelloWorldTransactor: ContracthelloWorldTransactor{contract: contract}, ContracthelloWorldFilterer: ContracthelloWorldFilterer{contract: contract}}, nil
}

// ContracthelloWorld is an auto generated Go binding around an Ethereum contract.
type ContracthelloWorld struct {
	ContracthelloWorldCaller     // Read-only binding to the contract
	ContracthelloWorldTransactor // Write-only binding to the contract
	ContracthelloWorldFilterer   // Log filterer for contract events
}

// ContracthelloWorldCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContracthelloWorldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContracthelloWorldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContracthelloWorldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContracthelloWorldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContracthelloWorldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContracthelloWorldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContracthelloWorldSession struct {
	Contract     *ContracthelloWorld // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContracthelloWorldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContracthelloWorldCallerSession struct {
	Contract *ContracthelloWorldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContracthelloWorldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContracthelloWorldTransactorSession struct {
	Contract     *ContracthelloWorldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContracthelloWorldRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContracthelloWorldRaw struct {
	Contract *ContracthelloWorld // Generic contract binding to access the raw methods on
}

// ContracthelloWorldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContracthelloWorldCallerRaw struct {
	Contract *ContracthelloWorldCaller // Generic read-only contract binding to access the raw methods on
}

// ContracthelloWorldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContracthelloWorldTransactorRaw struct {
	Contract *ContracthelloWorldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracthelloWorld creates a new instance of ContracthelloWorld, bound to a specific deployed contract.
func NewContracthelloWorld(address common.Address, backend bind.ContractBackend) (*ContracthelloWorld, error) {
	contract, err := bindContracthelloWorld(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContracthelloWorld{ContracthelloWorldCaller: ContracthelloWorldCaller{contract: contract}, ContracthelloWorldTransactor: ContracthelloWorldTransactor{contract: contract}, ContracthelloWorldFilterer: ContracthelloWorldFilterer{contract: contract}}, nil
}

// NewContracthelloWorldCaller creates a new read-only instance of ContracthelloWorld, bound to a specific deployed contract.
func NewContracthelloWorldCaller(address common.Address, caller bind.ContractCaller) (*ContracthelloWorldCaller, error) {
	contract, err := bindContracthelloWorld(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContracthelloWorldCaller{contract: contract}, nil
}

// NewContracthelloWorldTransactor creates a new write-only instance of ContracthelloWorld, bound to a specific deployed contract.
func NewContracthelloWorldTransactor(address common.Address, transactor bind.ContractTransactor) (*ContracthelloWorldTransactor, error) {
	contract, err := bindContracthelloWorld(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContracthelloWorldTransactor{contract: contract}, nil
}

// NewContracthelloWorldFilterer creates a new log filterer instance of ContracthelloWorld, bound to a specific deployed contract.
func NewContracthelloWorldFilterer(address common.Address, filterer bind.ContractFilterer) (*ContracthelloWorldFilterer, error) {
	contract, err := bindContracthelloWorld(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContracthelloWorldFilterer{contract: contract}, nil
}

// bindContracthelloWorld binds a generic wrapper to an already deployed contract.
func bindContracthelloWorld(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContracthelloWorldMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContracthelloWorld *ContracthelloWorldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContracthelloWorld.Contract.ContracthelloWorldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContracthelloWorld *ContracthelloWorldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.ContracthelloWorldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContracthelloWorld *ContracthelloWorldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.ContracthelloWorldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContracthelloWorld *ContracthelloWorldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContracthelloWorld.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContracthelloWorld *ContracthelloWorldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContracthelloWorld *ContracthelloWorldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.contract.Transact(opts, method, params...)
}

// GetAVSEpochIdentifier is a free data retrieval call binding the contract method 0xe0938414.
//
// Solidity: function getAVSEpochIdentifier(address avsAddr) view returns(string)
func (_ContracthelloWorld *ContracthelloWorldCaller) GetAVSEpochIdentifier(opts *bind.CallOpts, avsAddr common.Address) (string, error) {
	var out []interface{}
	err := _ContracthelloWorld.contract.Call(opts, &out, "getAVSEpochIdentifier", avsAddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAVSEpochIdentifier is a free data retrieval call binding the contract method 0xe0938414.
//
// Solidity: function getAVSEpochIdentifier(address avsAddr) view returns(string)
func (_ContracthelloWorld *ContracthelloWorldSession) GetAVSEpochIdentifier(avsAddr common.Address) (string, error) {
	return _ContracthelloWorld.Contract.GetAVSEpochIdentifier(&_ContracthelloWorld.CallOpts, avsAddr)
}

// GetAVSEpochIdentifier is a free data retrieval call binding the contract method 0xe0938414.
//
// Solidity: function getAVSEpochIdentifier(address avsAddr) view returns(string)
func (_ContracthelloWorld *ContracthelloWorldCallerSession) GetAVSEpochIdentifier(avsAddr common.Address) (string, error) {
	return _ContracthelloWorld.Contract.GetAVSEpochIdentifier(&_ContracthelloWorld.CallOpts, avsAddr)
}

// GetAVSUSDValue is a free data retrieval call binding the contract method 0xdcf61b2c.
//
// Solidity: function getAVSUSDValue(address avsAddr) view returns(uint256)
func (_ContracthelloWorld *ContracthelloWorldCaller) GetAVSUSDValue(opts *bind.CallOpts, avsAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContracthelloWorld.contract.Call(opts, &out, "getAVSUSDValue", avsAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAVSUSDValue is a free data retrieval call binding the contract method 0xdcf61b2c.
//
// Solidity: function getAVSUSDValue(address avsAddr) view returns(uint256)
func (_ContracthelloWorld *ContracthelloWorldSession) GetAVSUSDValue(avsAddr common.Address) (*big.Int, error) {
	return _ContracthelloWorld.Contract.GetAVSUSDValue(&_ContracthelloWorld.CallOpts, avsAddr)
}

// GetAVSUSDValue is a free data retrieval call binding the contract method 0xdcf61b2c.
//
// Solidity: function getAVSUSDValue(address avsAddr) view returns(uint256)
func (_ContracthelloWorld *ContracthelloWorldCallerSession) GetAVSUSDValue(avsAddr common.Address) (*big.Int, error) {
	return _ContracthelloWorld.Contract.GetAVSUSDValue(&_ContracthelloWorld.CallOpts, avsAddr)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0x992907fb.
//
// Solidity: function getCurrentEpoch(string epochIdentifier) view returns(int64)
func (_ContracthelloWorld *ContracthelloWorldCaller) GetCurrentEpoch(opts *bind.CallOpts, epochIdentifier string) (int64, error) {
	var out []interface{}
	err := _ContracthelloWorld.contract.Call(opts, &out, "getCurrentEpoch", epochIdentifier)

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0x992907fb.
//
// Solidity: function getCurrentEpoch(string epochIdentifier) view returns(int64)
func (_ContracthelloWorld *ContracthelloWorldSession) GetCurrentEpoch(epochIdentifier string) (int64, error) {
	return _ContracthelloWorld.Contract.GetCurrentEpoch(&_ContracthelloWorld.CallOpts, epochIdentifier)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0x992907fb.
//
// Solidity: function getCurrentEpoch(string epochIdentifier) view returns(int64)
func (_ContracthelloWorld *ContracthelloWorldCallerSession) GetCurrentEpoch(epochIdentifier string) (int64, error) {
	return _ContracthelloWorld.Contract.GetCurrentEpoch(&_ContracthelloWorld.CallOpts, epochIdentifier)
}

// GetOperatorOptedUSDValue is a free data retrieval call binding the contract method 0x4d568f24.
//
// Solidity: function getOperatorOptedUSDValue(address avsAddr, address operatorAddr) view returns(uint256)
func (_ContracthelloWorld *ContracthelloWorldCaller) GetOperatorOptedUSDValue(opts *bind.CallOpts, avsAddr common.Address, operatorAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContracthelloWorld.contract.Call(opts, &out, "getOperatorOptedUSDValue", avsAddr, operatorAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorOptedUSDValue is a free data retrieval call binding the contract method 0x4d568f24.
//
// Solidity: function getOperatorOptedUSDValue(address avsAddr, address operatorAddr) view returns(uint256)
func (_ContracthelloWorld *ContracthelloWorldSession) GetOperatorOptedUSDValue(avsAddr common.Address, operatorAddr common.Address) (*big.Int, error) {
	return _ContracthelloWorld.Contract.GetOperatorOptedUSDValue(&_ContracthelloWorld.CallOpts, avsAddr, operatorAddr)
}

// GetOperatorOptedUSDValue is a free data retrieval call binding the contract method 0x4d568f24.
//
// Solidity: function getOperatorOptedUSDValue(address avsAddr, address operatorAddr) view returns(uint256)
func (_ContracthelloWorld *ContracthelloWorldCallerSession) GetOperatorOptedUSDValue(avsAddr common.Address, operatorAddr common.Address) (*big.Int, error) {
	return _ContracthelloWorld.Contract.GetOperatorOptedUSDValue(&_ContracthelloWorld.CallOpts, avsAddr, operatorAddr)
}

// GetOptInOperators is a free data retrieval call binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) view returns(string[])
func (_ContracthelloWorld *ContracthelloWorldCaller) GetOptInOperators(opts *bind.CallOpts, avsAddress common.Address) ([]string, error) {
	var out []interface{}
	err := _ContracthelloWorld.contract.Call(opts, &out, "getOptInOperators", avsAddress)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetOptInOperators is a free data retrieval call binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) view returns(string[])
func (_ContracthelloWorld *ContracthelloWorldSession) GetOptInOperators(avsAddress common.Address) ([]string, error) {
	return _ContracthelloWorld.Contract.GetOptInOperators(&_ContracthelloWorld.CallOpts, avsAddress)
}

// GetOptInOperators is a free data retrieval call binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) view returns(string[])
func (_ContracthelloWorld *ContracthelloWorldCallerSession) GetOptInOperators(avsAddress common.Address) ([]string, error) {
	return _ContracthelloWorld.Contract.GetOptInOperators(&_ContracthelloWorld.CallOpts, avsAddress)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x9943aa27.
//
// Solidity: function getRegisteredPubkey(address operator, address avsAddr) view returns(bytes)
func (_ContracthelloWorld *ContracthelloWorldCaller) GetRegisteredPubkey(opts *bind.CallOpts, operator common.Address, avsAddr common.Address) ([]byte, error) {
	var out []interface{}
	err := _ContracthelloWorld.contract.Call(opts, &out, "getRegisteredPubkey", operator, avsAddr)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x9943aa27.
//
// Solidity: function getRegisteredPubkey(address operator, address avsAddr) view returns(bytes)
func (_ContracthelloWorld *ContracthelloWorldSession) GetRegisteredPubkey(operator common.Address, avsAddr common.Address) ([]byte, error) {
	return _ContracthelloWorld.Contract.GetRegisteredPubkey(&_ContracthelloWorld.CallOpts, operator, avsAddr)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x9943aa27.
//
// Solidity: function getRegisteredPubkey(address operator, address avsAddr) view returns(bytes)
func (_ContracthelloWorld *ContracthelloWorldCallerSession) GetRegisteredPubkey(operator common.Address, avsAddr common.Address) ([]byte, error) {
	return _ContracthelloWorld.Contract.GetRegisteredPubkey(&_ContracthelloWorld.CallOpts, operator, avsAddr)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0xe2906f3d.
//
// Solidity: function getTaskInfo(address taskAddr, uint64 taskID) view returns(uint64[])
func (_ContracthelloWorld *ContracthelloWorldCaller) GetTaskInfo(opts *bind.CallOpts, taskAddr common.Address, taskID uint64) ([]uint64, error) {
	var out []interface{}
	err := _ContracthelloWorld.contract.Call(opts, &out, "getTaskInfo", taskAddr, taskID)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetTaskInfo is a free data retrieval call binding the contract method 0xe2906f3d.
//
// Solidity: function getTaskInfo(address taskAddr, uint64 taskID) view returns(uint64[])
func (_ContracthelloWorld *ContracthelloWorldSession) GetTaskInfo(taskAddr common.Address, taskID uint64) ([]uint64, error) {
	return _ContracthelloWorld.Contract.GetTaskInfo(&_ContracthelloWorld.CallOpts, taskAddr, taskID)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0xe2906f3d.
//
// Solidity: function getTaskInfo(address taskAddr, uint64 taskID) view returns(uint64[])
func (_ContracthelloWorld *ContracthelloWorldCallerSession) GetTaskInfo(taskAddr common.Address, taskID uint64) ([]uint64, error) {
	return _ContracthelloWorld.Contract.GetTaskInfo(&_ContracthelloWorld.CallOpts, taskAddr, taskID)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_ContracthelloWorld *ContracthelloWorldCaller) IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ContracthelloWorld.contract.Call(opts, &out, "isOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_ContracthelloWorld *ContracthelloWorldSession) IsOperator(operator common.Address) (bool, error) {
	return _ContracthelloWorld.Contract.IsOperator(&_ContracthelloWorld.CallOpts, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_ContracthelloWorld *ContracthelloWorldCallerSession) IsOperator(operator common.Address) (bool, error) {
	return _ContracthelloWorld.Contract.IsOperator(&_ContracthelloWorld.CallOpts, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContracthelloWorld *ContracthelloWorldCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContracthelloWorld.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContracthelloWorld *ContracthelloWorldSession) Owner() (common.Address, error) {
	return _ContracthelloWorld.Contract.Owner(&_ContracthelloWorld.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContracthelloWorld *ContracthelloWorldCallerSession) Owner() (common.Address, error) {
	return _ContracthelloWorld.Contract.Owner(&_ContracthelloWorld.CallOpts)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xaf1991b3.
//
// Solidity: function createNewTask(string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(uint64)
func (_ContracthelloWorld *ContracthelloWorldTransactor) CreateNewTask(opts *bind.TransactOpts, name string, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint64, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _ContracthelloWorld.contract.Transact(opts, "createNewTask", name, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xaf1991b3.
//
// Solidity: function createNewTask(string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(uint64)
func (_ContracthelloWorld *ContracthelloWorldSession) CreateNewTask(name string, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint64, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.CreateNewTask(&_ContracthelloWorld.TransactOpts, name, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xaf1991b3.
//
// Solidity: function createNewTask(string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(uint64)
func (_ContracthelloWorld *ContracthelloWorldTransactorSession) CreateNewTask(name string, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint64, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.CreateNewTask(&_ContracthelloWorld.TransactOpts, name, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContracthelloWorld.contract.Transact(opts, "deregisterOperatorFromAVS")
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns(bool)
func (_ContracthelloWorld *ContracthelloWorldSession) DeregisterOperatorFromAVS() (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.DeregisterOperatorFromAVS(&_ContracthelloWorld.TransactOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactorSession) DeregisterOperatorFromAVS() (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.DeregisterOperatorFromAVS(&_ContracthelloWorld.TransactOpts)
}

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0x26135e9d.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, address taskContractAddress, uint8 phase) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactor) OperatorSubmitTask(opts *bind.TransactOpts, taskID uint64, taskResponse []byte, blsSignature []byte, taskContractAddress common.Address, phase uint8) (*types.Transaction, error) {
	return _ContracthelloWorld.contract.Transact(opts, "operatorSubmitTask", taskID, taskResponse, blsSignature, taskContractAddress, phase)
}

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0x26135e9d.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, address taskContractAddress, uint8 phase) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldSession) OperatorSubmitTask(taskID uint64, taskResponse []byte, blsSignature []byte, taskContractAddress common.Address, phase uint8) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.OperatorSubmitTask(&_ContracthelloWorld.TransactOpts, taskID, taskResponse, blsSignature, taskContractAddress, phase)
}

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0x26135e9d.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, address taskContractAddress, uint8 phase) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactorSession) OperatorSubmitTask(taskID uint64, taskResponse []byte, blsSignature []byte, taskContractAddress common.Address, phase uint8) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.OperatorSubmitTask(&_ContracthelloWorld.TransactOpts, taskID, taskResponse, blsSignature, taskContractAddress, phase)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0x0b70f322.
//
// Solidity: function registerAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactor) RegisterAVS(opts *bind.TransactOpts, params AVSParams) (*types.Transaction, error) {
	return _ContracthelloWorld.contract.Transact(opts, "registerAVS", params)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0x0b70f322.
//
// Solidity: function registerAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldSession) RegisterAVS(params AVSParams) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.RegisterAVS(&_ContracthelloWorld.TransactOpts, params)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0x0b70f322.
//
// Solidity: function registerAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactorSession) RegisterAVS(params AVSParams) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.RegisterAVS(&_ContracthelloWorld.TransactOpts, params)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xfd6edb4e.
//
// Solidity: function registerBLSPublicKey(address avsAddr, bytes pubKey, bytes pubKeyRegistrationSignature, bytes pubKeyRegistrationMessageHash) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, avsAddr common.Address, pubKey []byte, pubKeyRegistrationSignature []byte, pubKeyRegistrationMessageHash []byte) (*types.Transaction, error) {
	return _ContracthelloWorld.contract.Transact(opts, "registerBLSPublicKey", avsAddr, pubKey, pubKeyRegistrationSignature, pubKeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xfd6edb4e.
//
// Solidity: function registerBLSPublicKey(address avsAddr, bytes pubKey, bytes pubKeyRegistrationSignature, bytes pubKeyRegistrationMessageHash) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldSession) RegisterBLSPublicKey(avsAddr common.Address, pubKey []byte, pubKeyRegistrationSignature []byte, pubKeyRegistrationMessageHash []byte) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.RegisterBLSPublicKey(&_ContracthelloWorld.TransactOpts, avsAddr, pubKey, pubKeyRegistrationSignature, pubKeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xfd6edb4e.
//
// Solidity: function registerBLSPublicKey(address avsAddr, bytes pubKey, bytes pubKeyRegistrationSignature, bytes pubKeyRegistrationMessageHash) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactorSession) RegisterBLSPublicKey(avsAddr common.Address, pubKey []byte, pubKeyRegistrationSignature []byte, pubKeyRegistrationMessageHash []byte) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.RegisterBLSPublicKey(&_ContracthelloWorld.TransactOpts, avsAddr, pubKey, pubKeyRegistrationSignature, pubKeyRegistrationMessageHash)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContracthelloWorld.contract.Transact(opts, "registerOperatorToAVS")
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns(bool)
func (_ContracthelloWorld *ContracthelloWorldSession) RegisterOperatorToAVS() (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.RegisterOperatorToAVS(&_ContracthelloWorld.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactorSession) RegisterOperatorToAVS() (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.RegisterOperatorToAVS(&_ContracthelloWorld.TransactOpts)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x3a72b900.
//
// Solidity: function updateAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactor) UpdateAVS(opts *bind.TransactOpts, params AVSParams) (*types.Transaction, error) {
	return _ContracthelloWorld.contract.Transact(opts, "updateAVS", params)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x3a72b900.
//
// Solidity: function updateAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldSession) UpdateAVS(params AVSParams) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.UpdateAVS(&_ContracthelloWorld.TransactOpts, params)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x3a72b900.
//
// Solidity: function updateAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactorSession) UpdateAVS(params AVSParams) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.UpdateAVS(&_ContracthelloWorld.TransactOpts, params)
}

// ContracthelloWorldTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the ContracthelloWorld contract.
type ContracthelloWorldTaskCreatedIterator struct {
	Event *ContracthelloWorldTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContracthelloWorldTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContracthelloWorldTaskCreated)
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
		it.Event = new(ContracthelloWorldTaskCreated)
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
func (it *ContracthelloWorldTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContracthelloWorldTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContracthelloWorldTaskCreated represents a TaskCreated event raised by the ContracthelloWorld contract.
type ContracthelloWorldTaskCreated struct {
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
func (_ContracthelloWorld *ContracthelloWorldFilterer) FilterTaskCreated(opts *bind.FilterOpts) (*ContracthelloWorldTaskCreatedIterator, error) {

	logs, sub, err := _ContracthelloWorld.contract.FilterLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return &ContracthelloWorldTaskCreatedIterator{contract: _ContracthelloWorld.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0xc3d2056aaa49a6e50ff7a353c7777b5cda982f2bf1e2214af811056cef07232d.
//
// Solidity: event TaskCreated(uint256 taskId, address issuer, string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod)
func (_ContracthelloWorld *ContracthelloWorldFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *ContracthelloWorldTaskCreated) (event.Subscription, error) {

	logs, sub, err := _ContracthelloWorld.contract.WatchLogs(opts, "TaskCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContracthelloWorldTaskCreated)
				if err := _ContracthelloWorld.contract.UnpackLog(event, "TaskCreated", log); err != nil {
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
func (_ContracthelloWorld *ContracthelloWorldFilterer) ParseTaskCreated(log types.Log) (*ContracthelloWorldTaskCreated, error) {
	event := new(ContracthelloWorldTaskCreated)
	if err := _ContracthelloWorld.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
