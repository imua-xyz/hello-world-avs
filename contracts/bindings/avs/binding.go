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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"taskResponsePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskChallengePeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"thresholdPercentage\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\"}],\"name\":\"createNewTask\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregisterOperatorFromAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"}],\"name\":\"getAVSInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"}],\"name\":\"getAVSUSDValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"}],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"operatorAddr\",\"type\":\"string\"}],\"name\":\"getOperatorOptedUSDValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"avsAddress\",\"type\":\"address\"}],\"name\":\"getOptInOperators\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"}],\"name\":\"getRegisteredPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"taskID\",\"type\":\"uint64\"}],\"name\":\"getTaskInfo\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"taskID\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"taskResponse\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"}],\"name\":\"operatorSubmitTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"minStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"avsOwnerAddress\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"assetIds\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSelfDelegation\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64[]\",\"name\":\"params\",\"type\":\"uint64[]\"}],\"name\":\"registerAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyRegistrationSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyRegistrationMessageHash\",\"type\":\"bytes\"}],\"name\":\"registerBLSPublicKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerOperatorToAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"minStakeAmount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"taskAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"avsOwnerAddress\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"assetIds\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minSelfDelegation\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"epochIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64[]\",\"name\":\"params\",\"type\":\"uint64[]\"}],\"name\":\"updateAVS\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50600080546001600160a01b031916331790556117ac806100316000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063992907fb11610097578063c49bb52111610066578063c49bb52114610262578063dcf61b2c14610275578063de16bf4614610288578063e2906f3d1461029057600080fd5b8063992907fb146101f6578063af1991b31461021c578063b32f34b314610247578063c208dd991461025a57600080fd5b806355b42cbe116100d357806355b42cbe146101855780636d70f7ae146101a55780636f48e1a2146101b85780638da5cb5b146101cb57600080fd5b80631d4c80071461010557806326135e9d1461012e5780632d9d6a201461015157806354c77f7114610172575b600080fd5b610118610113366004610a50565b6102b0565b6040516101259190610b15565b60405180910390f35b61014161013c366004610b93565b61032b565b6040519015158152602001610125565b61016461015f366004610d03565b6103b3565b604051908152602001610125565b610141610180366004610d50565b610428565b610198610193366004610dfc565b6104a7565b6040516101259190610e30565b6101416101b3366004610a50565b61051a565b6101416101c6366004610f58565b610589565b6000546101de906001600160a01b031681565b6040516001600160a01b039091168152602001610125565b610209610204366004610dfc565b610624565b60405160079190910b8152602001610125565b61022f61022a366004611095565b61068f565b6040516001600160401b039091168152602001610125565b610198610255366004610a50565b6107f4565b610141610868565b610141610270366004610f58565b6108d1565b610164610283366004610a50565b610918565b610141610987565b6102a361029e36600461111e565b6109b1565b604051610125919061119a565b604051631d4c800760e01b81526001600160a01b038216600482015260609060009061090190631d4c800790602401600060405180830381865afa1580156102fc573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261032491908101906111fd565b9392505050565b60405163046d13b160e11b81526000908190610901906308da2762906103639033908d908d908d908d908d908d908d906004016112d6565b6020604051808303816000875af1158015610382573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103a6919061133c565b9998505050505050505050565b60405163016ceb5160e51b8152600090819061090190632d9d6a20906103df908790879060040161135e565b602060405180830381865afa1580156103fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104209190611382565b949350505050565b604051632a95fddf60e21b815260009081906109019063aa57f77c9061045a9033908a908a908a908a9060040161139b565b6020604051808303816000875af1158015610479573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049d919061133c565b9695505050505050565b604051632ada165f60e11b8152606090610901906355b42cbe906104cf908590600401610e30565b600060405180830381865afa1580156104ec573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526105149190810190611405565b92915050565b6040516336b87bd760e11b81526001600160a01b0382166004820152600090819061090190636d70f7ae90602401602060405180830381865afa158015610565573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610324919061133c565b6000806109016001600160a01b031663d9e5daa0338f8f8f8f8f8f8f8f8f8f8f6040518d63ffffffff1660e01b81526004016105d09c9b9a9998979695949392919061144d565b6020604051808303816000875af11580156105ef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610613919061133c565b9d9c50505050505050505050505050565b60405163992907fb60e01b815260009081906109019063992907fb9061064e908690600401610e30565b602060405180830381865afa15801561066b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103249190611533565b6040805160c081018252600060208083018290528883526001600160401b0388811684860152878116606085015285811660a08501528616608084015292519092839161090191638bf30a699133918c916106ec91889101611556565b60408051601f19818403018152828252805160209182012090830152016040516020818303038152906040528b8b8b8b6040518863ffffffff1660e01b815260040161073e97969594939291906115c3565b6020604051808303816000875af115801561075d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610781919061162e565b6001600160401b0381166020840181905283516040808601516060870151608088015160a089015193519697507fc3d2056aaa49a6e50ff7a353c7777b5cda982f2bf1e2214af811056cef07232d966107e196953395909493929161164b565b60405180910390a1979650505050505050565b604051633824e10560e21b81526001600160a01b03821660048201526060906000906109019063e093841490602401600060405180830381865afa158015610840573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261032491908101906116a9565b60405163d7a2398b60e01b815233600482015260009081906109019063d7a2398b906024015b6020604051808303816000875af11580156108ad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610514919061133c565b6000806109016001600160a01b031663cde09950338f8f8f8f8f8f8f8f8f8f8f6040518d63ffffffff1660e01b81526004016105d09c9b9a9998979695949392919061144d565b60405163373d86cb60e21b81526001600160a01b038216600482015260009081906109019063dcf61b2c90602401602060405180830381865afa158015610963573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103249190611382565b6040516351b27a6d60e11b815233600482015260009081906109019063a364f4da9060240161088e565b60405163e2906f3d60e01b81526001600160a01b03831660048201526001600160401b03821660248201526060906000906109019063e2906f3d90604401600060405180830381865afa158015610a0c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261042091908101906116dd565b80356001600160a01b0381168114610a4b57600080fd5b919050565b600060208284031215610a6257600080fd5b61032482610a34565b60005b83811015610a86578181015183820152602001610a6e565b50506000910152565b60008151808452610aa7816020860160208601610a6b565b601f01601f19169290920160200192915050565b60008282518085526020808601955060208260051b8401016020860160005b84811015610b0857601f19868403018952610af6838351610a8f565b98840198925090830190600101610ada565b5090979650505050505050565b6020815260006103246020830184610abb565b6001600160401b0381168114610b3d57600080fd5b50565b8035610a4b81610b28565b60008083601f840112610b5d57600080fd5b5081356001600160401b03811115610b7457600080fd5b602083019150836020828501011115610b8c57600080fd5b9250929050565b600080600080600080600060a0888a031215610bae57600080fd5b8735610bb981610b28565b965060208801356001600160401b0380821115610bd557600080fd5b610be18b838c01610b4b565b909850965060408a0135915080821115610bfa57600080fd5b50610c078a828b01610b4b565b9095509350610c1a905060608901610a34565b9150608088013560ff81168114610c3057600080fd5b8091505092959891949750929550565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715610c7e57610c7e610c40565b604052919050565b60006001600160401b03821115610c9f57610c9f610c40565b50601f01601f191660200190565b600082601f830112610cbe57600080fd5b8135610cd1610ccc82610c86565b610c56565b818152846020838601011115610ce657600080fd5b816020850160208301376000918101602001919091529392505050565b60008060408385031215610d1657600080fd5b610d1f83610a34565b915060208301356001600160401b03811115610d3a57600080fd5b610d4685828601610cad565b9150509250929050565b60008060008060808587031215610d6657600080fd5b84356001600160401b0380821115610d7d57600080fd5b610d8988838901610cad565b95506020870135915080821115610d9f57600080fd5b610dab88838901610cad565b94506040870135915080821115610dc157600080fd5b610dcd88838901610cad565b93506060870135915080821115610de357600080fd5b50610df087828801610cad565b91505092959194509250565b600060208284031215610e0e57600080fd5b81356001600160401b03811115610e2457600080fd5b61042084828501610cad565b6020815260006103246020830184610a8f565b60006001600160401b03821115610e5c57610e5c610c40565b5060051b60200190565b600082601f830112610e7757600080fd5b81356020610e87610ccc83610e43565b82815260059290921b84018101918181019086841115610ea657600080fd5b8286015b84811015610ee55780356001600160401b03811115610ec95760008081fd5b610ed78986838b0101610cad565b845250918301918301610eaa565b509695505050505050565b600082601f830112610f0157600080fd5b81356020610f11610ccc83610e43565b8083825260208201915060208460051b870101935086841115610f3357600080fd5b602086015b84811015610ee5578035610f4b81610b28565b8352918301918301610f38565b60008060008060008060008060008060006101608c8e031215610f7a57600080fd5b6001600160401b03808d351115610f9057600080fd5b610f9d8e8e358f01610cad565b9b50610fab60208e01610b40565b9a50610fb960408e01610a34565b9950610fc760608e01610a34565b9850610fd560808e01610a34565b97508060a08e01351115610fe857600080fd5b610ff88e60a08f01358f01610e66565b96508060c08e0135111561100b57600080fd5b61101b8e60c08f01358f01610e66565b955061102960e08e01610b40565b94506110386101008e01610b40565b9350806101208e0135111561104c57600080fd5b61105d8e6101208f01358f01610cad565b9250806101408e0135111561107157600080fd5b506110838d6101408e01358e01610ef0565b90509295989b509295989b9093969950565b600080600080600060a086880312156110ad57600080fd5b85356001600160401b038111156110c357600080fd5b6110cf88828901610cad565b95505060208601356110e081610b28565b935060408601356110f081610b28565b9250606086013561110081610b28565b9150608086013561111081610b28565b809150509295509295909350565b6000806040838503121561113157600080fd5b61113a83610a34565b9150602083013561114a81610b28565b809150509250929050565b60008151808452602080850194506020840160005b8381101561118f5781516001600160401b03168752958201959082019060010161116a565b509495945050505050565b6020815260006103246020830184611155565b60006111bb610ccc84610c86565b90508281528383830111156111cf57600080fd5b610324836020830184610a6b565b600082601f8301126111ee57600080fd5b610324838351602085016111ad565b6000602080838503121561121057600080fd5b82516001600160401b038082111561122757600080fd5b818501915085601f83011261123b57600080fd5b8151611249610ccc82610e43565b81815260059190911b8301840190848101908883111561126857600080fd5b8585015b838110156112a0578051858111156112845760008081fd5b6112928b89838a01016111dd565b84525091860191860161126c565b5098975050505050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b600060018060a01b03808b1683526001600160401b038a16602084015260c0604084015261130860c08401898b6112ad565b838103606085015261131b81888a6112ad565b959091166080840152505060ff9190911660a0909101529695505050505050565b60006020828403121561134e57600080fd5b8151801515811461032457600080fd5b6001600160a01b038316815260406020820181905260009061042090830184610a8f565b60006020828403121561139457600080fd5b5051919050565b6001600160a01b038616815260a0602082018190526000906113bf90830187610a8f565b82810360408401526113d18187610a8f565b905082810360608401526113e58186610a8f565b905082810360808401526113f98185610a8f565b98975050505050505050565b60006020828403121561141757600080fd5b81516001600160401b0381111561142d57600080fd5b8201601f8101841361143e57600080fd5b610420848251602084016111ad565b6001600160a01b038d168152610180602082018190526000906114728382018f610a8f565b6001600160401b038e16604085015290506001600160a01b038c1660608401526001600160a01b038b1660808401526001600160a01b038a1660a084015282810360c08401526114c2818a610abb565b905082810360e08401526114d68189610abb565b6001600160401b03881661010085015290506001600160401b03861661012084015282810361014084015261150b8186610a8f565b90508281036101608401526115208185611155565b9f9e505050505050505050505050505050565b60006020828403121561154557600080fd5b81518060070b811461032457600080fd5b602081526000825160c0602084015261157260e0840182610a8f565b90506020840151604084015260408401516001600160401b0380821660608601528060608701511660808601528060808701511660a08601528060a08701511660c086015250508091505092915050565b6001600160a01b038816815260e0602082018190526000906115e790830189610a8f565b82810360408401526115f98189610a8f565b6001600160401b0397881660608501529587166080840152505091841660a083015290921660c0909201919091529392505050565b60006020828403121561164057600080fd5b815161032481610b28565b8781526001600160a01b038716602082015260e06040820181905260009061167590830188610a8f565b6001600160401b03968716606084015294861660808301525091841660a083015290921660c0909201919091529392505050565b6000602082840312156116bb57600080fd5b81516001600160401b038111156116d157600080fd5b610420848285016111dd565b600060208083850312156116f057600080fd5b82516001600160401b0381111561170657600080fd5b8301601f8101851361171757600080fd5b8051611725610ccc82610e43565b81815260059190911b8201830190838101908783111561174457600080fd5b928401925b8284101561176b57835161175c81610b28565b82529284019290840190611749565b97965050505050505056fea2646970667358221220693a3a834fa2cc3758e5f8131b2446aea247372bebef469942a1d8d15a294ecd64736f6c63430008190033",
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
