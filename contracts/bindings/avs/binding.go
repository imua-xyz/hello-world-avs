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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"name\":\"createNewTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregisterOperatorFromAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"}],\"name\":\"getAVSInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"}],\"name\":\"getAVSUSDValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"operatorAddr\",\"type\":\"string\"}],\"name\":\"getOperatorOptedUSDValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddress\",\"type\":\"address\"}],\"name\":\"getOptInOperators\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"}],\"name\":\"getRegisteredPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"taskID\",\"type\":\"uint64\"}],\"name\":\"getTaskInfo\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"taskID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"taskResponse\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"stage\",\"type\":\"string\"}],\"name\":\"operatorSubmitTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"minStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"avsOwnerAddress\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"assetIds\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSelfDelegation\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64[]\",\"name\":\"params\",\"type\":\"uint64[]\"}],\"name\":\"registerAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyRegistrationSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyRegistrationMessageHash\",\"type\":\"bytes\"}],\"name\":\"registerBLSPublicKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerOperatorToAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metaInfo\",\"type\":\"string\"}],\"name\":\"registerOperatorToExocore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"minStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"avsOwnerAddress\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"assetIds\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSelfDelegation\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64[]\",\"name\":\"params\",\"type\":\"uint64[]\"}],\"name\":\"updateAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506124698061005f6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063af1991b311610097578063dcf61b2c11610066578063dcf61b2c14610321578063de16bf4614610351578063e2906f3d1461036f578063f32982731461039f57610100565b8063af1991b314610273578063b32f34b3146102a3578063c208dd99146102d3578063c49bb521146102f157610100565b806354c77f71116100d357806354c77f71146101c557806355b42cbe146101f55780636f48e1a2146102255780638da5cb5b1461025557610100565b80631c275bf4146101055780631d4c8007146101355780631f041bd1146101655780632d9d6a2014610195575b600080fd5b61011f600480360381019061011a9190610fb9565b6103cf565b60405161012c919061101d565b60405180910390f35b61014f600480360381019061014a9190611096565b61045d565b60405161015c9190611204565b60405180910390f35b61017f600480360381019061017a9190610fb9565b6104ed565b60405161018c919061101d565b60405180910390f35b6101af60048036038101906101aa9190611226565b610577565b6040516101bc919061129b565b60405180910390f35b6101df60048036038101906101da9190611357565b610604565b6040516101ec919061101d565b60405180910390f35b61020f600480360381019061020a9190610fb9565b61069b565b60405161021c9190611483565b60405180910390f35b61023f600480360381019061023a919061168e565b61073e565b60405161024c919061101d565b60405180910390f35b61025d6107ea565b60405161026a919061181d565b60405180910390f35b61028d60048036038101906102889190611838565b61080e565b60405161029a919061101d565b60405180910390f35b6102bd60048036038101906102b89190611096565b6109f4565b6040516102ca9190611919565b60405180910390f35b6102db610a84565b6040516102e8919061101d565b60405180910390f35b61030b6004803603810190610306919061168e565b610b0e565b604051610318919061101d565b60405180910390f35b61033b60048036038101906103369190611096565b610bba565b604051610348919061129b565b60405180910390f35b610359610c44565b604051610366919061101d565b60405180910390f35b6103896004803603810190610384919061193b565b610cce565b6040516103969190611a39565b60405180910390f35b6103b960048036038101906103b49190611ab6565b610d61565b6040516103c6919061101d565b60405180910390f35b60008061090173ffffffffffffffffffffffffffffffffffffffff16635607f15633856040518363ffffffff1660e01b815260040161040f929190611b8e565b6020604051808303816000875af115801561042e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104529190611bea565b905080915050919050565b6060600061090173ffffffffffffffffffffffffffffffffffffffff16631d4c8007846040518263ffffffff1660e01b815260040161049c919061181d565b600060405180830381865afa1580156104b9573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906104e29190611d3c565b905080915050919050565b60008061090173ffffffffffffffffffffffffffffffffffffffff16631f041bd1846040518263ffffffff1660e01b815260040161052b9190611919565b602060405180830381865afa158015610548573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056c9190611bea565b905080915050919050565b60008061090173ffffffffffffffffffffffffffffffffffffffff16632d9d6a2085856040518363ffffffff1660e01b81526004016105b7929190611b8e565b602060405180830381865afa1580156105d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105f89190611db1565b90508091505092915050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663aa57f77c33888888886040518663ffffffff1660e01b815260040161064a959493929190611dde565b6020604051808303816000875af1158015610669573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061068d9190611bea565b905080915050949350505050565b6060600061090173ffffffffffffffffffffffffffffffffffffffff166355b42cbe846040518263ffffffff1660e01b81526004016106da9190611919565b600060405180830381865afa1580156106f7573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906107209190611ebd565b9050808060200190518101906107369190611ebd565b915050919050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663d9e5daa0338f8f8f8f8f8f8f8f8f8f8f6040518d63ffffffff1660e01b81526004016107929c9b9a99989796959493929190611f15565b6020604051808303816000875af11580156107b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d59190611bea565b9050809150509b9a5050505050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000610818610e01565b86816000018190525085816040019067ffffffffffffffff16908167ffffffffffffffff168152505084816060019067ffffffffffffffff16908167ffffffffffffffff1681525050828160a0019067ffffffffffffffff16908167ffffffffffffffff168152505083816080019067ffffffffffffffff16908167ffffffffffffffff168152505060008061090173ffffffffffffffffffffffffffffffffffffffff16638bf30a69338b866040516020016108d5919061208a565b604051602081830303815290604052805190602001206040516020016108fb91906120d7565b6040516020818303038152906040528c8c8c8c6040518863ffffffff1660e01b815260040161093097969594939291906120f2565b60408051808303816000875af115801561094e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109729190612184565b915091508067ffffffffffffffff168360200181815250507fc3d2056aaa49a6e50ff7a353c7777b5cda982f2bf1e2214af811056cef07232d83602001513385600001518660400151876060015188608001518960a001516040516109dd97969594939291906121c4565b60405180910390a181935050505095945050505050565b6060600061090173ffffffffffffffffffffffffffffffffffffffff1663b32f34b3846040518263ffffffff1660e01b8152600401610a33919061181d565b600060405180830381865afa158015610a50573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610a79919061223a565b905080915050919050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663d7a2398b336040518263ffffffff1660e01b8152600401610ac2919061181d565b6020604051808303816000875af1158015610ae1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b059190611bea565b90508091505090565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663cde09950338f8f8f8f8f8f8f8f8f8f8f6040518d63ffffffff1660e01b8152600401610b629c9b9a99989796959493929190611f15565b6020604051808303816000875af1158015610b81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba59190611bea565b9050809150509b9a5050505050505050505050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663dcf61b2c846040518263ffffffff1660e01b8152600401610bf8919061181d565b602060405180830381865afa158015610c15573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c399190611db1565b905080915050919050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663a364f4da336040518263ffffffff1660e01b8152600401610c82919061181d565b6020604051808303816000875af1158015610ca1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc59190611bea565b90508091505090565b6060600061090173ffffffffffffffffffffffffffffffffffffffff1663e2906f3d85856040518363ffffffff1660e01b8152600401610d0f929190612283565b600060405180830381865afa158015610d2c573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610d559190612343565b90508091505092915050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663e538f58c338b8b8b8b8b8b8b6040518963ffffffff1660e01b8152600401610dad9897969594939291906123b9565b6020604051808303816000875af1158015610dcc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df09190611bea565b905080915050979650505050505050565b6040518060c001604052806060815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610ec682610e7d565b810181811067ffffffffffffffff82111715610ee557610ee4610e8e565b5b80604052505050565b6000610ef8610e5f565b9050610f048282610ebd565b919050565b600067ffffffffffffffff821115610f2457610f23610e8e565b5b610f2d82610e7d565b9050602081019050919050565b82818337600083830152505050565b6000610f5c610f5784610f09565b610eee565b905082815260208101848484011115610f7857610f77610e78565b5b610f83848285610f3a565b509392505050565b600082601f830112610fa057610f9f610e73565b5b8135610fb0848260208601610f49565b91505092915050565b600060208284031215610fcf57610fce610e69565b5b600082013567ffffffffffffffff811115610fed57610fec610e6e565b5b610ff984828501610f8b565b91505092915050565b60008115159050919050565b61101781611002565b82525050565b6000602082019050611032600083018461100e565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061106382611038565b9050919050565b61107381611058565b811461107e57600080fd5b50565b6000813590506110908161106a565b92915050565b6000602082840312156110ac576110ab610e69565b5b60006110ba84828501611081565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561112957808201518184015260208101905061110e565b60008484015250505050565b6000611140826110ef565b61114a81856110fa565b935061115a81856020860161110b565b61116381610e7d565b840191505092915050565b600061117a8383611135565b905092915050565b6000602082019050919050565b600061119a826110c3565b6111a481856110ce565b9350836020820285016111b6856110df565b8060005b858110156111f257848403895281516111d3858261116e565b94506111de83611182565b925060208a019950506001810190506111ba565b50829750879550505050505092915050565b6000602082019050818103600083015261121e818461118f565b905092915050565b6000806040838503121561123d5761123c610e69565b5b600061124b85828601611081565b925050602083013567ffffffffffffffff81111561126c5761126b610e6e565b5b61127885828601610f8b565b9150509250929050565b6000819050919050565b61129581611282565b82525050565b60006020820190506112b0600083018461128c565b92915050565b600067ffffffffffffffff8211156112d1576112d0610e8e565b5b6112da82610e7d565b9050602081019050919050565b60006112fa6112f5846112b6565b610eee565b90508281526020810184848401111561131657611315610e78565b5b611321848285610f3a565b509392505050565b600082601f83011261133e5761133d610e73565b5b813561134e8482602086016112e7565b91505092915050565b6000806000806080858703121561137157611370610e69565b5b600085013567ffffffffffffffff81111561138f5761138e610e6e565b5b61139b87828801610f8b565b945050602085013567ffffffffffffffff8111156113bc576113bb610e6e565b5b6113c887828801611329565b935050604085013567ffffffffffffffff8111156113e9576113e8610e6e565b5b6113f587828801611329565b925050606085013567ffffffffffffffff81111561141657611415610e6e565b5b61142287828801611329565b91505092959194509250565b600081519050919050565b600082825260208201905092915050565b60006114558261142e565b61145f8185611439565b935061146f81856020860161110b565b61147881610e7d565b840191505092915050565b6000602082019050818103600083015261149d818461144a565b905092915050565b600067ffffffffffffffff82169050919050565b6114c2816114a5565b81146114cd57600080fd5b50565b6000813590506114df816114b9565b92915050565b600067ffffffffffffffff821115611500576114ff610e8e565b5b602082029050602081019050919050565b600080fd5b6000611529611524846114e5565b610eee565b9050808382526020820190506020840283018581111561154c5761154b611511565b5b835b8181101561159357803567ffffffffffffffff81111561157157611570610e73565b5b80860161157e8982610f8b565b8552602085019450505060208101905061154e565b5050509392505050565b600082601f8301126115b2576115b1610e73565b5b81356115c2848260208601611516565b91505092915050565b600067ffffffffffffffff8211156115e6576115e5610e8e565b5b602082029050602081019050919050565b600061160a611605846115cb565b610eee565b9050808382526020820190506020840283018581111561162d5761162c611511565b5b835b81811015611656578061164288826114d0565b84526020840193505060208101905061162f565b5050509392505050565b600082601f83011261167557611674610e73565b5b81356116858482602086016115f7565b91505092915050565b60008060008060008060008060008060006101608c8e0312156116b4576116b3610e69565b5b60008c013567ffffffffffffffff8111156116d2576116d1610e6e565b5b6116de8e828f01610f8b565b9b505060206116ef8e828f016114d0565b9a505060406117008e828f01611081565b99505060606117118e828f01611081565b98505060806117228e828f01611081565b97505060a08c013567ffffffffffffffff81111561174357611742610e6e565b5b61174f8e828f0161159d565b96505060c08c013567ffffffffffffffff8111156117705761176f610e6e565b5b61177c8e828f0161159d565b95505060e061178d8e828f016114d0565b94505061010061179f8e828f016114d0565b9350506101208c013567ffffffffffffffff8111156117c1576117c0610e6e565b5b6117cd8e828f01610f8b565b9250506101408c013567ffffffffffffffff8111156117ef576117ee610e6e565b5b6117fb8e828f01611660565b9150509295989b509295989b9093969950565b61181781611058565b82525050565b6000602082019050611832600083018461180e565b92915050565b600080600080600060a0868803121561185457611853610e69565b5b600086013567ffffffffffffffff81111561187257611871610e6e565b5b61187e88828901610f8b565b955050602061188f888289016114d0565b94505060406118a0888289016114d0565b93505060606118b1888289016114d0565b92505060806118c2888289016114d0565b9150509295509295909350565b600082825260208201905092915050565b60006118eb826110ef565b6118f581856118cf565b935061190581856020860161110b565b61190e81610e7d565b840191505092915050565b6000602082019050818103600083015261193381846118e0565b905092915050565b6000806040838503121561195257611951610e69565b5b600061196085828601611081565b9250506020611971858286016114d0565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6119b0816114a5565b82525050565b60006119c283836119a7565b60208301905092915050565b6000602082019050919050565b60006119e68261197b565b6119f08185611986565b93506119fb83611997565b8060005b83811015611a2c578151611a1388826119b6565b9750611a1e836119ce565b9250506001810190506119ff565b5085935050505092915050565b60006020820190508181036000830152611a5381846119db565b905092915050565b600080fd5b60008083601f840112611a7657611a75610e73565b5b8235905067ffffffffffffffff811115611a9357611a92611a5b565b5b602083019150836001820283011115611aaf57611aae611511565b5b9250929050565b600080600080600080600060a0888a031215611ad557611ad4610e69565b5b6000611ae38a828b016114d0565b975050602088013567ffffffffffffffff811115611b0457611b03610e6e565b5b611b108a828b01611a60565b9650965050604088013567ffffffffffffffff811115611b3357611b32610e6e565b5b611b3f8a828b01611a60565b94509450506060611b528a828b01611081565b925050608088013567ffffffffffffffff811115611b7357611b72610e6e565b5b611b7f8a828b01610f8b565b91505092959891949750929550565b6000604082019050611ba3600083018561180e565b8181036020830152611bb581846118e0565b90509392505050565b611bc781611002565b8114611bd257600080fd5b50565b600081519050611be481611bbe565b92915050565b600060208284031215611c0057611bff610e69565b5b6000611c0e84828501611bd5565b91505092915050565b6000611c2a611c2584610f09565b610eee565b905082815260208101848484011115611c4657611c45610e78565b5b611c5184828561110b565b509392505050565b600082601f830112611c6e57611c6d610e73565b5b8151611c7e848260208601611c17565b91505092915050565b6000611c9a611c95846114e5565b610eee565b90508083825260208201905060208402830185811115611cbd57611cbc611511565b5b835b81811015611d0457805167ffffffffffffffff811115611ce257611ce1610e73565b5b808601611cef8982611c59565b85526020850194505050602081019050611cbf565b5050509392505050565b600082601f830112611d2357611d22610e73565b5b8151611d33848260208601611c87565b91505092915050565b600060208284031215611d5257611d51610e69565b5b600082015167ffffffffffffffff811115611d7057611d6f610e6e565b5b611d7c84828501611d0e565b91505092915050565b611d8e81611282565b8114611d9957600080fd5b50565b600081519050611dab81611d85565b92915050565b600060208284031215611dc757611dc6610e69565b5b6000611dd584828501611d9c565b91505092915050565b600060a082019050611df3600083018861180e565b8181036020830152611e0581876118e0565b90508181036040830152611e19818661144a565b90508181036060830152611e2d818561144a565b90508181036080830152611e41818461144a565b90509695505050505050565b6000611e60611e5b846112b6565b610eee565b905082815260208101848484011115611e7c57611e7b610e78565b5b611e8784828561110b565b509392505050565b600082601f830112611ea457611ea3610e73565b5b8151611eb4848260208601611e4d565b91505092915050565b600060208284031215611ed357611ed2610e69565b5b600082015167ffffffffffffffff811115611ef157611ef0610e6e565b5b611efd84828501611e8f565b91505092915050565b611f0f816114a5565b82525050565b600061018082019050611f2b600083018f61180e565b8181036020830152611f3d818e6118e0565b9050611f4c604083018d611f06565b611f59606083018c61180e565b611f66608083018b61180e565b611f7360a083018a61180e565b81810360c0830152611f85818961118f565b905081810360e0830152611f99818861118f565b9050611fa9610100830187611f06565b611fb7610120830186611f06565b818103610140830152611fca81856118e0565b9050818103610160830152611fdf81846119db565b90509d9c50505050505050505050505050565b611ffb81611282565b82525050565b600060c083016000830151848203600086015261201e8282611135565b91505060208301516120336020860182611ff2565b50604083015161204660408601826119a7565b50606083015161205960608601826119a7565b50608083015161206c60808601826119a7565b5060a083015161207f60a08601826119a7565b508091505092915050565b600060208201905081810360008301526120a48184612001565b905092915050565b6000819050919050565b6000819050919050565b6120d16120cc826120ac565b6120b6565b82525050565b60006120e382846120c0565b60208201915081905092915050565b600060e082019050612107600083018a61180e565b818103602083015261211981896118e0565b9050818103604083015261212d818861144a565b905061213c6060830187611f06565b6121496080830186611f06565b61215660a0830185611f06565b61216360c0830184611f06565b98975050505050505050565b60008151905061217e816114b9565b92915050565b6000806040838503121561219b5761219a610e69565b5b60006121a985828601611bd5565b92505060206121ba8582860161216f565b9150509250929050565b600060e0820190506121d9600083018a61128c565b6121e6602083018961180e565b81810360408301526121f881886118e0565b90506122076060830187611f06565b6122146080830186611f06565b61222160a0830185611f06565b61222e60c0830184611f06565b98975050505050505050565b6000602082840312156122505761224f610e69565b5b600082015167ffffffffffffffff81111561226e5761226d610e6e565b5b61227a84828501611c59565b91505092915050565b6000604082019050612298600083018561180e565b6122a56020830184611f06565b9392505050565b60006122bf6122ba846115cb565b610eee565b905080838252602082019050602084028301858111156122e2576122e1611511565b5b835b8181101561230b57806122f7888261216f565b8452602084019350506020810190506122e4565b5050509392505050565b600082601f83011261232a57612329610e73565b5b815161233a8482602086016122ac565b91505092915050565b60006020828403121561235957612358610e69565b5b600082015167ffffffffffffffff81111561237757612376610e6e565b5b61238384828501612315565b91505092915050565b60006123988385611439565b93506123a5838584610f3a565b6123ae83610e7d565b840190509392505050565b600060c0820190506123ce600083018b61180e565b6123db602083018a611f06565b81810360408301526123ee81888a61238c565b9050818103606083015261240381868861238c565b9050612412608083018561180e565b81810360a083015261242481846118e0565b9050999850505050505050505056fea2646970667358221220caa793214054491a8f87d518443a1ecf8d12a4c090a222d70198850aa1e54c3164736f6c63430008190033",
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

// GetAVSInfo is a free data retrieval call binding the contract method 0xb32f34b3.
//
// Solidity: function getAVSInfo(address avsAddr) view returns(string)
func (_Contractavsservice *ContractavsserviceCaller) GetAVSInfo(opts *bind.CallOpts, avsAddr common.Address) (string, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "getAVSInfo", avsAddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAVSInfo is a free data retrieval call binding the contract method 0xb32f34b3.
//
// Solidity: function getAVSInfo(address avsAddr) view returns(string)
func (_Contractavsservice *ContractavsserviceSession) GetAVSInfo(avsAddr common.Address) (string, error) {
	return _Contractavsservice.Contract.GetAVSInfo(&_Contractavsservice.CallOpts, avsAddr)
}

// GetAVSInfo is a free data retrieval call binding the contract method 0xb32f34b3.
//
// Solidity: function getAVSInfo(address avsAddr) view returns(string)
func (_Contractavsservice *ContractavsserviceCallerSession) GetAVSInfo(avsAddr common.Address) (string, error) {
	return _Contractavsservice.Contract.GetAVSInfo(&_Contractavsservice.CallOpts, avsAddr)
}

// GetAVSUSDValue is a free data retrieval call binding the contract method 0xdcf61b2c.
//
// Solidity: function getAVSUSDValue(address avsAddr) view returns(uint256)
func (_Contractavsservice *ContractavsserviceCaller) GetAVSUSDValue(opts *bind.CallOpts, avsAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "getAVSUSDValue", avsAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAVSUSDValue is a free data retrieval call binding the contract method 0xdcf61b2c.
//
// Solidity: function getAVSUSDValue(address avsAddr) view returns(uint256)
func (_Contractavsservice *ContractavsserviceSession) GetAVSUSDValue(avsAddr common.Address) (*big.Int, error) {
	return _Contractavsservice.Contract.GetAVSUSDValue(&_Contractavsservice.CallOpts, avsAddr)
}

// GetAVSUSDValue is a free data retrieval call binding the contract method 0xdcf61b2c.
//
// Solidity: function getAVSUSDValue(address avsAddr) view returns(uint256)
func (_Contractavsservice *ContractavsserviceCallerSession) GetAVSUSDValue(avsAddr common.Address) (*big.Int, error) {
	return _Contractavsservice.Contract.GetAVSUSDValue(&_Contractavsservice.CallOpts, avsAddr)
}

// GetOperatorOptedUSDValue is a free data retrieval call binding the contract method 0x2d9d6a20.
//
// Solidity: function getOperatorOptedUSDValue(address avsAddr, string operatorAddr) view returns(uint256)
func (_Contractavsservice *ContractavsserviceCaller) GetOperatorOptedUSDValue(opts *bind.CallOpts, avsAddr common.Address, operatorAddr string) (*big.Int, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "getOperatorOptedUSDValue", avsAddr, operatorAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorOptedUSDValue is a free data retrieval call binding the contract method 0x2d9d6a20.
//
// Solidity: function getOperatorOptedUSDValue(address avsAddr, string operatorAddr) view returns(uint256)
func (_Contractavsservice *ContractavsserviceSession) GetOperatorOptedUSDValue(avsAddr common.Address, operatorAddr string) (*big.Int, error) {
	return _Contractavsservice.Contract.GetOperatorOptedUSDValue(&_Contractavsservice.CallOpts, avsAddr, operatorAddr)
}

// GetOperatorOptedUSDValue is a free data retrieval call binding the contract method 0x2d9d6a20.
//
// Solidity: function getOperatorOptedUSDValue(address avsAddr, string operatorAddr) view returns(uint256)
func (_Contractavsservice *ContractavsserviceCallerSession) GetOperatorOptedUSDValue(avsAddr common.Address, operatorAddr string) (*big.Int, error) {
	return _Contractavsservice.Contract.GetOperatorOptedUSDValue(&_Contractavsservice.CallOpts, avsAddr, operatorAddr)
}

// GetOptInOperators is a free data retrieval call binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) view returns(string[])
func (_Contractavsservice *ContractavsserviceCaller) GetOptInOperators(opts *bind.CallOpts, avsAddress common.Address) ([]string, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "getOptInOperators", avsAddress)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetOptInOperators is a free data retrieval call binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) view returns(string[])
func (_Contractavsservice *ContractavsserviceSession) GetOptInOperators(avsAddress common.Address) ([]string, error) {
	return _Contractavsservice.Contract.GetOptInOperators(&_Contractavsservice.CallOpts, avsAddress)
}

// GetOptInOperators is a free data retrieval call binding the contract method 0x1d4c8007.
//
// Solidity: function getOptInOperators(address avsAddress) view returns(string[])
func (_Contractavsservice *ContractavsserviceCallerSession) GetOptInOperators(avsAddress common.Address) ([]string, error) {
	return _Contractavsservice.Contract.GetOptInOperators(&_Contractavsservice.CallOpts, avsAddress)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x55b42cbe.
//
// Solidity: function getRegisteredPubkey(string operator) view returns(bytes)
func (_Contractavsservice *ContractavsserviceCaller) GetRegisteredPubkey(opts *bind.CallOpts, operator string) ([]byte, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "getRegisteredPubkey", operator)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x55b42cbe.
//
// Solidity: function getRegisteredPubkey(string operator) view returns(bytes)
func (_Contractavsservice *ContractavsserviceSession) GetRegisteredPubkey(operator string) ([]byte, error) {
	return _Contractavsservice.Contract.GetRegisteredPubkey(&_Contractavsservice.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x55b42cbe.
//
// Solidity: function getRegisteredPubkey(string operator) view returns(bytes)
func (_Contractavsservice *ContractavsserviceCallerSession) GetRegisteredPubkey(operator string) ([]byte, error) {
	return _Contractavsservice.Contract.GetRegisteredPubkey(&_Contractavsservice.CallOpts, operator)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0xe2906f3d.
//
// Solidity: function getTaskInfo(address taskAddr, uint64 taskID) view returns(uint64[])
func (_Contractavsservice *ContractavsserviceCaller) GetTaskInfo(opts *bind.CallOpts, taskAddr common.Address, taskID uint64) ([]uint64, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "getTaskInfo", taskAddr, taskID)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetTaskInfo is a free data retrieval call binding the contract method 0xe2906f3d.
//
// Solidity: function getTaskInfo(address taskAddr, uint64 taskID) view returns(uint64[])
func (_Contractavsservice *ContractavsserviceSession) GetTaskInfo(taskAddr common.Address, taskID uint64) ([]uint64, error) {
	return _Contractavsservice.Contract.GetTaskInfo(&_Contractavsservice.CallOpts, taskAddr, taskID)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0xe2906f3d.
//
// Solidity: function getTaskInfo(address taskAddr, uint64 taskID) view returns(uint64[])
func (_Contractavsservice *ContractavsserviceCallerSession) GetTaskInfo(taskAddr common.Address, taskID uint64) ([]uint64, error) {
	return _Contractavsservice.Contract.GetTaskInfo(&_Contractavsservice.CallOpts, taskAddr, taskID)
}

// IsOperator is a free data retrieval call binding the contract method 0x1f041bd1.
//
// Solidity: function isOperator(string operator) view returns(bool)
func (_Contractavsservice *ContractavsserviceCaller) IsOperator(opts *bind.CallOpts, operator string) (bool, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "isOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x1f041bd1.
//
// Solidity: function isOperator(string operator) view returns(bool)
func (_Contractavsservice *ContractavsserviceSession) IsOperator(operator string) (bool, error) {
	return _Contractavsservice.Contract.IsOperator(&_Contractavsservice.CallOpts, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x1f041bd1.
//
// Solidity: function isOperator(string operator) view returns(bool)
func (_Contractavsservice *ContractavsserviceCallerSession) IsOperator(operator string) (bool, error) {
	return _Contractavsservice.Contract.IsOperator(&_Contractavsservice.CallOpts, operator)
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

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0xf3298273.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, address taskContractAddress, string stage) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) OperatorSubmitTask(opts *bind.TransactOpts, taskID uint64, taskResponse []byte, blsSignature []byte, taskContractAddress common.Address, stage string) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "operatorSubmitTask", taskID, taskResponse, blsSignature, taskContractAddress, stage)
}

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0xf3298273.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, address taskContractAddress, string stage) returns(bool)
func (_Contractavsservice *ContractavsserviceSession) OperatorSubmitTask(taskID uint64, taskResponse []byte, blsSignature []byte, taskContractAddress common.Address, stage string) (*types.Transaction, error) {
	return _Contractavsservice.Contract.OperatorSubmitTask(&_Contractavsservice.TransactOpts, taskID, taskResponse, blsSignature, taskContractAddress, stage)
}

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0xf3298273.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, address taskContractAddress, string stage) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) OperatorSubmitTask(taskID uint64, taskResponse []byte, blsSignature []byte, taskContractAddress common.Address, stage string) (*types.Transaction, error) {
	return _Contractavsservice.Contract.OperatorSubmitTask(&_Contractavsservice.TransactOpts, taskID, taskResponse, blsSignature, taskContractAddress, stage)
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

// RegisterOperatorToExocore is a paid mutator transaction binding the contract method 0x1c275bf4.
//
// Solidity: function registerOperatorToExocore(string metaInfo) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactor) RegisterOperatorToExocore(opts *bind.TransactOpts, metaInfo string) (*types.Transaction, error) {
	return _Contractavsservice.contract.Transact(opts, "registerOperatorToExocore", metaInfo)
}

// RegisterOperatorToExocore is a paid mutator transaction binding the contract method 0x1c275bf4.
//
// Solidity: function registerOperatorToExocore(string metaInfo) returns(bool)
func (_Contractavsservice *ContractavsserviceSession) RegisterOperatorToExocore(metaInfo string) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterOperatorToExocore(&_Contractavsservice.TransactOpts, metaInfo)
}

// RegisterOperatorToExocore is a paid mutator transaction binding the contract method 0x1c275bf4.
//
// Solidity: function registerOperatorToExocore(string metaInfo) returns(bool)
func (_Contractavsservice *ContractavsserviceTransactorSession) RegisterOperatorToExocore(metaInfo string) (*types.Transaction, error) {
	return _Contractavsservice.Contract.RegisterOperatorToExocore(&_Contractavsservice.TransactOpts, metaInfo)
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
