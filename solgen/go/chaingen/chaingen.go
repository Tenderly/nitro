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

// CacheManagerEntry is an auto generated low-level Go binding around an user-defined struct.
type CacheManagerEntry struct {
	Code [32]byte
	Size uint64
	Bid  *big.Int
}

// CacheManagerMetaData contains all meta data concerning the CacheManager contract.
var CacheManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"}],\"name\":\"AlreadyCached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"asm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"queueSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cacheSize\",\"type\":\"uint256\"}],\"name\":\"AsmTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"},{\"internalType\":\"uint192\",\"name\":\"min\",\"type\":\"uint192\"}],\"name\":\"BidTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BidsArePaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"}],\"name\":\"MakeSpaceTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotChainOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"DeleteBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"InsertBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"SetCacheSize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"decay\",\"type\":\"uint64\"}],\"name\":\"SetDecayRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cacheSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"entries\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"code\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"evictAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"evictPrograms\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEntries\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"code\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"}],\"internalType\":\"structCacheManager.Entry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"}],\"name\":\"getMinBid\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"min\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"}],\"name\":\"getMinBid\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"min\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"getMinBid\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"min\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"}],\"name\":\"getSmallestEntries\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"code\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"bid\",\"type\":\"uint192\"}],\"internalType\":\"structCacheManager.Entry[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"initCacheSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initDecay\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"size\",\"type\":\"uint64\"}],\"name\":\"makeSpace\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"space\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newSize\",\"type\":\"uint64\"}],\"name\":\"setCacheSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDecay\",\"type\":\"uint64\"}],\"name\":\"setDecayRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sweepFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b50608051611dda61003060003960006105270152611dda6000f3fe60806040526004361061011f5760003560e01c8063b187bd26116100a0578063c77ed13e11610064578063c77ed13e14610353578063cadb43e214610373578063d29b303e14610393578063e4940157146103b3578063e9c1bc0f146103c657600080fd5b8063b187bd261461027c578063b30906d4146102ad578063bae6c2ad146102f9578063c1c013c414610320578063c565a2081461033357600080fd5b806354fac919116100e757806354fac919146101de5780635c32e9431461021d5780635c975abb14610232578063674a64e014610247578063a8d6fe041461026757600080fd5b806317be85c31461012457806320f2f3451461014f5780632dd4f5661461017157806332052a9b146101915780633f4ba83a146101c9575b600080fd5b34801561013057600080fd5b506101396103e6565b6040516101469190611acb565b60405180910390f35b34801561015b57600080fd5b5061016f61016a366004611b52565b610473565b005b34801561017d57600080fd5b5061016f61018c366004611b85565b6105fa565b34801561019d57600080fd5b506101b16101ac366004611bb5565b6106e3565b6040516001600160c01b039091168152602001610146565b3480156101d557600080fd5b5061016f6106fe565b3480156101ea57600080fd5b5060035461020590600160801b90046001600160401b031681565b6040516001600160401b039091168152602001610146565b34801561022957600080fd5b5061016f6107c9565b34801561023e57600080fd5b5061016f610875565b34801561025357600080fd5b50600354610205906001600160401b031681565b34801561027357600080fd5b5061016f610946565b34801561028857600080fd5b5060035461029d90600160c01b900460ff1681565b6040519015158152602001610146565b3480156102b957600080fd5b506102cd6102c8366004611bd2565b610a20565b604080519384526001600160401b0390921660208401526001600160c01b031690820152606001610146565b34801561030557600080fd5b5060035461020590600160401b90046001600160401b031681565b61020561032e366004611b85565b610a69565b34801561033f57600080fd5b506101b161034e366004611bd2565b610b02565b34801561035f57600080fd5b5061016f61036e366004611b85565b610b10565b34801561037f57600080fd5b5061016f61038e366004611bd2565b610bfe565b34801561039f57600080fd5b506101b16103ae366004611b85565b610ce4565b61016f6103c1366004611bb5565b610ec9565b3480156103d257600080fd5b506101396103e1366004611bd2565b610f5a565b60606002805480602002602001604051908101604052809291908181526020016000905b8282101561046a5760008481526020908190206040805160608101825260028602909201805483526001908101546001600160401b03811684860152600160401b90046001600160c01b031691830191909152908352909201910161040a565b50505050905090565b600054610100900460ff1661048e5760005460ff1615610492565b303b155b6104fa5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b600054610100900460ff1615801561051c576000805461ffff19166101011790555b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156105aa5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b60648201526084016104f1565b600380546001600160401b03848116600160801b0277ffffffffffffffff0000000000000000ffffffffffffffff199092169086161717905580156105f5576000805461ff00191690555b505050565b6040516304ddefed60e31b8152606b906326ef7f689061061e903390600401611beb565b60206040518083038186803b15801561063657600080fd5b505afa15801561064a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066e9190611bff565b61068d5733604051639531eff160e01b81526004016104f19190611beb565b6003805467ffffffffffffffff19166001600160401b0383169081179091556040519081527fca22875e098f3b9c06ff3950c0cded621c968253a16623e890165451094c1839906020015b60405180910390a150565b60006106f8826001600160a01b03163f610b02565b92915050565b6040516304ddefed60e31b8152606b906326ef7f6890610722903390600401611beb565b60206040518083038186803b15801561073a57600080fd5b505afa15801561074e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107729190611bff565b6107915733604051639531eff160e01b81526004016104f19190611beb565b6003805460ff60c01b191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b6040516304ddefed60e31b8152606b906326ef7f68906107ed903390600401611beb565b60206040518083038186803b15801561080557600080fd5b505afa158015610819573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083d9190611bff565b61085c5733604051639531eff160e01b81526004016104f19190611beb565b610867600019610bfe565b61087360026000611a90565b565b6040516304ddefed60e31b8152606b906326ef7f6890610899903390600401611beb565b60206040518083038186803b1580156108b157600080fd5b505afa1580156108c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e99190611bff565b6109085733604051639531eff160e01b81526004016104f19190611beb565b6003805460ff60c01b1916600160c01b1790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b600080606b6001600160a01b0316632d9125e96040518163ffffffff1660e01b815260040160206040518083038186803b15801561098357600080fd5b505afa158015610997573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109bb9190611c21565b6001600160a01b03164760405160006040518083038185875af1925050503d8060008114610a05576040519150601f19603f3d011682016040523d82523d6000602084013e610a0a565b606091505b509150915081610a1c57805160208201fd5b5050565b60028181548110610a3057600080fd5b6000918252602090912060029091020180546001909101549091506001600160401b03811690600160401b90046001600160c01b031683565b600354600090600160c01b900460ff1615610a97576040516323d5725b60e21b815260040160405180910390fd5b625000006001600160401b0383161115610ad85760405163e6b801f360e01b81526001600160401b03831660048201526250000060248201526044016104f1565b610ae1826110c3565b50506003546106f8906001600160401b03600160401b820481169116611c54565b60006106f86103ae83611185565b6040516304ddefed60e31b8152606b906326ef7f6890610b34903390600401611beb565b60206040518083038186803b158015610b4c57600080fd5b505afa158015610b60573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b849190611bff565b610ba35733604051639531eff160e01b81526004016104f19190611beb565b6003805467ffffffffffffffff60801b1916600160801b6001600160401b038416908102919091179091556040519081527fd5ad38a519f54c97117f5a79fa7e82b03f32d2719f3ce4a27d4b561217cfea0c906020016106d8565b6040516304ddefed60e31b8152606b906326ef7f6890610c22903390600401611beb565b60206040518083038186803b158015610c3a57600080fd5b505afa158015610c4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c729190611bff565b610c915733604051639531eff160e01b81526004016104f19190611beb565b60015415801590610ca25750600081115b15610ce157600080610cbf610cb76001611223565b604081901c91565b91509150610ccd8282611233565b610cd8600184611c7c565b92505050610c91565b50565b6003546000906001600160401b039081169083161115610d355760035460405163bcc27c3760e01b81526001600160401b0380851660048301526000602483015290911660448201526064016104f1565b6110006001600160401b0383161015610d5057611000610d52565b815b600354909250600090610d76908490600160401b90046001600160401b0316611c93565b6003546001600160401b039182169250168111610d965750600092915050565b600354600090610daf906001600160401b031683611c7c565b905060006110006001610dc28285611cbe565b610dcc9190611c7c565b610dd69190611cd6565b90506000610de382610f5a565b905060005b8151811015610e8957818181518110610e0357610e03611cf8565b6020026020010151602001516001600160401b03168411610e4357818181518110610e3057610e30611cf8565b6020026020010151604001519550610e89565b818181518110610e5557610e55611cf8565b6020026020010151602001516001600160401b031684610e759190611c7c565b935080610e8181611d0e565b915050610de8565b506000610e946113d5565b905080866001600160c01b03161015610eb4575060009695505050505050565b610ebe8187611d29565b979650505050505050565b600354600160c01b900460ff1615610ef4576040516323d5725b60e21b815260040160405180910390fd5b6001600160a01b0381163f610f08816113fa565b15610f295760405163c7e2d8e560e01b8152600481018290526024016104f1565b6000610f3482611185565b9050600080610f42836110c3565b91509150610f53828686868561146d565b5050505050565b606081610f6660015490565b1015610f725760015491505b6000610f7f60018461172f565b905080516001600160401b03811115610f9a57610f9a611d49565b604051908082528060200260200182016040528015610fe557816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610fb85790505b50915060005b81518110156110bc57600061101b83838151811061100b5761100b611cf8565b6020026020010151604081901c91565b9150506002816001600160401b03168154811061103a5761103a611cf8565b60009182526020918290206040805160608101825260029390930290910180548352600101546001600160401b03811693830193909352600160401b9092046001600160c01b031691810191909152845185908490811061109d5761109d611cf8565b60200260200101819052505080806110b490611d0e565b915050610feb565b5050919050565b6000806110cf346118d4565b60025460035491935091506000906001600160401b03165b6003546001600160401b0380831691611109918891600160401b900416611c93565b6001600160401b0316111561113857611125610cb76001611223565b935091506111338284611233565b6110e7565b816001600160c01b0316846001600160c01b0316101561117e57604051631be6e1c960e31b81526001600160c01b038086166004830152831660248201526044016104f1565b5050915091565b604051634089267f60e01b8152600481018290526000908190607190634089267f9060240160206040518083038186803b1580156111c257600080fd5b505afa1580156111d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111fa9190611d5f565b905061100063ffffffff821610156112145761100061121c565b8063ffffffff165b9392505050565b600061121c826000806002611916565b60006002826001600160401b03168154811061125157611251611cf8565b60009182526020918290206040805160608101825260029390930290910180548084526001909101546001600160401b03811694840194909452600160401b9093046001600160c01b0316828201525163ce97201360e01b81526004810192909252915060729063ce97201390602401600060405180830381600087803b1580156112db57600080fd5b505af11580156112ef573d6000803e3d6000fd5b505050508060200151600360088282829054906101000a90046001600160401b031661131b9190611c54565b92506101000a8154816001600160401b0302191690836001600160401b0316021790555080600001517f65905594d332f592fa6d4b86efc250c300a286b9d4f07f2ae89c3147dc4f39e78483602001516040516113969291906001600160c01b039290921682526001600160401b0316602082015260400190565b60405180910390a26002826001600160401b0316815481106113ba576113ba611cf8565b60009182526020822060029091020181815560010155505050565b6003546000906113f590600160801b90046001600160401b031642611d85565b905090565b60405163a72f179b60e01b81526004810182905260009060729063a72f179b9060240160206040518083038186803b15801561143557600080fd5b505afa158015611449573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f89190611bff565b6003546001600160401b038082169161148f918591600160401b900416611c93565b6001600160401b031611156114dc5760035460405163bcc27c3760e01b81526001600160401b038085166004830152600160401b83048116602483015290911660448201526064016104f1565b604080516060810182528481526001600160401b03841660208201526001600160c01b03871681830152905163739d64f960e11b815260729063e73ac9f290611529908890600401611beb565b600060405180830381600087803b15801561154357600080fd5b505af1158015611557573d6000803e3d6000fd5b506115829250505067ffffffffffffffff19604088901b166001600160401b03841617600190611a82565b82600360088282829054906101000a90046001600160401b03166115a69190611c93565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550600280549050826001600160401b0316141561166957600280546001810182556000829052825191027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace810191909155602082015160408301516001600160c01b0316600160401b026001600160401b03909116177f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf909101556116cb565b806002836001600160401b03168154811061168657611686611cf8565b60009182526020918290208351600290920201908155908201516040909201516001600160c01b0316600160401b026001600160401b03909216919091176001909101555b604080516001600160a01b03871681526001600160c01b03881660208201526001600160401b03851681830152905185917fb9271ce6a232cb5e0010e10fc10b38fe5d25dd27f8c03beef068a581cfc21bec919081900360600190a2505050505050565b60606117ef565b6001820360011c61174d818360069190911b015190565b8510830261175b575061177c565b600681811b830180516020918201519286901b850190815201529150611736565b600682901b81018481526020018390525b50505050565b600060015b838110156117e2576001810184118101600690811b84015182821b8501511190910180821b840180516020918201519385901b860190815201919091529050600181811b01611798565b50610f5385858385611736565b5060408051600084815260208082208654815487821882891002821860051b86018085019182529096019390935292939084019190808314155b80156118bb578151845283602001935081841415611846576118bb565b602082015160019060011b0183811061188a57600182039150611884611872838560069190911b015190565b600684901b8501602001518486611793565b50611829565b61189981870154828486611793565b600101838114611884576118b281870154828486611736565b50600101611829565b505050601f198382030160051c83526040525092915050565b6000806118df6113d5565b6118e99084611cbe565b90506001600160c01b038111156106f857604051631edd0da560e31b8152600481018290526024016104f1565b6000808554866000526020600020600080198661196e5787611937576119f0565b87841461195057505060018281018955935081816119fe565b825489811061195f57506119fe565b600396509450600190506119fe565b600287116119b15783611980576119f0565b60028714156119a557600184039350838a55838301549850889450836119a5576119fe565b508154935060016119fe565b60038714156119c957505060018201885581816119fe565b889450836119d6576119fe565b82548981106119e557506119fe565b9450600190506119fe565b63a6ca772e6000526004601cfd5b5b83811015611a385780830154600182018085015480831087831011611a245750829050815b938501939093555050600181811b016119ff565b8115611a66576001820360011c935083830154808a10611a585750611a66565b808385015550839150611a38565b6001810115611a755788828401555b5050505094509492505050565b61178d828260006003611916565b5080546000825560020290600052602060002090810190610ce191905b80821115611ac75760008082556001820155600201611aad565b5090565b602080825282518282018190526000919060409081850190868401855b82811015611b2957815180518552868101516001600160401b0316878601528501516001600160c01b03168585015260609093019290850190600101611ae8565b5091979650505050505050565b80356001600160401b0381168114611b4d57600080fd5b919050565b60008060408385031215611b6557600080fd5b611b6e83611b36565b9150611b7c60208401611b36565b90509250929050565b600060208284031215611b9757600080fd5b61121c82611b36565b6001600160a01b0381168114610ce157600080fd5b600060208284031215611bc757600080fd5b813561121c81611ba0565b600060208284031215611be457600080fd5b5035919050565b6001600160a01b0391909116815260200190565b600060208284031215611c1157600080fd5b8151801515811461121c57600080fd5b600060208284031215611c3357600080fd5b815161121c81611ba0565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b0383811690831681811015611c7457611c74611c3e565b039392505050565b600082821015611c8e57611c8e611c3e565b500390565b60006001600160401b03808316818516808303821115611cb557611cb5611c3e565b01949350505050565b60008219821115611cd157611cd1611c3e565b500190565b600082611cf357634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052603260045260246000fd5b6000600019821415611d2257611d22611c3e565b5060010190565b60006001600160c01b0383811690831681811015611c7457611c74611c3e565b634e487b7160e01b600052604160045260246000fd5b600060208284031215611d7157600080fd5b815163ffffffff8116811461121c57600080fd5b6000816000190483118215151615611d9f57611d9f611c3e565b50029056fea2646970667358221220d7ed98768aa449a7866f396db73379cc3dd387848af7a0bb3023d924f36d4cea64736f6c63430008090033",
}

// CacheManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use CacheManagerMetaData.ABI instead.
var CacheManagerABI = CacheManagerMetaData.ABI

// CacheManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CacheManagerMetaData.Bin instead.
var CacheManagerBin = CacheManagerMetaData.Bin

// DeployCacheManager deploys a new Ethereum contract, binding an instance of CacheManager to it.
func DeployCacheManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CacheManager, error) {
	parsed, err := CacheManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CacheManagerBin), backend)
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
// Solidity: function entries(uint256 ) view returns(bytes32 code, uint64 size, uint192 bid)
func (_CacheManager *CacheManagerCaller) Entries(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Code [32]byte
	Size uint64
	Bid  *big.Int
}, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "entries", arg0)

	outstruct := new(struct {
		Code [32]byte
		Size uint64
		Bid  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Code = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Size = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Bid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) view returns(bytes32 code, uint64 size, uint192 bid)
func (_CacheManager *CacheManagerSession) Entries(arg0 *big.Int) (struct {
	Code [32]byte
	Size uint64
	Bid  *big.Int
}, error) {
	return _CacheManager.Contract.Entries(&_CacheManager.CallOpts, arg0)
}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) view returns(bytes32 code, uint64 size, uint192 bid)
func (_CacheManager *CacheManagerCallerSession) Entries(arg0 *big.Int) (struct {
	Code [32]byte
	Size uint64
	Bid  *big.Int
}, error) {
	return _CacheManager.Contract.Entries(&_CacheManager.CallOpts, arg0)
}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes32,uint64,uint192)[])
func (_CacheManager *CacheManagerCaller) GetEntries(opts *bind.CallOpts) ([]CacheManagerEntry, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getEntries")

	if err != nil {
		return *new([]CacheManagerEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]CacheManagerEntry)).(*[]CacheManagerEntry)

	return out0, err

}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes32,uint64,uint192)[])
func (_CacheManager *CacheManagerSession) GetEntries() ([]CacheManagerEntry, error) {
	return _CacheManager.Contract.GetEntries(&_CacheManager.CallOpts)
}

// GetEntries is a free data retrieval call binding the contract method 0x17be85c3.
//
// Solidity: function getEntries() view returns((bytes32,uint64,uint192)[])
func (_CacheManager *CacheManagerCallerSession) GetEntries() ([]CacheManagerEntry, error) {
	return _CacheManager.Contract.GetEntries(&_CacheManager.CallOpts)
}

// GetMinBid32052a9b is a free data retrieval call binding the contract method 0x32052a9b.
//
// Solidity: function getMinBid(address program) view returns(uint192 min)
func (_CacheManager *CacheManagerCaller) GetMinBid32052a9b(opts *bind.CallOpts, program common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getMinBid", program)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinBid32052a9b is a free data retrieval call binding the contract method 0x32052a9b.
//
// Solidity: function getMinBid(address program) view returns(uint192 min)
func (_CacheManager *CacheManagerSession) GetMinBid32052a9b(program common.Address) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBid32052a9b(&_CacheManager.CallOpts, program)
}

// GetMinBid32052a9b is a free data retrieval call binding the contract method 0x32052a9b.
//
// Solidity: function getMinBid(address program) view returns(uint192 min)
func (_CacheManager *CacheManagerCallerSession) GetMinBid32052a9b(program common.Address) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBid32052a9b(&_CacheManager.CallOpts, program)
}

// GetMinBidc565a208 is a free data retrieval call binding the contract method 0xc565a208.
//
// Solidity: function getMinBid(bytes32 codehash) view returns(uint192 min)
func (_CacheManager *CacheManagerCaller) GetMinBidc565a208(opts *bind.CallOpts, codehash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getMinBid0", codehash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinBidc565a208 is a free data retrieval call binding the contract method 0xc565a208.
//
// Solidity: function getMinBid(bytes32 codehash) view returns(uint192 min)
func (_CacheManager *CacheManagerSession) GetMinBidc565a208(codehash [32]byte) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBidc565a208(&_CacheManager.CallOpts, codehash)
}

// GetMinBidc565a208 is a free data retrieval call binding the contract method 0xc565a208.
//
// Solidity: function getMinBid(bytes32 codehash) view returns(uint192 min)
func (_CacheManager *CacheManagerCallerSession) GetMinBidc565a208(codehash [32]byte) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBidc565a208(&_CacheManager.CallOpts, codehash)
}

// GetMinBidd29b303e is a free data retrieval call binding the contract method 0xd29b303e.
//
// Solidity: function getMinBid(uint64 size) view returns(uint192 min)
func (_CacheManager *CacheManagerCaller) GetMinBidd29b303e(opts *bind.CallOpts, size uint64) (*big.Int, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getMinBid1", size)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinBidd29b303e is a free data retrieval call binding the contract method 0xd29b303e.
//
// Solidity: function getMinBid(uint64 size) view returns(uint192 min)
func (_CacheManager *CacheManagerSession) GetMinBidd29b303e(size uint64) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBidd29b303e(&_CacheManager.CallOpts, size)
}

// GetMinBidd29b303e is a free data retrieval call binding the contract method 0xd29b303e.
//
// Solidity: function getMinBid(uint64 size) view returns(uint192 min)
func (_CacheManager *CacheManagerCallerSession) GetMinBidd29b303e(size uint64) (*big.Int, error) {
	return _CacheManager.Contract.GetMinBidd29b303e(&_CacheManager.CallOpts, size)
}

// GetSmallestEntries is a free data retrieval call binding the contract method 0xe9c1bc0f.
//
// Solidity: function getSmallestEntries(uint256 k) view returns((bytes32,uint64,uint192)[] result)
func (_CacheManager *CacheManagerCaller) GetSmallestEntries(opts *bind.CallOpts, k *big.Int) ([]CacheManagerEntry, error) {
	var out []interface{}
	err := _CacheManager.contract.Call(opts, &out, "getSmallestEntries", k)

	if err != nil {
		return *new([]CacheManagerEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]CacheManagerEntry)).(*[]CacheManagerEntry)

	return out0, err

}

// GetSmallestEntries is a free data retrieval call binding the contract method 0xe9c1bc0f.
//
// Solidity: function getSmallestEntries(uint256 k) view returns((bytes32,uint64,uint192)[] result)
func (_CacheManager *CacheManagerSession) GetSmallestEntries(k *big.Int) ([]CacheManagerEntry, error) {
	return _CacheManager.Contract.GetSmallestEntries(&_CacheManager.CallOpts, k)
}

// GetSmallestEntries is a free data retrieval call binding the contract method 0xe9c1bc0f.
//
// Solidity: function getSmallestEntries(uint256 k) view returns((bytes32,uint64,uint192)[] result)
func (_CacheManager *CacheManagerCallerSession) GetSmallestEntries(k *big.Int) ([]CacheManagerEntry, error) {
	return _CacheManager.Contract.GetSmallestEntries(&_CacheManager.CallOpts, k)
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

// Initialize is a paid mutator transaction binding the contract method 0x20f2f345.
//
// Solidity: function initialize(uint64 initCacheSize, uint64 initDecay) returns()
func (_CacheManager *CacheManagerTransactor) Initialize(opts *bind.TransactOpts, initCacheSize uint64, initDecay uint64) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "initialize", initCacheSize, initDecay)
}

// Initialize is a paid mutator transaction binding the contract method 0x20f2f345.
//
// Solidity: function initialize(uint64 initCacheSize, uint64 initDecay) returns()
func (_CacheManager *CacheManagerSession) Initialize(initCacheSize uint64, initDecay uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.Initialize(&_CacheManager.TransactOpts, initCacheSize, initDecay)
}

// Initialize is a paid mutator transaction binding the contract method 0x20f2f345.
//
// Solidity: function initialize(uint64 initCacheSize, uint64 initDecay) returns()
func (_CacheManager *CacheManagerTransactorSession) Initialize(initCacheSize uint64, initDecay uint64) (*types.Transaction, error) {
	return _CacheManager.Contract.Initialize(&_CacheManager.TransactOpts, initCacheSize, initDecay)
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

// PlaceBid is a paid mutator transaction binding the contract method 0xe4940157.
//
// Solidity: function placeBid(address program) payable returns()
func (_CacheManager *CacheManagerTransactor) PlaceBid(opts *bind.TransactOpts, program common.Address) (*types.Transaction, error) {
	return _CacheManager.contract.Transact(opts, "placeBid", program)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xe4940157.
//
// Solidity: function placeBid(address program) payable returns()
func (_CacheManager *CacheManagerSession) PlaceBid(program common.Address) (*types.Transaction, error) {
	return _CacheManager.Contract.PlaceBid(&_CacheManager.TransactOpts, program)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xe4940157.
//
// Solidity: function placeBid(address program) payable returns()
func (_CacheManager *CacheManagerTransactorSession) PlaceBid(program common.Address) (*types.Transaction, error) {
	return _CacheManager.Contract.PlaceBid(&_CacheManager.TransactOpts, program)
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
	Program  common.Address
	Bid      *big.Int
	Size     uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInsertBid is a free log retrieval operation binding the contract event 0xb9271ce6a232cb5e0010e10fc10b38fe5d25dd27f8c03beef068a581cfc21bec.
//
// Solidity: event InsertBid(bytes32 indexed codehash, address program, uint192 bid, uint64 size)
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

// WatchInsertBid is a free log subscription operation binding the contract event 0xb9271ce6a232cb5e0010e10fc10b38fe5d25dd27f8c03beef068a581cfc21bec.
//
// Solidity: event InsertBid(bytes32 indexed codehash, address program, uint192 bid, uint64 size)
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

// ParseInsertBid is a log parse operation binding the contract event 0xb9271ce6a232cb5e0010e10fc10b38fe5d25dd27f8c03beef068a581cfc21bec.
//
// Solidity: event InsertBid(bytes32 indexed codehash, address program, uint192 bid, uint64 size)
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
