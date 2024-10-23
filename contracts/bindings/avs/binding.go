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

// ContracthelloWorldMetaData contains all meta data concerning the ContracthelloWorld contract.
var ContracthelloWorldMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"name\":\"createNewTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregisterOperatorFromAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"}],\"name\":\"getAVSInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"}],\"name\":\"getAVSUSDValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"}],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"operatorAddr\",\"type\":\"string\"}],\"name\":\"getOperatorOptedUSDValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddress\",\"type\":\"address\"}],\"name\":\"getOptInOperators\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"}],\"name\":\"getRegisteredPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"taskID\",\"type\":\"uint64\"}],\"name\":\"getTaskInfo\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"taskID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"taskResponse\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"stage\",\"type\":\"string\"}],\"name\":\"operatorSubmitTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"minStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"avsOwnerAddress\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"assetIds\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSelfDelegation\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64[]\",\"name\":\"params\",\"type\":\"uint64[]\"}],\"name\":\"registerAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyRegistrationSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyRegistrationMessageHash\",\"type\":\"bytes\"}],\"name\":\"registerBLSPublicKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerOperatorToAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metaInfo\",\"type\":\"string\"}],\"name\":\"registerOperatorToExocore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"minStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"avsOwnerAddress\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"assetIds\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSelfDelegation\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64[]\",\"name\":\"params\",\"type\":\"uint64[]\"}],\"name\":\"updateAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b505f80546001600160a01b031916331790556117748061002d5f395ff3fe608060405234801561000f575f80fd5b5060043610610106575f3560e01c8063992907fb1161009e578063c49bb5211161006e578063c49bb5211461024a578063dcf61b2c1461025d578063de16bf4614610270578063e2906f3d14610278578063f329827314610298575f80fd5b8063992907fb146101f6578063af1991b31461021c578063b32f34b31461022f578063c208dd9914610242575f80fd5b806354c77f71116100d957806354c77f711461018657806355b42cbe146101995780636f48e1a2146101b95780638da5cb5b146101cc575f80fd5b80631c275bf41461010a5780631d4c8007146101325780631f041bd1146101525780632d9d6a2014610165575b5f80fd5b61011d610118366004610b13565b6102ab565b60405190151581526020015b60405180910390f35b610145610140366004610b5f565b61031d565b6040516101299190610bfe565b61011d610160366004610b13565b61038c565b610178610173366004610c10565b6103d0565b604051908152602001610129565b61011d610194366004610c5a565b610442565b6101ac6101a7366004610b13565b6104bd565b6040516101299190610cff565b61011d6101c7366004610e3f565b61053e565b5f546101de906001600160a01b031681565b6040516001600160a01b039091168152602001610129565b610209610204366004610b13565b6105d5565b60405160079190910b8152602001610129565b61011d61022a366004610f70565b61063d565b6101ac61023d366004610b5f565b6107a0565b61011d61080f565b61011d610258366004610e3f565b61087a565b61017861026b366004610b5f565b6108c0565b61011d61092c565b61028b610286366004610ff4565b610955565b604051610129919061106c565b61011d6102a63660046110c2565b6109d3565b604051632b03f8ab60e11b81525f90819061090190635607f156906102d69033908790600401611173565b6020604051808303815f875af11580156102f2573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061031691906111a5565b9392505050565b604051631d4c800760e01b81526001600160a01b03821660048201526060905f9061090190631d4c8007906024015f60405180830381865afa158015610365573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526103169190810190611212565b604051631f041bd160e01b81525f90819061090190631f041bd1906103b5908690600401610cff565b602060405180830381865afa1580156102f2573d5f803e3d5ffd5b60405163016ceb5160e51b81525f90819061090190632d9d6a20906103fb9087908790600401611173565b602060405180830381865afa158015610416573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061043a91906112bb565b949350505050565b604051632a95fddf60e21b81525f9081906109019063aa57f77c906104739033908a908a908a908a906004016112d2565b6020604051808303815f875af115801561048f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104b391906111a5565b9695505050505050565b604051632ada165f60e11b81526060905f90610901906355b42cbe906104e7908690600401610cff565b5f60405180830381865afa158015610501573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610528919081019061133b565b905080806020019051810190610316919061133b565b5f806109016001600160a01b031663d9e5daa0338f8f8f8f8f8f8f8f8f8f8f6040518d63ffffffff1660e01b81526004016105849c9b9a9998979695949392919061137f565b6020604051808303815f875af11580156105a0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105c491906111a5565b9d9c50505050505050505050505050565b60405163992907fb60e01b81525f9081906109019063992907fb906105fe908690600401610cff565b602060405180830381865afa158015610619573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103169190611464565b6040805160c0810182525f60208083018290528883526001600160401b0388811684860152878116606085015285811660a085015286166080840152925190928391829161090191638bf30a699133918d9161069b91899101611484565b60408051601f19818403018152828252805160209182012090830152016040516020818303038152906040528c8c8c8c6040518863ffffffff1660e01b81526004016106ed97969594939291906114f0565b60408051808303815f875af1158015610708573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061072c919061155a565b6001600160401b038116602086018190528551604080880151606089015160808a015160a08b015193519799509597507fc3d2056aaa49a6e50ff7a353c7777b5cda982f2bf1e2214af811056cef07232d9661078c963395949091611584565b60405180910390a150979650505050505050565b60405163b32f34b360e01b81526001600160a01b03821660048201526060905f906109019063b32f34b3906024015f60405180830381865afa1580156107e8573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261031691908101906115e1565b60405163d7a2398b60e01b81523360048201525f9081906109019063d7a2398b906024015b6020604051808303815f875af1158015610850573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061087491906111a5565b92915050565b5f806109016001600160a01b031663cde09950338f8f8f8f8f8f8f8f8f8f8f6040518d63ffffffff1660e01b81526004016105849c9b9a9998979695949392919061137f565b60405163373d86cb60e21b81526001600160a01b03821660048201525f9081906109019063dcf61b2c90602401602060405180830381865afa158015610908573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061031691906112bb565b6040516351b27a6d60e11b81523360048201525f9081906109019063a364f4da90602401610834565b60405163e2906f3d60e01b81526001600160a01b03831660048201526001600160401b03821660248201526060905f906109019063e2906f3d906044015f60405180830381865afa1580156109ac573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261043a9190810190611612565b60405163394e3d6360e21b81525f9081906109019063e538f58c90610a0a9033908d908d908d908d908d908d908d906004016116ce565b6020604051808303815f875af1158015610a26573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a4a91906111a5565b9998505050505050505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715610a9357610a93610a57565b604052919050565b5f6001600160401b03821115610ab357610ab3610a57565b50601f01601f191660200190565b5f82601f830112610ad0575f80fd5b8135610ae3610ade82610a9b565b610a6b565b818152846020838601011115610af7575f80fd5b816020850160208301375f918101602001919091529392505050565b5f60208284031215610b23575f80fd5b81356001600160401b03811115610b38575f80fd5b61043a84828501610ac1565b80356001600160a01b0381168114610b5a575f80fd5b919050565b5f60208284031215610b6f575f80fd5b61031682610b44565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f8282518085526020808601955060208260051b840101602086015f5b84811015610bf157601f19868403018952610bdf838351610b78565b98840198925090830190600101610bc3565b5090979650505050505050565b602081525f6103166020830184610ba6565b5f8060408385031215610c21575f80fd5b610c2a83610b44565b915060208301356001600160401b03811115610c44575f80fd5b610c5085828601610ac1565b9150509250929050565b5f805f8060808587031215610c6d575f80fd5b84356001600160401b0380821115610c83575f80fd5b610c8f88838901610ac1565b95506020870135915080821115610ca4575f80fd5b610cb088838901610ac1565b94506040870135915080821115610cc5575f80fd5b610cd188838901610ac1565b93506060870135915080821115610ce6575f80fd5b50610cf387828801610ac1565b91505092959194509250565b602081525f6103166020830184610b78565b6001600160401b0381168114610d25575f80fd5b50565b8035610b5a81610d11565b5f6001600160401b03821115610d4b57610d4b610a57565b5060051b60200190565b5f82601f830112610d64575f80fd5b81356020610d74610ade83610d33565b82815260059290921b84018101918181019086841115610d92575f80fd5b8286015b84811015610dcf5780356001600160401b03811115610db3575f80fd5b610dc18986838b0101610ac1565b845250918301918301610d96565b509695505050505050565b5f82601f830112610de9575f80fd5b81356020610df9610ade83610d33565b8083825260208201915060208460051b870101935086841115610e1a575f80fd5b602086015b84811015610dcf578035610e3281610d11565b8352918301918301610e1f565b5f805f805f805f805f805f6101608c8e031215610e5a575f80fd5b6001600160401b03808d351115610e6f575f80fd5b610e7c8e8e358f01610ac1565b9b50610e8a60208e01610d28565b9a50610e9860408e01610b44565b9950610ea660608e01610b44565b9850610eb460808e01610b44565b97508060a08e01351115610ec6575f80fd5b610ed68e60a08f01358f01610d55565b96508060c08e01351115610ee8575f80fd5b610ef88e60c08f01358f01610d55565b9550610f0660e08e01610d28565b9450610f156101008e01610d28565b9350806101208e01351115610f28575f80fd5b610f398e6101208f01358f01610ac1565b9250806101408e01351115610f4c575f80fd5b50610f5e8d6101408e01358e01610dda565b90509295989b509295989b9093969950565b5f805f805f60a08688031215610f84575f80fd5b85356001600160401b03811115610f99575f80fd5b610fa588828901610ac1565b9550506020860135610fb681610d11565b93506040860135610fc681610d11565b92506060860135610fd681610d11565b91506080860135610fe681610d11565b809150509295509295909350565b5f8060408385031215611005575f80fd5b61100e83610b44565b9150602083013561101e81610d11565b809150509250929050565b5f815180845260208085019450602084015f5b838110156110615781516001600160401b03168752958201959082019060010161103c565b509495945050505050565b602081525f6103166020830184611029565b5f8083601f84011261108e575f80fd5b5081356001600160401b038111156110a4575f80fd5b6020830191508360208285010111156110bb575f80fd5b9250929050565b5f805f805f805f60a0888a0312156110d8575f80fd5b87356110e381610d11565b965060208801356001600160401b03808211156110fe575f80fd5b61110a8b838c0161107e565b909850965060408a0135915080821115611122575f80fd5b61112e8b838c0161107e565b909650945084915061114260608b01610b44565b935060808a0135915080821115611157575f80fd5b506111648a828b01610ac1565b91505092959891949750929550565b6001600160a01b03831681526040602082018190525f9061043a90830184610b78565b80518015158114610b5a575f80fd5b5f602082840312156111b5575f80fd5b61031682611196565b5f6111cb610ade84610a9b565b90508281528383830111156111de575f80fd5b8282602083015e5f602084830101529392505050565b5f82601f830112611203575f80fd5b610316838351602085016111be565b5f6020808385031215611223575f80fd5b82516001600160401b0380821115611239575f80fd5b818501915085601f83011261124c575f80fd5b815161125a610ade82610d33565b81815260059190911b83018401908481019088831115611278575f80fd5b8585015b838110156112ae57805185811115611292575f80fd5b6112a08b89838a01016111f4565b84525091860191860161127c565b5098975050505050505050565b5f602082840312156112cb575f80fd5b5051919050565b6001600160a01b038616815260a0602082018190525f906112f590830187610b78565b82810360408401526113078187610b78565b9050828103606084015261131b8186610b78565b9050828103608084015261132f8185610b78565b98975050505050505050565b5f6020828403121561134b575f80fd5b81516001600160401b03811115611360575f80fd5b8201601f81018413611370575f80fd5b61043a848251602084016111be565b6001600160a01b038d168152610180602082018190525f906113a38382018f610b78565b6001600160401b038e16604085015290506001600160a01b038c1660608401526001600160a01b038b1660808401526001600160a01b038a1660a084015282810360c08401526113f3818a610ba6565b905082810360e08401526114078189610ba6565b6001600160401b03881661010085015290506001600160401b03861661012084015282810361014084015261143c8186610b78565b90508281036101608401526114518185611029565b9f9e505050505050505050505050505050565b5f60208284031215611474575f80fd5b81518060070b8114610316575f80fd5b602081525f825160c0602084015261149f60e0840182610b78565b90506020840151604084015260408401516001600160401b0380821660608601528060608701511660808601528060808701511660a08601528060a08701511660c086015250508091505092915050565b6001600160a01b038816815260e0602082018190525f9061151390830189610b78565b82810360408401526115258189610b78565b6001600160401b0397881660608501529587166080840152505091841660a083015290921660c0909201919091529392505050565b5f806040838503121561156b575f80fd5b61157483611196565b9150602083015161101e81610d11565b8781526001600160a01b038716602082015260e0604082018190525f906115ad90830188610b78565b6001600160401b03968716606084015294861660808301525091841660a083015290921660c0909201919091529392505050565b5f602082840312156115f1575f80fd5b81516001600160401b03811115611606575f80fd5b61043a848285016111f4565b5f6020808385031215611623575f80fd5b82516001600160401b03811115611638575f80fd5b8301601f81018513611648575f80fd5b8051611656610ade82610d33565b81815260059190911b82018301908381019087831115611674575f80fd5b928401925b8284101561169b57835161168c81610d11565b82529284019290840190611679565b979650505050505050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b5f60018060a01b03808b1683526001600160401b038a16602084015260c060408401526116ff60c08401898b6116a6565b838103606085015261171281888a6116a6565b9050818616608085015283810360a085015261172e8186610b78565b9c9b50505050505050505050505056fea26469706673582212209a992f7deb5659b1b8f61727e9b4ad2efedbd6d5bbd79c1e3cd9a9a4cdf4df0f64736f6c63430008190033",
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

// GetAVSInfo is a free data retrieval call binding the contract method 0xb32f34b3.
//
// Solidity: function getAVSInfo(address avsAddr) view returns(string)
func (_ContracthelloWorld *ContracthelloWorldCaller) GetAVSInfo(opts *bind.CallOpts, avsAddr common.Address) (string, error) {
	var out []interface{}
	err := _ContracthelloWorld.contract.Call(opts, &out, "getAVSInfo", avsAddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAVSInfo is a free data retrieval call binding the contract method 0xb32f34b3.
//
// Solidity: function getAVSInfo(address avsAddr) view returns(string)
func (_ContracthelloWorld *ContracthelloWorldSession) GetAVSInfo(avsAddr common.Address) (string, error) {
	return _ContracthelloWorld.Contract.GetAVSInfo(&_ContracthelloWorld.CallOpts, avsAddr)
}

// GetAVSInfo is a free data retrieval call binding the contract method 0xb32f34b3.
//
// Solidity: function getAVSInfo(address avsAddr) view returns(string)
func (_ContracthelloWorld *ContracthelloWorldCallerSession) GetAVSInfo(avsAddr common.Address) (string, error) {
	return _ContracthelloWorld.Contract.GetAVSInfo(&_ContracthelloWorld.CallOpts, avsAddr)
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

// GetOperatorOptedUSDValue is a free data retrieval call binding the contract method 0x2d9d6a20.
//
// Solidity: function getOperatorOptedUSDValue(address avsAddr, string operatorAddr) view returns(uint256)
func (_ContracthelloWorld *ContracthelloWorldCaller) GetOperatorOptedUSDValue(opts *bind.CallOpts, avsAddr common.Address, operatorAddr string) (*big.Int, error) {
	var out []interface{}
	err := _ContracthelloWorld.contract.Call(opts, &out, "getOperatorOptedUSDValue", avsAddr, operatorAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorOptedUSDValue is a free data retrieval call binding the contract method 0x2d9d6a20.
//
// Solidity: function getOperatorOptedUSDValue(address avsAddr, string operatorAddr) view returns(uint256)
func (_ContracthelloWorld *ContracthelloWorldSession) GetOperatorOptedUSDValue(avsAddr common.Address, operatorAddr string) (*big.Int, error) {
	return _ContracthelloWorld.Contract.GetOperatorOptedUSDValue(&_ContracthelloWorld.CallOpts, avsAddr, operatorAddr)
}

// GetOperatorOptedUSDValue is a free data retrieval call binding the contract method 0x2d9d6a20.
//
// Solidity: function getOperatorOptedUSDValue(address avsAddr, string operatorAddr) view returns(uint256)
func (_ContracthelloWorld *ContracthelloWorldCallerSession) GetOperatorOptedUSDValue(avsAddr common.Address, operatorAddr string) (*big.Int, error) {
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

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x55b42cbe.
//
// Solidity: function getRegisteredPubkey(string operator) view returns(bytes)
func (_ContracthelloWorld *ContracthelloWorldCaller) GetRegisteredPubkey(opts *bind.CallOpts, operator string) ([]byte, error) {
	var out []interface{}
	err := _ContracthelloWorld.contract.Call(opts, &out, "getRegisteredPubkey", operator)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x55b42cbe.
//
// Solidity: function getRegisteredPubkey(string operator) view returns(bytes)
func (_ContracthelloWorld *ContracthelloWorldSession) GetRegisteredPubkey(operator string) ([]byte, error) {
	return _ContracthelloWorld.Contract.GetRegisteredPubkey(&_ContracthelloWorld.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x55b42cbe.
//
// Solidity: function getRegisteredPubkey(string operator) view returns(bytes)
func (_ContracthelloWorld *ContracthelloWorldCallerSession) GetRegisteredPubkey(operator string) ([]byte, error) {
	return _ContracthelloWorld.Contract.GetRegisteredPubkey(&_ContracthelloWorld.CallOpts, operator)
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

// IsOperator is a free data retrieval call binding the contract method 0x1f041bd1.
//
// Solidity: function isOperator(string operator) view returns(bool)
func (_ContracthelloWorld *ContracthelloWorldCaller) IsOperator(opts *bind.CallOpts, operator string) (bool, error) {
	var out []interface{}
	err := _ContracthelloWorld.contract.Call(opts, &out, "isOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x1f041bd1.
//
// Solidity: function isOperator(string operator) view returns(bool)
func (_ContracthelloWorld *ContracthelloWorldSession) IsOperator(operator string) (bool, error) {
	return _ContracthelloWorld.Contract.IsOperator(&_ContracthelloWorld.CallOpts, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x1f041bd1.
//
// Solidity: function isOperator(string operator) view returns(bool)
func (_ContracthelloWorld *ContracthelloWorldCallerSession) IsOperator(operator string) (bool, error) {
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
// Solidity: function createNewTask(string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactor) CreateNewTask(opts *bind.TransactOpts, name string, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint64, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _ContracthelloWorld.contract.Transact(opts, "createNewTask", name, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xaf1991b3.
//
// Solidity: function createNewTask(string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldSession) CreateNewTask(name string, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint64, taskStatisticalPeriod uint64) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.CreateNewTask(&_ContracthelloWorld.TransactOpts, name, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xaf1991b3.
//
// Solidity: function createNewTask(string name, uint64 taskResponsePeriod, uint64 taskChallengePeriod, uint64 thresholdPercentage, uint64 taskStatisticalPeriod) returns(bool)
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

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0xf3298273.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, address taskContractAddress, string stage) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactor) OperatorSubmitTask(opts *bind.TransactOpts, taskID uint64, taskResponse []byte, blsSignature []byte, taskContractAddress common.Address, stage string) (*types.Transaction, error) {
	return _ContracthelloWorld.contract.Transact(opts, "operatorSubmitTask", taskID, taskResponse, blsSignature, taskContractAddress, stage)
}

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0xf3298273.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, address taskContractAddress, string stage) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldSession) OperatorSubmitTask(taskID uint64, taskResponse []byte, blsSignature []byte, taskContractAddress common.Address, stage string) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.OperatorSubmitTask(&_ContracthelloWorld.TransactOpts, taskID, taskResponse, blsSignature, taskContractAddress, stage)
}

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0xf3298273.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, address taskContractAddress, string stage) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactorSession) OperatorSubmitTask(taskID uint64, taskResponse []byte, blsSignature []byte, taskContractAddress common.Address, stage string) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.OperatorSubmitTask(&_ContracthelloWorld.TransactOpts, taskID, taskResponse, blsSignature, taskContractAddress, stage)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0xc49bb521.
//
// Solidity: function registerAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactor) RegisterAVS(opts *bind.TransactOpts, avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _ContracthelloWorld.contract.Transact(opts, "registerAVS", avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0xc49bb521.
//
// Solidity: function registerAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldSession) RegisterAVS(avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.RegisterAVS(&_ContracthelloWorld.TransactOpts, avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0xc49bb521.
//
// Solidity: function registerAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactorSession) RegisterAVS(avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.RegisterAVS(&_ContracthelloWorld.TransactOpts, avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x54c77f71.
//
// Solidity: function registerBLSPublicKey(string name, bytes pubKey, bytes pubKeyRegistrationSignature, bytes pubKeyRegistrationMessageHash) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, name string, pubKey []byte, pubKeyRegistrationSignature []byte, pubKeyRegistrationMessageHash []byte) (*types.Transaction, error) {
	return _ContracthelloWorld.contract.Transact(opts, "registerBLSPublicKey", name, pubKey, pubKeyRegistrationSignature, pubKeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x54c77f71.
//
// Solidity: function registerBLSPublicKey(string name, bytes pubKey, bytes pubKeyRegistrationSignature, bytes pubKeyRegistrationMessageHash) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldSession) RegisterBLSPublicKey(name string, pubKey []byte, pubKeyRegistrationSignature []byte, pubKeyRegistrationMessageHash []byte) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.RegisterBLSPublicKey(&_ContracthelloWorld.TransactOpts, name, pubKey, pubKeyRegistrationSignature, pubKeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x54c77f71.
//
// Solidity: function registerBLSPublicKey(string name, bytes pubKey, bytes pubKeyRegistrationSignature, bytes pubKeyRegistrationMessageHash) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactorSession) RegisterBLSPublicKey(name string, pubKey []byte, pubKeyRegistrationSignature []byte, pubKeyRegistrationMessageHash []byte) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.RegisterBLSPublicKey(&_ContracthelloWorld.TransactOpts, name, pubKey, pubKeyRegistrationSignature, pubKeyRegistrationMessageHash)
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

// RegisterOperatorToExocore is a paid mutator transaction binding the contract method 0x1c275bf4.
//
// Solidity: function registerOperatorToExocore(string metaInfo) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactor) RegisterOperatorToExocore(opts *bind.TransactOpts, metaInfo string) (*types.Transaction, error) {
	return _ContracthelloWorld.contract.Transact(opts, "registerOperatorToExocore", metaInfo)
}

// RegisterOperatorToExocore is a paid mutator transaction binding the contract method 0x1c275bf4.
//
// Solidity: function registerOperatorToExocore(string metaInfo) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldSession) RegisterOperatorToExocore(metaInfo string) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.RegisterOperatorToExocore(&_ContracthelloWorld.TransactOpts, metaInfo)
}

// RegisterOperatorToExocore is a paid mutator transaction binding the contract method 0x1c275bf4.
//
// Solidity: function registerOperatorToExocore(string metaInfo) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactorSession) RegisterOperatorToExocore(metaInfo string) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.RegisterOperatorToExocore(&_ContracthelloWorld.TransactOpts, metaInfo)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x6f48e1a2.
//
// Solidity: function updateAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactor) UpdateAVS(opts *bind.TransactOpts, avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _ContracthelloWorld.contract.Transact(opts, "updateAVS", avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x6f48e1a2.
//
// Solidity: function updateAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldSession) UpdateAVS(avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.UpdateAVS(&_ContracthelloWorld.TransactOpts, avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x6f48e1a2.
//
// Solidity: function updateAVS(string avsName, uint64 minStakeAmount, address taskAddr, address slashAddr, address rewardAddr, string[] avsOwnerAddress, string[] assetIds, uint64 avsUnbondingPeriod, uint64 minSelfDelegation, string epochIdentifier, uint64[] params) returns(bool)
func (_ContracthelloWorld *ContracthelloWorldTransactorSession) UpdateAVS(avsName string, minStakeAmount uint64, taskAddr common.Address, slashAddr common.Address, rewardAddr common.Address, avsOwnerAddress []string, assetIds []string, avsUnbondingPeriod uint64, minSelfDelegation uint64, epochIdentifier string, params []uint64) (*types.Transaction, error) {
	return _ContracthelloWorld.Contract.UpdateAVS(&_ContracthelloWorld.TransactOpts, avsName, minStakeAmount, taskAddr, slashAddr, rewardAddr, avsOwnerAddress, assetIds, avsUnbondingPeriod, minSelfDelegation, epochIdentifier, params)
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
