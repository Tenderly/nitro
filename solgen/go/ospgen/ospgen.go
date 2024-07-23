// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ospgen

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/tenderly/nitro/go-ethereum"
	"github.com/tenderly/nitro/go-ethereum/accounts/abi"
	"github.com/tenderly/nitro/go-ethereum/accounts/abi/bind"
	"github.com/tenderly/nitro/go-ethereum/common"
	"github.com/tenderly/nitro/go-ethereum/core/types"
	"github.com/tenderly/nitro/go-ethereum/event"
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

// ExecutionContext is an auto generated low-level Go binding around an user-defined struct.
type ExecutionContext struct {
	MaxInboxMessagesRead *big.Int
	Bridge               common.Address
}

// Instruction is an auto generated low-level Go binding around an user-defined struct.
type Instruction struct {
	Opcode       uint16
	ArgumentData *big.Int
}

// Machine is an auto generated low-level Go binding around an user-defined struct.
type Machine struct {
	Status          uint8
	ValueStack      ValueStack
	ValueMultiStack MultiStack
	InternalStack   ValueStack
	FrameStack      StackFrameWindow
	FrameMultiStack MultiStack
	GlobalStateHash [32]byte
	ModuleIdx       uint32
	FunctionIdx     uint32
	FunctionPc      uint32
	RecoveryPc      [32]byte
	ModulesRoot     [32]byte
}

// Module is an auto generated low-level Go binding around an user-defined struct.
type Module struct {
	GlobalsMerkleRoot   [32]byte
	ModuleMemory        ModuleMemory
	TablesMerkleRoot    [32]byte
	FunctionsMerkleRoot [32]byte
	ExtraHash           [32]byte
	InternalsOffset     uint32
}

// ModuleMemory is an auto generated low-level Go binding around an user-defined struct.
type ModuleMemory struct {
	Size       uint64
	MaxSize    uint64
	MerkleRoot [32]byte
}

// MultiStack is an auto generated low-level Go binding around an user-defined struct.
type MultiStack struct {
	InactiveStackHash [32]byte
	RemainingHash     [32]byte
}

// StackFrame is an auto generated low-level Go binding around an user-defined struct.
type StackFrame struct {
	ReturnPc              Value
	LocalsMerkleRoot      [32]byte
	CallerModule          uint32
	CallerModuleInternals uint32
}

// StackFrameWindow is an auto generated low-level Go binding around an user-defined struct.
type StackFrameWindow struct {
	Proved        []StackFrame
	RemainingHash [32]byte
}

// Value is an auto generated low-level Go binding around an user-defined struct.
type Value struct {
	ValueType uint8
	Contents  *big.Int
}

// ValueArray is an auto generated low-level Go binding around an user-defined struct.
type ValueArray struct {
	Inner []Value
}

// ValueStack is an auto generated low-level Go binding around an user-defined struct.
type ValueStack struct {
	Proved        ValueArray
	RemainingHash [32]byte
}

// HashProofHelperMetaData contains all meta data concerning the HashProofHelper contract.
var HashProofHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"NotProven\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"part\",\"type\":\"bytes\"}],\"name\":\"PreimagePartProven\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"clearSplitProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"getPreimagePart\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"keccakStates\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"part\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"}],\"name\":\"proveWithFullPreimage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"name\":\"proveWithSplitPreimage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"fullHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611cea806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063740085d71461005c57806379754cba14610085578063ae364ac2146100a6578063b7465799146100b0578063d4e5dd2b146100d2575b600080fd5b61006f61006a3660046118e0565b6100e5565b60405161007c9190611959565b60405180910390f35b6100986100933660046119bb565b6101d8565b60405190815260200161007c565b6100ae6106cb565b005b6100c36100be366004611a16565b610713565b60405161007c93929190611a3f565b6100986100e0366004611a71565b6107c9565b6000828152602081815260408083206001600160401b0385168452909152902080546060919060ff16610142576040516309cb23c960e11b8152600481018590526001600160401b03841660248201526044015b60405180910390fd5b80600101805461015190611ac4565b80601f016020809104026020016040519081016040528092919081815260200182805461017d90611ac4565b80156101ca5780601f1061019f576101008083540402835291602001916101ca565b820191906000526020600020905b8154815290600101906020018083116101ad57829003601f168201915b505050505091505092915050565b60006001821615156002831615610230573360009081526001602081905260408220805467ffffffffffffffff191681559190610217908301826117a2565b6102256002830160006117df565b600982016000905550505b80806102445750610242608886611b15565b155b6102845760405162461bcd60e51b81526020600482015260116024820152701393d517d09313d0d2d7d0531251d39151607a1b6044820152606401610139565b3360009081526001602052604090206009810154806102bc57815467ffffffffffffffff19166001600160401b038716178255610306565b81546001600160401b038781169116146103065760405162461bcd60e51b815260206004820152600b60248201526a1112519197d3d19194d15560aa1b6044820152606401610139565b61031282898986610920565b8061032760206001600160401b038916611b3f565b11801561034057508160090154866001600160401b0316105b1561045a57600081876001600160401b0316111561036e5761036b826001600160401b038916611b57565b90505b60008261038560206001600160401b038b16611b3f565b61038f9190611b57565b90508881111561039c5750875b815b8181101561045657846001018b8b838181106103bc576103bc611b6e565b9050013560f81c60f81b90808054806103d490611ac4565b80601f81146103e2576103f8565b83600052602060002060ff1984168155603f9350505b506002919091019091558154600116156104215790600052602060002090602091828204019190065b909190919091601f036101000a81548160ff02191690600160f81b84040217905550808061044e90611b84565b91505061039e565b5050505b8261046c5750600092506106c3915050565b60005b602081101561053c576000610485600883611b9f565b9050610492600582611b9f565b61049d600583611b15565b6104a8906005611bb3565b6104b29190611b3f565b905060006104c1600884611b15565b6104cc906008611bb3565b8560020183601981106104e1576104e1611b6e565b60048104909101546001600160401b036008600390931683026101000a9091041690911c9150610512908490611bb3565b61051d9060f8611b57565b60ff909116901b9590951794508061053481611b84565b91505061046f565b50604051806040016040528060011515815260200183600101805461056090611ac4565b80601f016020809104026020016040519081016040528092919081815260200182805461058c90611ac4565b80156105d95780601f106105ae576101008083540402835291602001916105d9565b820191906000526020600020905b8154815290600101906020018083116105bc57829003601f168201915b50505091909252505060008581526020818152604080832086546001600160401b0316845282529091208251815460ff1916901515178155828201518051919261062b926001850192909101906117ee565b505082546040516001600160401b03909116915085907ff88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c90610671906001870190611bd2565b60405180910390a33360009081526001602081905260408220805467ffffffffffffffff1916815591906106a7908301826117a2565b6106b56002830160006117df565b600982016000905550505050505b949350505050565b3360009081526001602081905260408220805467ffffffffffffffff1916815591906106f9908301826117a2565b6107076002830160006117df565b60098201600090555050565b6001602081905260009182526040909120805491810180546001600160401b039093169261074090611ac4565b80601f016020809104026020016040519081016040528092919081815260200182805461076c90611ac4565b80156107b95780601f1061078e576101008083540402835291602001916107b9565b820191906000526020600020905b81548152906001019060200180831161079c57829003601f168201915b5050505050908060090154905083565b600083836040516107db929190611c7a565b604051908190039020905060606001600160401b03831684111561087957600061080e6001600160401b03851686611b57565b9050602081111561081d575060205b856001600160401b038516866108338483611b3f565b9261084093929190611c8a565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929450505050505b6040805180820182526001808252602080830185815260008781528083528581206001600160401b038a1682528352949094208351815460ff1916901515178155935180519394936108d29385019291909101906117ee565b50905050826001600160401b0316827ff88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c836040516109109190611959565b60405180910390a3509392505050565b828290508460090160008282546109379190611b3f565b90915550505b81158015610949575080155b1561095357610ba2565b60005b6088811015610a7a5760008382101561098c5784848381811061097b5761097b611b6e565b919091013560f81c91506109af9050565b81841415610998576001175b6109a460016088611b57565b8214156109af576080175b60006109bc600884611b9f565b90506109c9600582611b9f565b6109d4600583611b15565b6109df906005611bb3565b6109e99190611b3f565b90506109f6600884611b15565b610a01906008611bb3565b6001600160401b03168260ff166001600160401b0316901b876002018260198110610a2e57610a2e611b6e565b6004810490910180546001600160401b0360086003909416939093026101000a808204841690941883168402929093021990921617905550819050610a7281611b84565b915050610956565b50610a83611872565b60005b6019811015610af557856002018160198110610aa457610aa4611b6e565b600491828204019190066008029054906101000a90046001600160401b03166001600160401b0316828260198110610ade57610ade611b6e565b602002015280610aed81611b84565b915050610a86565b50610aff81610ba8565b905060005b6019811015610b7b57818160198110610b1f57610b1f611b6e565b6020020151866002018260198110610b3957610b39611b6e565b600491828204019190066008026101000a8154816001600160401b0302191690836001600160401b031602179055508080610b7390611b84565b915050610b04565b506088831015610b8b5750610ba2565b610b988360888187611c8a565b935093505061093d565b50505050565b610bb0611872565b610bb8611891565b610bc0611891565b610bc8611872565b600060405180610300016040528060018152602001618082815260200167800000000000808a8152602001678000000080008000815260200161808b81526020016380000001815260200167800000008000808181526020016780000000000080098152602001608a81526020016088815260200163800080098152602001638000000a8152602001638000808b815260200167800000000000008b8152602001678000000000008089815260200167800000000000800381526020016780000000000080028152602001678000000000000080815260200161800a815260200167800000008000000a81526020016780000000800080818152602001678000000000008080815260200163800000018152602001678000000080008008815250905060005b6018811015611797576080878101516060808a01516040808c01516020808e01518e511890911890921890931889526101208b01516101008c015160e08d015160c08e015160a08f0151181818189089018190526101c08b01516101a08c01516101808d01516101608e01516101408f0151181818189289019283526102608b01516102408c01516102208d01516102008e01516101e08f015118181818918901919091526103008a01516102e08b01516102c08c01516102a08d01516102808e0151181818189288018390526001600160401b0360028202166001603f1b91829004179092188652510485600260200201516002026001600160401b03161785600060200201511884600160200201526001603f1b856003602002015181610e1957610e19611aff565b0485600360200201516002026001600160401b03161785600160200201511884600260200201526001603f1b856004602002015181610e5a57610e5a611aff565b0485600460200201516002026001600160401b03161785600260058110610e8357610e83611b6e565b602002015118606085015284516001603f1b9086516060808901519390920460029091026001600160401b031617909118608086810191825286518a5118808b5287516020808d018051909218825289516040808f0180519092189091528a518e8801805190911890528a51948e0180519095189094528901805160a08e0180519091189052805160c08e0180519091189052805160e08e018051909118905280516101008e0180519091189052516101208d018051909118905291880180516101408d018051909118905280516101608d018051909118905280516101808d018051909118905280516101a08d0180519091189052516101c08c018051909118905292870180516101e08c018051909118905280516102008c018051909118905280516102208c018051909118905280516102408c0180519091189052516102608b018051909118905281516102808b018051909118905281516102a08b018051909118905281516102c08b018051909118905281516102e08b018051909118905290516103008a01805190911890529084525163100000009060208901516001600160401b03641000000000909102169190041761010084015260408701516001603d1b9060408901516001600160401b03600890910216919004176101608401526060870151628000009060608901516001600160401b036502000000000090910216919004176102608401526080870151654000000000009060808901516001600160401b036204000090910216919004176102c084015260a08701516001603f1b900487600560200201516002026001600160401b031617836002601981106110f3576110f3611b6e565b602002015260c08701516210000081046001602c1b9091026001600160401b039081169190911760a085015260e0880151664000000000000081046104009091028216176101a08501526101008801516208000081046520000000000090910282161761020085015261012088015160048082029092166001603e1b909104176103008501526101408801516101408901516001600160401b036001603e1b909102169190041760808401526101608701516001603a1b906101608901516001600160401b036040909102169190041760e084015261018087015162200000906101808901516001600160401b036001602b1b90910216919004176101408401526101a08701516602000000000000906101a08901516001600160401b0361800090910216919004176102408401526101c08701516008906101c08901516001600160401b036001603d1b90910216919004176102a08401526101e0870151641000000000906101e08901516001600160401b03631000000090910216919004176020840152610200808801516102008901516001600160401b0366800000000000009091021691900417610120840152610220870151648000000000906102208901516001600160401b03630200000090910216919004176101808401526102408701516001602b1b906102408901516001600160401b036220000090910216919004176101e0840152610260870151610100906102608901516001600160401b03600160381b90910216919004176102e0840152610280870151642000000000906102808901516001600160401b036308000000909102169190041760608401526102a08701516001602c1b906102a08901516001600160401b0362100000909102169190041760c08401526102c08701516302000000906102c08901516001600160401b0364800000000090910216919004176101c08401526102e0870151600160381b906102e08901516001600160401b036101009091021691900417610220840152610300870151660400000000000090048760186020020151614000026001600160401b031617836014602002015282600a602002015183600560200201511916836000602002015118876000602002015282600b602002015183600660200201511916836001602002015118876001602002015282600c602002015183600760200201511916836002602002015118876002602002015282600d602002015183600860200201511916836003602002015118876003602002015282600e602002015183600960200201511916836004602002015118876004602002015282600f602002015183600a602002015119168360056020020151188760056020020152826010602002015183600b602002015119168360066020020151188760066020020152826011602002015183600c602002015119168360076020020151188760076020020152826012602002015183600d602002015119168360086020020151188760086020020152826013602002015183600e602002015119168360096020020151188760096020020152826014602002015183600f6020020151191683600a60200201511887600a602002015282601560200201518360106020020151191683600b60200201511887600b602002015282601660200201518360116020020151191683600c60200201511887600c602002015282601760200201518360126020020151191683600d60200201511887600d602002015282601860200201518360136020020151191683600e60200201511887600e602002015282600060200201518360146020020151191683600f60200201511887600f602002015282600160200201518360156020020151191683601060200201511887601060200201528260026020020151836016602002015119168360116020020151188760116020020152826003602002015183601760200201511916836012602002015118876012602002015282600460200201518360186020020151191683601360200201511887601360200201528260056020020151836000602002015119168360146020020151188760146020020152826006602002015183600160200201511916836015602002015118876015602002015282600760200201518360026020020151191683601660200201511887601660200201528260086020020151836003602002015119168360176020020151188760176020020152826009602002015183600460200201511916836018602002015118876018602002015281816018811061178557611785611b6e565b60200201518751188752600101610cee565b509495945050505050565b5080546117ae90611ac4565b6000825580601f106117be575050565b601f0160209004906000526020600020908101906117dc91906118af565b50565b506117dc9060078101906118af565b8280546117fa90611ac4565b90600052602060002090601f01602090048101928261181c5760008555611862565b82601f1061183557805160ff1916838001178555611862565b82800160010185558215611862579182015b82811115611862578251825591602001919060010190611847565b5061186e9291506118af565b5090565b6040518061032001604052806019906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b5b8082111561186e57600081556001016118b0565b80356001600160401b03811681146118db57600080fd5b919050565b600080604083850312156118f357600080fd5b82359150611903602084016118c4565b90509250929050565b6000815180845260005b8181101561193257602081850181015186830182015201611916565b81811115611944576000602083870101525b50601f01601f19169290920160200192915050565b60208152600061196c602083018461190c565b9392505050565b60008083601f84011261198557600080fd5b5081356001600160401b0381111561199c57600080fd5b6020830191508360208285010111156119b457600080fd5b9250929050565b600080600080606085870312156119d157600080fd5b84356001600160401b038111156119e757600080fd5b6119f387828801611973565b9095509350611a069050602086016118c4565b9396929550929360400135925050565b600060208284031215611a2857600080fd5b81356001600160a01b038116811461196c57600080fd5b6001600160401b0384168152606060208201526000611a61606083018561190c565b9050826040830152949350505050565b600080600060408486031215611a8657600080fd5b83356001600160401b03811115611a9c57600080fd5b611aa886828701611973565b9094509250611abb9050602085016118c4565b90509250925092565b600181811c90821680611ad857607f821691505b60208210811415611af957634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601260045260246000fd5b600082611b2457611b24611aff565b500690565b634e487b7160e01b600052601160045260246000fd5b60008219821115611b5257611b52611b29565b500190565b600082821015611b6957611b69611b29565b500390565b634e487b7160e01b600052603260045260246000fd5b6000600019821415611b9857611b98611b29565b5060010190565b600082611bae57611bae611aff565b500490565b6000816000190483118215151615611bcd57611bcd611b29565b500290565b600060208083526000845481600182811c915080831680611bf457607f831692505b858310811415611c1257634e487b7160e01b85526022600452602485fd5b878601838152602001818015611c2f5760018114611c4057611c6b565b60ff19861682528782019650611c6b565b60008b81526020902060005b86811015611c6557815484820152908501908901611c4c565b83019750505b50949998505050505050505050565b8183823760009101908152919050565b60008085851115611c9a57600080fd5b83861115611ca757600080fd5b505082019391909203915056fea26469706673582212202f2dcb3af29934bcb41f7861970718df9ad65661f9f6cb26c88c963224a1e4c264736f6c63430008090033",
}

// HashProofHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use HashProofHelperMetaData.ABI instead.
var HashProofHelperABI = HashProofHelperMetaData.ABI

// HashProofHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HashProofHelperMetaData.Bin instead.
var HashProofHelperBin = HashProofHelperMetaData.Bin

// DeployHashProofHelper deploys a new Ethereum contract, binding an instance of HashProofHelper to it.
func DeployHashProofHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashProofHelper, error) {
	parsed, err := HashProofHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HashProofHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HashProofHelper{HashProofHelperCaller: HashProofHelperCaller{contract: contract}, HashProofHelperTransactor: HashProofHelperTransactor{contract: contract}, HashProofHelperFilterer: HashProofHelperFilterer{contract: contract}}, nil
}

// HashProofHelper is an auto generated Go binding around an Ethereum contract.
type HashProofHelper struct {
	HashProofHelperCaller     // Read-only binding to the contract
	HashProofHelperTransactor // Write-only binding to the contract
	HashProofHelperFilterer   // Log filterer for contract events
}

// HashProofHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashProofHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashProofHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashProofHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashProofHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashProofHelperSession struct {
	Contract     *HashProofHelper  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashProofHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashProofHelperCallerSession struct {
	Contract *HashProofHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// HashProofHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashProofHelperTransactorSession struct {
	Contract     *HashProofHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// HashProofHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashProofHelperRaw struct {
	Contract *HashProofHelper // Generic contract binding to access the raw methods on
}

// HashProofHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashProofHelperCallerRaw struct {
	Contract *HashProofHelperCaller // Generic read-only contract binding to access the raw methods on
}

// HashProofHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashProofHelperTransactorRaw struct {
	Contract *HashProofHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashProofHelper creates a new instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelper(address common.Address, backend bind.ContractBackend) (*HashProofHelper, error) {
	contract, err := bindHashProofHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashProofHelper{HashProofHelperCaller: HashProofHelperCaller{contract: contract}, HashProofHelperTransactor: HashProofHelperTransactor{contract: contract}, HashProofHelperFilterer: HashProofHelperFilterer{contract: contract}}, nil
}

// NewHashProofHelperCaller creates a new read-only instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperCaller(address common.Address, caller bind.ContractCaller) (*HashProofHelperCaller, error) {
	contract, err := bindHashProofHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperCaller{contract: contract}, nil
}

// NewHashProofHelperTransactor creates a new write-only instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*HashProofHelperTransactor, error) {
	contract, err := bindHashProofHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperTransactor{contract: contract}, nil
}

// NewHashProofHelperFilterer creates a new log filterer instance of HashProofHelper, bound to a specific deployed contract.
func NewHashProofHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*HashProofHelperFilterer, error) {
	contract, err := bindHashProofHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperFilterer{contract: contract}, nil
}

// bindHashProofHelper binds a generic wrapper to an already deployed contract.
func bindHashProofHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HashProofHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashProofHelper *HashProofHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashProofHelper.Contract.HashProofHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashProofHelper *HashProofHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.Contract.HashProofHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashProofHelper *HashProofHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashProofHelper.Contract.HashProofHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashProofHelper *HashProofHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashProofHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashProofHelper *HashProofHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashProofHelper *HashProofHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashProofHelper.Contract.contract.Transact(opts, method, params...)
}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperCaller) GetPreimagePart(opts *bind.CallOpts, fullHash [32]byte, offset uint64) ([]byte, error) {
	var out []interface{}
	err := _HashProofHelper.contract.Call(opts, &out, "getPreimagePart", fullHash, offset)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperSession) GetPreimagePart(fullHash [32]byte, offset uint64) ([]byte, error) {
	return _HashProofHelper.Contract.GetPreimagePart(&_HashProofHelper.CallOpts, fullHash, offset)
}

// GetPreimagePart is a free data retrieval call binding the contract method 0x740085d7.
//
// Solidity: function getPreimagePart(bytes32 fullHash, uint64 offset) view returns(bytes)
func (_HashProofHelper *HashProofHelperCallerSession) GetPreimagePart(fullHash [32]byte, offset uint64) ([]byte, error) {
	return _HashProofHelper.Contract.GetPreimagePart(&_HashProofHelper.CallOpts, fullHash, offset)
}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperCaller) KeccakStates(opts *bind.CallOpts, arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	var out []interface{}
	err := _HashProofHelper.contract.Call(opts, &out, "keccakStates", arg0)

	outstruct := new(struct {
		Offset uint64
		Part   []byte
		Length *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Offset = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Part = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Length = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperSession) KeccakStates(arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	return _HashProofHelper.Contract.KeccakStates(&_HashProofHelper.CallOpts, arg0)
}

// KeccakStates is a free data retrieval call binding the contract method 0xb7465799.
//
// Solidity: function keccakStates(address ) view returns(uint64 offset, bytes part, uint256 length)
func (_HashProofHelper *HashProofHelperCallerSession) KeccakStates(arg0 common.Address) (struct {
	Offset uint64
	Part   []byte
	Length *big.Int
}, error) {
	return _HashProofHelper.Contract.KeccakStates(&_HashProofHelper.CallOpts, arg0)
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperTransactor) ClearSplitProof(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "clearSplitProof")
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperSession) ClearSplitProof() (*types.Transaction, error) {
	return _HashProofHelper.Contract.ClearSplitProof(&_HashProofHelper.TransactOpts)
}

// ClearSplitProof is a paid mutator transaction binding the contract method 0xae364ac2.
//
// Solidity: function clearSplitProof() returns()
func (_HashProofHelper *HashProofHelperTransactorSession) ClearSplitProof() (*types.Transaction, error) {
	return _HashProofHelper.Contract.ClearSplitProof(&_HashProofHelper.TransactOpts)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactor) ProveWithFullPreimage(opts *bind.TransactOpts, data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "proveWithFullPreimage", data, offset)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperSession) ProveWithFullPreimage(data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithFullPreimage(&_HashProofHelper.TransactOpts, data, offset)
}

// ProveWithFullPreimage is a paid mutator transaction binding the contract method 0xd4e5dd2b.
//
// Solidity: function proveWithFullPreimage(bytes data, uint64 offset) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactorSession) ProveWithFullPreimage(data []byte, offset uint64) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithFullPreimage(&_HashProofHelper.TransactOpts, data, offset)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactor) ProveWithSplitPreimage(opts *bind.TransactOpts, data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.contract.Transact(opts, "proveWithSplitPreimage", data, offset, flags)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperSession) ProveWithSplitPreimage(data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithSplitPreimage(&_HashProofHelper.TransactOpts, data, offset, flags)
}

// ProveWithSplitPreimage is a paid mutator transaction binding the contract method 0x79754cba.
//
// Solidity: function proveWithSplitPreimage(bytes data, uint64 offset, uint256 flags) returns(bytes32 fullHash)
func (_HashProofHelper *HashProofHelperTransactorSession) ProveWithSplitPreimage(data []byte, offset uint64, flags *big.Int) (*types.Transaction, error) {
	return _HashProofHelper.Contract.ProveWithSplitPreimage(&_HashProofHelper.TransactOpts, data, offset, flags)
}

// HashProofHelperPreimagePartProvenIterator is returned from FilterPreimagePartProven and is used to iterate over the raw logs and unpacked data for PreimagePartProven events raised by the HashProofHelper contract.
type HashProofHelperPreimagePartProvenIterator struct {
	Event *HashProofHelperPreimagePartProven // Event containing the contract specifics and raw log

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
func (it *HashProofHelperPreimagePartProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashProofHelperPreimagePartProven)
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
		it.Event = new(HashProofHelperPreimagePartProven)
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
func (it *HashProofHelperPreimagePartProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashProofHelperPreimagePartProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashProofHelperPreimagePartProven represents a PreimagePartProven event raised by the HashProofHelper contract.
type HashProofHelperPreimagePartProven struct {
	FullHash [32]byte
	Offset   uint64
	Part     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPreimagePartProven is a free log retrieval operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) FilterPreimagePartProven(opts *bind.FilterOpts, fullHash [][32]byte, offset []uint64) (*HashProofHelperPreimagePartProvenIterator, error) {

	var fullHashRule []interface{}
	for _, fullHashItem := range fullHash {
		fullHashRule = append(fullHashRule, fullHashItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _HashProofHelper.contract.FilterLogs(opts, "PreimagePartProven", fullHashRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return &HashProofHelperPreimagePartProvenIterator{contract: _HashProofHelper.contract, event: "PreimagePartProven", logs: logs, sub: sub}, nil
}

// WatchPreimagePartProven is a free log subscription operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) WatchPreimagePartProven(opts *bind.WatchOpts, sink chan<- *HashProofHelperPreimagePartProven, fullHash [][32]byte, offset []uint64) (event.Subscription, error) {

	var fullHashRule []interface{}
	for _, fullHashItem := range fullHash {
		fullHashRule = append(fullHashRule, fullHashItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _HashProofHelper.contract.WatchLogs(opts, "PreimagePartProven", fullHashRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashProofHelperPreimagePartProven)
				if err := _HashProofHelper.contract.UnpackLog(event, "PreimagePartProven", log); err != nil {
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

// ParsePreimagePartProven is a log parse operation binding the contract event 0xf88493e8ac6179d3c1ba8712068367d7ecdd6f30d3b5de01198e7a449fe2802c.
//
// Solidity: event PreimagePartProven(bytes32 indexed fullHash, uint64 indexed offset, bytes part)
func (_HashProofHelper *HashProofHelperFilterer) ParsePreimagePartProven(log types.Log) (*HashProofHelperPreimagePartProven, error) {
	event := new(HashProofHelperPreimagePartProven)
	if err := _HashProofHelper.contract.UnpackLog(event, "PreimagePartProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOneStepProofEntryMetaData contains all meta data concerning the IOneStepProofEntry contract.
var IOneStepProofEntryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"}],\"name\":\"getEndMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"}],\"name\":\"getStartMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"machineStep\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveOneStep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IOneStepProofEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use IOneStepProofEntryMetaData.ABI instead.
var IOneStepProofEntryABI = IOneStepProofEntryMetaData.ABI

// IOneStepProofEntry is an auto generated Go binding around an Ethereum contract.
type IOneStepProofEntry struct {
	IOneStepProofEntryCaller     // Read-only binding to the contract
	IOneStepProofEntryTransactor // Write-only binding to the contract
	IOneStepProofEntryFilterer   // Log filterer for contract events
}

// IOneStepProofEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOneStepProofEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOneStepProofEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOneStepProofEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProofEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOneStepProofEntrySession struct {
	Contract     *IOneStepProofEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IOneStepProofEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOneStepProofEntryCallerSession struct {
	Contract *IOneStepProofEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IOneStepProofEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOneStepProofEntryTransactorSession struct {
	Contract     *IOneStepProofEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IOneStepProofEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOneStepProofEntryRaw struct {
	Contract *IOneStepProofEntry // Generic contract binding to access the raw methods on
}

// IOneStepProofEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOneStepProofEntryCallerRaw struct {
	Contract *IOneStepProofEntryCaller // Generic read-only contract binding to access the raw methods on
}

// IOneStepProofEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOneStepProofEntryTransactorRaw struct {
	Contract *IOneStepProofEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOneStepProofEntry creates a new instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntry(address common.Address, backend bind.ContractBackend) (*IOneStepProofEntry, error) {
	contract, err := bindIOneStepProofEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntry{IOneStepProofEntryCaller: IOneStepProofEntryCaller{contract: contract}, IOneStepProofEntryTransactor: IOneStepProofEntryTransactor{contract: contract}, IOneStepProofEntryFilterer: IOneStepProofEntryFilterer{contract: contract}}, nil
}

// NewIOneStepProofEntryCaller creates a new read-only instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryCaller(address common.Address, caller bind.ContractCaller) (*IOneStepProofEntryCaller, error) {
	contract, err := bindIOneStepProofEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryCaller{contract: contract}, nil
}

// NewIOneStepProofEntryTransactor creates a new write-only instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*IOneStepProofEntryTransactor, error) {
	contract, err := bindIOneStepProofEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryTransactor{contract: contract}, nil
}

// NewIOneStepProofEntryFilterer creates a new log filterer instance of IOneStepProofEntry, bound to a specific deployed contract.
func NewIOneStepProofEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*IOneStepProofEntryFilterer, error) {
	contract, err := bindIOneStepProofEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOneStepProofEntryFilterer{contract: contract}, nil
}

// bindIOneStepProofEntry binds a generic wrapper to an already deployed contract.
func bindIOneStepProofEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProofEntry *IOneStepProofEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.IOneStepProofEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProofEntry *IOneStepProofEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProofEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProofEntry *IOneStepProofEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProofEntry *IOneStepProofEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProofEntry.Contract.contract.Transact(opts, method, params...)
}

// GetEndMachineHash is a free data retrieval call binding the contract method 0xd8558b87.
//
// Solidity: function getEndMachineHash(uint8 status, bytes32 globalStateHash) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntryCaller) GetEndMachineHash(opts *bind.CallOpts, status uint8, globalStateHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IOneStepProofEntry.contract.Call(opts, &out, "getEndMachineHash", status, globalStateHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEndMachineHash is a free data retrieval call binding the contract method 0xd8558b87.
//
// Solidity: function getEndMachineHash(uint8 status, bytes32 globalStateHash) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntrySession) GetEndMachineHash(status uint8, globalStateHash [32]byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.GetEndMachineHash(&_IOneStepProofEntry.CallOpts, status, globalStateHash)
}

// GetEndMachineHash is a free data retrieval call binding the contract method 0xd8558b87.
//
// Solidity: function getEndMachineHash(uint8 status, bytes32 globalStateHash) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntryCallerSession) GetEndMachineHash(status uint8, globalStateHash [32]byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.GetEndMachineHash(&_IOneStepProofEntry.CallOpts, status, globalStateHash)
}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntryCaller) GetStartMachineHash(opts *bind.CallOpts, globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IOneStepProofEntry.contract.Call(opts, &out, "getStartMachineHash", globalStateHash, wasmModuleRoot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntrySession) GetStartMachineHash(globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.GetStartMachineHash(&_IOneStepProofEntry.CallOpts, globalStateHash, wasmModuleRoot)
}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_IOneStepProofEntry *IOneStepProofEntryCallerSession) GetStartMachineHash(globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.GetStartMachineHash(&_IOneStepProofEntry.CallOpts, globalStateHash, wasmModuleRoot)
}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntryCaller) ProveOneStep(opts *bind.CallOpts, execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	var out []interface{}
	err := _IOneStepProofEntry.contract.Call(opts, &out, "proveOneStep", execCtx, machineStep, beforeHash, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntrySession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.ProveOneStep(&_IOneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_IOneStepProofEntry *IOneStepProofEntryCallerSession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _IOneStepProofEntry.Contract.ProveOneStep(&_IOneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// IOneStepProverMetaData contains all meta data concerning the IOneStepProver contract.
var IOneStepProverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"instruction\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"result\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"resultMod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IOneStepProverABI is the input ABI used to generate the binding from.
// Deprecated: Use IOneStepProverMetaData.ABI instead.
var IOneStepProverABI = IOneStepProverMetaData.ABI

// IOneStepProver is an auto generated Go binding around an Ethereum contract.
type IOneStepProver struct {
	IOneStepProverCaller     // Read-only binding to the contract
	IOneStepProverTransactor // Write-only binding to the contract
	IOneStepProverFilterer   // Log filterer for contract events
}

// IOneStepProverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOneStepProverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOneStepProverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOneStepProverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOneStepProverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOneStepProverSession struct {
	Contract     *IOneStepProver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOneStepProverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOneStepProverCallerSession struct {
	Contract *IOneStepProverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IOneStepProverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOneStepProverTransactorSession struct {
	Contract     *IOneStepProverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IOneStepProverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOneStepProverRaw struct {
	Contract *IOneStepProver // Generic contract binding to access the raw methods on
}

// IOneStepProverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOneStepProverCallerRaw struct {
	Contract *IOneStepProverCaller // Generic read-only contract binding to access the raw methods on
}

// IOneStepProverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOneStepProverTransactorRaw struct {
	Contract *IOneStepProverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOneStepProver creates a new instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProver(address common.Address, backend bind.ContractBackend) (*IOneStepProver, error) {
	contract, err := bindIOneStepProver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOneStepProver{IOneStepProverCaller: IOneStepProverCaller{contract: contract}, IOneStepProverTransactor: IOneStepProverTransactor{contract: contract}, IOneStepProverFilterer: IOneStepProverFilterer{contract: contract}}, nil
}

// NewIOneStepProverCaller creates a new read-only instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverCaller(address common.Address, caller bind.ContractCaller) (*IOneStepProverCaller, error) {
	contract, err := bindIOneStepProver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverCaller{contract: contract}, nil
}

// NewIOneStepProverTransactor creates a new write-only instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverTransactor(address common.Address, transactor bind.ContractTransactor) (*IOneStepProverTransactor, error) {
	contract, err := bindIOneStepProver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverTransactor{contract: contract}, nil
}

// NewIOneStepProverFilterer creates a new log filterer instance of IOneStepProver, bound to a specific deployed contract.
func NewIOneStepProverFilterer(address common.Address, filterer bind.ContractFilterer) (*IOneStepProverFilterer, error) {
	contract, err := bindIOneStepProver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOneStepProverFilterer{contract: contract}, nil
}

// bindIOneStepProver binds a generic wrapper to an already deployed contract.
func bindIOneStepProver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IOneStepProverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProver *IOneStepProverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProver.Contract.IOneStepProverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProver *IOneStepProverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProver.Contract.IOneStepProverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProver *IOneStepProverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProver.Contract.IOneStepProverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOneStepProver *IOneStepProverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOneStepProver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOneStepProver *IOneStepProverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOneStepProver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOneStepProver *IOneStepProverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOneStepProver.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverCaller) ExecuteOneStep(opts *bind.CallOpts, execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	var out []interface{}
	err := _IOneStepProver.contract.Call(opts, &out, "executeOneStep", execCtx, mach, mod, instruction, proof)

	outstruct := new(struct {
		Result    Machine
		ResultMod Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Result = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.ResultMod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverSession) ExecuteOneStep(execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	return _IOneStepProver.Contract.ExecuteOneStep(&_IOneStepProver.CallOpts, execCtx, mach, mod, instruction, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod, (uint16,uint256) instruction, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) result, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) resultMod)
func (_IOneStepProver *IOneStepProverCallerSession) ExecuteOneStep(execCtx ExecutionContext, mach Machine, mod Module, instruction Instruction, proof []byte) (struct {
	Result    Machine
	ResultMod Module
}, error) {
	return _IOneStepProver.Contract.ExecuteOneStep(&_IOneStepProver.CallOpts, execCtx, mach, mod, instruction, proof)
}

// OneStepProofEntryMetaData contains all meta data concerning the OneStepProofEntry contract.
var OneStepProofEntryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"prover0_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverMem_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverMath_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProver\",\"name\":\"proverHostIo_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"}],\"name\":\"getEndMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"}],\"name\":\"getStartMachineHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"machineStep\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"proveOneStep\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"afterHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prover0\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverHostIo\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverMath\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proverMem\",\"outputs\":[{\"internalType\":\"contractIOneStepProver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002f2f38038062002f2f8339810160408190526200003491620000a5565b600080546001600160a01b039586166001600160a01b031991821617909155600180549486169482169490941790935560028054928516928416929092179091556003805491909316911617905562000102565b80516001600160a01b0381168114620000a057600080fd5b919050565b60008060008060808587031215620000bc57600080fd5b620000c78562000088565b9350620000d76020860162000088565b9250620000e76040860162000088565b9150620000f76060860162000088565b905092959194509250565b612e1d80620001126000396000f3fe608060405234801561001057600080fd5b506004361061006d5760003560e01c806304997be4146100725780631f128bc01461009857806330a5509f146100c35780635d3adcfb146100d65780635f52fd7c146100e957806366e5d9c3146100fc578063d8558b871461010f575b600080fd5b6100856100803660046121c8565b610122565b6040519081526020015b60405180910390f35b6001546100ab906001600160a01b031681565b6040516001600160a01b03909116815260200161008f565b6000546100ab906001600160a01b031681565b6100856100e43660046121ea565b6102c8565b6003546100ab906001600160a01b031681565b6002546100ab906001600160a01b031681565b61008561011d366004612295565b61095a565b60408051600380825260808201909252600091829190816020015b604080518082019091526000808252602082015281526020019060019003908161013d57505060408051808201825260008082526020918201819052825180840190935260048352908201529091508160008151811061019f5761019f6122d7565b60200260200101819052506101b46000610a5b565b816001815181106101c7576101c76122d7565b60200260200101819052506101dc6000610a5b565b816002815181106101ef576101ef6122d7565b6020908102919091018101919091526040805180830182528381528151808301909252808252600092820192909252610226612067565b604080518082018252606080825260006020808401829052845180860186528082018390526000198082528651610180810188528481529283018990529582018190529281018690526080810184905260a0810183905260c081018c905260e081018290526101008101829052610120810191909152610140810193909352610160830189905290916102b881610a8e565b9750505050505050505b92915050565b60006102d2612085565b6102da61213c565b6040805160208101909152606081526040805180820190915260008082526020820152600061030a888883610ccc565b90955090508861031986610a8e565b146103615760405162461bcd60e51b815260206004820152601360248201527209a828690929c8abe848a8c9ea48abe9082a69606b1b60448201526064015b60405180910390fd5b600085516003811115610376576103766122ed565b146103905761038485610a8e565b95505050505050610951565b650800000000006103a28b6001612319565b14156103b5576002855261038485610a8e565b6103c0888883610f28565b90945090506103d0888883610fec565b80925081945050508461016001516103fd8660e0015163ffffffff1686866110c69092919063ffffffff16565b146104395760405162461bcd60e51b815260206004820152600c60248201526b1353d115531154d7d493d3d560a21b6044820152606401610358565b60606104516040518060200160405280606081525090565b60408051602081019091526060815261046b8b8b86611111565b9450925061047a8b8b86610fec565b945091506104898b8b86610fec565b809550819250505060006104bf60408a61012001516104a89190612347565b63ffffffff1685856112109092919063ffffffff16565b905060006104e38a610100015163ffffffff1683856112569092919063ffffffff16565b90508860600151811461052d5760405162461bcd60e51b815260206004820152601260248201527110905117d1955390d51253d394d7d493d3d560721b6044820152606401610358565b8460408b6101200151610540919061236a565b63ffffffff1681518110610556576105566122d7565b6020026020010151965050505050508787829080926105779392919061238d565b975097505060008460e0015163ffffffff1690506001856101200181815161059f91906123b7565b63ffffffff1690525081516000602861ffff8316108015906105c65750603561ffff831611155b806105e65750603661ffff8316108015906105e65750603e61ffff831611155b806105f5575061ffff8216603f145b80610604575061ffff82166040145b1561061b57506001546001600160a01b0316610832565b61ffff821660451480610632575061ffff82166050145b806106605750604661ffff8316108015906106605750610654600960466123df565b61ffff168261ffff1611155b8061068e5750606761ffff83161080159061068e5750610682600260676123df565b61ffff168261ffff1611155b806106ae5750606a61ffff8316108015906106ae5750607861ffff831611155b806106dc5750605161ffff8316108015906106dc57506106d0600960516123df565b61ffff168261ffff1611155b8061070a5750607961ffff83161080159061070a57506106fe600260796123df565b61ffff168261ffff1611155b8061072a5750607c61ffff83161080159061072a5750608a61ffff831611155b80610739575061ffff821660a7145b80610756575061ffff821660ac1480610756575061ffff821660ad145b80610776575060c061ffff831610801590610776575060c461ffff831611155b80610796575060bc61ffff831610801590610796575060bf61ffff831611155b156107ad57506002546001600160a01b0316610832565b61801061ffff8316108015906107c9575061801361ffff831611155b806107eb575061802061ffff8316108015906107eb575061802461ffff831611155b8061080d575061803061ffff83161080159061080d575061803261ffff831611155b1561082457506003546001600160a01b0316610832565b506000546001600160a01b03165b806001600160a01b0316633604366f8e8989888f8f6040518763ffffffff1660e01b81526004016108689695949392919061253e565b60006040518083038186803b15801561088057600080fd5b505afa158015610894573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526108bc9190810190612b67565b9097509550600061ffff831661802314806108dc575061ffff8316618024145b15905080156108f7576108f08685896110c6565b6101608901525b60028851600381111561090c5761090c6122ed565b148015610920575061014088015160001914155b1561093d5761092e886112cb565b6109378861134f565b50600088525b61094688610a8e565b985050505050505050505b95945050505050565b60006001836003811115610970576109706122ed565b14156109a457816040516020016109879190612cc1565b6040516020818303038152906040528051906020012090506102c2565b60028360038111156109b8576109b86122ed565b14156109e2576040516f26b0b1b434b7329032b93937b932b21d60811b6020820152603001610987565b60038360038111156109f6576109f66122ed565b1415610a20576040516f26b0b1b434b732903a37b7903330b91d60811b6020820152603001610987565b60405162461bcd60e51b815260206004820152601060248201526f4241445f424c4f434b5f53544154555360801b6044820152606401610358565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b60008082516003811115610aa457610aa46122ed565b1415610bc2576000610ad3610abc846020015161137f565b610140850151604086015191906000191415611404565b90506000610afe610ae785608001516114f7565b61014086015160a087015191906000191415611404565b9050600082610b10866060015161137f565b60c087015160e0808901516101008a01516101208b01516101408c01516101608d01516040516f26b0b1b434b73290393ab73734b7339d60811b602082015260308101999099526050890197909752607088018a905260908801959095526001600160e01b031992841b831660b088015290831b821660b487015290911b1660b884015260bc83015260dc82015260fc0160408051601f19818403018152919052805160209091012095945050505050565b600182516003811115610bd757610bd76122ed565b1415610c0f578160c00151604051602001610bf29190612cc1565b604051602081830303815290604052805190602001209050919050565b600282516003811115610c2457610c246122ed565b1415610c4e576040516f26b0b1b434b7329032b93937b932b21d60811b6020820152603001610bf2565b600382516003811115610c6357610c636122ed565b1415610c8d576040516f26b0b1b434b732903a37b7903330b91d60811b6020820152603001610bf2565b60405162461bcd60e51b815260206004820152600f60248201526e4241445f4d4143485f53544154555360881b6044820152606401610358565b919050565b610cd4612085565b81600080610ce3878785611590565b9350905060ff8116610cf85760009150610d75565b8060ff1660011415610d0d5760019150610d75565b8060ff1660021415610d225760029150610d75565b8060ff1660031415610d375760039150610d75565b60405162461bcd60e51b8152602060048201526013602482015272554e4b4e4f574e5f4d4143485f53544154555360681b6044820152606401610358565b50610d7e612067565b610d86612067565b60408051808201909152600080825260208201526040805180820190915260608152600060208201526040805180820190915260008082526020820152610dce8b8b896115c6565b97509450610ddd8b8b896116c5565b97509250610dec8b8b896115c6565b97509350610dfb8b8b8961171b565b97509150610e0a8b8b896116c5565b8098508192505050604051806101800160405280876003811115610e3057610e306122ed565b8152602081019690965260408601939093526060850193909352608084015260a0830191909152600060c0830181905260e08301819052610100830181905261012083018190526101408301819052610160909201919091529250610e989050858583611843565b60c08401919091529050610ead85858361185f565b63ffffffff90911660e08401529050610ec785858361185f565b63ffffffff9091166101008401529050610ee285858361185f565b63ffffffff9091166101208401529050610efd858583611843565b6101408401919091529050610f13858583611843565b61016084019190915291959194509092505050565b610f3061213c565b60408051606081018252600080825260208201819052918101829052839190600080600080610f608b8b89611843565b97509550610f6f8b8b896118c3565b97509450610f7e8b8b89611843565b97509350610f8d8b8b89611843565b97509250610f9c8b8b89611843565b97509150610fab8b8b8961185f565b6040805160c081018252988952602089019790975295870194909452506060850191909152608084015263ffffffff1660a083015290969095509350505050565b604080516020810190915260608152816000611009868684611590565b92509050600060ff82166001600160401b0381111561102a5761102a6122c1565b604051908082528060200260200182016040528015611053578160200160208202803683370190505b50905060005b8260ff168160ff1610156110aa57611072888886611843565b838360ff1681518110611087576110876122d7565b6020026020010181965082815250505080806110a290612ce6565b915050611059565b5060405180602001604052808281525093505050935093915050565b600061110784846110d68561193e565b6040518060400160405280601381526020017226b7b23ab6329036b2b935b632903a3932b29d60691b8152506119ba565b90505b9392505050565b6060816000611121868684611590565b9250905060ff81166001600160401b03811115611140576111406122c1565b60405190808252806020026020018201604052801561118557816020015b604080518082019091526000808252602082015281526020019060019003908161115e5790505b50925060005b8160ff16811015611206576000806111a4898987611acc565b955091506111b3898987611b25565b809650819250505060405180604001604052808361ffff168152602001828152508684815181106111e6576111e66122d7565b6020026020010181905250505080806111fe90612d06565b91505061118b565b5050935093915050565b6000611107848461122085611b7a565b6040518060400160405280601881526020017724b739ba393ab1ba34b7b71036b2b935b632903a3932b29d60411b8152506119ba565b60405168233ab731ba34b7b71d60b91b602082015260298101829052600090819060490160405160208183030381529060405280519060200120905061095185858360405180604001604052806015815260200174233ab731ba34b7b71036b2b935b632903a3932b29d60591b8152506119ba565b60408101515160a0820151516000198114806112e8575060001982145b156112f557505060029052565b61130283608001516114f7565b60a08401515260208301516113169061137f565b60408401515260808301516113319082602082015260609052565b50602091820151808301919091526040805192830190526060825252565b60006113638283610140015160001c611d19565b61136f57506000919050565b5060001961014090910152600190565b60208101518151515160005b818110156113fd5783516113a8906113a39083611d5b565b611d93565b6040516b2b30b63ab29039ba30b1b59d60a11b6020820152602c810191909152604c8101849052606c0160405160208183030381529060405280519060200120925080806113f590612d06565b91505061138b565b5050919050565b60006000198314156114545760405162461bcd60e51b81526020600482015260196024820152784d554c5449535441434b5f4e4f535441434b5f41435449564560381b6044820152606401610358565b81156114dd57835160001914156114a75760405162461bcd60e51b815260206004820152601760248201527626aaa62a24a9aa20a1a5afa727a9aa20a1a5afa6a0a4a760491b6044820152606401610358565b83516020808601516040516114c0939287929101612d21565b60405160208183030381529060405280519060200120905061110a565b83516020808601516040516114c093879390929101612d21565b602081015160005b82515181101561158a5761152f83600001518281518110611522576115226122d7565b6020026020010151611db0565b6040517129ba30b1b590333930b6b29039ba30b1b59d60711b6020820152603281019190915260528101839052607201604051602081830303815290604052805190602001209150808061158290612d06565b9150506114ff565b50919050565b6000818484828181106115a5576115a56122d7565b919091013560f81c92508190506115bb81612d06565b915050935093915050565b6115ce612067565b8160006115dc868684611843565b9250905060006115ed878785611b25565b935090506000816001600160401b0381111561160b5761160b6122c1565b60405190808252806020026020018201604052801561165057816020015b60408051808201909152600080825260208201528152602001906001900390816116295790505b50905060005b815181101561169e5761166a898987611e20565b83838151811061167c5761167c6122d7565b602002602001018197508290525050808061169690612d06565b915050611656565b50604080516060810182529081019182529081526020810192909252509590945092505050565b60408051808201909152600080825260208201528160006116e7868684611843565b9250905060006116f8878785611843565b604080518082019091529384526020840191909152919791965090945050505050565b60408051808201909152606081526000602082015281600061173e868684611843565b925090506060868684818110611756576117566122d7565b909101356001600160f81b0319161590506117de578261177581612d06565b604080516001808252818301909252919550909150816020015b611797612194565b81526020019060019003908161178f5790505090506117b7878785611f1c565b826000815181106117ca576117ca6122d7565b602002602001018195508290525050611822565b826117e881612d06565b6040805160008082526020820190925291955090915061181e565b61180b612194565b8152602001906001900390816118035790505b5090505b60405180604001604052808281526020018381525093505050935093915050565b60008181611852868684611b25565b9097909650945050505050565b600081815b60048110156118ba5760088363ffffffff16901b925085858381811061188c5761188c6122d7565b919091013560f81c939093179250816118a481612d06565b92505080806118b290612d06565b915050611864565b50935093915050565b604080516060810182526000808252602082018190529181019190915281600080806118f0888886611fb5565b945092506118ff888886611fb5565b9450915061190e888886611843565b604080516060810182526001600160401b0396871681529490951660208501529383015250969095509350505050565b600081600001516119528360200151612013565b6040808501516060860151608087015160a08801519351610bf2969594906020016626b7b23ab6329d60c91b81526007810196909652602786019490945260478501929092526067840152608783015260e01b6001600160e01b03191660a782015260ab0190565b8160005b855151811015611a835760018516611a1f578282876000015183815181106119e8576119e86122d7565b6020026020010151604051602001611a0293929190612d4d565b604051602081830303815290604052805190602001209150611a6a565b8286600001518281518110611a3657611a366122d7565b602002602001015183604051602001611a5193929190612d4d565b6040516020818303038152906040528051906020012091505b60019490941c9380611a7b81612d06565b9150506119be565b508315611ac45760405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b6044820152606401610358565b949350505050565b600081815b60028110156118ba5760088361ffff16901b9250858583818110611af757611af76122d7565b919091013560f81c93909317925081611b0f81612d06565b9250508080611b1d90612d06565b915050611ad1565b600081815b60208110156118ba57600883901b9250858583818110611b4c57611b4c6122d7565b919091013560f81c93909317925081611b6481612d06565b9250508080611b7290612d06565b915050611b2a565b60008082516022611b8b9190612d93565b611b9690600e612319565b6001600160401b03811115611bad57611bad6122c1565b6040519080825280601f01601f191660200182016040528015611bd7576020820181803683370190505b5090506c24b739ba393ab1ba34b7b7399d60991b60208201526000600d9050835160f81b828281518110611c0d57611c0d6122d7565b60200101906001600160f81b031916908160001a90535080611c2e81612d06565b91505060005b8451811015611d09576000858281518110611c5157611c516122d7565b602002602001015190506008816000015161ffff16901c60f81b848481518110611c7d57611c7d6122d7565b60200101906001600160f81b031916908160001a905350805160f81b84611ca5856001612319565b81518110611cb557611cb56122d7565b60200101906001600160f81b031916908160001a905350611cd7600284612319565b6020808301518683018201819052919450611cf29085612319565b935050508080611d0190612d06565b915050611c34565b5050805160209091012092915050565b6000606082901c15611d2d575060006102c2565b5063ffffffff818116610120840152602082901c811661010084015260409190911c1660e090910152600190565b60408051808201909152600080825260208201528251805183908110611d8357611d836122d7565b6020026020010151905092915050565b600081600001518260200151604051602001610bf2929190612db2565b6000611dbf8260000151611d93565b602080840151604080860151606087015191516b29ba30b1b590333930b6b29d60a11b94810194909452602c840194909452604c8301919091526001600160e01b031960e093841b8116606c840152921b9091166070820152607401610bf2565b6040805180820190915260008082526020820152816000858583818110611e4957611e496122d7565b919091013560f81c9150829050611e5f81612d06565b925050611e6a600690565b6006811115611e7b57611e7b6122ed565b60ff168160ff161115611ec15760405162461bcd60e51b815260206004820152600e60248201526d4241445f56414c55455f5459504560901b6044820152606401610358565b6000611ece878785611b25565b809450819250505060405180604001604052808360ff166006811115611ef657611ef66122ed565b6006811115611f0757611f076122ed565b81526020018281525093505050935093915050565b611f24612194565b81611f3f604080518082019091526000808252602082015290565b6000806000611f4f898987611e20565b95509350611f5e898987611843565b95509250611f6d89898761185f565b95509150611f7c89898761185f565b60408051608081018252968752602087019590955263ffffffff9384169486019490945290911660608401525090969095509350505050565b600081815b60088110156118ba576008836001600160401b0316901b9250858583818110611fe557611fe56122d7565b919091013560f81c93909317925081611ffd81612d06565b925050808061200b90612d06565b915050611fba565b805160208083015160408085015190516626b2b6b7b93c9d60c91b938101939093526001600160c01b031960c094851b811660278501529190931b16602f8201526037810191909152600090605701610bf2565b60408051606080820183529181019182529081526000602082015290565b60408051610180810190915280600081526020016120a1612067565b81526020016120c0604080518082019091526000808252602082015290565b81526020016120cd612067565b81526020016120ed60408051808201909152606081526000602082015290565b815260200161210c604080518082019091526000808252602082015290565b815260006020820181905260408201819052606082018190526080820181905260a0820181905260c09091015290565b6040805160c081019091526000815260208101612172604080516060810182526000808252602082018190529181019190915290565b8152600060208201819052604082018190526060820181905260809091015290565b6040805160c08101825260006080820181815260a08301829052825260208201819052918101829052606081019190915290565b600080604083850312156121db57600080fd5b50508035926020909101359150565b600080600080600085870360a081121561220357600080fd5b604081121561221157600080fd5b50859450604086013593506060860135925060808601356001600160401b038082111561223d57600080fd5b818801915088601f83011261225157600080fd5b81358181111561226057600080fd5b89602082850101111561227257600080fd5b9699959850939650602001949392505050565b6004811061229257600080fd5b50565b600080604083850312156122a857600080fd5b82356122b381612285565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000821982111561232c5761232c612303565b500190565b634e487b7160e01b600052601260045260246000fd5b600063ffffffff8084168061235e5761235e612331565b92169190910492915050565b600063ffffffff8084168061238157612381612331565b92169190910692915050565b6000808585111561239d57600080fd5b838611156123aa57600080fd5b5050820193919092039150565b600063ffffffff8083168185168083038211156123d6576123d6612303565b01949350505050565b600061ffff8083168185168083038211156123d6576123d6612303565b6004811061240c5761240c6122ed565b9052565b805160078110612422576124226122ed565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b8084101561247f5761246b828651612410565b938201936001939093019290850190612458565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b828110156124fe5784516124ca858251612410565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a0909301926001016124b5565b509687015197909601969096525093949350505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8635815260006101c060208901356001600160a01b03811680821461256257600080fd5b806020860152505080604084015261257d81840189516123fc565b6020880151816101e085015261259761038085018261242f565b91505060408801516125b761020085018280518252602090810151910152565b5060608801516101bf1980858403016102408601526125d6838361242f565b925060808a015191508085840301610260860152506125f58282612493565b91505060a088015161261561028085018280518252602090810151910152565b5060c08801516102c084015260e088015163ffffffff81166102e08501525061010088015163ffffffff81166103008501525061012088015163ffffffff811661032085015250610140880151610340840152610160808901516103608501526126e160608501898051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a0830152608081015160c083015263ffffffff60a08201511660e08301525050565b6126fb81850188805161ffff168252602090810151910152565b508281036101a0840152612710818587612515565b9998505050505050505050565b604080519081016001600160401b038111828210171561273f5761273f6122c1565b60405290565b604051602081016001600160401b038111828210171561273f5761273f6122c1565b604051608081016001600160401b038111828210171561273f5761273f6122c1565b60405160c081016001600160401b038111828210171561273f5761273f6122c1565b604051606081016001600160401b038111828210171561273f5761273f6122c1565b60405161018081016001600160401b038111828210171561273f5761273f6122c1565b604051601f8201601f191681016001600160401b0381118282101715612818576128186122c1565b604052919050565b8051610cc781612285565b60006001600160401b03821115612844576128446122c1565b5060051b60200190565b60006040828403121561286057600080fd5b61286861271d565b905081516007811061287957600080fd5b808252506020820151602082015292915050565b600060408083850312156128a057600080fd5b6128a861271d565b915082516001600160401b03808211156128c157600080fd5b818501915060208083880312156128d757600080fd5b6128df612745565b8351838111156128ee57600080fd5b80850194505087601f85011261290357600080fd5b835192506129186129138461282b565b6127f0565b83815260069390931b8401820192828101908985111561293757600080fd5b948301945b8486101561295d5761294e8a8761284e565b8252948601949083019061293c565b8252508552948501519484019490945250909392505050565b60006040828403121561298857600080fd5b61299061271d565b9050815181526020820151602082015292915050565b805163ffffffff81168114610cc757600080fd5b600060408083850312156129cd57600080fd5b6129d561271d565b915082516001600160401b038111156129ed57600080fd5b8301601f810185136129fe57600080fd5b80516020612a0e6129138361282b565b82815260a09283028401820192828201919089851115612a2d57600080fd5b948301945b84861015612a965780868b031215612a4a5760008081fd5b612a52612767565b612a5c8b8861284e565b815287870151858201526060612a738189016129a6565b89830152612a83608089016129a6565b9082015283529485019491830191612a32565b50808752505080860151818601525050505092915050565b80516001600160401b0381168114610cc757600080fd5b6000818303610100811215612ad957600080fd5b612ae1612789565b8351815291506060601f1982011215612af957600080fd5b50612b026127ab565b612b0e60208401612aae565b8152612b1c60408401612aae565b602082015260608301516040820152806020830152506080820151604082015260a0820151606082015260c08201516080820152612b5c60e083016129a6565b60a082015292915050565b600080610120808486031215612b7c57600080fd5b83516001600160401b0380821115612b9357600080fd5b908501906101c08288031215612ba857600080fd5b612bb06127cd565b612bb983612820565b8152602083015182811115612bcd57600080fd5b612bd98982860161288d565b602083015250612bec8860408501612976565b6040820152608083015182811115612c0357600080fd5b612c0f8982860161288d565b60608301525060a083015182811115612c2757600080fd5b612c33898286016129ba565b608083015250612c468860c08501612976565b60a082015261010091508183015160c0820152612c648484016129a6565b60e0820152610140612c778185016129a6565b838301526101609250612c8b8385016129a6565b8583015261018084015181830152506101a08301518282015280955050505050612cb88460208501612ac5565b90509250929050565b7026b0b1b434b732903334b734b9b432b21d60791b8152601181019190915260310190565b600060ff821660ff811415612cfd57612cfd612303565b60010192915050565b6000600019821415612d1a57612d1a612303565b5060010190565b6a36bab63a34b9ba30b1b59d60a91b8152600b810193909352602b830191909152604b820152606b0190565b6000845160005b81811015612d6e5760208188018101518583015201612d54565b81811115612d7d576000828501525b5091909101928352506020820152604001919050565b6000816000190483118215151615612dad57612dad612303565b500290565b652b30b63ab29d60d11b8152600060078410612dd057612dd06122ed565b5060f89290921b600683015260078201526027019056fea26469706673582212206630cc3ca936424a379d3ba5328e6633b67f26f2f4c58f0334dbe3c122002ad364736f6c63430008090033",
}

// OneStepProofEntryABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProofEntryMetaData.ABI instead.
var OneStepProofEntryABI = OneStepProofEntryMetaData.ABI

// OneStepProofEntryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProofEntryMetaData.Bin instead.
var OneStepProofEntryBin = OneStepProofEntryMetaData.Bin

// DeployOneStepProofEntry deploys a new Ethereum contract, binding an instance of OneStepProofEntry to it.
func DeployOneStepProofEntry(auth *bind.TransactOpts, backend bind.ContractBackend, prover0_ common.Address, proverMem_ common.Address, proverMath_ common.Address, proverHostIo_ common.Address) (common.Address, *types.Transaction, *OneStepProofEntry, error) {
	parsed, err := OneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProofEntryBin), backend, prover0_, proverMem_, proverMath_, proverHostIo_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProofEntry{OneStepProofEntryCaller: OneStepProofEntryCaller{contract: contract}, OneStepProofEntryTransactor: OneStepProofEntryTransactor{contract: contract}, OneStepProofEntryFilterer: OneStepProofEntryFilterer{contract: contract}}, nil
}

// OneStepProofEntry is an auto generated Go binding around an Ethereum contract.
type OneStepProofEntry struct {
	OneStepProofEntryCaller     // Read-only binding to the contract
	OneStepProofEntryTransactor // Write-only binding to the contract
	OneStepProofEntryFilterer   // Log filterer for contract events
}

// OneStepProofEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofEntrySession struct {
	Contract     *OneStepProofEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OneStepProofEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofEntryCallerSession struct {
	Contract *OneStepProofEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OneStepProofEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofEntryTransactorSession struct {
	Contract     *OneStepProofEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OneStepProofEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofEntryRaw struct {
	Contract *OneStepProofEntry // Generic contract binding to access the raw methods on
}

// OneStepProofEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofEntryCallerRaw struct {
	Contract *OneStepProofEntryCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofEntryTransactorRaw struct {
	Contract *OneStepProofEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProofEntry creates a new instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntry(address common.Address, backend bind.ContractBackend) (*OneStepProofEntry, error) {
	contract, err := bindOneStepProofEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntry{OneStepProofEntryCaller: OneStepProofEntryCaller{contract: contract}, OneStepProofEntryTransactor: OneStepProofEntryTransactor{contract: contract}, OneStepProofEntryFilterer: OneStepProofEntryFilterer{contract: contract}}, nil
}

// NewOneStepProofEntryCaller creates a new read-only instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofEntryCaller, error) {
	contract, err := bindOneStepProofEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryCaller{contract: contract}, nil
}

// NewOneStepProofEntryTransactor creates a new write-only instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofEntryTransactor, error) {
	contract, err := bindOneStepProofEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryTransactor{contract: contract}, nil
}

// NewOneStepProofEntryFilterer creates a new log filterer instance of OneStepProofEntry, bound to a specific deployed contract.
func NewOneStepProofEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofEntryFilterer, error) {
	contract, err := bindOneStepProofEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryFilterer{contract: contract}, nil
}

// bindOneStepProofEntry binds a generic wrapper to an already deployed contract.
func bindOneStepProofEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProofEntryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntry *OneStepProofEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntry.Contract.OneStepProofEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntry *OneStepProofEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.OneStepProofEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntry *OneStepProofEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.OneStepProofEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntry *OneStepProofEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntry *OneStepProofEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntry *OneStepProofEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntry.Contract.contract.Transact(opts, method, params...)
}

// GetEndMachineHash is a free data retrieval call binding the contract method 0xd8558b87.
//
// Solidity: function getEndMachineHash(uint8 status, bytes32 globalStateHash) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntryCaller) GetEndMachineHash(opts *bind.CallOpts, status uint8, globalStateHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "getEndMachineHash", status, globalStateHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEndMachineHash is a free data retrieval call binding the contract method 0xd8558b87.
//
// Solidity: function getEndMachineHash(uint8 status, bytes32 globalStateHash) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntrySession) GetEndMachineHash(status uint8, globalStateHash [32]byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.GetEndMachineHash(&_OneStepProofEntry.CallOpts, status, globalStateHash)
}

// GetEndMachineHash is a free data retrieval call binding the contract method 0xd8558b87.
//
// Solidity: function getEndMachineHash(uint8 status, bytes32 globalStateHash) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) GetEndMachineHash(status uint8, globalStateHash [32]byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.GetEndMachineHash(&_OneStepProofEntry.CallOpts, status, globalStateHash)
}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntryCaller) GetStartMachineHash(opts *bind.CallOpts, globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "getStartMachineHash", globalStateHash, wasmModuleRoot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntrySession) GetStartMachineHash(globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.GetStartMachineHash(&_OneStepProofEntry.CallOpts, globalStateHash, wasmModuleRoot)
}

// GetStartMachineHash is a free data retrieval call binding the contract method 0x04997be4.
//
// Solidity: function getStartMachineHash(bytes32 globalStateHash, bytes32 wasmModuleRoot) pure returns(bytes32)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) GetStartMachineHash(globalStateHash [32]byte, wasmModuleRoot [32]byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.GetStartMachineHash(&_OneStepProofEntry.CallOpts, globalStateHash, wasmModuleRoot)
}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProveOneStep(opts *bind.CallOpts, execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proveOneStep", execCtx, machineStep, beforeHash, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntrySession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.ProveOneStep(&_OneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// ProveOneStep is a free data retrieval call binding the contract method 0x5d3adcfb.
//
// Solidity: function proveOneStep((uint256,address) execCtx, uint256 machineStep, bytes32 beforeHash, bytes proof) view returns(bytes32 afterHash)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProveOneStep(execCtx ExecutionContext, machineStep *big.Int, beforeHash [32]byte, proof []byte) ([32]byte, error) {
	return _OneStepProofEntry.Contract.ProveOneStep(&_OneStepProofEntry.CallOpts, execCtx, machineStep, beforeHash, proof)
}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) Prover0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "prover0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) Prover0() (common.Address, error) {
	return _OneStepProofEntry.Contract.Prover0(&_OneStepProofEntry.CallOpts)
}

// Prover0 is a free data retrieval call binding the contract method 0x30a5509f.
//
// Solidity: function prover0() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) Prover0() (common.Address, error) {
	return _OneStepProofEntry.Contract.Prover0(&_OneStepProofEntry.CallOpts)
}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverHostIo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverHostIo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverHostIo() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverHostIo(&_OneStepProofEntry.CallOpts)
}

// ProverHostIo is a free data retrieval call binding the contract method 0x5f52fd7c.
//
// Solidity: function proverHostIo() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverHostIo() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverHostIo(&_OneStepProofEntry.CallOpts)
}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverMath(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverMath")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverMath() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMath(&_OneStepProofEntry.CallOpts)
}

// ProverMath is a free data retrieval call binding the contract method 0x66e5d9c3.
//
// Solidity: function proverMath() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverMath() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMath(&_OneStepProofEntry.CallOpts)
}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCaller) ProverMem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneStepProofEntry.contract.Call(opts, &out, "proverMem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntrySession) ProverMem() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMem(&_OneStepProofEntry.CallOpts)
}

// ProverMem is a free data retrieval call binding the contract method 0x1f128bc0.
//
// Solidity: function proverMem() view returns(address)
func (_OneStepProofEntry *OneStepProofEntryCallerSession) ProverMem() (common.Address, error) {
	return _OneStepProofEntry.Contract.ProverMem(&_OneStepProofEntry.CallOpts)
}

// OneStepProofEntryLibMetaData contains all meta data concerning the OneStepProofEntryLib contract.
var OneStepProofEntryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220223fd026c9ded55afa397377c4fd1d09617b50854daf739261c2555ea51d443264736f6c63430008090033",
}

// OneStepProofEntryLibABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProofEntryLibMetaData.ABI instead.
var OneStepProofEntryLibABI = OneStepProofEntryLibMetaData.ABI

// OneStepProofEntryLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProofEntryLibMetaData.Bin instead.
var OneStepProofEntryLibBin = OneStepProofEntryLibMetaData.Bin

// DeployOneStepProofEntryLib deploys a new Ethereum contract, binding an instance of OneStepProofEntryLib to it.
func DeployOneStepProofEntryLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProofEntryLib, error) {
	parsed, err := OneStepProofEntryLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProofEntryLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProofEntryLib{OneStepProofEntryLibCaller: OneStepProofEntryLibCaller{contract: contract}, OneStepProofEntryLibTransactor: OneStepProofEntryLibTransactor{contract: contract}, OneStepProofEntryLibFilterer: OneStepProofEntryLibFilterer{contract: contract}}, nil
}

// OneStepProofEntryLib is an auto generated Go binding around an Ethereum contract.
type OneStepProofEntryLib struct {
	OneStepProofEntryLibCaller     // Read-only binding to the contract
	OneStepProofEntryLibTransactor // Write-only binding to the contract
	OneStepProofEntryLibFilterer   // Log filterer for contract events
}

// OneStepProofEntryLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofEntryLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofEntryLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofEntryLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofEntryLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofEntryLibSession struct {
	Contract     *OneStepProofEntryLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OneStepProofEntryLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofEntryLibCallerSession struct {
	Contract *OneStepProofEntryLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// OneStepProofEntryLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofEntryLibTransactorSession struct {
	Contract     *OneStepProofEntryLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// OneStepProofEntryLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofEntryLibRaw struct {
	Contract *OneStepProofEntryLib // Generic contract binding to access the raw methods on
}

// OneStepProofEntryLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofEntryLibCallerRaw struct {
	Contract *OneStepProofEntryLibCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofEntryLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofEntryLibTransactorRaw struct {
	Contract *OneStepProofEntryLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProofEntryLib creates a new instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLib(address common.Address, backend bind.ContractBackend) (*OneStepProofEntryLib, error) {
	contract, err := bindOneStepProofEntryLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLib{OneStepProofEntryLibCaller: OneStepProofEntryLibCaller{contract: contract}, OneStepProofEntryLibTransactor: OneStepProofEntryLibTransactor{contract: contract}, OneStepProofEntryLibFilterer: OneStepProofEntryLibFilterer{contract: contract}}, nil
}

// NewOneStepProofEntryLibCaller creates a new read-only instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofEntryLibCaller, error) {
	contract, err := bindOneStepProofEntryLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibCaller{contract: contract}, nil
}

// NewOneStepProofEntryLibTransactor creates a new write-only instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofEntryLibTransactor, error) {
	contract, err := bindOneStepProofEntryLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibTransactor{contract: contract}, nil
}

// NewOneStepProofEntryLibFilterer creates a new log filterer instance of OneStepProofEntryLib, bound to a specific deployed contract.
func NewOneStepProofEntryLibFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofEntryLibFilterer, error) {
	contract, err := bindOneStepProofEntryLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofEntryLibFilterer{contract: contract}, nil
}

// bindOneStepProofEntryLib binds a generic wrapper to an already deployed contract.
func bindOneStepProofEntryLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProofEntryLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntryLib *OneStepProofEntryLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.OneStepProofEntryLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProofEntryLib *OneStepProofEntryLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProofEntryLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProofEntryLib *OneStepProofEntryLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProofEntryLib *OneStepProofEntryLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProofEntryLib.Contract.contract.Transact(opts, method, params...)
}

// OneStepProver0MetaData contains all meta data concerning the OneStepProver0 contract.
var OneStepProver0MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612c11806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80633604366f14610030575b600080fd5b61004361003e36600461209a565b61005a565b6040516100519291906122ae565b60405180910390f35b610062611f0b565b61006a611fec565b610073876127d6565b915061008436879003870187612912565b9050600061009560208701876129b4565b905061204461ffff82166100ac57506102eb6102cd565b61ffff8216600114156100c257506102f66102cd565b61ffff8216600f14156100d857506102fd6102cd565b61ffff8216601014156100ee57506103246102cd565b61ffff8216618009141561010557506103c06102cd565b61ffff821661800b141561011c575061043e6102cd565b61ffff821661800c141561013357506104ce6102cd565b61ffff821661800a141561014a57506106116102cd565b61ffff82166011141561016057506107026102cd565b61ffff821661800314156101775750610ae06102cd565b61ffff8216618004141561018e5750610b206102cd565b61ffff8216602014156101a45750610b7e6102cd565b61ffff8216602114156101ba5750610bc06102cd565b61ffff8216602314156101d05750610c056102cd565b61ffff8216602414156101e65750610c2d6102cd565b61ffff821661800214156101fd5750610c5d6102cd565b61ffff8216601a14156102135750610cfa6102cd565b61ffff8216601b14156102295750610d076102cd565b604161ffff8316108015906102435750604461ffff831611155b156102515750610d766102cd565b61ffff8216618005148061026a575061ffff8216618006145b156102785750610e6a6102cd565b61ffff8216618008141561028f5750610f3d6102cd565b60405162461bcd60e51b815260206004820152600e60248201526d494e56414c49445f4f50434f444560901b60448201526064015b60405180910390fd5b6102de84848989898663ffffffff16565b5050965096945050505050565b505060029092525050565b5050505050565b600061030c8660800151610f4c565b805190915061031c908790610fec565b505050505050565b61033b61033086611065565b6020870151906110e2565b600061034a86608001516110ee565b905061036761035c826040015161113a565b6020880151906110e2565b61037761035c826060015161113a565b602084013563ffffffff811681146103a15760405162461bcd60e51b81526004016102c4906129d8565b63ffffffff166101008701525050600061012090940193909352505050565b6103cc61033086611065565b6103dc6103308660e0015161113a565b6103ec6103308560a0015161113a565b6020808401359081901c604082901c156104185760405162461bcd60e51b81526004016102c4906129ff565b63ffffffff90811660e08801521661010086015250506000610120909301929092525050565b61044a61033086611065565b600061045986608001516110ee565b905061046b61035c826040015161113a565b61047b61035c826060015161113a565b6020808501359081901c604082901c156104a75760405162461bcd60e51b81526004016102c4906129ff565b63ffffffff90811660e0890152166101008701525050600061012090940193909352505050565b60008360200135905060006104ee6104e9886020015161116d565b61118c565b90506104f8611fec565b604080516020810190915260608152600061051487878361121d565b90935090506105248787836112e1565b6101608c015191935091506105448363ffffffff8088169087906113bb16565b1461059c5760405162461bcd60e51b815260206004820152602260248201527f43524f53535f4d4f44554c455f494e5445524e414c5f4d4f44554c45535f524f60448201526113d560f21b60648201526084016102c4565b6105b36105a88b611065565b60208c0151906110e2565b6105c36105a88b60e0015161113a565b6105d36105a88a60a0015161113a565b63ffffffff841660e08b015260a08301516105ee9086612a4c565b63ffffffff166101008b0152505060006101209098019790975250505050505050565b61061d61033086611065565b61062d6103308660e0015161113a565b61063d6103308560a0015161113a565b600061064c86608001516110ee565b9050806060015163ffffffff166000141561066b5750600285526102f6565b602084013563ffffffff811681146106c55760405162461bcd60e51b815260206004820152601d60248201527f4241445f43414c4c45525f494e5445524e414c5f43414c4c5f4441544100000060448201526064016102c4565b604082015163ffffffff1660e088015260608201516106e5908290612a4c565b63ffffffff16610100880152505060006101208601525050505050565b6000806107156104e9886020015161116d565b90506000806000808060006107366040518060200160405280606081525090565b6107418b8b87611406565b955093506107508b8b8761146d565b90965094506107608b8b87611489565b9550925061076f8b8b87611406565b9550915061077e8b8b8761146d565b909750945061078e8b8b876112e1565b6040516d21b0b6361034b73234b932b1ba1d60911b60208201526001600160c01b031960c088901b16602e8201526036810189905290965090915060009060560160408051601f19818403018152919052805160209182012091508d013581146108335760405162461bcd60e51b81526020600482015260166024820152754241445f43414c4c5f494e4449524543545f4441544160501b60448201526064016102c4565b610849826001600160401b03871686868c6114bf565b90508d6040015181146108905760405162461bcd60e51b815260206004820152600f60248201526e10905117d51050931154d7d493d3d5608a1b60448201526064016102c4565b826001600160401b03168963ffffffff16106108ba57505060028d52506102f69650505050505050565b505050505060006108db604080518082019091526000808252602082015290565b6040805160208101909152606081526108f58a8a8661146d565b945092506109048a8a86611561565b945091506109138a8a866112e1565b9450905060006109308263ffffffff808b16908790879061165d16565b90508681146109755760405162461bcd60e51b815260206004820152601160248201527010905117d153115351539514d7d493d3d5607a1b60448201526064016102c4565b8584146109a5578d60025b908160038111156109935761099361217f565b815250505050505050505050506102f6565b6004835160068111156109ba576109ba61217f565b14156109c8578d6002610980565b6005835160068111156109dd576109dd61217f565b1415610a3c576020830151985063ffffffff89168914610a375760405162461bcd60e51b81526020600482015260156024820152744241445f46554e435f5245465f434f4e54454e545360581b60448201526064016102c4565b610a74565b60405162461bcd60e51b815260206004820152600d60248201526c4241445f454c454d5f5459504560981b60448201526064016102c4565b5050505050505050610a8861035c87611065565b6000610a9787608001516110ee565b9050610ab4610aa9826040015161113a565b6020890151906110e2565b610ac4610aa9826060015161113a565b5063ffffffff1661010086015260006101208601525050505050565b602083013563ffffffff81168114610b0a5760405162461bcd60e51b81526004016102c4906129d8565b63ffffffff166101209095019490945250505050565b6000610b326104e9876020015161116d565b905063ffffffff81161561031c57602084013563ffffffff81168114610b6a5760405162461bcd60e51b81526004016102c4906129d8565b63ffffffff16610120870152505050505050565b6000610b8d86608001516110ee565b90506000610ba58260200151866020013586866116f9565b6020880151909150610bb790826110e2565b50505050505050565b6000610bcf866020015161116d565b90506000610be087608001516110ee565b9050610bf781602001518660200135848787611791565b602090910152505050505050565b6000610c1b8560000151856020013585856116f9565b602087015190915061031c90826110e2565b6000610c3c866020015161116d565b9050610c5385600001518560200135838686611791565b9094525050505050565b6000610c6c866020015161116d565b90506000610c7d876020015161116d565b90506000610c8e886020015161116d565b905060006040518060800160405280838152602001886020013560001b8152602001610cb98561118c565b63ffffffff168152602001610ccd8661118c565b63ffffffff168152509050610cef818a6080015161182b90919063ffffffff16565b505050505050505050565b61031c856020015161116d565b6000610d196104e9876020015161116d565b90506000610d2a876020015161116d565b90506000610d3b886020015161116d565b905063ffffffff831615610d5d576020880151610d5890826110e2565b610d6c565b6020880151610d6c90836110e2565b5050505050505050565b6000610d8560208501856129b4565b9050600061ffff821660411415610d9e57506000610e21565b61ffff821660421415610db357506001610e21565b61ffff821660431415610dc857506002610e21565b61ffff821660441415610ddd57506003610e21565b60405162461bcd60e51b8152602060048201526019602482015278434f4e53545f505553485f494e56414c49445f4f50434f444560381b60448201526064016102c4565b610bb76040518060400160405280836006811115610e4157610e4161217f565b815260200187602001356001600160401b031681525088602001516110e290919063ffffffff16565b6040805180820190915260008082526020820152618005610e8e60208601866129b4565b61ffff161415610ebc57610ea5866020015161116d565b6060870151909150610eb790826110e2565b61031c565b618006610ecc60208601866129b4565b61ffff161415610ef557610ee3866060015161116d565b6020870151909150610eb790826110e2565b60405162461bcd60e51b815260206004820152601c60248201527f4d4f56455f494e5445524e414c5f494e56414c49445f4f50434f44450000000060448201526064016102c4565b6000610c1b8660200151611912565b610f5461204e565b815151600114610f765760405162461bcd60e51b81526004016102c490612a74565b81518051600090610f8957610f89612a9f565b6020026020010151905060006001600160401b03811115610fac57610fac612426565b604051908082528060200260200182016040528015610fe557816020015b610fd261204e565b815260200190600190039081610fca5790505b5090915290565b6004815160068111156110015761100161217f565b1415611025578160025b9081600381111561101e5761101e61217f565b9052505050565b60068151600681111561103a5761103a61217f565b146110475781600261100b565b611055828260200151611940565b6110615781600261100b565b5050565b60408051808201909152600080825260208201526110dc8261012001518361010001518460e001516040805180820190915260008082526020820152506040805180820182526006815263ffffffff94909416602093841b67ffffffff00000000161791901b63ffffffff60401b16179082015290565b92915050565b81516110619082611982565b6110f661204e565b8151516001146111185760405162461bcd60e51b81526004016102c490612a74565b8151805160009061112b5761112b612a9f565b60200260200101519050919050565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b604080518082019091526000808252602082015281516110dc90611a4b565b602081015160009081835160068111156111a8576111a861217f565b146111df5760405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b60448201526064016102c4565b64010000000081106110dc5760405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b60448201526064016102c4565b611225611fec565b604080516060810182526000808252602082018190529181018290528391906000806000806112558b8b8961146d565b975095506112648b8b89611b54565b975094506112738b8b8961146d565b975093506112828b8b8961146d565b975092506112918b8b8961146d565b975091506112a08b8b89611bcf565b6040805160c081018252988952602089019790975295870194909452506060850191909152608084015263ffffffff1660a083015290969095509350505050565b6040805160208101909152606081528160006112fe868684611489565b92509050600060ff82166001600160401b0381111561131f5761131f612426565b604051908082528060200260200182016040528015611348578160200160208202803683370190505b50905060005b8260ff168160ff16101561139f5761136788888661146d565b838360ff168151811061137c5761137c612a9f565b60200260200101819650828152505050808061139790612ab5565b91505061134e565b5060405180602001604052808281525093505050935093915050565b60006113fc84846113cb85611c2a565b6040518060400160405280601381526020017226b7b23ab6329036b2b935b632903a3932b29d60691b815250611cc3565b90505b9392505050565b600081815b6008811015611464576008836001600160401b0316901b925085858381811061143657611436612a9f565b919091013560f81c9390931792508161144e81612ad5565b925050808061145c90612ad5565b91505061140b565b50935093915050565b6000818161147c868684611dcd565b9097909650945050505050565b60008184848281811061149e5761149e612a9f565b919091013560f81c92508190506114b481612ad5565b915050935093915050565b604051652a30b136329d60d11b60208201526001600160f81b031960f885901b1660268201526001600160c01b031960c084901b166027820152602f81018290526000908190604f01604051602081830303815290604052805190602001209050611556878783604051806040016040528060128152602001712a30b136329036b2b935b632903a3932b29d60711b815250611cc3565b979650505050505050565b604080518082019091526000808252602082015281600085858381811061158a5761158a612a9f565b919091013560f81c91508290506115a081612ad5565b9250506115ab600690565b60068111156115bc576115bc61217f565b60ff168160ff1611156116025760405162461bcd60e51b815260206004820152600e60248201526d4241445f56414c55455f5459504560901b60448201526064016102c4565b600061160f878785611dcd565b809450819250505060405180604001604052808360ff1660068111156116375761163761217f565b60068111156116485761164861217f565b81526020018281525093505050935093915050565b6000808361166a84611e22565b6040516d2a30b136329032b632b6b2b73a1d60911b6020820152602e810192909252604e820152606e016040516020818303038152906040528051906020012090506116ed8686836040518060400160405280601a81526020017f5461626c6520656c656d656e74206d65726b6c6520747265653a000000000000815250611cc3565b9150505b949350505050565b60408051808201909152600080825260208201526000611729604080518082019091526000808252602082015290565b604080516020810190915260608152611743868685611561565b935091506117528686856112e1565b935090506000611763828985611e3f565b90508881146117845760405162461bcd60e51b81526004016102c490612af0565b5090979650505050505050565b60006117ad604080518082019091526000808252602082015290565b60006117c56040518060200160405280606081525090565b6117d0868684611561565b90935091506117e08686846112e1565b9250905060006117f1828a86611e3f565b90508981146118125760405162461bcd60e51b81526004016102c490612af0565b61181d828a8a611e3f565b9a9950505050505050505050565b81515160009061183c906001612b1b565b6001600160401b0381111561185357611853612426565b60405190808252806020026020018201604052801561188c57816020015b61187961204e565b8152602001906001900390816118715790505b50905060005b8351518110156118e85783518051829081106118b0576118b0612a9f565b60200260200101518282815181106118ca576118ca612a9f565b602002602001018190525080806118e090612ad5565b915050611892565b5081818460000151518151811061190157611901612a9f565b602090810291909101015290915250565b6040805180820190915260008082526020820152815151516113ff611938600183612b33565b845190611e7f565b6000606082901c15611954575060006110dc565b5063ffffffff818116610120840152602082901c811661010084015260409190911c1660e090910152600190565b815151600090611993906001612b1b565b6001600160401b038111156119aa576119aa612426565b6040519080825280602002602001820160405280156119ef57816020015b60408051808201909152600080825260208201528152602001906001900390816119c85790505b50905060005b8351518110156118e8578351805182908110611a1357611a13612a9f565b6020026020010151828281518110611a2d57611a2d612a9f565b60200260200101819052508080611a4390612ad5565b9150506119f5565b604080518082019091526000808252602082015281518051611a6f90600190612b33565b81518110611a7f57611a7f612a9f565b6020026020010151905060006001836000015151611a9d9190612b33565b6001600160401b03811115611ab457611ab4612426565b604051908082528060200260200182016040528015611af957816020015b6040805180820190915260008082526020820152815260200190600190039081611ad25790505b50905060005b8151811015610fe5578351805182908110611b1c57611b1c612a9f565b6020026020010151828281518110611b3657611b36612a9f565b60200260200101819052508080611b4c90612ad5565b915050611aff565b60408051606081018252600080825260208201819052918101919091528160008080611b81888886611406565b94509250611b90888886611406565b94509150611b9f88888661146d565b604080516060810182526001600160401b0396871681529490951660208501529383015250969095509350505050565b600081815b60048110156114645760088363ffffffff16901b9250858583818110611bfc57611bfc612a9f565b919091013560f81c93909317925081611c1481612ad5565b9250508080611c2290612ad5565b915050611bd4565b60008160000151611c3e8360200151611eb7565b6040808501516060860151608087015160a08801519351611ca6969594906020016626b7b23ab6329d60c91b81526007810196909652602786019490945260478501929092526067840152608783015260e01b6001600160e01b03191660a782015260ab0190565b604051602081830303815290604052805190602001209050919050565b8160005b855151811015611d8c5760018516611d2857828287600001518381518110611cf157611cf1612a9f565b6020026020010151604051602001611d0b93929190612b4a565b604051602081830303815290604052805190602001209150611d73565b8286600001518281518110611d3f57611d3f612a9f565b602002602001015183604051602001611d5a93929190612b4a565b6040516020818303038152906040528051906020012091505b60019490941c9380611d8481612ad5565b915050611cc7565b5083156116f15760405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b60448201526064016102c4565b600081815b602081101561146457600883901b9250858583818110611df457611df4612a9f565b919091013560f81c93909317925081611e0c81612ad5565b9250508080611e1a90612ad5565b915050611dd2565b600081600001518260200151604051602001611ca6929190612b90565b60006113fc8484611e4f85611e22565b604051806040016040528060128152602001712b30b63ab29036b2b935b632903a3932b29d60711b815250611cc3565b60408051808201909152600080825260208201528251805183908110611ea757611ea7612a9f565b6020026020010151905092915050565b805160208083015160408085015190516626b2b6b7b93c9d60c91b938101939093526001600160c01b031960c094851b811660278501529190931b16602f8201526037810191909152600090605701611ca6565b6040805161018081019091528060008152602001611f4060408051606080820183529181019182529081526000602082015290565b8152604080518082018252600080825260208083019190915283015201611f7e60408051606080820183529181019182529081526000602082015290565b8152602001611fa3604051806040016040528060608152602001600080191681525090565b815260408051808201825260008082526020808301829052840191909152908201819052606082018190526080820181905260a0820181905260c0820181905260e09091015290565b6040805160c081019091526000815260208101612022604080516060810182526000808252602082018190529181019190915290565b8152600060208201819052604082018190526060820181905260809091015290565b61204c612bc5565b565b6040805160c08101825260006080820181815260a08301829052825260208201819052918101829052606081019190915290565b60006040828403121561209457600080fd5b50919050565b6000806000806000808688036101c0808212156120b657600080fd5b6120c08a8a612082565b975060408901356001600160401b03808211156120dc57600080fd5b818b01915082828d0312156120f057600080fd5b819850610100605f198501121561210657600080fd5b60608b01975061211a8c6101608d01612082565b96506101a08b013593508084111561213157600080fd5b838b0193508b601f85011261214557600080fd5b833592508083111561215657600080fd5b505089602082840101111561216a57600080fd5b60208201935080925050509295509295509295565b634e487b7160e01b600052602160045260246000fd5b600481106121a5576121a561217f565b9052565b8051600781106121bb576121bb61217f565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b80841015612218576122048286516121a9565b9382019360019390930192908501906121f1565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b828110156122975784516122638582516121a9565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a09093019260010161224e565b509687015197909601969096525093949350505050565b60006101208083526122c38184018651612195565b60208501516101c061014081818701526122e16102e08701846121c8565b925060408801516101606123018189018380518252602090810151910152565b60608a0151915061011f1980898703016101a08a015261232186846121c8565b955060808b015192508089870301858a01525061233e858361222c565b60a08b015180516101e08b015260208101516102008b0152909550935060c08a015161022089015260e08a015163ffffffff81166102408a015293506101008a015163ffffffff81166102608a015293509489015163ffffffff811661028089015294918901516102a0880152508701516102c08601525091506113ff905060208301848051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a0830152608081015160c083015263ffffffff60a08201511660e08301525050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561245e5761245e612426565b60405290565b604051602081016001600160401b038111828210171561245e5761245e612426565b604051608081016001600160401b038111828210171561245e5761245e612426565b60405161018081016001600160401b038111828210171561245e5761245e612426565b60405160c081016001600160401b038111828210171561245e5761245e612426565b604051606081016001600160401b038111828210171561245e5761245e612426565b604051601f8201601f191681016001600160401b038111828210171561253757612537612426565b604052919050565b80356004811061254e57600080fd5b919050565b60006001600160401b0382111561256c5761256c612426565b5060051b60200190565b60006040828403121561258857600080fd5b61259061243c565b90508135600781106125a157600080fd5b808252506020820135602082015292915050565b600060408083850312156125c857600080fd5b6125d061243c565b915082356001600160401b03808211156125e957600080fd5b818501915060208083880312156125ff57600080fd5b612607612464565b83358381111561261657600080fd5b80850194505087601f85011261262b57600080fd5b8335925061264061263b84612553565b61250f565b83815260069390931b8401820192828101908985111561265f57600080fd5b948301945b84861015612685576126768a87612576565b82529486019490830190612664565b8252508552948501359484019490945250909392505050565b6000604082840312156126b057600080fd5b6126b861243c565b9050813581526020820135602082015292915050565b803563ffffffff8116811461254e57600080fd5b600060408083850312156126f557600080fd5b6126fd61243c565b915082356001600160401b0381111561271557600080fd5b8301601f8101851361272657600080fd5b8035602061273661263b83612553565b82815260a0928302840182019282820191908985111561275557600080fd5b948301945b848610156127be5780868b0312156127725760008081fd5b61277a612486565b6127848b88612576565b81528787013585820152606061279b8189016126ce565b898301526127ab608089016126ce565b908201528352948501949183019161275a565b50808752505080860135818601525050505092915050565b60006101c082360312156127e957600080fd5b6127f16124a8565b6127fa8361253f565b815260208301356001600160401b038082111561281657600080fd5b612822368387016125b5565b6020840152612834366040870161269e565b6040840152608085013591508082111561284d57600080fd5b612859368387016125b5565b606084015260a085013591508082111561287257600080fd5b5061287f368286016126e2565b6080830152506128923660c0850161269e565b60a08201526101008084013560c08301526101206128b18186016126ce565b60e08401526101406128c48187016126ce565b8385015261016092506128d88387016126ce565b91840191909152610180850135908301526101a090930135928101929092525090565b80356001600160401b038116811461254e57600080fd5b600081830361010081121561292657600080fd5b61292e6124cb565b833581526060601f198301121561294457600080fd5b61294c6124ed565b915061295a602085016128fb565b8252612968604085016128fb565b6020830152606084013560408301528160208201526080840135604082015260a0840135606082015260c084013560808201526129a760e085016126ce565b60a0820152949350505050565b6000602082840312156129c657600080fd5b813561ffff811681146113ff57600080fd5b6020808252600d908201526c4241445f43414c4c5f4441544160981b604082015260600190565b6020808252601a908201527f4241445f43524f53535f4d4f44554c455f43414c4c5f44415441000000000000604082015260600190565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff808316818516808303821115612a6b57612a6b612a36565b01949350505050565b6020808252601190820152700848288beae929c889eaebe988a9c8ea89607b1b604082015260600190565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff811415612acc57612acc612a36565b60010192915050565b6000600019821415612ae957612ae9612a36565b5060010190565b60208082526011908201527015d493d391d7d3515492d31157d493d3d5607a1b604082015260600190565b60008219821115612b2e57612b2e612a36565b500190565b600082821015612b4557612b45612a36565b500390565b6000845160005b81811015612b6b5760208188018101518583015201612b51565b81811115612b7a576000828501525b5091909101928352506020820152604001919050565b652b30b63ab29d60d11b8152600060078410612bae57612bae61217f565b5060f89290921b6006830152600782015260270190565b634e487b7160e01b600052605160045260246000fdfea2646970667358221220fc103f19f529f25606dcf95d221632948bea361b65291f312c2bdf94f716271c64736f6c63430008090033",
}

// OneStepProver0ABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProver0MetaData.ABI instead.
var OneStepProver0ABI = OneStepProver0MetaData.ABI

// OneStepProver0Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProver0MetaData.Bin instead.
var OneStepProver0Bin = OneStepProver0MetaData.Bin

// DeployOneStepProver0 deploys a new Ethereum contract, binding an instance of OneStepProver0 to it.
func DeployOneStepProver0(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProver0, error) {
	parsed, err := OneStepProver0MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProver0Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProver0{OneStepProver0Caller: OneStepProver0Caller{contract: contract}, OneStepProver0Transactor: OneStepProver0Transactor{contract: contract}, OneStepProver0Filterer: OneStepProver0Filterer{contract: contract}}, nil
}

// OneStepProver0 is an auto generated Go binding around an Ethereum contract.
type OneStepProver0 struct {
	OneStepProver0Caller     // Read-only binding to the contract
	OneStepProver0Transactor // Write-only binding to the contract
	OneStepProver0Filterer   // Log filterer for contract events
}

// OneStepProver0Caller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProver0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProver0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProver0Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProver0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProver0Session struct {
	Contract     *OneStepProver0   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProver0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProver0CallerSession struct {
	Contract *OneStepProver0Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OneStepProver0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProver0TransactorSession struct {
	Contract     *OneStepProver0Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OneStepProver0Raw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProver0Raw struct {
	Contract *OneStepProver0 // Generic contract binding to access the raw methods on
}

// OneStepProver0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProver0CallerRaw struct {
	Contract *OneStepProver0Caller // Generic read-only contract binding to access the raw methods on
}

// OneStepProver0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProver0TransactorRaw struct {
	Contract *OneStepProver0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProver0 creates a new instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0(address common.Address, backend bind.ContractBackend) (*OneStepProver0, error) {
	contract, err := bindOneStepProver0(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0{OneStepProver0Caller: OneStepProver0Caller{contract: contract}, OneStepProver0Transactor: OneStepProver0Transactor{contract: contract}, OneStepProver0Filterer: OneStepProver0Filterer{contract: contract}}, nil
}

// NewOneStepProver0Caller creates a new read-only instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Caller(address common.Address, caller bind.ContractCaller) (*OneStepProver0Caller, error) {
	contract, err := bindOneStepProver0(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Caller{contract: contract}, nil
}

// NewOneStepProver0Transactor creates a new write-only instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Transactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProver0Transactor, error) {
	contract, err := bindOneStepProver0(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Transactor{contract: contract}, nil
}

// NewOneStepProver0Filterer creates a new log filterer instance of OneStepProver0, bound to a specific deployed contract.
func NewOneStepProver0Filterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProver0Filterer, error) {
	contract, err := bindOneStepProver0(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProver0Filterer{contract: contract}, nil
}

// bindOneStepProver0 binds a generic wrapper to an already deployed contract.
func bindOneStepProver0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProver0MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProver0 *OneStepProver0Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProver0.Contract.OneStepProver0Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProver0 *OneStepProver0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProver0.Contract.OneStepProver0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProver0 *OneStepProver0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProver0.Contract.OneStepProver0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProver0 *OneStepProver0CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProver0.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProver0 *OneStepProver0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProver0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProver0 *OneStepProver0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProver0.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0Caller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProver0.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0Session) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProver0.Contract.ExecuteOneStep(&_OneStepProver0.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProver0 *OneStepProver0CallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProver0.Contract.ExecuteOneStep(&_OneStepProver0.CallOpts, arg0, startMach, startMod, inst, proof)
}

// OneStepProverHostIoMetaData contains all meta data concerning the OneStepProverHostIo contract.
var OneStepProverHostIoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"execCtx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50613b99806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80633604366f14610030575b600080fd5b61004361003e366004612e1c565b61005a565b604051610051929190613030565b60405180910390f35b610062612c84565b61006a612d65565b61007387613558565b915061008436879003870187613694565b905060006100956020870187613736565b9050612dbd61801061ffff8316108015906100b6575061801361ffff831611155b156100c457506101e06101c1565b61ffff821661802014156100db57506103286101c1565b61ffff821661802114156100f257506109e56101c1565b61ffff821661802214156101095750610d0a6101c1565b61ffff821661802314156101205750610d166101c1565b61ffff821661802414156101375750610e3c6101c1565b61ffff8216618030141561014e5750610ee46101c1565b61ffff821661803114156101655750610f2b6101c1565b61ffff8216618032141561017c5750610f826101c1565b60405162461bcd60e51b8152602060048201526015602482015274494e56414c49445f4d454d4f52595f4f50434f444560581b60448201526064015b60405180910390fd5b6101d38a85858a8a8a8763ffffffff16565b5050965096945050505050565b60006101ef6020850185613736565b90506101f9612dc7565b6000610206858583610ff6565b60c08a01519193509150610219836110d1565b146102595760405162461bcd60e51b815260206004820152601060248201526f4241445f474c4f42414c5f535441544560801b60448201526064016101b8565b61ffff83166180101480610272575061ffff8316618011145b156102945761028f8888848961028a8987818d61375a565b611152565b61030c565b61ffff831661801214156102ac5761028f88836112d7565b61ffff831661801314156102c45761028f8883611385565b60405162461bcd60e51b815260206004820152601a60248201527f494e56414c49445f474c4f42414c53544154455f4f50434f444500000000000060448201526064016101b8565b610315826110d1565b60c0909801979097525050505050505050565b600061033f61033a87602001516113fb565b611420565b63ffffffff169050600061035961033a88602001516113fb565b63ffffffff16905061036c60208361379a565b1515806103925750602080870151516001600160401b0316906103909083906137c4565b115b806103a657506103a360208261379a565b15155b156103cd578660025b908160038111156103c2576103c2612f01565b8152505050506109dd565b60006103da6020836137dc565b90506000806103f56040518060200160405280606081525090565b60208a015161040790858a8a876114b1565b909450909250905060606000898986818110610425576104256137f0565b919091013560f81c915085905061043b81613806565b95505060208b01356105165760ff81166104fe5736600061045e8b88818f61375a565b91509150858282604051610473929190613821565b6040518091039020146104985760405162461bcd60e51b81526004016101b890613831565b60006104a58b60206137c4565b9050818111156104b25750805b6104be818c848661375a565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092975061096195505050505050565b60405162461bcd60e51b81526004016101b890613857565b8a60200135600114156105c75760ff8116156105445760405162461bcd60e51b81526004016101b890613857565b3660006105538b88818f61375a565b91509150856002838360405161056a929190613821565b602060405180830381855afa158015610587573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906105aa9190613887565b146104985760405162461bcd60e51b81526004016101b890613831565b8a60200135600214156109215760ff8116156105f55760405162461bcd60e51b81526004016101b890613857565b3660006106048b88818f61375a565b90925090508561061860206000848661375a565b610621916138a0565b146106655760405162461bcd60e51b8152602060048201526014602482015273096b48ebea0a49e9e8cbeaea49e9c8ebe9082a6960631b60448201526064016101b8565b600080600080600a6001600160a01b03168686604051610686929190613821565b600060405180830381855afa9150503d80600081146106c1576040519150601f19603f3d011682016040523d82523d6000602084013e6106c6565b606091505b50915091508161070c5760405162461bcd60e51b815260206004820152601160248201527024a72b20a624a22fa5ad23afa82927a7a360791b60448201526064016101b8565b60008151116107565760405162461bcd60e51b81526020600482015260166024820152754b5a475f505245434f4d50494c455f4d495353494e4760501b60448201526064016101b8565b8080602001905181019061076a91906138be565b9094509250507f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001821490506107d75760405162461bcd60e51b8152602060048201526013602482015272554e4b4e4f574e5f424c535f4d4f44554c555360681b60448201526064016101b8565b6107e28260206138e2565b8c1015610918576000806107f760208f6137dc565b905060015b8481101561082657600192831b92828116141561081a576001831792505b600191821c911b6107fc565b506000610838856401000000006137dc565b905061084483826138e2565b905060006108737f16a2a19edfe81f20d09b681922c813b4b63683508c2280b93829971f439f0d2b838761154b565b905080610884604060208a8c61375a565b61088d916138a0565b146108ce5760405162461bcd60e51b815260206004820152601160248201527025ad23afa82927a7a32faba927a723afad60791b60448201526064016101b8565b6108dc60606040898b61375a565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929c50505050505050505b50505050610961565b60405162461bcd60e51b8152602060048201526015602482015274554e4b4e4f574e5f505245494d4147455f5459504560581b60448201526064016101b8565b60005b82518110156109a5576109918582858481518110610984576109846137f0565b016020015160f81c61167f565b94508061099d81613806565b915050610964565b506109b1838786611704565b60208d01516040015281516109d4906109c990611783565b60208f0151906117b6565b50505050505050505b505050505050565b60006109f761033a87602001516113fb565b63ffffffff1690506000610a1161033a88602001516113fb565b63ffffffff1690506000610a30610a2b89602001516113fb565b6117c6565b6001600160401b031690506020860135158015610a4e575088358110155b15610a76578760035b90816003811115610a6a57610a6a612f01565b815250505050506109dd565b602080880151516001600160401b031690610a929084906137c4565b1180610aa75750610aa460208361379a565b15155b15610ab457876002610a57565b6000610ac16020846137dc565b9050600080610adc6040518060200160405280606081525090565b60208b0151610aee90858b8b876114b1565b9094509092509050888884818110610b0857610b086137f0565b909101356001600160f81b031916159050610b5b5760405162461bcd60e51b81526020600482015260136024820152722aa725a727aba72fa4a72127ac2fa82927a7a360691b60448201526064016101b8565b82610b6581613806565b935050612dbd6000808c602001351415610b83576118579150610bc3565b60018c602001351415610b9a57611b589150610bc3565b8d60025b90816003811115610bb157610bb1612f01565b815250505050505050505050506109dd565b610be38f888d8d89908092610bda9392919061375a565b8663ffffffff16565b905080610bf2578d6002610b9e565b505082881015610c385760405162461bcd60e51b81526020600482015260116024820152702120a22fa6a2a9a9a0a3a2afa82927a7a360791b60448201526064016101b8565b6000610c44848a613901565b905060005b60208163ffffffff16108015610c6d575081610c6b63ffffffff83168b6137c4565b105b15610cc657610cb28463ffffffff83168d8d82610c8a8f8c6137c4565b610c9491906137c4565b818110610ca357610ca36137f0565b919091013560f81c905061167f565b935080610cbe81613918565b915050610c49565b610cd1838786611704565b60208e015160400152610cf9610ce682611783565b8f602001516117b690919063ffffffff16565b505050505050505050505050505050565b50506001909252505050565b60006040518060400160405280601381526020017226b7b23ab6329036b2b935b632903a3932b29d60691b8152509050600086610160015190506000610d6261033a89602001516113fb565b63ffffffff169050610d81818860200151611dfa90919063ffffffff16565b610d8d57876002610a57565b600080610dad610d9e6020856137dc565b60208b015190898960006114b1565b5091509150600080610dc18c848b8b611e2f565b92505091506000610ddd836001610dd891906137c4565b61200f565b90508015610e0857610dfd87610df48560016137c4565b8760008c61202f565b6101608e0152610e26565b610e1f610e168460016137c4565b8390878b6120d9565b6101608e01525b6109d46109c9610e378560016137c4565b611783565b60408051808201909152601381527226b7b23ab6329036b2b935b632903a3932b29d60691b6020820152600080610e7588828787611e2f565b50915091506000610e858361200f565b90508015610ec45781518051610e9d90600190613901565b81518110610ead57610ead6137f0565b602002602001015189610160018181525050610ed8565b610ed182846000876120d9565b6101608a01525b50505050505050505050565b61014085015160001914610f11578460025b90816003811115610f0957610f09612f01565b9052506109dd565b610f1e8560a001516121e3565b6109dd85604001516121e3565b61014085015160001914610f4157846002610ef6565b60a0850151516000191415610f5857846002610ef6565b610f6785604001518383612226565b60a08501516109dd90610f7d836040818761375a565b612226565b60a0850151516000191415610f9957846002610ef6565b6020830135610fc7576101408501516000191415610fb957846002610ef6565b600019610140860152610fed565b61014085015160001914610fdd57846002610ef6565b610feb856020850135612314565b505b6109dd85612387565b610ffe612dc7565b81611007612dec565b61100f612dec565b60005b600260ff8216101561105a5761102988888661240b565b848360ff166002811061103e5761103e6137f0565b60200201919091529350806110528161393c565b915050611012565b5060005b600260ff821610156110b457611075888886612427565b838360ff166002811061108a5761108a6137f0565b6001600160401b0390931660209390930201919091529350806110ac8161393c565b91505061105e565b506040805180820190915291825260208201529590945092505050565b80518051602091820151828401518051908401516040516c23b637b130b61039ba30ba329d60991b95810195909552602d850193909352604d8401919091526001600160c01b031960c091821b8116606d85015291901b166075820152600090607d015b604051602081830303815290604052805190602001209050919050565b600061116461033a88602001516113fb565b63ffffffff169050600061117e61033a89602001516113fb565b9050600263ffffffff821610611196578760026103af565b60208701516111a59083611dfa565b6111b1578760026103af565b60006111be6020846137dc565b90506000806111d96040518060200160405280606081525090565b60208b01516111eb90858a8a876114b1565b909450909250905061801061120360208b018b613736565b61ffff1614156112485761123a848b600001518763ffffffff166002811061122d5761122d6137f0565b6020020151839190611704565b60208c0151604001526112c9565b61801161125860208b018b613736565b61ffff161415611287578951829063ffffffff87166002811061127d5761127d6137f0565b60200201526112c9565b60405162461bcd60e51b81526020600482015260176024820152764241445f474c4f42414c5f53544154455f4f50434f444560481b60448201526064016101b8565b505050505050505050505050565b60006112e961033a84602001516113fb565b9050600263ffffffff821610611318578260025b9081600381111561131057611310612f01565b905250505050565b61138061137583602001518363ffffffff166002811061133a5761133a6137f0565b602002015160408051808201909152600080825260208201525060408051808201909152600181526001600160401b03909116602082015290565b6020850151906117b6565b505050565b6000611397610a2b84602001516113fb565b905060006113ab61033a85602001516113fb565b9050600263ffffffff8216106113c5575050600290915250565b8183602001518263ffffffff16600281106113e2576113e26137f0565b6001600160401b03909216602092909202015250505050565b6040805180820190915260008082526020820152815161141a9061248e565b92915050565b6020810151600090818351600681111561143c5761143c612f01565b146114735760405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b60448201526064016101b8565b640100000000811061141a5760405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b60448201526064016101b8565b6000806114ca6040518060200160405280606081525090565b8391506114d886868461240b565b90935091506114e886868461259e565b9250905060006114f9828986611704565b90508860400151811461153f5760405162461bcd60e51b815260206004820152600e60248201526d15d493d391d7d3515357d493d3d560921b60448201526064016101b8565b50955095509592505050565b60408051602080820181905281830181905260608201526080810185905260a0810184905260c08082018490528251808303909101815260e090910191829052600091829081906005906115a090859061398c565b600060405180830381855afa9150503d80600081146115db576040519150601f19603f3d011682016040523d82523d6000602084013e6115e0565b606091505b5091509150816116225760405162461bcd60e51b815260206004820152600d60248201526c1353d111561417d19052531151609a1b60448201526064016101b8565b80516020146116695760405162461bcd60e51b815260206004820152601360248201527209a9e888ab0a0beaea49e9c8ebe988a9c8ea89606b1b60448201526064016101b8565b611672816139a8565b93505050505b9392505050565b6000602083106116c95760405162461bcd60e51b81526020600482015260156024820152740848288bea68aa8be988a828cbe84b2a88abe9288b605b1b60448201526064016101b8565b6000836116d860016020613901565b6116e29190613901565b6116ed9060086138e2565b60ff848116821b911b198616179150509392505050565b6040516b26b2b6b7b93c903632b0b31d60a11b6020820152602c81018290526000908190604c0160405160208183030381529060405280519060200120905061177a8585836040518060400160405280601381526020017226b2b6b7b93c9036b2b935b632903a3932b29d60691b8152506120d9565b95945050505050565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b81516117c29082612678565b5050565b60208101516000906001835160068111156117e3576117e3612f01565b1461181a5760405162461bcd60e51b81526020600482015260076024820152661393d517d24d8d60ca1b60448201526064016101b8565b600160401b811061141a5760405162461bcd60e51b815260206004820152600760248201526610905117d24d8d60ca1b60448201526064016101b8565b6000602882101561189f5760405162461bcd60e51b81526020600482015260126024820152712120a22fa9a2a8a4a72127ac2fa82927a7a360711b60448201526064016101b8565b60006118ad84846020612427565b5080915050600084846040516118c4929190613821565b60405190819003902090506000806001600160401b03881615611984576118f160408a0160208b016139cc565b6001600160a01b03166316bf557961190a60018b6139f5565b6040516001600160e01b031960e084901b1681526001600160401b03909116600482015260240160206040518083038186803b15801561194957600080fd5b505afa15801561195d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119819190613887565b91505b6001600160401b03841615611a36576119a360408a0160208b016139cc565b6001600160a01b031663d5719dc26119bc6001876139f5565b6040516001600160e01b031960e084901b1681526001600160401b03909116600482015260240160206040518083038186803b1580156119fb57600080fd5b505afa158015611a0f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a339190613887565b90505b604080516020810184905290810184905260608101829052600090608001604051602081830303815290604052805190602001209050896020016020810190611a7f91906139cc565b6040516316bf557960e01b81526001600160401b038b1660048201526001600160a01b0391909116906316bf55799060240160206040518083038186803b158015611ac957600080fd5b505afa158015611add573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b019190613887565b8114611b465760405162461bcd60e51b81526020600482015260146024820152734241445f534551494e424f585f4d45535341474560601b60448201526064016101b8565b6001955050505050505b949350505050565b60006071821015611b9f5760405162461bcd60e51b81526020600482015260116024820152702120a22fa222a620aca2a22fa82927a7a360791b60448201526064016101b8565b60006001600160401b03851615611c5357611bc060408701602088016139cc565b6001600160a01b031663d5719dc2611bd96001886139f5565b6040516001600160e01b031960e084901b1681526001600160401b03909116600482015260240160206040518083038186803b158015611c1857600080fd5b505afa158015611c2c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c509190613887565b90505b6000611c62846071818861375a565b604051611c70929190613821565b60405180910390209050600085856000818110611c8f57611c8f6137f0565b9050013560f81c60f81b90506000611ca98787600161276b565b50905060008282611cbe607160218b8d61375a565b87604051602001611cd3959493929190613a1d565b60408051601f19818403018152828252805160209182012083820189905283830181905282518085038401815260609094019092528251920191909120909150611d2360408c0160208d016139cc565b604051636ab8cee160e11b81526001600160401b038c1660048201526001600160a01b03919091169063d5719dc29060240160206040518083038186803b158015611d6d57600080fd5b505afa158015611d81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611da59190613887565b8114611de95760405162461bcd60e51b81526020600482015260136024820152724241445f44454c415945445f4d45535341474560681b60448201526064016101b8565b5060019a9950505050505050505050565b81516000906001600160401b0316611e138360206137c4565b111580156116785750611e2760208361379a565b159392505050565b6000611e476040518060200160405280606081525090565b60408051602081019091526060815260408051808201909152601381527226b7b23ab6329036b2b935b632903a3932b29d60691b6020820152610160880151611e8e612d65565b6000611e9b89898c6127c0565b9a509150611eaa89898c612884565b9a509050611eb989898c61259e565b9a5063ffffffff8083169850909650600090611edb9088908a9086906128df16565b9050838114611f225760405162461bcd60e51b81526020600482015260136024820152722ba927a723afa927a7aa2fa327a92fa622a0a360691b60448201526064016101b8565b5050506000611f37866001610dd891906137c4565b90508015611f9057611f4a8660016137c4565b8551516001901b14611f8b5760405162461bcd60e51b815260206004820152600a6024820152692ba927a723afa622a0a360b11b60448201526064016101b8565b612002565b611f9b88888b61259e565b995093506000611fb9611faf8860016137c4565b86906000876120d9565b90508281146120005760405162461bcd60e51b815260206004820152601360248201527257524f4e475f524f4f545f464f525f5a45524f60681b60448201526064016101b8565b505b5050509450945094915050565b6000811580159061141a5750612026600183613901565b82161592915050565b600083855b60018111156120a15783828660405160200161205293929190613a5c565b60405160208183030381529060405280519060200120915083858660405160200161207f93929190613a5c565b60408051601f198184030181529190528051602090910120945060011c612034565b8388836040516020016120b693929190613a5c565b604051602081830303815290604052805190602001209250505095945050505050565b8160005b8551518110156121a2576001851661213e57828287600001518381518110612107576121076137f0565b602002602001015160405160200161212193929190613a5c565b604051602081830303815290604052805190602001209150612189565b8286600001518281518110612155576121556137f0565b60200260200101518360405160200161217093929190613a5c565b6040516020818303038152906040528051906020012091505b60019490941c938061219a81613806565b9150506120dd565b508315611b505760405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b60448201526064016101b8565b805160001914612220578051602080830151604051612203939201613a83565b60408051601f198184030181529190528051602091820120908201525b60009052565b6000808061223585858561240b565b9350915061224485858561240b565b935090506000198214156122975780156122705760405162461bcd60e51b81526004016101b890613aa5565b6020860151156122925760405162461bcd60e51b81526004016101b890613aa5565b612307565b856020015182826040516020016122af929190613a83565b60405160208183030381529060405280519060200120146123075760405162461bcd60e51b8152602060048201526012602482015271057524f4e475f434f5448524541445f504f560741b60448201526064016101b8565b6020860152909352505050565b6101408201516000906000191461232d5750600061141a565b600060408460e0015163ffffffff16901b9050602084610100015163ffffffff16901b811790506001838561012001516123679190613ad3565b6123719190613afb565b63ffffffff161761014084015250600192915050565b60408101515160a0820151516000198114806123a4575060001982145b156123b1578260026112fd565b6123be8360800151612920565b60a08401515260208301516123d2906129b9565b60408401515260808301516123ed9082602082015260609052565b50602091820151808301919091526040805192830190526060825252565b6000818161241a86868461276b565b9097909650945050505050565b600081815b6008811015612485576008836001600160401b0316901b9250858583818110612457576124576137f0565b919091013560f81c9390931792508161246f81613806565b925050808061247d90613806565b91505061242c565b50935093915050565b6040805180820190915260008082526020820152815180516124b290600190613901565b815181106124c2576124c26137f0565b60200260200101519050600060018360000151516124e09190613901565b6001600160401b038111156124f7576124f76131a8565b60405190808252806020026020018201604052801561253c57816020015b60408051808201909152600080825260208201528152602001906001900390816125155790505b50905060005b815181101561259757835180518290811061255f5761255f6137f0565b6020026020010151828281518110612579576125796137f0565b6020026020010181905250808061258f90613806565b915050612542565b5090915290565b6040805160208101909152606081528160006125bb868684612a3e565b92509050600060ff82166001600160401b038111156125dc576125dc6131a8565b604051908082528060200260200182016040528015612605578160200160208202803683370190505b50905060005b8260ff168160ff16101561265c5761262488888661240b565b838360ff1681518110612639576126396137f0565b6020026020010181965082815250505080806126549061393c565b91505061260b565b5060405180602001604052808281525093505050935093915050565b8151516000906126899060016137c4565b6001600160401b038111156126a0576126a06131a8565b6040519080825280602002602001820160405280156126e557816020015b60408051808201909152600080825260208201528152602001906001900390816126be5790505b50905060005b835151811015612741578351805182908110612709576127096137f0565b6020026020010151828281518110612723576127236137f0565b6020026020010181905250808061273990613806565b9150506126eb565b5081818460000151518151811061275a5761275a6137f0565b602090810291909101015290915250565b600081815b602081101561248557600883901b9250858583818110612792576127926137f0565b919091013560f81c939093179250816127aa81613806565b92505080806127b890613806565b915050612770565b6127c8612d65565b604080516060810182526000808252602082018190529181018290528391906000806000806127f88b8b8961240b565b975095506128078b8b89612a74565b975094506128168b8b8961240b565b975093506128258b8b8961240b565b975092506128348b8b8961240b565b975091506128438b8b89612884565b6040805160c081018252988952602089019790975295870194909452506060850191909152608084015263ffffffff1660a083015290969095509350505050565b600081815b60048110156124855760088363ffffffff16901b92508585838181106128b1576128b16137f0565b919091013560f81c939093179250816128c981613806565b92505080806128d790613806565b915050612889565b6000611b5084846128ef85612aef565b6040518060400160405280601381526020017226b7b23ab6329036b2b935b632903a3932b29d60691b8152506120d9565b602081015160005b8251518110156129b3576129588360000151828151811061294b5761294b6137f0565b6020026020010151612b6b565b6040517129ba30b1b590333930b6b29039ba30b1b59d60711b602082015260328101919091526052810183905260720160405160208183030381529060405280519060200120915080806129ab90613806565b915050612928565b50919050565b60208101518151515160005b81811015612a375783516129e2906129dd9083612bdb565b612c13565b6040516b2b30b63ab29039ba30b1b59d60a11b6020820152602c810191909152604c8101849052606c016040516020818303038152906040528051906020012092508080612a2f90613806565b9150506129c5565b5050919050565b600081848482818110612a5357612a536137f0565b919091013560f81c9250819050612a6981613806565b915050935093915050565b60408051606081018252600080825260208201819052918101919091528160008080612aa1888886612427565b94509250612ab0888886612427565b94509150612abf88888661240b565b604080516060810182526001600160401b0396871681529490951660208501529383015250969095509350505050565b60008160000151612b038360200151612c30565b6040808501516060860151608087015160a08801519351611135969594906020016626b7b23ab6329d60c91b81526007810196909652602786019490945260478501929092526067840152608783015260e01b6001600160e01b03191660a782015260ab0190565b6000612b7a8260000151612c13565b602080840151604080860151606087015191516b29ba30b1b590333930b6b29d60a11b94810194909452602c840194909452604c8301919091526001600160e01b031960e093841b8116606c840152921b9091166070820152607401611135565b60408051808201909152600080825260208201528251805183908110612c0357612c036137f0565b6020026020010151905092915050565b600081600001518260200151604051602001611135929190613b18565b805160208083015160408085015190516626b2b6b7b93c9d60c91b938101939093526001600160c01b031960c094851b811660278501529190931b16602f8201526037810191909152600090605701611135565b6040805161018081019091528060008152602001612cb960408051606080820183529181019182529081526000602082015290565b8152604080518082018252600080825260208083019190915283015201612cf760408051606080820183529181019182529081526000602082015290565b8152602001612d1c604051806040016040528060608152602001600080191681525090565b815260408051808201825260008082526020808301829052840191909152908201819052606082018190526080820181905260a0820181905260c0820181905260e09091015290565b6040805160c081019091526000815260208101612d9b604080516060810182526000808252602082018190529181019190915290565b8152600060208201819052604082018190526060820181905260809091015290565b612dc5613b4d565b565b6040518060400160405280612dda612dec565b8152602001612de7612dec565b905290565b60405180604001604052806002906020820280368337509192915050565b6000604082840312156129b357600080fd5b6000806000806000808688036101c080821215612e3857600080fd5b612e428a8a612e0a565b975060408901356001600160401b0380821115612e5e57600080fd5b818b01915082828d031215612e7257600080fd5b819850610100605f1985011215612e8857600080fd5b60608b019750612e9c8c6101608d01612e0a565b96506101a08b0135935080841115612eb357600080fd5b838b0193508b601f850112612ec757600080fd5b8335925080831115612ed857600080fd5b5050896020828401011115612eec57600080fd5b60208201935080925050509295509295509295565b634e487b7160e01b600052602160045260246000fd5b60048110612f2757612f27612f01565b9052565b805160078110612f3d57612f3d612f01565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b80841015612f9a57612f86828651612f2b565b938201936001939093019290850190612f73565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b82811015613019578451612fe5858251612f2b565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a090930192600101612fd0565b509687015197909601969096525093949350505050565b60006101208083526130458184018651612f17565b60208501516101c061014081818701526130636102e0870184612f4a565b925060408801516101606130838189018380518252602090810151910152565b60608a0151915061011f1980898703016101a08a01526130a38684612f4a565b955060808b015192508089870301858a0152506130c08583612fae565b60a08b015180516101e08b015260208101516102008b0152909550935060c08a015161022089015260e08a015163ffffffff81166102408a015293506101008a015163ffffffff81166102608a015293509489015163ffffffff811661028089015294918901516102a0880152508701516102c0860152509150611678905060208301848051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a0830152608081015160c083015263ffffffff60a08201511660e08301525050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156131e0576131e06131a8565b60405290565b604051602081016001600160401b03811182821017156131e0576131e06131a8565b604051608081016001600160401b03811182821017156131e0576131e06131a8565b60405161018081016001600160401b03811182821017156131e0576131e06131a8565b60405160c081016001600160401b03811182821017156131e0576131e06131a8565b604051606081016001600160401b03811182821017156131e0576131e06131a8565b604051601f8201601f191681016001600160401b03811182821017156132b9576132b96131a8565b604052919050565b8035600481106132d057600080fd5b919050565b60006001600160401b038211156132ee576132ee6131a8565b5060051b60200190565b60006040828403121561330a57600080fd5b6133126131be565b905081356007811061332357600080fd5b808252506020820135602082015292915050565b6000604080838503121561334a57600080fd5b6133526131be565b915082356001600160401b038082111561336b57600080fd5b8185019150602080838803121561338157600080fd5b6133896131e6565b83358381111561339857600080fd5b80850194505087601f8501126133ad57600080fd5b833592506133c26133bd846132d5565b613291565b83815260069390931b840182019282810190898511156133e157600080fd5b948301945b84861015613407576133f88a876132f8565b825294860194908301906133e6565b8252508552948501359484019490945250909392505050565b60006040828403121561343257600080fd5b61343a6131be565b9050813581526020820135602082015292915050565b803563ffffffff811681146132d057600080fd5b6000604080838503121561347757600080fd5b61347f6131be565b915082356001600160401b0381111561349757600080fd5b8301601f810185136134a857600080fd5b803560206134b86133bd836132d5565b82815260a092830284018201928282019190898511156134d757600080fd5b948301945b848610156135405780868b0312156134f45760008081fd5b6134fc613208565b6135068b886132f8565b81528787013585820152606061351d818901613450565b8983015261352d60808901613450565b90820152835294850194918301916134dc565b50808752505080860135818601525050505092915050565b60006101c0823603121561356b57600080fd5b61357361322a565b61357c836132c1565b815260208301356001600160401b038082111561359857600080fd5b6135a436838701613337565b60208401526135b63660408701613420565b604084015260808501359150808211156135cf57600080fd5b6135db36838701613337565b606084015260a08501359150808211156135f457600080fd5b5061360136828601613464565b6080830152506136143660c08501613420565b60a08201526101008084013560c0830152610120613633818601613450565b60e0840152610140613646818701613450565b83850152610160925061365a838701613450565b91840191909152610180850135908301526101a090930135928101929092525090565b80356001600160401b03811681146132d057600080fd5b60008183036101008112156136a857600080fd5b6136b061324d565b833581526060601f19830112156136c657600080fd5b6136ce61326f565b91506136dc6020850161367d565b82526136ea6040850161367d565b6020830152606084013560408301528160208201526080840135604082015260a0840135606082015260c0840135608082015261372960e08501613450565b60a0820152949350505050565b60006020828403121561374857600080fd5b813561ffff8116811461167857600080fd5b6000808585111561376a57600080fd5b8386111561377757600080fd5b5050820193919092039150565b634e487b7160e01b600052601260045260246000fd5b6000826137a9576137a9613784565b500690565b634e487b7160e01b600052601160045260246000fd5b600082198211156137d7576137d76137ae565b500190565b6000826137eb576137eb613784565b500490565b634e487b7160e01b600052603260045260246000fd5b600060001982141561381a5761381a6137ae565b5060010190565b8183823760009101908152919050565b6020808252600c908201526b4241445f505245494d41474560a01b604082015260600190565b6020808252601690820152752aa725a727aba72fa82922a4a6a0a3a2afa82927a7a360511b604082015260600190565b60006020828403121561389957600080fd5b5051919050565b8035602083101561141a57600019602084900360031b1b1692915050565b600080604083850312156138d157600080fd5b505080516020909101519092909150565b60008160001904831182151516156138fc576138fc6137ae565b500290565b600082821015613913576139136137ae565b500390565b600063ffffffff80831681811415613932576139326137ae565b6001019392505050565b600060ff821660ff811415613953576139536137ae565b60010192915050565b60005b8381101561397757818101518382015260200161395f565b83811115613986576000848401525b50505050565b6000825161399e81846020870161395c565b9190910192915050565b805160208083015191908110156129b35760001960209190910360031b1b16919050565b6000602082840312156139de57600080fd5b81356001600160a01b038116811461167857600080fd5b60006001600160401b0383811690831681811015613a1557613a156137ae565b039392505050565b6001600160f81b031986168152606085901b6bffffffffffffffffffffffff191660018201528284601583013760159201918201526035019392505050565b60008451613a6e81846020890161395c565b91909101928352506020820152604001919050565b6831b7ba343932b0b21d60b91b81526009810192909252602982015260490190565b60208082526014908201527357524f4e475f434f5448524541445f454d50545960601b604082015260600190565b600063ffffffff808316818516808303821115613af257613af26137ae565b01949350505050565b600063ffffffff83811690831681811015613a1557613a156137ae565b652b30b63ab29d60d11b8152600060078410613b3657613b36612f01565b5060f89290921b6006830152600782015260270190565b634e487b7160e01b600052605160045260246000fdfea26469706673582212205d08bb1a4317b988ca889fd9b2948ff407a0aaf15d9e442789c6ce9d58fd6f6a64736f6c63430008090033",
}

// OneStepProverHostIoABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverHostIoMetaData.ABI instead.
var OneStepProverHostIoABI = OneStepProverHostIoMetaData.ABI

// OneStepProverHostIoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverHostIoMetaData.Bin instead.
var OneStepProverHostIoBin = OneStepProverHostIoMetaData.Bin

// DeployOneStepProverHostIo deploys a new Ethereum contract, binding an instance of OneStepProverHostIo to it.
func DeployOneStepProverHostIo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProverHostIo, error) {
	parsed, err := OneStepProverHostIoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverHostIoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverHostIo{OneStepProverHostIoCaller: OneStepProverHostIoCaller{contract: contract}, OneStepProverHostIoTransactor: OneStepProverHostIoTransactor{contract: contract}, OneStepProverHostIoFilterer: OneStepProverHostIoFilterer{contract: contract}}, nil
}

// OneStepProverHostIo is an auto generated Go binding around an Ethereum contract.
type OneStepProverHostIo struct {
	OneStepProverHostIoCaller     // Read-only binding to the contract
	OneStepProverHostIoTransactor // Write-only binding to the contract
	OneStepProverHostIoFilterer   // Log filterer for contract events
}

// OneStepProverHostIoCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverHostIoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverHostIoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverHostIoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverHostIoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverHostIoSession struct {
	Contract     *OneStepProverHostIo // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OneStepProverHostIoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverHostIoCallerSession struct {
	Contract *OneStepProverHostIoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OneStepProverHostIoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverHostIoTransactorSession struct {
	Contract     *OneStepProverHostIoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OneStepProverHostIoRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverHostIoRaw struct {
	Contract *OneStepProverHostIo // Generic contract binding to access the raw methods on
}

// OneStepProverHostIoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverHostIoCallerRaw struct {
	Contract *OneStepProverHostIoCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverHostIoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverHostIoTransactorRaw struct {
	Contract *OneStepProverHostIoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverHostIo creates a new instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIo(address common.Address, backend bind.ContractBackend) (*OneStepProverHostIo, error) {
	contract, err := bindOneStepProverHostIo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIo{OneStepProverHostIoCaller: OneStepProverHostIoCaller{contract: contract}, OneStepProverHostIoTransactor: OneStepProverHostIoTransactor{contract: contract}, OneStepProverHostIoFilterer: OneStepProverHostIoFilterer{contract: contract}}, nil
}

// NewOneStepProverHostIoCaller creates a new read-only instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverHostIoCaller, error) {
	contract, err := bindOneStepProverHostIo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoCaller{contract: contract}, nil
}

// NewOneStepProverHostIoTransactor creates a new write-only instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverHostIoTransactor, error) {
	contract, err := bindOneStepProverHostIo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoTransactor{contract: contract}, nil
}

// NewOneStepProverHostIoFilterer creates a new log filterer instance of OneStepProverHostIo, bound to a specific deployed contract.
func NewOneStepProverHostIoFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverHostIoFilterer, error) {
	contract, err := bindOneStepProverHostIo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverHostIoFilterer{contract: contract}, nil
}

// bindOneStepProverHostIo binds a generic wrapper to an already deployed contract.
func bindOneStepProverHostIo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverHostIoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverHostIo *OneStepProverHostIoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.OneStepProverHostIoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverHostIo *OneStepProverHostIoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverHostIo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverHostIo *OneStepProverHostIoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverHostIo *OneStepProverHostIoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverHostIo.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoCaller) ExecuteOneStep(opts *bind.CallOpts, execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverHostIo.contract.Call(opts, &out, "executeOneStep", execCtx, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoSession) ExecuteOneStep(execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverHostIo.Contract.ExecuteOneStep(&_OneStepProverHostIo.CallOpts, execCtx, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) execCtx, (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) view returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverHostIo *OneStepProverHostIoCallerSession) ExecuteOneStep(execCtx ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverHostIo.Contract.ExecuteOneStep(&_OneStepProverHostIo.CallOpts, execCtx, startMach, startMod, inst, proof)
}

// OneStepProverMathMetaData contains all meta data concerning the OneStepProverMath contract.
var OneStepProverMathMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612469806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80633604366f14610030575b600080fd5b61004361003e3660046118ed565b61005a565b604051610051929190611b01565b60405180910390f35b6100626117ea565b6040805160c081018252600080825282516060808201855282825260208083018490528286018490528401919091529282018190529181018290526080810182905260a08101919091526100b587612024565b91506100c636879003870187612160565b905060006100d76020870187612202565b90506118cb61ffff8216604514806100f3575061ffff82166050145b1561010157506103106102f2565b604661ffff831610801590610129575061011d6009604661223c565b61ffff168261ffff1611155b1561013757506104296102f2565b606761ffff83161080159061015f57506101536002606761223c565b61ffff168261ffff1611155b1561016d575061050c6102f2565b606a61ffff8316108015906101875750607861ffff831611155b1561019557506105746102f2565b605161ffff8316108015906101bd57506101b16009605161223c565b61ffff168261ffff1611155b156101cb57506107616102f2565b607961ffff8316108015906101f357506101e76002607961223c565b61ffff168261ffff1611155b1561020157506107c66102f2565b607c61ffff83161080159061021b5750608a61ffff831611155b1561022957506108196102f2565b61ffff821660a7141561023f57506109d46102f2565b61ffff821660ac1480610256575061ffff821660ad145b1561026457506109f56102f2565b60c061ffff83161080159061027e575060c461ffff831611155b1561028c5750610a496102f2565b60bc61ffff8316108015906102a6575060bf61ffff831611155b156102b45750610c5e6102f2565b60405162461bcd60e51b815260206004820152600e60248201526d494e56414c49445f4f50434f444560901b60448201526064015b60405180910390fd5b61030384848989898663ffffffff16565b5050965096945050505050565b600061031f8660200151610de9565b905060456103306020860186612202565b61ffff1614156103715760008151600681111561034f5761034f6119d2565b1461036c5760405162461bcd60e51b81526004016102e990612262565b6103ee565b60506103806020860186612202565b61ffff1614156103bc5760018151600681111561039f5761039f6119d2565b1461036c5760405162461bcd60e51b81526004016102e990612283565b60405162461bcd60e51b81526020600482015260076024820152662120a22fa2a8ad60c91b60448201526064016102e9565b600081602001516000141561040557506001610409565b5060005b61042061041582610e0e565b602089015190610e41565b50505050505050565b600061044061043b8760200151610de9565b610e51565b9050600061045461043b8860200151610de9565b9050600060466104676020880188612202565b61047191906122a4565b905060008061ffff83166002148061048d575061ffff83166004145b8061049c575061ffff83166006145b806104ab575061ffff83166008145b156104cb576104b984610ec8565b91506104c485610ec8565b90506104d9565b505063ffffffff8083169084165b60006104e6838386610ef4565b90506104ff6104f48261109d565b60208d015190610e41565b5050505050505050505050565b600061051e61043b8760200151610de9565b9050600060676105316020870187612202565b61053b91906122a4565b905060006105518363ffffffff168360206110d0565b905061056a61055f82610e0e565b60208a015190610e41565b5050505050505050565b600061058661043b8760200151610de9565b9050600061059a61043b8860200151610de9565b9050600080606a6105ae6020890189612202565b6105b891906122a4565b90508061ffff16600314156106365763ffffffff841615806105f057508260030b637fffffff191480156105f057508360030b600019145b15610619578860025b9081600381111561060c5761060c6119d2565b815250505050505061075a565b8360030b8360030b8161062e5761062e6122c7565b05915061073e565b8061ffff16600514156106735763ffffffff8416610656578860026105f9565b8360030b8360030b8161066b5761066b6122c7565b07915061073e565b8061ffff16600a14156106935763ffffffff8316601f85161b915061073e565b8061ffff16600c14156106b35763ffffffff8316601f85161c915061073e565b8061ffff16600b14156106d157600383900b601f85161d915061073e565b8061ffff16600d14156106ef576106e88385611294565b915061073e565b8061ffff16600e1415610706576106e883856112d6565b6000806107208563ffffffff168763ffffffff1685611318565b91509150801561073a575050600289525061075a92505050565b5091505b61075561074a83610e0e565b60208b015190610e41565b505050505b5050505050565b60006107786107738760200151610de9565b61149e565b9050600061078c6107738860200151610de9565b90506000605161079f6020880188612202565b6107a991906122a4565b905060006107b8838584610ef4565b905061075561074a8261109d565b60006107d86107738760200151610de9565b9050600060796107eb6020870187612202565b6107f591906122a4565b90506000610805838360406110d0565b63ffffffff16905061056a61055f82611515565b600061082b6107738760200151610de9565b9050600061083f6107738860200151610de9565b9050600080607c6108536020890189612202565b61085d91906122a4565b90508061ffff16600314156108c6576001600160401b038416158061089c57508260070b677fffffffffffffff1914801561089c57508360070b600019145b156108a9578860026105f9565b8360070b8360070b816108be576108be6122c7565b0591506109c8565b8061ffff1660051415610906576001600160401b0384166108e9578860026105f9565b8360070b8360070b816108fe576108fe6122c7565b0791506109c8565b8061ffff16600a1415610929576001600160401b038316603f85161b91506109c8565b8061ffff16600c141561094c576001600160401b038316603f85161c91506109c8565b8061ffff16600b141561096a57600783900b603f85161d91506109c8565b8061ffff16600d141561098857610981838561154b565b91506109c8565b8061ffff16600e141561099f576109818385611599565b60006109ac848684611318565b909350905080156109c6575050600288525061075a915050565b505b61075561074a83611515565b60006109e66107738760200151610de9565b90508061042061041582610e0e565b6000610a0761043b8760200151610de9565b9050600060ac610a1a6020870187612202565b61ffff161415610a3457610a2d82610ec8565b9050610a3d565b5063ffffffff81165b61042061041582611515565b60008060c0610a5b6020870187612202565b61ffff161415610a715750600090506008610b48565b60c1610a806020870187612202565b61ffff161415610a965750600090506010610b48565b60c2610aa56020870187612202565b61ffff161415610abb5750600190506008610b48565b60c3610aca6020870187612202565b61ffff161415610ae05750600190506010610b48565b60c4610aef6020870187612202565b61ffff161415610b055750600190506020610b48565b60405162461bcd60e51b8152602060048201526018602482015277494e56414c49445f455854454e445f53414d455f5459504560401b60448201526064016102e9565b600080836006811115610b5d57610b5d6119d2565b1415610b6e575063ffffffff610b78565b506001600160401b035b6000610b878960200151610de9565b9050836006811115610b9b57610b9b6119d2565b81516006811115610bae57610bae6119d2565b14610bf75760405162461bcd60e51b81526020600482015260196024820152784241445f455854454e445f53414d455f545950455f5459504560381b60448201526064016102e9565b6000610c0a600160ff861681901b6122dd565b602083018051821690529050610c216001856122f4565b60ff166001901b826020015116600014610c4357602082018051821985161790525b60208a0151610c529083610e41565b50505050505050505050565b60008060bc610c706020870187612202565b61ffff161415610c865750600090506002610d33565b60bd610c956020870187612202565b61ffff161415610cab5750600190506003610d33565b60be610cba6020870187612202565b61ffff161415610cd05750600290506000610d33565b60bf610cdf6020870187612202565b61ffff161415610cf55750600390506001610d33565b60405162461bcd60e51b81526020600482015260136024820152721253959053125117d491525395115494149155606a1b60448201526064016102e9565b6000610d428860200151610de9565b9050816006811115610d5657610d566119d2565b81516006811115610d6957610d696119d2565b14610db15760405162461bcd60e51b8152602060048201526018602482015277494e56414c49445f5245494e544552505245545f5459504560401b60448201526064016102e9565b80836006811115610dc457610dc46119d2565b90816006811115610dd757610dd76119d2565b905250602088015161056a9082610e41565b60408051808201909152600080825260208201528151610e08906115e7565b92915050565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b8151610e4d90826116f7565b5050565b60208101516000908183516006811115610e6d57610e6d6119d2565b14610e8a5760405162461bcd60e51b81526004016102e990612262565b6401000000008110610e085760405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b60448201526064016102e9565b60006380000000821615610eea575063ffffffff1667ffffffff000000001790565b5063ffffffff1690565b600061ffff8216610f1b57826001600160401b0316846001600160401b0316149050611096565b61ffff821660011415610f4557826001600160401b0316846001600160401b031614159050611096565b61ffff821660021415610f62578260070b8460070b129050611096565b61ffff821660031415610f8b57826001600160401b0316846001600160401b0316109050611096565b61ffff821660041415610fa8578260070b8460070b139050611096565b61ffff821660051415610fd157826001600160401b0316846001600160401b0316119050611096565b61ffff821660061415610fef578260070b8460070b13159050611096565b61ffff82166007141561101957826001600160401b0316846001600160401b031611159050611096565b61ffff821660081415611037578260070b8460070b12159050611096565b61ffff82166009141561106157826001600160401b0316846001600160401b031610159050611096565b60405162461bcd60e51b815260206004820152600a6024820152690424144204952454c4f560b41b60448201526064016102e9565b9392505050565b604080518082019091526000808252602082015281156110c157610e086001610e0e565b610e086000610e0e565b919050565b60008161ffff16602014806110e957508161ffff166040145b6111305760405162461bcd60e51b8152602060048201526018602482015277057524f4e4720555345204f462067656e65726963556e4f760441b60448201526064016102e9565b61ffff83166111a15761ffff82165b60008163ffffffff16118015611174575061115b600182612317565b63ffffffff166001901b856001600160401b0316166000145b1561118b57611184600182612317565b905061113f565b6111998161ffff8516612317565b915050611096565b61ffff8316600114156111fa5760005b8261ffff168163ffffffff161080156111dc5750600163ffffffff82161b85166001600160401b0316155b156111f3576111ec600182612334565b90506111b1565b9050611096565b61ffff831660021415611260576000805b8361ffff168263ffffffff16101561125757600163ffffffff83161b86166001600160401b03161561124557611242600182612334565b90505b8161124f81612353565b92505061120b565b91506110969050565b60405162461bcd60e51b815260206004820152600960248201526804241442049556e4f760bc1b60448201526064016102e9565b60006112a1602083612377565b91506112ae826020612317565b63ffffffff168363ffffffff16901c8263ffffffff168463ffffffff16901b17905092915050565b60006112e3602083612377565b91506112f0826020612317565b63ffffffff168363ffffffff16901b8263ffffffff168463ffffffff16901c17905092915050565b60008061ffff83166113305750508282016000611496565b8261ffff16600114156113495750508183036000611496565b8261ffff16600214156113625750508282026000611496565b8261ffff16600414156113b6576001600160401b0384166113895750600090506001611496565b836001600160401b0316856001600160401b0316816113aa576113aa6122c7565b04600091509150611496565b8261ffff166006141561140a576001600160401b0384166113dd5750600090506001611496565b836001600160401b0316856001600160401b0316816113fe576113fe6122c7565b06600091509150611496565b8261ffff16600714156114235750508282166000611496565b8261ffff166008141561143c5750508282176000611496565b8261ffff16600914156114555750508282186000611496565b60405162461bcd60e51b81526020600482015260166024820152750494e56414c49445f47454e455249435f42494e5f4f560541b60448201526064016102e9565b935093915050565b60208101516000906001835160068111156114bb576114bb6119d2565b146114d85760405162461bcd60e51b81526004016102e990612283565b600160401b8110610e085760405162461bcd60e51b815260206004820152600760248201526610905117d24d8d60ca1b60448201526064016102e9565b60408051808201909152600080825260208201525060408051808201909152600181526001600160401b03909116602082015290565b600061155860408361239a565b91506115658260406123b4565b6001600160401b0316836001600160401b0316901c826001600160401b0316846001600160401b0316901b17905092915050565b60006115a660408361239a565b91506115b38260406123b4565b6001600160401b0316836001600160401b0316901b826001600160401b0316846001600160401b0316901c17905092915050565b60408051808201909152600080825260208201528151805161160b906001906122dd565b8151811061161b5761161b6123d4565b602002602001015190506000600183600001515161163991906122dd565b6001600160401b0381111561165057611650611c79565b60405190808252806020026020018201604052801561169557816020015b604080518082019091526000808252602082015281526020019060019003908161166e5790505b50905060005b81518110156116f05783518051829081106116b8576116b86123d4565b60200260200101518282815181106116d2576116d26123d4565b602002602001018190525080806116e8906123ea565b91505061169b565b5090915290565b815151600090611708906001612405565b6001600160401b0381111561171f5761171f611c79565b60405190808252806020026020018201604052801561176457816020015b604080518082019091526000808252602082015281526020019060019003908161173d5790505b50905060005b8351518110156117c0578351805182908110611788576117886123d4565b60200260200101518282815181106117a2576117a26123d4565b602002602001018190525080806117b8906123ea565b91505061176a565b508181846000015151815181106117d9576117d96123d4565b602090810291909101015290915250565b604080516101808101909152806000815260200161181f60408051606080820183529181019182529081526000602082015290565b815260408051808201825260008082526020808301919091528301520161185d60408051606080820183529181019182529081526000602082015290565b8152602001611882604051806040016040528060608152602001600080191681525090565b815260408051808201825260008082526020808301829052840191909152908201819052606082018190526080820181905260a0820181905260c0820181905260e09091015290565b6118d361241d565b565b6000604082840312156118e757600080fd5b50919050565b6000806000806000808688036101c08082121561190957600080fd5b6119138a8a6118d5565b975060408901356001600160401b038082111561192f57600080fd5b818b01915082828d03121561194357600080fd5b819850610100605f198501121561195957600080fd5b60608b01975061196d8c6101608d016118d5565b96506101a08b013593508084111561198457600080fd5b838b0193508b601f85011261199857600080fd5b83359250808311156119a957600080fd5b50508960208284010111156119bd57600080fd5b60208201935080925050509295509295509295565b634e487b7160e01b600052602160045260246000fd5b600481106119f8576119f86119d2565b9052565b805160078110611a0e57611a0e6119d2565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b80841015611a6b57611a578286516119fc565b938201936001939093019290850190611a44565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b82811015611aea578451611ab68582516119fc565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a090930192600101611aa1565b509687015197909601969096525093949350505050565b6000610120808352611b1681840186516119e8565b60208501516101c06101408181870152611b346102e0870184611a1b565b92506040880151610160611b548189018380518252602090810151910152565b60608a0151915061011f1980898703016101a08a0152611b748684611a1b565b955060808b015192508089870301858a015250611b918583611a7f565b60a08b015180516101e08b015260208101516102008b0152909550935060c08a015161022089015260e08a015163ffffffff81166102408a015293506101008a015163ffffffff81166102608a015293509489015163ffffffff811661028089015294918901516102a0880152508701516102c0860152509150611096905060208301848051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a0830152608081015160c083015263ffffffff60a08201511660e08301525050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715611cb157611cb1611c79565b60405290565b604051602081016001600160401b0381118282101715611cb157611cb1611c79565b604051608081016001600160401b0381118282101715611cb157611cb1611c79565b60405161018081016001600160401b0381118282101715611cb157611cb1611c79565b60405160c081016001600160401b0381118282101715611cb157611cb1611c79565b604051606081016001600160401b0381118282101715611cb157611cb1611c79565b604051601f8201601f191681016001600160401b0381118282101715611d8a57611d8a611c79565b604052919050565b8035600481106110cb57600080fd5b60006001600160401b03821115611dba57611dba611c79565b5060051b60200190565b600060408284031215611dd657600080fd5b611dde611c8f565b9050813560078110611def57600080fd5b808252506020820135602082015292915050565b60006040808385031215611e1657600080fd5b611e1e611c8f565b915082356001600160401b0380821115611e3757600080fd5b81850191506020808388031215611e4d57600080fd5b611e55611cb7565b833583811115611e6457600080fd5b80850194505087601f850112611e7957600080fd5b83359250611e8e611e8984611da1565b611d62565b83815260069390931b84018201928281019089851115611ead57600080fd5b948301945b84861015611ed357611ec48a87611dc4565b82529486019490830190611eb2565b8252508552948501359484019490945250909392505050565b600060408284031215611efe57600080fd5b611f06611c8f565b9050813581526020820135602082015292915050565b803563ffffffff811681146110cb57600080fd5b60006040808385031215611f4357600080fd5b611f4b611c8f565b915082356001600160401b03811115611f6357600080fd5b8301601f81018513611f7457600080fd5b80356020611f84611e8983611da1565b82815260a09283028401820192828201919089851115611fa357600080fd5b948301945b8486101561200c5780868b031215611fc05760008081fd5b611fc8611cd9565b611fd28b88611dc4565b815287870135858201526060611fe9818901611f1c565b89830152611ff960808901611f1c565b9082015283529485019491830191611fa8565b50808752505080860135818601525050505092915050565b60006101c0823603121561203757600080fd5b61203f611cfb565b61204883611d92565b815260208301356001600160401b038082111561206457600080fd5b61207036838701611e03565b60208401526120823660408701611eec565b6040840152608085013591508082111561209b57600080fd5b6120a736838701611e03565b606084015260a08501359150808211156120c057600080fd5b506120cd36828601611f30565b6080830152506120e03660c08501611eec565b60a08201526101008084013560c08301526101206120ff818601611f1c565b60e0840152610140612112818701611f1c565b838501526101609250612126838701611f1c565b91840191909152610180850135908301526101a090930135928101929092525090565b80356001600160401b03811681146110cb57600080fd5b600081830361010081121561217457600080fd5b61217c611d1e565b833581526060601f198301121561219257600080fd5b61219a611d40565b91506121a860208501612149565b82526121b660408501612149565b6020830152606084013560408301528160208201526080840135604082015260a0840135606082015260c084013560808201526121f560e08501611f1c565b60a0820152949350505050565b60006020828403121561221457600080fd5b813561ffff8116811461109657600080fd5b634e487b7160e01b600052601160045260246000fd5b600061ffff80831681851680830382111561225957612259612226565b01949350505050565b6020808252600790820152662727aa2fa4999960c91b604082015260600190565b6020808252600790820152661393d517d24d8d60ca1b604082015260600190565b600061ffff838116908316818110156122bf576122bf612226565b039392505050565b634e487b7160e01b600052601260045260246000fd5b6000828210156122ef576122ef612226565b500390565b600060ff821660ff84168082101561230e5761230e612226565b90039392505050565b600063ffffffff838116908316818110156122bf576122bf612226565b600063ffffffff80831681851680830382111561225957612259612226565b600063ffffffff8083168181141561236d5761236d612226565b6001019392505050565b600063ffffffff8084168061238e5761238e6122c7565b92169190910692915050565b60006001600160401b038084168061238e5761238e6122c7565b60006001600160401b03838116908316818110156122bf576122bf612226565b634e487b7160e01b600052603260045260246000fd5b60006000198214156123fe576123fe612226565b5060010190565b6000821982111561241857612418612226565b500190565b634e487b7160e01b600052605160045260246000fdfea2646970667358221220fc8fd4b42f5280ec9682c5446179e074592ebc574b83ff120770d3faceaa97a964736f6c63430008090033",
}

// OneStepProverMathABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverMathMetaData.ABI instead.
var OneStepProverMathABI = OneStepProverMathMetaData.ABI

// OneStepProverMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverMathMetaData.Bin instead.
var OneStepProverMathBin = OneStepProverMathMetaData.Bin

// DeployOneStepProverMath deploys a new Ethereum contract, binding an instance of OneStepProverMath to it.
func DeployOneStepProverMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProverMath, error) {
	parsed, err := OneStepProverMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverMath{OneStepProverMathCaller: OneStepProverMathCaller{contract: contract}, OneStepProverMathTransactor: OneStepProverMathTransactor{contract: contract}, OneStepProverMathFilterer: OneStepProverMathFilterer{contract: contract}}, nil
}

// OneStepProverMath is an auto generated Go binding around an Ethereum contract.
type OneStepProverMath struct {
	OneStepProverMathCaller     // Read-only binding to the contract
	OneStepProverMathTransactor // Write-only binding to the contract
	OneStepProverMathFilterer   // Log filterer for contract events
}

// OneStepProverMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverMathSession struct {
	Contract     *OneStepProverMath // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OneStepProverMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverMathCallerSession struct {
	Contract *OneStepProverMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OneStepProverMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverMathTransactorSession struct {
	Contract     *OneStepProverMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OneStepProverMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverMathRaw struct {
	Contract *OneStepProverMath // Generic contract binding to access the raw methods on
}

// OneStepProverMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverMathCallerRaw struct {
	Contract *OneStepProverMathCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverMathTransactorRaw struct {
	Contract *OneStepProverMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverMath creates a new instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMath(address common.Address, backend bind.ContractBackend) (*OneStepProverMath, error) {
	contract, err := bindOneStepProverMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMath{OneStepProverMathCaller: OneStepProverMathCaller{contract: contract}, OneStepProverMathTransactor: OneStepProverMathTransactor{contract: contract}, OneStepProverMathFilterer: OneStepProverMathFilterer{contract: contract}}, nil
}

// NewOneStepProverMathCaller creates a new read-only instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverMathCaller, error) {
	contract, err := bindOneStepProverMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathCaller{contract: contract}, nil
}

// NewOneStepProverMathTransactor creates a new write-only instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverMathTransactor, error) {
	contract, err := bindOneStepProverMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathTransactor{contract: contract}, nil
}

// NewOneStepProverMathFilterer creates a new log filterer instance of OneStepProverMath, bound to a specific deployed contract.
func NewOneStepProverMathFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverMathFilterer, error) {
	contract, err := bindOneStepProverMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMathFilterer{contract: contract}, nil
}

// bindOneStepProverMath binds a generic wrapper to an already deployed contract.
func bindOneStepProverMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMath *OneStepProverMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMath.Contract.OneStepProverMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMath *OneStepProverMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.OneStepProverMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMath *OneStepProverMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.OneStepProverMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMath *OneStepProverMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMath *OneStepProverMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMath *OneStepProverMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMath.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathCaller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverMath.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMath.Contract.ExecuteOneStep(&_OneStepProverMath.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMath *OneStepProverMathCallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMath.Contract.ExecuteOneStep(&_OneStepProverMath.CallOpts, arg0, startMach, startMod, inst, proof)
}

// OneStepProverMemoryMetaData contains all meta data concerning the OneStepProverMemory contract.
var OneStepProverMemoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxInboxMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge\",\"type\":\"address\"}],\"internalType\":\"structExecutionContext\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"startMach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"startMod\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"opcode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"argumentData\",\"type\":\"uint256\"}],\"internalType\":\"structInstruction\",\"name\":\"inst\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"executeOneStep\",\"outputs\":[{\"components\":[{\"internalType\":\"enumMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"valueStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"valueMultiStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue[]\",\"name\":\"inner\",\"type\":\"tuple[]\"}],\"internalType\":\"structValueArray\",\"name\":\"proved\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structValueStack\",\"name\":\"internalStack\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumValueType\",\"name\":\"valueType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"contents\",\"type\":\"uint256\"}],\"internalType\":\"structValue\",\"name\":\"returnPc\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"localsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"callerModule\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"callerModuleInternals\",\"type\":\"uint32\"}],\"internalType\":\"structStackFrame[]\",\"name\":\"proved\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structStackFrameWindow\",\"name\":\"frameStack\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"inactiveStackHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"remainingHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMultiStack\",\"name\":\"frameMultiStack\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"globalStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"moduleIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionIdx\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"functionPc\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recoveryPc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"modulesRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structMachine\",\"name\":\"mach\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"globalsMerkleRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxSize\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structModuleMemory\",\"name\":\"moduleMemory\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"tablesMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"functionsMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"extraHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"internalsOffset\",\"type\":\"uint32\"}],\"internalType\":\"structModule\",\"name\":\"mod\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611f84806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80633604366f14610030575b600080fd5b61004361003e366004611444565b61005a565b604051610051929190611658565b60405180910390f35b610062611341565b6040805160c081018252600080825282516060808201855282825260208083018490528286018490528401919091529282018190529181018290526080810182905260a08101919091526100b587611b80565b91506100c636879003870187611cbc565b905060006100d76020870187611d5e565b9050611422602861ffff8316108015906100f65750603561ffff831611155b1561010457506101bb61019d565b603661ffff83161080159061011e5750603e61ffff831611155b1561012c575061062661019d565b61ffff8216603f141561014257506109d061019d565b61ffff8216604014156101585750610a0861019d565b60405162461bcd60e51b8152602060048201526015602482015274494e56414c49445f4d454d4f52595f4f50434f444560581b60448201526064015b60405180910390fd5b6101ae84848989898663ffffffff16565b5050965096945050505050565b6000808060286101ce6020880188611d5e565b61ffff1614156101e7575060009150600490508161043c565b60296101f66020880188611d5e565b61ffff16141561021057506001915060089050600061043c565b602a61021f6020880188611d5e565b61ffff16141561023957506002915060049050600061043c565b602b6102486020880188611d5e565b61ffff16141561026257506003915060089050600061043c565b602c6102716020880188611d5e565b61ffff16141561028a575060009150600190508061043c565b602d6102996020880188611d5e565b61ffff1614156102b2575060009150600190508161043c565b602e6102c16020880188611d5e565b61ffff1614156102db57506000915060029050600161043c565b602f6102ea6020880188611d5e565b61ffff161415610303575060009150600290508161043c565b60306103126020880188611d5e565b61ffff16141561032a5750600191508190508061043c565b60316103396020880188611d5e565b61ffff161415610352575060019150819050600061043c565b60326103616020880188611d5e565b61ffff16141561037a575060019150600290508161043c565b60336103896020880188611d5e565b61ffff1614156103a357506001915060029050600061043c565b60346103b26020880188611d5e565b61ffff1614156103cb575060019150600490508161043c565b60356103da6020880188611d5e565b61ffff1614156103f457506001915060049050600061043c565b60405162461bcd60e51b815260206004820152601a60248201527f494e56414c49445f4d454d4f52595f4c4f41445f4f50434f44450000000000006044820152606401610194565b600061045361044e8a60200151610ab7565b610adc565b6104679063ffffffff166020890135611d98565b602089015190915060009081906104829084878b8b86610b6d565b5091509150811561049d575050600289525061061f92505050565b8084156105dc578560011480156104c5575060008760068111156104c3576104c3611529565b145b156104db578060000b63ffffffff1690506105dc565b8560011480156104fc575060018760068111156104fa576104fa611529565b145b156105095760000b6105dc565b85600214801561052a5750600087600681111561052857610528611529565b145b15610540578060010b63ffffffff1690506105dc565b8560021480156105615750600187600681111561055f5761055f611529565b145b1561056e5760010b6105dc565b85600414801561058f5750600187600681111561058d5761058d611529565b145b1561059c5760030b6105dc565b60405162461bcd60e51b815260206004820152601560248201527410905117d491505117d096551154d7d4d251d39151605a1b6044820152606401610194565b61061760405180604001604052808960068111156105fc576105fc611529565b81526001600160401b0384166020918201528e015190610c45565b505050505050505b5050505050565b6000808060366106396020880188611d5e565b61ffff16141561064f57506004915060006107be565b603761065e6020880188611d5e565b61ffff16141561067457506008915060016107be565b60386106836020880188611d5e565b61ffff16141561069957506004915060026107be565b60396106a86020880188611d5e565b61ffff1614156106be57506008915060036107be565b603a6106cd6020880188611d5e565b61ffff1614156106e357506001915060006107be565b603b6106f26020880188611d5e565b61ffff16141561070857506002915060006107be565b603c6107176020880188611d5e565b61ffff16141561072c575060019150816107be565b603d61073b6020880188611d5e565b61ffff16141561075157506002915060016107be565b603e6107606020880188611d5e565b61ffff16141561077657506004915060016107be565b60405162461bcd60e51b815260206004820152601b60248201527f494e56414c49445f4d454d4f52595f53544f52455f4f50434f444500000000006044820152606401610194565b60006107cd8960200151610ab7565b90508160068111156107e1576107e1611529565b815160068111156107f4576107f4611529565b146108325760405162461bcd60e51b815260206004820152600e60248201526d4241445f53544f52455f5459504560901b6044820152606401610194565b806020015192506008846001600160401b0316101561087d576001610858856008611db0565b6001600160401b031660016001600160401b0316901b6108789190611ddf565b831692505b5050600061089161044e8960200151610ab7565b6108a59063ffffffff166020880135611d98565b90508660200151600001516001600160401b0316836001600160401b0316826108ce9190611d98565b11156108e0575050600286525061061f565b604080516020810190915260608152600090600019906000805b876001600160401b03168110156109ad5760006109178288611d98565b90506000610926602083611e1d565b905085811461096b57600019861461094d57610943858786610c55565b60208f0151604001525b61095e8e60200151828e8e8b610cd6565b9098509196509094509250845b6000610978602084611e31565b905061098585828c610d70565b945060088a6001600160401b0316901c995050505080806109a590611e45565b9150506108fa565b506109b9828483610c55565b60208c015160400152505050505050505050505050565b6020840151516000906109e7906201000090611e60565b9050610a006109f582610df5565b602088015190610c45565b505050505050565b602084015151600090610a1f906201000090611e60565b90506000610a3361044e8860200151610ab7565b90506000610a4a63ffffffff808416908516611d98565b90508660200151602001516001600160401b03168111610a9f57610a716201000082611e86565b60208801516001600160401b039091169052610a9a610a8f84610df5565b60208a015190610c45565b610aad565b610aad610a8f600019610df5565b5050505050505050565b60408051808201909152600080825260208201528151610ad690610e28565b92915050565b60208101516000908183516006811115610af857610af8611529565b14610b2f5760405162461bcd60e51b81526020600482015260076024820152662727aa2fa4999960c91b6044820152606401610194565b6401000000008110610ad65760405162461bcd60e51b81526020600482015260076024820152662120a22fa4999960c91b6044820152606401610194565b8551600090819081906001600160401b0316610b89888a611d98565b1115610b9e5750600191506000905082610c39565b600019600080805b8a811015610c2c576000610bba828e611d98565b90506000610bc9602083611e1d565b9050858114610be957610bdf8f828e8e8e610cd6565b509a509095509350845b6000610bf6602084611e31565b9050610c03846008611e86565b610c0d8783610f38565b60ff16901b851794505050508080610c2490611e45565b915050610ba6565b5060009550935085925050505b96509650969350505050565b8151610c519082610fb2565b5050565b6040516b26b2b6b7b93c903632b0b31d60a11b6020820152602c81018290526000908190604c01604051602081830303815290604052805190602001209050610ccb8585836040518060400160405280601381526020017226b2b6b7b93c9036b2b935b632903a3932b29d60691b8152506110a5565b9150505b9392505050565b600080610cef6040518060200160405280606081525090565b839150610cfd8686846111b7565b9093509150610d0d8686846111d3565b925090506000610d1e828986610c55565b905088604001518114610d645760405162461bcd60e51b815260206004820152600e60248201526d15d493d391d7d3515357d493d3d560921b6044820152606401610194565b50955095509592505050565b600060208310610dba5760405162461bcd60e51b81526020600482015260156024820152740848288bea68aa8be988a828cbe84b2a88abe9288b605b1b6044820152606401610194565b600083610dc960016020611ea5565b610dd39190611ea5565b610dde906008611e86565b60ff848116821b911b198616179150509392505050565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b604080518082019091526000808252602082015281518051610e4c90600190611ea5565b81518110610e5c57610e5c611ebc565b6020026020010151905060006001836000015151610e7a9190611ea5565b6001600160401b03811115610e9157610e916117d0565b604051908082528060200260200182016040528015610ed657816020015b6040805180820190915260008082526020820152815260200190600190039081610eaf5790505b50905060005b8151811015610f31578351805182908110610ef957610ef9611ebc565b6020026020010151828281518110610f1357610f13611ebc565b60200260200101819052508080610f2990611e45565b915050610edc565b5090915290565b600060208210610f835760405162461bcd60e51b81526020600482015260166024820152750848288bea0aa9898be988a828cbe84b2a88abe9288b60531b6044820152606401610194565b600082610f9260016020611ea5565b610f9c9190611ea5565b610fa7906008611e86565b9390931c9392505050565b815151600090610fc3906001611d98565b6001600160401b03811115610fda57610fda6117d0565b60405190808252806020026020018201604052801561101f57816020015b6040805180820190915260008082526020820152815260200190600190039081610ff85790505b50905060005b83515181101561107b57835180518290811061104357611043611ebc565b602002602001015182828151811061105d5761105d611ebc565b6020026020010181905250808061107390611e45565b915050611025565b5081818460000151518151811061109457611094611ebc565b602090810291909101015290915250565b8160005b85515181101561116e576001851661110a578282876000015183815181106110d3576110d3611ebc565b60200260200101516040516020016110ed93929190611ed2565b604051602081830303815290604052805190602001209150611155565b828660000151828151811061112157611121611ebc565b60200260200101518360405160200161113c93929190611ed2565b6040516020818303038152906040528051906020012091505b60019490941c938061116681611e45565b9150506110a9565b5083156111af5760405162461bcd60e51b815260206004820152600f60248201526e141493d3d197d513d3d7d4d213d495608a1b6044820152606401610194565b949350505050565b600081816111c68686846112ad565b9097909650945050505050565b6040805160208101909152606081528160006111f086868461130b565b92509050600060ff82166001600160401b03811115611211576112116117d0565b60405190808252806020026020018201604052801561123a578160200160208202803683370190505b50905060005b8260ff168160ff161015611291576112598888866111b7565b838360ff168151811061126e5761126e611ebc565b60200260200101819650828152505050808061128990611f18565b915050611240565b5060405180602001604052808281525093505050935093915050565b600081815b602081101561130257600883901b92508585838181106112d4576112d4611ebc565b919091013560f81c939093179250816112ec81611e45565b92505080806112fa90611e45565b9150506112b2565b50935093915050565b60008184848281811061132057611320611ebc565b919091013560f81c925081905061133681611e45565b915050935093915050565b604080516101808101909152806000815260200161137660408051606080820183529181019182529081526000602082015290565b81526040805180820182526000808252602080830191909152830152016113b460408051606080820183529181019182529081526000602082015290565b81526020016113d9604051806040016040528060608152602001600080191681525090565b815260408051808201825260008082526020808301829052840191909152908201819052606082018190526080820181905260a0820181905260c0820181905260e09091015290565b61142a611f38565b565b60006040828403121561143e57600080fd5b50919050565b6000806000806000808688036101c08082121561146057600080fd5b61146a8a8a61142c565b975060408901356001600160401b038082111561148657600080fd5b818b01915082828d03121561149a57600080fd5b819850610100605f19850112156114b057600080fd5b60608b0197506114c48c6101608d0161142c565b96506101a08b01359350808411156114db57600080fd5b838b0193508b601f8501126114ef57600080fd5b833592508083111561150057600080fd5b505089602082840101111561151457600080fd5b60208201935080925050509295509295509295565b634e487b7160e01b600052602160045260246000fd5b6004811061154f5761154f611529565b9052565b80516007811061156557611565611529565b8252602090810151910152565b805160408084529051602084830181905281516060860181905260009392820191849160808801905b808410156115c2576115ae828651611553565b93820193600193909301929085019061159b565b509581015196019590955250919392505050565b8051604080845281518482018190526000926060916020918201918388019190865b8281101561164157845161160d858251611553565b80830151858901528781015163ffffffff90811688870152908701511660808501529381019360a0909301926001016115f8565b509687015197909601969096525093949350505050565b600061012080835261166d818401865161153f565b60208501516101c0610140818187015261168b6102e0870184611572565b925060408801516101606116ab8189018380518252602090810151910152565b60608a0151915061011f1980898703016101a08a01526116cb8684611572565b955060808b015192508089870301858a0152506116e885836115d6565b60a08b015180516101e08b015260208101516102008b0152909550935060c08a015161022089015260e08a015163ffffffff81166102408a015293506101008a015163ffffffff81166102608a015293509489015163ffffffff811661028089015294918901516102a0880152508701516102c0860152509150610ccf905060208301848051825260208101516001600160401b0380825116602085015280602083015116604085015250604081015160608401525060408101516080830152606081015160a0830152608081015160c083015263ffffffff60a08201511660e08301525050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715611808576118086117d0565b60405290565b604051602081016001600160401b0381118282101715611808576118086117d0565b604051608081016001600160401b0381118282101715611808576118086117d0565b60405161018081016001600160401b0381118282101715611808576118086117d0565b60405160c081016001600160401b0381118282101715611808576118086117d0565b604051606081016001600160401b0381118282101715611808576118086117d0565b604051601f8201601f191681016001600160401b03811182821017156118e1576118e16117d0565b604052919050565b8035600481106118f857600080fd5b919050565b60006001600160401b03821115611916576119166117d0565b5060051b60200190565b60006040828403121561193257600080fd5b61193a6117e6565b905081356007811061194b57600080fd5b808252506020820135602082015292915050565b6000604080838503121561197257600080fd5b61197a6117e6565b915082356001600160401b038082111561199357600080fd5b818501915060208083880312156119a957600080fd5b6119b161180e565b8335838111156119c057600080fd5b80850194505087601f8501126119d557600080fd5b833592506119ea6119e5846118fd565b6118b9565b83815260069390931b84018201928281019089851115611a0957600080fd5b948301945b84861015611a2f57611a208a87611920565b82529486019490830190611a0e565b8252508552948501359484019490945250909392505050565b600060408284031215611a5a57600080fd5b611a626117e6565b9050813581526020820135602082015292915050565b803563ffffffff811681146118f857600080fd5b60006040808385031215611a9f57600080fd5b611aa76117e6565b915082356001600160401b03811115611abf57600080fd5b8301601f81018513611ad057600080fd5b80356020611ae06119e5836118fd565b82815260a09283028401820192828201919089851115611aff57600080fd5b948301945b84861015611b685780868b031215611b1c5760008081fd5b611b24611830565b611b2e8b88611920565b815287870135858201526060611b45818901611a78565b89830152611b5560808901611a78565b9082015283529485019491830191611b04565b50808752505080860135818601525050505092915050565b60006101c08236031215611b9357600080fd5b611b9b611852565b611ba4836118e9565b815260208301356001600160401b0380821115611bc057600080fd5b611bcc3683870161195f565b6020840152611bde3660408701611a48565b60408401526080850135915080821115611bf757600080fd5b611c033683870161195f565b606084015260a0850135915080821115611c1c57600080fd5b50611c2936828601611a8c565b608083015250611c3c3660c08501611a48565b60a08201526101008084013560c0830152610120611c5b818601611a78565b60e0840152610140611c6e818701611a78565b838501526101609250611c82838701611a78565b91840191909152610180850135908301526101a090930135928101929092525090565b80356001600160401b03811681146118f857600080fd5b6000818303610100811215611cd057600080fd5b611cd8611875565b833581526060601f1983011215611cee57600080fd5b611cf6611897565b9150611d0460208501611ca5565b8252611d1260408501611ca5565b6020830152606084013560408301528160208201526080840135604082015260a0840135606082015260c08401356080820152611d5160e08501611a78565b60a0820152949350505050565b600060208284031215611d7057600080fd5b813561ffff81168114610ccf57600080fd5b634e487b7160e01b600052601160045260246000fd5b60008219821115611dab57611dab611d82565b500190565b60006001600160401b0380831681851681830481118215151615611dd657611dd6611d82565b02949350505050565b60006001600160401b0383811690831681811015611dff57611dff611d82565b039392505050565b634e487b7160e01b600052601260045260246000fd5b600082611e2c57611e2c611e07565b500490565b600082611e4057611e40611e07565b500690565b6000600019821415611e5957611e59611d82565b5060010190565b60006001600160401b0380841680611e7a57611e7a611e07565b92169190910492915050565b6000816000190483118215151615611ea057611ea0611d82565b500290565b600082821015611eb757611eb7611d82565b500390565b634e487b7160e01b600052603260045260246000fd5b6000845160005b81811015611ef35760208188018101518583015201611ed9565b81811115611f02576000828501525b5091909101928352506020820152604001919050565b600060ff821660ff811415611f2f57611f2f611d82565b60010192915050565b634e487b7160e01b600052605160045260246000fdfea264697066735822122075656d2789d3f1dcea2eac8ac49fb7a432ca87fcd7794f8df78e6331d332953c64736f6c63430008090033",
}

// OneStepProverMemoryABI is the input ABI used to generate the binding from.
// Deprecated: Use OneStepProverMemoryMetaData.ABI instead.
var OneStepProverMemoryABI = OneStepProverMemoryMetaData.ABI

// OneStepProverMemoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneStepProverMemoryMetaData.Bin instead.
var OneStepProverMemoryBin = OneStepProverMemoryMetaData.Bin

// DeployOneStepProverMemory deploys a new Ethereum contract, binding an instance of OneStepProverMemory to it.
func DeployOneStepProverMemory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProverMemory, error) {
	parsed, err := OneStepProverMemoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneStepProverMemoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProverMemory{OneStepProverMemoryCaller: OneStepProverMemoryCaller{contract: contract}, OneStepProverMemoryTransactor: OneStepProverMemoryTransactor{contract: contract}, OneStepProverMemoryFilterer: OneStepProverMemoryFilterer{contract: contract}}, nil
}

// OneStepProverMemory is an auto generated Go binding around an Ethereum contract.
type OneStepProverMemory struct {
	OneStepProverMemoryCaller     // Read-only binding to the contract
	OneStepProverMemoryTransactor // Write-only binding to the contract
	OneStepProverMemoryFilterer   // Log filterer for contract events
}

// OneStepProverMemoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProverMemoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProverMemoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProverMemoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProverMemorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProverMemorySession struct {
	Contract     *OneStepProverMemory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OneStepProverMemoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProverMemoryCallerSession struct {
	Contract *OneStepProverMemoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// OneStepProverMemoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProverMemoryTransactorSession struct {
	Contract     *OneStepProverMemoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// OneStepProverMemoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProverMemoryRaw struct {
	Contract *OneStepProverMemory // Generic contract binding to access the raw methods on
}

// OneStepProverMemoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProverMemoryCallerRaw struct {
	Contract *OneStepProverMemoryCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProverMemoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProverMemoryTransactorRaw struct {
	Contract *OneStepProverMemoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProverMemory creates a new instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemory(address common.Address, backend bind.ContractBackend) (*OneStepProverMemory, error) {
	contract, err := bindOneStepProverMemory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemory{OneStepProverMemoryCaller: OneStepProverMemoryCaller{contract: contract}, OneStepProverMemoryTransactor: OneStepProverMemoryTransactor{contract: contract}, OneStepProverMemoryFilterer: OneStepProverMemoryFilterer{contract: contract}}, nil
}

// NewOneStepProverMemoryCaller creates a new read-only instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryCaller(address common.Address, caller bind.ContractCaller) (*OneStepProverMemoryCaller, error) {
	contract, err := bindOneStepProverMemory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryCaller{contract: contract}, nil
}

// NewOneStepProverMemoryTransactor creates a new write-only instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProverMemoryTransactor, error) {
	contract, err := bindOneStepProverMemory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryTransactor{contract: contract}, nil
}

// NewOneStepProverMemoryFilterer creates a new log filterer instance of OneStepProverMemory, bound to a specific deployed contract.
func NewOneStepProverMemoryFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProverMemoryFilterer, error) {
	contract, err := bindOneStepProverMemory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProverMemoryFilterer{contract: contract}, nil
}

// bindOneStepProverMemory binds a generic wrapper to an already deployed contract.
func bindOneStepProverMemory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneStepProverMemoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMemory.Contract.OneStepProverMemoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.OneStepProverMemoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMemory *OneStepProverMemoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.OneStepProverMemoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProverMemory *OneStepProverMemoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneStepProverMemory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProverMemory *OneStepProverMemoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProverMemory *OneStepProverMemoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProverMemory.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemoryCaller) ExecuteOneStep(opts *bind.CallOpts, arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	var out []interface{}
	err := _OneStepProverMemory.contract.Call(opts, &out, "executeOneStep", arg0, startMach, startMod, inst, proof)

	outstruct := new(struct {
		Mach Machine
		Mod  Module
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Mach = *abi.ConvertType(out[0], new(Machine)).(*Machine)
	outstruct.Mod = *abi.ConvertType(out[1], new(Module)).(*Module)

	return *outstruct, err

}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemorySession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMemory.Contract.ExecuteOneStep(&_OneStepProverMemory.CallOpts, arg0, startMach, startMod, inst, proof)
}

// ExecuteOneStep is a free data retrieval call binding the contract method 0x3604366f.
//
// Solidity: function executeOneStep((uint256,address) , (uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) startMach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) startMod, (uint16,uint256) inst, bytes proof) pure returns((uint8,(((uint8,uint256)[]),bytes32),(bytes32,bytes32),(((uint8,uint256)[]),bytes32),(((uint8,uint256),bytes32,uint32,uint32)[],bytes32),(bytes32,bytes32),bytes32,uint32,uint32,uint32,bytes32,bytes32) mach, (bytes32,(uint64,uint64,bytes32),bytes32,bytes32,bytes32,uint32) mod)
func (_OneStepProverMemory *OneStepProverMemoryCallerSession) ExecuteOneStep(arg0 ExecutionContext, startMach Machine, startMod Module, inst Instruction, proof []byte) (struct {
	Mach Machine
	Mod  Module
}, error) {
	return _OneStepProverMemory.Contract.ExecuteOneStep(&_OneStepProverMemory.CallOpts, arg0, startMach, startMod, inst, proof)
}
