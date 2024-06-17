// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chaingen

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

// CacheManagerMetaData contains all meta data concerning the CacheManager contract.
var CacheManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initCacheSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initDecay\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"}],\"name\":\"AlreadyCached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"asm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"queueSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cacheSize\",\"type\":\"uint256\"}],\"name\":\"AsmTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"},{\"internalType\":\"uint192\",\"name\":\"min\",\"type\":\"uint192\"}],\"name\":\"BidTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BidsArePaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"}],\"name\":\"MakeSpaceTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotChainOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"DeleteBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"InsertBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"SetCacheSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"decay\",\"type\":\"uint64\"}],\"name\":\"SetDecayRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cacheSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"entries\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"code\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"evictAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"evictPrograms\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"makeSpace\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"space\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSize\",\"type\":\"uint64\"}],\"name\":\"setCacheSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDecay\",\"type\":\"uint64\"}],\"name\":\"setDecayRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sweepFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161146538038061146583398101604081905261002f91610097565b600280546001600160401b03928316600160801b027fffffffffffffffff0000000000000000ffffffffffffffff000000000000000090911692909316919091179190911790556100ca565b80516001600160401b038116811461009257600080fd5b919050565b600080604083850312156100aa57600080fd5b6100b38361007b565b91506100c16020840161007b565b90509250929050565b61138c806100d96000396000f3fe6080604052600436106100bd5760003560e01c8063a8d6fe041161006f578063a8d6fe041461019a578063b187bd26146101af578063b30906d4146101e0578063bae6c2ad1461021d578063c1c013c414610244578063c77ed13e14610257578063cadb43e21461027757600080fd5b80632dd4f566146100c25780633f4ba83a146100e4578063497ecfc5146100f957806354fac9191461010c5780635c32e943146101505780635c975abb14610165578063674a64e01461017a575b600080fd5b3480156100ce57600080fd5b506100e26100dd3660046111a0565b610297565b005b3480156100f057600080fd5b506100e2610389565b6100e26101073660046111c9565b610454565b34801561011857600080fd5b5060025461013390600160801b90046001600160401b031681565b6040516001600160401b0390911681526020015b60405180910390f35b34801561015c57600080fd5b506100e26104d8565b34801561017157600080fd5b506100e2610584565b34801561018657600080fd5b50600254610133906001600160401b031681565b3480156101a657600080fd5b506100e2610655565b3480156101bb57600080fd5b506002546101d090600160c01b900460ff1681565b6040519015158152602001610147565b3480156101ec57600080fd5b506102006101fb3660046111c9565b61072f565b604080519283526001600160401b03909116602083015201610147565b34801561022957600080fd5b5060025461013390600160401b90046001600160401b031681565b6101336102523660046111a0565b610766565b34801561026357600080fd5b506100e26102723660046111a0565b610805565b34801561028357600080fd5b506100e26102923660046111c9565b6108f3565b6040516304ddefed60e31b8152606b906326ef7f68906102bb9033906004016111e2565b60206040518083038186803b1580156102d357600080fd5b505afa1580156102e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030b91906111f6565b6103335733604051639531eff160e01b815260040161032a91906111e2565b60405180910390fd5b6002805467ffffffffffffffff19166001600160401b0383169081179091556040519081527fca22875e098f3b9c06ff3950c0cded621c968253a16623e890165451094c1839906020015b60405180910390a150565b6040516304ddefed60e31b8152606b906326ef7f68906103ad9033906004016111e2565b60206040518083038186803b1580156103c557600080fd5b505afa1580156103d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103fd91906111f6565b61041c5733604051639531eff160e01b815260040161032a91906111e2565b6002805460ff60c01b191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b600254600160c01b900460ff161561047f576040516323d5725b60e21b815260040160405180910390fd5b610488816109d9565b156104a95760405163c7e2d8e560e01b81526004810182905260240161032a565b60006104b482610a4c565b90506000806104c283610aea565b915091506104d282858584610bcc565b50505050565b6040516304ddefed60e31b8152606b906326ef7f68906104fc9033906004016111e2565b60206040518083038186803b15801561051457600080fd5b505afa158015610528573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061054c91906111f6565b61056b5733604051639531eff160e01b815260040161032a91906111e2565b6105766000196108f3565b61058260016000611158565b565b6040516304ddefed60e31b8152606b906326ef7f68906105a89033906004016111e2565b60206040518083038186803b1580156105c057600080fd5b505afa1580156105d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105f891906111f6565b6106175733604051639531eff160e01b815260040161032a91906111e2565b6002805460ff60c01b1916600160c01b1790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b600080606b6001600160a01b0316632d9125e96040518163ffffffff1660e01b815260040160206040518083038186803b15801561069257600080fd5b505afa1580156106a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ca9190611218565b6001600160a01b03164760405160006040518083038185875af1925050503d8060008114610714576040519150601f19603f3d011682016040523d82523d6000602084013e610719565b606091505b50915091508161072b57805160208201fd5b5050565b6001818154811061073f57600080fd5b6000918252602090912060029091020180546001909101549091506001600160401b031682565b600254600090600160c01b900460ff1615610794576040516323d5725b60e21b815260040160405180910390fd5b625000006001600160401b03831611156107d55760405163e6b801f360e01b81526001600160401b038316600482015262500000602482015260440161032a565b6107de82610aea565b50506002546107ff906001600160401b03600160401b820481169116611257565b92915050565b6040516304ddefed60e31b8152606b906326ef7f68906108299033906004016111e2565b60206040518083038186803b15801561084157600080fd5b505afa158015610855573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087991906111f6565b6108985733604051639531eff160e01b815260040161032a91906111e2565b6002805467ffffffffffffffff60801b1916600160801b6001600160401b038416908102919091179091556040519081527fd5ad38a519f54c97117f5a79fa7e82b03f32d2719f3ce4a27d4b561217cfea0c9060200161037e565b6040516304ddefed60e31b8152606b906326ef7f68906109179033906004016111e2565b60206040518083038186803b15801561092f57600080fd5b505afa158015610943573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096791906111f6565b6109865733604051639531eff160e01b815260040161032a91906111e2565b600054158015906109975750600081115b156109d6576000806109b46109ac6000610e4f565b604081901c91565b915091506109c28282610e66565b6109cd60018461127f565b92505050610986565b50565b60405163a72f179b60e01b81526004810182905260009060729063a72f179b9060240160206040518083038186803b158015610a1457600080fd5b505afa158015610a28573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ff91906111f6565b604051634089267f60e01b8152600481018290526000908190607190634089267f9060240160206040518083038186803b158015610a8957600080fd5b505afa158015610a9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac19190611296565b90506110008163ffffffff161015610adb57611000610add565b805b63ffffffff169392505050565b6002546000908190610b0c90600160801b90046001600160401b0316426112bc565b610b1690346112db565b60015460025491935091506000906001600160401b03165b6002546001600160401b0380831691610b50918891600160401b9004166112f3565b6001600160401b03161115610b7f57610b6c6109ac6000610e4f565b93509150610b7a8284610e66565b610b2e565b816001600160c01b0316846001600160c01b03161015610bc557604051631be6e1c960e31b81526001600160c01b0380861660048301528316602482015260440161032a565b5050915091565b6002546001600160401b0380821691610bee918591600160401b9004166112f3565b6001600160401b03161115610c3b5760025460405163bcc27c3760e01b81526001600160401b038085166004830152600160401b830481166024830152909116604482015260640161032a565b6040805180820182528481526001600160401b03841660208201529051634ceac81760e01b815260048101859052607290634ceac81790602401600060405180830381600087803b158015610c8f57600080fd5b505af1158015610ca3573d6000803e3d6000fd5b50610cce9250505067ffffffffffffffff19604087901b166001600160401b03841617600090610fe2565b82600260088282829054906101000a90046001600160401b0316610cf291906112f3565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550600180549050826001600160401b03161415610db35760018054808201825560009190915281517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf660029092029182015560208201517fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf7909101805467ffffffffffffffff19166001600160401b03909216919091179055610e0e565b806001836001600160401b031681548110610dd057610dd061131e565b60009182526020918290208351600292909202019081559101516001909101805467ffffffffffffffff19166001600160401b039092169190911790555b837fb5660cfc6c0afd838b3dc96244fa4b22b2e68203736b0ef80acef9e70cab30498685604051610e40929190611334565b60405180910390a25050505050565b6000610e5f826000806002610fec565b9392505050565b60006001826001600160401b031681548110610e8457610e8461131e565b600091825260209182902060408051808201825260029390930290910180548084526001909101546001600160401b0316938301939093525163ce97201360e01b81526004810192909252915060729063ce97201390602401600060405180830381600087803b158015610ef757600080fd5b505af1158015610f0b573d6000803e3d6000fd5b505050508060200151600260088282829054906101000a90046001600160401b0316610f379190611257565b92506101000a8154816001600160401b0302191690836001600160401b0316021790555080600001517f65905594d332f592fa6d4b86efc250c300a286b9d4f07f2ae89c3147dc4f39e7848360200151604051610f95929190611334565b60405180910390a26001826001600160401b031681548110610fb957610fb961131e565b600091825260208220600290910201908155600101805467ffffffffffffffff19169055505050565b6104d28282600060035b60008085548660005260206000206000801986611044578761100d576110c6565b87841461102657505060018281018955935081816110d4565b825489811061103557506110d4565b600396509450600190506110d4565b600287116110875783611056576110c6565b600287141561107b57600184039350838a558383015498508894508361107b576110d4565b508154935060016110d4565b600387141561109f57505060018201885581816110d4565b889450836110ac576110d4565b82548981106110bb57506110d4565b9450600190506110d4565b63a6ca772e6000526004601cfd5b5b8381101561110e57808301546001820180850154808310878310116110fa5750829050815b938501939093555050600181811b016110d5565b811561113c576001820360011c935083830154808a1061112e575061113c565b80838501555083915061110e565b600181011561114b5788828401555b5050505094509492505050565b50805460008255600202906000526020600020908101906109d691905b8082111561119c576000815560018101805467ffffffffffffffff19169055600201611175565b5090565b6000602082840312156111b257600080fd5b81356001600160401b0381168114610e5f57600080fd5b6000602082840312156111db57600080fd5b5035919050565b6001600160a01b0391909116815260200190565b60006020828403121561120857600080fd5b81518015158114610e5f57600080fd5b60006020828403121561122a57600080fd5b81516001600160a01b0381168114610e5f57600080fd5b634e487b7160e01b600052601160045260246000fd5b60006001600160401b038381169083168181101561127757611277611241565b039392505050565b60008282101561129157611291611241565b500390565b6000602082840312156112a857600080fd5b815163ffffffff81168114610e5f57600080fd5b60008160001904831182151516156112d6576112d6611241565b500290565b600082198211156112ee576112ee611241565b500190565b60006001600160401b0380831681851680830382111561131557611315611241565b01949350505050565b634e487b7160e01b600052603260045260246000fd5b6001600160c01b039290921682526001600160401b031660208201526040019056fea2646970667358221220a36636e0cdbf6146e01331e67841585dfe74280aaa17fa75a90a74c4556be69464736f6c63430008090033",
}

// CacheManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use CacheManagerMetaData.ABI instead.
var CacheManagerABI = CacheManagerMetaData.ABI

// CacheManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CacheManagerMetaData.Bin instead.
var CacheManagerBin = CacheManagerMetaData.Bin

// DeployCacheManager deploys a new Ethereum contract, binding an instance of CacheManager to it.
func DeployCacheManager(auth *bind.TransactOpts, backend bind.ContractBackend, initCacheSize uint64, initDecay uint64) (common.Address, *types.Transaction, *CacheManager, error) {
	parsed, err := CacheManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CacheManagerBin), backend, initCacheSize, initDecay)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CacheManager{CacheManagerCaller: CacheManagerCaller{contract: contract}, CacheManagerTransactor: CacheManagerTransactor{contract: contract}, CacheManagerFilterer: CacheManagerFilterer{contract: contract}}, nil
}

// CacheManager is an auto generated Go binding around an Ethereum contract.
type CacheManager struct {
	CacheManagerCaller     // Read-only binding to the contract
	CacheManagerTransactor // Write-only binding to the contract
	CacheManagerFilterer   // Log filterer for contract events
}

// CacheManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CacheManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CacheManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CacheManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CacheManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CacheManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CacheManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CacheManagerSession struct {
	Contract     *CacheManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CacheManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CacheManagerCallerSession struct {
	Contract *CacheManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CacheManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CacheManagerTransactorSession struct {
	Contract     *CacheManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CacheManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CacheManagerRaw struct {
	Contract *CacheManager // Generic contract binding to access the raw methods on
}

// CacheManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CacheManagerCallerRaw struct {
	Contract *CacheManagerCaller // Generic read-only contract binding to access the raw methods on
}

// CacheManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CacheManagerTransactorRaw struct {
	Contract *CacheManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCacheManager creates a new instance of CacheManager, bound to a specific deployed contract.
func NewCacheManager(address common.Address, backend bind.ContractBackend) (*CacheManager, error) {
	contract, err := bindCacheManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CacheManager{CacheManagerCaller: CacheManagerCaller{contract: contract}, CacheManagerTransactor: CacheManagerTransactor{contract: contract}, CacheManagerFilterer: CacheManagerFilterer{contract: contract}}, nil
}

// NewCacheManagerCaller creates a new read-only instance of CacheManager, bound to a specific deployed contract.
func NewCacheManagerCaller(address common.Address, caller bind.ContractCaller) (*CacheManagerCaller, error) {
	contract, err := bindCacheManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CacheManagerCaller{contract: contract}, nil
}

// NewCacheManagerTransactor creates a new write-only instance of CacheManager, bound to a specific deployed contract.
func NewCacheManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*CacheManagerTransactor, error) {
	contract, err := bindCacheManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CacheManagerTransactor{contract: contract}, nil
}

// NewCacheManagerFilterer creates a new log filterer instance of CacheManager, bound to a specific deployed contract.
func NewCacheManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*CacheManagerFilterer, error) {
	contract, err := bindCacheManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CacheManagerFilterer{contract: contract}, nil
}

// bindCacheManager binds a generic wrapper to an already deployed contract.
func bindCacheManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CacheManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CacheManager *CacheManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CacheManager.Contract.CacheManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CacheManager *CacheManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.Contract.CacheManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CacheManager *CacheManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CacheManager.Contract.CacheManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CacheManager *CacheManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CacheManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CacheManager *CacheManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CacheManager *CacheManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CacheManager.Contract.contract.Transact(opts, method, params...)
}

// CacheSize is a free data retrieval call binding the contract method 0x674a64e0.
//
// Solidity: function cacheSize() view returns(uint64)
func (_CacheManager *CacheManagerCaller) CacheSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "cacheSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CacheSize is a free data retrieval call binding the contract method 0x674a64e0.
//
// Solidity: function cacheSize() view returns(uint64)
func (_CacheManager *CacheManagerSession) CacheSize() (uint64, error) {
	return _CacheManager.Contract.CacheSize(&_CacheManager.CallOpts)
}

// CacheSize is a free data retrieval call binding the contract method 0x674a64e0.
//
// Solidity: function cacheSize() view returns(uint64)
func (_CacheManager *CacheManagerCallerSession) CacheSize() (uint64, error) {
	return _CacheManager.Contract.CacheSize(&_CacheManager.CallOpts)
}

// Decay is a free data retrieval call binding the contract method 0x54fac919.
//
// Solidity: function decay() view returns(uint64)
func (_CacheManager *CacheManagerCaller) Decay(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "decay")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Decay is a free data retrieval call binding the contract method 0x54fac919.
//
// Solidity: function decay() view returns(uint64)
func (_CacheManager *CacheManagerSession) Decay() (uint64, error) {
	return _CacheManager.Contract.Decay(&_CacheManager.CallOpts)
}

// Decay is a free data retrieval call binding the contract method 0x54fac919.
//
// Solidity: function decay() view returns(uint64)
func (_CacheManager *CacheManagerCallerSession) Decay() (uint64, error) {
	return _CacheManager.Contract.Decay(&_CacheManager.CallOpts)
}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) view returns(bytes32 code, uint64 size)
func (_CacheManager *CacheManagerCaller) Entries(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Code [32]byte
	Size uint64
}, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "entries", arg0)

	outstruct := new(struct {
		Code [32]byte
		Size uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Code = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Size = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) view returns(bytes32 code, uint64 size)
func (_CacheManager *CacheManagerSession) Entries(arg0 *big.Int) (struct {
	Code [32]byte
	Size uint64
}, error) {
	return _CacheManager.Contract.Entries(&_CacheManager.CallOpts, arg0)
}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) view returns(bytes32 code, uint64 size)
func (_CacheManager *CacheManagerCallerSession) Entries(arg0 *big.Int) (struct {
	Code [32]byte
	Size uint64
}, error) {
	return _CacheManager.Contract.Entries(&_CacheManager.CallOpts, arg0)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CacheManager *CacheManagerCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CacheManager *CacheManagerSession) IsPaused() (bool, error) {
	return _CacheManager.Contract.IsPaused(&_CacheManager.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CacheManager *CacheManagerCallerSession) IsPaused() (bool, error) {
	return _CacheManager.Contract.IsPaused(&_CacheManager.CallOpts)
}

// QueueSize is a free data retrieval call binding the contract method 0xbae6c2ad.
//
// Solidity: function queueSize() view returns(uint64)
func (_CacheManager *CacheManagerCaller) QueueSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "queueSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// QueueSize is a free data retrieval call binding the contract method 0xbae6c2ad.
//
// Solidity: function queueSize() view returns(uint64)
func (_CacheManager *CacheManagerSession) QueueSize() (uint64, error) {
	return _CacheManager.Contract.QueueSize(&_CacheManager.CallOpts)
}

// QueueSize is a free data retrieval call binding the contract method 0xbae6c2ad.
//
// Solidity: function queueSize() view returns(uint64)
func (_CacheManager *CacheManagerCallerSession) QueueSize() (uint64, error) {
	return _CacheManager.Contract.QueueSize(&_CacheManager.CallOpts)
}

// EvictAll is a paid mutator transaction binding the contract method 0x5c32e943.
//
// Solidity: function evictAll() returns()
func (_CacheManager *CacheManagerTransactor) EvictAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "evictAll")
}

// EvictAll is a paid mutator transaction binding the contract method 0x5c32e943.
//
// Solidity: function evictAll() returns()
func (_CacheManager *CacheManagerSession) EvictAll() (*types.Transaction, error) {
	return _CacheManager.Contract.EvictAll(&_CacheManager.TransactOpts)
}

// EvictAll is a paid mutator transaction binding the contract method 0x5c32e943.
//
// Solidity: function evictAll() returns()
func (_CacheManager *CacheManagerTransactorSession) EvictAll() (*types.Transaction, error) {
	return _CacheManager.Contract.EvictAll(&_CacheManager.TransactOpts)
}

// EvictPrograms is a paid mutator transaction binding the contract method 0xcadb43e2.
//
// Solidity: function evictPrograms(uint256 count) returns()
func (_CacheManager *CacheManagerTransactor) EvictPrograms(opts *bind.TransactOpts, count *big.Int) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "evictPrograms", count)
}

// EvictPrograms is a paid mutator transaction binding the contract method 0xcadb43e2.
//
// Solidity: function evictPrograms(uint256 count) returns()
func (_CacheManager *CacheManagerSession) EvictPrograms(count *big.Int) (*types.Transaction, error) {
	return _CacheManager.Contract.EvictPrograms(&_CacheManager.TransactOpts, count)
}

// EvictPrograms is a paid mutator transaction binding the contract method 0xcadb43e2.
//
// Solidity: function evictPrograms(uint256 count) returns()
func (_CacheManager *CacheManagerTransactorSession) EvictPrograms(count *big.Int) (*types.Transaction, error) {
	return _CacheManager.Contract.EvictPrograms(&_CacheManager.TransactOpts, count)
}

// MakeSpace is a paid mutator transaction binding the contract method 0xc1c013c4.
//
// Solidity: function makeSpace(uint64 size) payable returns(uint64 space)
func (_CacheManager *CacheManagerTransactor) MakeSpace(opts *bind.TransactOpts, size uint64) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "makeSpace", size)
}

// MakeSpace is a paid mutator transaction binding the contract method 0xc1c013c4.
//
// Solidity: function makeSpace(uint64 size) payable returns(uint64 space)
func (_CacheManager *CacheManagerSession) MakeSpace(size uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.MakeSpace(&_CacheManager.TransactOpts, size)
}

// MakeSpace is a paid mutator transaction binding the contract method 0xc1c013c4.
//
// Solidity: function makeSpace(uint64 size) payable returns(uint64 space)
func (_CacheManager *CacheManagerTransactorSession) MakeSpace(size uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.MakeSpace(&_CacheManager.TransactOpts, size)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns()
func (_CacheManager *CacheManagerTransactor) Paused(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "paused")
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns()
func (_CacheManager *CacheManagerSession) Paused() (*types.Transaction, error) {
	return _CacheManager.Contract.Paused(&_CacheManager.TransactOpts)
}

// Paused is a paid mutator transaction binding the contract method 0x5c975abb.
//
// Solidity: function paused() returns()
func (_CacheManager *CacheManagerTransactorSession) Paused() (*types.Transaction, error) {
	return _CacheManager.Contract.Paused(&_CacheManager.TransactOpts)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x497ecfc5.
//
// Solidity: function placeBid(bytes32 codehash) payable returns()
func (_CacheManager *CacheManagerTransactor) PlaceBid(opts *bind.TransactOpts, codehash [32]byte) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "placeBid", codehash)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x497ecfc5.
//
// Solidity: function placeBid(bytes32 codehash) payable returns()
func (_CacheManager *CacheManagerSession) PlaceBid(codehash [32]byte) (*types.Transaction, error) {
	return _CacheManager.Contract.PlaceBid(&_CacheManager.TransactOpts, codehash)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x497ecfc5.
//
// Solidity: function placeBid(bytes32 codehash) payable returns()
func (_CacheManager *CacheManagerTransactorSession) PlaceBid(codehash [32]byte) (*types.Transaction, error) {
	return _CacheManager.Contract.PlaceBid(&_CacheManager.TransactOpts, codehash)
}

// SetCacheSize is a paid mutator transaction binding the contract method 0x2dd4f566.
//
// Solidity: function setCacheSize(uint64 newSize) returns()
func (_CacheManager *CacheManagerTransactor) SetCacheSize(opts *bind.TransactOpts, newSize uint64) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "setCacheSize", newSize)
}

// SetCacheSize is a paid mutator transaction binding the contract method 0x2dd4f566.
//
// Solidity: function setCacheSize(uint64 newSize) returns()
func (_CacheManager *CacheManagerSession) SetCacheSize(newSize uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.SetCacheSize(&_CacheManager.TransactOpts, newSize)
}

// SetCacheSize is a paid mutator transaction binding the contract method 0x2dd4f566.
//
// Solidity: function setCacheSize(uint64 newSize) returns()
func (_CacheManager *CacheManagerTransactorSession) SetCacheSize(newSize uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.SetCacheSize(&_CacheManager.TransactOpts, newSize)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xc77ed13e.
//
// Solidity: function setDecayRate(uint64 newDecay) returns()
func (_CacheManager *CacheManagerTransactor) SetDecayRate(opts *bind.TransactOpts, newDecay uint64) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "setDecayRate", newDecay)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xc77ed13e.
//
// Solidity: function setDecayRate(uint64 newDecay) returns()
func (_CacheManager *CacheManagerSession) SetDecayRate(newDecay uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.SetDecayRate(&_CacheManager.TransactOpts, newDecay)
}

// SetDecayRate is a paid mutator transaction binding the contract method 0xc77ed13e.
//
// Solidity: function setDecayRate(uint64 newDecay) returns()
func (_CacheManager *CacheManagerTransactorSession) SetDecayRate(newDecay uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.SetDecayRate(&_CacheManager.TransactOpts, newDecay)
}

// SweepFunds is a paid mutator transaction binding the contract method 0xa8d6fe04.
//
// Solidity: function sweepFunds() returns()
func (_CacheManager *CacheManagerTransactor) SweepFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "sweepFunds")
}

// SweepFunds is a paid mutator transaction binding the contract method 0xa8d6fe04.
//
// Solidity: function sweepFunds() returns()
func (_CacheManager *CacheManagerSession) SweepFunds() (*types.Transaction, error) {
	return _CacheManager.Contract.SweepFunds(&_CacheManager.TransactOpts)
}

// SweepFunds is a paid mutator transaction binding the contract method 0xa8d6fe04.
//
// Solidity: function sweepFunds() returns()
func (_CacheManager *CacheManagerTransactorSession) SweepFunds() (*types.Transaction, error) {
	return _CacheManager.Contract.SweepFunds(&_CacheManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CacheManager *CacheManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CacheManager *CacheManagerSession) Unpause() (*types.Transaction, error) {
	return _CacheManager.Contract.Unpause(&_CacheManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CacheManager *CacheManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _CacheManager.Contract.Unpause(&_CacheManager.TransactOpts)
}

// CacheManagerDeleteBidIterator is returned from FilterDeleteBid and is used to iterate over the raw logs and unpacked data for DeleteBid events raised by the CacheManager contract.
type CacheManagerDeleteBidIterator struct {
	Event *CacheManagerDeleteBid // Event containing the contract specifics and raw log

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
func (it *CacheManagerDeleteBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerDeleteBid)
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
		it.Event = new(CacheManagerDeleteBid)
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
func (it *CacheManagerDeleteBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerDeleteBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerDeleteBid represents a DeleteBid event raised by the CacheManager contract.
type CacheManagerDeleteBid struct {
	Codehash [32]byte
	Bid      *big.Int
	Size     uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteBid is a free log retrieval operation binding the contract event 0x65905594d332f592fa6d4b86efc250c300a286b9d4f07f2ae89c3147dc4f39e7.
//
// Solidity: event DeleteBid(bytes32 indexed codehash, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) FilterDeleteBid(opts *bind.FilterOpts, codehash [][32]byte) (*CacheManagerDeleteBidIterator, error) {

	var codehashRule []interface{}
	for _, codehashItem := range codehash {
		codehashRule = append(codehashRule, codehashItem)
	}

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "DeleteBid", codehashRule)
	if err != nil {
		return nil, err
	}
	return &CacheManagerDeleteBidIterator{contract: _CacheManager.contract, event: "DeleteBid", logs: logs, sub: sub}, nil
}

// WatchDeleteBid is a free log subscription operation binding the contract event 0x65905594d332f592fa6d4b86efc250c300a286b9d4f07f2ae89c3147dc4f39e7.
//
// Solidity: event DeleteBid(bytes32 indexed codehash, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) WatchDeleteBid(opts *bind.WatchOpts, sink chan<- *CacheManagerDeleteBid, codehash [][32]byte) (event.Subscription, error) {

	var codehashRule []interface{}
	for _, codehashItem := range codehash {
		codehashRule = append(codehashRule, codehashItem)
	}

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "DeleteBid", codehashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerDeleteBid)
				if err := _CacheManager.contract.UnpackLog(event, "DeleteBid", log); err != nil {
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

// ParseDeleteBid is a log parse operation binding the contract event 0x65905594d332f592fa6d4b86efc250c300a286b9d4f07f2ae89c3147dc4f39e7.
//
// Solidity: event DeleteBid(bytes32 indexed codehash, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) ParseDeleteBid(log types.Log) (*CacheManagerDeleteBid, error) {
	event := new(CacheManagerDeleteBid)
	if err := _CacheManager.contract.UnpackLog(event, "DeleteBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerInsertBidIterator is returned from FilterInsertBid and is used to iterate over the raw logs and unpacked data for InsertBid events raised by the CacheManager contract.
type CacheManagerInsertBidIterator struct {
	Event *CacheManagerInsertBid // Event containing the contract specifics and raw log

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
func (it *CacheManagerInsertBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerInsertBid)
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
		it.Event = new(CacheManagerInsertBid)
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
func (it *CacheManagerInsertBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerInsertBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerInsertBid represents a InsertBid event raised by the CacheManager contract.
type CacheManagerInsertBid struct {
	Codehash [32]byte
	Bid      *big.Int
	Size     uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInsertBid is a free log retrieval operation binding the contract event 0xb5660cfc6c0afd838b3dc96244fa4b22b2e68203736b0ef80acef9e70cab3049.
//
// Solidity: event InsertBid(bytes32 indexed codehash, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) FilterInsertBid(opts *bind.FilterOpts, codehash [][32]byte) (*CacheManagerInsertBidIterator, error) {

	var codehashRule []interface{}
	for _, codehashItem := range codehash {
		codehashRule = append(codehashRule, codehashItem)
	}

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "InsertBid", codehashRule)
	if err != nil {
		return nil, err
	}
	return &CacheManagerInsertBidIterator{contract: _CacheManager.contract, event: "InsertBid", logs: logs, sub: sub}, nil
}

// WatchInsertBid is a free log subscription operation binding the contract event 0xb5660cfc6c0afd838b3dc96244fa4b22b2e68203736b0ef80acef9e70cab3049.
//
// Solidity: event InsertBid(bytes32 indexed codehash, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) WatchInsertBid(opts *bind.WatchOpts, sink chan<- *CacheManagerInsertBid, codehash [][32]byte) (event.Subscription, error) {

	var codehashRule []interface{}
	for _, codehashItem := range codehash {
		codehashRule = append(codehashRule, codehashItem)
	}

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "InsertBid", codehashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerInsertBid)
				if err := _CacheManager.contract.UnpackLog(event, "InsertBid", log); err != nil {
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

// ParseInsertBid is a log parse operation binding the contract event 0xb5660cfc6c0afd838b3dc96244fa4b22b2e68203736b0ef80acef9e70cab3049.
//
// Solidity: event InsertBid(bytes32 indexed codehash, uint192 bid, uint64 size)
func (_CacheManager *CacheManagerFilterer) ParseInsertBid(log types.Log) (*CacheManagerInsertBid, error) {
	event := new(CacheManagerInsertBid)
	if err := _CacheManager.contract.UnpackLog(event, "InsertBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the CacheManager contract.
type CacheManagerPauseIterator struct {
	Event *CacheManagerPause // Event containing the contract specifics and raw log

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
func (it *CacheManagerPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerPause)
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
		it.Event = new(CacheManagerPause)
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
func (it *CacheManagerPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerPause represents a Pause event raised by the CacheManager contract.
type CacheManagerPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CacheManager *CacheManagerFilterer) FilterPause(opts *bind.FilterOpts) (*CacheManagerPauseIterator, error) {

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &CacheManagerPauseIterator{contract: _CacheManager.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CacheManager *CacheManagerFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *CacheManagerPause) (event.Subscription, error) {

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerPause)
				if err := _CacheManager.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_CacheManager *CacheManagerFilterer) ParsePause(log types.Log) (*CacheManagerPause, error) {
	event := new(CacheManagerPause)
	if err := _CacheManager.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerSetCacheSizeIterator is returned from FilterSetCacheSize and is used to iterate over the raw logs and unpacked data for SetCacheSize events raised by the CacheManager contract.
type CacheManagerSetCacheSizeIterator struct {
	Event *CacheManagerSetCacheSize // Event containing the contract specifics and raw log

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
func (it *CacheManagerSetCacheSizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerSetCacheSize)
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
		it.Event = new(CacheManagerSetCacheSize)
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
func (it *CacheManagerSetCacheSizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerSetCacheSizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerSetCacheSize represents a SetCacheSize event raised by the CacheManager contract.
type CacheManagerSetCacheSize struct {
	Size uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetCacheSize is a free log retrieval operation binding the contract event 0xca22875e098f3b9c06ff3950c0cded621c968253a16623e890165451094c1839.
//
// Solidity: event SetCacheSize(uint64 size)
func (_CacheManager *CacheManagerFilterer) FilterSetCacheSize(opts *bind.FilterOpts) (*CacheManagerSetCacheSizeIterator, error) {

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "SetCacheSize")
	if err != nil {
		return nil, err
	}
	return &CacheManagerSetCacheSizeIterator{contract: _CacheManager.contract, event: "SetCacheSize", logs: logs, sub: sub}, nil
}

// WatchSetCacheSize is a free log subscription operation binding the contract event 0xca22875e098f3b9c06ff3950c0cded621c968253a16623e890165451094c1839.
//
// Solidity: event SetCacheSize(uint64 size)
func (_CacheManager *CacheManagerFilterer) WatchSetCacheSize(opts *bind.WatchOpts, sink chan<- *CacheManagerSetCacheSize) (event.Subscription, error) {

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "SetCacheSize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerSetCacheSize)
				if err := _CacheManager.contract.UnpackLog(event, "SetCacheSize", log); err != nil {
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

// ParseSetCacheSize is a log parse operation binding the contract event 0xca22875e098f3b9c06ff3950c0cded621c968253a16623e890165451094c1839.
//
// Solidity: event SetCacheSize(uint64 size)
func (_CacheManager *CacheManagerFilterer) ParseSetCacheSize(log types.Log) (*CacheManagerSetCacheSize, error) {
	event := new(CacheManagerSetCacheSize)
	if err := _CacheManager.contract.UnpackLog(event, "SetCacheSize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerSetDecayRateIterator is returned from FilterSetDecayRate and is used to iterate over the raw logs and unpacked data for SetDecayRate events raised by the CacheManager contract.
type CacheManagerSetDecayRateIterator struct {
	Event *CacheManagerSetDecayRate // Event containing the contract specifics and raw log

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
func (it *CacheManagerSetDecayRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerSetDecayRate)
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
		it.Event = new(CacheManagerSetDecayRate)
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
func (it *CacheManagerSetDecayRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerSetDecayRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerSetDecayRate represents a SetDecayRate event raised by the CacheManager contract.
type CacheManagerSetDecayRate struct {
	Decay uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetDecayRate is a free log retrieval operation binding the contract event 0xd5ad38a519f54c97117f5a79fa7e82b03f32d2719f3ce4a27d4b561217cfea0c.
//
// Solidity: event SetDecayRate(uint64 decay)
func (_CacheManager *CacheManagerFilterer) FilterSetDecayRate(opts *bind.FilterOpts) (*CacheManagerSetDecayRateIterator, error) {

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "SetDecayRate")
	if err != nil {
		return nil, err
	}
	return &CacheManagerSetDecayRateIterator{contract: _CacheManager.contract, event: "SetDecayRate", logs: logs, sub: sub}, nil
}

// WatchSetDecayRate is a free log subscription operation binding the contract event 0xd5ad38a519f54c97117f5a79fa7e82b03f32d2719f3ce4a27d4b561217cfea0c.
//
// Solidity: event SetDecayRate(uint64 decay)
func (_CacheManager *CacheManagerFilterer) WatchSetDecayRate(opts *bind.WatchOpts, sink chan<- *CacheManagerSetDecayRate) (event.Subscription, error) {

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "SetDecayRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerSetDecayRate)
				if err := _CacheManager.contract.UnpackLog(event, "SetDecayRate", log); err != nil {
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

// ParseSetDecayRate is a log parse operation binding the contract event 0xd5ad38a519f54c97117f5a79fa7e82b03f32d2719f3ce4a27d4b561217cfea0c.
//
// Solidity: event SetDecayRate(uint64 decay)
func (_CacheManager *CacheManagerFilterer) ParseSetDecayRate(log types.Log) (*CacheManagerSetDecayRate, error) {
	event := new(CacheManagerSetDecayRate)
	if err := _CacheManager.contract.UnpackLog(event, "SetDecayRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CacheManagerUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the CacheManager contract.
type CacheManagerUnpauseIterator struct {
	Event *CacheManagerUnpause // Event containing the contract specifics and raw log

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
func (it *CacheManagerUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CacheManagerUnpause)
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
		it.Event = new(CacheManagerUnpause)
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
func (it *CacheManagerUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CacheManagerUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CacheManagerUnpause represents a Unpause event raised by the CacheManager contract.
type CacheManagerUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CacheManager *CacheManagerFilterer) FilterUnpause(opts *bind.FilterOpts) (*CacheManagerUnpauseIterator, error) {

	logs, sub, err := _CacheManager.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &CacheManagerUnpauseIterator{contract: _CacheManager.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CacheManager *CacheManagerFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *CacheManagerUnpause) (event.Subscription, error) {

	logs, sub, err := _CacheManager.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CacheManagerUnpause)
				if err := _CacheManager.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_CacheManager *CacheManagerFilterer) ParseUnpause(log types.Log) (*CacheManagerUnpause, error) {
	event := new(CacheManagerUnpause)
	if err := _CacheManager.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
