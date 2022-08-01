// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocksgen

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
)

// ChallengeLibChallenge is an auto generated low-level Go binding around an user-defined struct.
type ChallengeLibChallenge struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}

// ChallengeLibParticipant is an auto generated low-level Go binding around an user-defined struct.
type ChallengeLibParticipant struct {
	Addr     common.Address
	TimeLeft *big.Int
}

// ChallengeLibSegmentSelection is an auto generated low-level Go binding around an user-defined struct.
type ChallengeLibSegmentSelection struct {
	OldSegmentsStart  *big.Int
	OldSegmentsLength *big.Int
	OldSegments       [][32]byte
	ChallengePosition *big.Int
}

// GlobalState is an auto generated low-level Go binding around an user-defined struct.
type GlobalState struct {
	Bytes32Vals [2][32]byte
	U64Vals     [2]uint64
}

// ISequencerInboxMaxTimeVariation is an auto generated low-level Go binding around an user-defined struct.
type ISequencerInboxMaxTimeVariation struct {
	DelayBlocks   *big.Int
	FutureBlocks  *big.Int
	DelaySeconds  *big.Int
	FutureSeconds *big.Int
}

// ISequencerInboxTimeBounds is an auto generated low-level Go binding around an user-defined struct.
type ISequencerInboxTimeBounds struct {
	MinTimestamp   uint64
	MaxTimestamp   uint64
	MinBlockNumber uint64
	MaxBlockNumber uint64
}

// BridgeStubMetaData contains all meta data concerning the BridgeStub contract.
var BridgeStubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BridgeCallTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"InboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"OutboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSequencerInbox\",\"type\":\"address\"}],\"name\":\"SequencerInboxUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptFundsFromOldBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeOutbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedDelayedInboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"allowedDelayedInboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedOutboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedOutboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delayedInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"enqueueDelayedMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"}],\"name\":\"enqueueSequencerMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seqMessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"acc\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setDelayedInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencerInbox\",\"type\":\"address\"}],\"name\":\"setSequencerInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"batchPoster\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"submitBatchSpendingReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610bd5806100206000396000f3fe6080604052600436106101035760003560e01c80639e5d4c49116100a0578063d5719dc211610064578063d5719dc21461032b578063e76f5c8d1461034b578063e77145f41461019c578063eca067ad1461036b578063ee35f3271461038057600080fd5b80639e5d4c491461026c578063ab5d89431461029a578063ae60bd13146102ba578063cb23bcb5146102f6578063cee3d7281461030b57600080fd5b806284120c1461010857806316bf55791461012c578063413b35bd1461014c57806347fb24c51461017c5780634f61f8501461019e5780637a88b107146101be578063845bde82146101e15780638db5993b14610221578063945e114714610234575b600080fd5b34801561011457600080fd5b506005545b6040519081526020015b60405180910390f35b34801561013857600080fd5b5061011961014736600461091c565b6103a0565b34801561015857600080fd5b5061016c610167366004610951565b6103c1565b6040519015158152602001610123565b34801561018857600080fd5b5061019c610197366004610973565b6103e4565b005b3480156101aa57600080fd5b5061019c6101b9366004610951565b610610565b3480156101ca57600080fd5b506101196101d93660046109af565b600092915050565b3480156101ed57600080fd5b506102016101fc3660046109d9565b610664565b604080519485526020850193909352918301526060820152608001610123565b61011961022f3660046109fb565b610743565b34801561024057600080fd5b5061025461024f36600461091c565b6107ac565b6040516001600160a01b039091168152602001610123565b34801561027857600080fd5b5061028c610287366004610a3f565b6107d6565b604051610123929190610ac6565b3480156102a657600080fd5b50600354610254906001600160a01b031681565b3480156102c657600080fd5b5061016c6102d5366004610951565b6001600160a01b031660009081526020819052604090206001015460ff1690565b34801561030257600080fd5b506102546103c1565b34801561031757600080fd5b5061019c610326366004610973565b6107f2565b34801561033757600080fd5b5061011961034636600461091c565b61080a565b34801561035757600080fd5b5061025461036636600461091c565b61081a565b34801561037757600080fd5b50600454610119565b34801561038c57600080fd5b50600654610254906001600160a01b031681565b600581815481106103b057600080fd5b600091825260209091200154905081565b600060405162461bcd60e51b81526004016103db90610b25565b60405180910390fd5b6001600160a01b03821660008181526020818152604091829020600181015492518515158152909360ff90931692917f6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521910160405180910390a28080156104485750825b8061045a57508015801561045a575082155b156104655750505050565b82156104f15760408051808201825260018054825260208083018281526001600160a01b0389166000818152928390529482209351845551928201805460ff1916931515939093179092558054808201825591527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b031916909117905561060a565b60018054610500908290610b4e565b8154811061051057610510610b73565b6000918252602090912001548254600180546001600160a01b0390931692909190811061053f5761053f610b73565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160000154600080600185600001548154811061058c5761058c610b73565b60009182526020808320909101546001600160a01b0316835282019290925260400190205560018054806105c2576105c2610b89565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038616825281905260408120908155600101805460ff191690555b50505050565b600680546001600160a01b0319166001600160a01b0383169081179091556040519081527f8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a9060200160405180910390a150565b60055460008080831561069f576005805461068190600190610b4e565b8154811061069157610691610b73565b906000526020600020015492505b84156106d05760046106b2600187610b4e565b815481106106c2576106c2610b73565b906000526020600020015491505b60408051602081018590529081018790526060810183905260800160408051601f198184030181529190528051602090910120600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001819055939692955090935050565b3360009081526020819052604081206001015460ff166107965760405162461bcd60e51b815260206004820152600e60248201526d09c9ea8be8ca49e9abe929c849eb60931b60448201526064016103db565b6107a484844342488761082a565b949350505050565b600281815481106107bc57600080fd5b6000918252602090912001546001600160a01b0316905081565b6000606060405162461bcd60e51b81526004016103db90610b25565b60405162461bcd60e51b81526004016103db90610b25565b600481815481106103b057600080fd5b600181815481106107bc57600080fd5b60045460408051600060208083018290526021830182905260358301829052603d8301829052604583018290526065830182905260858084018790528451808503909101815260a590930190935281519190920120909190600082156108b5576004610897600185610b4e565b815481106108a7576108a7610b73565b906000526020600020015490505b6040805160208082019390935280820193909352805180840382018152606090930190528151910120600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0155979650505050505050565b60006020828403121561092e57600080fd5b5035919050565b80356001600160a01b038116811461094c57600080fd5b919050565b60006020828403121561096357600080fd5b61096c82610935565b9392505050565b6000806040838503121561098657600080fd5b61098f83610935565b9150602083013580151581146109a457600080fd5b809150509250929050565b600080604083850312156109c257600080fd5b6109cb83610935565b946020939093013593505050565b600080604083850312156109ec57600080fd5b50508035926020909101359150565b600080600060608486031215610a1057600080fd5b833560ff81168114610a2157600080fd5b9250610a2f60208501610935565b9150604084013590509250925092565b60008060008060608587031215610a5557600080fd5b610a5e85610935565b935060208501359250604085013567ffffffffffffffff80821115610a8257600080fd5b818701915087601f830112610a9657600080fd5b813581811115610aa557600080fd5b886020828501011115610ab757600080fd5b95989497505060200194505050565b821515815260006020604081840152835180604085015260005b81811015610afc57858101830151858201606001528201610ae0565b81811115610b0e576000606083870101525b50601f01601f191692909201606001949350505050565b6020808252600f908201526e1393d517d253541311535153951151608a1b604082015260600190565b600082821015610b6e57634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052603160045260246000fdfea2646970667358221220e5518a44ad856cd2e19e1a2d65d0ff84d55d892fbeb542dd62106ffc8d62e03364736f6c63430008090033",
}

// BridgeStubABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeStubMetaData.ABI instead.
var BridgeStubABI = BridgeStubMetaData.ABI

// BridgeStubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeStubMetaData.Bin instead.
var BridgeStubBin = BridgeStubMetaData.Bin

// DeployBridgeStub deploys a new Ethereum contract, binding an instance of BridgeStub to it.
func DeployBridgeStub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeStub, error) {
	parsed, err := BridgeStubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeStubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeStub{BridgeStubCaller: BridgeStubCaller{contract: contract}, BridgeStubTransactor: BridgeStubTransactor{contract: contract}, BridgeStubFilterer: BridgeStubFilterer{contract: contract}}, nil
}

// BridgeStub is an auto generated Go binding around an Ethereum contract.
type BridgeStub struct {
	BridgeStubCaller     // Read-only binding to the contract
	BridgeStubTransactor // Write-only binding to the contract
	BridgeStubFilterer   // Log filterer for contract events
}

// BridgeStubCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeStubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeStubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeStubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeStubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeStubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeStubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeStubSession struct {
	Contract     *BridgeStub       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeStubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeStubCallerSession struct {
	Contract *BridgeStubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BridgeStubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeStubTransactorSession struct {
	Contract     *BridgeStubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BridgeStubRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeStubRaw struct {
	Contract *BridgeStub // Generic contract binding to access the raw methods on
}

// BridgeStubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeStubCallerRaw struct {
	Contract *BridgeStubCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeStubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeStubTransactorRaw struct {
	Contract *BridgeStubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeStub creates a new instance of BridgeStub, bound to a specific deployed contract.
func NewBridgeStub(address common.Address, backend bind.ContractBackend) (*BridgeStub, error) {
	contract, err := bindBridgeStub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeStub{BridgeStubCaller: BridgeStubCaller{contract: contract}, BridgeStubTransactor: BridgeStubTransactor{contract: contract}, BridgeStubFilterer: BridgeStubFilterer{contract: contract}}, nil
}

// NewBridgeStubCaller creates a new read-only instance of BridgeStub, bound to a specific deployed contract.
func NewBridgeStubCaller(address common.Address, caller bind.ContractCaller) (*BridgeStubCaller, error) {
	contract, err := bindBridgeStub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeStubCaller{contract: contract}, nil
}

// NewBridgeStubTransactor creates a new write-only instance of BridgeStub, bound to a specific deployed contract.
func NewBridgeStubTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeStubTransactor, error) {
	contract, err := bindBridgeStub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeStubTransactor{contract: contract}, nil
}

// NewBridgeStubFilterer creates a new log filterer instance of BridgeStub, bound to a specific deployed contract.
func NewBridgeStubFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeStubFilterer, error) {
	contract, err := bindBridgeStub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeStubFilterer{contract: contract}, nil
}

// bindBridgeStub binds a generic wrapper to an already deployed contract.
func bindBridgeStub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeStubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeStub *BridgeStubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeStub.Contract.BridgeStubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeStub *BridgeStubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeStub.Contract.BridgeStubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeStub *BridgeStubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeStub.Contract.BridgeStubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeStub *BridgeStubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeStub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeStub *BridgeStubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeStub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeStub *BridgeStubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeStub.Contract.contract.Transact(opts, method, params...)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeStub *BridgeStubCaller) ActiveOutbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "activeOutbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeStub *BridgeStubSession) ActiveOutbox() (common.Address, error) {
	return _BridgeStub.Contract.ActiveOutbox(&_BridgeStub.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeStub *BridgeStubCallerSession) ActiveOutbox() (common.Address, error) {
	return _BridgeStub.Contract.ActiveOutbox(&_BridgeStub.CallOpts)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubCaller) AllowedDelayedInboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "allowedDelayedInboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeStub.Contract.AllowedDelayedInboxList(&_BridgeStub.CallOpts, arg0)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubCallerSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeStub.Contract.AllowedDelayedInboxList(&_BridgeStub.CallOpts, arg0)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeStub *BridgeStubCaller) AllowedDelayedInboxes(opts *bind.CallOpts, inbox common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "allowedDelayedInboxes", inbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeStub *BridgeStubSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeStub.Contract.AllowedDelayedInboxes(&_BridgeStub.CallOpts, inbox)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeStub *BridgeStubCallerSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeStub.Contract.AllowedDelayedInboxes(&_BridgeStub.CallOpts, inbox)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubCaller) AllowedOutboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "allowedOutboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeStub.Contract.AllowedOutboxList(&_BridgeStub.CallOpts, arg0)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubCallerSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeStub.Contract.AllowedOutboxList(&_BridgeStub.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address ) pure returns(bool)
func (_BridgeStub *BridgeStubCaller) AllowedOutboxes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "allowedOutboxes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address ) pure returns(bool)
func (_BridgeStub *BridgeStubSession) AllowedOutboxes(arg0 common.Address) (bool, error) {
	return _BridgeStub.Contract.AllowedOutboxes(&_BridgeStub.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address ) pure returns(bool)
func (_BridgeStub *BridgeStubCallerSession) AllowedOutboxes(arg0 common.Address) (bool, error) {
	return _BridgeStub.Contract.AllowedOutboxes(&_BridgeStub.CallOpts, arg0)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubCaller) DelayedInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "delayedInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeStub.Contract.DelayedInboxAccs(&_BridgeStub.CallOpts, arg0)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubCallerSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeStub.Contract.DelayedInboxAccs(&_BridgeStub.CallOpts, arg0)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCaller) DelayedMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "delayedMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.DelayedMessageCount(&_BridgeStub.CallOpts)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCallerSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.DelayedMessageCount(&_BridgeStub.CallOpts)
}

// ExecuteCall is a free data retrieval call binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address , uint256 , bytes ) pure returns(bool, bytes)
func (_BridgeStub *BridgeStubCaller) ExecuteCall(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (bool, []byte, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "executeCall", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// ExecuteCall is a free data retrieval call binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address , uint256 , bytes ) pure returns(bool, bytes)
func (_BridgeStub *BridgeStubSession) ExecuteCall(arg0 common.Address, arg1 *big.Int, arg2 []byte) (bool, []byte, error) {
	return _BridgeStub.Contract.ExecuteCall(&_BridgeStub.CallOpts, arg0, arg1, arg2)
}

// ExecuteCall is a free data retrieval call binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address , uint256 , bytes ) pure returns(bool, bytes)
func (_BridgeStub *BridgeStubCallerSession) ExecuteCall(arg0 common.Address, arg1 *big.Int, arg2 []byte) (bool, []byte, error) {
	return _BridgeStub.Contract.ExecuteCall(&_BridgeStub.CallOpts, arg0, arg1, arg2)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() pure returns(address)
func (_BridgeStub *BridgeStubCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() pure returns(address)
func (_BridgeStub *BridgeStubSession) Rollup() (common.Address, error) {
	return _BridgeStub.Contract.Rollup(&_BridgeStub.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() pure returns(address)
func (_BridgeStub *BridgeStubCallerSession) Rollup() (common.Address, error) {
	return _BridgeStub.Contract.Rollup(&_BridgeStub.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeStub *BridgeStubCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeStub *BridgeStubSession) SequencerInbox() (common.Address, error) {
	return _BridgeStub.Contract.SequencerInbox(&_BridgeStub.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeStub *BridgeStubCallerSession) SequencerInbox() (common.Address, error) {
	return _BridgeStub.Contract.SequencerInbox(&_BridgeStub.CallOpts)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubCaller) SequencerInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "sequencerInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeStub.Contract.SequencerInboxAccs(&_BridgeStub.CallOpts, arg0)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubCallerSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeStub.Contract.SequencerInboxAccs(&_BridgeStub.CallOpts, arg0)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCaller) SequencerMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "sequencerMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.SequencerMessageCount(&_BridgeStub.CallOpts)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCallerSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.SequencerMessageCount(&_BridgeStub.CallOpts)
}

// SetOutbox is a free data retrieval call binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address , bool ) pure returns()
func (_BridgeStub *BridgeStubCaller) SetOutbox(opts *bind.CallOpts, arg0 common.Address, arg1 bool) error {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "setOutbox", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// SetOutbox is a free data retrieval call binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address , bool ) pure returns()
func (_BridgeStub *BridgeStubSession) SetOutbox(arg0 common.Address, arg1 bool) error {
	return _BridgeStub.Contract.SetOutbox(&_BridgeStub.CallOpts, arg0, arg1)
}

// SetOutbox is a free data retrieval call binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address , bool ) pure returns()
func (_BridgeStub *BridgeStubCallerSession) SetOutbox(arg0 common.Address, arg1 bool) error {
	return _BridgeStub.Contract.SetOutbox(&_BridgeStub.CallOpts, arg0, arg1)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeStub *BridgeStubTransactor) AcceptFundsFromOldBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "acceptFundsFromOldBridge")
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeStub *BridgeStubSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeStub.Contract.AcceptFundsFromOldBridge(&_BridgeStub.TransactOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeStub *BridgeStubTransactorSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeStub.Contract.AcceptFundsFromOldBridge(&_BridgeStub.TransactOpts)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeStub *BridgeStubTransactor) EnqueueDelayedMessage(opts *bind.TransactOpts, kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "enqueueDelayedMessage", kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeStub *BridgeStubSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.Contract.EnqueueDelayedMessage(&_BridgeStub.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeStub *BridgeStubTransactorSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.Contract.EnqueueDelayedMessage(&_BridgeStub.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x845bde82.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeStub *BridgeStubTransactor) EnqueueSequencerMessage(opts *bind.TransactOpts, dataHash [32]byte, afterDelayedMessagesRead *big.Int) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "enqueueSequencerMessage", dataHash, afterDelayedMessagesRead)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x845bde82.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeStub *BridgeStubSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int) (*types.Transaction, error) {
	return _BridgeStub.Contract.EnqueueSequencerMessage(&_BridgeStub.TransactOpts, dataHash, afterDelayedMessagesRead)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x845bde82.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeStub *BridgeStubTransactorSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int) (*types.Transaction, error) {
	return _BridgeStub.Contract.EnqueueSequencerMessage(&_BridgeStub.TransactOpts, dataHash, afterDelayedMessagesRead)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeStub *BridgeStubTransactor) SetDelayedInbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "setDelayedInbox", inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeStub *BridgeStubSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeStub.Contract.SetDelayedInbox(&_BridgeStub.TransactOpts, inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeStub *BridgeStubTransactorSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeStub.Contract.SetDelayedInbox(&_BridgeStub.TransactOpts, inbox, enabled)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeStub *BridgeStubTransactor) SetSequencerInbox(opts *bind.TransactOpts, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "setSequencerInbox", _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeStub *BridgeStubSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeStub.Contract.SetSequencerInbox(&_BridgeStub.TransactOpts, _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeStub *BridgeStubTransactorSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeStub.Contract.SetSequencerInbox(&_BridgeStub.TransactOpts, _sequencerInbox)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeStub *BridgeStubTransactor) SubmitBatchSpendingReport(opts *bind.TransactOpts, batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "submitBatchSpendingReport", batchPoster, dataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeStub *BridgeStubSession) SubmitBatchSpendingReport(batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.Contract.SubmitBatchSpendingReport(&_BridgeStub.TransactOpts, batchPoster, dataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeStub *BridgeStubTransactorSession) SubmitBatchSpendingReport(batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.Contract.SubmitBatchSpendingReport(&_BridgeStub.TransactOpts, batchPoster, dataHash)
}

// BridgeStubBridgeCallTriggeredIterator is returned from FilterBridgeCallTriggered and is used to iterate over the raw logs and unpacked data for BridgeCallTriggered events raised by the BridgeStub contract.
type BridgeStubBridgeCallTriggeredIterator struct {
	Event *BridgeStubBridgeCallTriggered // Event containing the contract specifics and raw log

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
func (it *BridgeStubBridgeCallTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubBridgeCallTriggered)
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
		it.Event = new(BridgeStubBridgeCallTriggered)
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
func (it *BridgeStubBridgeCallTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubBridgeCallTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubBridgeCallTriggered represents a BridgeCallTriggered event raised by the BridgeStub contract.
type BridgeStubBridgeCallTriggered struct {
	Outbox common.Address
	To     common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeCallTriggered is a free log retrieval operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeStub *BridgeStubFilterer) FilterBridgeCallTriggered(opts *bind.FilterOpts, outbox []common.Address, to []common.Address) (*BridgeStubBridgeCallTriggeredIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeStubBridgeCallTriggeredIterator{contract: _BridgeStub.contract, event: "BridgeCallTriggered", logs: logs, sub: sub}, nil
}

// WatchBridgeCallTriggered is a free log subscription operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeStub *BridgeStubFilterer) WatchBridgeCallTriggered(opts *bind.WatchOpts, sink chan<- *BridgeStubBridgeCallTriggered, outbox []common.Address, to []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubBridgeCallTriggered)
				if err := _BridgeStub.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
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

// ParseBridgeCallTriggered is a log parse operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeStub *BridgeStubFilterer) ParseBridgeCallTriggered(log types.Log) (*BridgeStubBridgeCallTriggered, error) {
	event := new(BridgeStubBridgeCallTriggered)
	if err := _BridgeStub.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeStubInboxToggleIterator is returned from FilterInboxToggle and is used to iterate over the raw logs and unpacked data for InboxToggle events raised by the BridgeStub contract.
type BridgeStubInboxToggleIterator struct {
	Event *BridgeStubInboxToggle // Event containing the contract specifics and raw log

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
func (it *BridgeStubInboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubInboxToggle)
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
		it.Event = new(BridgeStubInboxToggle)
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
func (it *BridgeStubInboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubInboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubInboxToggle represents a InboxToggle event raised by the BridgeStub contract.
type BridgeStubInboxToggle struct {
	Inbox   common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInboxToggle is a free log retrieval operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) FilterInboxToggle(opts *bind.FilterOpts, inbox []common.Address) (*BridgeStubInboxToggleIterator, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeStubInboxToggleIterator{contract: _BridgeStub.contract, event: "InboxToggle", logs: logs, sub: sub}, nil
}

// WatchInboxToggle is a free log subscription operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) WatchInboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeStubInboxToggle, inbox []common.Address) (event.Subscription, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubInboxToggle)
				if err := _BridgeStub.contract.UnpackLog(event, "InboxToggle", log); err != nil {
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

// ParseInboxToggle is a log parse operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) ParseInboxToggle(log types.Log) (*BridgeStubInboxToggle, error) {
	event := new(BridgeStubInboxToggle)
	if err := _BridgeStub.contract.UnpackLog(event, "InboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeStubMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the BridgeStub contract.
type BridgeStubMessageDeliveredIterator struct {
	Event *BridgeStubMessageDelivered // Event containing the contract specifics and raw log

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
func (it *BridgeStubMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubMessageDelivered)
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
		it.Event = new(BridgeStubMessageDelivered)
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
func (it *BridgeStubMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubMessageDelivered represents a MessageDelivered event raised by the BridgeStub contract.
type BridgeStubMessageDelivered struct {
	MessageIndex    *big.Int
	BeforeInboxAcc  [32]byte
	Inbox           common.Address
	Kind            uint8
	Sender          common.Address
	MessageDataHash [32]byte
	BaseFeeL1       *big.Int
	Timestamp       uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeStub *BridgeStubFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (*BridgeStubMessageDeliveredIterator, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &BridgeStubMessageDeliveredIterator{contract: _BridgeStub.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeStub *BridgeStubFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *BridgeStubMessageDelivered, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubMessageDelivered)
				if err := _BridgeStub.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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

// ParseMessageDelivered is a log parse operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeStub *BridgeStubFilterer) ParseMessageDelivered(log types.Log) (*BridgeStubMessageDelivered, error) {
	event := new(BridgeStubMessageDelivered)
	if err := _BridgeStub.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeStubOutboxToggleIterator is returned from FilterOutboxToggle and is used to iterate over the raw logs and unpacked data for OutboxToggle events raised by the BridgeStub contract.
type BridgeStubOutboxToggleIterator struct {
	Event *BridgeStubOutboxToggle // Event containing the contract specifics and raw log

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
func (it *BridgeStubOutboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubOutboxToggle)
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
		it.Event = new(BridgeStubOutboxToggle)
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
func (it *BridgeStubOutboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubOutboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubOutboxToggle represents a OutboxToggle event raised by the BridgeStub contract.
type BridgeStubOutboxToggle struct {
	Outbox  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOutboxToggle is a free log retrieval operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) FilterOutboxToggle(opts *bind.FilterOpts, outbox []common.Address) (*BridgeStubOutboxToggleIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeStubOutboxToggleIterator{contract: _BridgeStub.contract, event: "OutboxToggle", logs: logs, sub: sub}, nil
}

// WatchOutboxToggle is a free log subscription operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) WatchOutboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeStubOutboxToggle, outbox []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubOutboxToggle)
				if err := _BridgeStub.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
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

// ParseOutboxToggle is a log parse operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) ParseOutboxToggle(log types.Log) (*BridgeStubOutboxToggle, error) {
	event := new(BridgeStubOutboxToggle)
	if err := _BridgeStub.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeStubSequencerInboxUpdatedIterator is returned from FilterSequencerInboxUpdated and is used to iterate over the raw logs and unpacked data for SequencerInboxUpdated events raised by the BridgeStub contract.
type BridgeStubSequencerInboxUpdatedIterator struct {
	Event *BridgeStubSequencerInboxUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeStubSequencerInboxUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubSequencerInboxUpdated)
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
		it.Event = new(BridgeStubSequencerInboxUpdated)
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
func (it *BridgeStubSequencerInboxUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubSequencerInboxUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubSequencerInboxUpdated represents a SequencerInboxUpdated event raised by the BridgeStub contract.
type BridgeStubSequencerInboxUpdated struct {
	NewSequencerInbox common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSequencerInboxUpdated is a free log retrieval operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeStub *BridgeStubFilterer) FilterSequencerInboxUpdated(opts *bind.FilterOpts) (*BridgeStubSequencerInboxUpdatedIterator, error) {

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeStubSequencerInboxUpdatedIterator{contract: _BridgeStub.contract, event: "SequencerInboxUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerInboxUpdated is a free log subscription operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeStub *BridgeStubFilterer) WatchSequencerInboxUpdated(opts *bind.WatchOpts, sink chan<- *BridgeStubSequencerInboxUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubSequencerInboxUpdated)
				if err := _BridgeStub.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
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

// ParseSequencerInboxUpdated is a log parse operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeStub *BridgeStubFilterer) ParseSequencerInboxUpdated(log types.Log) (*BridgeStubSequencerInboxUpdated, error) {
	event := new(BridgeStubSequencerInboxUpdated)
	if err := _BridgeStub.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxStubMetaData contains all meta data concerning the InboxStub contract.
var InboxStubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"createRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"unsafeCreateRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061082c806100206000396000f3fe6080604052600436106100bd5760003560e01c806367ef3ab81161006f57806367ef3ab8146101785780636e6e8a6a1461016a5780638a631aa614610186578063b75436bb146101a1578063c474d2c5146101c1578063c4d66de8146101e1578063e78cea921461020157600080fd5b80630f4d14e9146100c25780631fe927cf146100e8578063439370b1146101085780635075788b146101105780635c975abb1461012b5780635e9167581461015c578063679b6ded1461016a575b600080fd5b6100d56100d0366004610474565b610239565b6040519081526020015b60405180910390f35b3480156100f457600080fd5b506100d56101033660046104d5565b61027b565b6100d5610239565b34801561011c57600080fd5b506100d56100d036600461052b565b34801561013757600080fd5b5060005461014c90600160a01b900460ff1681565b60405190151581526020016100df565b6100d56100d03660046105a7565b6100d56100d0366004610610565b6100d56100d03660046106b4565b34801561019257600080fd5b506100d56100d0366004610726565b3480156101ad57600080fd5b506100d56101bc3660046104d5565b610314565b3480156101cd57600080fd5b506101df6101dc36600461077a565b50565b005b3480156101ed57600080fd5b506101df6101fc36600461077a565b610370565b34801561020d57600080fd5b50600054610221906001600160a01b031681565b6040516001600160a01b0390911681526020016100df565b60405162461bcd60e51b815260206004820152600f60248201526e1393d517d253541311535153951151608a1b60448201526000906064015b60405180910390fd5b60003332146102ba5760405162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b6044820152606401610272565b60006102df60033386866040516102d292919061079e565b60405180910390206103da565b60405190915081907fab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c90600090a29392505050565b60008061032d60033386866040516102d292919061079e565b9050807fff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b85856040516103619291906107ae565b60405180910390a29392505050565b6000546001600160a01b0316156103b85760405162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b6044820152606401610272565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b60008054604051638db5993b60e01b815260ff861660048201526001600160a01b0385811660248301526044820185905290911690638db5993b9034906064016020604051808303818588803b15801561043357600080fd5b505af1158015610447573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061046c91906107dd565b949350505050565b60006020828403121561048657600080fd5b5035919050565b60008083601f84011261049f57600080fd5b5081356001600160401b038111156104b657600080fd5b6020830191508360208285010111156104ce57600080fd5b9250929050565b600080602083850312156104e857600080fd5b82356001600160401b038111156104fe57600080fd5b61050a8582860161048d565b90969095509350505050565b6001600160a01b03811681146101dc57600080fd5b600080600080600080600060c0888a03121561054657600080fd5b873596506020880135955060408801359450606088013561056681610516565b93506080880135925060a08801356001600160401b0381111561058857600080fd5b6105948a828b0161048d565b989b979a50959850939692959293505050565b6000806000806000608086880312156105bf57600080fd5b853594506020860135935060408601356105d881610516565b925060608601356001600160401b038111156105f357600080fd5b6105ff8882890161048d565b969995985093965092949392505050565b60008060008060008060008060006101008a8c03121561062f57600080fd5b893561063a81610516565b985060208a0135975060408a0135965060608a013561065881610516565b955060808a013561066881610516565b945060a08a0135935060c08a0135925060e08a01356001600160401b0381111561069157600080fd5b61069d8c828d0161048d565b915080935050809150509295985092959850929598565b60008060008060008060a087890312156106cd57600080fd5b86359550602087013594506040870135935060608701356106ed81610516565b925060808701356001600160401b0381111561070857600080fd5b61071489828a0161048d565b979a9699509497509295939492505050565b60008060008060008060a0878903121561073f57600080fd5b8635955060208701359450604087013561075881610516565b93506060870135925060808701356001600160401b0381111561070857600080fd5b60006020828403121561078c57600080fd5b813561079781610516565b9392505050565b8183823760009101908152919050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b6000602082840312156107ef57600080fd5b505191905056fea2646970667358221220c0a9d1034f8f30ca5241a9f4d95735db7439aa946a17e05972ea21d9dbc08b5964736f6c63430008090033",
}

// InboxStubABI is the input ABI used to generate the binding from.
// Deprecated: Use InboxStubMetaData.ABI instead.
var InboxStubABI = InboxStubMetaData.ABI

// InboxStubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InboxStubMetaData.Bin instead.
var InboxStubBin = InboxStubMetaData.Bin

// DeployInboxStub deploys a new Ethereum contract, binding an instance of InboxStub to it.
func DeployInboxStub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InboxStub, error) {
	parsed, err := InboxStubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InboxStubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InboxStub{InboxStubCaller: InboxStubCaller{contract: contract}, InboxStubTransactor: InboxStubTransactor{contract: contract}, InboxStubFilterer: InboxStubFilterer{contract: contract}}, nil
}

// InboxStub is an auto generated Go binding around an Ethereum contract.
type InboxStub struct {
	InboxStubCaller     // Read-only binding to the contract
	InboxStubTransactor // Write-only binding to the contract
	InboxStubFilterer   // Log filterer for contract events
}

// InboxStubCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxStubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxStubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxStubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxStubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxStubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxStubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxStubSession struct {
	Contract     *InboxStub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxStubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxStubCallerSession struct {
	Contract *InboxStubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// InboxStubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxStubTransactorSession struct {
	Contract     *InboxStubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// InboxStubRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxStubRaw struct {
	Contract *InboxStub // Generic contract binding to access the raw methods on
}

// InboxStubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxStubCallerRaw struct {
	Contract *InboxStubCaller // Generic read-only contract binding to access the raw methods on
}

// InboxStubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxStubTransactorRaw struct {
	Contract *InboxStubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInboxStub creates a new instance of InboxStub, bound to a specific deployed contract.
func NewInboxStub(address common.Address, backend bind.ContractBackend) (*InboxStub, error) {
	contract, err := bindInboxStub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InboxStub{InboxStubCaller: InboxStubCaller{contract: contract}, InboxStubTransactor: InboxStubTransactor{contract: contract}, InboxStubFilterer: InboxStubFilterer{contract: contract}}, nil
}

// NewInboxStubCaller creates a new read-only instance of InboxStub, bound to a specific deployed contract.
func NewInboxStubCaller(address common.Address, caller bind.ContractCaller) (*InboxStubCaller, error) {
	contract, err := bindInboxStub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxStubCaller{contract: contract}, nil
}

// NewInboxStubTransactor creates a new write-only instance of InboxStub, bound to a specific deployed contract.
func NewInboxStubTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxStubTransactor, error) {
	contract, err := bindInboxStub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxStubTransactor{contract: contract}, nil
}

// NewInboxStubFilterer creates a new log filterer instance of InboxStub, bound to a specific deployed contract.
func NewInboxStubFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxStubFilterer, error) {
	contract, err := bindInboxStub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxStubFilterer{contract: contract}, nil
}

// bindInboxStub binds a generic wrapper to an already deployed contract.
func bindInboxStub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxStubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InboxStub *InboxStubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InboxStub.Contract.InboxStubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InboxStub *InboxStubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxStub.Contract.InboxStubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InboxStub *InboxStubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InboxStub.Contract.InboxStubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InboxStub *InboxStubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InboxStub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InboxStub *InboxStubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxStub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InboxStub *InboxStubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InboxStub.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_InboxStub *InboxStubCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_InboxStub *InboxStubSession) Bridge() (common.Address, error) {
	return _InboxStub.Contract.Bridge(&_InboxStub.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_InboxStub *InboxStubCallerSession) Bridge() (common.Address, error) {
	return _InboxStub.Contract.Bridge(&_InboxStub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InboxStub *InboxStubCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InboxStub *InboxStubSession) Paused() (bool, error) {
	return _InboxStub.Contract.Paused(&_InboxStub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InboxStub *InboxStubCallerSession) Paused() (bool, error) {
	return _InboxStub.Contract.Paused(&_InboxStub.CallOpts)
}

// SendContractTransaction is a free data retrieval call binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCaller) SendContractTransaction(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*big.Int, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "sendContractTransaction", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SendContractTransaction is a free data retrieval call binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubSession) SendContractTransaction(arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendContractTransaction(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendContractTransaction is a free data retrieval call binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCallerSession) SendContractTransaction(arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendContractTransaction(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendUnsignedTransaction is a free data retrieval call binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCaller) SendUnsignedTransaction(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "sendUnsignedTransaction", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SendUnsignedTransaction is a free data retrieval call binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubSession) SendUnsignedTransaction(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendUnsignedTransaction(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// SendUnsignedTransaction is a free data retrieval call binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCallerSession) SendUnsignedTransaction(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendUnsignedTransaction(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) CreateRetryableTicket(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "createRetryableTicket", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubSession) CreateRetryableTicket(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.CreateRetryableTicket(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) CreateRetryableTicket(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.CreateRetryableTicket(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) DepositEth(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "depositEth", arg0)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 ) payable returns(uint256)
func (_InboxStub *InboxStubSession) DepositEth(arg0 *big.Int) (*types.Transaction, error) {
	return _InboxStub.Contract.DepositEth(&_InboxStub.TransactOpts, arg0)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) DepositEth(arg0 *big.Int) (*types.Transaction, error) {
	return _InboxStub.Contract.DepositEth(&_InboxStub.TransactOpts, arg0)
}

// DepositEth0 is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_InboxStub *InboxStubTransactor) DepositEth0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "depositEth0")
}

// DepositEth0 is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_InboxStub *InboxStubSession) DepositEth0() (*types.Transaction, error) {
	return _InboxStub.Contract.DepositEth0(&_InboxStub.TransactOpts)
}

// DepositEth0 is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) DepositEth0() (*types.Transaction, error) {
	return _InboxStub.Contract.DepositEth0(&_InboxStub.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_InboxStub *InboxStubTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "initialize", _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_InboxStub *InboxStubSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _InboxStub.Contract.Initialize(&_InboxStub.TransactOpts, _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_InboxStub *InboxStubTransactorSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _InboxStub.Contract.Initialize(&_InboxStub.TransactOpts, _bridge)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_InboxStub *InboxStubTransactor) PostUpgradeInit(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "postUpgradeInit", _bridge)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_InboxStub *InboxStubSession) PostUpgradeInit(_bridge common.Address) (*types.Transaction, error) {
	return _InboxStub.Contract.PostUpgradeInit(&_InboxStub.TransactOpts, _bridge)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_InboxStub *InboxStubTransactorSession) PostUpgradeInit(_bridge common.Address) (*types.Transaction, error) {
	return _InboxStub.Contract.PostUpgradeInit(&_InboxStub.TransactOpts, _bridge)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) SendL1FundedContractTransaction(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "sendL1FundedContractTransaction", arg0, arg1, arg2, arg3)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubSession) SendL1FundedContractTransaction(arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedContractTransaction(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) SendL1FundedContractTransaction(arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedContractTransaction(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) SendL1FundedUnsignedTransaction(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "sendL1FundedUnsignedTransaction", arg0, arg1, arg2, arg3, arg4)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubSession) SendL1FundedUnsignedTransaction(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedUnsignedTransaction(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) SendL1FundedUnsignedTransaction(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedUnsignedTransaction(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL2Message(&_InboxStub.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL2Message(&_InboxStub.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubTransactor) SendL2MessageFromOrigin(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "sendL2MessageFromOrigin", messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL2MessageFromOrigin(&_InboxStub.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubTransactorSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL2MessageFromOrigin(&_InboxStub.TransactOpts, messageData)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) UnsafeCreateRetryableTicket(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "unsafeCreateRetryableTicket", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubSession) UnsafeCreateRetryableTicket(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.UnsafeCreateRetryableTicket(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) UnsafeCreateRetryableTicket(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.UnsafeCreateRetryableTicket(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// InboxStubInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the InboxStub contract.
type InboxStubInboxMessageDeliveredIterator struct {
	Event *InboxStubInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *InboxStubInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxStubInboxMessageDelivered)
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
		it.Event = new(InboxStubInboxMessageDelivered)
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
func (it *InboxStubInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxStubInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxStubInboxMessageDelivered represents a InboxMessageDelivered event raised by the InboxStub contract.
type InboxStubInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_InboxStub *InboxStubFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxStubInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _InboxStub.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxStubInboxMessageDeliveredIterator{contract: _InboxStub.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_InboxStub *InboxStubFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *InboxStubInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _InboxStub.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxStubInboxMessageDelivered)
				if err := _InboxStub.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_InboxStub *InboxStubFilterer) ParseInboxMessageDelivered(log types.Log) (*InboxStubInboxMessageDelivered, error) {
	event := new(InboxStubInboxMessageDelivered)
	if err := _InboxStub.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxStubInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the InboxStub contract.
type InboxStubInboxMessageDeliveredFromOriginIterator struct {
	Event *InboxStubInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *InboxStubInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxStubInboxMessageDeliveredFromOrigin)
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
		it.Event = new(InboxStubInboxMessageDeliveredFromOrigin)
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
func (it *InboxStubInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxStubInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxStubInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the InboxStub contract.
type InboxStubInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_InboxStub *InboxStubFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxStubInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _InboxStub.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxStubInboxMessageDeliveredFromOriginIterator{contract: _InboxStub.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_InboxStub *InboxStubFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *InboxStubInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _InboxStub.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxStubInboxMessageDeliveredFromOrigin)
				if err := _InboxStub.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_InboxStub *InboxStubFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*InboxStubInboxMessageDeliveredFromOrigin, error) {
	event := new(InboxStubInboxMessageDeliveredFromOrigin)
	if err := _InboxStub.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockResultReceiverMetaData contains all meta data concerning the MockResultReceiver contract.
var MockResultReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIChallengeManager\",\"name\":\"manager_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"challengeIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"ChallengeCompleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"challengeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengeIndex_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser_\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"startAndEndMachineStatuses_\",\"type\":\"uint8[2]\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState[2]\",\"name\":\"startAndEndGlobalStates_\",\"type\":\"tuple[2]\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"asserter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"asserterTimeLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengerTimeLeft_\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"contractIChallengeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161050f38038061050f83398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b61047c806100936000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630357aa491461006757806314eab5e71461007c578063481c6a75146100ad578063d6853748146100d8578063dfbf53ae146100ef578063e82898b314610102575b600080fd5b61007a610075366004610235565b610115565b005b61008f61008a36600461028a565b61017b565b60405167ffffffffffffffff90911681526020015b60405180910390f35b6000546100c0906001600160a01b031681565b6040516001600160a01b0390911681526020016100a4565b6100e160035481565b6040519081526020016100a4565b6001546100c0906001600160a01b031681565b6002546100c0906001600160a01b031681565b600180546001600160a01b03199081166001600160a01b0385811691821790935560028054909216928416928317909155600385905560405185907f88cb1f3fe351f3ac338db9c36bff1ece1750423c7ae6dfc427cd194b1c69b12790600090a4505050565b600080546040516314eab5e760e01b81526001600160a01b03909116906314eab5e7906101ba908c908c908c908c908c908c908c908c90600401610395565b602060405180830381600087803b1580156101d457600080fd5b505af11580156101e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020c9190610422565b9998505050505050505050565b80356001600160a01b038116811461023057600080fd5b919050565b60008060006060848603121561024a57600080fd5b8335925061025a60208501610219565b915061026860408501610219565b90509250925092565b67ffffffffffffffff8116811461028757600080fd5b50565b600080600080600080600080610200898b0312156102a757600080fd5b88359750606089018a8111156102bc57600080fd5b60208a0197506101608a018b8111156102d457600080fd5b909650356102e181610271565b94506102f06101808a01610219565b93506102ff6101a08a01610219565b92506101c089013591506101e089013590509295985092959890939650565b806000805b6002808210610332575061038e565b6040808588378681018481529085019084905b8382101561037757823561035881610271565b67ffffffffffffffff1681526020928301926001929092019101610345565b505050608095860195939093019250600101610323565b5050505050565b888152610200810160208083018a6000805b60028110156103d0578235600481106103be578283fd5b845292840192918401916001016103a7565b50505050506103e2606083018961031e565b67ffffffffffffffff969096166101608201526001600160a01b03948516610180820152929093166101a08301526101c08201526101e001529392505050565b60006020828403121561043457600080fd5b815161043f81610271565b939250505056fea2646970667358221220037a0493ffc4d0973a2c5f49249627ee37a71e136a7d427afa3a9e46638e789164736f6c63430008090033",
}

// MockResultReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use MockResultReceiverMetaData.ABI instead.
var MockResultReceiverABI = MockResultReceiverMetaData.ABI

// MockResultReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockResultReceiverMetaData.Bin instead.
var MockResultReceiverBin = MockResultReceiverMetaData.Bin

// DeployMockResultReceiver deploys a new Ethereum contract, binding an instance of MockResultReceiver to it.
func DeployMockResultReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, manager_ common.Address) (common.Address, *types.Transaction, *MockResultReceiver, error) {
	parsed, err := MockResultReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockResultReceiverBin), backend, manager_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockResultReceiver{MockResultReceiverCaller: MockResultReceiverCaller{contract: contract}, MockResultReceiverTransactor: MockResultReceiverTransactor{contract: contract}, MockResultReceiverFilterer: MockResultReceiverFilterer{contract: contract}}, nil
}

// MockResultReceiver is an auto generated Go binding around an Ethereum contract.
type MockResultReceiver struct {
	MockResultReceiverCaller     // Read-only binding to the contract
	MockResultReceiverTransactor // Write-only binding to the contract
	MockResultReceiverFilterer   // Log filterer for contract events
}

// MockResultReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockResultReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockResultReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockResultReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockResultReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockResultReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockResultReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockResultReceiverSession struct {
	Contract     *MockResultReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MockResultReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockResultReceiverCallerSession struct {
	Contract *MockResultReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MockResultReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockResultReceiverTransactorSession struct {
	Contract     *MockResultReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MockResultReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockResultReceiverRaw struct {
	Contract *MockResultReceiver // Generic contract binding to access the raw methods on
}

// MockResultReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockResultReceiverCallerRaw struct {
	Contract *MockResultReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// MockResultReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockResultReceiverTransactorRaw struct {
	Contract *MockResultReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockResultReceiver creates a new instance of MockResultReceiver, bound to a specific deployed contract.
func NewMockResultReceiver(address common.Address, backend bind.ContractBackend) (*MockResultReceiver, error) {
	contract, err := bindMockResultReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockResultReceiver{MockResultReceiverCaller: MockResultReceiverCaller{contract: contract}, MockResultReceiverTransactor: MockResultReceiverTransactor{contract: contract}, MockResultReceiverFilterer: MockResultReceiverFilterer{contract: contract}}, nil
}

// NewMockResultReceiverCaller creates a new read-only instance of MockResultReceiver, bound to a specific deployed contract.
func NewMockResultReceiverCaller(address common.Address, caller bind.ContractCaller) (*MockResultReceiverCaller, error) {
	contract, err := bindMockResultReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockResultReceiverCaller{contract: contract}, nil
}

// NewMockResultReceiverTransactor creates a new write-only instance of MockResultReceiver, bound to a specific deployed contract.
func NewMockResultReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*MockResultReceiverTransactor, error) {
	contract, err := bindMockResultReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockResultReceiverTransactor{contract: contract}, nil
}

// NewMockResultReceiverFilterer creates a new log filterer instance of MockResultReceiver, bound to a specific deployed contract.
func NewMockResultReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*MockResultReceiverFilterer, error) {
	contract, err := bindMockResultReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockResultReceiverFilterer{contract: contract}, nil
}

// bindMockResultReceiver binds a generic wrapper to an already deployed contract.
func bindMockResultReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockResultReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockResultReceiver *MockResultReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockResultReceiver.Contract.MockResultReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockResultReceiver *MockResultReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.MockResultReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockResultReceiver *MockResultReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.MockResultReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockResultReceiver *MockResultReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockResultReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockResultReceiver *MockResultReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockResultReceiver *MockResultReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.contract.Transact(opts, method, params...)
}

// ChallengeIndex is a free data retrieval call binding the contract method 0xd6853748.
//
// Solidity: function challengeIndex() view returns(uint256)
func (_MockResultReceiver *MockResultReceiverCaller) ChallengeIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockResultReceiver.contract.Call(opts, &out, "challengeIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeIndex is a free data retrieval call binding the contract method 0xd6853748.
//
// Solidity: function challengeIndex() view returns(uint256)
func (_MockResultReceiver *MockResultReceiverSession) ChallengeIndex() (*big.Int, error) {
	return _MockResultReceiver.Contract.ChallengeIndex(&_MockResultReceiver.CallOpts)
}

// ChallengeIndex is a free data retrieval call binding the contract method 0xd6853748.
//
// Solidity: function challengeIndex() view returns(uint256)
func (_MockResultReceiver *MockResultReceiverCallerSession) ChallengeIndex() (*big.Int, error) {
	return _MockResultReceiver.Contract.ChallengeIndex(&_MockResultReceiver.CallOpts)
}

// Loser is a free data retrieval call binding the contract method 0xe82898b3.
//
// Solidity: function loser() view returns(address)
func (_MockResultReceiver *MockResultReceiverCaller) Loser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockResultReceiver.contract.Call(opts, &out, "loser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Loser is a free data retrieval call binding the contract method 0xe82898b3.
//
// Solidity: function loser() view returns(address)
func (_MockResultReceiver *MockResultReceiverSession) Loser() (common.Address, error) {
	return _MockResultReceiver.Contract.Loser(&_MockResultReceiver.CallOpts)
}

// Loser is a free data retrieval call binding the contract method 0xe82898b3.
//
// Solidity: function loser() view returns(address)
func (_MockResultReceiver *MockResultReceiverCallerSession) Loser() (common.Address, error) {
	return _MockResultReceiver.Contract.Loser(&_MockResultReceiver.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_MockResultReceiver *MockResultReceiverCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockResultReceiver.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_MockResultReceiver *MockResultReceiverSession) Manager() (common.Address, error) {
	return _MockResultReceiver.Contract.Manager(&_MockResultReceiver.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_MockResultReceiver *MockResultReceiverCallerSession) Manager() (common.Address, error) {
	return _MockResultReceiver.Contract.Manager(&_MockResultReceiver.CallOpts)
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_MockResultReceiver *MockResultReceiverCaller) Winner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockResultReceiver.contract.Call(opts, &out, "winner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_MockResultReceiver *MockResultReceiverSession) Winner() (common.Address, error) {
	return _MockResultReceiver.Contract.Winner(&_MockResultReceiver.CallOpts)
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_MockResultReceiver *MockResultReceiverCallerSession) Winner() (common.Address, error) {
	return _MockResultReceiver.Contract.Winner(&_MockResultReceiver.CallOpts)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex_, address winner_, address loser_) returns()
func (_MockResultReceiver *MockResultReceiverTransactor) CompleteChallenge(opts *bind.TransactOpts, challengeIndex_ *big.Int, winner_ common.Address, loser_ common.Address) (*types.Transaction, error) {
	return _MockResultReceiver.contract.Transact(opts, "completeChallenge", challengeIndex_, winner_, loser_)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex_, address winner_, address loser_) returns()
func (_MockResultReceiver *MockResultReceiverSession) CompleteChallenge(challengeIndex_ *big.Int, winner_ common.Address, loser_ common.Address) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.CompleteChallenge(&_MockResultReceiver.TransactOpts, challengeIndex_, winner_, loser_)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex_, address winner_, address loser_) returns()
func (_MockResultReceiver *MockResultReceiverTransactorSession) CompleteChallenge(challengeIndex_ *big.Int, winner_ common.Address, loser_ common.Address) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.CompleteChallenge(&_MockResultReceiver.TransactOpts, challengeIndex_, winner_, loser_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_MockResultReceiver *MockResultReceiverTransactor) CreateChallenge(opts *bind.TransactOpts, wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _MockResultReceiver.contract.Transact(opts, "createChallenge", wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_MockResultReceiver *MockResultReceiverSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.CreateChallenge(&_MockResultReceiver.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_MockResultReceiver *MockResultReceiverTransactorSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.CreateChallenge(&_MockResultReceiver.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// MockResultReceiverChallengeCompletedIterator is returned from FilterChallengeCompleted and is used to iterate over the raw logs and unpacked data for ChallengeCompleted events raised by the MockResultReceiver contract.
type MockResultReceiverChallengeCompletedIterator struct {
	Event *MockResultReceiverChallengeCompleted // Event containing the contract specifics and raw log

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
func (it *MockResultReceiverChallengeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockResultReceiverChallengeCompleted)
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
		it.Event = new(MockResultReceiverChallengeCompleted)
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
func (it *MockResultReceiverChallengeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockResultReceiverChallengeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockResultReceiverChallengeCompleted represents a ChallengeCompleted event raised by the MockResultReceiver contract.
type MockResultReceiverChallengeCompleted struct {
	ChallengeIndex *big.Int
	Winner         common.Address
	Loser          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChallengeCompleted is a free log retrieval operation binding the contract event 0x88cb1f3fe351f3ac338db9c36bff1ece1750423c7ae6dfc427cd194b1c69b127.
//
// Solidity: event ChallengeCompleted(uint256 indexed challengeIndex, address indexed winner, address indexed loser)
func (_MockResultReceiver *MockResultReceiverFilterer) FilterChallengeCompleted(opts *bind.FilterOpts, challengeIndex []*big.Int, winner []common.Address, loser []common.Address) (*MockResultReceiverChallengeCompletedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var loserRule []interface{}
	for _, loserItem := range loser {
		loserRule = append(loserRule, loserItem)
	}

	logs, sub, err := _MockResultReceiver.contract.FilterLogs(opts, "ChallengeCompleted", challengeIndexRule, winnerRule, loserRule)
	if err != nil {
		return nil, err
	}
	return &MockResultReceiverChallengeCompletedIterator{contract: _MockResultReceiver.contract, event: "ChallengeCompleted", logs: logs, sub: sub}, nil
}

// WatchChallengeCompleted is a free log subscription operation binding the contract event 0x88cb1f3fe351f3ac338db9c36bff1ece1750423c7ae6dfc427cd194b1c69b127.
//
// Solidity: event ChallengeCompleted(uint256 indexed challengeIndex, address indexed winner, address indexed loser)
func (_MockResultReceiver *MockResultReceiverFilterer) WatchChallengeCompleted(opts *bind.WatchOpts, sink chan<- *MockResultReceiverChallengeCompleted, challengeIndex []*big.Int, winner []common.Address, loser []common.Address) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var loserRule []interface{}
	for _, loserItem := range loser {
		loserRule = append(loserRule, loserItem)
	}

	logs, sub, err := _MockResultReceiver.contract.WatchLogs(opts, "ChallengeCompleted", challengeIndexRule, winnerRule, loserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockResultReceiverChallengeCompleted)
				if err := _MockResultReceiver.contract.UnpackLog(event, "ChallengeCompleted", log); err != nil {
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

// ParseChallengeCompleted is a log parse operation binding the contract event 0x88cb1f3fe351f3ac338db9c36bff1ece1750423c7ae6dfc427cd194b1c69b127.
//
// Solidity: event ChallengeCompleted(uint256 indexed challengeIndex, address indexed winner, address indexed loser)
func (_MockResultReceiver *MockResultReceiverFilterer) ParseChallengeCompleted(log types.Log) (*MockResultReceiverChallengeCompleted, error) {
	event := new(MockResultReceiverChallengeCompleted)
	if err := _MockResultReceiver.contract.UnpackLog(event, "ChallengeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubMetaData contains all meta data concerning the SequencerInboxStub contract.
var SequencerInboxStubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation_\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"AlreadyValidDASKeyset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"BadSequencerNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataNotAuthenticated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDataLength\",\"type\":\"uint256\"}],\"name\":\"DataTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayedBackwards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayedTooFar\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceIncludeBlockTooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceIncludeTimeTooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HadZeroInit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectMessagePreimage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"NoSuchKeyset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBatchPoster\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOrigin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keysetHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidateKeyset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"OwnerFunctionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchSequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SequencerBatchData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchSequenceNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxBlockNumber\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structISequencerInbox.TimeBounds\",\"name\":\"timeBounds\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"enumISequencerInbox.BatchDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"}],\"name\":\"SequencerBatchDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keysetHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"keysetBytes\",\"type\":\"bytes\"}],\"name\":\"SetValidKeyset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DATA_AUTHENTICATED_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HEADER_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addInitMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"}],\"name\":\"addSequencerL2Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"}],\"name\":\"addSequencerL2BatchFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dasKeySetInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidKeyset\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"creationBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint64[2]\",\"name\":\"l1BlockAndTime\",\"type\":\"uint64[2]\"},{\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"forceInclusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"getKeysetCreationBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"inboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation_\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"invalidateKeysetHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBatchPoster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"isValidKeysetHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTimeVariation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isBatchPoster_\",\"type\":\"bool\"}],\"name\":\"setIsBatchPoster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation_\",\"type\":\"tuple\"}],\"name\":\"setMaxTimeVariation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"keysetBytes\",\"type\":\"bytes\"}],\"name\":\"setValidKeyset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDelayedMessagesRead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052306080523480156200001557600080fd5b5060405162001f3b38038062001f3b8339810160408190526200003891620000bb565b600180546001600160a01b039485166001600160a01b03199182161782556002805490911633179055815160045560208083015160055560408084015160065560609093015160075592909316600090815260039092529020805460ff191690911790556200016e565b6001600160a01b0381168114620000b857600080fd5b50565b600080600083850360c0811215620000d257600080fd5b8451620000df81620000a2565b6020860151909450620000f281620000a2565b92506080603f19820112156200010757600080fd5b50604051608081016001600160401b03811182821017156200013957634e487b7160e01b600052604160045260246000fd5b806040525060408501518152606085015160208201526080850151604082015260a08501516060820152809150509250925092565b608051611db16200018a600039600061045f0152611db16000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c806371c3e6fe116100b8578063d1ce8da81161007c578063d1ce8da8146102ca578063d9dd67ab146102dd578063e5a358c8146102f0578063e78cea9214610314578063ebea461d14610327578063f19815781461035d57600080fd5b806371c3e6fe1461024d5780637fa3a40e146102705780638442086014610279578063b31761f81461028c578063cb23bcb51461029f57600080fd5b806327957a49116100ff57806327957a49146101ba57806361a93f8f146101c25780636e7df3e7146101d55780636f12b0c9146101e8578063715ea34b146101fb57600080fd5b806306f130561461013c5780631637be481461015757806317c9c5771461018a5780631f7a92b214610194578063258f0495146101a7575b600080fd5b610144610370565b6040519081526020015b60405180910390f35b61017a6101653660046117c9565b60009081526008602052604090205460ff1690565b604051901515815260200161014e565b6101926103f0565b005b6101926101a23660046117fa565b610454565b6101446101b53660046117c9565b61060c565b610144602881565b6101926101d0366004611883565b610675565b6101926101e33660046118fb565b610854565b6101926101f6366004611883565b61095a565b61022e6102093660046117c9565b60086020526000908152604090205460ff81169061010090046001600160401b031682565b6040805192151583526001600160401b0390911660208301520161014e565b61017a61025b366004611934565b60036020526000908152604090205460ff1681565b61014460005481565b6101926102873660046117c9565b610a74565b61019261029a366004611958565b610bc7565b6002546102b2906001600160a01b031681565b6040516001600160a01b03909116815260200161014e565b6101926102d83660046119cb565b610cc7565b6101446102eb3660046117c9565b610f0f565b6102fb600160fe1b81565b6040516001600160f81b0319909116815260200161014e565b6001546102b2906001600160a01b031681565b60045460055460065460075461033d9392919084565b60408051948552602085019390935291830152606082015260800161014e565b61019261036b366004611a0c565b610f92565b600154604080516221048360e21b815290516000926001600160a01b0316916284120c916004808301926020929190829003018186803b1580156103b357600080fd5b505afa1580156103c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103eb9190611a7c565b905090565b6000806103fd60006112c5565b91509150600080600080610413866000806112f1565b9350935093509350808385600080516020611d3c833981519152856000548a60026040516104449493929190611a95565b60405180910390a4505050505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156104e75760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b60648201526084015b60405180910390fd5b6001546001600160a01b03161561051157604051633bcd329760e21b815260040160405180910390fd5b6001600160a01b03821661053857604051631ad0f74360e01b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0384169081179091556040805163cb23bcb560e01b8152905163cb23bcb591600480820192602092909190829003018186803b15801561058c57600080fd5b505afa1580156105a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c49190611b09565b600280546001600160a01b0319166001600160a01b03929092169190911790558035600490815560208201356005556040820135600655606082013560075581905b50505050565b600081815260086020908152604080832081518083019092525460ff81161515825261010090046001600160401b0316918101829052906106625760405162f20c5d60e01b8152600481018490526024016104de565b602001516001600160401b031692915050565b8060005a3360009081526003602052604090205490915060ff161580156106a757506002546001600160a01b03163314155b156106c557604051632dd9fc9760e01b815260040160405180910390fd5b6000806106d388888861157d565b915091506000806000806106e9868b60006112f1565b93509350935093508c841461071b5760405163ac7411c960e01b815260048101859052602481018e90526044016104de565b80838e600080516020611d3c833981519152858e8a60016040516107429493929190611a95565b60405180910390a48c7ffe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c208d8d60405161077c929190611b26565b60405180910390a2505050506001600160a01b03841615915061084b9050576000333214156107a85750365b826001600160a01b031663e3db8a49335a6107c39086611b6b565b6040516001600160e01b031960e085901b1681526001600160a01b039092166004830152602482015260448101849052606401602060405180830381600087803b15801561081057600080fd5b505af1158015610824573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108489190611b82565b50505b50505050505050565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156108a257600080fd5b505afa1580156108b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108da9190611b09565b6001600160a01b0316336001600160a01b03161461091c57600254604051631194af8760e11b81526104de9133916001600160a01b0390911690600401611b9f565b6001600160a01b038216600090815260036020526040808220805460ff191684151517905551600191600080516020611d5c83398151915291a25050565b8060005a90503332146109805760405163feb3d07160e01b815260040160405180910390fd5b3360009081526003602052604090205460ff166109b057604051632dd9fc9760e01b815260040160405180910390fd5b6000806109be88888861157d565b909250905060008080806109d3868b8d6112f1565b93509350935093508c8414610a055760405163ac7411c960e01b815260048101859052602481018e90526044016104de565b80838e600080516020611d3c833981519152856000548a6000604051610a2e9493929190611a95565b60405180910390a4505050506001600160a01b03841615915061084b9050576000333214156107a8575036826001600160a01b031663e3db8a49335a6107c39086611b6b565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610ac257600080fd5b505afa158015610ad6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610afa9190611b09565b6001600160a01b0316336001600160a01b031614610b3c57600254604051631194af8760e11b81526104de9133916001600160a01b0390911690600401611b9f565b60008181526008602052604090205460ff16610b6d5760405162f20c5d60e01b8152600481018290526024016104de565b600081815260086020526040808220805460ff191690555182917f5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a91a2604051600390600080516020611d5c83398151915290600090a250565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610c1557600080fd5b505afa158015610c29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c4d9190611b09565b6001600160a01b0316336001600160a01b031614610c8f57600254604051631194af8760e11b81526104de9133916001600160a01b0390911690600401611b9f565b80516004556020810151600555604080820151600655606082015160075551600090600080516020611d5c833981519152908290a250565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610d1557600080fd5b505afa158015610d29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d4d9190611b09565b6001600160a01b0316336001600160a01b031614610d8f57600254604051631194af8760e11b81526104de9133916001600160a01b0390911690600401611b9f565b60008282604051610da1929190611bb9565b604051908190038120607f60f91b6020830152602182015260410160408051601f1981840301815291905280516020909101209050600160ff1b8118620100008310610e255760405162461bcd60e51b81526020600482015260136024820152726b657973657420697320746f6f206c6172676560681b60448201526064016104de565b60008181526008602052604090205460ff1615610e5857604051637d17eeed60e11b8152600481018290526024016104de565b60408051808201825260018152436001600160401b0390811660208084019182526000868152600890915284902092518354915168ffffffffffffffffff1990921690151568ffffffffffffffff0019161761010091909216021790555181907fabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef65572290610ee79087908790611b26565b60405180910390a2604051600290600080516020611d5c83398151915290600090a250505050565b6001546040516316bf557960e01b8152600481018390526000916001600160a01b0316906316bf55799060240160206040518083038186803b158015610f5457600080fd5b505afa158015610f68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f8c9190611a7c565b92915050565b6000548611610fb457604051633eb9f37d60e11b815260040160405180910390fd5b60006110648684610fc86020890189611bdf565b610fd860408a0160208b01611bdf565b610fe360018d611b6b565b6040805160f89690961b6001600160f81b03191660208088019190915260609590951b6001600160601b031916602187015260c093841b6001600160c01b031990811660358801529290931b909116603d85015260458401526065830188905260858084018790528151808503909101815260a59093019052815191012090565b60045490915043906110796020880188611bdf565b6001600160401b031661108c9190611c08565b106110aa5760405163ad3515d960e01b815260040160405180910390fd5b60065442906110bf6040880160208901611bdf565b6001600160401b03166110d29190611c08565b106110f05760405163c76d17e560e01b815260040160405180910390fd5b60006001881115611188576001546001600160a01b031663d5719dc261111760028b611b6b565b6040518263ffffffff1660e01b815260040161113591815260200190565b60206040518083038186803b15801561114d57600080fd5b505afa158015611161573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111859190611a7c565b90505b60408051602080820184905281830185905282518083038401815260609092019092528051910120600180546001600160a01b03169063d5719dc2906111ce908c611b6b565b6040518263ffffffff1660e01b81526004016111ec91815260200190565b60206040518083038186803b15801561120457600080fd5b505afa158015611218573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123c9190611a7c565b1461125a576040516313947fd760e01b815260040160405180910390fd5b6000806112668a6112c5565b9150915060008060008061127c868f60006112f1565b9350935093509350808385600080516020611d3c833981519152856000548a60026040516112ad9493929190611a95565b60405180910390a45050505050505050505050505050565b60006112cf6117a2565b6000806112db856116f0565b8151602090920191909120969095509350505050565b60008060008060005486101561131a57604051633eb9f37d60e11b815260040160405180910390fd5b600160009054906101000a90046001600160a01b03166001600160a01b031663eca067ad6040518163ffffffff1660e01b815260040160206040518083038186803b15801561136857600080fd5b505afa15801561137c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113a09190611a7c565b8611156113c05760405163925f8bd360e01b815260040160405180910390fd5b60015460405163422def4160e11b815260048101899052602481018890526001600160a01b039091169063845bde8290604401608060405180830381600087803b15801561140d57600080fd5b505af1158015611421573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114459190611c20565b60008a9055929650909450925090508415611574576040805142602082015233606081901b6001600160601b03191692820192909252605481018990526074810186905248609482015260009060b40160408051808303601f190181529082905260015481516020830120637a88b10760e01b84526001600160a01b0386811660048601526024850191909152919350600092911690637a88b10790604401602060405180830381600087803b1580156114fe57600080fd5b505af1158015611512573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115369190611a7c565b9050807fff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b836040516115689190611c82565b60405180910390a25050505b93509350935093565b60006115876117a2565b84846000611596826028611c08565b90506201cccc8111156115c857604051634634691b60e01b8152600481018290526201cccc60248201526044016104de565b81158015906116005750600160fe1b8084846000816115e9576115e9611bc9565b9050013560f81c60f81b166001600160f81b031916145b1561161e57604051631f97007f60e01b815260040160405180910390fd5b6021821080159061164c57508282600081811061163d5761163d611bc9565b90910135600160ff1b16151590505b156116a0576000611661602160018587611cb5565b61166a91611cdf565b60008181526008602052604090205490915060ff1661169e5760405162f20c5d60e01b8152600481018290526024016104de565b505b6000806116ac886116f0565b915091506000828b8b6040516020016116c793929190611cfd565b60408051808303601f1901815291905280516020909101209b919a509098505050505050505050565b60606116fa6117a2565b6000611704611797565b9050600081600001518260200151836040015184606001518860405160200161176c95949392919060c095861b6001600160c01b0319908116825294861b8516600882015292851b8416601084015290841b8316601883015290921b16602082015260280190565b6040516020818303038152906040529050602881511461178e5761178e611d25565b94909350915050565b61179f6117a2565b90565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6000602082840312156117db57600080fd5b5035919050565b6001600160a01b03811681146117f757600080fd5b50565b60008082840360a081121561180e57600080fd5b8335611819816117e2565b92506080601f198201121561182d57600080fd5b506020830190509250929050565b60008083601f84011261184d57600080fd5b5081356001600160401b0381111561186457600080fd5b60208301915083602082850101111561187c57600080fd5b9250929050565b60008060008060006080868803121561189b57600080fd5b8535945060208601356001600160401b038111156118b857600080fd5b6118c48882890161183b565b9095509350506040860135915060608601356118df816117e2565b809150509295509295909350565b80151581146117f757600080fd5b6000806040838503121561190e57600080fd5b8235611919816117e2565b91506020830135611929816118ed565b809150509250929050565b60006020828403121561194657600080fd5b8135611951816117e2565b9392505050565b60006080828403121561196a57600080fd5b604051608081018181106001600160401b038211171561199a57634e487b7160e01b600052604160045260246000fd5b8060405250823581526020830135602082015260408301356040820152606083013560608201528091505092915050565b600080602083850312156119de57600080fd5b82356001600160401b038111156119f457600080fd5b611a008582860161183b565b90969095509350505050565b60008060008060008060e08789031215611a2557600080fd5b86359550602087013560ff81168114611a3d57600080fd5b94506080870188811115611a5057600080fd5b60408801945035925060a0870135611a67816117e2565b8092505060c087013590509295509295509295565b600060208284031215611a8e57600080fd5b5051919050565b600060e0820190508582528460208301526001600160401b038085511660408401528060208601511660608401528060408601511660808401528060608601511660a08401525060038310611afa57634e487b7160e01b600052602160045260246000fd5b8260c083015295945050505050565b600060208284031215611b1b57600080fd5b8151611951816117e2565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b634e487b7160e01b600052601160045260246000fd5b600082821015611b7d57611b7d611b55565b500390565b600060208284031215611b9457600080fd5b8151611951816118ed565b6001600160a01b0392831681529116602082015260400190565b8183823760009101908152919050565b634e487b7160e01b600052603260045260246000fd5b600060208284031215611bf157600080fd5b81356001600160401b038116811461195157600080fd5b60008219821115611c1b57611c1b611b55565b500190565b60008060008060808587031215611c3657600080fd5b505082516020840151604085015160609095015191969095509092509050565b60005b83811015611c71578181015183820152602001611c59565b838111156106065750506000910152565b6020815260008251806020840152611ca1816040850160208701611c56565b601f01601f19169190910160400192915050565b60008085851115611cc557600080fd5b83861115611cd257600080fd5b5050820193919092039150565b80356020831015610f8c57600019602084900360031b1b1692915050565b60008451611d0f818460208901611c56565b8201838582376000930192835250909392505050565b634e487b7160e01b600052600160045260246000fdfe7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7ea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456ea26469706673582212205be3d8f514750d0d72105b153b10cd4fa41c23fd7cee0fe8d30927edf84e6b8564736f6c63430008090033",
}

// SequencerInboxStubABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerInboxStubMetaData.ABI instead.
var SequencerInboxStubABI = SequencerInboxStubMetaData.ABI

// SequencerInboxStubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SequencerInboxStubMetaData.Bin instead.
var SequencerInboxStubBin = SequencerInboxStubMetaData.Bin

// DeploySequencerInboxStub deploys a new Ethereum contract, binding an instance of SequencerInboxStub to it.
func DeploySequencerInboxStub(auth *bind.TransactOpts, backend bind.ContractBackend, bridge_ common.Address, sequencer_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (common.Address, *types.Transaction, *SequencerInboxStub, error) {
	parsed, err := SequencerInboxStubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SequencerInboxStubBin), backend, bridge_, sequencer_, maxTimeVariation_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SequencerInboxStub{SequencerInboxStubCaller: SequencerInboxStubCaller{contract: contract}, SequencerInboxStubTransactor: SequencerInboxStubTransactor{contract: contract}, SequencerInboxStubFilterer: SequencerInboxStubFilterer{contract: contract}}, nil
}

// SequencerInboxStub is an auto generated Go binding around an Ethereum contract.
type SequencerInboxStub struct {
	SequencerInboxStubCaller     // Read-only binding to the contract
	SequencerInboxStubTransactor // Write-only binding to the contract
	SequencerInboxStubFilterer   // Log filterer for contract events
}

// SequencerInboxStubCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerInboxStubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxStubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerInboxStubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxStubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerInboxStubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxStubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerInboxStubSession struct {
	Contract     *SequencerInboxStub // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SequencerInboxStubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerInboxStubCallerSession struct {
	Contract *SequencerInboxStubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SequencerInboxStubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerInboxStubTransactorSession struct {
	Contract     *SequencerInboxStubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SequencerInboxStubRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerInboxStubRaw struct {
	Contract *SequencerInboxStub // Generic contract binding to access the raw methods on
}

// SequencerInboxStubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerInboxStubCallerRaw struct {
	Contract *SequencerInboxStubCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerInboxStubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerInboxStubTransactorRaw struct {
	Contract *SequencerInboxStubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencerInboxStub creates a new instance of SequencerInboxStub, bound to a specific deployed contract.
func NewSequencerInboxStub(address common.Address, backend bind.ContractBackend) (*SequencerInboxStub, error) {
	contract, err := bindSequencerInboxStub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStub{SequencerInboxStubCaller: SequencerInboxStubCaller{contract: contract}, SequencerInboxStubTransactor: SequencerInboxStubTransactor{contract: contract}, SequencerInboxStubFilterer: SequencerInboxStubFilterer{contract: contract}}, nil
}

// NewSequencerInboxStubCaller creates a new read-only instance of SequencerInboxStub, bound to a specific deployed contract.
func NewSequencerInboxStubCaller(address common.Address, caller bind.ContractCaller) (*SequencerInboxStubCaller, error) {
	contract, err := bindSequencerInboxStub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubCaller{contract: contract}, nil
}

// NewSequencerInboxStubTransactor creates a new write-only instance of SequencerInboxStub, bound to a specific deployed contract.
func NewSequencerInboxStubTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerInboxStubTransactor, error) {
	contract, err := bindSequencerInboxStub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubTransactor{contract: contract}, nil
}

// NewSequencerInboxStubFilterer creates a new log filterer instance of SequencerInboxStub, bound to a specific deployed contract.
func NewSequencerInboxStubFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerInboxStubFilterer, error) {
	contract, err := bindSequencerInboxStub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubFilterer{contract: contract}, nil
}

// bindSequencerInboxStub binds a generic wrapper to an already deployed contract.
func bindSequencerInboxStub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SequencerInboxStubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInboxStub *SequencerInboxStubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInboxStub.Contract.SequencerInboxStubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInboxStub *SequencerInboxStubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SequencerInboxStubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInboxStub *SequencerInboxStubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SequencerInboxStubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInboxStub *SequencerInboxStubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInboxStub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInboxStub *SequencerInboxStubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInboxStub *SequencerInboxStubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.contract.Transact(opts, method, params...)
}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCaller) DATAAUTHENTICATEDFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "DATA_AUTHENTICATED_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubSession) DATAAUTHENTICATEDFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.DATAAUTHENTICATEDFLAG(&_SequencerInboxStub.CallOpts)
}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) DATAAUTHENTICATEDFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.DATAAUTHENTICATEDFLAG(&_SequencerInboxStub.CallOpts)
}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCaller) HEADERLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "HEADER_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubSession) HEADERLENGTH() (*big.Int, error) {
	return _SequencerInboxStub.Contract.HEADERLENGTH(&_SequencerInboxStub.CallOpts)
}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) HEADERLENGTH() (*big.Int, error) {
	return _SequencerInboxStub.Contract.HEADERLENGTH(&_SequencerInboxStub.CallOpts)
}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCaller) BatchCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "batchCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubSession) BatchCount() (*big.Int, error) {
	return _SequencerInboxStub.Contract.BatchCount(&_SequencerInboxStub.CallOpts)
}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) BatchCount() (*big.Int, error) {
	return _SequencerInboxStub.Contract.BatchCount(&_SequencerInboxStub.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubSession) Bridge() (common.Address, error) {
	return _SequencerInboxStub.Contract.Bridge(&_SequencerInboxStub.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) Bridge() (common.Address, error) {
	return _SequencerInboxStub.Contract.Bridge(&_SequencerInboxStub.CallOpts)
}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(bool isValidKeyset, uint64 creationBlock)
func (_SequencerInboxStub *SequencerInboxStubCaller) DasKeySetInfo(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "dasKeySetInfo", arg0)

	outstruct := new(struct {
		IsValidKeyset bool
		CreationBlock uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsValidKeyset = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CreationBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(bool isValidKeyset, uint64 creationBlock)
func (_SequencerInboxStub *SequencerInboxStubSession) DasKeySetInfo(arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	return _SequencerInboxStub.Contract.DasKeySetInfo(&_SequencerInboxStub.CallOpts, arg0)
}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(bool isValidKeyset, uint64 creationBlock)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) DasKeySetInfo(arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	return _SequencerInboxStub.Contract.DasKeySetInfo(&_SequencerInboxStub.CallOpts, arg0)
}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCaller) GetKeysetCreationBlock(opts *bind.CallOpts, ksHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "getKeysetCreationBlock", ksHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubSession) GetKeysetCreationBlock(ksHash [32]byte) (*big.Int, error) {
	return _SequencerInboxStub.Contract.GetKeysetCreationBlock(&_SequencerInboxStub.CallOpts, ksHash)
}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) GetKeysetCreationBlock(ksHash [32]byte) (*big.Int, error) {
	return _SequencerInboxStub.Contract.GetKeysetCreationBlock(&_SequencerInboxStub.CallOpts, ksHash)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInboxStub *SequencerInboxStubCaller) InboxAccs(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "inboxAccs", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInboxStub *SequencerInboxStubSession) InboxAccs(index *big.Int) ([32]byte, error) {
	return _SequencerInboxStub.Contract.InboxAccs(&_SequencerInboxStub.CallOpts, index)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) InboxAccs(index *big.Int) ([32]byte, error) {
	return _SequencerInboxStub.Contract.InboxAccs(&_SequencerInboxStub.CallOpts, index)
}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCaller) IsBatchPoster(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "isBatchPoster", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubSession) IsBatchPoster(arg0 common.Address) (bool, error) {
	return _SequencerInboxStub.Contract.IsBatchPoster(&_SequencerInboxStub.CallOpts, arg0)
}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) IsBatchPoster(arg0 common.Address) (bool, error) {
	return _SequencerInboxStub.Contract.IsBatchPoster(&_SequencerInboxStub.CallOpts, arg0)
}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCaller) IsValidKeysetHash(opts *bind.CallOpts, ksHash [32]byte) (bool, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "isValidKeysetHash", ksHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubSession) IsValidKeysetHash(ksHash [32]byte) (bool, error) {
	return _SequencerInboxStub.Contract.IsValidKeysetHash(&_SequencerInboxStub.CallOpts, ksHash)
}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) IsValidKeysetHash(ksHash [32]byte) (bool, error) {
	return _SequencerInboxStub.Contract.IsValidKeysetHash(&_SequencerInboxStub.CallOpts, ksHash)
}

// MaxTimeVariation is a free data retrieval call binding the contract method 0xebea461d.
//
// Solidity: function maxTimeVariation() view returns(uint256 delayBlocks, uint256 futureBlocks, uint256 delaySeconds, uint256 futureSeconds)
func (_SequencerInboxStub *SequencerInboxStubCaller) MaxTimeVariation(opts *bind.CallOpts) (struct {
	DelayBlocks   *big.Int
	FutureBlocks  *big.Int
	DelaySeconds  *big.Int
	FutureSeconds *big.Int
}, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "maxTimeVariation")

	outstruct := new(struct {
		DelayBlocks   *big.Int
		FutureBlocks  *big.Int
		DelaySeconds  *big.Int
		FutureSeconds *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DelayBlocks = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FutureBlocks = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DelaySeconds = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FutureSeconds = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MaxTimeVariation is a free data retrieval call binding the contract method 0xebea461d.
//
// Solidity: function maxTimeVariation() view returns(uint256 delayBlocks, uint256 futureBlocks, uint256 delaySeconds, uint256 futureSeconds)
func (_SequencerInboxStub *SequencerInboxStubSession) MaxTimeVariation() (struct {
	DelayBlocks   *big.Int
	FutureBlocks  *big.Int
	DelaySeconds  *big.Int
	FutureSeconds *big.Int
}, error) {
	return _SequencerInboxStub.Contract.MaxTimeVariation(&_SequencerInboxStub.CallOpts)
}

// MaxTimeVariation is a free data retrieval call binding the contract method 0xebea461d.
//
// Solidity: function maxTimeVariation() view returns(uint256 delayBlocks, uint256 futureBlocks, uint256 delaySeconds, uint256 futureSeconds)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) MaxTimeVariation() (struct {
	DelayBlocks   *big.Int
	FutureBlocks  *big.Int
	DelaySeconds  *big.Int
	FutureSeconds *big.Int
}, error) {
	return _SequencerInboxStub.Contract.MaxTimeVariation(&_SequencerInboxStub.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubSession) Rollup() (common.Address, error) {
	return _SequencerInboxStub.Contract.Rollup(&_SequencerInboxStub.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) Rollup() (common.Address, error) {
	return _SequencerInboxStub.Contract.Rollup(&_SequencerInboxStub.CallOpts)
}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCaller) TotalDelayedMessagesRead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "totalDelayedMessagesRead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubSession) TotalDelayedMessagesRead() (*big.Int, error) {
	return _SequencerInboxStub.Contract.TotalDelayedMessagesRead(&_SequencerInboxStub.CallOpts)
}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) TotalDelayedMessagesRead() (*big.Int, error) {
	return _SequencerInboxStub.Contract.TotalDelayedMessagesRead(&_SequencerInboxStub.CallOpts)
}

// AddInitMessage is a paid mutator transaction binding the contract method 0x17c9c577.
//
// Solidity: function addInitMessage() returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) AddInitMessage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "addInitMessage")
}

// AddInitMessage is a paid mutator transaction binding the contract method 0x17c9c577.
//
// Solidity: function addInitMessage() returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddInitMessage() (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddInitMessage(&_SequencerInboxStub.TransactOpts)
}

// AddInitMessage is a paid mutator transaction binding the contract method 0x17c9c577.
//
// Solidity: function addInitMessage() returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) AddInitMessage() (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddInitMessage(&_SequencerInboxStub.TransactOpts)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x61a93f8f.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) AddSequencerL2Batch(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "addSequencerL2Batch", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x61a93f8f.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddSequencerL2Batch(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2Batch(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x61a93f8f.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) AddSequencerL2Batch(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2Batch(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) AddSequencerL2BatchFromOrigin(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "addSequencerL2BatchFromOrigin", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddSequencerL2BatchFromOrigin(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromOrigin(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) AddSequencerL2BatchFromOrigin(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromOrigin(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) ForceInclusion(opts *bind.TransactOpts, _totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "forceInclusion", _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.ForceInclusion(&_SequencerInboxStub.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.ForceInclusion(&_SequencerInboxStub.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f7a92b2.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) Initialize(opts *bind.TransactOpts, bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "initialize", bridge_, maxTimeVariation_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f7a92b2.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) Initialize(bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.Initialize(&_SequencerInboxStub.TransactOpts, bridge_, maxTimeVariation_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f7a92b2.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) Initialize(bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.Initialize(&_SequencerInboxStub.TransactOpts, bridge_, maxTimeVariation_)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) InvalidateKeysetHash(opts *bind.TransactOpts, ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "invalidateKeysetHash", ksHash)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) InvalidateKeysetHash(ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.InvalidateKeysetHash(&_SequencerInboxStub.TransactOpts, ksHash)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) InvalidateKeysetHash(ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.InvalidateKeysetHash(&_SequencerInboxStub.TransactOpts, ksHash)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) SetIsBatchPoster(opts *bind.TransactOpts, addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "setIsBatchPoster", addr, isBatchPoster_)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) SetIsBatchPoster(addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetIsBatchPoster(&_SequencerInboxStub.TransactOpts, addr, isBatchPoster_)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) SetIsBatchPoster(addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetIsBatchPoster(&_SequencerInboxStub.TransactOpts, addr, isBatchPoster_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) SetMaxTimeVariation(opts *bind.TransactOpts, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "setMaxTimeVariation", maxTimeVariation_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) SetMaxTimeVariation(maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetMaxTimeVariation(&_SequencerInboxStub.TransactOpts, maxTimeVariation_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) SetMaxTimeVariation(maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetMaxTimeVariation(&_SequencerInboxStub.TransactOpts, maxTimeVariation_)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) SetValidKeyset(opts *bind.TransactOpts, keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "setValidKeyset", keysetBytes)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) SetValidKeyset(keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetValidKeyset(&_SequencerInboxStub.TransactOpts, keysetBytes)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) SetValidKeyset(keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetValidKeyset(&_SequencerInboxStub.TransactOpts, keysetBytes)
}

// SequencerInboxStubInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the SequencerInboxStub contract.
type SequencerInboxStubInboxMessageDeliveredIterator struct {
	Event *SequencerInboxStubInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubInboxMessageDelivered)
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
		it.Event = new(SequencerInboxStubInboxMessageDelivered)
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
func (it *SequencerInboxStubInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubInboxMessageDelivered represents a InboxMessageDelivered event raised by the SequencerInboxStub contract.
type SequencerInboxStubInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*SequencerInboxStubInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubInboxMessageDeliveredIterator{contract: _SequencerInboxStub.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubInboxMessageDelivered)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseInboxMessageDelivered(log types.Log) (*SequencerInboxStubInboxMessageDelivered, error) {
	event := new(SequencerInboxStubInboxMessageDelivered)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the SequencerInboxStub contract.
type SequencerInboxStubInboxMessageDeliveredFromOriginIterator struct {
	Event *SequencerInboxStubInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubInboxMessageDeliveredFromOrigin)
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
		it.Event = new(SequencerInboxStubInboxMessageDeliveredFromOrigin)
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
func (it *SequencerInboxStubInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the SequencerInboxStub contract.
type SequencerInboxStubInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*SequencerInboxStubInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubInboxMessageDeliveredFromOriginIterator{contract: _SequencerInboxStub.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubInboxMessageDeliveredFromOrigin)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*SequencerInboxStubInboxMessageDeliveredFromOrigin, error) {
	event := new(SequencerInboxStubInboxMessageDeliveredFromOrigin)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubInvalidateKeysetIterator is returned from FilterInvalidateKeyset and is used to iterate over the raw logs and unpacked data for InvalidateKeyset events raised by the SequencerInboxStub contract.
type SequencerInboxStubInvalidateKeysetIterator struct {
	Event *SequencerInboxStubInvalidateKeyset // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubInvalidateKeysetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubInvalidateKeyset)
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
		it.Event = new(SequencerInboxStubInvalidateKeyset)
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
func (it *SequencerInboxStubInvalidateKeysetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubInvalidateKeysetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubInvalidateKeyset represents a InvalidateKeyset event raised by the SequencerInboxStub contract.
type SequencerInboxStubInvalidateKeyset struct {
	KeysetHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvalidateKeyset is a free log retrieval operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterInvalidateKeyset(opts *bind.FilterOpts, keysetHash [][32]byte) (*SequencerInboxStubInvalidateKeysetIterator, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "InvalidateKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubInvalidateKeysetIterator{contract: _SequencerInboxStub.contract, event: "InvalidateKeyset", logs: logs, sub: sub}, nil
}

// WatchInvalidateKeyset is a free log subscription operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchInvalidateKeyset(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubInvalidateKeyset, keysetHash [][32]byte) (event.Subscription, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "InvalidateKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubInvalidateKeyset)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "InvalidateKeyset", log); err != nil {
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

// ParseInvalidateKeyset is a log parse operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseInvalidateKeyset(log types.Log) (*SequencerInboxStubInvalidateKeyset, error) {
	event := new(SequencerInboxStubInvalidateKeyset)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "InvalidateKeyset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubOwnerFunctionCalledIterator is returned from FilterOwnerFunctionCalled and is used to iterate over the raw logs and unpacked data for OwnerFunctionCalled events raised by the SequencerInboxStub contract.
type SequencerInboxStubOwnerFunctionCalledIterator struct {
	Event *SequencerInboxStubOwnerFunctionCalled // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubOwnerFunctionCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubOwnerFunctionCalled)
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
		it.Event = new(SequencerInboxStubOwnerFunctionCalled)
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
func (it *SequencerInboxStubOwnerFunctionCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubOwnerFunctionCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubOwnerFunctionCalled represents a OwnerFunctionCalled event raised by the SequencerInboxStub contract.
type SequencerInboxStubOwnerFunctionCalled struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOwnerFunctionCalled is a free log retrieval operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterOwnerFunctionCalled(opts *bind.FilterOpts, id []*big.Int) (*SequencerInboxStubOwnerFunctionCalledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "OwnerFunctionCalled", idRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubOwnerFunctionCalledIterator{contract: _SequencerInboxStub.contract, event: "OwnerFunctionCalled", logs: logs, sub: sub}, nil
}

// WatchOwnerFunctionCalled is a free log subscription operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchOwnerFunctionCalled(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubOwnerFunctionCalled, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "OwnerFunctionCalled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubOwnerFunctionCalled)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
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

// ParseOwnerFunctionCalled is a log parse operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseOwnerFunctionCalled(log types.Log) (*SequencerInboxStubOwnerFunctionCalled, error) {
	event := new(SequencerInboxStubOwnerFunctionCalled)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubSequencerBatchDataIterator is returned from FilterSequencerBatchData and is used to iterate over the raw logs and unpacked data for SequencerBatchData events raised by the SequencerInboxStub contract.
type SequencerInboxStubSequencerBatchDataIterator struct {
	Event *SequencerInboxStubSequencerBatchData // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubSequencerBatchDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubSequencerBatchData)
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
		it.Event = new(SequencerInboxStubSequencerBatchData)
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
func (it *SequencerInboxStubSequencerBatchDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubSequencerBatchDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubSequencerBatchData represents a SequencerBatchData event raised by the SequencerInboxStub contract.
type SequencerInboxStubSequencerBatchData struct {
	BatchSequenceNumber *big.Int
	Data                []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchData is a free log retrieval operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterSequencerBatchData(opts *bind.FilterOpts, batchSequenceNumber []*big.Int) (*SequencerInboxStubSequencerBatchDataIterator, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "SequencerBatchData", batchSequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubSequencerBatchDataIterator{contract: _SequencerInboxStub.contract, event: "SequencerBatchData", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchData is a free log subscription operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchSequencerBatchData(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubSequencerBatchData, batchSequenceNumber []*big.Int) (event.Subscription, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "SequencerBatchData", batchSequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubSequencerBatchData)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "SequencerBatchData", log); err != nil {
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

// ParseSequencerBatchData is a log parse operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseSequencerBatchData(log types.Log) (*SequencerInboxStubSequencerBatchData, error) {
	event := new(SequencerInboxStubSequencerBatchData)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "SequencerBatchData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubSequencerBatchDeliveredIterator is returned from FilterSequencerBatchDelivered and is used to iterate over the raw logs and unpacked data for SequencerBatchDelivered events raised by the SequencerInboxStub contract.
type SequencerInboxStubSequencerBatchDeliveredIterator struct {
	Event *SequencerInboxStubSequencerBatchDelivered // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubSequencerBatchDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubSequencerBatchDelivered)
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
		it.Event = new(SequencerInboxStubSequencerBatchDelivered)
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
func (it *SequencerInboxStubSequencerBatchDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubSequencerBatchDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubSequencerBatchDelivered represents a SequencerBatchDelivered event raised by the SequencerInboxStub contract.
type SequencerInboxStubSequencerBatchDelivered struct {
	BatchSequenceNumber      *big.Int
	BeforeAcc                [32]byte
	AfterAcc                 [32]byte
	DelayedAcc               [32]byte
	AfterDelayedMessagesRead *big.Int
	TimeBounds               ISequencerInboxTimeBounds
	DataLocation             uint8
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchDelivered is a free log retrieval operation binding the contract event 0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed batchSequenceNumber, bytes32 indexed beforeAcc, bytes32 indexed afterAcc, bytes32 delayedAcc, uint256 afterDelayedMessagesRead, (uint64,uint64,uint64,uint64) timeBounds, uint8 dataLocation)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterSequencerBatchDelivered(opts *bind.FilterOpts, batchSequenceNumber []*big.Int, beforeAcc [][32]byte, afterAcc [][32]byte) (*SequencerInboxStubSequencerBatchDeliveredIterator, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}
	var afterAccRule []interface{}
	for _, afterAccItem := range afterAcc {
		afterAccRule = append(afterAccRule, afterAccItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "SequencerBatchDelivered", batchSequenceNumberRule, beforeAccRule, afterAccRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubSequencerBatchDeliveredIterator{contract: _SequencerInboxStub.contract, event: "SequencerBatchDelivered", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchDelivered is a free log subscription operation binding the contract event 0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed batchSequenceNumber, bytes32 indexed beforeAcc, bytes32 indexed afterAcc, bytes32 delayedAcc, uint256 afterDelayedMessagesRead, (uint64,uint64,uint64,uint64) timeBounds, uint8 dataLocation)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchSequencerBatchDelivered(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubSequencerBatchDelivered, batchSequenceNumber []*big.Int, beforeAcc [][32]byte, afterAcc [][32]byte) (event.Subscription, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}
	var afterAccRule []interface{}
	for _, afterAccItem := range afterAcc {
		afterAccRule = append(afterAccRule, afterAccItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "SequencerBatchDelivered", batchSequenceNumberRule, beforeAccRule, afterAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubSequencerBatchDelivered)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
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

// ParseSequencerBatchDelivered is a log parse operation binding the contract event 0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed batchSequenceNumber, bytes32 indexed beforeAcc, bytes32 indexed afterAcc, bytes32 delayedAcc, uint256 afterDelayedMessagesRead, (uint64,uint64,uint64,uint64) timeBounds, uint8 dataLocation)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseSequencerBatchDelivered(log types.Log) (*SequencerInboxStubSequencerBatchDelivered, error) {
	event := new(SequencerInboxStubSequencerBatchDelivered)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubSetValidKeysetIterator is returned from FilterSetValidKeyset and is used to iterate over the raw logs and unpacked data for SetValidKeyset events raised by the SequencerInboxStub contract.
type SequencerInboxStubSetValidKeysetIterator struct {
	Event *SequencerInboxStubSetValidKeyset // Event containing the contract specifics and raw log

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
func (it *SequencerInboxStubSetValidKeysetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubSetValidKeyset)
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
		it.Event = new(SequencerInboxStubSetValidKeyset)
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
func (it *SequencerInboxStubSetValidKeysetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubSetValidKeysetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubSetValidKeyset represents a SetValidKeyset event raised by the SequencerInboxStub contract.
type SequencerInboxStubSetValidKeyset struct {
	KeysetHash  [32]byte
	KeysetBytes []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetValidKeyset is a free log retrieval operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterSetValidKeyset(opts *bind.FilterOpts, keysetHash [][32]byte) (*SequencerInboxStubSetValidKeysetIterator, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "SetValidKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubSetValidKeysetIterator{contract: _SequencerInboxStub.contract, event: "SetValidKeyset", logs: logs, sub: sub}, nil
}

// WatchSetValidKeyset is a free log subscription operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchSetValidKeyset(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubSetValidKeyset, keysetHash [][32]byte) (event.Subscription, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "SetValidKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubSetValidKeyset)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "SetValidKeyset", log); err != nil {
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

// ParseSetValidKeyset is a log parse operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseSetValidKeyset(log types.Log) (*SequencerInboxStubSetValidKeyset, error) {
	event := new(SequencerInboxStubSetValidKeyset)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "SetValidKeyset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleMetaData contains all meta data concerning the Simple contract.
var SimpleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"CounterEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NullEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"}],\"name\":\"RedeemedEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"checkBlockHashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitNullEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementEmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementRedeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506103ba806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630e8c389f146100675780635677c11e1461007157806361bc221a1461008c5780639ff5ccac146100b7578063b226a964146100bf578063d09de08a146100c7575b600080fd5b61006f6100cf565b005b6100796101c6565b6040519081526020015b60405180910390f35b60005461009f906001600160401b031681565b6040516001600160401b039091168152602001610083565b61006f610229565b61006f610296565b61006f6102c1565b600080546001600160401b031690806100e783610316565b91906101000a8154816001600160401b0302191690836001600160401b03160217905550507f773c78bf96e65f61c1a2622b47d76e78bfe70dd59cf4f11470c4c121c315941333606e6001600160a01b031663de4ba2b36040518163ffffffff1660e01b815260040160206040518083038186803b15801561016857600080fd5b505afa15801561017c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101a0919061033d565b604080516001600160a01b039384168152929091166020830152015b60405180910390a1565b60006101d360024361036d565b406101df60014361036d565b4014156102245760405162461bcd60e51b815260206004820152600f60248201526e0a6829a8abe84989e8696be9082a69608b1b604482015260640160405180910390fd5b504390565b600080546001600160401b0316908061024183610316565b82546101009290920a6001600160401b03818102199093169183160217909155600054604051911681527fa45d7e79cb3c6044f30c8dd891e6571301d6b8b6618df519c987905ec70742e791506020016101bc565b6040517f6f59c82101949290205a9ae9d0c657e6dae1a71c301ae76d385c2792294585fe90600090a1565b600080546001600160401b031690806102d983610316565b91906101000a8154816001600160401b0302191690836001600160401b0316021790555050565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b038083168181141561033357610333610300565b6001019392505050565b60006020828403121561034f57600080fd5b81516001600160a01b038116811461036657600080fd5b9392505050565b60008282101561037f5761037f610300565b50039056fea264697066735822122083a0a53edc158d4dfdf7a4a5009bf860cb7cdde4597b10f3477a518dce43936b64736f6c63430008090033",
}

// SimpleABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleMetaData.ABI instead.
var SimpleABI = SimpleMetaData.ABI

// SimpleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleMetaData.Bin instead.
var SimpleBin = SimpleMetaData.Bin

// DeploySimple deploys a new Ethereum contract, binding an instance of Simple to it.
func DeploySimple(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Simple, error) {
	parsed, err := SimpleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Simple{SimpleCaller: SimpleCaller{contract: contract}, SimpleTransactor: SimpleTransactor{contract: contract}, SimpleFilterer: SimpleFilterer{contract: contract}}, nil
}

// Simple is an auto generated Go binding around an Ethereum contract.
type Simple struct {
	SimpleCaller     // Read-only binding to the contract
	SimpleTransactor // Write-only binding to the contract
	SimpleFilterer   // Log filterer for contract events
}

// SimpleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleSession struct {
	Contract     *Simple           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleCallerSession struct {
	Contract *SimpleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SimpleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleTransactorSession struct {
	Contract     *SimpleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleRaw struct {
	Contract *Simple // Generic contract binding to access the raw methods on
}

// SimpleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleCallerRaw struct {
	Contract *SimpleCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleTransactorRaw struct {
	Contract *SimpleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimple creates a new instance of Simple, bound to a specific deployed contract.
func NewSimple(address common.Address, backend bind.ContractBackend) (*Simple, error) {
	contract, err := bindSimple(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simple{SimpleCaller: SimpleCaller{contract: contract}, SimpleTransactor: SimpleTransactor{contract: contract}, SimpleFilterer: SimpleFilterer{contract: contract}}, nil
}

// NewSimpleCaller creates a new read-only instance of Simple, bound to a specific deployed contract.
func NewSimpleCaller(address common.Address, caller bind.ContractCaller) (*SimpleCaller, error) {
	contract, err := bindSimple(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleCaller{contract: contract}, nil
}

// NewSimpleTransactor creates a new write-only instance of Simple, bound to a specific deployed contract.
func NewSimpleTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleTransactor, error) {
	contract, err := bindSimple(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTransactor{contract: contract}, nil
}

// NewSimpleFilterer creates a new log filterer instance of Simple, bound to a specific deployed contract.
func NewSimpleFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleFilterer, error) {
	contract, err := bindSimple(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleFilterer{contract: contract}, nil
}

// bindSimple binds a generic wrapper to an already deployed contract.
func bindSimple(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simple *SimpleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simple.Contract.SimpleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simple *SimpleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.Contract.SimpleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simple *SimpleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simple.Contract.SimpleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simple *SimpleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simple.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simple *SimpleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simple *SimpleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simple.Contract.contract.Transact(opts, method, params...)
}

// CheckBlockHashes is a free data retrieval call binding the contract method 0x5677c11e.
//
// Solidity: function checkBlockHashes() view returns(uint256)
func (_Simple *SimpleCaller) CheckBlockHashes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Simple.contract.Call(opts, &out, "checkBlockHashes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckBlockHashes is a free data retrieval call binding the contract method 0x5677c11e.
//
// Solidity: function checkBlockHashes() view returns(uint256)
func (_Simple *SimpleSession) CheckBlockHashes() (*big.Int, error) {
	return _Simple.Contract.CheckBlockHashes(&_Simple.CallOpts)
}

// CheckBlockHashes is a free data retrieval call binding the contract method 0x5677c11e.
//
// Solidity: function checkBlockHashes() view returns(uint256)
func (_Simple *SimpleCallerSession) CheckBlockHashes() (*big.Int, error) {
	return _Simple.Contract.CheckBlockHashes(&_Simple.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint64)
func (_Simple *SimpleCaller) Counter(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Simple.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint64)
func (_Simple *SimpleSession) Counter() (uint64, error) {
	return _Simple.Contract.Counter(&_Simple.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint64)
func (_Simple *SimpleCallerSession) Counter() (uint64, error) {
	return _Simple.Contract.Counter(&_Simple.CallOpts)
}

// EmitNullEvent is a paid mutator transaction binding the contract method 0xb226a964.
//
// Solidity: function emitNullEvent() returns()
func (_Simple *SimpleTransactor) EmitNullEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "emitNullEvent")
}

// EmitNullEvent is a paid mutator transaction binding the contract method 0xb226a964.
//
// Solidity: function emitNullEvent() returns()
func (_Simple *SimpleSession) EmitNullEvent() (*types.Transaction, error) {
	return _Simple.Contract.EmitNullEvent(&_Simple.TransactOpts)
}

// EmitNullEvent is a paid mutator transaction binding the contract method 0xb226a964.
//
// Solidity: function emitNullEvent() returns()
func (_Simple *SimpleTransactorSession) EmitNullEvent() (*types.Transaction, error) {
	return _Simple.Contract.EmitNullEvent(&_Simple.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Simple *SimpleTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Simple *SimpleSession) Increment() (*types.Transaction, error) {
	return _Simple.Contract.Increment(&_Simple.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Simple *SimpleTransactorSession) Increment() (*types.Transaction, error) {
	return _Simple.Contract.Increment(&_Simple.TransactOpts)
}

// IncrementEmit is a paid mutator transaction binding the contract method 0x9ff5ccac.
//
// Solidity: function incrementEmit() returns()
func (_Simple *SimpleTransactor) IncrementEmit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "incrementEmit")
}

// IncrementEmit is a paid mutator transaction binding the contract method 0x9ff5ccac.
//
// Solidity: function incrementEmit() returns()
func (_Simple *SimpleSession) IncrementEmit() (*types.Transaction, error) {
	return _Simple.Contract.IncrementEmit(&_Simple.TransactOpts)
}

// IncrementEmit is a paid mutator transaction binding the contract method 0x9ff5ccac.
//
// Solidity: function incrementEmit() returns()
func (_Simple *SimpleTransactorSession) IncrementEmit() (*types.Transaction, error) {
	return _Simple.Contract.IncrementEmit(&_Simple.TransactOpts)
}

// IncrementRedeem is a paid mutator transaction binding the contract method 0x0e8c389f.
//
// Solidity: function incrementRedeem() returns()
func (_Simple *SimpleTransactor) IncrementRedeem(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "incrementRedeem")
}

// IncrementRedeem is a paid mutator transaction binding the contract method 0x0e8c389f.
//
// Solidity: function incrementRedeem() returns()
func (_Simple *SimpleSession) IncrementRedeem() (*types.Transaction, error) {
	return _Simple.Contract.IncrementRedeem(&_Simple.TransactOpts)
}

// IncrementRedeem is a paid mutator transaction binding the contract method 0x0e8c389f.
//
// Solidity: function incrementRedeem() returns()
func (_Simple *SimpleTransactorSession) IncrementRedeem() (*types.Transaction, error) {
	return _Simple.Contract.IncrementRedeem(&_Simple.TransactOpts)
}

// SimpleCounterEventIterator is returned from FilterCounterEvent and is used to iterate over the raw logs and unpacked data for CounterEvent events raised by the Simple contract.
type SimpleCounterEventIterator struct {
	Event *SimpleCounterEvent // Event containing the contract specifics and raw log

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
func (it *SimpleCounterEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleCounterEvent)
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
		it.Event = new(SimpleCounterEvent)
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
func (it *SimpleCounterEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleCounterEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleCounterEvent represents a CounterEvent event raised by the Simple contract.
type SimpleCounterEvent struct {
	Count uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCounterEvent is a free log retrieval operation binding the contract event 0xa45d7e79cb3c6044f30c8dd891e6571301d6b8b6618df519c987905ec70742e7.
//
// Solidity: event CounterEvent(uint64 count)
func (_Simple *SimpleFilterer) FilterCounterEvent(opts *bind.FilterOpts) (*SimpleCounterEventIterator, error) {

	logs, sub, err := _Simple.contract.FilterLogs(opts, "CounterEvent")
	if err != nil {
		return nil, err
	}
	return &SimpleCounterEventIterator{contract: _Simple.contract, event: "CounterEvent", logs: logs, sub: sub}, nil
}

// WatchCounterEvent is a free log subscription operation binding the contract event 0xa45d7e79cb3c6044f30c8dd891e6571301d6b8b6618df519c987905ec70742e7.
//
// Solidity: event CounterEvent(uint64 count)
func (_Simple *SimpleFilterer) WatchCounterEvent(opts *bind.WatchOpts, sink chan<- *SimpleCounterEvent) (event.Subscription, error) {

	logs, sub, err := _Simple.contract.WatchLogs(opts, "CounterEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleCounterEvent)
				if err := _Simple.contract.UnpackLog(event, "CounterEvent", log); err != nil {
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

// ParseCounterEvent is a log parse operation binding the contract event 0xa45d7e79cb3c6044f30c8dd891e6571301d6b8b6618df519c987905ec70742e7.
//
// Solidity: event CounterEvent(uint64 count)
func (_Simple *SimpleFilterer) ParseCounterEvent(log types.Log) (*SimpleCounterEvent, error) {
	event := new(SimpleCounterEvent)
	if err := _Simple.contract.UnpackLog(event, "CounterEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleNullEventIterator is returned from FilterNullEvent and is used to iterate over the raw logs and unpacked data for NullEvent events raised by the Simple contract.
type SimpleNullEventIterator struct {
	Event *SimpleNullEvent // Event containing the contract specifics and raw log

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
func (it *SimpleNullEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleNullEvent)
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
		it.Event = new(SimpleNullEvent)
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
func (it *SimpleNullEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleNullEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleNullEvent represents a NullEvent event raised by the Simple contract.
type SimpleNullEvent struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNullEvent is a free log retrieval operation binding the contract event 0x6f59c82101949290205a9ae9d0c657e6dae1a71c301ae76d385c2792294585fe.
//
// Solidity: event NullEvent()
func (_Simple *SimpleFilterer) FilterNullEvent(opts *bind.FilterOpts) (*SimpleNullEventIterator, error) {

	logs, sub, err := _Simple.contract.FilterLogs(opts, "NullEvent")
	if err != nil {
		return nil, err
	}
	return &SimpleNullEventIterator{contract: _Simple.contract, event: "NullEvent", logs: logs, sub: sub}, nil
}

// WatchNullEvent is a free log subscription operation binding the contract event 0x6f59c82101949290205a9ae9d0c657e6dae1a71c301ae76d385c2792294585fe.
//
// Solidity: event NullEvent()
func (_Simple *SimpleFilterer) WatchNullEvent(opts *bind.WatchOpts, sink chan<- *SimpleNullEvent) (event.Subscription, error) {

	logs, sub, err := _Simple.contract.WatchLogs(opts, "NullEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleNullEvent)
				if err := _Simple.contract.UnpackLog(event, "NullEvent", log); err != nil {
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

// ParseNullEvent is a log parse operation binding the contract event 0x6f59c82101949290205a9ae9d0c657e6dae1a71c301ae76d385c2792294585fe.
//
// Solidity: event NullEvent()
func (_Simple *SimpleFilterer) ParseNullEvent(log types.Log) (*SimpleNullEvent, error) {
	event := new(SimpleNullEvent)
	if err := _Simple.contract.UnpackLog(event, "NullEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRedeemedEventIterator is returned from FilterRedeemedEvent and is used to iterate over the raw logs and unpacked data for RedeemedEvent events raised by the Simple contract.
type SimpleRedeemedEventIterator struct {
	Event *SimpleRedeemedEvent // Event containing the contract specifics and raw log

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
func (it *SimpleRedeemedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRedeemedEvent)
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
		it.Event = new(SimpleRedeemedEvent)
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
func (it *SimpleRedeemedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRedeemedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRedeemedEvent represents a RedeemedEvent event raised by the Simple contract.
type SimpleRedeemedEvent struct {
	Caller   common.Address
	Redeemer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedeemedEvent is a free log retrieval operation binding the contract event 0x773c78bf96e65f61c1a2622b47d76e78bfe70dd59cf4f11470c4c121c3159413.
//
// Solidity: event RedeemedEvent(address caller, address redeemer)
func (_Simple *SimpleFilterer) FilterRedeemedEvent(opts *bind.FilterOpts) (*SimpleRedeemedEventIterator, error) {

	logs, sub, err := _Simple.contract.FilterLogs(opts, "RedeemedEvent")
	if err != nil {
		return nil, err
	}
	return &SimpleRedeemedEventIterator{contract: _Simple.contract, event: "RedeemedEvent", logs: logs, sub: sub}, nil
}

// WatchRedeemedEvent is a free log subscription operation binding the contract event 0x773c78bf96e65f61c1a2622b47d76e78bfe70dd59cf4f11470c4c121c3159413.
//
// Solidity: event RedeemedEvent(address caller, address redeemer)
func (_Simple *SimpleFilterer) WatchRedeemedEvent(opts *bind.WatchOpts, sink chan<- *SimpleRedeemedEvent) (event.Subscription, error) {

	logs, sub, err := _Simple.contract.WatchLogs(opts, "RedeemedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRedeemedEvent)
				if err := _Simple.contract.UnpackLog(event, "RedeemedEvent", log); err != nil {
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

// ParseRedeemedEvent is a log parse operation binding the contract event 0x773c78bf96e65f61c1a2622b47d76e78bfe70dd59cf4f11470c4c121c3159413.
//
// Solidity: event RedeemedEvent(address caller, address redeemer)
func (_Simple *SimpleFilterer) ParseRedeemedEvent(log types.Log) (*SimpleRedeemedEvent, error) {
	event := new(SimpleRedeemedEvent)
	if err := _Simple.contract.UnpackLog(event, "RedeemedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleProxyMetaData contains all meta data concerning the SimpleProxy contract.
var SimpleProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161011d38038061011d83398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051609561008860003960006017015260956000f3fe608060405236601057600e6013565b005b600e5b603a7f0000000000000000000000000000000000000000000000000000000000000000603c565b565b3660008037600080366000845af43d6000803e808015605a573d6000f35b3d6000fdfea2646970667358221220b9512068685fa7c7af2de743123c43df93094268afa6ab531deaff05f774b5c264736f6c63430008090033",
}

// SimpleProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleProxyMetaData.ABI instead.
var SimpleProxyABI = SimpleProxyMetaData.ABI

// SimpleProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleProxyMetaData.Bin instead.
var SimpleProxyBin = SimpleProxyMetaData.Bin

// DeploySimpleProxy deploys a new Ethereum contract, binding an instance of SimpleProxy to it.
func DeploySimpleProxy(auth *bind.TransactOpts, backend bind.ContractBackend, impl_ common.Address) (common.Address, *types.Transaction, *SimpleProxy, error) {
	parsed, err := SimpleProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleProxyBin), backend, impl_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleProxy{SimpleProxyCaller: SimpleProxyCaller{contract: contract}, SimpleProxyTransactor: SimpleProxyTransactor{contract: contract}, SimpleProxyFilterer: SimpleProxyFilterer{contract: contract}}, nil
}

// SimpleProxy is an auto generated Go binding around an Ethereum contract.
type SimpleProxy struct {
	SimpleProxyCaller     // Read-only binding to the contract
	SimpleProxyTransactor // Write-only binding to the contract
	SimpleProxyFilterer   // Log filterer for contract events
}

// SimpleProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleProxySession struct {
	Contract     *SimpleProxy      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleProxyCallerSession struct {
	Contract *SimpleProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SimpleProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleProxyTransactorSession struct {
	Contract     *SimpleProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SimpleProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleProxyRaw struct {
	Contract *SimpleProxy // Generic contract binding to access the raw methods on
}

// SimpleProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleProxyCallerRaw struct {
	Contract *SimpleProxyCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleProxyTransactorRaw struct {
	Contract *SimpleProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleProxy creates a new instance of SimpleProxy, bound to a specific deployed contract.
func NewSimpleProxy(address common.Address, backend bind.ContractBackend) (*SimpleProxy, error) {
	contract, err := bindSimpleProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleProxy{SimpleProxyCaller: SimpleProxyCaller{contract: contract}, SimpleProxyTransactor: SimpleProxyTransactor{contract: contract}, SimpleProxyFilterer: SimpleProxyFilterer{contract: contract}}, nil
}

// NewSimpleProxyCaller creates a new read-only instance of SimpleProxy, bound to a specific deployed contract.
func NewSimpleProxyCaller(address common.Address, caller bind.ContractCaller) (*SimpleProxyCaller, error) {
	contract, err := bindSimpleProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleProxyCaller{contract: contract}, nil
}

// NewSimpleProxyTransactor creates a new write-only instance of SimpleProxy, bound to a specific deployed contract.
func NewSimpleProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleProxyTransactor, error) {
	contract, err := bindSimpleProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleProxyTransactor{contract: contract}, nil
}

// NewSimpleProxyFilterer creates a new log filterer instance of SimpleProxy, bound to a specific deployed contract.
func NewSimpleProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleProxyFilterer, error) {
	contract, err := bindSimpleProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleProxyFilterer{contract: contract}, nil
}

// bindSimpleProxy binds a generic wrapper to an already deployed contract.
func bindSimpleProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleProxy *SimpleProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleProxy.Contract.SimpleProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleProxy *SimpleProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleProxy.Contract.SimpleProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleProxy *SimpleProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleProxy.Contract.SimpleProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleProxy *SimpleProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleProxy *SimpleProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleProxy *SimpleProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleProxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleProxy *SimpleProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SimpleProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleProxy *SimpleProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SimpleProxy.Contract.Fallback(&_SimpleProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleProxy *SimpleProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SimpleProxy.Contract.Fallback(&_SimpleProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleProxy *SimpleProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleProxy *SimpleProxySession) Receive() (*types.Transaction, error) {
	return _SimpleProxy.Contract.Receive(&_SimpleProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleProxy *SimpleProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _SimpleProxy.Contract.Receive(&_SimpleProxy.TransactOpts)
}

// SingleExecutionChallengeMetaData contains all meta data concerning the SingleExecutionChallenge contract.
var SingleExecutionChallengeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"osp_\",\"type\":\"address\"},{\"internalType\":\"contractIChallengeResultReceiver\",\"name\":\"resultReceiver_\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessagesRead_\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[2]\",\"name\":\"startAndEndHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256\",\"name\":\"numSteps_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asserter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"asserterTimeLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengerTimeLeft_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"challengeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"chainHashes\",\"type\":\"bytes32[]\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIChallengeManager.ChallengeTerminationType\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"ChallengeEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockSteps\",\"type\":\"uint256\"}],\"name\":\"ExecutionChallengeBegun\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"startState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"endState\",\"type\":\"tuple\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"newSegments\",\"type\":\"bytes32[]\"}],\"name\":\"bisectExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"machineStatuses\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"globalStateHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256\",\"name\":\"numSteps\",\"type\":\"uint256\"}],\"name\":\"challengeExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"challengeInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"internalType\":\"structChallengeLib.Challenge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"challenges\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"clearChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"startAndEndMachineStatuses_\",\"type\":\"uint8[2]\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState[2]\",\"name\":\"startAndEndGlobalStates_\",\"type\":\"tuple[2]\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"asserter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"asserterTimeLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengerTimeLeft_\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"currentResponder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"currentResponderTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChallengeResultReceiver\",\"name\":\"resultReceiver_\",\"type\":\"address\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"sequencerInbox_\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"osp_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"isTimedOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProveExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"osp\",\"outputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resultReceiver\",\"outputs\":[{\"internalType\":\"contractIChallengeResultReceiver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalChallengesCreated\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052306080523480156200001557600080fd5b5060405162003539380380620035398339810160408190526200003891620002bc565b600580546001600160a01b03808c166001600160a01b03199283161790925560028054928b1692909116919091179055600080548190819062000084906001600160401b0316620003cd565b82546101009290920a6001600160401b03818102199093168284169182021790935560009283526001602090815260408085206007810180546001600160401b031916958f16959095179094558051600280825260608201835293965093949392918301908036833750508a518251929350918391506000906200010c576200010c62000403565b602090810291909101015288600160200201518160018151811062000135576200013562000403565b60200260200101818152505060006200015c60008a846200024360201b62001a411760201c565b600684018190556040805180820182526001600160a01b038b811680835260209283018b90526002880180546001600160a01b03199081169092179055600388018b905583518085018552918c168083529190920189905286549091161785556001850187905542600486015560078501805460ff60401b1916680200000000000000001790555190915081906001600160401b038616907f86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d689062000228906000908e90889062000419565b60405180910390a350505050505050505050505050620004b4565b60008383836040516020016200025c9392919062000470565b6040516020818303038152906040528051906020012090509392505050565b6001600160a01b03811681146200029157600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b8051620002b7816200027b565b919050565b60008060008060008060008060006101408a8c031215620002dc57600080fd5b8951620002e9816200027b565b809950506020808b0151620002fe816200027b565b60408c01519099506001600160401b0380821682146200031d57600080fd5b8199508d607f8e01126200033057600080fd5b604051915060408201828110828211171562000350576200035062000294565b604052508060a08d018e8111156200036757600080fd5b60608e015b818110156200038557805183529184019184016200036c565b50519198509096506200039e91505060c08b01620002aa565b9350620003ae60e08b01620002aa565b92506101008a015191506101208a015190509295985092959850929598565b60006001600160401b0382811680821415620003f957634e487b7160e01b600052601160045260246000fd5b6001019392505050565b634e487b7160e01b600052603260045260246000fd5b6000606082018583526020858185015260606040850152818551808452608086019150828701935060005b81811015620004625784518352938301939183019160010162000444565b509098975050505050505050565b83815260006020848184015260408301845182860160005b82811015620004a65781518452928401929084019060010162000488565b509198975050505050505050565b608051613069620004d0600039600061126b01526130696000f3fe608060405234801561001057600080fd5b50600436106100eb5760003560e01c80639ede42b9116100925780639ede42b91461025c578063a521b0321461027f578063c8fdb65b14610292578063d248d124146102cd578063e78cea92146102e0578063ee35f327146102f3578063f26a62c614610306578063f8c8765e14610319578063fb7be0a11461032c57600080fd5b806314eab5e7146100f05780631b45c86a1461012057806323a9ef23146101355780633504f1d71461016057806356e9df97146101735780635ef489e6146101865780637fd07a9c146101995780638f1d3776146101b9575b600080fd5b6101036100fe36600461263e565b61033f565b6040516001600160401b0390911681526020015b60405180910390f35b61013361012e3660046126d1565b610647565b005b6101486101433660046126d1565b610717565b6040516001600160a01b039091168152602001610117565b600254610148906001600160a01b031681565b6101336101813660046126d1565b61073b565b600054610103906001600160401b031681565b6101ac6101a73660046126d1565b6108a9565b604051610117919061272e565b6102496101c73660046127a0565b6001602081815260009283526040928390208351808501855281546001600160a01b0390811682529382015481840152845180860190955260028201549093168452600381015491840191909152600481015460058201546006830154600790930154939493919290916001600160401b03811690600160401b900460ff1687565b60405161011797969594939291906127b9565b61026f61026a3660046126d1565b610982565b6040519015158152602001610117565b61013361028d366004612816565b6109a9565b6102bf6102a03660046126d1565b6001600160401b03166000908152600160208190526040909120015490565b604051908152602001610117565b6101336102db3660046128ba565b610e1f565b600454610148906001600160a01b031681565b600354610148906001600160a01b031681565b600554610148906001600160a01b031681565b61013361032736600461294c565b611260565b61013361033a3660046129a8565b6113d1565b6002546000906001600160a01b031633146103945760405162461bcd60e51b815260206004820152601060248201526f13d3931657d493d313155417d0d2105360821b60448201526064015b60405180910390fd5b6040805160028082526060820183526000926020830190803683370190505090506103ea6103c560208b018b612a4c565b6103e58a60005b608002018036038101906103e09190612b0b565b611a78565b611af9565b816000815181106103fd576103fd612a36565b602090810291909101015261042c8960016020020160208101906104219190612a4c565b6103e58a60016103cc565b8160018151811061043f5761043f612a36565b60209081029190910101526000805481908190610464906001600160401b0316612bb9565b91906101000a8154816001600160401b0302191690836001600160401b031602179055905060006001600160401b0316816001600160401b031614156104ac576104ac612be0565b6001600160401b0381166000908152600160205260408120600581018d9055906104e66104e1368d90038d0160808e01612b0b565b611c1d565b905060026104fa60408e0160208f01612a4c565b600381111561050b5761050b612704565b14806105395750600061052e610529368e90038e0160808f01612b0b565b611c32565b6001600160401b0316115b1561054c578061054881612bb9565b9150505b6007820180546040805180820182526001600160a01b038d811680835260209283018d90526002880180546001600160a01b03199081169092179055600388018d905583518085018552918e16808352919092018b90528654909116178555600185018990554260048601556001600160401b0384811668ffffffffffffffffff1990931692909217600160401b179092559051908416907f76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a90610616908e906080820190612c40565b60405180910390a26106348360008c6001600160401b031687611c41565b5090925050505b98975050505050505050565b60006001600160401b038216600090815260016020526040902060070154600160401b900460ff16600281111561068057610680612704565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906106c05760405162461bcd60e51b815260040161038b9190612c5c565b506106ca81610982565b6107095760405162461bcd60e51b815260206004820152601060248201526f54494d454f55545f444541444c494e4560801b604482015260640161038b565b610714816000611cd7565b50565b6001600160401b03166000908152600160205260409020546001600160a01b031690565b6002546001600160a01b031633146107885760405162461bcd60e51b815260206004820152601060248201526f2727aa2fa922a9afa922a1a2a4ab22a960811b604482015260640161038b565b60006001600160401b038216600090815260016020526040902060070154600160401b900460ff1660028111156107c1576107c1612704565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906108015760405162461bcd60e51b815260040161038b9190612c5c565b506001600160401b038116600081815260016020819052604080832080546001600160a01b031990811682559281018490556002810180549093169092556003808301849055600483018490556005830184905560068301939093556007909101805468ffffffffffffffffff19169055517ffdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f409161089e91612cb1565b60405180910390a250565b6108b1612599565b6001600160401b0382811660009081526001602081815260409283902083516101208101855281546001600160a01b0390811660e0830190815294830154610100830152938152845180860186526002808401549095168152600383015481850152928101929092526004810154938201939093526005830154606082015260068301546080820152600783015493841660a08201529260c0840191600160401b90910460ff169081111561096857610968612704565b600281111561097957610979612704565b90525092915050565b6001600160401b03811660009081526001602052604081206109a390611e05565b92915050565b6001600160401b0384166000908152600160205260408120859185916109ce84610717565b6001600160a01b0316336001600160a01b0316146109fe5760405162461bcd60e51b815260040161038b90612ccb565b610a0784610982565b15610a245760405162461bcd60e51b815260040161038b90612cf0565b6000826002811115610a3857610a38612704565b1415610aa65760006007820154600160401b900460ff166002811115610a6057610a60612704565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b81525090610aa05760405162461bcd60e51b815260040161038b9190612c5c565b50610b65565b6001826002811115610aba57610aba612704565b1415610b045760016007820154600160401b900460ff166002811115610ae257610ae2612704565b14610aff5760405162461bcd60e51b815260040161038b90612d17565b610b65565b6002826002811115610b1857610b18612704565b1415610b5d5760026007820154600160401b900460ff166002811115610b4057610b40612704565b14610aff5760405162461bcd60e51b815260040161038b90612d3f565b610b65612be0565b610bb383356020850135610b7c6040870187612d6b565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250611a4192505050565b816006015414610bd55760405162461bcd60e51b815260040161038b90612dbb565b6002610be46040850185612d6b565b90501080610c0f57506001610bfc6040850185612d6b565b610c07929150612dde565b836060013510155b15610c2c5760405162461bcd60e51b815260040161038b90612df5565b600080610c3889611e1d565b9150915060018111610c785760405162461bcd60e51b81526020600482015260096024820152681513d3d7d4d213d49560ba1b604482015260640161038b565b806028811115610c86575060285b610c91816001612e20565b8814610cce5760405162461bcd60e51b815260206004820152600c60248201526b57524f4e475f44454752454560a01b604482015260640161038b565b50610d188989896000818110610ce657610ce6612a36565b602002919091013590508a8a610cfd600182612dde565b818110610d0c57610d0c612a36565b90506020020135611eae565b610d578a83838b8b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250611c4192505050565b50600090505b6007820154600160401b900460ff166002811115610d7d57610d7d612704565b1415610d895750610e16565b6040805180820190915281546001600160a01b03168152600182015460208201526004820154610db99042612dde565b81602001818151610dca9190612dde565b90525060028201805483546001600160a01b038083166001600160a01b031992831617865560038601805460018801558551929093169116179091556020909101519055426004909101555b50505050505050565b6001600160401b038416600090815260016020526040902084908490600290610e4784610717565b6001600160a01b0316336001600160a01b031614610e775760405162461bcd60e51b815260040161038b90612ccb565b610e8084610982565b15610e9d5760405162461bcd60e51b815260040161038b90612cf0565b6000826002811115610eb157610eb1612704565b1415610f1f5760006007820154600160401b900460ff166002811115610ed957610ed9612704565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b81525090610f195760405162461bcd60e51b815260040161038b9190612c5c565b50610fde565b6001826002811115610f3357610f33612704565b1415610f7d5760016007820154600160401b900460ff166002811115610f5b57610f5b612704565b14610f785760405162461bcd60e51b815260040161038b90612d17565b610fde565b6002826002811115610f9157610f91612704565b1415610fd65760026007820154600160401b900460ff166002811115610fb957610fb9612704565b14610f785760405162461bcd60e51b815260040161038b90612d3f565b610fde612be0565b610ff583356020850135610b7c6040870187612d6b565b8160060154146110175760405162461bcd60e51b815260040161038b90612dbb565b60026110266040850185612d6b565b905010806110515750600161103e6040850185612d6b565b611049929150612dde565b836060013510155b1561106e5760405162461bcd60e51b815260040161038b90612df5565b6001600160401b038816600090815260016020526040812090806110918a611e1d565b9092509050600181146110b65760405162461bcd60e51b815260040161038b90612e38565b5060055460408051808201825260078501546001600160401b031681526004546001600160a01b0390811660208301526000931691635d3adcfb91908590611100908f018f612d6b565b8f6060013581811061111457611114612a36565b905060200201358d8d6040518663ffffffff1660e01b815260040161113d959493929190612e5a565b60206040518083038186803b15801561115557600080fd5b505afa158015611169573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061118d9190612eb1565b905061119c60408b018b612d6b565b6111ab60608d01356001612e20565b8181106111ba576111ba612a36565b905060200201358114156111ff5760405162461bcd60e51b815260206004820152600c60248201526b14d0535157d3d4d417d1539160a21b604482015260640161038b565b6040516001600160401b038c16907fc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e90600090a26112548b6001600160401b0316600090815260016020526040812060060155565b5060009150610d5d9050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156112ee5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b606482015260840161038b565b6002546001600160a01b0316156113365760405162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b604482015260640161038b565b6001600160a01b0384166113815760405162461bcd60e51b81526020600482015260126024820152712727afa922a9aaa62a2fa922a1a2a4ab22a960711b604482015260640161038b565b600280546001600160a01b039586166001600160a01b0319918216179091556003805494861694821694909417909355600480549285169284169290921790915560058054919093169116179055565b6001600160401b0385166000908152600160208190526040909120869186916113f984610717565b6001600160a01b0316336001600160a01b0316146114295760405162461bcd60e51b815260040161038b90612ccb565b61143284610982565b1561144f5760405162461bcd60e51b815260040161038b90612cf0565b600082600281111561146357611463612704565b14156114d15760006007820154600160401b900460ff16600281111561148b5761148b612704565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906114cb5760405162461bcd60e51b815260040161038b9190612c5c565b50611590565b60018260028111156114e5576114e5612704565b141561152f5760016007820154600160401b900460ff16600281111561150d5761150d612704565b1461152a5760405162461bcd60e51b815260040161038b90612d17565b611590565b600282600281111561154357611543612704565b14156115885760026007820154600160401b900460ff16600281111561156b5761156b612704565b1461152a5760405162461bcd60e51b815260040161038b90612d3f565b611590612be0565b6115a783356020850135610b7c6040870187612d6b565b8160060154146115c95760405162461bcd60e51b815260040161038b90612dbb565b60026115d86040850185612d6b565b90501080611603575060016115f06040850185612d6b565b6115fb929150612dde565b836060013510155b156116205760405162461bcd60e51b815260040161038b90612df5565b60018510156116675760405162461bcd60e51b815260206004820152601360248201527210d2105313115391d157d513d3d7d4d213d495606a1b604482015260640161038b565b650800000000008511156116b25760405162461bcd60e51b81526020600482015260126024820152714348414c4c454e47455f544f4f5f4c4f4e4760701b604482015260640161038b565b6116f4886116d46116c660208b018b612a4c565b8960005b6020020135611af9565b6116ef6116e760408c0160208d01612a4c565b8a60016116ca565b611eae565b6001600160401b038916600090815260016020526040812090806117178b611e1d565b915091508060011461173b5760405162461bcd60e51b815260040161038b90612e38565b600161174a60208c018c612a4c565b600381111561175b5761175b612704565b146118155761177060408b0160208c01612a4c565b600381111561178157611781612704565b61178e60208c018c612a4c565b600381111561179f5761179f612704565b1480156117b05750883560208a0135145b6117ec5760405162461bcd60e51b815260206004820152600d60248201526c48414c5445445f4348414e474560981b604482015260640161038b565b61180d8c6001600160401b0316600090815260016020526040812060060155565b50505061197c565b600261182760408c0160208d01612a4c565b600381111561183857611838612704565b141561188157883560208a0135146118815760405162461bcd60e51b815260206004820152600c60248201526b4552524f525f4348414e474560a01b604482015260640161038b565b60408051600280825260608201835260009260208301908036833750505060058501549091506118b3908b3590611f83565b816000815181106118c6576118c6612a36565b60209081029190910101526118f48b60016020020160208101906118ea9190612a4c565b60208c013561210f565b8160018151811061190757611907612a36565b602090810291909101015260078401805460ff60401b1916600160411b1790556119348d60008b84611c41565b8c6001600160401b03167f24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db8460405161196f91815260200190565b60405180910390a2505050505b60006007820154600160401b900460ff16600281111561199e5761199e612704565b14156119aa5750611a37565b6040805180820190915281546001600160a01b031681526001820154602082015260048201546119da9042612dde565b816020018181516119eb9190612dde565b90525060028201805483546001600160a01b038083166001600160a01b031992831617865560038601805460018801558551929093169116179091556020909101519055426004909101555b5050505050505050565b6000838383604051602001611a5893929190612eca565b6040516020818303038152906040528051906020012090505b9392505050565b80518051602091820151828401518051908401516040516c23b637b130b61039ba30ba329d60991b95810195909552602d850193909352604d8401919091526001600160c01b031960c091821b8116606d85015291901b166075820152600090607d015b604051602081830303815290604052805190602001209050919050565b60006001836003811115611b0f57611b0f612704565b1415611b55576040516b213637b1b59039ba30ba329d60a11b6020820152602c8101839052604c015b6040516020818303038152906040528051906020012090506109a3565b6002836003811115611b6957611b69612704565b1415611b9f5760405174213637b1b59039ba30ba32961032b93937b932b21d60591b602082015260358101839052605501611b38565b6003836003811115611bb357611bb3612704565b1415611be25760405174213637b1b59039ba30ba3296103a37b7903330b91d60591b6020820152603501611b38565b60405162461bcd60e51b815260206004820152601060248201526f4241445f424c4f434b5f53544154555360801b604482015260640161038b565b6020810151600090815b602002015192915050565b60208101516000906001611c27565b6001821015611c5257611c52612be0565b600281511015611c6457611c64612be0565b6000611c71848484611a41565b6001600160401b038616600081815260016020526040908190206006018390555191925082917f86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d6890611cc890889088908890612f0c565b60405180910390a35050505050565b6001600160401b03821660008181526001602081905260408083206002808201805483546001600160a01b0319808216865596850188905595811690915560038301869055600480840187905560058401879055600684019690965560078301805468ffffffffffffffffff1916905590549251630357aa4960e01b8152948501959095526001600160a01b03948516602485018190529285166044850181905290949293909290911690630357aa4990606401600060405180830381600087803b158015611da557600080fd5b505af1158015611db9573d6000803e3d6000fd5b50505050846001600160401b03167ffdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f4085604051611df69190612cb1565b60405180910390a25050505050565b6001810154600090611e16836121b8565b1192915050565b600080806001611e306040860186612d6b565b611e3b929150612dde565b9050611e4b816020860135612f77565b9150611e5b606085013583612f8b565b611e66908535612e20565b92506002611e776040860186612d6b565b611e82929150612dde565b84606001351415611ea857611e9b816020860135612faa565b611ea59083612e20565b91505b50915091565b81611ebc6040850185612d6b565b8560600135818110611ed057611ed0612a36565b9050602002013514611f125760405162461bcd60e51b815260206004820152600b60248201526a15d493d391d7d4d510549560aa1b604482015260640161038b565b80611f206040850185612d6b565b611f2f60608701356001612e20565b818110611f3e57611f3e612a36565b905060200201351415611f7e5760405162461bcd60e51b815260206004820152600860248201526714d0535157d1539160c21b604482015260640161038b565b505050565b60408051600380825260808201909252600091829190816020015b6040805180820190915260008082526020820152815260200190600190039081611f9e57505060408051808201825260008082526020918201819052825180840190935260048352908201529091508160008151811061200057612000612a36565b602002602001018190525061201560006121ca565b8160018151811061202857612028612a36565b602002602001018190525061203d60006121ca565b8160028151811061205057612050612a36565b60209081029190910181019190915260408051808301825283815281518083019092528082526000928201929092526120a060408051606080820183529181019182529081526000602082015290565b604080518082018252606080825260006020808401829052845161012081018652828152908101879052938401859052908301829052608083018a905260a0830181905260c0830181905260e0830152610100820188905290612102816121fd565b9998505050505050505050565b6000600183600381111561212557612125612704565b141561213c5781604051602001611b389190612fbe565b600283600381111561215057612150612704565b141561217a576040516f26b0b1b434b7329032b93937b932b21d60811b6020820152603001611b38565b600383600381111561218e5761218e612704565b1415611be2576040516f26b0b1b434b732903a37b7903330b91d60811b6020820152603001611b38565b60008160040154426109a39190612dde565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b6000808251600381111561221357612213612704565b14156122c95761222682602001516123b6565b61223383604001516123b6565b612240846060015161243b565b608085015160a086015160c087015160e0808901516101008a01516040516f26b0b1b434b73290393ab73734b7339d60811b602082015260308101999099526050890197909752607088019590955260908701939093526001600160e01b031991831b821660b0870152821b811660b486015291901b1660b883015260bc82015260dc01611adc565b6001825160038111156122de576122de612704565b14156122f9578160800151604051602001611adc9190612fbe565b60028251600381111561230e5761230e612704565b1415612338576040516f26b0b1b434b7329032b93937b932b21d60811b6020820152603001611adc565b60038251600381111561234d5761234d612704565b1415612377576040516f26b0b1b434b732903a37b7903330b91d60811b6020820152603001611adc565b60405162461bcd60e51b815260206004820152600f60248201526e4241445f4d4143485f53544154555360881b604482015260640161038b565b919050565b60208101518151515160005b818110156124345783516123df906123da90836124d4565b61250c565b6040516b2b30b63ab29039ba30b1b59d60a11b6020820152602c810191909152604c8101849052606c01604051602081830303815290604052805190602001209250808061242c90612fe3565b9150506123c2565b5050919050565b602081015160005b8251518110156124ce576124738360000151828151811061246657612466612a36565b6020026020010151612529565b6040517129ba30b1b590333930b6b29039ba30b1b59d60711b602082015260328101919091526052810183905260720160405160208183030381529060405280519060200120915080806124c690612fe3565b915050612443565b50919050565b604080518082019091526000808252602082015282518051839081106124fc576124fc612a36565b6020026020010151905092915050565b600081600001518260200151604051602001611adc929190612ffe565b6000612538826000015161250c565b602080840151604080860151606087015191516b29ba30b1b590333930b6b29d60a11b94810194909452602c840194909452604c8301919091526001600160e01b031960e093841b8116606c840152921b9091166070820152607401611adc565b604080516101208101909152600060e08201818152610100830191909152819081526020016125d8604080518082019091526000808252602082015290565b815260006020820181905260408201819052606082018190526080820181905260a09091015290565b80604081018310156109a357600080fd5b80356001600160401b03811681146123b157600080fd5b6001600160a01b038116811461071457600080fd5b600080600080600080600080610200898b03121561265b57600080fd5b8835975061266c8a60208b01612601565b965061016089018a81111561268057600080fd5b60608a01965061268f81612612565b9550506101808901356126a181612629565b93506101a08901356126b281612629565b979a96995094979396929592945050506101c0820135916101e0013590565b6000602082840312156126e357600080fd5b611a7182612612565b80516001600160a01b03168252602090810151910152565b634e487b7160e01b600052602160045260246000fd5b6003811061272a5761272a612704565b9052565b6000610120820190506127428284516126ec565b602083015161275460408401826126ec565b5060408301516080830152606083015160a0830152608083015160c08301526001600160401b0360a08401511660e083015260c083015161279961010084018261271a565b5092915050565b6000602082840312156127b257600080fd5b5035919050565b61012081016127c8828a6126ec565b6127d560408301896126ec565b8660808301528560a08301528460c08301526001600160401b03841660e083015261063b61010083018461271a565b6000608082840312156124ce57600080fd5b6000806000806060858703121561282c57600080fd5b61283585612612565b935060208501356001600160401b038082111561285157600080fd5b61285d88838901612804565b9450604087013591508082111561287357600080fd5b818701915087601f83011261288757600080fd5b81358181111561289657600080fd5b8860208260051b85010111156128ab57600080fd5b95989497505060200194505050565b600080600080606085870312156128d057600080fd5b6128d985612612565b935060208501356001600160401b03808211156128f557600080fd5b61290188838901612804565b9450604087013591508082111561291757600080fd5b818701915087601f83011261292b57600080fd5b81358181111561293a57600080fd5b8860208285010111156128ab57600080fd5b6000806000806080858703121561296257600080fd5b843561296d81612629565b9350602085013561297d81612629565b9250604085013561298d81612629565b9150606085013561299d81612629565b939692955090935050565b600080600080600060e086880312156129c057600080fd5b6129c986612612565b945060208601356001600160401b038111156129e457600080fd5b6129f088828901612804565b945050612a008760408801612601565b9250612a0f8760808801612601565b9497939650919460c0013592915050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600060208284031215612a5e57600080fd5b813560048110611a7157600080fd5b604080519081016001600160401b0381118282101715612a8f57612a8f612a20565b60405290565b600082601f830112612aa657600080fd5b604051604081018181106001600160401b0382111715612ac857612ac8612a20565b8060405250806040840185811115612adf57600080fd5b845b81811015612b0057612af281612612565b835260209283019201612ae1565b509195945050505050565b600060808284031215612b1d57600080fd5b604051604081018181106001600160401b0382111715612b3f57612b3f612a20565b604052601f83018413612b5157600080fd5b612b59612a6d565b806040850186811115612b6b57600080fd5b855b81811015612b85578035845260209384019301612b6d565b50818452612b938782612a95565b6020850152509195945050505050565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b0380831681811415612bd657612bd6612ba3565b6001019392505050565b634e487b7160e01b600052600160045260246000fd5b604081833760006040838101828152908301915b6002811015612c39576001600160401b03612c2484612612565b16825260209283019290910190600101612c0a565b5050505050565b6101008101612c4f8285612bf6565b611a716080830184612bf6565b600060208083528351808285015260005b81811015612c8957858101830151858201604001528201612c6d565b81811115612c9b576000604083870101525b50601f01601f1916929092016040019392505050565b6020810160048310612cc557612cc5612704565b91905290565b6020808252600b908201526a21a420a62fa9a2a72222a960a91b604082015260600190565b6020808252600d908201526c4348414c5f444541444c494e4560981b604082015260600190565b6020808252600e908201526d4348414c5f4e4f545f424c4f434b60901b604082015260600190565b60208082526012908201527121a420a62fa727aa2fa2ac22a1aaaa24a7a760711b604082015260600190565b6000808335601e19843603018112612d8257600080fd5b8301803591506001600160401b03821115612d9c57600080fd5b6020019150600581901b3603821315612db457600080fd5b9250929050565b6020808252600990820152684249535f535441544560b81b604082015260600190565b600082821015612df057612df0612ba3565b500390565b6020808252601190820152704241445f4348414c4c454e47455f504f5360781b604082015260600190565b60008219821115612e3357612e33612ba3565b500190565b602080825260089082015267544f4f5f4c4f4e4760c01b604082015260600190565b8551815260018060a01b03602087015116602082015284604082015283606082015260a060808201528160a0820152818360c0830137600081830160c090810191909152601f909201601f19160101949350505050565b600060208284031215612ec357600080fd5b5051919050565b83815260006020848184015260408301845182860160005b82811015612efe57815184529284019290840190600101612ee2565b509198975050505050505050565b6000606082018583526020858185015260606040850152818551808452608086019150828701935060005b81811015612f5357845183529383019391830191600101612f37565b509098975050505050505050565b634e487b7160e01b600052601260045260246000fd5b600082612f8657612f86612f61565b500490565b6000816000190483118215151615612fa557612fa5612ba3565b500290565b600082612fb957612fb9612f61565b500690565b7026b0b1b434b732903334b734b9b432b21d60791b8152601181019190915260310190565b6000600019821415612ff757612ff7612ba3565b5060010190565b652b30b63ab29d60d11b815260006007841061301c5761301c612704565b5060f89290921b600683015260078201526027019056fea2646970667358221220714b52fdf00cdce958451acf97ad279c581c89d2a946c01657f967ffcf410fff64736f6c63430008090033",
}

// SingleExecutionChallengeABI is the input ABI used to generate the binding from.
// Deprecated: Use SingleExecutionChallengeMetaData.ABI instead.
var SingleExecutionChallengeABI = SingleExecutionChallengeMetaData.ABI

// SingleExecutionChallengeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SingleExecutionChallengeMetaData.Bin instead.
var SingleExecutionChallengeBin = SingleExecutionChallengeMetaData.Bin

// DeploySingleExecutionChallenge deploys a new Ethereum contract, binding an instance of SingleExecutionChallenge to it.
func DeploySingleExecutionChallenge(auth *bind.TransactOpts, backend bind.ContractBackend, osp_ common.Address, resultReceiver_ common.Address, maxInboxMessagesRead_ uint64, startAndEndHashes [2][32]byte, numSteps_ *big.Int, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (common.Address, *types.Transaction, *SingleExecutionChallenge, error) {
	parsed, err := SingleExecutionChallengeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SingleExecutionChallengeBin), backend, osp_, resultReceiver_, maxInboxMessagesRead_, startAndEndHashes, numSteps_, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SingleExecutionChallenge{SingleExecutionChallengeCaller: SingleExecutionChallengeCaller{contract: contract}, SingleExecutionChallengeTransactor: SingleExecutionChallengeTransactor{contract: contract}, SingleExecutionChallengeFilterer: SingleExecutionChallengeFilterer{contract: contract}}, nil
}

// SingleExecutionChallenge is an auto generated Go binding around an Ethereum contract.
type SingleExecutionChallenge struct {
	SingleExecutionChallengeCaller     // Read-only binding to the contract
	SingleExecutionChallengeTransactor // Write-only binding to the contract
	SingleExecutionChallengeFilterer   // Log filterer for contract events
}

// SingleExecutionChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SingleExecutionChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleExecutionChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SingleExecutionChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleExecutionChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SingleExecutionChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleExecutionChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SingleExecutionChallengeSession struct {
	Contract     *SingleExecutionChallenge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SingleExecutionChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SingleExecutionChallengeCallerSession struct {
	Contract *SingleExecutionChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// SingleExecutionChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SingleExecutionChallengeTransactorSession struct {
	Contract     *SingleExecutionChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// SingleExecutionChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SingleExecutionChallengeRaw struct {
	Contract *SingleExecutionChallenge // Generic contract binding to access the raw methods on
}

// SingleExecutionChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SingleExecutionChallengeCallerRaw struct {
	Contract *SingleExecutionChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// SingleExecutionChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SingleExecutionChallengeTransactorRaw struct {
	Contract *SingleExecutionChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSingleExecutionChallenge creates a new instance of SingleExecutionChallenge, bound to a specific deployed contract.
func NewSingleExecutionChallenge(address common.Address, backend bind.ContractBackend) (*SingleExecutionChallenge, error) {
	contract, err := bindSingleExecutionChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallenge{SingleExecutionChallengeCaller: SingleExecutionChallengeCaller{contract: contract}, SingleExecutionChallengeTransactor: SingleExecutionChallengeTransactor{contract: contract}, SingleExecutionChallengeFilterer: SingleExecutionChallengeFilterer{contract: contract}}, nil
}

// NewSingleExecutionChallengeCaller creates a new read-only instance of SingleExecutionChallenge, bound to a specific deployed contract.
func NewSingleExecutionChallengeCaller(address common.Address, caller bind.ContractCaller) (*SingleExecutionChallengeCaller, error) {
	contract, err := bindSingleExecutionChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeCaller{contract: contract}, nil
}

// NewSingleExecutionChallengeTransactor creates a new write-only instance of SingleExecutionChallenge, bound to a specific deployed contract.
func NewSingleExecutionChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*SingleExecutionChallengeTransactor, error) {
	contract, err := bindSingleExecutionChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeTransactor{contract: contract}, nil
}

// NewSingleExecutionChallengeFilterer creates a new log filterer instance of SingleExecutionChallenge, bound to a specific deployed contract.
func NewSingleExecutionChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*SingleExecutionChallengeFilterer, error) {
	contract, err := bindSingleExecutionChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeFilterer{contract: contract}, nil
}

// bindSingleExecutionChallenge binds a generic wrapper to an already deployed contract.
func bindSingleExecutionChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SingleExecutionChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SingleExecutionChallenge *SingleExecutionChallengeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SingleExecutionChallenge.Contract.SingleExecutionChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SingleExecutionChallenge *SingleExecutionChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.SingleExecutionChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SingleExecutionChallenge *SingleExecutionChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.SingleExecutionChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SingleExecutionChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) Bridge() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.Bridge(&_SingleExecutionChallenge.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) Bridge() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.Bridge(&_SingleExecutionChallenge.CallOpts)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) ChallengeInfo(opts *bind.CallOpts, challengeIndex uint64) (ChallengeLibChallenge, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "challengeInfo", challengeIndex)

	if err != nil {
		return *new(ChallengeLibChallenge), err
	}

	out0 := *abi.ConvertType(out[0], new(ChallengeLibChallenge)).(*ChallengeLibChallenge)

	return out0, err

}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) ChallengeInfo(challengeIndex uint64) (ChallengeLibChallenge, error) {
	return _SingleExecutionChallenge.Contract.ChallengeInfo(&_SingleExecutionChallenge.CallOpts, challengeIndex)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) ChallengeInfo(challengeIndex uint64) (ChallengeLibChallenge, error) {
	return _SingleExecutionChallenge.Contract.ChallengeInfo(&_SingleExecutionChallenge.CallOpts, challengeIndex)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) Challenges(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "challenges", arg0)

	outstruct := new(struct {
		Current            ChallengeLibParticipant
		Next               ChallengeLibParticipant
		LastMoveTimestamp  *big.Int
		WasmModuleRoot     [32]byte
		ChallengeStateHash [32]byte
		MaxInboxMessages   uint64
		Mode               uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Current = *abi.ConvertType(out[0], new(ChallengeLibParticipant)).(*ChallengeLibParticipant)
	outstruct.Next = *abi.ConvertType(out[1], new(ChallengeLibParticipant)).(*ChallengeLibParticipant)
	outstruct.LastMoveTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.WasmModuleRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.ChallengeStateHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.MaxInboxMessages = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.Mode = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) Challenges(arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	return _SingleExecutionChallenge.Contract.Challenges(&_SingleExecutionChallenge.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) Challenges(arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	return _SingleExecutionChallenge.Contract.Challenges(&_SingleExecutionChallenge.CallOpts, arg0)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) CurrentResponder(opts *bind.CallOpts, challengeIndex uint64) (common.Address, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "currentResponder", challengeIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _SingleExecutionChallenge.Contract.CurrentResponder(&_SingleExecutionChallenge.CallOpts, challengeIndex)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _SingleExecutionChallenge.Contract.CurrentResponder(&_SingleExecutionChallenge.CallOpts, challengeIndex)
}

// CurrentResponderTimeLeft is a free data retrieval call binding the contract method 0xc8fdb65b.
//
// Solidity: function currentResponderTimeLeft(uint64 challengeIndex) view returns(uint256)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) CurrentResponderTimeLeft(opts *bind.CallOpts, challengeIndex uint64) (*big.Int, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "currentResponderTimeLeft", challengeIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentResponderTimeLeft is a free data retrieval call binding the contract method 0xc8fdb65b.
//
// Solidity: function currentResponderTimeLeft(uint64 challengeIndex) view returns(uint256)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) CurrentResponderTimeLeft(challengeIndex uint64) (*big.Int, error) {
	return _SingleExecutionChallenge.Contract.CurrentResponderTimeLeft(&_SingleExecutionChallenge.CallOpts, challengeIndex)
}

// CurrentResponderTimeLeft is a free data retrieval call binding the contract method 0xc8fdb65b.
//
// Solidity: function currentResponderTimeLeft(uint64 challengeIndex) view returns(uint256)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) CurrentResponderTimeLeft(challengeIndex uint64) (*big.Int, error) {
	return _SingleExecutionChallenge.Contract.CurrentResponderTimeLeft(&_SingleExecutionChallenge.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) IsTimedOut(opts *bind.CallOpts, challengeIndex uint64) (bool, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "isTimedOut", challengeIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _SingleExecutionChallenge.Contract.IsTimedOut(&_SingleExecutionChallenge.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _SingleExecutionChallenge.Contract.IsTimedOut(&_SingleExecutionChallenge.CallOpts, challengeIndex)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) Osp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "osp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) Osp() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.Osp(&_SingleExecutionChallenge.CallOpts)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) Osp() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.Osp(&_SingleExecutionChallenge.CallOpts)
}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) ResultReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "resultReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) ResultReceiver() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.ResultReceiver(&_SingleExecutionChallenge.CallOpts)
}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) ResultReceiver() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.ResultReceiver(&_SingleExecutionChallenge.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) SequencerInbox() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.SequencerInbox(&_SingleExecutionChallenge.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) SequencerInbox() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.SequencerInbox(&_SingleExecutionChallenge.CallOpts)
}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) TotalChallengesCreated(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "totalChallengesCreated")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) TotalChallengesCreated() (uint64, error) {
	return _SingleExecutionChallenge.Contract.TotalChallengesCreated(&_SingleExecutionChallenge.CallOpts)
}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) TotalChallengesCreated() (uint64, error) {
	return _SingleExecutionChallenge.Contract.TotalChallengesCreated(&_SingleExecutionChallenge.CallOpts)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactor) BisectExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _SingleExecutionChallenge.contract.Transact(opts, "bisectExecution", challengeIndex, selection, newSegments)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) BisectExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.BisectExecution(&_SingleExecutionChallenge.TransactOpts, challengeIndex, selection, newSegments)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorSession) BisectExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.BisectExecution(&_SingleExecutionChallenge.TransactOpts, challengeIndex, selection, newSegments)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactor) ChallengeExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _SingleExecutionChallenge.contract.Transact(opts, "challengeExecution", challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) ChallengeExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.ChallengeExecution(&_SingleExecutionChallenge.TransactOpts, challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorSession) ChallengeExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.ChallengeExecution(&_SingleExecutionChallenge.TransactOpts, challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactor) ClearChallenge(opts *bind.TransactOpts, challengeIndex uint64) (*types.Transaction, error) {
	return _SingleExecutionChallenge.contract.Transact(opts, "clearChallenge", challengeIndex)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) ClearChallenge(challengeIndex uint64) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.ClearChallenge(&_SingleExecutionChallenge.TransactOpts, challengeIndex)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorSession) ClearChallenge(challengeIndex uint64) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.ClearChallenge(&_SingleExecutionChallenge.TransactOpts, challengeIndex)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactor) CreateChallenge(opts *bind.TransactOpts, wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _SingleExecutionChallenge.contract.Transact(opts, "createChallenge", wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.CreateChallenge(&_SingleExecutionChallenge.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.CreateChallenge(&_SingleExecutionChallenge.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactor) Initialize(opts *bind.TransactOpts, resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _SingleExecutionChallenge.contract.Transact(opts, "initialize", resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.Initialize(&_SingleExecutionChallenge.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.Initialize(&_SingleExecutionChallenge.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactor) OneStepProveExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _SingleExecutionChallenge.contract.Transact(opts, "oneStepProveExecution", challengeIndex, selection, proof)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) OneStepProveExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.OneStepProveExecution(&_SingleExecutionChallenge.TransactOpts, challengeIndex, selection, proof)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorSession) OneStepProveExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.OneStepProveExecution(&_SingleExecutionChallenge.TransactOpts, challengeIndex, selection, proof)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactor) Timeout(opts *bind.TransactOpts, challengeIndex uint64) (*types.Transaction, error) {
	return _SingleExecutionChallenge.contract.Transact(opts, "timeout", challengeIndex)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) Timeout(challengeIndex uint64) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.Timeout(&_SingleExecutionChallenge.TransactOpts, challengeIndex)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorSession) Timeout(challengeIndex uint64) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.Timeout(&_SingleExecutionChallenge.TransactOpts, challengeIndex)
}

// SingleExecutionChallengeBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeBisectedIterator struct {
	Event *SingleExecutionChallengeBisected // Event containing the contract specifics and raw log

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
func (it *SingleExecutionChallengeBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleExecutionChallengeBisected)
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
		it.Event = new(SingleExecutionChallengeBisected)
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
func (it *SingleExecutionChallengeBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleExecutionChallengeBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleExecutionChallengeBisected represents a Bisected event raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeBisected struct {
	ChallengeIndex          uint64
	ChallengeRoot           [32]byte
	ChallengedSegmentStart  *big.Int
	ChallengedSegmentLength *big.Int
	ChainHashes             [][32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBisected is a free log retrieval operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) FilterBisected(opts *bind.FilterOpts, challengeIndex []uint64, challengeRoot [][32]byte) (*SingleExecutionChallengeBisectedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.FilterLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeBisectedIterator{contract: _SingleExecutionChallenge.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *SingleExecutionChallengeBisected, challengeIndex []uint64, challengeRoot [][32]byte) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.WatchLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleExecutionChallengeBisected)
				if err := _SingleExecutionChallenge.contract.UnpackLog(event, "Bisected", log); err != nil {
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

// ParseBisected is a log parse operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) ParseBisected(log types.Log) (*SingleExecutionChallengeBisected, error) {
	event := new(SingleExecutionChallengeBisected)
	if err := _SingleExecutionChallenge.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingleExecutionChallengeChallengeEndedIterator is returned from FilterChallengeEnded and is used to iterate over the raw logs and unpacked data for ChallengeEnded events raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeChallengeEndedIterator struct {
	Event *SingleExecutionChallengeChallengeEnded // Event containing the contract specifics and raw log

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
func (it *SingleExecutionChallengeChallengeEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleExecutionChallengeChallengeEnded)
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
		it.Event = new(SingleExecutionChallengeChallengeEnded)
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
func (it *SingleExecutionChallengeChallengeEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleExecutionChallengeChallengeEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleExecutionChallengeChallengeEnded represents a ChallengeEnded event raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeChallengeEnded struct {
	ChallengeIndex uint64
	Kind           uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChallengeEnded is a free log retrieval operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) FilterChallengeEnded(opts *bind.FilterOpts, challengeIndex []uint64) (*SingleExecutionChallengeChallengeEndedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.FilterLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeChallengeEndedIterator{contract: _SingleExecutionChallenge.contract, event: "ChallengeEnded", logs: logs, sub: sub}, nil
}

// WatchChallengeEnded is a free log subscription operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) WatchChallengeEnded(opts *bind.WatchOpts, sink chan<- *SingleExecutionChallengeChallengeEnded, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.WatchLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleExecutionChallengeChallengeEnded)
				if err := _SingleExecutionChallenge.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
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

// ParseChallengeEnded is a log parse operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) ParseChallengeEnded(log types.Log) (*SingleExecutionChallengeChallengeEnded, error) {
	event := new(SingleExecutionChallengeChallengeEnded)
	if err := _SingleExecutionChallenge.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingleExecutionChallengeExecutionChallengeBegunIterator is returned from FilterExecutionChallengeBegun and is used to iterate over the raw logs and unpacked data for ExecutionChallengeBegun events raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeExecutionChallengeBegunIterator struct {
	Event *SingleExecutionChallengeExecutionChallengeBegun // Event containing the contract specifics and raw log

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
func (it *SingleExecutionChallengeExecutionChallengeBegunIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleExecutionChallengeExecutionChallengeBegun)
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
		it.Event = new(SingleExecutionChallengeExecutionChallengeBegun)
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
func (it *SingleExecutionChallengeExecutionChallengeBegunIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleExecutionChallengeExecutionChallengeBegunIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleExecutionChallengeExecutionChallengeBegun represents a ExecutionChallengeBegun event raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeExecutionChallengeBegun struct {
	ChallengeIndex uint64
	BlockSteps     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExecutionChallengeBegun is a free log retrieval operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) FilterExecutionChallengeBegun(opts *bind.FilterOpts, challengeIndex []uint64) (*SingleExecutionChallengeExecutionChallengeBegunIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.FilterLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeExecutionChallengeBegunIterator{contract: _SingleExecutionChallenge.contract, event: "ExecutionChallengeBegun", logs: logs, sub: sub}, nil
}

// WatchExecutionChallengeBegun is a free log subscription operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) WatchExecutionChallengeBegun(opts *bind.WatchOpts, sink chan<- *SingleExecutionChallengeExecutionChallengeBegun, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.WatchLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleExecutionChallengeExecutionChallengeBegun)
				if err := _SingleExecutionChallenge.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
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

// ParseExecutionChallengeBegun is a log parse operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) ParseExecutionChallengeBegun(log types.Log) (*SingleExecutionChallengeExecutionChallengeBegun, error) {
	event := new(SingleExecutionChallengeExecutionChallengeBegun)
	if err := _SingleExecutionChallenge.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingleExecutionChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeInitiatedChallengeIterator struct {
	Event *SingleExecutionChallengeInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *SingleExecutionChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleExecutionChallengeInitiatedChallenge)
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
		it.Event = new(SingleExecutionChallengeInitiatedChallenge)
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
func (it *SingleExecutionChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleExecutionChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleExecutionChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeInitiatedChallenge struct {
	ChallengeIndex uint64
	StartState     GlobalState
	EndState       GlobalState
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts, challengeIndex []uint64) (*SingleExecutionChallengeInitiatedChallengeIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.FilterLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeInitiatedChallengeIterator{contract: _SingleExecutionChallenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *SingleExecutionChallengeInitiatedChallenge, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.WatchLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleExecutionChallengeInitiatedChallenge)
				if err := _SingleExecutionChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*SingleExecutionChallengeInitiatedChallenge, error) {
	event := new(SingleExecutionChallengeInitiatedChallenge)
	if err := _SingleExecutionChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingleExecutionChallengeOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeOneStepProofCompletedIterator struct {
	Event *SingleExecutionChallengeOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *SingleExecutionChallengeOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleExecutionChallengeOneStepProofCompleted)
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
		it.Event = new(SingleExecutionChallengeOneStepProofCompleted)
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
func (it *SingleExecutionChallengeOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleExecutionChallengeOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleExecutionChallengeOneStepProofCompleted represents a OneStepProofCompleted event raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeOneStepProofCompleted struct {
	ChallengeIndex uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts, challengeIndex []uint64) (*SingleExecutionChallengeOneStepProofCompletedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.FilterLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeOneStepProofCompletedIterator{contract: _SingleExecutionChallenge.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *SingleExecutionChallengeOneStepProofCompleted, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.WatchLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleExecutionChallengeOneStepProofCompleted)
				if err := _SingleExecutionChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) ParseOneStepProofCompleted(log types.Log) (*SingleExecutionChallengeOneStepProofCompleted, error) {
	event := new(SingleExecutionChallengeOneStepProofCompleted)
	if err := _SingleExecutionChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
