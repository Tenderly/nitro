// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challengegen

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

// ChallengeLibMetaData contains all meta data concerning the ChallengeLib contract.
var ChallengeLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122061be17751f5e97fb60b2309e008b8d79e1acd48d559da96540694217eab782c864736f6c63430008090033",
}

// ChallengeLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ChallengeLibMetaData.ABI instead.
var ChallengeLibABI = ChallengeLibMetaData.ABI

// ChallengeLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChallengeLibMetaData.Bin instead.
var ChallengeLibBin = ChallengeLibMetaData.Bin

// DeployChallengeLib deploys a new Ethereum contract, binding an instance of ChallengeLib to it.
func DeployChallengeLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeLib, error) {
	parsed, err := ChallengeLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChallengeLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeLib{ChallengeLibCaller: ChallengeLibCaller{contract: contract}, ChallengeLibTransactor: ChallengeLibTransactor{contract: contract}, ChallengeLibFilterer: ChallengeLibFilterer{contract: contract}}, nil
}

// ChallengeLib is an auto generated Go binding around an Ethereum contract.
type ChallengeLib struct {
	ChallengeLibCaller     // Read-only binding to the contract
	ChallengeLibTransactor // Write-only binding to the contract
	ChallengeLibFilterer   // Log filterer for contract events
}

// ChallengeLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeLibSession struct {
	Contract     *ChallengeLib     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeLibCallerSession struct {
	Contract *ChallengeLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ChallengeLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeLibTransactorSession struct {
	Contract     *ChallengeLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ChallengeLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeLibRaw struct {
	Contract *ChallengeLib // Generic contract binding to access the raw methods on
}

// ChallengeLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeLibCallerRaw struct {
	Contract *ChallengeLibCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeLibTransactorRaw struct {
	Contract *ChallengeLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeLib creates a new instance of ChallengeLib, bound to a specific deployed contract.
func NewChallengeLib(address common.Address, backend bind.ContractBackend) (*ChallengeLib, error) {
	contract, err := bindChallengeLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeLib{ChallengeLibCaller: ChallengeLibCaller{contract: contract}, ChallengeLibTransactor: ChallengeLibTransactor{contract: contract}, ChallengeLibFilterer: ChallengeLibFilterer{contract: contract}}, nil
}

// NewChallengeLibCaller creates a new read-only instance of ChallengeLib, bound to a specific deployed contract.
func NewChallengeLibCaller(address common.Address, caller bind.ContractCaller) (*ChallengeLibCaller, error) {
	contract, err := bindChallengeLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeLibCaller{contract: contract}, nil
}

// NewChallengeLibTransactor creates a new write-only instance of ChallengeLib, bound to a specific deployed contract.
func NewChallengeLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeLibTransactor, error) {
	contract, err := bindChallengeLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeLibTransactor{contract: contract}, nil
}

// NewChallengeLibFilterer creates a new log filterer instance of ChallengeLib, bound to a specific deployed contract.
func NewChallengeLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeLibFilterer, error) {
	contract, err := bindChallengeLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeLibFilterer{contract: contract}, nil
}

// bindChallengeLib binds a generic wrapper to an already deployed contract.
func bindChallengeLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChallengeLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeLib *ChallengeLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeLib.Contract.ChallengeLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeLib *ChallengeLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeLib.Contract.ChallengeLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeLib *ChallengeLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeLib.Contract.ChallengeLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeLib *ChallengeLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeLib *ChallengeLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeLib *ChallengeLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeLib.Contract.contract.Transact(opts, method, params...)
}

// ChallengeManagerMetaData contains all meta data concerning the ChallengeManager contract.
var ChallengeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"challengeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"chainHashes\",\"type\":\"bytes32[]\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIChallengeManager.ChallengeTerminationType\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"ChallengeEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockSteps\",\"type\":\"uint256\"}],\"name\":\"ExecutionChallengeBegun\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"startState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"endState\",\"type\":\"tuple\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"newSegments\",\"type\":\"bytes32[]\"}],\"name\":\"bisectExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"machineStatuses\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"globalStateHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256\",\"name\":\"numSteps\",\"type\":\"uint256\"}],\"name\":\"challengeExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"challengeInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"internalType\":\"structChallengeLib.Challenge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"challenges\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"clearChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"startAndEndMachineStatuses_\",\"type\":\"uint8[2]\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState[2]\",\"name\":\"startAndEndGlobalStates_\",\"type\":\"tuple[2]\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"asserter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"asserterTimeLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengerTimeLeft_\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"currentResponder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"}],\"name\":\"getOsp\",\"outputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChallengeResultReceiver\",\"name\":\"resultReceiver_\",\"type\":\"address\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"sequencerInbox_\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"osp_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"isTimedOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProveExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"osp\",\"outputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ospCond\",\"outputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"osp_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"condRoot\",\"type\":\"bytes32\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"condOsp\",\"type\":\"address\"}],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resultReceiver\",\"outputs\":[{\"internalType\":\"contractIChallengeResultReceiver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalChallengesCreated\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b50608051612cab6100376000396000818161082501526114350152612cab6000f3fe608060405234801561001057600080fd5b50600436106101115760003560e01c80638f1d3776116100ad578063e78cea9211610071578063e78cea921461031a578063ee35f3271461032d578063f26a62c614610340578063f8c8765e14610353578063fb7be0a11461036657600080fd5b80638f1d3776146102055780639ede42b9146102a8578063a521b032146102cb578063d248d124146102de578063dc74bf8b146102f157600080fd5b806314eab5e7146101165780631b45c86a1461014657806323a9ef231461015b5780633504f1d7146101865780633690b011146101995780635038934d146101ac57806356e9df97146101bf5780635ef489e6146101d25780637fd07a9c146101e5575b600080fd5b610129610124366004612241565b610379565b6040516001600160401b0390911681526020015b60405180910390f35b6101596101543660046122d4565b6106f0565b005b61016e6101693660046122d4565b6107c0565b6040516001600160a01b03909116815260200161013d565b60025461016e906001600160a01b031681565b61016e6101a73660046122ef565b6107e4565b6101596101ba366004612308565b61081a565b6101596101cd3660046122d4565b610900565b600054610129906001600160401b031681565b6101f86101f33660046122d4565b610a6e565b60405161013d919061238c565b6102956102133660046122ef565b6001602081815260009283526040928390208351808501855281546001600160a01b0390811682529382015481840152845180860190955260028201549093168452600381015491840191909152600481015460058201546006830154600790930154939493919290916001600160401b03811690600160401b900460ff1687565b60405161013d97969594939291906123fe565b6102bb6102b63660046122d4565b610b47565b604051901515815260200161013d565b6101596102d9366004612461565b610b68565b6101596102ec366004612505565b610fde565b61016e6102ff3660046122ef565b6006602052600090815260409020546001600160a01b031681565b60045461016e906001600160a01b031681565b60035461016e906001600160a01b031681565b60055461016e906001600160a01b031681565b610159610361366004612597565b61142a565b6101596103743660046125f3565b611556565b6002546000906001600160a01b031633146103ce5760405162461bcd60e51b815260206004820152601060248201526f13d3931657d493d313155417d0d2105360821b60448201526064015b60405180910390fd5b6040805160028082526060820183526000926020830190803683370190505090506104936103ff60208b018b612697565b61048e8a60005b6080020180360381019061041a9190612756565b8051805160209182015192820151805190830151604080516c23b637b130b61039ba30ba329d60991b81870152602d810194909452604d8401959095526001600160c01b031960c092831b8116606d850152911b1660758201528251808203605d018152607d909101909252815191012090565b611cc4565b816000815181106104a6576104a6612681565b60209081029190910101526104d58960016020020160208101906104ca9190612697565b61048e8a6001610406565b816001815181106104e8576104e8612681565b6020908102919091010152600080548190819061050d906001600160401b0316612804565b91906101000a8154816001600160401b0302191690836001600160401b031602179055905060006001600160401b0316816001600160401b031614156105555761055561282b565b6001600160401b0381166000908152600160205260408120600581018d90559061058f61058a368d90038d0160808e01612756565b611de8565b905060026105a360408e0160208f01612697565b60038111156105b4576105b4612362565b14806105e2575060006105d76105d2368e90038e0160808f01612756565b611dfd565b6001600160401b0316115b156105f557806105f181612804565b9150505b6007820180546040805180820182526001600160a01b038d811680835260209283018d90526002880180546001600160a01b03199081169092179055600388018d905583518085018552918e16808352919092018b90528654909116178555600185018990554260048601556001600160401b0384811668ffffffffffffffffff1990931692909217600160401b179092559051908416907f76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a906106bf908e90608082019061288b565b60405180910390a26106dd8360008c6001600160401b031687611e0c565b5090925050505b98975050505050505050565b60006001600160401b038216600090815260016020526040902060070154600160401b900460ff16600281111561072957610729612362565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906107695760405162461bcd60e51b81526004016103c591906128a7565b5061077381610b47565b6107b25760405162461bcd60e51b815260206004820152601060248201526f54494d454f55545f444541444c494e4560801b60448201526064016103c5565b6107bd816000611ea2565b50565b6001600160401b03166000908152600160205260409020546001600160a01b031690565b6000818152600660205260408120546001600160a01b0316806108145750506005546001600160a01b0316919050565b92915050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156108635760405162461bcd60e51b81526004016103c5906128fc565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61038054336001600160a01b038216146108c057604051631194af8760e11b81523360048201526001600160a01b03821660248201526044016103c5565b505060009182526006602052604090912080546001600160a01b039283166001600160a01b03199182161790915560058054939092169216919091179055565b6002546001600160a01b0316331461094d5760405162461bcd60e51b815260206004820152601060248201526f2727aa2fa922a9afa922a1a2a4ab22a960811b60448201526064016103c5565b60006001600160401b038216600090815260016020526040902060070154600160401b900460ff16600281111561098657610986612362565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906109c65760405162461bcd60e51b81526004016103c591906128a7565b506001600160401b038116600081815260016020819052604080832080546001600160a01b031990811682559281018490556002810180549093169092556003808301849055600483018490556005830184905560068301939093556007909101805468ffffffffffffffffff19169055517ffdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f4091610a6391612958565b60405180910390a250565b610a76612197565b6001600160401b0382811660009081526001602081815260409283902083516101208101855281546001600160a01b0390811660e0830190815294830154610100830152938152845180860186526002808401549095168152600383015481850152928101929092526004810154938201939093526005830154606082015260068301546080820152600783015493841660a08201529260c0840191600160401b90910460ff1690811115610b2d57610b2d612362565b6002811115610b3e57610b3e612362565b90525092915050565b6001600160401b038116600090815260016020526040812061081490611fd0565b6001600160401b038416600090815260016020526040812085918591610b8d846107c0565b6001600160a01b0316336001600160a01b031614610bbd5760405162461bcd60e51b81526004016103c59061296b565b610bc684610b47565b15610be35760405162461bcd60e51b81526004016103c590612990565b6000826002811115610bf757610bf7612362565b1415610c655760006007820154600160401b900460ff166002811115610c1f57610c1f612362565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b81525090610c5f5760405162461bcd60e51b81526004016103c591906128a7565b50610d24565b6001826002811115610c7957610c79612362565b1415610cc35760016007820154600160401b900460ff166002811115610ca157610ca1612362565b14610cbe5760405162461bcd60e51b81526004016103c5906129b7565b610d24565b6002826002811115610cd757610cd7612362565b1415610d1c5760026007820154600160401b900460ff166002811115610cff57610cff612362565b14610cbe5760405162461bcd60e51b81526004016103c5906129df565b610d2461282b565b610d7283356020850135610d3b6040870187612a0b565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250611fe892505050565b816006015414610d945760405162461bcd60e51b81526004016103c590612a5b565b6002610da36040850185612a0b565b90501080610dce57506001610dbb6040850185612a0b565b610dc6929150612a7e565b836060013510155b15610deb5760405162461bcd60e51b81526004016103c590612a95565b600080610df78961201f565b9150915060018111610e375760405162461bcd60e51b81526020600482015260096024820152681513d3d7d4d213d49560ba1b60448201526064016103c5565b806028811115610e45575060285b610e50816001612ac0565b8814610e8d5760405162461bcd60e51b815260206004820152600c60248201526b57524f4e475f44454752454560a01b60448201526064016103c5565b50610ed78989896000818110610ea557610ea5612681565b602002919091013590508a8a610ebc600182612a7e565b818110610ecb57610ecb612681565b905060200201356120b0565b610f168a83838b8b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250611e0c92505050565b50600090505b6007820154600160401b900460ff166002811115610f3c57610f3c612362565b1415610f485750610fd5565b6040805180820190915281546001600160a01b03168152600182015460208201526004820154610f789042612a7e565b81602001818151610f899190612a7e565b90525060028201805483546001600160a01b038083166001600160a01b031992831617865560038601805460018801558551929093169116179091556020909101519055426004909101555b50505050505050565b6001600160401b038416600090815260016020526040902084908490600290611006846107c0565b6001600160a01b0316336001600160a01b0316146110365760405162461bcd60e51b81526004016103c59061296b565b61103f84610b47565b1561105c5760405162461bcd60e51b81526004016103c590612990565b600082600281111561107057611070612362565b14156110de5760006007820154600160401b900460ff16600281111561109857611098612362565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906110d85760405162461bcd60e51b81526004016103c591906128a7565b5061119d565b60018260028111156110f2576110f2612362565b141561113c5760016007820154600160401b900460ff16600281111561111a5761111a612362565b146111375760405162461bcd60e51b81526004016103c5906129b7565b61119d565b600282600281111561115057611150612362565b14156111955760026007820154600160401b900460ff16600281111561117857611178612362565b146111375760405162461bcd60e51b81526004016103c5906129df565b61119d61282b565b6111b483356020850135610d3b6040870187612a0b565b8160060154146111d65760405162461bcd60e51b81526004016103c590612a5b565b60026111e56040850185612a0b565b90501080611210575060016111fd6040850185612a0b565b611208929150612a7e565b836060013510155b1561122d5760405162461bcd60e51b81526004016103c590612a95565b6001600160401b038816600090815260016020526040812090806112508a61201f565b9092509050600181146112755760405162461bcd60e51b81526004016103c590612ad8565b50600061128583600501546107e4565b60408051808201825260078601546001600160401b031681526004546001600160a01b0390811660208301529290921691635d3adcfb9185906112ca908f018f612a0b565b8f606001358181106112de576112de612681565b905060200201358d8d6040518663ffffffff1660e01b8152600401611307959493929190612afa565b60206040518083038186803b15801561131f57600080fd5b505afa158015611333573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113579190612b51565b905061136660408b018b612a0b565b61137560608d01356001612ac0565b81811061138457611384612681565b905060200201358114156113c95760405162461bcd60e51b815260206004820152600c60248201526b14d0535157d3d4d417d1539160a21b60448201526064016103c5565b6040516001600160401b038c16907fc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e90600090a261141e8b6001600160401b0316600090815260016020526040812060060155565b5060009150610f1c9050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156114735760405162461bcd60e51b81526004016103c5906128fc565b6002546001600160a01b0316156114bb5760405162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b60448201526064016103c5565b6001600160a01b0384166115065760405162461bcd60e51b81526020600482015260126024820152712727afa922a9aaa62a2fa922a1a2a4ab22a960711b60448201526064016103c5565b600280546001600160a01b039586166001600160a01b0319918216179091556003805494861694821694909417909355600480549285169284169290921790915560058054919093169116179055565b6001600160401b03851660009081526001602081905260409091208691869161157e846107c0565b6001600160a01b0316336001600160a01b0316146115ae5760405162461bcd60e51b81526004016103c59061296b565b6115b784610b47565b156115d45760405162461bcd60e51b81526004016103c590612990565b60008260028111156115e8576115e8612362565b14156116565760006007820154600160401b900460ff16600281111561161057611610612362565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906116505760405162461bcd60e51b81526004016103c591906128a7565b50611715565b600182600281111561166a5761166a612362565b14156116b45760016007820154600160401b900460ff16600281111561169257611692612362565b146116af5760405162461bcd60e51b81526004016103c5906129b7565b611715565b60028260028111156116c8576116c8612362565b141561170d5760026007820154600160401b900460ff1660028111156116f0576116f0612362565b146116af5760405162461bcd60e51b81526004016103c5906129df565b61171561282b565b61172c83356020850135610d3b6040870187612a0b565b81600601541461174e5760405162461bcd60e51b81526004016103c590612a5b565b600261175d6040850185612a0b565b90501080611788575060016117756040850185612a0b565b611780929150612a7e565b836060013510155b156117a55760405162461bcd60e51b81526004016103c590612a95565b60018510156117ec5760405162461bcd60e51b815260206004820152601360248201527210d2105313115391d157d513d3d7d4d213d495606a1b60448201526064016103c5565b650800000000008511156118375760405162461bcd60e51b81526020600482015260126024820152714348414c4c454e47455f544f4f5f4c4f4e4760701b60448201526064016103c5565b6118798861185961184b60208b018b612697565b8960005b6020020135611cc4565b61187461186c60408c0160208d01612697565b8a600161184f565b6120b0565b6001600160401b0389166000908152600160205260408120908061189c8b61201f565b91509150806001146118c05760405162461bcd60e51b81526004016103c590612ad8565b60016118cf60208c018c612697565b60038111156118e0576118e0612362565b1461199a576118f560408b0160208c01612697565b600381111561190657611906612362565b61191360208c018c612697565b600381111561192457611924612362565b1480156119355750883560208a0135145b6119715760405162461bcd60e51b815260206004820152600d60248201526c48414c5445445f4348414e474560981b60448201526064016103c5565b6119928c6001600160401b0316600090815260016020526040812060060155565b505050611bff565b60026119ac60408c0160208d01612697565b60038111156119bd576119bd612362565b1415611a0657883560208a013514611a065760405162461bcd60e51b815260206004820152600c60248201526b4552524f525f4348414e474560a01b60448201526064016103c5565b6040805160028082526060820183526000926020830190803683370190505090506000611a3685600501546107e4565b60058601546040516301265ef960e21b81528d35600482015260248101919091529091506001600160a01b038216906304997be49060440160206040518083038186803b158015611a8657600080fd5b505afa158015611a9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611abe9190612b51565b82600081518110611ad157611ad1612681565b60209081029190910101526001600160a01b03811663d8558b878d6001602002016020810190611b019190612697565b8d600160200201356040518363ffffffff1660e01b8152600401611b26929190612b6a565b60206040518083038186803b158015611b3e57600080fd5b505afa158015611b52573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b769190612b51565b82600181518110611b8957611b89612681565b602090810291909101015260078501805460ff60401b1916600160411b179055611bb68e60008c85611e0c565b8d6001600160401b03167f24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db85604051611bf191815260200190565b60405180910390a250505050505b60006007820154600160401b900460ff166002811115611c2157611c21612362565b1415611c2d5750611cba565b6040805180820190915281546001600160a01b03168152600182015460208201526004820154611c5d9042612a7e565b81602001818151611c6e9190612a7e565b90525060028201805483546001600160a01b038083166001600160a01b031992831617865560038601805460018801558551929093169116179091556020909101519055426004909101555b5050505050505050565b60006001836003811115611cda57611cda612362565b1415611d20576040516b213637b1b59039ba30ba329d60a11b6020820152602c8101839052604c015b604051602081830303815290604052805190602001209050610814565b6002836003811115611d3457611d34612362565b1415611d6a5760405174213637b1b59039ba30ba32961032b93937b932b21d60591b602082015260358101839052605501611d03565b6003836003811115611d7e57611d7e612362565b1415611dad5760405174213637b1b59039ba30ba3296103a37b7903330b91d60591b6020820152603501611d03565b60405162461bcd60e51b815260206004820152601060248201526f4241445f424c4f434b5f53544154555360801b60448201526064016103c5565b6020810151600090815b602002015192915050565b60208101516000906001611df2565b6001821015611e1d57611e1d61282b565b600281511015611e2f57611e2f61282b565b6000611e3c848484611fe8565b6001600160401b038616600081815260016020526040908190206006018390555191925082917f86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d6890611e9390889088908890612b81565b60405180910390a35050505050565b6001600160401b03821660008181526001602081905260408083206002808201805483546001600160a01b0319808216865596850188905595811690915560038301869055600480840187905560058401879055600684019690965560078301805468ffffffffffffffffff1916905590549251630357aa4960e01b8152948501959095526001600160a01b03948516602485018190529285166044850181905290949293909290911690630357aa4990606401600060405180830381600087803b158015611f7057600080fd5b505af1158015611f84573d6000803e3d6000fd5b50505050846001600160401b03167ffdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f4085604051611fc19190612958565b60405180910390a25050505050565b6001810154600090611fe183612185565b1192915050565b6000838383604051602001611fff93929190612bd6565b6040516020818303038152906040528051906020012090505b9392505050565b6000808060016120326040860186612a0b565b61203d929150612a7e565b905061204d816020860135612c2e565b915061205d606085013583612c42565b612068908535612ac0565b925060026120796040860186612a0b565b612084929150612a7e565b846060013514156120aa5761209d816020860135612c61565b6120a79083612ac0565b91505b50915091565b816120be6040850185612a0b565b85606001358181106120d2576120d2612681565b90506020020135146121145760405162461bcd60e51b815260206004820152600b60248201526a15d493d391d7d4d510549560aa1b60448201526064016103c5565b806121226040850185612a0b565b61213160608701356001612ac0565b81811061214057612140612681565b9050602002013514156121805760405162461bcd60e51b815260206004820152600860248201526714d0535157d1539160c21b60448201526064016103c5565b505050565b60008160040154426108149190612a7e565b604080516101208101909152600060e08201818152610100830191909152819081526020016121d6604080518082019091526000808252602082015290565b815260006020820181905260408201819052606082018190526080820181905260a09091015290565b806040810183101561081457600080fd5b80356001600160401b038116811461222757600080fd5b919050565b6001600160a01b03811681146107bd57600080fd5b600080600080600080600080610200898b03121561225e57600080fd5b8835975061226f8a60208b016121ff565b965061016089018a81111561228357600080fd5b60608a01965061229281612210565b9550506101808901356122a48161222c565b93506101a08901356122b58161222c565b979a96995094979396929592945050506101c0820135916101e0013590565b6000602082840312156122e657600080fd5b61201882612210565b60006020828403121561230157600080fd5b5035919050565b60008060006060848603121561231d57600080fd5b83356123288161222c565b925060208401359150604084013561233f8161222c565b809150509250925092565b80516001600160a01b03168252602090810151910152565b634e487b7160e01b600052602160045260246000fd5b6003811061238857612388612362565b9052565b6000610120820190506123a082845161234a565b60208301516123b2604084018261234a565b5060408301516080830152606083015160a0830152608083015160c08301526001600160401b0360a08401511660e083015260c08301516123f7610100840182612378565b5092915050565b610120810161240d828a61234a565b61241a604083018961234a565b8660808301528560a08301528460c08301526001600160401b03841660e08301526106e4610100830184612378565b60006080828403121561245b57600080fd5b50919050565b6000806000806060858703121561247757600080fd5b61248085612210565b935060208501356001600160401b038082111561249c57600080fd5b6124a888838901612449565b945060408701359150808211156124be57600080fd5b818701915087601f8301126124d257600080fd5b8135818111156124e157600080fd5b8860208260051b85010111156124f657600080fd5b95989497505060200194505050565b6000806000806060858703121561251b57600080fd5b61252485612210565b935060208501356001600160401b038082111561254057600080fd5b61254c88838901612449565b9450604087013591508082111561256257600080fd5b818701915087601f83011261257657600080fd5b81358181111561258557600080fd5b8860208285010111156124f657600080fd5b600080600080608085870312156125ad57600080fd5b84356125b88161222c565b935060208501356125c88161222c565b925060408501356125d88161222c565b915060608501356125e88161222c565b939692955090935050565b600080600080600060e0868803121561260b57600080fd5b61261486612210565b945060208601356001600160401b0381111561262f57600080fd5b61263b88828901612449565b94505061264b87604088016121ff565b925061265a87608088016121ff565b9497939650919460c0013592915050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6000602082840312156126a957600080fd5b81356004811061201857600080fd5b604080519081016001600160401b03811182821017156126da576126da61266b565b60405290565b600082601f8301126126f157600080fd5b604051604081018181106001600160401b03821117156127135761271361266b565b806040525080604084018581111561272a57600080fd5b845b8181101561274b5761273d81612210565b83526020928301920161272c565b509195945050505050565b60006080828403121561276857600080fd5b604051604081018181106001600160401b038211171561278a5761278a61266b565b604052601f8301841361279c57600080fd5b6127a46126b8565b8060408501868111156127b657600080fd5b855b818110156127d05780358452602093840193016127b8565b508184526127de87826126e0565b6020850152509195945050505050565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b0380831681811415612821576128216127ee565b6001019392505050565b634e487b7160e01b600052600160045260246000fd5b604081833760006040838101828152908301915b6002811015612884576001600160401b0361286f84612210565b16825260209283019290910190600101612855565b5050505050565b610100810161289a8285612841565b6120186080830184612841565b600060208083528351808285015260005b818110156128d4578581018301518582016040015282016128b8565b818111156128e6576000604083870101525b50601f01601f1916929092016040019392505050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b600481106107bd576107bd612362565b6020810161296583612948565b91905290565b6020808252600b908201526a21a420a62fa9a2a72222a960a91b604082015260600190565b6020808252600d908201526c4348414c5f444541444c494e4560981b604082015260600190565b6020808252600e908201526d4348414c5f4e4f545f424c4f434b60901b604082015260600190565b60208082526012908201527121a420a62fa727aa2fa2ac22a1aaaa24a7a760711b604082015260600190565b6000808335601e19843603018112612a2257600080fd5b8301803591506001600160401b03821115612a3c57600080fd5b6020019150600581901b3603821315612a5457600080fd5b9250929050565b6020808252600990820152684249535f535441544560b81b604082015260600190565b600082821015612a9057612a906127ee565b500390565b6020808252601190820152704241445f4348414c4c454e47455f504f5360781b604082015260600190565b60008219821115612ad357612ad36127ee565b500190565b602080825260089082015267544f4f5f4c4f4e4760c01b604082015260600190565b8551815260018060a01b03602087015116602082015284604082015283606082015260a060808201528160a0820152818360c0830137600081830160c090810191909152601f909201601f19160101949350505050565b600060208284031215612b6357600080fd5b5051919050565b60408101612b7784612948565b9281526020015290565b6000606082018583526020858185015260606040850152818551808452608086019150828701935060005b81811015612bc857845183529383019391830191600101612bac565b509098975050505050505050565b83815260006020848184015260408301845182860160005b82811015612c0a57815184529284019290840190600101612bee565b509198975050505050505050565b634e487b7160e01b600052601260045260246000fd5b600082612c3d57612c3d612c18565b500490565b6000816000190483118215151615612c5c57612c5c6127ee565b500290565b600082612c7057612c70612c18565b50069056fea2646970667358221220680bb49499063d1f05e7666296daf42f599170836c96c96494066a1b9d0c737f64736f6c63430008090033",
}

// ChallengeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ChallengeManagerMetaData.ABI instead.
var ChallengeManagerABI = ChallengeManagerMetaData.ABI

// ChallengeManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChallengeManagerMetaData.Bin instead.
var ChallengeManagerBin = ChallengeManagerMetaData.Bin

// DeployChallengeManager deploys a new Ethereum contract, binding an instance of ChallengeManager to it.
func DeployChallengeManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeManager, error) {
	parsed, err := ChallengeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChallengeManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeManager{ChallengeManagerCaller: ChallengeManagerCaller{contract: contract}, ChallengeManagerTransactor: ChallengeManagerTransactor{contract: contract}, ChallengeManagerFilterer: ChallengeManagerFilterer{contract: contract}}, nil
}

// ChallengeManager is an auto generated Go binding around an Ethereum contract.
type ChallengeManager struct {
	ChallengeManagerCaller     // Read-only binding to the contract
	ChallengeManagerTransactor // Write-only binding to the contract
	ChallengeManagerFilterer   // Log filterer for contract events
}

// ChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeManagerSession struct {
	Contract     *ChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeManagerCallerSession struct {
	Contract *ChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeManagerTransactorSession struct {
	Contract     *ChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeManagerRaw struct {
	Contract *ChallengeManager // Generic contract binding to access the raw methods on
}

// ChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeManagerCallerRaw struct {
	Contract *ChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeManagerTransactorRaw struct {
	Contract *ChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeManager creates a new instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManager(address common.Address, backend bind.ContractBackend) (*ChallengeManager, error) {
	contract, err := bindChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeManager{ChallengeManagerCaller: ChallengeManagerCaller{contract: contract}, ChallengeManagerTransactor: ChallengeManagerTransactor{contract: contract}, ChallengeManagerFilterer: ChallengeManagerFilterer{contract: contract}}, nil
}

// NewChallengeManagerCaller creates a new read-only instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*ChallengeManagerCaller, error) {
	contract, err := bindChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerCaller{contract: contract}, nil
}

// NewChallengeManagerTransactor creates a new write-only instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeManagerTransactor, error) {
	contract, err := bindChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerTransactor{contract: contract}, nil
}

// NewChallengeManagerFilterer creates a new log filterer instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeManagerFilterer, error) {
	contract, err := bindChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerFilterer{contract: contract}, nil
}

// bindChallengeManager binds a generic wrapper to an already deployed contract.
func bindChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChallengeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManager *ChallengeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeManager.Contract.ChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManager *ChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManager *ChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManager *ChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManager *ChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManager *ChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ChallengeManager *ChallengeManagerCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ChallengeManager *ChallengeManagerSession) Bridge() (common.Address, error) {
	return _ChallengeManager.Contract.Bridge(&_ChallengeManager.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ChallengeManager *ChallengeManagerCallerSession) Bridge() (common.Address, error) {
	return _ChallengeManager.Contract.Bridge(&_ChallengeManager.CallOpts)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_ChallengeManager *ChallengeManagerCaller) ChallengeInfo(opts *bind.CallOpts, challengeIndex uint64) (ChallengeLibChallenge, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "challengeInfo", challengeIndex)

	if err != nil {
		return *new(ChallengeLibChallenge), err
	}

	out0 := *abi.ConvertType(out[0], new(ChallengeLibChallenge)).(*ChallengeLibChallenge)

	return out0, err

}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_ChallengeManager *ChallengeManagerSession) ChallengeInfo(challengeIndex uint64) (ChallengeLibChallenge, error) {
	return _ChallengeManager.Contract.ChallengeInfo(&_ChallengeManager.CallOpts, challengeIndex)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_ChallengeManager *ChallengeManagerCallerSession) ChallengeInfo(challengeIndex uint64) (ChallengeLibChallenge, error) {
	return _ChallengeManager.Contract.ChallengeInfo(&_ChallengeManager.CallOpts, challengeIndex)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_ChallengeManager *ChallengeManagerCaller) Challenges(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "challenges", arg0)

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
func (_ChallengeManager *ChallengeManagerSession) Challenges(arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	return _ChallengeManager.Contract.Challenges(&_ChallengeManager.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_ChallengeManager *ChallengeManagerCallerSession) Challenges(arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	return _ChallengeManager.Contract.Challenges(&_ChallengeManager.CallOpts, arg0)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_ChallengeManager *ChallengeManagerCaller) CurrentResponder(opts *bind.CallOpts, challengeIndex uint64) (common.Address, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "currentResponder", challengeIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_ChallengeManager *ChallengeManagerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _ChallengeManager.Contract.CurrentResponder(&_ChallengeManager.CallOpts, challengeIndex)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_ChallengeManager *ChallengeManagerCallerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _ChallengeManager.Contract.CurrentResponder(&_ChallengeManager.CallOpts, challengeIndex)
}

// GetOsp is a free data retrieval call binding the contract method 0x3690b011.
//
// Solidity: function getOsp(bytes32 wasmModuleRoot) view returns(address)
func (_ChallengeManager *ChallengeManagerCaller) GetOsp(opts *bind.CallOpts, wasmModuleRoot [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "getOsp", wasmModuleRoot)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOsp is a free data retrieval call binding the contract method 0x3690b011.
//
// Solidity: function getOsp(bytes32 wasmModuleRoot) view returns(address)
func (_ChallengeManager *ChallengeManagerSession) GetOsp(wasmModuleRoot [32]byte) (common.Address, error) {
	return _ChallengeManager.Contract.GetOsp(&_ChallengeManager.CallOpts, wasmModuleRoot)
}

// GetOsp is a free data retrieval call binding the contract method 0x3690b011.
//
// Solidity: function getOsp(bytes32 wasmModuleRoot) view returns(address)
func (_ChallengeManager *ChallengeManagerCallerSession) GetOsp(wasmModuleRoot [32]byte) (common.Address, error) {
	return _ChallengeManager.Contract.GetOsp(&_ChallengeManager.CallOpts, wasmModuleRoot)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_ChallengeManager *ChallengeManagerCaller) IsTimedOut(opts *bind.CallOpts, challengeIndex uint64) (bool, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "isTimedOut", challengeIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_ChallengeManager *ChallengeManagerSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _ChallengeManager.Contract.IsTimedOut(&_ChallengeManager.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_ChallengeManager *ChallengeManagerCallerSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _ChallengeManager.Contract.IsTimedOut(&_ChallengeManager.CallOpts, challengeIndex)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_ChallengeManager *ChallengeManagerCaller) Osp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "osp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_ChallengeManager *ChallengeManagerSession) Osp() (common.Address, error) {
	return _ChallengeManager.Contract.Osp(&_ChallengeManager.CallOpts)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_ChallengeManager *ChallengeManagerCallerSession) Osp() (common.Address, error) {
	return _ChallengeManager.Contract.Osp(&_ChallengeManager.CallOpts)
}

// OspCond is a free data retrieval call binding the contract method 0xdc74bf8b.
//
// Solidity: function ospCond(bytes32 ) view returns(address)
func (_ChallengeManager *ChallengeManagerCaller) OspCond(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "ospCond", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OspCond is a free data retrieval call binding the contract method 0xdc74bf8b.
//
// Solidity: function ospCond(bytes32 ) view returns(address)
func (_ChallengeManager *ChallengeManagerSession) OspCond(arg0 [32]byte) (common.Address, error) {
	return _ChallengeManager.Contract.OspCond(&_ChallengeManager.CallOpts, arg0)
}

// OspCond is a free data retrieval call binding the contract method 0xdc74bf8b.
//
// Solidity: function ospCond(bytes32 ) view returns(address)
func (_ChallengeManager *ChallengeManagerCallerSession) OspCond(arg0 [32]byte) (common.Address, error) {
	return _ChallengeManager.Contract.OspCond(&_ChallengeManager.CallOpts, arg0)
}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_ChallengeManager *ChallengeManagerCaller) ResultReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "resultReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_ChallengeManager *ChallengeManagerSession) ResultReceiver() (common.Address, error) {
	return _ChallengeManager.Contract.ResultReceiver(&_ChallengeManager.CallOpts)
}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_ChallengeManager *ChallengeManagerCallerSession) ResultReceiver() (common.Address, error) {
	return _ChallengeManager.Contract.ResultReceiver(&_ChallengeManager.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_ChallengeManager *ChallengeManagerCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_ChallengeManager *ChallengeManagerSession) SequencerInbox() (common.Address, error) {
	return _ChallengeManager.Contract.SequencerInbox(&_ChallengeManager.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_ChallengeManager *ChallengeManagerCallerSession) SequencerInbox() (common.Address, error) {
	return _ChallengeManager.Contract.SequencerInbox(&_ChallengeManager.CallOpts)
}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_ChallengeManager *ChallengeManagerCaller) TotalChallengesCreated(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ChallengeManager.contract.Call(opts, &out, "totalChallengesCreated")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_ChallengeManager *ChallengeManagerSession) TotalChallengesCreated() (uint64, error) {
	return _ChallengeManager.Contract.TotalChallengesCreated(&_ChallengeManager.CallOpts)
}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_ChallengeManager *ChallengeManagerCallerSession) TotalChallengesCreated() (uint64, error) {
	return _ChallengeManager.Contract.TotalChallengesCreated(&_ChallengeManager.CallOpts)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_ChallengeManager *ChallengeManagerTransactor) BisectExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "bisectExecution", challengeIndex, selection, newSegments)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_ChallengeManager *ChallengeManagerSession) BisectExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.BisectExecution(&_ChallengeManager.TransactOpts, challengeIndex, selection, newSegments)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) BisectExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.BisectExecution(&_ChallengeManager.TransactOpts, challengeIndex, selection, newSegments)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_ChallengeManager *ChallengeManagerTransactor) ChallengeExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "challengeExecution", challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_ChallengeManager *ChallengeManagerSession) ChallengeExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeExecution(&_ChallengeManager.TransactOpts, challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) ChallengeExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeExecution(&_ChallengeManager.TransactOpts, challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_ChallengeManager *ChallengeManagerTransactor) ClearChallenge(opts *bind.TransactOpts, challengeIndex uint64) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "clearChallenge", challengeIndex)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_ChallengeManager *ChallengeManagerSession) ClearChallenge(challengeIndex uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ClearChallenge(&_ChallengeManager.TransactOpts, challengeIndex)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) ClearChallenge(challengeIndex uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ClearChallenge(&_ChallengeManager.TransactOpts, challengeIndex)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_ChallengeManager *ChallengeManagerTransactor) CreateChallenge(opts *bind.TransactOpts, wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "createChallenge", wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_ChallengeManager *ChallengeManagerSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _ChallengeManager.Contract.CreateChallenge(&_ChallengeManager.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_ChallengeManager *ChallengeManagerTransactorSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _ChallengeManager.Contract.CreateChallenge(&_ChallengeManager.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_ChallengeManager *ChallengeManagerTransactor) Initialize(opts *bind.TransactOpts, resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "initialize", resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_ChallengeManager *ChallengeManagerSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.Initialize(&_ChallengeManager.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.Initialize(&_ChallengeManager.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_ChallengeManager *ChallengeManagerTransactor) OneStepProveExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "oneStepProveExecution", challengeIndex, selection, proof)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_ChallengeManager *ChallengeManagerSession) OneStepProveExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.OneStepProveExecution(&_ChallengeManager.TransactOpts, challengeIndex, selection, proof)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) OneStepProveExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.OneStepProveExecution(&_ChallengeManager.TransactOpts, challengeIndex, selection, proof)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x5038934d.
//
// Solidity: function postUpgradeInit(address osp_, bytes32 condRoot, address condOsp) returns()
func (_ChallengeManager *ChallengeManagerTransactor) PostUpgradeInit(opts *bind.TransactOpts, osp_ common.Address, condRoot [32]byte, condOsp common.Address) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "postUpgradeInit", osp_, condRoot, condOsp)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x5038934d.
//
// Solidity: function postUpgradeInit(address osp_, bytes32 condRoot, address condOsp) returns()
func (_ChallengeManager *ChallengeManagerSession) PostUpgradeInit(osp_ common.Address, condRoot [32]byte, condOsp common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.PostUpgradeInit(&_ChallengeManager.TransactOpts, osp_, condRoot, condOsp)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x5038934d.
//
// Solidity: function postUpgradeInit(address osp_, bytes32 condRoot, address condOsp) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) PostUpgradeInit(osp_ common.Address, condRoot [32]byte, condOsp common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.PostUpgradeInit(&_ChallengeManager.TransactOpts, osp_, condRoot, condOsp)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_ChallengeManager *ChallengeManagerTransactor) Timeout(opts *bind.TransactOpts, challengeIndex uint64) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "timeout", challengeIndex)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_ChallengeManager *ChallengeManagerSession) Timeout(challengeIndex uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.Timeout(&_ChallengeManager.TransactOpts, challengeIndex)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) Timeout(challengeIndex uint64) (*types.Transaction, error) {
	return _ChallengeManager.Contract.Timeout(&_ChallengeManager.TransactOpts, challengeIndex)
}

// ChallengeManagerBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the ChallengeManager contract.
type ChallengeManagerBisectedIterator struct {
	Event *ChallengeManagerBisected // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerBisected)
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
		it.Event = new(ChallengeManagerBisected)
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
func (it *ChallengeManagerBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerBisected represents a Bisected event raised by the ChallengeManager contract.
type ChallengeManagerBisected struct {
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
func (_ChallengeManager *ChallengeManagerFilterer) FilterBisected(opts *bind.FilterOpts, challengeIndex []uint64, challengeRoot [][32]byte) (*ChallengeManagerBisectedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerBisectedIterator{contract: _ChallengeManager.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_ChallengeManager *ChallengeManagerFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *ChallengeManagerBisected, challengeIndex []uint64, challengeRoot [][32]byte) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerBisected)
				if err := _ChallengeManager.contract.UnpackLog(event, "Bisected", log); err != nil {
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
func (_ChallengeManager *ChallengeManagerFilterer) ParseBisected(log types.Log) (*ChallengeManagerBisected, error) {
	event := new(ChallengeManagerBisected)
	if err := _ChallengeManager.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeManagerChallengeEndedIterator is returned from FilterChallengeEnded and is used to iterate over the raw logs and unpacked data for ChallengeEnded events raised by the ChallengeManager contract.
type ChallengeManagerChallengeEndedIterator struct {
	Event *ChallengeManagerChallengeEnded // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerChallengeEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerChallengeEnded)
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
		it.Event = new(ChallengeManagerChallengeEnded)
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
func (it *ChallengeManagerChallengeEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerChallengeEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerChallengeEnded represents a ChallengeEnded event raised by the ChallengeManager contract.
type ChallengeManagerChallengeEnded struct {
	ChallengeIndex uint64
	Kind           uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChallengeEnded is a free log retrieval operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_ChallengeManager *ChallengeManagerFilterer) FilterChallengeEnded(opts *bind.FilterOpts, challengeIndex []uint64) (*ChallengeManagerChallengeEndedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerChallengeEndedIterator{contract: _ChallengeManager.contract, event: "ChallengeEnded", logs: logs, sub: sub}, nil
}

// WatchChallengeEnded is a free log subscription operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_ChallengeManager *ChallengeManagerFilterer) WatchChallengeEnded(opts *bind.WatchOpts, sink chan<- *ChallengeManagerChallengeEnded, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerChallengeEnded)
				if err := _ChallengeManager.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
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
func (_ChallengeManager *ChallengeManagerFilterer) ParseChallengeEnded(log types.Log) (*ChallengeManagerChallengeEnded, error) {
	event := new(ChallengeManagerChallengeEnded)
	if err := _ChallengeManager.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeManagerExecutionChallengeBegunIterator is returned from FilterExecutionChallengeBegun and is used to iterate over the raw logs and unpacked data for ExecutionChallengeBegun events raised by the ChallengeManager contract.
type ChallengeManagerExecutionChallengeBegunIterator struct {
	Event *ChallengeManagerExecutionChallengeBegun // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerExecutionChallengeBegunIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerExecutionChallengeBegun)
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
		it.Event = new(ChallengeManagerExecutionChallengeBegun)
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
func (it *ChallengeManagerExecutionChallengeBegunIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerExecutionChallengeBegunIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerExecutionChallengeBegun represents a ExecutionChallengeBegun event raised by the ChallengeManager contract.
type ChallengeManagerExecutionChallengeBegun struct {
	ChallengeIndex uint64
	BlockSteps     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExecutionChallengeBegun is a free log retrieval operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_ChallengeManager *ChallengeManagerFilterer) FilterExecutionChallengeBegun(opts *bind.FilterOpts, challengeIndex []uint64) (*ChallengeManagerExecutionChallengeBegunIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerExecutionChallengeBegunIterator{contract: _ChallengeManager.contract, event: "ExecutionChallengeBegun", logs: logs, sub: sub}, nil
}

// WatchExecutionChallengeBegun is a free log subscription operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_ChallengeManager *ChallengeManagerFilterer) WatchExecutionChallengeBegun(opts *bind.WatchOpts, sink chan<- *ChallengeManagerExecutionChallengeBegun, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerExecutionChallengeBegun)
				if err := _ChallengeManager.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
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
func (_ChallengeManager *ChallengeManagerFilterer) ParseExecutionChallengeBegun(log types.Log) (*ChallengeManagerExecutionChallengeBegun, error) {
	event := new(ChallengeManagerExecutionChallengeBegun)
	if err := _ChallengeManager.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeManagerInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the ChallengeManager contract.
type ChallengeManagerInitiatedChallengeIterator struct {
	Event *ChallengeManagerInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerInitiatedChallenge)
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
		it.Event = new(ChallengeManagerInitiatedChallenge)
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
func (it *ChallengeManagerInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerInitiatedChallenge represents a InitiatedChallenge event raised by the ChallengeManager contract.
type ChallengeManagerInitiatedChallenge struct {
	ChallengeIndex uint64
	StartState     GlobalState
	EndState       GlobalState
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_ChallengeManager *ChallengeManagerFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts, challengeIndex []uint64) (*ChallengeManagerInitiatedChallengeIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerInitiatedChallengeIterator{contract: _ChallengeManager.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_ChallengeManager *ChallengeManagerFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeManagerInitiatedChallenge, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerInitiatedChallenge)
				if err := _ChallengeManager.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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
func (_ChallengeManager *ChallengeManagerFilterer) ParseInitiatedChallenge(log types.Log) (*ChallengeManagerInitiatedChallenge, error) {
	event := new(ChallengeManagerInitiatedChallenge)
	if err := _ChallengeManager.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeManagerOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the ChallengeManager contract.
type ChallengeManagerOneStepProofCompletedIterator struct {
	Event *ChallengeManagerOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerOneStepProofCompleted)
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
		it.Event = new(ChallengeManagerOneStepProofCompleted)
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
func (it *ChallengeManagerOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerOneStepProofCompleted represents a OneStepProofCompleted event raised by the ChallengeManager contract.
type ChallengeManagerOneStepProofCompleted struct {
	ChallengeIndex uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_ChallengeManager *ChallengeManagerFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts, challengeIndex []uint64) (*ChallengeManagerOneStepProofCompletedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerOneStepProofCompletedIterator{contract: _ChallengeManager.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_ChallengeManager *ChallengeManagerFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *ChallengeManagerOneStepProofCompleted, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerOneStepProofCompleted)
				if err := _ChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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
func (_ChallengeManager *ChallengeManagerFilterer) ParseOneStepProofCompleted(log types.Log) (*ChallengeManagerOneStepProofCompleted, error) {
	event := new(ChallengeManagerOneStepProofCompleted)
	if err := _ChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IChallengeManagerMetaData contains all meta data concerning the IChallengeManager contract.
var IChallengeManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"challengeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"chainHashes\",\"type\":\"bytes32[]\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIChallengeManager.ChallengeTerminationType\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"ChallengeEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockSteps\",\"type\":\"uint256\"}],\"name\":\"ExecutionChallengeBegun\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"startState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"endState\",\"type\":\"tuple\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex_\",\"type\":\"uint64\"}],\"name\":\"challengeInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"internalType\":\"structChallengeLib.Challenge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex_\",\"type\":\"uint64\"}],\"name\":\"clearChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"startAndEndMachineStatuses_\",\"type\":\"uint8[2]\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState[2]\",\"name\":\"startAndEndGlobalStates_\",\"type\":\"tuple[2]\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"asserter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"asserterTimeLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengerTimeLeft_\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"currentResponder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"}],\"name\":\"getOsp\",\"outputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChallengeResultReceiver\",\"name\":\"resultReceiver_\",\"type\":\"address\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"sequencerInbox_\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"osp_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"isTimedOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"osp\",\"outputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"osp_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"condRoot\",\"type\":\"bytes32\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"condOsp\",\"type\":\"address\"}],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex_\",\"type\":\"uint64\"}],\"name\":\"timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IChallengeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IChallengeManagerMetaData.ABI instead.
var IChallengeManagerABI = IChallengeManagerMetaData.ABI

// IChallengeManager is an auto generated Go binding around an Ethereum contract.
type IChallengeManager struct {
	IChallengeManagerCaller     // Read-only binding to the contract
	IChallengeManagerTransactor // Write-only binding to the contract
	IChallengeManagerFilterer   // Log filterer for contract events
}

// IChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeManagerSession struct {
	Contract     *IChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeManagerCallerSession struct {
	Contract *IChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeManagerTransactorSession struct {
	Contract     *IChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeManagerRaw struct {
	Contract *IChallengeManager // Generic contract binding to access the raw methods on
}

// IChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeManagerCallerRaw struct {
	Contract *IChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactorRaw struct {
	Contract *IChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeManager creates a new instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManager(address common.Address, backend bind.ContractBackend) (*IChallengeManager, error) {
	contract, err := bindIChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeManager{IChallengeManagerCaller: IChallengeManagerCaller{contract: contract}, IChallengeManagerTransactor: IChallengeManagerTransactor{contract: contract}, IChallengeManagerFilterer: IChallengeManagerFilterer{contract: contract}}, nil
}

// NewIChallengeManagerCaller creates a new read-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*IChallengeManagerCaller, error) {
	contract, err := bindIChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerCaller{contract: contract}, nil
}

// NewIChallengeManagerTransactor creates a new write-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeManagerTransactor, error) {
	contract, err := bindIChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerTransactor{contract: contract}, nil
}

// NewIChallengeManagerFilterer creates a new log filterer instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeManagerFilterer, error) {
	contract, err := bindIChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerFilterer{contract: contract}, nil
}

// bindIChallengeManager binds a generic wrapper to an already deployed contract.
func bindIChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IChallengeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.IChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex_) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_IChallengeManager *IChallengeManagerCaller) ChallengeInfo(opts *bind.CallOpts, challengeIndex_ uint64) (ChallengeLibChallenge, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "challengeInfo", challengeIndex_)

	if err != nil {
		return *new(ChallengeLibChallenge), err
	}

	out0 := *abi.ConvertType(out[0], new(ChallengeLibChallenge)).(*ChallengeLibChallenge)

	return out0, err

}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex_) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_IChallengeManager *IChallengeManagerSession) ChallengeInfo(challengeIndex_ uint64) (ChallengeLibChallenge, error) {
	return _IChallengeManager.Contract.ChallengeInfo(&_IChallengeManager.CallOpts, challengeIndex_)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex_) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_IChallengeManager *IChallengeManagerCallerSession) ChallengeInfo(challengeIndex_ uint64) (ChallengeLibChallenge, error) {
	return _IChallengeManager.Contract.ChallengeInfo(&_IChallengeManager.CallOpts, challengeIndex_)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_IChallengeManager *IChallengeManagerCaller) CurrentResponder(opts *bind.CallOpts, challengeIndex uint64) (common.Address, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "currentResponder", challengeIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_IChallengeManager *IChallengeManagerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _IChallengeManager.Contract.CurrentResponder(&_IChallengeManager.CallOpts, challengeIndex)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_IChallengeManager *IChallengeManagerCallerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _IChallengeManager.Contract.CurrentResponder(&_IChallengeManager.CallOpts, challengeIndex)
}

// GetOsp is a free data retrieval call binding the contract method 0x3690b011.
//
// Solidity: function getOsp(bytes32 wasmModuleRoot) view returns(address)
func (_IChallengeManager *IChallengeManagerCaller) GetOsp(opts *bind.CallOpts, wasmModuleRoot [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "getOsp", wasmModuleRoot)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOsp is a free data retrieval call binding the contract method 0x3690b011.
//
// Solidity: function getOsp(bytes32 wasmModuleRoot) view returns(address)
func (_IChallengeManager *IChallengeManagerSession) GetOsp(wasmModuleRoot [32]byte) (common.Address, error) {
	return _IChallengeManager.Contract.GetOsp(&_IChallengeManager.CallOpts, wasmModuleRoot)
}

// GetOsp is a free data retrieval call binding the contract method 0x3690b011.
//
// Solidity: function getOsp(bytes32 wasmModuleRoot) view returns(address)
func (_IChallengeManager *IChallengeManagerCallerSession) GetOsp(wasmModuleRoot [32]byte) (common.Address, error) {
	return _IChallengeManager.Contract.GetOsp(&_IChallengeManager.CallOpts, wasmModuleRoot)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_IChallengeManager *IChallengeManagerCaller) IsTimedOut(opts *bind.CallOpts, challengeIndex uint64) (bool, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "isTimedOut", challengeIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_IChallengeManager *IChallengeManagerSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _IChallengeManager.Contract.IsTimedOut(&_IChallengeManager.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_IChallengeManager *IChallengeManagerCallerSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _IChallengeManager.Contract.IsTimedOut(&_IChallengeManager.CallOpts, challengeIndex)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_IChallengeManager *IChallengeManagerCaller) Osp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IChallengeManager.contract.Call(opts, &out, "osp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_IChallengeManager *IChallengeManagerSession) Osp() (common.Address, error) {
	return _IChallengeManager.Contract.Osp(&_IChallengeManager.CallOpts)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_IChallengeManager *IChallengeManagerCallerSession) Osp() (common.Address, error) {
	return _IChallengeManager.Contract.Osp(&_IChallengeManager.CallOpts)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex_) returns()
func (_IChallengeManager *IChallengeManagerTransactor) ClearChallenge(opts *bind.TransactOpts, challengeIndex_ uint64) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "clearChallenge", challengeIndex_)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex_) returns()
func (_IChallengeManager *IChallengeManagerSession) ClearChallenge(challengeIndex_ uint64) (*types.Transaction, error) {
	return _IChallengeManager.Contract.ClearChallenge(&_IChallengeManager.TransactOpts, challengeIndex_)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex_) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) ClearChallenge(challengeIndex_ uint64) (*types.Transaction, error) {
	return _IChallengeManager.Contract.ClearChallenge(&_IChallengeManager.TransactOpts, challengeIndex_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_IChallengeManager *IChallengeManagerTransactor) CreateChallenge(opts *bind.TransactOpts, wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "createChallenge", wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_IChallengeManager *IChallengeManagerSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _IChallengeManager.Contract.CreateChallenge(&_IChallengeManager.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_IChallengeManager *IChallengeManagerTransactorSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _IChallengeManager.Contract.CreateChallenge(&_IChallengeManager.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_IChallengeManager *IChallengeManagerTransactor) Initialize(opts *bind.TransactOpts, resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "initialize", resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_IChallengeManager *IChallengeManagerSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Initialize(&_IChallengeManager.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Initialize(&_IChallengeManager.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x5038934d.
//
// Solidity: function postUpgradeInit(address osp_, bytes32 condRoot, address condOsp) returns()
func (_IChallengeManager *IChallengeManagerTransactor) PostUpgradeInit(opts *bind.TransactOpts, osp_ common.Address, condRoot [32]byte, condOsp common.Address) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "postUpgradeInit", osp_, condRoot, condOsp)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x5038934d.
//
// Solidity: function postUpgradeInit(address osp_, bytes32 condRoot, address condOsp) returns()
func (_IChallengeManager *IChallengeManagerSession) PostUpgradeInit(osp_ common.Address, condRoot [32]byte, condOsp common.Address) (*types.Transaction, error) {
	return _IChallengeManager.Contract.PostUpgradeInit(&_IChallengeManager.TransactOpts, osp_, condRoot, condOsp)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x5038934d.
//
// Solidity: function postUpgradeInit(address osp_, bytes32 condRoot, address condOsp) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) PostUpgradeInit(osp_ common.Address, condRoot [32]byte, condOsp common.Address) (*types.Transaction, error) {
	return _IChallengeManager.Contract.PostUpgradeInit(&_IChallengeManager.TransactOpts, osp_, condRoot, condOsp)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex_) returns()
func (_IChallengeManager *IChallengeManagerTransactor) Timeout(opts *bind.TransactOpts, challengeIndex_ uint64) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "timeout", challengeIndex_)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex_) returns()
func (_IChallengeManager *IChallengeManagerSession) Timeout(challengeIndex_ uint64) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Timeout(&_IChallengeManager.TransactOpts, challengeIndex_)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex_) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) Timeout(challengeIndex_ uint64) (*types.Transaction, error) {
	return _IChallengeManager.Contract.Timeout(&_IChallengeManager.TransactOpts, challengeIndex_)
}

// IChallengeManagerBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the IChallengeManager contract.
type IChallengeManagerBisectedIterator struct {
	Event *IChallengeManagerBisected // Event containing the contract specifics and raw log

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
func (it *IChallengeManagerBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IChallengeManagerBisected)
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
		it.Event = new(IChallengeManagerBisected)
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
func (it *IChallengeManagerBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IChallengeManagerBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IChallengeManagerBisected represents a Bisected event raised by the IChallengeManager contract.
type IChallengeManagerBisected struct {
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
func (_IChallengeManager *IChallengeManagerFilterer) FilterBisected(opts *bind.FilterOpts, challengeIndex []uint64, challengeRoot [][32]byte) (*IChallengeManagerBisectedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _IChallengeManager.contract.FilterLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerBisectedIterator{contract: _IChallengeManager.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_IChallengeManager *IChallengeManagerFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *IChallengeManagerBisected, challengeIndex []uint64, challengeRoot [][32]byte) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _IChallengeManager.contract.WatchLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IChallengeManagerBisected)
				if err := _IChallengeManager.contract.UnpackLog(event, "Bisected", log); err != nil {
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
func (_IChallengeManager *IChallengeManagerFilterer) ParseBisected(log types.Log) (*IChallengeManagerBisected, error) {
	event := new(IChallengeManagerBisected)
	if err := _IChallengeManager.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IChallengeManagerChallengeEndedIterator is returned from FilterChallengeEnded and is used to iterate over the raw logs and unpacked data for ChallengeEnded events raised by the IChallengeManager contract.
type IChallengeManagerChallengeEndedIterator struct {
	Event *IChallengeManagerChallengeEnded // Event containing the contract specifics and raw log

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
func (it *IChallengeManagerChallengeEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IChallengeManagerChallengeEnded)
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
		it.Event = new(IChallengeManagerChallengeEnded)
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
func (it *IChallengeManagerChallengeEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IChallengeManagerChallengeEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IChallengeManagerChallengeEnded represents a ChallengeEnded event raised by the IChallengeManager contract.
type IChallengeManagerChallengeEnded struct {
	ChallengeIndex uint64
	Kind           uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChallengeEnded is a free log retrieval operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_IChallengeManager *IChallengeManagerFilterer) FilterChallengeEnded(opts *bind.FilterOpts, challengeIndex []uint64) (*IChallengeManagerChallengeEndedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.FilterLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerChallengeEndedIterator{contract: _IChallengeManager.contract, event: "ChallengeEnded", logs: logs, sub: sub}, nil
}

// WatchChallengeEnded is a free log subscription operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_IChallengeManager *IChallengeManagerFilterer) WatchChallengeEnded(opts *bind.WatchOpts, sink chan<- *IChallengeManagerChallengeEnded, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.WatchLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IChallengeManagerChallengeEnded)
				if err := _IChallengeManager.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
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
func (_IChallengeManager *IChallengeManagerFilterer) ParseChallengeEnded(log types.Log) (*IChallengeManagerChallengeEnded, error) {
	event := new(IChallengeManagerChallengeEnded)
	if err := _IChallengeManager.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IChallengeManagerExecutionChallengeBegunIterator is returned from FilterExecutionChallengeBegun and is used to iterate over the raw logs and unpacked data for ExecutionChallengeBegun events raised by the IChallengeManager contract.
type IChallengeManagerExecutionChallengeBegunIterator struct {
	Event *IChallengeManagerExecutionChallengeBegun // Event containing the contract specifics and raw log

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
func (it *IChallengeManagerExecutionChallengeBegunIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IChallengeManagerExecutionChallengeBegun)
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
		it.Event = new(IChallengeManagerExecutionChallengeBegun)
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
func (it *IChallengeManagerExecutionChallengeBegunIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IChallengeManagerExecutionChallengeBegunIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IChallengeManagerExecutionChallengeBegun represents a ExecutionChallengeBegun event raised by the IChallengeManager contract.
type IChallengeManagerExecutionChallengeBegun struct {
	ChallengeIndex uint64
	BlockSteps     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExecutionChallengeBegun is a free log retrieval operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_IChallengeManager *IChallengeManagerFilterer) FilterExecutionChallengeBegun(opts *bind.FilterOpts, challengeIndex []uint64) (*IChallengeManagerExecutionChallengeBegunIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.FilterLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerExecutionChallengeBegunIterator{contract: _IChallengeManager.contract, event: "ExecutionChallengeBegun", logs: logs, sub: sub}, nil
}

// WatchExecutionChallengeBegun is a free log subscription operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_IChallengeManager *IChallengeManagerFilterer) WatchExecutionChallengeBegun(opts *bind.WatchOpts, sink chan<- *IChallengeManagerExecutionChallengeBegun, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.WatchLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IChallengeManagerExecutionChallengeBegun)
				if err := _IChallengeManager.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
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
func (_IChallengeManager *IChallengeManagerFilterer) ParseExecutionChallengeBegun(log types.Log) (*IChallengeManagerExecutionChallengeBegun, error) {
	event := new(IChallengeManagerExecutionChallengeBegun)
	if err := _IChallengeManager.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IChallengeManagerInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the IChallengeManager contract.
type IChallengeManagerInitiatedChallengeIterator struct {
	Event *IChallengeManagerInitiatedChallenge // Event containing the contract specifics and raw log

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
func (it *IChallengeManagerInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IChallengeManagerInitiatedChallenge)
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
		it.Event = new(IChallengeManagerInitiatedChallenge)
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
func (it *IChallengeManagerInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IChallengeManagerInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IChallengeManagerInitiatedChallenge represents a InitiatedChallenge event raised by the IChallengeManager contract.
type IChallengeManagerInitiatedChallenge struct {
	ChallengeIndex uint64
	StartState     GlobalState
	EndState       GlobalState
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_IChallengeManager *IChallengeManagerFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts, challengeIndex []uint64) (*IChallengeManagerInitiatedChallengeIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.FilterLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerInitiatedChallengeIterator{contract: _IChallengeManager.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_IChallengeManager *IChallengeManagerFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *IChallengeManagerInitiatedChallenge, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.WatchLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IChallengeManagerInitiatedChallenge)
				if err := _IChallengeManager.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
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
func (_IChallengeManager *IChallengeManagerFilterer) ParseInitiatedChallenge(log types.Log) (*IChallengeManagerInitiatedChallenge, error) {
	event := new(IChallengeManagerInitiatedChallenge)
	if err := _IChallengeManager.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IChallengeManagerOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the IChallengeManager contract.
type IChallengeManagerOneStepProofCompletedIterator struct {
	Event *IChallengeManagerOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *IChallengeManagerOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IChallengeManagerOneStepProofCompleted)
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
		it.Event = new(IChallengeManagerOneStepProofCompleted)
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
func (it *IChallengeManagerOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IChallengeManagerOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IChallengeManagerOneStepProofCompleted represents a OneStepProofCompleted event raised by the IChallengeManager contract.
type IChallengeManagerOneStepProofCompleted struct {
	ChallengeIndex uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_IChallengeManager *IChallengeManagerFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts, challengeIndex []uint64) (*IChallengeManagerOneStepProofCompletedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.FilterLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerOneStepProofCompletedIterator{contract: _IChallengeManager.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_IChallengeManager *IChallengeManagerFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *IChallengeManagerOneStepProofCompleted, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _IChallengeManager.contract.WatchLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IChallengeManagerOneStepProofCompleted)
				if err := _IChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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
func (_IChallengeManager *IChallengeManagerFilterer) ParseOneStepProofCompleted(log types.Log) (*IChallengeManagerOneStepProofCompleted, error) {
	event := new(IChallengeManagerOneStepProofCompleted)
	if err := _IChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IChallengeResultReceiverMetaData contains all meta data concerning the IChallengeResultReceiver contract.
var IChallengeResultReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengeIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IChallengeResultReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IChallengeResultReceiverMetaData.ABI instead.
var IChallengeResultReceiverABI = IChallengeResultReceiverMetaData.ABI

// IChallengeResultReceiver is an auto generated Go binding around an Ethereum contract.
type IChallengeResultReceiver struct {
	IChallengeResultReceiverCaller     // Read-only binding to the contract
	IChallengeResultReceiverTransactor // Write-only binding to the contract
	IChallengeResultReceiverFilterer   // Log filterer for contract events
}

// IChallengeResultReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeResultReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeResultReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeResultReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeResultReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeResultReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeResultReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeResultReceiverSession struct {
	Contract     *IChallengeResultReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IChallengeResultReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeResultReceiverCallerSession struct {
	Contract *IChallengeResultReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IChallengeResultReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeResultReceiverTransactorSession struct {
	Contract     *IChallengeResultReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IChallengeResultReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeResultReceiverRaw struct {
	Contract *IChallengeResultReceiver // Generic contract binding to access the raw methods on
}

// IChallengeResultReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeResultReceiverCallerRaw struct {
	Contract *IChallengeResultReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeResultReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeResultReceiverTransactorRaw struct {
	Contract *IChallengeResultReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeResultReceiver creates a new instance of IChallengeResultReceiver, bound to a specific deployed contract.
func NewIChallengeResultReceiver(address common.Address, backend bind.ContractBackend) (*IChallengeResultReceiver, error) {
	contract, err := bindIChallengeResultReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeResultReceiver{IChallengeResultReceiverCaller: IChallengeResultReceiverCaller{contract: contract}, IChallengeResultReceiverTransactor: IChallengeResultReceiverTransactor{contract: contract}, IChallengeResultReceiverFilterer: IChallengeResultReceiverFilterer{contract: contract}}, nil
}

// NewIChallengeResultReceiverCaller creates a new read-only instance of IChallengeResultReceiver, bound to a specific deployed contract.
func NewIChallengeResultReceiverCaller(address common.Address, caller bind.ContractCaller) (*IChallengeResultReceiverCaller, error) {
	contract, err := bindIChallengeResultReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeResultReceiverCaller{contract: contract}, nil
}

// NewIChallengeResultReceiverTransactor creates a new write-only instance of IChallengeResultReceiver, bound to a specific deployed contract.
func NewIChallengeResultReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeResultReceiverTransactor, error) {
	contract, err := bindIChallengeResultReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeResultReceiverTransactor{contract: contract}, nil
}

// NewIChallengeResultReceiverFilterer creates a new log filterer instance of IChallengeResultReceiver, bound to a specific deployed contract.
func NewIChallengeResultReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeResultReceiverFilterer, error) {
	contract, err := bindIChallengeResultReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeResultReceiverFilterer{contract: contract}, nil
}

// bindIChallengeResultReceiver binds a generic wrapper to an already deployed contract.
func bindIChallengeResultReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IChallengeResultReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeResultReceiver *IChallengeResultReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeResultReceiver.Contract.IChallengeResultReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeResultReceiver *IChallengeResultReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeResultReceiver.Contract.IChallengeResultReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeResultReceiver *IChallengeResultReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeResultReceiver.Contract.IChallengeResultReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeResultReceiver *IChallengeResultReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IChallengeResultReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeResultReceiver *IChallengeResultReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeResultReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeResultReceiver *IChallengeResultReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeResultReceiver.Contract.contract.Transact(opts, method, params...)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex, address winner, address loser) returns()
func (_IChallengeResultReceiver *IChallengeResultReceiverTransactor) CompleteChallenge(opts *bind.TransactOpts, challengeIndex *big.Int, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IChallengeResultReceiver.contract.Transact(opts, "completeChallenge", challengeIndex, winner, loser)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex, address winner, address loser) returns()
func (_IChallengeResultReceiver *IChallengeResultReceiverSession) CompleteChallenge(challengeIndex *big.Int, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IChallengeResultReceiver.Contract.CompleteChallenge(&_IChallengeResultReceiver.TransactOpts, challengeIndex, winner, loser)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex, address winner, address loser) returns()
func (_IChallengeResultReceiver *IChallengeResultReceiverTransactorSession) CompleteChallenge(challengeIndex *big.Int, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _IChallengeResultReceiver.Contract.CompleteChallenge(&_IChallengeResultReceiver.TransactOpts, challengeIndex, winner, loser)
}
