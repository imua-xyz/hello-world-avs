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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"name\":\"createNewTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregisterOperatorFromAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"}],\"name\":\"getAVSInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"}],\"name\":\"getAVSUSDValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"}],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"operatorAddr\",\"type\":\"string\"}],\"name\":\"getOperatorOptedUSDValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddress\",\"type\":\"address\"}],\"name\":\"getOptInOperators\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"}],\"name\":\"getRegisteredPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"taskID\",\"type\":\"uint64\"}],\"name\":\"getTaskInfo\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"taskID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"taskResponse\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"stage\",\"type\":\"string\"}],\"name\":\"operatorSubmitTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"minStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"avsOwnerAddress\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"assetIds\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSelfDelegation\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64[]\",\"name\":\"params\",\"type\":\"uint64[]\"}],\"name\":\"registerAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyRegistrationSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyRegistrationMessageHash\",\"type\":\"bytes\"}],\"name\":\"registerBLSPublicKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerOperatorToAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metaInfo\",\"type\":\"string\"}],\"name\":\"registerOperatorToExocore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"minStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"avsOwnerAddress\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"assetIds\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSelfDelegation\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64[]\",\"name\":\"params\",\"type\":\"uint64[]\"}],\"name\":\"updateAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506125be8061005f6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063992907fb116100a2578063c49bb52111610071578063c49bb5211461032c578063dcf61b2c1461035c578063de16bf461461038c578063e2906f3d146103aa578063f3298273146103da5761010b565b8063992907fb1461027e578063af1991b3146102ae578063b32f34b3146102de578063c208dd991461030e5761010b565b806354c77f71116100de57806354c77f71146101d057806355b42cbe146102005780636f48e1a2146102305780638da5cb5b146102605761010b565b80631c275bf4146101105780631d4c8007146101405780631f041bd1146101705780632d9d6a20146101a0575b600080fd5b61012a6004803603810190610125919061107e565b61040a565b60405161013791906110e2565b60405180910390f35b61015a6004803603810190610155919061115b565b610498565b60405161016791906112c9565b60405180910390f35b61018a6004803603810190610185919061107e565b610528565b60405161019791906110e2565b60405180910390f35b6101ba60048036038101906101b591906112eb565b6105b2565b6040516101c79190611360565b60405180910390f35b6101ea60048036038101906101e5919061141c565b61063f565b6040516101f791906110e2565b60405180910390f35b61021a6004803603810190610215919061107e565b6106d6565b6040516102279190611548565b60405180910390f35b61024a60048036038101906102459190611753565b610779565b60405161025791906110e2565b60405180910390f35b610268610825565b60405161027591906118e2565b60405180910390f35b6102986004803603810190610293919061107e565b610849565b6040516102a59190611919565b60405180910390f35b6102c860048036038101906102c39190611934565b6108d3565b6040516102d591906110e2565b60405180910390f35b6102f860048036038101906102f3919061115b565b610ab9565b6040516103059190611a15565b60405180910390f35b610316610b49565b60405161032391906110e2565b60405180910390f35b61034660048036038101906103419190611753565b610bd3565b60405161035391906110e2565b60405180910390f35b6103766004803603810190610371919061115b565b610c7f565b6040516103839190611360565b60405180910390f35b610394610d09565b6040516103a191906110e2565b60405180910390f35b6103c460048036038101906103bf9190611a37565b610d93565b6040516103d19190611b35565b60405180910390f35b6103f460048036038101906103ef9190611bb2565b610e26565b60405161040191906110e2565b60405180910390f35b60008061090173ffffffffffffffffffffffffffffffffffffffff16635607f15633856040518363ffffffff1660e01b815260040161044a929190611c8a565b6020604051808303816000875af1158015610469573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061048d9190611ce6565b905080915050919050565b6060600061090173ffffffffffffffffffffffffffffffffffffffff16631d4c8007846040518263ffffffff1660e01b81526004016104d791906118e2565b600060405180830381865afa1580156104f4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061051d9190611e38565b905080915050919050565b60008061090173ffffffffffffffffffffffffffffffffffffffff16631f041bd1846040518263ffffffff1660e01b81526004016105669190611a15565b602060405180830381865afa158015610583573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105a79190611ce6565b905080915050919050565b60008061090173ffffffffffffffffffffffffffffffffffffffff16632d9d6a2085856040518363ffffffff1660e01b81526004016105f2929190611c8a565b602060405180830381865afa15801561060f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106339190611ead565b90508091505092915050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663aa57f77c33888888886040518663ffffffff1660e01b8152600401610685959493929190611eda565b6020604051808303816000875af11580156106a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106c89190611ce6565b905080915050949350505050565b6060600061090173ffffffffffffffffffffffffffffffffffffffff166355b42cbe846040518263ffffffff1660e01b81526004016107159190611a15565b600060405180830381865afa158015610732573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061075b9190611fb9565b9050808060200190518101906107719190611fb9565b915050919050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663d9e5daa0338f8f8f8f8f8f8f8f8f8f8f6040518d63ffffffff1660e01b81526004016107cd9c9b9a99989796959493929190612011565b6020604051808303816000875af11580156107ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108109190611ce6565b9050809150509b9a5050505050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663992907fb846040518263ffffffff1660e01b81526004016108879190611a15565b602060405180830381865afa1580156108a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c8919061211a565b905080915050919050565b60006108dd610ec6565b86816000018190525085816040019067ffffffffffffffff16908167ffffffffffffffff168152505084816060019067ffffffffffffffff16908167ffffffffffffffff1681525050828160a0019067ffffffffffffffff16908167ffffffffffffffff168152505083816080019067ffffffffffffffff16908167ffffffffffffffff168152505060008061090173ffffffffffffffffffffffffffffffffffffffff16638bf30a69338b8660405160200161099a91906121df565b604051602081830303815290604052805190602001206040516020016109c0919061222c565b6040516020818303038152906040528c8c8c8c6040518863ffffffff1660e01b81526004016109f59796959493929190612247565b60408051808303816000875af1158015610a13573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a3791906122d9565b915091508067ffffffffffffffff168360200181815250507fc3d2056aaa49a6e50ff7a353c7777b5cda982f2bf1e2214af811056cef07232d83602001513385600001518660400151876060015188608001518960a00151604051610aa29796959493929190612319565b60405180910390a181935050505095945050505050565b6060600061090173ffffffffffffffffffffffffffffffffffffffff1663b32f34b3846040518263ffffffff1660e01b8152600401610af891906118e2565b600060405180830381865afa158015610b15573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610b3e919061238f565b905080915050919050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663d7a2398b336040518263ffffffff1660e01b8152600401610b8791906118e2565b6020604051808303816000875af1158015610ba6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bca9190611ce6565b90508091505090565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663cde09950338f8f8f8f8f8f8f8f8f8f8f6040518d63ffffffff1660e01b8152600401610c279c9b9a99989796959493929190612011565b6020604051808303816000875af1158015610c46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c6a9190611ce6565b9050809150509b9a5050505050505050505050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663dcf61b2c846040518263ffffffff1660e01b8152600401610cbd91906118e2565b602060405180830381865afa158015610cda573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cfe9190611ead565b905080915050919050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663a364f4da336040518263ffffffff1660e01b8152600401610d4791906118e2565b6020604051808303816000875af1158015610d66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d8a9190611ce6565b90508091505090565b6060600061090173ffffffffffffffffffffffffffffffffffffffff1663e2906f3d85856040518363ffffffff1660e01b8152600401610dd49291906123d8565b600060405180830381865afa158015610df1573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610e1a9190612498565b90508091505092915050565b60008061090173ffffffffffffffffffffffffffffffffffffffff1663e538f58c338b8b8b8b8b8b8b6040518963ffffffff1660e01b8152600401610e7298979695949392919061250e565b6020604051808303816000875af1158015610e91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb59190611ce6565b905080915050979650505050505050565b6040518060c001604052806060815260200160008152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f8b82610f42565b810181811067ffffffffffffffff82111715610faa57610fa9610f53565b5b80604052505050565b6000610fbd610f24565b9050610fc98282610f82565b919050565b600067ffffffffffffffff821115610fe957610fe8610f53565b5b610ff282610f42565b9050602081019050919050565b82818337600083830152505050565b600061102161101c84610fce565b610fb3565b90508281526020810184848401111561103d5761103c610f3d565b5b611048848285610fff565b509392505050565b600082601f83011261106557611064610f38565b5b813561107584826020860161100e565b91505092915050565b60006020828403121561109457611093610f2e565b5b600082013567ffffffffffffffff8111156110b2576110b1610f33565b5b6110be84828501611050565b91505092915050565b60008115159050919050565b6110dc816110c7565b82525050565b60006020820190506110f760008301846110d3565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611128826110fd565b9050919050565b6111388161111d565b811461114357600080fd5b50565b6000813590506111558161112f565b92915050565b60006020828403121561117157611170610f2e565b5b600061117f84828501611146565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156111ee5780820151818401526020810190506111d3565b60008484015250505050565b6000611205826111b4565b61120f81856111bf565b935061121f8185602086016111d0565b61122881610f42565b840191505092915050565b600061123f83836111fa565b905092915050565b6000602082019050919050565b600061125f82611188565b6112698185611193565b93508360208202850161127b856111a4565b8060005b858110156112b757848403895281516112988582611233565b94506112a383611247565b925060208a0199505060018101905061127f565b50829750879550505050505092915050565b600060208201905081810360008301526112e38184611254565b905092915050565b6000806040838503121561130257611301610f2e565b5b600061131085828601611146565b925050602083013567ffffffffffffffff81111561133157611330610f33565b5b61133d85828601611050565b9150509250929050565b6000819050919050565b61135a81611347565b82525050565b60006020820190506113756000830184611351565b92915050565b600067ffffffffffffffff82111561139657611395610f53565b5b61139f82610f42565b9050602081019050919050565b60006113bf6113ba8461137b565b610fb3565b9050828152602081018484840111156113db576113da610f3d565b5b6113e6848285610fff565b509392505050565b600082601f83011261140357611402610f38565b5b81356114138482602086016113ac565b91505092915050565b6000806000806080858703121561143657611435610f2e565b5b600085013567ffffffffffffffff81111561145457611453610f33565b5b61146087828801611050565b945050602085013567ffffffffffffffff81111561148157611480610f33565b5b61148d878288016113ee565b935050604085013567ffffffffffffffff8111156114ae576114ad610f33565b5b6114ba878288016113ee565b925050606085013567ffffffffffffffff8111156114db576114da610f33565b5b6114e7878288016113ee565b91505092959194509250565b600081519050919050565b600082825260208201905092915050565b600061151a826114f3565b61152481856114fe565b93506115348185602086016111d0565b61153d81610f42565b840191505092915050565b60006020820190508181036000830152611562818461150f565b905092915050565b600067ffffffffffffffff82169050919050565b6115878161156a565b811461159257600080fd5b50565b6000813590506115a48161157e565b92915050565b600067ffffffffffffffff8211156115c5576115c4610f53565b5b602082029050602081019050919050565b600080fd5b60006115ee6115e9846115aa565b610fb3565b90508083825260208201905060208402830185811115611611576116106115d6565b5b835b8181101561165857803567ffffffffffffffff81111561163657611635610f38565b5b8086016116438982611050565b85526020850194505050602081019050611613565b5050509392505050565b600082601f83011261167757611676610f38565b5b81356116878482602086016115db565b91505092915050565b600067ffffffffffffffff8211156116ab576116aa610f53565b5b602082029050602081019050919050565b60006116cf6116ca84611690565b610fb3565b905080838252602082019050602084028301858111156116f2576116f16115d6565b5b835b8181101561171b57806117078882611595565b8452602084019350506020810190506116f4565b5050509392505050565b600082601f83011261173a57611739610f38565b5b813561174a8482602086016116bc565b91505092915050565b60008060008060008060008060008060006101608c8e03121561177957611778610f2e565b5b60008c013567ffffffffffffffff81111561179757611796610f33565b5b6117a38e828f01611050565b9b505060206117b48e828f01611595565b9a505060406117c58e828f01611146565b99505060606117d68e828f01611146565b98505060806117e78e828f01611146565b97505060a08c013567ffffffffffffffff81111561180857611807610f33565b5b6118148e828f01611662565b96505060c08c013567ffffffffffffffff81111561183557611834610f33565b5b6118418e828f01611662565b95505060e06118528e828f01611595565b9450506101006118648e828f01611595565b9350506101208c013567ffffffffffffffff81111561188657611885610f33565b5b6118928e828f01611050565b9250506101408c013567ffffffffffffffff8111156118b4576118b3610f33565b5b6118c08e828f01611725565b9150509295989b509295989b9093969950565b6118dc8161111d565b82525050565b60006020820190506118f760008301846118d3565b92915050565b60008160070b9050919050565b611913816118fd565b82525050565b600060208201905061192e600083018461190a565b92915050565b600080600080600060a086880312156119505761194f610f2e565b5b600086013567ffffffffffffffff81111561196e5761196d610f33565b5b61197a88828901611050565b955050602061198b88828901611595565b945050604061199c88828901611595565b93505060606119ad88828901611595565b92505060806119be88828901611595565b9150509295509295909350565b600082825260208201905092915050565b60006119e7826111b4565b6119f181856119cb565b9350611a018185602086016111d0565b611a0a81610f42565b840191505092915050565b60006020820190508181036000830152611a2f81846119dc565b905092915050565b60008060408385031215611a4e57611a4d610f2e565b5b6000611a5c85828601611146565b9250506020611a6d85828601611595565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b611aac8161156a565b82525050565b6000611abe8383611aa3565b60208301905092915050565b6000602082019050919050565b6000611ae282611a77565b611aec8185611a82565b9350611af783611a93565b8060005b83811015611b28578151611b0f8882611ab2565b9750611b1a83611aca565b925050600181019050611afb565b5085935050505092915050565b60006020820190508181036000830152611b4f8184611ad7565b905092915050565b600080fd5b60008083601f840112611b7257611b71610f38565b5b8235905067ffffffffffffffff811115611b8f57611b8e611b57565b5b602083019150836001820283011115611bab57611baa6115d6565b5b9250929050565b600080600080600080600060a0888a031215611bd157611bd0610f2e565b5b6000611bdf8a828b01611595565b975050602088013567ffffffffffffffff811115611c0057611bff610f33565b5b611c0c8a828b01611b5c565b9650965050604088013567ffffffffffffffff811115611c2f57611c2e610f33565b5b611c3b8a828b01611b5c565b94509450506060611c4e8a828b01611146565b925050608088013567ffffffffffffffff811115611c6f57611c6e610f33565b5b611c7b8a828b01611050565b91505092959891949750929550565b6000604082019050611c9f60008301856118d3565b8181036020830152611cb181846119dc565b90509392505050565b611cc3816110c7565b8114611cce57600080fd5b50565b600081519050611ce081611cba565b92915050565b600060208284031215611cfc57611cfb610f2e565b5b6000611d0a84828501611cd1565b91505092915050565b6000611d26611d2184610fce565b610fb3565b905082815260208101848484011115611d4257611d41610f3d565b5b611d4d8482856111d0565b509392505050565b600082601f830112611d6a57611d69610f38565b5b8151611d7a848260208601611d13565b91505092915050565b6000611d96611d91846115aa565b610fb3565b90508083825260208201905060208402830185811115611db957611db86115d6565b5b835b81811015611e0057805167ffffffffffffffff811115611dde57611ddd610f38565b5b808601611deb8982611d55565b85526020850194505050602081019050611dbb565b5050509392505050565b600082601f830112611e1f57611e1e610f38565b5b8151611e2f848260208601611d83565b91505092915050565b600060208284031215611e4e57611e4d610f2e565b5b600082015167ffffffffffffffff811115611e6c57611e6b610f33565b5b611e7884828501611e0a565b91505092915050565b611e8a81611347565b8114611e9557600080fd5b50565b600081519050611ea781611e81565b92915050565b600060208284031215611ec357611ec2610f2e565b5b6000611ed184828501611e98565b91505092915050565b600060a082019050611eef60008301886118d3565b8181036020830152611f0181876119dc565b90508181036040830152611f15818661150f565b90508181036060830152611f29818561150f565b90508181036080830152611f3d818461150f565b90509695505050505050565b6000611f5c611f578461137b565b610fb3565b905082815260208101848484011115611f7857611f77610f3d565b5b611f838482856111d0565b509392505050565b600082601f830112611fa057611f9f610f38565b5b8151611fb0848260208601611f49565b91505092915050565b600060208284031215611fcf57611fce610f2e565b5b600082015167ffffffffffffffff811115611fed57611fec610f33565b5b611ff984828501611f8b565b91505092915050565b61200b8161156a565b82525050565b600061018082019050612027600083018f6118d3565b8181036020830152612039818e6119dc565b9050612048604083018d612002565b612055606083018c6118d3565b612062608083018b6118d3565b61206f60a083018a6118d3565b81810360c08301526120818189611254565b905081810360e08301526120958188611254565b90506120a5610100830187612002565b6120b3610120830186612002565b8181036101408301526120c681856119dc565b90508181036101608301526120db8184611ad7565b90509d9c50505050505050505050505050565b6120f7816118fd565b811461210257600080fd5b50565b600081519050612114816120ee565b92915050565b6000602082840312156121305761212f610f2e565b5b600061213e84828501612105565b91505092915050565b61215081611347565b82525050565b600060c083016000830151848203600086015261217382826111fa565b91505060208301516121886020860182612147565b50604083015161219b6040860182611aa3565b5060608301516121ae6060860182611aa3565b5060808301516121c16080860182611aa3565b5060a08301516121d460a0860182611aa3565b508091505092915050565b600060208201905081810360008301526121f98184612156565b905092915050565b6000819050919050565b6000819050919050565b61222661222182612201565b61220b565b82525050565b60006122388284612215565b60208201915081905092915050565b600060e08201905061225c600083018a6118d3565b818103602083015261226e81896119dc565b90508181036040830152612282818861150f565b90506122916060830187612002565b61229e6080830186612002565b6122ab60a0830185612002565b6122b860c0830184612002565b98975050505050505050565b6000815190506122d38161157e565b92915050565b600080604083850312156122f0576122ef610f2e565b5b60006122fe85828601611cd1565b925050602061230f858286016122c4565b9150509250929050565b600060e08201905061232e600083018a611351565b61233b60208301896118d3565b818103604083015261234d81886119dc565b905061235c6060830187612002565b6123696080830186612002565b61237660a0830185612002565b61238360c0830184612002565b98975050505050505050565b6000602082840312156123a5576123a4610f2e565b5b600082015167ffffffffffffffff8111156123c3576123c2610f33565b5b6123cf84828501611d55565b91505092915050565b60006040820190506123ed60008301856118d3565b6123fa6020830184612002565b9392505050565b600061241461240f84611690565b610fb3565b90508083825260208201905060208402830185811115612437576124366115d6565b5b835b81811015612460578061244c88826122c4565b845260208401935050602081019050612439565b5050509392505050565b600082601f83011261247f5761247e610f38565b5b815161248f848260208601612401565b91505092915050565b6000602082840312156124ae576124ad610f2e565b5b600082015167ffffffffffffffff8111156124cc576124cb610f33565b5b6124d88482850161246a565b91505092915050565b60006124ed83856114fe565b93506124fa838584610fff565b61250383610f42565b840190509392505050565b600060c082019050612523600083018b6118d3565b612530602083018a612002565b818103604083015261254381888a6124e1565b905081810360608301526125588186886124e1565b905061256760808301856118d3565b81810360a083015261257981846119dc565b9050999850505050505050505056fea26469706673582212200099a5bdca66726b9abfd92687c5fc0903be55228ab0abcaca0297e766d1b1dc64736f6c63430008190033",
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

// GetCurrentEpoch is a free data retrieval call binding the contract method 0x992907fb.
//
// Solidity: function getCurrentEpoch(string epochIdentifier) view returns(int64)
func (_Contractavsservice *ContractavsserviceCaller) GetCurrentEpoch(opts *bind.CallOpts, epochIdentifier string) (int64, error) {
	var out []interface{}
	err := _Contractavsservice.contract.Call(opts, &out, "getCurrentEpoch", epochIdentifier)

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0x992907fb.
//
// Solidity: function getCurrentEpoch(string epochIdentifier) view returns(int64)
func (_Contractavsservice *ContractavsserviceSession) GetCurrentEpoch(epochIdentifier string) (int64, error) {
	return _Contractavsservice.Contract.GetCurrentEpoch(&_Contractavsservice.CallOpts, epochIdentifier)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0x992907fb.
//
// Solidity: function getCurrentEpoch(string epochIdentifier) view returns(int64)
func (_Contractavsservice *ContractavsserviceCallerSession) GetCurrentEpoch(epochIdentifier string) (int64, error) {
	return _Contractavsservice.Contract.GetCurrentEpoch(&_Contractavsservice.CallOpts, epochIdentifier)
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
