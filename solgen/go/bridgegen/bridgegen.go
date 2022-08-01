// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgegen

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
)

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

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"}],\"name\":\"InvalidOutboxSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotDelayedInbox\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotOutbox\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotRollupOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotSequencerInbox\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BridgeCallTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"InboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"OutboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSequencerInbox\",\"type\":\"address\"}],\"name\":\"SequencerInboxUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptFundsFromOldBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeOutbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedDelayedInboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"allowedDelayedInboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedOutboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"}],\"name\":\"allowedOutboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delayedInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"enqueueDelayedMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"}],\"name\":\"enqueueSequencerMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seqMessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"acc\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"rollup_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setDelayedInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencerInbox\",\"type\":\"address\"}],\"name\":\"setSequencerInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"submitBatchSpendingReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b506080516115026100306000396000610c3601526115026000f3fe60806040526004361061011e5760003560e01c8063ab5d8943116100a0578063d5719dc211610064578063d5719dc214610375578063e76f5c8d14610395578063e77145f4146101d3578063eca067ad146103b5578063ee35f327146103ca57600080fd5b8063ab5d8943146102c3578063ae60bd13146102d8578063c4d66de814610315578063cb23bcb514610335578063cee3d7281461035557600080fd5b80637a88b107116100e75780637a88b107146101f5578063845bde82146102155780638db5993b14610255578063945e1147146102685780639e5d4c491461029557600080fd5b806284120c1461012357806316bf557914610147578063413b35bd1461016757806347fb24c5146101b35780634f61f850146101d5575b600080fd5b34801561012f57600080fd5b506007545b6040519081526020015b60405180910390f35b34801561015357600080fd5b506101346101623660046111d8565b6103ea565b34801561017357600080fd5b506101a3610182366004611209565b6001600160a01b031660009081526002602052604090206001015460ff1690565b604051901515815260200161013e565b3480156101bf57600080fd5b506101d36101ce366004611226565b61040b565b005b3480156101e157600080fd5b506101d36101f0366004611209565b610716565b34801561020157600080fd5b50610134610210366004611264565b61083b565b34801561022157600080fd5b50610235610230366004611290565b610881565b60408051948552602085019390935291830152606082015260800161013e565b6101346102633660046112b2565b610994565b34801561027457600080fd5b506102886102833660046111d8565b6109df565b60405161013e91906112f9565b3480156102a157600080fd5b506102b56102b036600461130d565b610a09565b60405161013e929190611396565b3480156102cf57600080fd5b50610288610b61565b3480156102e457600080fd5b506101a36102f3366004611209565b6001600160a01b03166000908152600160208190526040909120015460ff1690565b34801561032157600080fd5b506101d3610330366004611209565b610b87565b34801561034157600080fd5b50600854610288906001600160a01b031681565b34801561036157600080fd5b506101d3610370366004611226565b610cfb565b34801561038157600080fd5b506101346103903660046111d8565b611029565b3480156103a157600080fd5b506102886103b03660046111d8565b611039565b3480156103c157600080fd5b50600654610134565b3480156103d657600080fd5b50600954610288906001600160a01b031681565b600781815481106103fa57600080fd5b600091825260209091200154905081565b6008546001600160a01b031633146104e35760085460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b15801561046257600080fd5b505afa158015610476573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049a91906113f5565b9050336001600160a01b038216146104e157600854604051630739600760e01b81526104d89133916001600160a01b03909116908490600401611412565b60405180910390fd5b505b6001600160a01b0382166000818152600160208181526040928390209182015492518515158152919360ff90931692917f6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521910160405180910390a28080156105485750825b8061055a57508015801561055a575082155b156105655750505050565b82156105f357604080518082018252600380548252600160208084018281526001600160a01b038a166000818152928490529582209451855551938201805460ff1916941515949094179093558154908101825591527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b031916909117905561070f565b6003805461060390600190611435565b815481106106135761061361145a565b6000918252602090912001548254600380546001600160a01b039093169290919081106106425761064261145a565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081600001546001600060038560000154815481106106905761069061145a565b60009182526020808320909101546001600160a01b0316835282019290925260400190205560038054806106c6576106c6611470565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526001908190526040822091825501805460ff191690555b50505b5050565b6008546001600160a01b031633146107e55760085460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b15801561076d57600080fd5b505afa158015610781573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107a591906113f5565b9050336001600160a01b038216146107e357600854604051630739600760e01b81526104d89133916001600160a01b03909116908490600401611412565b505b600980546001600160a01b0319166001600160a01b0383161790556040517f8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a906108309083906112f9565b60405180910390a150565b6009546000906001600160a01b0316331461086b573360405163223e13c160e21b81526004016104d891906112f9565b61087a600d8443424887611049565b9392505050565b6009546000908190819081906001600160a01b031633146108b7573360405163223e13c160e21b81526004016104d891906112f9565b600754935083156108f057600780546108d290600190611435565b815481106108e2576108e261145a565b906000526020600020015492505b8415610921576006610903600187611435565b815481106109135761091361145a565b906000526020600020015491505b60408051602081018590529081018790526060810183905260800160408051601f198184030181529190528051602090910120600780546001810182556000919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68801819055939692955090935050565b3360009081526001602081905260408220015460ff166109c9573360405163b6c60ea360e01b81526004016104d891906112f9565b6109d7848443424887611049565b949350505050565b600481815481106109ef57600080fd5b6000918252602090912001546001600160a01b0316905081565b3360009081526002602052604081206001015460609060ff16610a4157336040516332ea82ab60e01b81526004016104d891906112f9565b8215801590610a5857506001600160a01b0386163b155b15610a78578560405163b5cf5b8f60e01b81526004016104d891906112f9565b600580546001600160a01b0319811633179091556040516001600160a01b03918216918816908790610aad9088908890611486565b60006040518083038185875af1925050503d8060008114610aea576040519150601f19603f3d011682016040523d82523d6000602084013e610aef565b606091505b50600580546001600160a01b0319166001600160a01b038581169190911790915560405192955090935088169033907f2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d46690610b4f908a908a908a90611496565b60405180910390a35094509492505050565b6005546000906001600160a01b0390811690811415610b8257600091505090565b919050565b600054610100900460ff16610ba25760005460ff1615610ba6565b303b155b610c095760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016104d8565b600054610100900460ff16158015610c2b576000805461ffff19166101011790555b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161415610cb95760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b60648201526084016104d8565b600580546001600160a01b036001600160a01b03199182168117909255600880549091169184169190911790558015610712576000805461ff00191690555050565b6008546001600160a01b03163314610dca5760085460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b158015610d5257600080fd5b505afa158015610d66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d8a91906113f5565b9050336001600160a01b03821614610dc857600854604051630739600760e01b81526104d89133916001600160a01b03909116908490600401611412565b505b6001600160a01b038281161415610df6578160405163077abed160e41b81526004016104d891906112f9565b6001600160a01b038216600081815260026020908152604091829020600181015492518515158152909360ff90931692917f49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa910160405180910390a2808015610e5c5750825b80610e6e575080158015610e6e575082155b15610e795750505050565b8215610f0857604080518082018252600480548252600160208084018281526001600160a01b038a16600081815260029093529582209451855551938201805460ff1916941515949094179093558154908101825591527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b031916909117905561070f565b60048054610f1890600190611435565b81548110610f2857610f2861145a565b6000918252602090912001548254600480546001600160a01b03909316929091908110610f5757610f5761145a565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508160000154600260006004856000015481548110610fa557610fa561145a565b60009182526020808320909101546001600160a01b031683528201929092526040019020556004805480610fdb57610fdb611470565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526002905260408120908155600101805460ff1916905550505050565b600681815481106103fa57600080fd5b600381815481106109ef57600080fd5b600654604080516001600160f81b031960f88a901b166020808301919091526bffffffffffffffffffffffff1960608a901b1660218301526001600160c01b031960c089811b8216603585015288901b16603d830152604582018490526065820186905260858083018690528351808403909101815260a5909201909252805191012060009190600082156111035760066110e5600185611435565b815481106110f5576110f561145a565b906000526020600020015490505b6040805160208082018490528183018590528251808303840181526060830180855281519190920120600680546001810182556000919091527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f015533905260ff8c1660808201526001600160a01b038b1660a082015260c0810187905260e0810188905267ffffffffffffffff89166101008201529051829185917f5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1918190036101200190a3509098975050505050505050565b6000602082840312156111ea57600080fd5b5035919050565b6001600160a01b038116811461120657600080fd5b50565b60006020828403121561121b57600080fd5b813561087a816111f1565b6000806040838503121561123957600080fd5b8235611244816111f1565b91506020830135801515811461125957600080fd5b809150509250929050565b6000806040838503121561127757600080fd5b8235611282816111f1565b946020939093013593505050565b600080604083850312156112a357600080fd5b50508035926020909101359150565b6000806000606084860312156112c757600080fd5b833560ff811681146112d857600080fd5b925060208401356112e8816111f1565b929592945050506040919091013590565b6001600160a01b0391909116815260200190565b6000806000806060858703121561132357600080fd5b843561132e816111f1565b935060208501359250604085013567ffffffffffffffff8082111561135257600080fd5b818701915087601f83011261136657600080fd5b81358181111561137557600080fd5b88602082850101111561138757600080fd5b95989497505060200194505050565b821515815260006020604081840152835180604085015260005b818110156113cc578581018301518582016060015282016113b0565b818111156113de576000606083870101525b50601f01601f191692909201606001949350505050565b60006020828403121561140757600080fd5b815161087a816111f1565b6001600160a01b0393841681529183166020830152909116604082015260600190565b60008282101561145557634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b8183823760009101908152919050565b83815260406020820152816040820152818360608301376000818301606090810191909152601f909201601f191601019291505056fea2646970667358221220ada0d8025a0be6a2b8a84ad4c71656836bef296d759baab37befd6709c2d16db64736f6c63430008090033",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMetaData.Bin instead.
var BridgeBin = BridgeMetaData.Bin

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_Bridge *BridgeCaller) ActiveOutbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "activeOutbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_Bridge *BridgeSession) ActiveOutbox() (common.Address, error) {
	return _Bridge.Contract.ActiveOutbox(&_Bridge.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_Bridge *BridgeCallerSession) ActiveOutbox() (common.Address, error) {
	return _Bridge.Contract.ActiveOutbox(&_Bridge.CallOpts)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_Bridge *BridgeCaller) AllowedDelayedInboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allowedDelayedInboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_Bridge *BridgeSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.AllowedDelayedInboxList(&_Bridge.CallOpts, arg0)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_Bridge *BridgeCallerSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.AllowedDelayedInboxList(&_Bridge.CallOpts, arg0)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_Bridge *BridgeCaller) AllowedDelayedInboxes(opts *bind.CallOpts, inbox common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allowedDelayedInboxes", inbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_Bridge *BridgeSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _Bridge.Contract.AllowedDelayedInboxes(&_Bridge.CallOpts, inbox)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_Bridge *BridgeCallerSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _Bridge.Contract.AllowedDelayedInboxes(&_Bridge.CallOpts, inbox)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_Bridge *BridgeCaller) AllowedOutboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allowedOutboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_Bridge *BridgeSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.AllowedOutboxList(&_Bridge.CallOpts, arg0)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_Bridge *BridgeCallerSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _Bridge.Contract.AllowedOutboxList(&_Bridge.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_Bridge *BridgeCaller) AllowedOutboxes(opts *bind.CallOpts, outbox common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "allowedOutboxes", outbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_Bridge *BridgeSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _Bridge.Contract.AllowedOutboxes(&_Bridge.CallOpts, outbox)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_Bridge *BridgeCallerSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _Bridge.Contract.AllowedOutboxes(&_Bridge.CallOpts, outbox)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_Bridge *BridgeCaller) DelayedInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "delayedInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_Bridge *BridgeSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _Bridge.Contract.DelayedInboxAccs(&_Bridge.CallOpts, arg0)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_Bridge *BridgeCallerSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _Bridge.Contract.DelayedInboxAccs(&_Bridge.CallOpts, arg0)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_Bridge *BridgeCaller) DelayedMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "delayedMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_Bridge *BridgeSession) DelayedMessageCount() (*big.Int, error) {
	return _Bridge.Contract.DelayedMessageCount(&_Bridge.CallOpts)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_Bridge *BridgeCallerSession) DelayedMessageCount() (*big.Int, error) {
	return _Bridge.Contract.DelayedMessageCount(&_Bridge.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_Bridge *BridgeCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_Bridge *BridgeSession) Rollup() (common.Address, error) {
	return _Bridge.Contract.Rollup(&_Bridge.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_Bridge *BridgeCallerSession) Rollup() (common.Address, error) {
	return _Bridge.Contract.Rollup(&_Bridge.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_Bridge *BridgeCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_Bridge *BridgeSession) SequencerInbox() (common.Address, error) {
	return _Bridge.Contract.SequencerInbox(&_Bridge.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_Bridge *BridgeCallerSession) SequencerInbox() (common.Address, error) {
	return _Bridge.Contract.SequencerInbox(&_Bridge.CallOpts)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_Bridge *BridgeCaller) SequencerInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "sequencerInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_Bridge *BridgeSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _Bridge.Contract.SequencerInboxAccs(&_Bridge.CallOpts, arg0)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_Bridge *BridgeCallerSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _Bridge.Contract.SequencerInboxAccs(&_Bridge.CallOpts, arg0)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_Bridge *BridgeCaller) SequencerMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "sequencerMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_Bridge *BridgeSession) SequencerMessageCount() (*big.Int, error) {
	return _Bridge.Contract.SequencerMessageCount(&_Bridge.CallOpts)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_Bridge *BridgeCallerSession) SequencerMessageCount() (*big.Int, error) {
	return _Bridge.Contract.SequencerMessageCount(&_Bridge.CallOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_Bridge *BridgeTransactor) AcceptFundsFromOldBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "acceptFundsFromOldBridge")
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_Bridge *BridgeSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _Bridge.Contract.AcceptFundsFromOldBridge(&_Bridge.TransactOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_Bridge *BridgeTransactorSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _Bridge.Contract.AcceptFundsFromOldBridge(&_Bridge.TransactOpts)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_Bridge *BridgeTransactor) EnqueueDelayedMessage(opts *bind.TransactOpts, kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "enqueueDelayedMessage", kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_Bridge *BridgeSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.EnqueueDelayedMessage(&_Bridge.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_Bridge *BridgeTransactorSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.EnqueueDelayedMessage(&_Bridge.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x845bde82.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_Bridge *BridgeTransactor) EnqueueSequencerMessage(opts *bind.TransactOpts, dataHash [32]byte, afterDelayedMessagesRead *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "enqueueSequencerMessage", dataHash, afterDelayedMessagesRead)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x845bde82.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_Bridge *BridgeSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.EnqueueSequencerMessage(&_Bridge.TransactOpts, dataHash, afterDelayedMessagesRead)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x845bde82.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_Bridge *BridgeTransactorSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.EnqueueSequencerMessage(&_Bridge.TransactOpts, dataHash, afterDelayedMessagesRead)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_Bridge *BridgeTransactor) ExecuteCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "executeCall", to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_Bridge *BridgeSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteCall(&_Bridge.TransactOpts, to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_Bridge *BridgeTransactorSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteCall(&_Bridge.TransactOpts, to, value, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_Bridge *BridgeTransactor) Initialize(opts *bind.TransactOpts, rollup_ common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "initialize", rollup_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_Bridge *BridgeSession) Initialize(rollup_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, rollup_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_Bridge *BridgeTransactorSession) Initialize(rollup_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, rollup_)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeTransactor) SetDelayedInbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setDelayedInbox", inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetDelayedInbox(&_Bridge.TransactOpts, inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_Bridge *BridgeTransactorSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetDelayedInbox(&_Bridge.TransactOpts, inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_Bridge *BridgeTransactor) SetOutbox(opts *bind.TransactOpts, outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setOutbox", outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_Bridge *BridgeSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetOutbox(&_Bridge.TransactOpts, outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_Bridge *BridgeTransactorSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _Bridge.Contract.SetOutbox(&_Bridge.TransactOpts, outbox, enabled)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_Bridge *BridgeTransactor) SetSequencerInbox(opts *bind.TransactOpts, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setSequencerInbox", _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_Bridge *BridgeSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetSequencerInbox(&_Bridge.TransactOpts, _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_Bridge *BridgeTransactorSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetSequencerInbox(&_Bridge.TransactOpts, _sequencerInbox)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address sender, bytes32 messageDataHash) returns(uint256)
func (_Bridge *BridgeTransactor) SubmitBatchSpendingReport(opts *bind.TransactOpts, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "submitBatchSpendingReport", sender, messageDataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address sender, bytes32 messageDataHash) returns(uint256)
func (_Bridge *BridgeSession) SubmitBatchSpendingReport(sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.SubmitBatchSpendingReport(&_Bridge.TransactOpts, sender, messageDataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address sender, bytes32 messageDataHash) returns(uint256)
func (_Bridge *BridgeTransactorSession) SubmitBatchSpendingReport(sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.SubmitBatchSpendingReport(&_Bridge.TransactOpts, sender, messageDataHash)
}

// BridgeBridgeCallTriggeredIterator is returned from FilterBridgeCallTriggered and is used to iterate over the raw logs and unpacked data for BridgeCallTriggered events raised by the Bridge contract.
type BridgeBridgeCallTriggeredIterator struct {
	Event *BridgeBridgeCallTriggered // Event containing the contract specifics and raw log

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
func (it *BridgeBridgeCallTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeBridgeCallTriggered)
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
		it.Event = new(BridgeBridgeCallTriggered)
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
func (it *BridgeBridgeCallTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeBridgeCallTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeBridgeCallTriggered represents a BridgeCallTriggered event raised by the Bridge contract.
type BridgeBridgeCallTriggered struct {
	Outbox common.Address
	To     common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeCallTriggered is a free log retrieval operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_Bridge *BridgeFilterer) FilterBridgeCallTriggered(opts *bind.FilterOpts, outbox []common.Address, to []common.Address) (*BridgeBridgeCallTriggeredIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeBridgeCallTriggeredIterator{contract: _Bridge.contract, event: "BridgeCallTriggered", logs: logs, sub: sub}, nil
}

// WatchBridgeCallTriggered is a free log subscription operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_Bridge *BridgeFilterer) WatchBridgeCallTriggered(opts *bind.WatchOpts, sink chan<- *BridgeBridgeCallTriggered, outbox []common.Address, to []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeBridgeCallTriggered)
				if err := _Bridge.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseBridgeCallTriggered(log types.Log) (*BridgeBridgeCallTriggered, error) {
	event := new(BridgeBridgeCallTriggered)
	if err := _Bridge.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeInboxToggleIterator is returned from FilterInboxToggle and is used to iterate over the raw logs and unpacked data for InboxToggle events raised by the Bridge contract.
type BridgeInboxToggleIterator struct {
	Event *BridgeInboxToggle // Event containing the contract specifics and raw log

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
func (it *BridgeInboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeInboxToggle)
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
		it.Event = new(BridgeInboxToggle)
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
func (it *BridgeInboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeInboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeInboxToggle represents a InboxToggle event raised by the Bridge contract.
type BridgeInboxToggle struct {
	Inbox   common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInboxToggle is a free log retrieval operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_Bridge *BridgeFilterer) FilterInboxToggle(opts *bind.FilterOpts, inbox []common.Address) (*BridgeInboxToggleIterator, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeInboxToggleIterator{contract: _Bridge.contract, event: "InboxToggle", logs: logs, sub: sub}, nil
}

// WatchInboxToggle is a free log subscription operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_Bridge *BridgeFilterer) WatchInboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeInboxToggle, inbox []common.Address) (event.Subscription, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeInboxToggle)
				if err := _Bridge.contract.UnpackLog(event, "InboxToggle", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseInboxToggle(log types.Log) (*BridgeInboxToggle, error) {
	event := new(BridgeInboxToggle)
	if err := _Bridge.contract.UnpackLog(event, "InboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the Bridge contract.
type BridgeMessageDeliveredIterator struct {
	Event *BridgeMessageDelivered // Event containing the contract specifics and raw log

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
func (it *BridgeMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeMessageDelivered)
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
		it.Event = new(BridgeMessageDelivered)
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
func (it *BridgeMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeMessageDelivered represents a MessageDelivered event raised by the Bridge contract.
type BridgeMessageDelivered struct {
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
func (_Bridge *BridgeFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (*BridgeMessageDeliveredIterator, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &BridgeMessageDeliveredIterator{contract: _Bridge.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_Bridge *BridgeFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *BridgeMessageDelivered, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeMessageDelivered)
				if err := _Bridge.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseMessageDelivered(log types.Log) (*BridgeMessageDelivered, error) {
	event := new(BridgeMessageDelivered)
	if err := _Bridge.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOutboxToggleIterator is returned from FilterOutboxToggle and is used to iterate over the raw logs and unpacked data for OutboxToggle events raised by the Bridge contract.
type BridgeOutboxToggleIterator struct {
	Event *BridgeOutboxToggle // Event containing the contract specifics and raw log

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
func (it *BridgeOutboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOutboxToggle)
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
		it.Event = new(BridgeOutboxToggle)
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
func (it *BridgeOutboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOutboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOutboxToggle represents a OutboxToggle event raised by the Bridge contract.
type BridgeOutboxToggle struct {
	Outbox  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOutboxToggle is a free log retrieval operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_Bridge *BridgeFilterer) FilterOutboxToggle(opts *bind.FilterOpts, outbox []common.Address) (*BridgeOutboxToggleIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOutboxToggleIterator{contract: _Bridge.contract, event: "OutboxToggle", logs: logs, sub: sub}, nil
}

// WatchOutboxToggle is a free log subscription operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_Bridge *BridgeFilterer) WatchOutboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeOutboxToggle, outbox []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOutboxToggle)
				if err := _Bridge.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseOutboxToggle(log types.Log) (*BridgeOutboxToggle, error) {
	event := new(BridgeOutboxToggle)
	if err := _Bridge.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSequencerInboxUpdatedIterator is returned from FilterSequencerInboxUpdated and is used to iterate over the raw logs and unpacked data for SequencerInboxUpdated events raised by the Bridge contract.
type BridgeSequencerInboxUpdatedIterator struct {
	Event *BridgeSequencerInboxUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeSequencerInboxUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSequencerInboxUpdated)
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
		it.Event = new(BridgeSequencerInboxUpdated)
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
func (it *BridgeSequencerInboxUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSequencerInboxUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSequencerInboxUpdated represents a SequencerInboxUpdated event raised by the Bridge contract.
type BridgeSequencerInboxUpdated struct {
	NewSequencerInbox common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSequencerInboxUpdated is a free log retrieval operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_Bridge *BridgeFilterer) FilterSequencerInboxUpdated(opts *bind.FilterOpts) (*BridgeSequencerInboxUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeSequencerInboxUpdatedIterator{contract: _Bridge.contract, event: "SequencerInboxUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerInboxUpdated is a free log subscription operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_Bridge *BridgeFilterer) WatchSequencerInboxUpdated(opts *bind.WatchOpts, sink chan<- *BridgeSequencerInboxUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSequencerInboxUpdated)
				if err := _Bridge.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseSequencerInboxUpdated(log types.Log) (*BridgeSequencerInboxUpdated, error) {
	event := new(BridgeSequencerInboxUpdated)
	if err := _Bridge.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeMetaData contains all meta data concerning the IBridge contract.
var IBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BridgeCallTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"InboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"OutboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSequencerInbox\",\"type\":\"address\"}],\"name\":\"SequencerInboxUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptFundsFromOldBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeOutbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"allowedDelayedInboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"}],\"name\":\"allowedOutboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"delayedInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"enqueueDelayedMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"}],\"name\":\"enqueueSequencerMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seqMessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"acc\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"sequencerInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setDelayedInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencerInbox\",\"type\":\"address\"}],\"name\":\"setSequencerInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"batchPoster\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"submitBatchSpendingReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"msgNum\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IBridgeMetaData.ABI instead.
var IBridgeABI = IBridgeMetaData.ABI

// IBridge is an auto generated Go binding around an Ethereum contract.
type IBridge struct {
	IBridgeCaller     // Read-only binding to the contract
	IBridgeTransactor // Write-only binding to the contract
	IBridgeFilterer   // Log filterer for contract events
}

// IBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBridgeSession struct {
	Contract     *IBridge          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBridgeCallerSession struct {
	Contract *IBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBridgeTransactorSession struct {
	Contract     *IBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBridgeRaw struct {
	Contract *IBridge // Generic contract binding to access the raw methods on
}

// IBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBridgeCallerRaw struct {
	Contract *IBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBridgeTransactorRaw struct {
	Contract *IBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBridge creates a new instance of IBridge, bound to a specific deployed contract.
func NewIBridge(address common.Address, backend bind.ContractBackend) (*IBridge, error) {
	contract, err := bindIBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBridge{IBridgeCaller: IBridgeCaller{contract: contract}, IBridgeTransactor: IBridgeTransactor{contract: contract}, IBridgeFilterer: IBridgeFilterer{contract: contract}}, nil
}

// NewIBridgeCaller creates a new read-only instance of IBridge, bound to a specific deployed contract.
func NewIBridgeCaller(address common.Address, caller bind.ContractCaller) (*IBridgeCaller, error) {
	contract, err := bindIBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeCaller{contract: contract}, nil
}

// NewIBridgeTransactor creates a new write-only instance of IBridge, bound to a specific deployed contract.
func NewIBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IBridgeTransactor, error) {
	contract, err := bindIBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeTransactor{contract: contract}, nil
}

// NewIBridgeFilterer creates a new log filterer instance of IBridge, bound to a specific deployed contract.
func NewIBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IBridgeFilterer, error) {
	contract, err := bindIBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBridgeFilterer{contract: contract}, nil
}

// bindIBridge binds a generic wrapper to an already deployed contract.
func bindIBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridge *IBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridge.Contract.IBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridge *IBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridge.Contract.IBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridge *IBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridge.Contract.IBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridge *IBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridge *IBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridge *IBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridge.Contract.contract.Transact(opts, method, params...)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_IBridge *IBridgeCaller) ActiveOutbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "activeOutbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_IBridge *IBridgeSession) ActiveOutbox() (common.Address, error) {
	return _IBridge.Contract.ActiveOutbox(&_IBridge.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_IBridge *IBridgeCallerSession) ActiveOutbox() (common.Address, error) {
	return _IBridge.Contract.ActiveOutbox(&_IBridge.CallOpts)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_IBridge *IBridgeCaller) AllowedDelayedInboxes(opts *bind.CallOpts, inbox common.Address) (bool, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "allowedDelayedInboxes", inbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_IBridge *IBridgeSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _IBridge.Contract.AllowedDelayedInboxes(&_IBridge.CallOpts, inbox)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_IBridge *IBridgeCallerSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _IBridge.Contract.AllowedDelayedInboxes(&_IBridge.CallOpts, inbox)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_IBridge *IBridgeCaller) AllowedOutboxes(opts *bind.CallOpts, outbox common.Address) (bool, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "allowedOutboxes", outbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_IBridge *IBridgeSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _IBridge.Contract.AllowedOutboxes(&_IBridge.CallOpts, outbox)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_IBridge *IBridgeCallerSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _IBridge.Contract.AllowedOutboxes(&_IBridge.CallOpts, outbox)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 index) view returns(bytes32)
func (_IBridge *IBridgeCaller) DelayedInboxAccs(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "delayedInboxAccs", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 index) view returns(bytes32)
func (_IBridge *IBridgeSession) DelayedInboxAccs(index *big.Int) ([32]byte, error) {
	return _IBridge.Contract.DelayedInboxAccs(&_IBridge.CallOpts, index)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 index) view returns(bytes32)
func (_IBridge *IBridgeCallerSession) DelayedInboxAccs(index *big.Int) ([32]byte, error) {
	return _IBridge.Contract.DelayedInboxAccs(&_IBridge.CallOpts, index)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_IBridge *IBridgeCaller) DelayedMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "delayedMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_IBridge *IBridgeSession) DelayedMessageCount() (*big.Int, error) {
	return _IBridge.Contract.DelayedMessageCount(&_IBridge.CallOpts)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_IBridge *IBridgeCallerSession) DelayedMessageCount() (*big.Int, error) {
	return _IBridge.Contract.DelayedMessageCount(&_IBridge.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_IBridge *IBridgeCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_IBridge *IBridgeSession) Rollup() (common.Address, error) {
	return _IBridge.Contract.Rollup(&_IBridge.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_IBridge *IBridgeCallerSession) Rollup() (common.Address, error) {
	return _IBridge.Contract.Rollup(&_IBridge.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_IBridge *IBridgeCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_IBridge *IBridgeSession) SequencerInbox() (common.Address, error) {
	return _IBridge.Contract.SequencerInbox(&_IBridge.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_IBridge *IBridgeCallerSession) SequencerInbox() (common.Address, error) {
	return _IBridge.Contract.SequencerInbox(&_IBridge.CallOpts)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 index) view returns(bytes32)
func (_IBridge *IBridgeCaller) SequencerInboxAccs(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "sequencerInboxAccs", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 index) view returns(bytes32)
func (_IBridge *IBridgeSession) SequencerInboxAccs(index *big.Int) ([32]byte, error) {
	return _IBridge.Contract.SequencerInboxAccs(&_IBridge.CallOpts, index)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 index) view returns(bytes32)
func (_IBridge *IBridgeCallerSession) SequencerInboxAccs(index *big.Int) ([32]byte, error) {
	return _IBridge.Contract.SequencerInboxAccs(&_IBridge.CallOpts, index)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_IBridge *IBridgeCaller) SequencerMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBridge.contract.Call(opts, &out, "sequencerMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_IBridge *IBridgeSession) SequencerMessageCount() (*big.Int, error) {
	return _IBridge.Contract.SequencerMessageCount(&_IBridge.CallOpts)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_IBridge *IBridgeCallerSession) SequencerMessageCount() (*big.Int, error) {
	return _IBridge.Contract.SequencerMessageCount(&_IBridge.CallOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_IBridge *IBridgeTransactor) AcceptFundsFromOldBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "acceptFundsFromOldBridge")
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_IBridge *IBridgeSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _IBridge.Contract.AcceptFundsFromOldBridge(&_IBridge.TransactOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_IBridge *IBridgeTransactorSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _IBridge.Contract.AcceptFundsFromOldBridge(&_IBridge.TransactOpts)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_IBridge *IBridgeTransactor) EnqueueDelayedMessage(opts *bind.TransactOpts, kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "enqueueDelayedMessage", kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_IBridge *IBridgeSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _IBridge.Contract.EnqueueDelayedMessage(&_IBridge.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_IBridge *IBridgeTransactorSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _IBridge.Contract.EnqueueDelayedMessage(&_IBridge.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x845bde82.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_IBridge *IBridgeTransactor) EnqueueSequencerMessage(opts *bind.TransactOpts, dataHash [32]byte, afterDelayedMessagesRead *big.Int) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "enqueueSequencerMessage", dataHash, afterDelayedMessagesRead)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x845bde82.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_IBridge *IBridgeSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int) (*types.Transaction, error) {
	return _IBridge.Contract.EnqueueSequencerMessage(&_IBridge.TransactOpts, dataHash, afterDelayedMessagesRead)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x845bde82.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_IBridge *IBridgeTransactorSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int) (*types.Transaction, error) {
	return _IBridge.Contract.EnqueueSequencerMessage(&_IBridge.TransactOpts, dataHash, afterDelayedMessagesRead)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_IBridge *IBridgeTransactor) ExecuteCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "executeCall", to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_IBridge *IBridgeSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IBridge.Contract.ExecuteCall(&_IBridge.TransactOpts, to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_IBridge *IBridgeTransactorSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IBridge.Contract.ExecuteCall(&_IBridge.TransactOpts, to, value, data)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_IBridge *IBridgeTransactor) SetDelayedInbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "setDelayedInbox", inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_IBridge *IBridgeSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _IBridge.Contract.SetDelayedInbox(&_IBridge.TransactOpts, inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_IBridge *IBridgeTransactorSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _IBridge.Contract.SetDelayedInbox(&_IBridge.TransactOpts, inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address inbox, bool enabled) returns()
func (_IBridge *IBridgeTransactor) SetOutbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "setOutbox", inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address inbox, bool enabled) returns()
func (_IBridge *IBridgeSession) SetOutbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _IBridge.Contract.SetOutbox(&_IBridge.TransactOpts, inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address inbox, bool enabled) returns()
func (_IBridge *IBridgeTransactorSession) SetOutbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _IBridge.Contract.SetOutbox(&_IBridge.TransactOpts, inbox, enabled)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_IBridge *IBridgeTransactor) SetSequencerInbox(opts *bind.TransactOpts, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "setSequencerInbox", _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_IBridge *IBridgeSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _IBridge.Contract.SetSequencerInbox(&_IBridge.TransactOpts, _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_IBridge *IBridgeTransactorSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _IBridge.Contract.SetSequencerInbox(&_IBridge.TransactOpts, _sequencerInbox)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256 msgNum)
func (_IBridge *IBridgeTransactor) SubmitBatchSpendingReport(opts *bind.TransactOpts, batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _IBridge.contract.Transact(opts, "submitBatchSpendingReport", batchPoster, dataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256 msgNum)
func (_IBridge *IBridgeSession) SubmitBatchSpendingReport(batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _IBridge.Contract.SubmitBatchSpendingReport(&_IBridge.TransactOpts, batchPoster, dataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256 msgNum)
func (_IBridge *IBridgeTransactorSession) SubmitBatchSpendingReport(batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _IBridge.Contract.SubmitBatchSpendingReport(&_IBridge.TransactOpts, batchPoster, dataHash)
}

// IBridgeBridgeCallTriggeredIterator is returned from FilterBridgeCallTriggered and is used to iterate over the raw logs and unpacked data for BridgeCallTriggered events raised by the IBridge contract.
type IBridgeBridgeCallTriggeredIterator struct {
	Event *IBridgeBridgeCallTriggered // Event containing the contract specifics and raw log

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
func (it *IBridgeBridgeCallTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeBridgeCallTriggered)
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
		it.Event = new(IBridgeBridgeCallTriggered)
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
func (it *IBridgeBridgeCallTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeBridgeCallTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeBridgeCallTriggered represents a BridgeCallTriggered event raised by the IBridge contract.
type IBridgeBridgeCallTriggered struct {
	Outbox common.Address
	To     common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeCallTriggered is a free log retrieval operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_IBridge *IBridgeFilterer) FilterBridgeCallTriggered(opts *bind.FilterOpts, outbox []common.Address, to []common.Address) (*IBridgeBridgeCallTriggeredIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IBridgeBridgeCallTriggeredIterator{contract: _IBridge.contract, event: "BridgeCallTriggered", logs: logs, sub: sub}, nil
}

// WatchBridgeCallTriggered is a free log subscription operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_IBridge *IBridgeFilterer) WatchBridgeCallTriggered(opts *bind.WatchOpts, sink chan<- *IBridgeBridgeCallTriggered, outbox []common.Address, to []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeBridgeCallTriggered)
				if err := _IBridge.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseBridgeCallTriggered(log types.Log) (*IBridgeBridgeCallTriggered, error) {
	event := new(IBridgeBridgeCallTriggered)
	if err := _IBridge.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeInboxToggleIterator is returned from FilterInboxToggle and is used to iterate over the raw logs and unpacked data for InboxToggle events raised by the IBridge contract.
type IBridgeInboxToggleIterator struct {
	Event *IBridgeInboxToggle // Event containing the contract specifics and raw log

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
func (it *IBridgeInboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeInboxToggle)
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
		it.Event = new(IBridgeInboxToggle)
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
func (it *IBridgeInboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeInboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeInboxToggle represents a InboxToggle event raised by the IBridge contract.
type IBridgeInboxToggle struct {
	Inbox   common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInboxToggle is a free log retrieval operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_IBridge *IBridgeFilterer) FilterInboxToggle(opts *bind.FilterOpts, inbox []common.Address) (*IBridgeInboxToggleIterator, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return &IBridgeInboxToggleIterator{contract: _IBridge.contract, event: "InboxToggle", logs: logs, sub: sub}, nil
}

// WatchInboxToggle is a free log subscription operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_IBridge *IBridgeFilterer) WatchInboxToggle(opts *bind.WatchOpts, sink chan<- *IBridgeInboxToggle, inbox []common.Address) (event.Subscription, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeInboxToggle)
				if err := _IBridge.contract.UnpackLog(event, "InboxToggle", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseInboxToggle(log types.Log) (*IBridgeInboxToggle, error) {
	event := new(IBridgeInboxToggle)
	if err := _IBridge.contract.UnpackLog(event, "InboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the IBridge contract.
type IBridgeMessageDeliveredIterator struct {
	Event *IBridgeMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IBridgeMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeMessageDelivered)
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
		it.Event = new(IBridgeMessageDelivered)
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
func (it *IBridgeMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeMessageDelivered represents a MessageDelivered event raised by the IBridge contract.
type IBridgeMessageDelivered struct {
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
func (_IBridge *IBridgeFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (*IBridgeMessageDeliveredIterator, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &IBridgeMessageDeliveredIterator{contract: _IBridge.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_IBridge *IBridgeFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *IBridgeMessageDelivered, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeMessageDelivered)
				if err := _IBridge.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseMessageDelivered(log types.Log) (*IBridgeMessageDelivered, error) {
	event := new(IBridgeMessageDelivered)
	if err := _IBridge.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeOutboxToggleIterator is returned from FilterOutboxToggle and is used to iterate over the raw logs and unpacked data for OutboxToggle events raised by the IBridge contract.
type IBridgeOutboxToggleIterator struct {
	Event *IBridgeOutboxToggle // Event containing the contract specifics and raw log

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
func (it *IBridgeOutboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeOutboxToggle)
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
		it.Event = new(IBridgeOutboxToggle)
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
func (it *IBridgeOutboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeOutboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeOutboxToggle represents a OutboxToggle event raised by the IBridge contract.
type IBridgeOutboxToggle struct {
	Outbox  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOutboxToggle is a free log retrieval operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_IBridge *IBridgeFilterer) FilterOutboxToggle(opts *bind.FilterOpts, outbox []common.Address) (*IBridgeOutboxToggleIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return &IBridgeOutboxToggleIterator{contract: _IBridge.contract, event: "OutboxToggle", logs: logs, sub: sub}, nil
}

// WatchOutboxToggle is a free log subscription operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_IBridge *IBridgeFilterer) WatchOutboxToggle(opts *bind.WatchOpts, sink chan<- *IBridgeOutboxToggle, outbox []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeOutboxToggle)
				if err := _IBridge.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseOutboxToggle(log types.Log) (*IBridgeOutboxToggle, error) {
	event := new(IBridgeOutboxToggle)
	if err := _IBridge.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgeSequencerInboxUpdatedIterator is returned from FilterSequencerInboxUpdated and is used to iterate over the raw logs and unpacked data for SequencerInboxUpdated events raised by the IBridge contract.
type IBridgeSequencerInboxUpdatedIterator struct {
	Event *IBridgeSequencerInboxUpdated // Event containing the contract specifics and raw log

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
func (it *IBridgeSequencerInboxUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgeSequencerInboxUpdated)
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
		it.Event = new(IBridgeSequencerInboxUpdated)
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
func (it *IBridgeSequencerInboxUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgeSequencerInboxUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgeSequencerInboxUpdated represents a SequencerInboxUpdated event raised by the IBridge contract.
type IBridgeSequencerInboxUpdated struct {
	NewSequencerInbox common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSequencerInboxUpdated is a free log retrieval operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_IBridge *IBridgeFilterer) FilterSequencerInboxUpdated(opts *bind.FilterOpts) (*IBridgeSequencerInboxUpdatedIterator, error) {

	logs, sub, err := _IBridge.contract.FilterLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return &IBridgeSequencerInboxUpdatedIterator{contract: _IBridge.contract, event: "SequencerInboxUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerInboxUpdated is a free log subscription operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_IBridge *IBridgeFilterer) WatchSequencerInboxUpdated(opts *bind.WatchOpts, sink chan<- *IBridgeSequencerInboxUpdated) (event.Subscription, error) {

	logs, sub, err := _IBridge.contract.WatchLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgeSequencerInboxUpdated)
				if err := _IBridge.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
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
func (_IBridge *IBridgeFilterer) ParseSequencerInboxUpdated(log types.Log) (*IBridgeSequencerInboxUpdated, error) {
	event := new(IBridgeSequencerInboxUpdated)
	if err := _IBridge.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelayedMessageProviderMetaData contains all meta data concerning the IDelayedMessageProvider contract.
var IDelayedMessageProviderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"}]",
}

// IDelayedMessageProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use IDelayedMessageProviderMetaData.ABI instead.
var IDelayedMessageProviderABI = IDelayedMessageProviderMetaData.ABI

// IDelayedMessageProvider is an auto generated Go binding around an Ethereum contract.
type IDelayedMessageProvider struct {
	IDelayedMessageProviderCaller     // Read-only binding to the contract
	IDelayedMessageProviderTransactor // Write-only binding to the contract
	IDelayedMessageProviderFilterer   // Log filterer for contract events
}

// IDelayedMessageProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDelayedMessageProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelayedMessageProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDelayedMessageProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelayedMessageProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDelayedMessageProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDelayedMessageProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDelayedMessageProviderSession struct {
	Contract     *IDelayedMessageProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IDelayedMessageProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDelayedMessageProviderCallerSession struct {
	Contract *IDelayedMessageProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// IDelayedMessageProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDelayedMessageProviderTransactorSession struct {
	Contract     *IDelayedMessageProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// IDelayedMessageProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDelayedMessageProviderRaw struct {
	Contract *IDelayedMessageProvider // Generic contract binding to access the raw methods on
}

// IDelayedMessageProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDelayedMessageProviderCallerRaw struct {
	Contract *IDelayedMessageProviderCaller // Generic read-only contract binding to access the raw methods on
}

// IDelayedMessageProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDelayedMessageProviderTransactorRaw struct {
	Contract *IDelayedMessageProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDelayedMessageProvider creates a new instance of IDelayedMessageProvider, bound to a specific deployed contract.
func NewIDelayedMessageProvider(address common.Address, backend bind.ContractBackend) (*IDelayedMessageProvider, error) {
	contract, err := bindIDelayedMessageProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDelayedMessageProvider{IDelayedMessageProviderCaller: IDelayedMessageProviderCaller{contract: contract}, IDelayedMessageProviderTransactor: IDelayedMessageProviderTransactor{contract: contract}, IDelayedMessageProviderFilterer: IDelayedMessageProviderFilterer{contract: contract}}, nil
}

// NewIDelayedMessageProviderCaller creates a new read-only instance of IDelayedMessageProvider, bound to a specific deployed contract.
func NewIDelayedMessageProviderCaller(address common.Address, caller bind.ContractCaller) (*IDelayedMessageProviderCaller, error) {
	contract, err := bindIDelayedMessageProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDelayedMessageProviderCaller{contract: contract}, nil
}

// NewIDelayedMessageProviderTransactor creates a new write-only instance of IDelayedMessageProvider, bound to a specific deployed contract.
func NewIDelayedMessageProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*IDelayedMessageProviderTransactor, error) {
	contract, err := bindIDelayedMessageProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDelayedMessageProviderTransactor{contract: contract}, nil
}

// NewIDelayedMessageProviderFilterer creates a new log filterer instance of IDelayedMessageProvider, bound to a specific deployed contract.
func NewIDelayedMessageProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*IDelayedMessageProviderFilterer, error) {
	contract, err := bindIDelayedMessageProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDelayedMessageProviderFilterer{contract: contract}, nil
}

// bindIDelayedMessageProvider binds a generic wrapper to an already deployed contract.
func bindIDelayedMessageProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDelayedMessageProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDelayedMessageProvider *IDelayedMessageProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDelayedMessageProvider.Contract.IDelayedMessageProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDelayedMessageProvider *IDelayedMessageProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDelayedMessageProvider.Contract.IDelayedMessageProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDelayedMessageProvider *IDelayedMessageProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDelayedMessageProvider.Contract.IDelayedMessageProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDelayedMessageProvider *IDelayedMessageProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDelayedMessageProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDelayedMessageProvider *IDelayedMessageProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDelayedMessageProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDelayedMessageProvider *IDelayedMessageProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDelayedMessageProvider.Contract.contract.Transact(opts, method, params...)
}

// IDelayedMessageProviderInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the IDelayedMessageProvider contract.
type IDelayedMessageProviderInboxMessageDeliveredIterator struct {
	Event *IDelayedMessageProviderInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IDelayedMessageProviderInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelayedMessageProviderInboxMessageDelivered)
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
		it.Event = new(IDelayedMessageProviderInboxMessageDelivered)
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
func (it *IDelayedMessageProviderInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelayedMessageProviderInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelayedMessageProviderInboxMessageDelivered represents a InboxMessageDelivered event raised by the IDelayedMessageProvider contract.
type IDelayedMessageProviderInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_IDelayedMessageProvider *IDelayedMessageProviderFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*IDelayedMessageProviderInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IDelayedMessageProvider.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &IDelayedMessageProviderInboxMessageDeliveredIterator{contract: _IDelayedMessageProvider.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_IDelayedMessageProvider *IDelayedMessageProviderFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *IDelayedMessageProviderInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IDelayedMessageProvider.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelayedMessageProviderInboxMessageDelivered)
				if err := _IDelayedMessageProvider.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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
func (_IDelayedMessageProvider *IDelayedMessageProviderFilterer) ParseInboxMessageDelivered(log types.Log) (*IDelayedMessageProviderInboxMessageDelivered, error) {
	event := new(IDelayedMessageProviderInboxMessageDelivered)
	if err := _IDelayedMessageProvider.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDelayedMessageProviderInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the IDelayedMessageProvider contract.
type IDelayedMessageProviderInboxMessageDeliveredFromOriginIterator struct {
	Event *IDelayedMessageProviderInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *IDelayedMessageProviderInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDelayedMessageProviderInboxMessageDeliveredFromOrigin)
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
		it.Event = new(IDelayedMessageProviderInboxMessageDeliveredFromOrigin)
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
func (it *IDelayedMessageProviderInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDelayedMessageProviderInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDelayedMessageProviderInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the IDelayedMessageProvider contract.
type IDelayedMessageProviderInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_IDelayedMessageProvider *IDelayedMessageProviderFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*IDelayedMessageProviderInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IDelayedMessageProvider.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &IDelayedMessageProviderInboxMessageDeliveredFromOriginIterator{contract: _IDelayedMessageProvider.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_IDelayedMessageProvider *IDelayedMessageProviderFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *IDelayedMessageProviderInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IDelayedMessageProvider.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDelayedMessageProviderInboxMessageDeliveredFromOrigin)
				if err := _IDelayedMessageProvider.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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
func (_IDelayedMessageProvider *IDelayedMessageProviderFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*IDelayedMessageProviderInboxMessageDeliveredFromOrigin, error) {
	event := new(IDelayedMessageProviderInboxMessageDeliveredFromOrigin)
	if err := _IDelayedMessageProvider.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IInboxMetaData contains all meta data concerning the IInbox contract.
var IInboxMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"arbTxCallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submissionRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"valueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"}],\"name\":\"depositEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"arbTxCallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submissionRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"valueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"unsafeCreateRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use IInboxMetaData.ABI instead.
var IInboxABI = IInboxMetaData.ABI

// IInbox is an auto generated Go binding around an Ethereum contract.
type IInbox struct {
	IInboxCaller     // Read-only binding to the contract
	IInboxTransactor // Write-only binding to the contract
	IInboxFilterer   // Log filterer for contract events
}

// IInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInboxSession struct {
	Contract     *IInbox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInboxCallerSession struct {
	Contract *IInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInboxTransactorSession struct {
	Contract     *IInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInboxRaw struct {
	Contract *IInbox // Generic contract binding to access the raw methods on
}

// IInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInboxCallerRaw struct {
	Contract *IInboxCaller // Generic read-only contract binding to access the raw methods on
}

// IInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInboxTransactorRaw struct {
	Contract *IInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInbox creates a new instance of IInbox, bound to a specific deployed contract.
func NewIInbox(address common.Address, backend bind.ContractBackend) (*IInbox, error) {
	contract, err := bindIInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInbox{IInboxCaller: IInboxCaller{contract: contract}, IInboxTransactor: IInboxTransactor{contract: contract}, IInboxFilterer: IInboxFilterer{contract: contract}}, nil
}

// NewIInboxCaller creates a new read-only instance of IInbox, bound to a specific deployed contract.
func NewIInboxCaller(address common.Address, caller bind.ContractCaller) (*IInboxCaller, error) {
	contract, err := bindIInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInboxCaller{contract: contract}, nil
}

// NewIInboxTransactor creates a new write-only instance of IInbox, bound to a specific deployed contract.
func NewIInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IInboxTransactor, error) {
	contract, err := bindIInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInboxTransactor{contract: contract}, nil
}

// NewIInboxFilterer creates a new log filterer instance of IInbox, bound to a specific deployed contract.
func NewIInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IInboxFilterer, error) {
	contract, err := bindIInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInboxFilterer{contract: contract}, nil
}

// bindIInbox binds a generic wrapper to an already deployed contract.
func bindIInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInbox *IInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInbox.Contract.IInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInbox *IInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInbox.Contract.IInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInbox *IInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInbox.Contract.IInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInbox *IInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInbox *IInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInbox *IInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInbox.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IInbox *IInboxCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IInbox.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IInbox *IInboxSession) Bridge() (common.Address, error) {
	return _IInbox.Contract.Bridge(&_IInbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_IInbox *IInboxCallerSession) Bridge() (common.Address, error) {
	return _IInbox.Contract.Bridge(&_IInbox.CallOpts)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address to, uint256 arbTxCallValue, uint256 maxSubmissionCost, address submissionRefundAddress, address valueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_IInbox *IInboxTransactor) CreateRetryableTicket(opts *bind.TransactOpts, to common.Address, arbTxCallValue *big.Int, maxSubmissionCost *big.Int, submissionRefundAddress common.Address, valueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "createRetryableTicket", to, arbTxCallValue, maxSubmissionCost, submissionRefundAddress, valueRefundAddress, gasLimit, maxFeePerGas, data)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address to, uint256 arbTxCallValue, uint256 maxSubmissionCost, address submissionRefundAddress, address valueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_IInbox *IInboxSession) CreateRetryableTicket(to common.Address, arbTxCallValue *big.Int, maxSubmissionCost *big.Int, submissionRefundAddress common.Address, valueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.CreateRetryableTicket(&_IInbox.TransactOpts, to, arbTxCallValue, maxSubmissionCost, submissionRefundAddress, valueRefundAddress, gasLimit, maxFeePerGas, data)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address to, uint256 arbTxCallValue, uint256 maxSubmissionCost, address submissionRefundAddress, address valueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_IInbox *IInboxTransactorSession) CreateRetryableTicket(to common.Address, arbTxCallValue *big.Int, maxSubmissionCost *big.Int, submissionRefundAddress common.Address, valueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.CreateRetryableTicket(&_IInbox.TransactOpts, to, arbTxCallValue, maxSubmissionCost, submissionRefundAddress, valueRefundAddress, gasLimit, maxFeePerGas, data)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 maxSubmissionCost) payable returns(uint256)
func (_IInbox *IInboxTransactor) DepositEth(opts *bind.TransactOpts, maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "depositEth", maxSubmissionCost)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 maxSubmissionCost) payable returns(uint256)
func (_IInbox *IInboxSession) DepositEth(maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _IInbox.Contract.DepositEth(&_IInbox.TransactOpts, maxSubmissionCost)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 maxSubmissionCost) payable returns(uint256)
func (_IInbox *IInboxTransactorSession) DepositEth(maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _IInbox.Contract.DepositEth(&_IInbox.TransactOpts, maxSubmissionCost)
}

// DepositEth0 is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_IInbox *IInboxTransactor) DepositEth0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "depositEth0")
}

// DepositEth0 is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_IInbox *IInboxSession) DepositEth0() (*types.Transaction, error) {
	return _IInbox.Contract.DepositEth0(&_IInbox.TransactOpts)
}

// DepositEth0 is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_IInbox *IInboxTransactorSession) DepositEth0() (*types.Transaction, error) {
	return _IInbox.Contract.DepositEth0(&_IInbox.TransactOpts)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_IInbox *IInboxTransactor) PostUpgradeInit(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "postUpgradeInit", _bridge)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_IInbox *IInboxSession) PostUpgradeInit(_bridge common.Address) (*types.Transaction, error) {
	return _IInbox.Contract.PostUpgradeInit(&_IInbox.TransactOpts, _bridge)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_IInbox *IInboxTransactorSession) PostUpgradeInit(_bridge common.Address) (*types.Transaction, error) {
	return _IInbox.Contract.PostUpgradeInit(&_IInbox.TransactOpts, _bridge)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, uint256 value, bytes data) returns(uint256)
func (_IInbox *IInboxTransactor) SendContractTransaction(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendContractTransaction", gasLimit, maxFeePerGas, to, value, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, uint256 value, bytes data) returns(uint256)
func (_IInbox *IInboxSession) SendContractTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendContractTransaction(&_IInbox.TransactOpts, gasLimit, maxFeePerGas, to, value, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, uint256 value, bytes data) returns(uint256)
func (_IInbox *IInboxTransactorSession) SendContractTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendContractTransaction(&_IInbox.TransactOpts, gasLimit, maxFeePerGas, to, value, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, bytes data) payable returns(uint256)
func (_IInbox *IInboxTransactor) SendL1FundedContractTransaction(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendL1FundedContractTransaction", gasLimit, maxFeePerGas, to, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, bytes data) payable returns(uint256)
func (_IInbox *IInboxSession) SendL1FundedContractTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL1FundedContractTransaction(&_IInbox.TransactOpts, gasLimit, maxFeePerGas, to, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, bytes data) payable returns(uint256)
func (_IInbox *IInboxTransactorSession) SendL1FundedContractTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL1FundedContractTransaction(&_IInbox.TransactOpts, gasLimit, maxFeePerGas, to, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, bytes data) payable returns(uint256)
func (_IInbox *IInboxTransactor) SendL1FundedUnsignedTransaction(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendL1FundedUnsignedTransaction", gasLimit, maxFeePerGas, nonce, to, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, bytes data) payable returns(uint256)
func (_IInbox *IInboxSession) SendL1FundedUnsignedTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL1FundedUnsignedTransaction(&_IInbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, bytes data) payable returns(uint256)
func (_IInbox *IInboxTransactorSession) SendL1FundedUnsignedTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL1FundedUnsignedTransaction(&_IInbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, data)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_IInbox *IInboxTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_IInbox *IInboxSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL2Message(&_IInbox.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_IInbox *IInboxTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendL2Message(&_IInbox.TransactOpts, messageData)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, uint256 value, bytes data) returns(uint256)
func (_IInbox *IInboxTransactor) SendUnsignedTransaction(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "sendUnsignedTransaction", gasLimit, maxFeePerGas, nonce, to, value, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, uint256 value, bytes data) returns(uint256)
func (_IInbox *IInboxSession) SendUnsignedTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendUnsignedTransaction(&_IInbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, value, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, uint256 value, bytes data) returns(uint256)
func (_IInbox *IInboxTransactorSession) SendUnsignedTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.SendUnsignedTransaction(&_IInbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, value, data)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address to, uint256 arbTxCallValue, uint256 maxSubmissionCost, address submissionRefundAddress, address valueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_IInbox *IInboxTransactor) UnsafeCreateRetryableTicket(opts *bind.TransactOpts, to common.Address, arbTxCallValue *big.Int, maxSubmissionCost *big.Int, submissionRefundAddress common.Address, valueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.contract.Transact(opts, "unsafeCreateRetryableTicket", to, arbTxCallValue, maxSubmissionCost, submissionRefundAddress, valueRefundAddress, gasLimit, maxFeePerGas, data)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address to, uint256 arbTxCallValue, uint256 maxSubmissionCost, address submissionRefundAddress, address valueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_IInbox *IInboxSession) UnsafeCreateRetryableTicket(to common.Address, arbTxCallValue *big.Int, maxSubmissionCost *big.Int, submissionRefundAddress common.Address, valueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.UnsafeCreateRetryableTicket(&_IInbox.TransactOpts, to, arbTxCallValue, maxSubmissionCost, submissionRefundAddress, valueRefundAddress, gasLimit, maxFeePerGas, data)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address to, uint256 arbTxCallValue, uint256 maxSubmissionCost, address submissionRefundAddress, address valueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_IInbox *IInboxTransactorSession) UnsafeCreateRetryableTicket(to common.Address, arbTxCallValue *big.Int, maxSubmissionCost *big.Int, submissionRefundAddress common.Address, valueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _IInbox.Contract.UnsafeCreateRetryableTicket(&_IInbox.TransactOpts, to, arbTxCallValue, maxSubmissionCost, submissionRefundAddress, valueRefundAddress, gasLimit, maxFeePerGas, data)
}

// IInboxInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the IInbox contract.
type IInboxInboxMessageDeliveredIterator struct {
	Event *IInboxInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IInboxInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInboxInboxMessageDelivered)
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
		it.Event = new(IInboxInboxMessageDelivered)
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
func (it *IInboxInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInboxInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInboxInboxMessageDelivered represents a InboxMessageDelivered event raised by the IInbox contract.
type IInboxInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_IInbox *IInboxFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*IInboxInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &IInboxInboxMessageDeliveredIterator{contract: _IInbox.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_IInbox *IInboxFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *IInboxInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInboxInboxMessageDelivered)
				if err := _IInbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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
func (_IInbox *IInboxFilterer) ParseInboxMessageDelivered(log types.Log) (*IInboxInboxMessageDelivered, error) {
	event := new(IInboxInboxMessageDelivered)
	if err := _IInbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IInboxInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the IInbox contract.
type IInboxInboxMessageDeliveredFromOriginIterator struct {
	Event *IInboxInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *IInboxInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInboxInboxMessageDeliveredFromOrigin)
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
		it.Event = new(IInboxInboxMessageDeliveredFromOrigin)
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
func (it *IInboxInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInboxInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInboxInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the IInbox contract.
type IInboxInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_IInbox *IInboxFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*IInboxInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &IInboxInboxMessageDeliveredFromOriginIterator{contract: _IInbox.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_IInbox *IInboxFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *IInboxInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _IInbox.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInboxInboxMessageDeliveredFromOrigin)
				if err := _IInbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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
func (_IInbox *IInboxFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*IInboxInboxMessageDeliveredFromOrigin, error) {
	event := new(IInboxInboxMessageDeliveredFromOrigin)
	if err := _IInbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOutboxMetaData contains all meta data concerning the IOutbox contract.
var IOutboxMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"zero\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionIndex\",\"type\":\"uint256\"}],\"name\":\"OutBoxTransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"}],\"name\":\"SendRootUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"calculateItemHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"item\",\"type\":\"bytes32\"}],\"name\":\"calculateMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeTransactionSimulation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1BatchNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1EthBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1OutputId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sendRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2BlockHash\",\"type\":\"bytes32\"}],\"name\":\"updateSendRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IOutboxABI is the input ABI used to generate the binding from.
// Deprecated: Use IOutboxMetaData.ABI instead.
var IOutboxABI = IOutboxMetaData.ABI

// IOutbox is an auto generated Go binding around an Ethereum contract.
type IOutbox struct {
	IOutboxCaller     // Read-only binding to the contract
	IOutboxTransactor // Write-only binding to the contract
	IOutboxFilterer   // Log filterer for contract events
}

// IOutboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOutboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOutboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOutboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOutboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOutboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOutboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOutboxSession struct {
	Contract     *IOutbox          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOutboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOutboxCallerSession struct {
	Contract *IOutboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IOutboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOutboxTransactorSession struct {
	Contract     *IOutboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IOutboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOutboxRaw struct {
	Contract *IOutbox // Generic contract binding to access the raw methods on
}

// IOutboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOutboxCallerRaw struct {
	Contract *IOutboxCaller // Generic read-only contract binding to access the raw methods on
}

// IOutboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOutboxTransactorRaw struct {
	Contract *IOutboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOutbox creates a new instance of IOutbox, bound to a specific deployed contract.
func NewIOutbox(address common.Address, backend bind.ContractBackend) (*IOutbox, error) {
	contract, err := bindIOutbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOutbox{IOutboxCaller: IOutboxCaller{contract: contract}, IOutboxTransactor: IOutboxTransactor{contract: contract}, IOutboxFilterer: IOutboxFilterer{contract: contract}}, nil
}

// NewIOutboxCaller creates a new read-only instance of IOutbox, bound to a specific deployed contract.
func NewIOutboxCaller(address common.Address, caller bind.ContractCaller) (*IOutboxCaller, error) {
	contract, err := bindIOutbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOutboxCaller{contract: contract}, nil
}

// NewIOutboxTransactor creates a new write-only instance of IOutbox, bound to a specific deployed contract.
func NewIOutboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IOutboxTransactor, error) {
	contract, err := bindIOutbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOutboxTransactor{contract: contract}, nil
}

// NewIOutboxFilterer creates a new log filterer instance of IOutbox, bound to a specific deployed contract.
func NewIOutboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IOutboxFilterer, error) {
	contract, err := bindIOutbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOutboxFilterer{contract: contract}, nil
}

// bindIOutbox binds a generic wrapper to an already deployed contract.
func bindIOutbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOutboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOutbox *IOutboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOutbox.Contract.IOutboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOutbox *IOutboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOutbox.Contract.IOutboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOutbox *IOutboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOutbox.Contract.IOutboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOutbox *IOutboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOutbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOutbox *IOutboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOutbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOutbox *IOutboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOutbox.Contract.contract.Transact(opts, method, params...)
}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_IOutbox *IOutboxCaller) CalculateItemHash(opts *bind.CallOpts, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	var out []interface{}
	err := _IOutbox.contract.Call(opts, &out, "calculateItemHash", l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_IOutbox *IOutboxSession) CalculateItemHash(l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	return _IOutbox.Contract.CalculateItemHash(&_IOutbox.CallOpts, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_IOutbox *IOutboxCallerSession) CalculateItemHash(l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	return _IOutbox.Contract.CalculateItemHash(&_IOutbox.CallOpts, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_IOutbox *IOutboxCaller) CalculateMerkleRoot(opts *bind.CallOpts, proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IOutbox.contract.Call(opts, &out, "calculateMerkleRoot", proof, path, item)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_IOutbox *IOutboxSession) CalculateMerkleRoot(proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	return _IOutbox.Contract.CalculateMerkleRoot(&_IOutbox.CallOpts, proof, path, item)
}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_IOutbox *IOutboxCallerSession) CalculateMerkleRoot(proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	return _IOutbox.Contract.CalculateMerkleRoot(&_IOutbox.CallOpts, proof, path, item)
}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 ) view returns(bool)
func (_IOutbox *IOutboxCaller) IsSpent(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _IOutbox.contract.Call(opts, &out, "isSpent", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 ) view returns(bool)
func (_IOutbox *IOutboxSession) IsSpent(arg0 *big.Int) (bool, error) {
	return _IOutbox.Contract.IsSpent(&_IOutbox.CallOpts, arg0)
}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 ) view returns(bool)
func (_IOutbox *IOutboxCallerSession) IsSpent(arg0 *big.Int) (bool, error) {
	return _IOutbox.Contract.IsSpent(&_IOutbox.CallOpts, arg0)
}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() view returns(uint256)
func (_IOutbox *IOutboxCaller) L2ToL1BatchNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOutbox.contract.Call(opts, &out, "l2ToL1BatchNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() view returns(uint256)
func (_IOutbox *IOutboxSession) L2ToL1BatchNum() (*big.Int, error) {
	return _IOutbox.Contract.L2ToL1BatchNum(&_IOutbox.CallOpts)
}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() view returns(uint256)
func (_IOutbox *IOutboxCallerSession) L2ToL1BatchNum() (*big.Int, error) {
	return _IOutbox.Contract.L2ToL1BatchNum(&_IOutbox.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_IOutbox *IOutboxCaller) L2ToL1Block(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOutbox.contract.Call(opts, &out, "l2ToL1Block")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_IOutbox *IOutboxSession) L2ToL1Block() (*big.Int, error) {
	return _IOutbox.Contract.L2ToL1Block(&_IOutbox.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_IOutbox *IOutboxCallerSession) L2ToL1Block() (*big.Int, error) {
	return _IOutbox.Contract.L2ToL1Block(&_IOutbox.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_IOutbox *IOutboxCaller) L2ToL1EthBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOutbox.contract.Call(opts, &out, "l2ToL1EthBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_IOutbox *IOutboxSession) L2ToL1EthBlock() (*big.Int, error) {
	return _IOutbox.Contract.L2ToL1EthBlock(&_IOutbox.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_IOutbox *IOutboxCallerSession) L2ToL1EthBlock() (*big.Int, error) {
	return _IOutbox.Contract.L2ToL1EthBlock(&_IOutbox.CallOpts)
}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_IOutbox *IOutboxCaller) L2ToL1OutputId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IOutbox.contract.Call(opts, &out, "l2ToL1OutputId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_IOutbox *IOutboxSession) L2ToL1OutputId() ([32]byte, error) {
	return _IOutbox.Contract.L2ToL1OutputId(&_IOutbox.CallOpts)
}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_IOutbox *IOutboxCallerSession) L2ToL1OutputId() ([32]byte, error) {
	return _IOutbox.Contract.L2ToL1OutputId(&_IOutbox.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_IOutbox *IOutboxCaller) L2ToL1Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOutbox.contract.Call(opts, &out, "l2ToL1Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_IOutbox *IOutboxSession) L2ToL1Sender() (common.Address, error) {
	return _IOutbox.Contract.L2ToL1Sender(&_IOutbox.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_IOutbox *IOutboxCallerSession) L2ToL1Sender() (common.Address, error) {
	return _IOutbox.Contract.L2ToL1Sender(&_IOutbox.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_IOutbox *IOutboxCaller) L2ToL1Timestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOutbox.contract.Call(opts, &out, "l2ToL1Timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_IOutbox *IOutboxSession) L2ToL1Timestamp() (*big.Int, error) {
	return _IOutbox.Contract.L2ToL1Timestamp(&_IOutbox.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_IOutbox *IOutboxCallerSession) L2ToL1Timestamp() (*big.Int, error) {
	return _IOutbox.Contract.L2ToL1Timestamp(&_IOutbox.CallOpts)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_IOutbox *IOutboxTransactor) ExecuteTransaction(opts *bind.TransactOpts, proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IOutbox.contract.Transact(opts, "executeTransaction", proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_IOutbox *IOutboxSession) ExecuteTransaction(proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IOutbox.Contract.ExecuteTransaction(&_IOutbox.TransactOpts, proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_IOutbox *IOutboxTransactorSession) ExecuteTransaction(proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IOutbox.Contract.ExecuteTransaction(&_IOutbox.TransactOpts, proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransactionSimulation is a paid mutator transaction binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_IOutbox *IOutboxTransactor) ExecuteTransactionSimulation(opts *bind.TransactOpts, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IOutbox.contract.Transact(opts, "executeTransactionSimulation", index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransactionSimulation is a paid mutator transaction binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_IOutbox *IOutboxSession) ExecuteTransactionSimulation(index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IOutbox.Contract.ExecuteTransactionSimulation(&_IOutbox.TransactOpts, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransactionSimulation is a paid mutator transaction binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_IOutbox *IOutboxTransactorSession) ExecuteTransactionSimulation(index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _IOutbox.Contract.ExecuteTransactionSimulation(&_IOutbox.TransactOpts, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 sendRoot, bytes32 l2BlockHash) returns()
func (_IOutbox *IOutboxTransactor) UpdateSendRoot(opts *bind.TransactOpts, sendRoot [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _IOutbox.contract.Transact(opts, "updateSendRoot", sendRoot, l2BlockHash)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 sendRoot, bytes32 l2BlockHash) returns()
func (_IOutbox *IOutboxSession) UpdateSendRoot(sendRoot [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _IOutbox.Contract.UpdateSendRoot(&_IOutbox.TransactOpts, sendRoot, l2BlockHash)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 sendRoot, bytes32 l2BlockHash) returns()
func (_IOutbox *IOutboxTransactorSession) UpdateSendRoot(sendRoot [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _IOutbox.Contract.UpdateSendRoot(&_IOutbox.TransactOpts, sendRoot, l2BlockHash)
}

// IOutboxOutBoxTransactionExecutedIterator is returned from FilterOutBoxTransactionExecuted and is used to iterate over the raw logs and unpacked data for OutBoxTransactionExecuted events raised by the IOutbox contract.
type IOutboxOutBoxTransactionExecutedIterator struct {
	Event *IOutboxOutBoxTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *IOutboxOutBoxTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOutboxOutBoxTransactionExecuted)
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
		it.Event = new(IOutboxOutBoxTransactionExecuted)
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
func (it *IOutboxOutBoxTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOutboxOutBoxTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOutboxOutBoxTransactionExecuted represents a OutBoxTransactionExecuted event raised by the IOutbox contract.
type IOutboxOutBoxTransactionExecuted struct {
	To               common.Address
	L2Sender         common.Address
	Zero             *big.Int
	TransactionIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOutBoxTransactionExecuted is a free log retrieval operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_IOutbox *IOutboxFilterer) FilterOutBoxTransactionExecuted(opts *bind.FilterOpts, to []common.Address, l2Sender []common.Address, zero []*big.Int) (*IOutboxOutBoxTransactionExecutedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var zeroRule []interface{}
	for _, zeroItem := range zero {
		zeroRule = append(zeroRule, zeroItem)
	}

	logs, sub, err := _IOutbox.contract.FilterLogs(opts, "OutBoxTransactionExecuted", toRule, l2SenderRule, zeroRule)
	if err != nil {
		return nil, err
	}
	return &IOutboxOutBoxTransactionExecutedIterator{contract: _IOutbox.contract, event: "OutBoxTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchOutBoxTransactionExecuted is a free log subscription operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_IOutbox *IOutboxFilterer) WatchOutBoxTransactionExecuted(opts *bind.WatchOpts, sink chan<- *IOutboxOutBoxTransactionExecuted, to []common.Address, l2Sender []common.Address, zero []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var zeroRule []interface{}
	for _, zeroItem := range zero {
		zeroRule = append(zeroRule, zeroItem)
	}

	logs, sub, err := _IOutbox.contract.WatchLogs(opts, "OutBoxTransactionExecuted", toRule, l2SenderRule, zeroRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOutboxOutBoxTransactionExecuted)
				if err := _IOutbox.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
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

// ParseOutBoxTransactionExecuted is a log parse operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_IOutbox *IOutboxFilterer) ParseOutBoxTransactionExecuted(log types.Log) (*IOutboxOutBoxTransactionExecuted, error) {
	event := new(IOutboxOutBoxTransactionExecuted)
	if err := _IOutbox.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOutboxSendRootUpdatedIterator is returned from FilterSendRootUpdated and is used to iterate over the raw logs and unpacked data for SendRootUpdated events raised by the IOutbox contract.
type IOutboxSendRootUpdatedIterator struct {
	Event *IOutboxSendRootUpdated // Event containing the contract specifics and raw log

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
func (it *IOutboxSendRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOutboxSendRootUpdated)
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
		it.Event = new(IOutboxSendRootUpdated)
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
func (it *IOutboxSendRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOutboxSendRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOutboxSendRootUpdated represents a SendRootUpdated event raised by the IOutbox contract.
type IOutboxSendRootUpdated struct {
	BlockHash  [32]byte
	OutputRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSendRootUpdated is a free log retrieval operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed blockHash, bytes32 indexed outputRoot)
func (_IOutbox *IOutboxFilterer) FilterSendRootUpdated(opts *bind.FilterOpts, blockHash [][32]byte, outputRoot [][32]byte) (*IOutboxSendRootUpdatedIterator, error) {

	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}

	logs, sub, err := _IOutbox.contract.FilterLogs(opts, "SendRootUpdated", blockHashRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return &IOutboxSendRootUpdatedIterator{contract: _IOutbox.contract, event: "SendRootUpdated", logs: logs, sub: sub}, nil
}

// WatchSendRootUpdated is a free log subscription operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed blockHash, bytes32 indexed outputRoot)
func (_IOutbox *IOutboxFilterer) WatchSendRootUpdated(opts *bind.WatchOpts, sink chan<- *IOutboxSendRootUpdated, blockHash [][32]byte, outputRoot [][32]byte) (event.Subscription, error) {

	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}

	logs, sub, err := _IOutbox.contract.WatchLogs(opts, "SendRootUpdated", blockHashRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOutboxSendRootUpdated)
				if err := _IOutbox.contract.UnpackLog(event, "SendRootUpdated", log); err != nil {
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

// ParseSendRootUpdated is a log parse operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed blockHash, bytes32 indexed outputRoot)
func (_IOutbox *IOutboxFilterer) ParseSendRootUpdated(log types.Log) (*IOutboxSendRootUpdated, error) {
	event := new(IOutboxSendRootUpdated)
	if err := _IOutbox.contract.UnpackLog(event, "SendRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOwnableMetaData contains all meta data concerning the IOwnable contract.
var IOwnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IOwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use IOwnableMetaData.ABI instead.
var IOwnableABI = IOwnableMetaData.ABI

// IOwnable is an auto generated Go binding around an Ethereum contract.
type IOwnable struct {
	IOwnableCaller     // Read-only binding to the contract
	IOwnableTransactor // Write-only binding to the contract
	IOwnableFilterer   // Log filterer for contract events
}

// IOwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOwnableSession struct {
	Contract     *IOwnable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOwnableCallerSession struct {
	Contract *IOwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IOwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOwnableTransactorSession struct {
	Contract     *IOwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IOwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOwnableRaw struct {
	Contract *IOwnable // Generic contract binding to access the raw methods on
}

// IOwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOwnableCallerRaw struct {
	Contract *IOwnableCaller // Generic read-only contract binding to access the raw methods on
}

// IOwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOwnableTransactorRaw struct {
	Contract *IOwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOwnable creates a new instance of IOwnable, bound to a specific deployed contract.
func NewIOwnable(address common.Address, backend bind.ContractBackend) (*IOwnable, error) {
	contract, err := bindIOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOwnable{IOwnableCaller: IOwnableCaller{contract: contract}, IOwnableTransactor: IOwnableTransactor{contract: contract}, IOwnableFilterer: IOwnableFilterer{contract: contract}}, nil
}

// NewIOwnableCaller creates a new read-only instance of IOwnable, bound to a specific deployed contract.
func NewIOwnableCaller(address common.Address, caller bind.ContractCaller) (*IOwnableCaller, error) {
	contract, err := bindIOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnableCaller{contract: contract}, nil
}

// NewIOwnableTransactor creates a new write-only instance of IOwnable, bound to a specific deployed contract.
func NewIOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*IOwnableTransactor, error) {
	contract, err := bindIOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnableTransactor{contract: contract}, nil
}

// NewIOwnableFilterer creates a new log filterer instance of IOwnable, bound to a specific deployed contract.
func NewIOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*IOwnableFilterer, error) {
	contract, err := bindIOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOwnableFilterer{contract: contract}, nil
}

// bindIOwnable binds a generic wrapper to an already deployed contract.
func bindIOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwnable *IOwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOwnable.Contract.IOwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwnable *IOwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwnable.Contract.IOwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwnable *IOwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwnable.Contract.IOwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwnable *IOwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOwnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwnable *IOwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwnable *IOwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwnable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IOwnable *IOwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOwnable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IOwnable *IOwnableSession) Owner() (common.Address, error) {
	return _IOwnable.Contract.Owner(&_IOwnable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IOwnable *IOwnableCallerSession) Owner() (common.Address, error) {
	return _IOwnable.Contract.Owner(&_IOwnable.CallOpts)
}

// ISequencerInboxMetaData contains all meta data concerning the ISequencerInbox contract.
var ISequencerInboxMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keysetHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidateKeyset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"OwnerFunctionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchSequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SequencerBatchData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchSequenceNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxBlockNumber\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structISequencerInbox.TimeBounds\",\"name\":\"timeBounds\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"enumISequencerInbox.BatchDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"}],\"name\":\"SequencerBatchDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keysetHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"keysetBytes\",\"type\":\"bytes\"}],\"name\":\"SetValidKeyset\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"}],\"name\":\"addSequencerL2Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"inboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"invalidateKeysetHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isBatchPoster\",\"type\":\"bool\"}],\"name\":\"setIsBatchPoster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"timeVariation\",\"type\":\"tuple\"}],\"name\":\"setMaxTimeVariation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"keysetBytes\",\"type\":\"bytes\"}],\"name\":\"setValidKeyset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISequencerInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use ISequencerInboxMetaData.ABI instead.
var ISequencerInboxABI = ISequencerInboxMetaData.ABI

// ISequencerInbox is an auto generated Go binding around an Ethereum contract.
type ISequencerInbox struct {
	ISequencerInboxCaller     // Read-only binding to the contract
	ISequencerInboxTransactor // Write-only binding to the contract
	ISequencerInboxFilterer   // Log filterer for contract events
}

// ISequencerInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISequencerInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISequencerInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISequencerInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISequencerInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISequencerInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISequencerInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISequencerInboxSession struct {
	Contract     *ISequencerInbox  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISequencerInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISequencerInboxCallerSession struct {
	Contract *ISequencerInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ISequencerInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISequencerInboxTransactorSession struct {
	Contract     *ISequencerInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ISequencerInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISequencerInboxRaw struct {
	Contract *ISequencerInbox // Generic contract binding to access the raw methods on
}

// ISequencerInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISequencerInboxCallerRaw struct {
	Contract *ISequencerInboxCaller // Generic read-only contract binding to access the raw methods on
}

// ISequencerInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISequencerInboxTransactorRaw struct {
	Contract *ISequencerInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISequencerInbox creates a new instance of ISequencerInbox, bound to a specific deployed contract.
func NewISequencerInbox(address common.Address, backend bind.ContractBackend) (*ISequencerInbox, error) {
	contract, err := bindISequencerInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISequencerInbox{ISequencerInboxCaller: ISequencerInboxCaller{contract: contract}, ISequencerInboxTransactor: ISequencerInboxTransactor{contract: contract}, ISequencerInboxFilterer: ISequencerInboxFilterer{contract: contract}}, nil
}

// NewISequencerInboxCaller creates a new read-only instance of ISequencerInbox, bound to a specific deployed contract.
func NewISequencerInboxCaller(address common.Address, caller bind.ContractCaller) (*ISequencerInboxCaller, error) {
	contract, err := bindISequencerInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISequencerInboxCaller{contract: contract}, nil
}

// NewISequencerInboxTransactor creates a new write-only instance of ISequencerInbox, bound to a specific deployed contract.
func NewISequencerInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*ISequencerInboxTransactor, error) {
	contract, err := bindISequencerInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISequencerInboxTransactor{contract: contract}, nil
}

// NewISequencerInboxFilterer creates a new log filterer instance of ISequencerInbox, bound to a specific deployed contract.
func NewISequencerInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*ISequencerInboxFilterer, error) {
	contract, err := bindISequencerInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISequencerInboxFilterer{contract: contract}, nil
}

// bindISequencerInbox binds a generic wrapper to an already deployed contract.
func bindISequencerInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISequencerInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISequencerInbox *ISequencerInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISequencerInbox.Contract.ISequencerInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISequencerInbox *ISequencerInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.ISequencerInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISequencerInbox *ISequencerInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.ISequencerInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISequencerInbox *ISequencerInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISequencerInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISequencerInbox *ISequencerInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISequencerInbox *ISequencerInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.contract.Transact(opts, method, params...)
}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_ISequencerInbox *ISequencerInboxCaller) BatchCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISequencerInbox.contract.Call(opts, &out, "batchCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_ISequencerInbox *ISequencerInboxSession) BatchCount() (*big.Int, error) {
	return _ISequencerInbox.Contract.BatchCount(&_ISequencerInbox.CallOpts)
}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_ISequencerInbox *ISequencerInboxCallerSession) BatchCount() (*big.Int, error) {
	return _ISequencerInbox.Contract.BatchCount(&_ISequencerInbox.CallOpts)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_ISequencerInbox *ISequencerInboxCaller) InboxAccs(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ISequencerInbox.contract.Call(opts, &out, "inboxAccs", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_ISequencerInbox *ISequencerInboxSession) InboxAccs(index *big.Int) ([32]byte, error) {
	return _ISequencerInbox.Contract.InboxAccs(&_ISequencerInbox.CallOpts, index)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_ISequencerInbox *ISequencerInboxCallerSession) InboxAccs(index *big.Int) ([32]byte, error) {
	return _ISequencerInbox.Contract.InboxAccs(&_ISequencerInbox.CallOpts, index)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x61a93f8f.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_ISequencerInbox *ISequencerInboxTransactor) AddSequencerL2Batch(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _ISequencerInbox.contract.Transact(opts, "addSequencerL2Batch", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x61a93f8f.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_ISequencerInbox *ISequencerInboxSession) AddSequencerL2Batch(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.AddSequencerL2Batch(&_ISequencerInbox.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x61a93f8f.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_ISequencerInbox *ISequencerInboxTransactorSession) AddSequencerL2Batch(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.AddSequencerL2Batch(&_ISequencerInbox.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_ISequencerInbox *ISequencerInboxTransactor) InvalidateKeysetHash(opts *bind.TransactOpts, ksHash [32]byte) (*types.Transaction, error) {
	return _ISequencerInbox.contract.Transact(opts, "invalidateKeysetHash", ksHash)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_ISequencerInbox *ISequencerInboxSession) InvalidateKeysetHash(ksHash [32]byte) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.InvalidateKeysetHash(&_ISequencerInbox.TransactOpts, ksHash)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_ISequencerInbox *ISequencerInboxTransactorSession) InvalidateKeysetHash(ksHash [32]byte) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.InvalidateKeysetHash(&_ISequencerInbox.TransactOpts, ksHash)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster) returns()
func (_ISequencerInbox *ISequencerInboxTransactor) SetIsBatchPoster(opts *bind.TransactOpts, addr common.Address, isBatchPoster bool) (*types.Transaction, error) {
	return _ISequencerInbox.contract.Transact(opts, "setIsBatchPoster", addr, isBatchPoster)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster) returns()
func (_ISequencerInbox *ISequencerInboxSession) SetIsBatchPoster(addr common.Address, isBatchPoster bool) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.SetIsBatchPoster(&_ISequencerInbox.TransactOpts, addr, isBatchPoster)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster) returns()
func (_ISequencerInbox *ISequencerInboxTransactorSession) SetIsBatchPoster(addr common.Address, isBatchPoster bool) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.SetIsBatchPoster(&_ISequencerInbox.TransactOpts, addr, isBatchPoster)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) timeVariation) returns()
func (_ISequencerInbox *ISequencerInboxTransactor) SetMaxTimeVariation(opts *bind.TransactOpts, timeVariation ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _ISequencerInbox.contract.Transact(opts, "setMaxTimeVariation", timeVariation)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) timeVariation) returns()
func (_ISequencerInbox *ISequencerInboxSession) SetMaxTimeVariation(timeVariation ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.SetMaxTimeVariation(&_ISequencerInbox.TransactOpts, timeVariation)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) timeVariation) returns()
func (_ISequencerInbox *ISequencerInboxTransactorSession) SetMaxTimeVariation(timeVariation ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.SetMaxTimeVariation(&_ISequencerInbox.TransactOpts, timeVariation)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_ISequencerInbox *ISequencerInboxTransactor) SetValidKeyset(opts *bind.TransactOpts, keysetBytes []byte) (*types.Transaction, error) {
	return _ISequencerInbox.contract.Transact(opts, "setValidKeyset", keysetBytes)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_ISequencerInbox *ISequencerInboxSession) SetValidKeyset(keysetBytes []byte) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.SetValidKeyset(&_ISequencerInbox.TransactOpts, keysetBytes)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_ISequencerInbox *ISequencerInboxTransactorSession) SetValidKeyset(keysetBytes []byte) (*types.Transaction, error) {
	return _ISequencerInbox.Contract.SetValidKeyset(&_ISequencerInbox.TransactOpts, keysetBytes)
}

// ISequencerInboxInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the ISequencerInbox contract.
type ISequencerInboxInboxMessageDeliveredIterator struct {
	Event *ISequencerInboxInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *ISequencerInboxInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISequencerInboxInboxMessageDelivered)
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
		it.Event = new(ISequencerInboxInboxMessageDelivered)
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
func (it *ISequencerInboxInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISequencerInboxInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISequencerInboxInboxMessageDelivered represents a InboxMessageDelivered event raised by the ISequencerInbox contract.
type ISequencerInboxInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_ISequencerInbox *ISequencerInboxFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*ISequencerInboxInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _ISequencerInbox.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &ISequencerInboxInboxMessageDeliveredIterator{contract: _ISequencerInbox.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_ISequencerInbox *ISequencerInboxFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *ISequencerInboxInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _ISequencerInbox.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISequencerInboxInboxMessageDelivered)
				if err := _ISequencerInbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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
func (_ISequencerInbox *ISequencerInboxFilterer) ParseInboxMessageDelivered(log types.Log) (*ISequencerInboxInboxMessageDelivered, error) {
	event := new(ISequencerInboxInboxMessageDelivered)
	if err := _ISequencerInbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISequencerInboxInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the ISequencerInbox contract.
type ISequencerInboxInboxMessageDeliveredFromOriginIterator struct {
	Event *ISequencerInboxInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *ISequencerInboxInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISequencerInboxInboxMessageDeliveredFromOrigin)
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
		it.Event = new(ISequencerInboxInboxMessageDeliveredFromOrigin)
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
func (it *ISequencerInboxInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISequencerInboxInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISequencerInboxInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the ISequencerInbox contract.
type ISequencerInboxInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_ISequencerInbox *ISequencerInboxFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*ISequencerInboxInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _ISequencerInbox.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &ISequencerInboxInboxMessageDeliveredFromOriginIterator{contract: _ISequencerInbox.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_ISequencerInbox *ISequencerInboxFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *ISequencerInboxInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _ISequencerInbox.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISequencerInboxInboxMessageDeliveredFromOrigin)
				if err := _ISequencerInbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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
func (_ISequencerInbox *ISequencerInboxFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*ISequencerInboxInboxMessageDeliveredFromOrigin, error) {
	event := new(ISequencerInboxInboxMessageDeliveredFromOrigin)
	if err := _ISequencerInbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISequencerInboxInvalidateKeysetIterator is returned from FilterInvalidateKeyset and is used to iterate over the raw logs and unpacked data for InvalidateKeyset events raised by the ISequencerInbox contract.
type ISequencerInboxInvalidateKeysetIterator struct {
	Event *ISequencerInboxInvalidateKeyset // Event containing the contract specifics and raw log

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
func (it *ISequencerInboxInvalidateKeysetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISequencerInboxInvalidateKeyset)
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
		it.Event = new(ISequencerInboxInvalidateKeyset)
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
func (it *ISequencerInboxInvalidateKeysetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISequencerInboxInvalidateKeysetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISequencerInboxInvalidateKeyset represents a InvalidateKeyset event raised by the ISequencerInbox contract.
type ISequencerInboxInvalidateKeyset struct {
	KeysetHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvalidateKeyset is a free log retrieval operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_ISequencerInbox *ISequencerInboxFilterer) FilterInvalidateKeyset(opts *bind.FilterOpts, keysetHash [][32]byte) (*ISequencerInboxInvalidateKeysetIterator, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _ISequencerInbox.contract.FilterLogs(opts, "InvalidateKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return &ISequencerInboxInvalidateKeysetIterator{contract: _ISequencerInbox.contract, event: "InvalidateKeyset", logs: logs, sub: sub}, nil
}

// WatchInvalidateKeyset is a free log subscription operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_ISequencerInbox *ISequencerInboxFilterer) WatchInvalidateKeyset(opts *bind.WatchOpts, sink chan<- *ISequencerInboxInvalidateKeyset, keysetHash [][32]byte) (event.Subscription, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _ISequencerInbox.contract.WatchLogs(opts, "InvalidateKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISequencerInboxInvalidateKeyset)
				if err := _ISequencerInbox.contract.UnpackLog(event, "InvalidateKeyset", log); err != nil {
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
func (_ISequencerInbox *ISequencerInboxFilterer) ParseInvalidateKeyset(log types.Log) (*ISequencerInboxInvalidateKeyset, error) {
	event := new(ISequencerInboxInvalidateKeyset)
	if err := _ISequencerInbox.contract.UnpackLog(event, "InvalidateKeyset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISequencerInboxOwnerFunctionCalledIterator is returned from FilterOwnerFunctionCalled and is used to iterate over the raw logs and unpacked data for OwnerFunctionCalled events raised by the ISequencerInbox contract.
type ISequencerInboxOwnerFunctionCalledIterator struct {
	Event *ISequencerInboxOwnerFunctionCalled // Event containing the contract specifics and raw log

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
func (it *ISequencerInboxOwnerFunctionCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISequencerInboxOwnerFunctionCalled)
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
		it.Event = new(ISequencerInboxOwnerFunctionCalled)
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
func (it *ISequencerInboxOwnerFunctionCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISequencerInboxOwnerFunctionCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISequencerInboxOwnerFunctionCalled represents a OwnerFunctionCalled event raised by the ISequencerInbox contract.
type ISequencerInboxOwnerFunctionCalled struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOwnerFunctionCalled is a free log retrieval operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_ISequencerInbox *ISequencerInboxFilterer) FilterOwnerFunctionCalled(opts *bind.FilterOpts, id []*big.Int) (*ISequencerInboxOwnerFunctionCalledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ISequencerInbox.contract.FilterLogs(opts, "OwnerFunctionCalled", idRule)
	if err != nil {
		return nil, err
	}
	return &ISequencerInboxOwnerFunctionCalledIterator{contract: _ISequencerInbox.contract, event: "OwnerFunctionCalled", logs: logs, sub: sub}, nil
}

// WatchOwnerFunctionCalled is a free log subscription operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_ISequencerInbox *ISequencerInboxFilterer) WatchOwnerFunctionCalled(opts *bind.WatchOpts, sink chan<- *ISequencerInboxOwnerFunctionCalled, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ISequencerInbox.contract.WatchLogs(opts, "OwnerFunctionCalled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISequencerInboxOwnerFunctionCalled)
				if err := _ISequencerInbox.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
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
func (_ISequencerInbox *ISequencerInboxFilterer) ParseOwnerFunctionCalled(log types.Log) (*ISequencerInboxOwnerFunctionCalled, error) {
	event := new(ISequencerInboxOwnerFunctionCalled)
	if err := _ISequencerInbox.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISequencerInboxSequencerBatchDataIterator is returned from FilterSequencerBatchData and is used to iterate over the raw logs and unpacked data for SequencerBatchData events raised by the ISequencerInbox contract.
type ISequencerInboxSequencerBatchDataIterator struct {
	Event *ISequencerInboxSequencerBatchData // Event containing the contract specifics and raw log

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
func (it *ISequencerInboxSequencerBatchDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISequencerInboxSequencerBatchData)
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
		it.Event = new(ISequencerInboxSequencerBatchData)
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
func (it *ISequencerInboxSequencerBatchDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISequencerInboxSequencerBatchDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISequencerInboxSequencerBatchData represents a SequencerBatchData event raised by the ISequencerInbox contract.
type ISequencerInboxSequencerBatchData struct {
	BatchSequenceNumber *big.Int
	Data                []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchData is a free log retrieval operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_ISequencerInbox *ISequencerInboxFilterer) FilterSequencerBatchData(opts *bind.FilterOpts, batchSequenceNumber []*big.Int) (*ISequencerInboxSequencerBatchDataIterator, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}

	logs, sub, err := _ISequencerInbox.contract.FilterLogs(opts, "SequencerBatchData", batchSequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &ISequencerInboxSequencerBatchDataIterator{contract: _ISequencerInbox.contract, event: "SequencerBatchData", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchData is a free log subscription operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_ISequencerInbox *ISequencerInboxFilterer) WatchSequencerBatchData(opts *bind.WatchOpts, sink chan<- *ISequencerInboxSequencerBatchData, batchSequenceNumber []*big.Int) (event.Subscription, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}

	logs, sub, err := _ISequencerInbox.contract.WatchLogs(opts, "SequencerBatchData", batchSequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISequencerInboxSequencerBatchData)
				if err := _ISequencerInbox.contract.UnpackLog(event, "SequencerBatchData", log); err != nil {
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
func (_ISequencerInbox *ISequencerInboxFilterer) ParseSequencerBatchData(log types.Log) (*ISequencerInboxSequencerBatchData, error) {
	event := new(ISequencerInboxSequencerBatchData)
	if err := _ISequencerInbox.contract.UnpackLog(event, "SequencerBatchData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISequencerInboxSequencerBatchDeliveredIterator is returned from FilterSequencerBatchDelivered and is used to iterate over the raw logs and unpacked data for SequencerBatchDelivered events raised by the ISequencerInbox contract.
type ISequencerInboxSequencerBatchDeliveredIterator struct {
	Event *ISequencerInboxSequencerBatchDelivered // Event containing the contract specifics and raw log

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
func (it *ISequencerInboxSequencerBatchDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISequencerInboxSequencerBatchDelivered)
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
		it.Event = new(ISequencerInboxSequencerBatchDelivered)
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
func (it *ISequencerInboxSequencerBatchDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISequencerInboxSequencerBatchDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISequencerInboxSequencerBatchDelivered represents a SequencerBatchDelivered event raised by the ISequencerInbox contract.
type ISequencerInboxSequencerBatchDelivered struct {
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
func (_ISequencerInbox *ISequencerInboxFilterer) FilterSequencerBatchDelivered(opts *bind.FilterOpts, batchSequenceNumber []*big.Int, beforeAcc [][32]byte, afterAcc [][32]byte) (*ISequencerInboxSequencerBatchDeliveredIterator, error) {

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

	logs, sub, err := _ISequencerInbox.contract.FilterLogs(opts, "SequencerBatchDelivered", batchSequenceNumberRule, beforeAccRule, afterAccRule)
	if err != nil {
		return nil, err
	}
	return &ISequencerInboxSequencerBatchDeliveredIterator{contract: _ISequencerInbox.contract, event: "SequencerBatchDelivered", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchDelivered is a free log subscription operation binding the contract event 0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed batchSequenceNumber, bytes32 indexed beforeAcc, bytes32 indexed afterAcc, bytes32 delayedAcc, uint256 afterDelayedMessagesRead, (uint64,uint64,uint64,uint64) timeBounds, uint8 dataLocation)
func (_ISequencerInbox *ISequencerInboxFilterer) WatchSequencerBatchDelivered(opts *bind.WatchOpts, sink chan<- *ISequencerInboxSequencerBatchDelivered, batchSequenceNumber []*big.Int, beforeAcc [][32]byte, afterAcc [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ISequencerInbox.contract.WatchLogs(opts, "SequencerBatchDelivered", batchSequenceNumberRule, beforeAccRule, afterAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISequencerInboxSequencerBatchDelivered)
				if err := _ISequencerInbox.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
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
func (_ISequencerInbox *ISequencerInboxFilterer) ParseSequencerBatchDelivered(log types.Log) (*ISequencerInboxSequencerBatchDelivered, error) {
	event := new(ISequencerInboxSequencerBatchDelivered)
	if err := _ISequencerInbox.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISequencerInboxSetValidKeysetIterator is returned from FilterSetValidKeyset and is used to iterate over the raw logs and unpacked data for SetValidKeyset events raised by the ISequencerInbox contract.
type ISequencerInboxSetValidKeysetIterator struct {
	Event *ISequencerInboxSetValidKeyset // Event containing the contract specifics and raw log

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
func (it *ISequencerInboxSetValidKeysetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISequencerInboxSetValidKeyset)
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
		it.Event = new(ISequencerInboxSetValidKeyset)
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
func (it *ISequencerInboxSetValidKeysetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISequencerInboxSetValidKeysetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISequencerInboxSetValidKeyset represents a SetValidKeyset event raised by the ISequencerInbox contract.
type ISequencerInboxSetValidKeyset struct {
	KeysetHash  [32]byte
	KeysetBytes []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetValidKeyset is a free log retrieval operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_ISequencerInbox *ISequencerInboxFilterer) FilterSetValidKeyset(opts *bind.FilterOpts, keysetHash [][32]byte) (*ISequencerInboxSetValidKeysetIterator, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _ISequencerInbox.contract.FilterLogs(opts, "SetValidKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return &ISequencerInboxSetValidKeysetIterator{contract: _ISequencerInbox.contract, event: "SetValidKeyset", logs: logs, sub: sub}, nil
}

// WatchSetValidKeyset is a free log subscription operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_ISequencerInbox *ISequencerInboxFilterer) WatchSetValidKeyset(opts *bind.WatchOpts, sink chan<- *ISequencerInboxSetValidKeyset, keysetHash [][32]byte) (event.Subscription, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _ISequencerInbox.contract.WatchLogs(opts, "SetValidKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISequencerInboxSetValidKeyset)
				if err := _ISequencerInbox.contract.UnpackLog(event, "SetValidKeyset", log); err != nil {
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
func (_ISequencerInbox *ISequencerInboxFilterer) ParseSetValidKeyset(log types.Log) (*ISequencerInboxSetValidKeyset, error) {
	event := new(ISequencerInboxSetValidKeyset)
	if err := _ISequencerInbox.contract.UnpackLog(event, "SetValidKeyset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxMetaData contains all meta data concerning the Inbox contract.
var InboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDataLength\",\"type\":\"uint256\"}],\"name\":\"DataTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"InsufficientSubmissionCost\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"InsufficientValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"}],\"name\":\"NotAllowedOrigin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOrigin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotRollupOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2CallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"excessFeeRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callValueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"RetryableData\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"AllowListAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseFee\",\"type\":\"uint256\"}],\"name\":\"calculateRetryableSubmissionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2CallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"excessFeeRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callValueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2CallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"excessFeeRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callValueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createRetryableTicketNoRefundAliasRewrite\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"_sequencerInbox\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"user\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"val\",\"type\":\"bool[]\"}],\"name\":\"setAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_allowListEnabled\",\"type\":\"bool\"}],\"name\":\"setAllowListEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2CallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"excessFeeRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callValueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"unsafeCreateRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b506080516121f1610037600039600081816108f1015261110001526121f16000f3fe6080604052600436106101405760003560e01c806367ef3ab8116100b6578063babcc5391161006f578063babcc53914610307578063c474d2c514610337578063e3de72a514610357578063e78cea9214610377578063ee35f327146103a4578063efeadb6d146103c457600080fd5b806367ef3ab81461026c5780636e6e8a6a1461027f5780638456cb59146102925780638a631aa6146102a7578063a66b327d146102c7578063b75436bb146102e757600080fd5b8063439370b111610108578063439370b1146101e6578063485cc955146101ee5780635075788b1461020e5780635c975abb1461022e5780635e91675814610246578063679b6ded1461025957600080fd5b80630f4d14e9146101455780631b871c8d1461016b5780631fe927cf1461017e57806322bd5c1c1461019e5780633f4ba83a146101cf575b600080fd5b6101586101533660046118dd565b6103e4565b6040519081526020015b60405180910390f35b610158610179366004611953565b610471565b34801561018a57600080fd5b506101586101993660046119f7565b610504565b3480156101aa57600080fd5b506066546101bf90600160a01b900460ff1681565b6040519015158152602001610162565b3480156101db57600080fd5b506101e4610622565b005b610158610762565b3480156101fa57600080fd5b506101e4610209366004611a38565b610842565b34801561021a57600080fd5b50610158610229366004611a71565b6109b1565b34801561023a57600080fd5b5060335460ff166101bf565b610158610254366004611aed565b610a60565b610158610267366004611953565b610b0b565b61015861027a366004611b56565b610c33565b61015861028d366004611953565b610ce1565b34801561029e57600080fd5b506101e4610e29565b3480156102b357600080fd5b506101586102c2366004611bc8565b610f66565b3480156102d357600080fd5b506101586102e2366004611c1c565b611007565b3480156102f357600080fd5b506101586103023660046119f7565b61103f565b34801561031357600080fd5b506101bf610322366004611c3e565b60676020526000908152604090205460ff1681565b34801561034357600080fd5b506101e4610352366004611c3e565b6110f5565b34801561036357600080fd5b506101e4610372366004611d46565b6111f7565b34801561038357600080fd5b50606554610397906001600160a01b031681565b6040516101629190611e07565b3480156103b057600080fd5b50606654610397906001600160a01b031681565b3480156103d057600080fd5b506101e46103df366004611e1b565b611477565b60006103f260335460ff1690565b156104185760405162461bcd60e51b815260040161040f90611e36565b60405180910390fd5b606654600160a01b900460ff16801561044157503260009081526067602052604090205460ff16155b156104615732604051630f51ed7160e41b815260040161040f9190611e07565b610469610762565b90505b919050565b600061047f60335460ff1690565b1561049c5760405162461bcd60e51b815260040161040f90611e36565b606654600160a01b900460ff1680156104c557503260009081526067602052604090205460ff16155b156104e55732604051630f51ed7160e41b815260040161040f9190611e07565b6104f68a8a8a8a8a8a8a8a8a610ce1565b9a9950505050505050505050565b600061051260335460ff1690565b1561052f5760405162461bcd60e51b815260040161040f90611e36565b606654600160a01b900460ff16801561055857503260009081526067602052604090205460ff16155b156105785732604051630f51ed7160e41b815260040161040f9190611e07565b3332146105985760405163feb3d07160e01b815260040160405180910390fd5b6201cccc8211156105c857604051634634691b60e01b8152600481018390526201cccc602482015260440161040f565b60006105ed60033386866040516105e0929190611e60565b6040518091039020611656565b60405190915081907fab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c90600090a29392505050565b6065546040805163cb23bcb560e01b815290516000926001600160a01b03169163cb23bcb5916004808301926020929190829003018186803b15801561066757600080fd5b505afa15801561067b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069f9190611e70565b9050336001600160a01b03821614610757576000816001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156106ec57600080fd5b505afa158015610700573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107249190611e70565b9050336001600160a01b0382161461075557338282604051630739600760e01b815260040161040f93929190611e8d565b505b61075f611708565b50565b600061077060335460ff1690565b1561078d5760405162461bcd60e51b815260040161040f90611e36565b606654600160a01b900460ff1680156107b657503260009081526067602052604090205460ff16155b156107d65732604051630f51ed7160e41b815260040161040f9190611e07565b33803b1515806107e65750323314155b156107f957503361111161111160901b01015b6040516bffffffffffffffffffffffff19606083901b16602082015234603482015261083c90600c9033906054015b604051602081830303815290604052611795565b91505090565b600054610100900460ff1661085d5760005460ff1615610861565b303b155b6108c45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161040f565b600054610100900460ff161580156108e6576000805461ffff19166101011790555b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016141561092f5760405162461bcd60e51b815260040161040f90611eb0565b6065546001600160a01b03161561095957604051633bcd329760e21b815260040160405180910390fd5b606580546001600160a01b038086166001600160a01b031990921691909117909155606680546001600160a81b03191691841691909117905561099a611821565b80156109ac576000805461ff00191690555b505050565b60006109bf60335460ff1690565b156109dc5760405162461bcd60e51b815260040161040f90611e36565b606654600160a01b900460ff168015610a0557503260009081526067602052604090205460ff16155b15610a255732604051630f51ed7160e41b815260040161040f9190611e07565b610a5460033360008b8b8b8b6001600160a01b03168b8b8b604051602001610828989796959493929190611efc565b98975050505050505050565b6000610a6e60335460ff1690565b15610a8b5760405162461bcd60e51b815260040161040f90611e36565b606654600160a01b900460ff168015610ab457503260009081526067602052604090205460ff16155b15610ad45732604051630f51ed7160e41b815260040161040f9190611e07565b610b0160073360018989896001600160a01b0316348a8a6040516020016108289796959493929190611f42565b9695505050505050565b6000610b1960335460ff1690565b15610b365760405162461bcd60e51b815260040161040f90611e36565b606654600160a01b900460ff168015610b5f57503260009081526067602052604090205460ff16155b15610b7f5732604051630f51ed7160e41b815260040161040f9190611e07565b610b898486611f97565b610b938a8a611fb6565b610b9d9190611fb6565b341015610be557610bae8486611f97565b610bb88a8a611fb6565b610bc29190611fb6565b604051631c102d6360e21b8152600481019190915234602482015260440161040f565b6001600160a01b0387163b15610c045761111161111160901b01870196505b6001600160a01b0386163b156104e55761111161111160901b01860195506104f68a8a8a8a8a8a8a8a8a610ce1565b6000610c4160335460ff1690565b15610c5e5760405162461bcd60e51b815260040161040f90611e36565b606654600160a01b900460ff168015610c8757503260009081526067602052604090205460ff16155b15610ca75732604051630f51ed7160e41b815260040161040f9190611e07565b610cd660073360008a8a8a8a6001600160a01b0316348b8b604051602001610828989796959493929190611efc565b979650505050505050565b6000610cef60335460ff1690565b15610d0c5760405162461bcd60e51b815260040161040f90611e36565b606654600160a01b900460ff168015610d3557503260009081526067602052604090205460ff16155b15610d555732604051630f51ed7160e41b815260040161040f9190611e07565b8460011480610d645750836001145b15610d9857338a8a348b8b8b8b8b8b8b6040516307c266e360e01b815260040161040f9b9a99989796959493929190611fce565b6000610da48348611007565b905080891015610dd157604051637d6f91c560e11b815260048101829052602481018a905260440161040f565b610e1a6009338d6001600160a01b03168d348e8e6001600160a01b03168e6001600160a01b03168e8e8e8e90508f8f6040516020016108289b9a99989796959493929190612057565b9b9a5050505050505050505050565b6065546040805163cb23bcb560e01b815290516000926001600160a01b03169163cb23bcb5916004808301926020929190829003018186803b158015610e6e57600080fd5b505afa158015610e82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea69190611e70565b9050336001600160a01b03821614610f5e576000816001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610ef357600080fd5b505afa158015610f07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f2b9190611e70565b9050336001600160a01b03821614610f5c57338282604051630739600760e01b815260040161040f93929190611e8d565b505b61075f611852565b6000610f7460335460ff1690565b15610f915760405162461bcd60e51b815260040161040f90611e36565b606654600160a01b900460ff168015610fba57503260009081526067602052604090205460ff16155b15610fda5732604051630f51ed7160e41b815260040161040f9190611e07565b610cd660033360018a8a8a6001600160a01b03168a8a8a6040516020016108289796959493929190611f42565b600081156110155781611017565b485b611022846006611f97565b61102e90610578611fb6565b6110389190611f97565b9392505050565b600061104d60335460ff1690565b1561106a5760405162461bcd60e51b815260040161040f90611e36565b606654600160a01b900460ff16801561109357503260009081526067602052604090205460ff16155b156110b35732604051630f51ed7160e41b815260040161040f9190611e07565b61103860033385858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061179592505050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016141561113e5760405162461bcd60e51b815260040161040f90611eb0565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61038054336001600160a01b0382161461119b57604051631194af8760e11b81523360048201526001600160a01b038216602482015260440161040f565b600360005b8160ff168160ff1610156111c45760008155806111bc816120b1565b9150506111a0565b50506066805460ff60a01b191690555050606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546040805163cb23bcb560e01b815290516000926001600160a01b03169163cb23bcb5916004808301926020929190829003018186803b15801561123c57600080fd5b505afa158015611250573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112749190611e70565b9050336001600160a01b0382161461132c576000816001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156112c157600080fd5b505afa1580156112d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112f99190611e70565b9050336001600160a01b0382161461132a57338282604051630739600760e01b815260040161040f93929190611e8d565b505b815183511461136d5760405162461bcd60e51b815260206004820152600d60248201526c1253959053125117d253941555609a1b604482015260640161040f565b60005b83518110156114715782818151811061138b5761138b6120d1565b6020026020010151606760008684815181106113a9576113a96120d1565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff0219169083151502179055508381815181106113fa576113fa6120d1565b60200260200101516001600160a01b03167fd9739f45a01ce092c5cdb3d68f63d63d21676b1c6c0b4f9cbc6be4cf5449595a84838151811061143e5761143e6120d1565b6020026020010151604051611457911515815260200190565b60405180910390a280611469816120e7565b915050611370565b50505050565b6065546040805163cb23bcb560e01b815290516000926001600160a01b03169163cb23bcb5916004808301926020929190829003018186803b1580156114bc57600080fd5b505afa1580156114d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114f49190611e70565b9050336001600160a01b038216146115ac576000816001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561154157600080fd5b505afa158015611555573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115799190611e70565b9050336001600160a01b038216146115aa57338282604051630739600760e01b815260040161040f93929190611e8d565b505b606660149054906101000a900460ff16151582151514156115fd5760405162461bcd60e51b815260206004820152600b60248201526a1053149150511657d4d15560aa1b604482015260640161040f565b60668054831515600160a01b0260ff60a01b199091161790556040517f16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb39061164a90841515815260200190565b60405180910390a15050565b6065546000906001600160a01b0316638db5993b348661111161111160901b0187016040516001600160e01b031960e086901b16815260ff90921660048301526001600160a01b03166024820152604481018690526064016020604051808303818588803b1580156116c757600080fd5b505af11580156116db573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906117009190612102565b949350505050565b60335460ff166117515760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015260640161040f565b6033805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405161178b9190611e07565b60405180910390a1565b60006201cccc825111156117cb578151604051634634691b60e01b815260048101919091526201cccc602482015260440161040f565b60006117df85858580519060200120611656565b9050807fff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b84604051611811919061211b565b60405180910390a2949350505050565b600054610100900460ff166118485760405162461bcd60e51b815260040161040f90612170565b6118506118aa565b565b60335460ff16156118755760405162461bcd60e51b815260040161040f90611e36565b6033805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861177e3390565b600054610100900460ff166118d15760405162461bcd60e51b815260040161040f90612170565b6033805460ff19169055565b6000602082840312156118ef57600080fd5b5035919050565b6001600160a01b038116811461075f57600080fd5b60008083601f84011261191d57600080fd5b5081356001600160401b0381111561193457600080fd5b60208301915083602082850101111561194c57600080fd5b9250929050565b60008060008060008060008060006101008a8c03121561197257600080fd5b893561197d816118f6565b985060208a0135975060408a0135965060608a013561199b816118f6565b955060808a01356119ab816118f6565b945060a08a0135935060c08a0135925060e08a01356001600160401b038111156119d457600080fd5b6119e08c828d0161190b565b915080935050809150509295985092959850929598565b60008060208385031215611a0a57600080fd5b82356001600160401b03811115611a2057600080fd5b611a2c8582860161190b565b90969095509350505050565b60008060408385031215611a4b57600080fd5b8235611a56816118f6565b91506020830135611a66816118f6565b809150509250929050565b600080600080600080600060c0888a031215611a8c57600080fd5b8735965060208801359550604088013594506060880135611aac816118f6565b93506080880135925060a08801356001600160401b03811115611ace57600080fd5b611ada8a828b0161190b565b989b979a50959850939692959293505050565b600080600080600060808688031215611b0557600080fd5b85359450602086013593506040860135611b1e816118f6565b925060608601356001600160401b03811115611b3957600080fd5b611b458882890161190b565b969995985093965092949392505050565b60008060008060008060a08789031215611b6f57600080fd5b8635955060208701359450604087013593506060870135611b8f816118f6565b925060808701356001600160401b03811115611baa57600080fd5b611bb689828a0161190b565b979a9699509497509295939492505050565b60008060008060008060a08789031215611be157600080fd5b86359550602087013594506040870135611bfa816118f6565b93506060870135925060808701356001600160401b03811115611baa57600080fd5b60008060408385031215611c2f57600080fd5b50508035926020909101359150565b600060208284031215611c5057600080fd5b8135611038816118f6565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715611c9957611c99611c5b565b604052919050565b60006001600160401b03821115611cba57611cba611c5b565b5060051b60200190565b8035801515811461046c57600080fd5b600082601f830112611ce557600080fd5b81356020611cfa611cf583611ca1565b611c71565b82815260059290921b84018101918181019086841115611d1957600080fd5b8286015b84811015611d3b57611d2e81611cc4565b8352918301918301611d1d565b509695505050505050565b60008060408385031215611d5957600080fd5b82356001600160401b0380821115611d7057600080fd5b818501915085601f830112611d8457600080fd5b81356020611d94611cf583611ca1565b82815260059290921b84018101918181019089841115611db357600080fd5b948201945b83861015611dda578535611dcb816118f6565b82529482019490820190611db8565b96505086013592505080821115611df057600080fd5b50611dfd85828601611cd4565b9150509250929050565b6001600160a01b0391909116815260200190565b600060208284031215611e2d57600080fd5b61103882611cc4565b60208082526010908201526f14185d5cd8589b194e881c185d5cd95960821b604082015260600190565b8183823760009101908152919050565b600060208284031215611e8257600080fd5b8151611038816118f6565b6001600160a01b0393841681529183166020830152909116604082015260600190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b60ff60f81b8960f81b168152876001820152866021820152856041820152846061820152836081820152818360a18301376000910160a101908152979650505050505050565b60ff60f81b8860f81b16815286600182015285602182015284604182015283606182015281836081830137600091016081019081529695505050505050565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615611fb157611fb1611f81565b500290565b60008219821115611fc957611fc9611f81565b500190565b6001600160a01b038c811682528b81166020830152604082018b9052606082018a90526080820189905287811660a0830152861660c082015260e0810185905261010081018490526101406101208201819052810182905260006101608385828501376000838501820152601f909301601f19169091019091019b9a5050505050505050505050565b8b81528a60208201528960408201528860608201528760808201528660a08201528560c08201528460e08201528361010082015260006101208385828501376000929093019092019081529b9a5050505050505050505050565b600060ff821660ff8114156120c8576120c8611f81565b60010192915050565b634e487b7160e01b600052603260045260246000fd5b60006000198214156120fb576120fb611f81565b5060010190565b60006020828403121561211457600080fd5b5051919050565b600060208083528351808285015260005b818110156121485785810183015185820160400152820161212c565b8181111561215a576000604083870101525b50601f01601f1916929092016040019392505050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b60608201526080019056fea26469706673582212206bbb670a9790484a77aae3442d28021a335974f227e6fd971280640b736d466264736f6c63430008090033",
}

// InboxABI is the input ABI used to generate the binding from.
// Deprecated: Use InboxMetaData.ABI instead.
var InboxABI = InboxMetaData.ABI

// InboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InboxMetaData.Bin instead.
var InboxBin = InboxMetaData.Bin

// DeployInbox deploys a new Ethereum contract, binding an instance of Inbox to it.
func DeployInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Inbox, error) {
	parsed, err := InboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// Inbox is an auto generated Go binding around an Ethereum contract.
type Inbox struct {
	InboxCaller     // Read-only binding to the contract
	InboxTransactor // Write-only binding to the contract
	InboxFilterer   // Log filterer for contract events
}

// InboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxSession struct {
	Contract     *Inbox            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxCallerSession struct {
	Contract *InboxCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// InboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxTransactorSession struct {
	Contract     *InboxTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxRaw struct {
	Contract *Inbox // Generic contract binding to access the raw methods on
}

// InboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxCallerRaw struct {
	Contract *InboxCaller // Generic read-only contract binding to access the raw methods on
}

// InboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxTransactorRaw struct {
	Contract *InboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInbox creates a new instance of Inbox, bound to a specific deployed contract.
func NewInbox(address common.Address, backend bind.ContractBackend) (*Inbox, error) {
	contract, err := bindInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// NewInboxCaller creates a new read-only instance of Inbox, bound to a specific deployed contract.
func NewInboxCaller(address common.Address, caller bind.ContractCaller) (*InboxCaller, error) {
	contract, err := bindInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxCaller{contract: contract}, nil
}

// NewInboxTransactor creates a new write-only instance of Inbox, bound to a specific deployed contract.
func NewInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxTransactor, error) {
	contract, err := bindInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxTransactor{contract: contract}, nil
}

// NewInboxFilterer creates a new log filterer instance of Inbox, bound to a specific deployed contract.
func NewInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxFilterer, error) {
	contract, err := bindInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxFilterer{contract: contract}, nil
}

// bindInbox binds a generic wrapper to an already deployed contract.
func bindInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.InboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transact(opts, method, params...)
}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_Inbox *InboxCaller) AllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "allowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_Inbox *InboxSession) AllowListEnabled() (bool, error) {
	return _Inbox.Contract.AllowListEnabled(&_Inbox.CallOpts)
}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_Inbox *InboxCallerSession) AllowListEnabled() (bool, error) {
	return _Inbox.Contract.AllowListEnabled(&_Inbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxSession) Bridge() (common.Address, error) {
	return _Inbox.Contract.Bridge(&_Inbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxCallerSession) Bridge() (common.Address, error) {
	return _Inbox.Contract.Bridge(&_Inbox.CallOpts)
}

// CalculateRetryableSubmissionFee is a free data retrieval call binding the contract method 0xa66b327d.
//
// Solidity: function calculateRetryableSubmissionFee(uint256 dataLength, uint256 baseFee) view returns(uint256)
func (_Inbox *InboxCaller) CalculateRetryableSubmissionFee(opts *bind.CallOpts, dataLength *big.Int, baseFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "calculateRetryableSubmissionFee", dataLength, baseFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRetryableSubmissionFee is a free data retrieval call binding the contract method 0xa66b327d.
//
// Solidity: function calculateRetryableSubmissionFee(uint256 dataLength, uint256 baseFee) view returns(uint256)
func (_Inbox *InboxSession) CalculateRetryableSubmissionFee(dataLength *big.Int, baseFee *big.Int) (*big.Int, error) {
	return _Inbox.Contract.CalculateRetryableSubmissionFee(&_Inbox.CallOpts, dataLength, baseFee)
}

// CalculateRetryableSubmissionFee is a free data retrieval call binding the contract method 0xa66b327d.
//
// Solidity: function calculateRetryableSubmissionFee(uint256 dataLength, uint256 baseFee) view returns(uint256)
func (_Inbox *InboxCallerSession) CalculateRetryableSubmissionFee(dataLength *big.Int, baseFee *big.Int) (*big.Int, error) {
	return _Inbox.Contract.CalculateRetryableSubmissionFee(&_Inbox.CallOpts, dataLength, baseFee)
}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address ) view returns(bool)
func (_Inbox *InboxCaller) IsAllowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "isAllowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address ) view returns(bool)
func (_Inbox *InboxSession) IsAllowed(arg0 common.Address) (bool, error) {
	return _Inbox.Contract.IsAllowed(&_Inbox.CallOpts, arg0)
}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address ) view returns(bool)
func (_Inbox *InboxCallerSession) IsAllowed(arg0 common.Address) (bool, error) {
	return _Inbox.Contract.IsAllowed(&_Inbox.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxSession) Paused() (bool, error) {
	return _Inbox.Contract.Paused(&_Inbox.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxCallerSession) Paused() (bool, error) {
	return _Inbox.Contract.Paused(&_Inbox.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_Inbox *InboxCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_Inbox *InboxSession) SequencerInbox() (common.Address, error) {
	return _Inbox.Contract.SequencerInbox(&_Inbox.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_Inbox *InboxCallerSession) SequencerInbox() (common.Address, error) {
	return _Inbox.Contract.SequencerInbox(&_Inbox.CallOpts)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) CreateRetryableTicket(opts *bind.TransactOpts, to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "createRetryableTicket", to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) CreateRetryableTicket(to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.CreateRetryableTicket(&_Inbox.TransactOpts, to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) CreateRetryableTicket(to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.CreateRetryableTicket(&_Inbox.TransactOpts, to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// CreateRetryableTicketNoRefundAliasRewrite is a paid mutator transaction binding the contract method 0x1b871c8d.
//
// Solidity: function createRetryableTicketNoRefundAliasRewrite(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) CreateRetryableTicketNoRefundAliasRewrite(opts *bind.TransactOpts, to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "createRetryableTicketNoRefundAliasRewrite", to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// CreateRetryableTicketNoRefundAliasRewrite is a paid mutator transaction binding the contract method 0x1b871c8d.
//
// Solidity: function createRetryableTicketNoRefundAliasRewrite(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) CreateRetryableTicketNoRefundAliasRewrite(to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.CreateRetryableTicketNoRefundAliasRewrite(&_Inbox.TransactOpts, to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// CreateRetryableTicketNoRefundAliasRewrite is a paid mutator transaction binding the contract method 0x1b871c8d.
//
// Solidity: function createRetryableTicketNoRefundAliasRewrite(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) CreateRetryableTicketNoRefundAliasRewrite(to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.CreateRetryableTicketNoRefundAliasRewrite(&_Inbox.TransactOpts, to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 ) payable returns(uint256)
func (_Inbox *InboxTransactor) DepositEth(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "depositEth", arg0)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 ) payable returns(uint256)
func (_Inbox *InboxSession) DepositEth(arg0 *big.Int) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEth(&_Inbox.TransactOpts, arg0)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 ) payable returns(uint256)
func (_Inbox *InboxTransactorSession) DepositEth(arg0 *big.Int) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEth(&_Inbox.TransactOpts, arg0)
}

// DepositEth0 is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_Inbox *InboxTransactor) DepositEth0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "depositEth0")
}

// DepositEth0 is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_Inbox *InboxSession) DepositEth0() (*types.Transaction, error) {
	return _Inbox.Contract.DepositEth0(&_Inbox.TransactOpts)
}

// DepositEth0 is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_Inbox *InboxTransactorSession) DepositEth0() (*types.Transaction, error) {
	return _Inbox.Contract.DepositEth0(&_Inbox.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address _sequencerInbox) returns()
func (_Inbox *InboxTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "initialize", _bridge, _sequencerInbox)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address _sequencerInbox) returns()
func (_Inbox *InboxSession) Initialize(_bridge common.Address, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Initialize(&_Inbox.TransactOpts, _bridge, _sequencerInbox)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address _sequencerInbox) returns()
func (_Inbox *InboxTransactorSession) Initialize(_bridge common.Address, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Initialize(&_Inbox.TransactOpts, _bridge, _sequencerInbox)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxSession) Pause() (*types.Transaction, error) {
	return _Inbox.Contract.Pause(&_Inbox.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxTransactorSession) Pause() (*types.Transaction, error) {
	return _Inbox.Contract.Pause(&_Inbox.TransactOpts)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_Inbox *InboxTransactor) PostUpgradeInit(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "postUpgradeInit", _bridge)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_Inbox *InboxSession) PostUpgradeInit(_bridge common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.PostUpgradeInit(&_Inbox.TransactOpts, _bridge)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_Inbox *InboxTransactorSession) PostUpgradeInit(_bridge common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.PostUpgradeInit(&_Inbox.TransactOpts, _bridge)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxTransactor) SendContractTransaction(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendContractTransaction", gasLimit, maxFeePerGas, to, value, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxSession) SendContractTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendContractTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, to, value, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxTransactorSession) SendContractTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendContractTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, to, value, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) SendL1FundedContractTransaction(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL1FundedContractTransaction", gasLimit, maxFeePerGas, to, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) SendL1FundedContractTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedContractTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, to, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) SendL1FundedContractTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedContractTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, to, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) SendL1FundedUnsignedTransaction(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL1FundedUnsignedTransaction", gasLimit, maxFeePerGas, nonce, to, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) SendL1FundedUnsignedTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedUnsignedTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) SendL1FundedUnsignedTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedUnsignedTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, data)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_Inbox *InboxSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactor) SendL2MessageFromOrigin(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2MessageFromOrigin", messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_Inbox *InboxSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactorSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxTransactor) SendUnsignedTransaction(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendUnsignedTransaction", gasLimit, maxFeePerGas, nonce, to, value, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxSession) SendUnsignedTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendUnsignedTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, value, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxTransactorSession) SendUnsignedTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendUnsignedTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, value, data)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_Inbox *InboxTransactor) SetAllowList(opts *bind.TransactOpts, user []common.Address, val []bool) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "setAllowList", user, val)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_Inbox *InboxSession) SetAllowList(user []common.Address, val []bool) (*types.Transaction, error) {
	return _Inbox.Contract.SetAllowList(&_Inbox.TransactOpts, user, val)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_Inbox *InboxTransactorSession) SetAllowList(user []common.Address, val []bool) (*types.Transaction, error) {
	return _Inbox.Contract.SetAllowList(&_Inbox.TransactOpts, user, val)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_Inbox *InboxTransactor) SetAllowListEnabled(opts *bind.TransactOpts, _allowListEnabled bool) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "setAllowListEnabled", _allowListEnabled)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_Inbox *InboxSession) SetAllowListEnabled(_allowListEnabled bool) (*types.Transaction, error) {
	return _Inbox.Contract.SetAllowListEnabled(&_Inbox.TransactOpts, _allowListEnabled)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_Inbox *InboxTransactorSession) SetAllowListEnabled(_allowListEnabled bool) (*types.Transaction, error) {
	return _Inbox.Contract.SetAllowListEnabled(&_Inbox.TransactOpts, _allowListEnabled)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxSession) Unpause() (*types.Transaction, error) {
	return _Inbox.Contract.Unpause(&_Inbox.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxTransactorSession) Unpause() (*types.Transaction, error) {
	return _Inbox.Contract.Unpause(&_Inbox.TransactOpts)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) UnsafeCreateRetryableTicket(opts *bind.TransactOpts, to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "unsafeCreateRetryableTicket", to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) UnsafeCreateRetryableTicket(to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.UnsafeCreateRetryableTicket(&_Inbox.TransactOpts, to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) UnsafeCreateRetryableTicket(to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.UnsafeCreateRetryableTicket(&_Inbox.TransactOpts, to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// InboxAllowListAddressSetIterator is returned from FilterAllowListAddressSet and is used to iterate over the raw logs and unpacked data for AllowListAddressSet events raised by the Inbox contract.
type InboxAllowListAddressSetIterator struct {
	Event *InboxAllowListAddressSet // Event containing the contract specifics and raw log

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
func (it *InboxAllowListAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxAllowListAddressSet)
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
		it.Event = new(InboxAllowListAddressSet)
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
func (it *InboxAllowListAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxAllowListAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxAllowListAddressSet represents a AllowListAddressSet event raised by the Inbox contract.
type InboxAllowListAddressSet struct {
	User common.Address
	Val  bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAllowListAddressSet is a free log retrieval operation binding the contract event 0xd9739f45a01ce092c5cdb3d68f63d63d21676b1c6c0b4f9cbc6be4cf5449595a.
//
// Solidity: event AllowListAddressSet(address indexed user, bool val)
func (_Inbox *InboxFilterer) FilterAllowListAddressSet(opts *bind.FilterOpts, user []common.Address) (*InboxAllowListAddressSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "AllowListAddressSet", userRule)
	if err != nil {
		return nil, err
	}
	return &InboxAllowListAddressSetIterator{contract: _Inbox.contract, event: "AllowListAddressSet", logs: logs, sub: sub}, nil
}

// WatchAllowListAddressSet is a free log subscription operation binding the contract event 0xd9739f45a01ce092c5cdb3d68f63d63d21676b1c6c0b4f9cbc6be4cf5449595a.
//
// Solidity: event AllowListAddressSet(address indexed user, bool val)
func (_Inbox *InboxFilterer) WatchAllowListAddressSet(opts *bind.WatchOpts, sink chan<- *InboxAllowListAddressSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "AllowListAddressSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxAllowListAddressSet)
				if err := _Inbox.contract.UnpackLog(event, "AllowListAddressSet", log); err != nil {
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

// ParseAllowListAddressSet is a log parse operation binding the contract event 0xd9739f45a01ce092c5cdb3d68f63d63d21676b1c6c0b4f9cbc6be4cf5449595a.
//
// Solidity: event AllowListAddressSet(address indexed user, bool val)
func (_Inbox *InboxFilterer) ParseAllowListAddressSet(log types.Log) (*InboxAllowListAddressSet, error) {
	event := new(InboxAllowListAddressSet)
	if err := _Inbox.contract.UnpackLog(event, "AllowListAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxAllowListEnabledUpdatedIterator is returned from FilterAllowListEnabledUpdated and is used to iterate over the raw logs and unpacked data for AllowListEnabledUpdated events raised by the Inbox contract.
type InboxAllowListEnabledUpdatedIterator struct {
	Event *InboxAllowListEnabledUpdated // Event containing the contract specifics and raw log

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
func (it *InboxAllowListEnabledUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxAllowListEnabledUpdated)
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
		it.Event = new(InboxAllowListEnabledUpdated)
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
func (it *InboxAllowListEnabledUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxAllowListEnabledUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxAllowListEnabledUpdated represents a AllowListEnabledUpdated event raised by the Inbox contract.
type InboxAllowListEnabledUpdated struct {
	IsEnabled bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAllowListEnabledUpdated is a free log retrieval operation binding the contract event 0x16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3.
//
// Solidity: event AllowListEnabledUpdated(bool isEnabled)
func (_Inbox *InboxFilterer) FilterAllowListEnabledUpdated(opts *bind.FilterOpts) (*InboxAllowListEnabledUpdatedIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "AllowListEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return &InboxAllowListEnabledUpdatedIterator{contract: _Inbox.contract, event: "AllowListEnabledUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowListEnabledUpdated is a free log subscription operation binding the contract event 0x16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3.
//
// Solidity: event AllowListEnabledUpdated(bool isEnabled)
func (_Inbox *InboxFilterer) WatchAllowListEnabledUpdated(opts *bind.WatchOpts, sink chan<- *InboxAllowListEnabledUpdated) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "AllowListEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxAllowListEnabledUpdated)
				if err := _Inbox.contract.UnpackLog(event, "AllowListEnabledUpdated", log); err != nil {
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

// ParseAllowListEnabledUpdated is a log parse operation binding the contract event 0x16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3.
//
// Solidity: event AllowListEnabledUpdated(bool isEnabled)
func (_Inbox *InboxFilterer) ParseAllowListEnabledUpdated(log types.Log) (*InboxAllowListEnabledUpdated, error) {
	event := new(InboxAllowListEnabledUpdated)
	if err := _Inbox.contract.UnpackLog(event, "AllowListEnabledUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the Inbox contract.
type InboxInboxMessageDeliveredIterator struct {
	Event *InboxInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *InboxInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxInboxMessageDelivered)
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
		it.Event = new(InboxInboxMessageDelivered)
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
func (it *InboxInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxInboxMessageDelivered represents a InboxMessageDelivered event raised by the Inbox contract.
type InboxInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_Inbox *InboxFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxInboxMessageDeliveredIterator{contract: _Inbox.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_Inbox *InboxFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *InboxInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxInboxMessageDelivered)
				if err := _Inbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseInboxMessageDelivered(log types.Log) (*InboxInboxMessageDelivered, error) {
	event := new(InboxInboxMessageDelivered)
	if err := _Inbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the Inbox contract.
type InboxInboxMessageDeliveredFromOriginIterator struct {
	Event *InboxInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *InboxInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxInboxMessageDeliveredFromOrigin)
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
		it.Event = new(InboxInboxMessageDeliveredFromOrigin)
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
func (it *InboxInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the Inbox contract.
type InboxInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_Inbox *InboxFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxInboxMessageDeliveredFromOriginIterator{contract: _Inbox.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_Inbox *InboxFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *InboxInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxInboxMessageDeliveredFromOrigin)
				if err := _Inbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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
func (_Inbox *InboxFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*InboxInboxMessageDeliveredFromOrigin, error) {
	event := new(InboxInboxMessageDeliveredFromOrigin)
	if err := _Inbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Inbox contract.
type InboxPausedIterator struct {
	Event *InboxPaused // Event containing the contract specifics and raw log

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
func (it *InboxPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxPaused)
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
		it.Event = new(InboxPaused)
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
func (it *InboxPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxPaused represents a Paused event raised by the Inbox contract.
type InboxPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Inbox *InboxFilterer) FilterPaused(opts *bind.FilterOpts) (*InboxPausedIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &InboxPausedIterator{contract: _Inbox.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Inbox *InboxFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *InboxPaused) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxPaused)
				if err := _Inbox.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Inbox *InboxFilterer) ParsePaused(log types.Log) (*InboxPaused, error) {
	event := new(InboxPaused)
	if err := _Inbox.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Inbox contract.
type InboxUnpausedIterator struct {
	Event *InboxUnpaused // Event containing the contract specifics and raw log

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
func (it *InboxUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxUnpaused)
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
		it.Event = new(InboxUnpaused)
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
func (it *InboxUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxUnpaused represents a Unpaused event raised by the Inbox contract.
type InboxUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Inbox *InboxFilterer) FilterUnpaused(opts *bind.FilterOpts) (*InboxUnpausedIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &InboxUnpausedIterator{contract: _Inbox.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Inbox *InboxFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *InboxUnpaused) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxUnpaused)
				if err := _Inbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Inbox *InboxFilterer) ParseUnpaused(log types.Log) (*InboxUnpaused, error) {
	event := new(InboxUnpaused)
	if err := _Inbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessagesMetaData contains all meta data concerning the Messages contract.
var MessagesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204d366c3cf2631726630bd6d5f2cfea4f84d0863d919a483d76fab607f8a89a1c64736f6c63430008090033",
}

// MessagesABI is the input ABI used to generate the binding from.
// Deprecated: Use MessagesMetaData.ABI instead.
var MessagesABI = MessagesMetaData.ABI

// MessagesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessagesMetaData.Bin instead.
var MessagesBin = MessagesMetaData.Bin

// DeployMessages deploys a new Ethereum contract, binding an instance of Messages to it.
func DeployMessages(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Messages, error) {
	parsed, err := MessagesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessagesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Messages{MessagesCaller: MessagesCaller{contract: contract}, MessagesTransactor: MessagesTransactor{contract: contract}, MessagesFilterer: MessagesFilterer{contract: contract}}, nil
}

// Messages is an auto generated Go binding around an Ethereum contract.
type Messages struct {
	MessagesCaller     // Read-only binding to the contract
	MessagesTransactor // Write-only binding to the contract
	MessagesFilterer   // Log filterer for contract events
}

// MessagesCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessagesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessagesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessagesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessagesSession struct {
	Contract     *Messages         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessagesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessagesCallerSession struct {
	Contract *MessagesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MessagesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessagesTransactorSession struct {
	Contract     *MessagesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MessagesRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessagesRaw struct {
	Contract *Messages // Generic contract binding to access the raw methods on
}

// MessagesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessagesCallerRaw struct {
	Contract *MessagesCaller // Generic read-only contract binding to access the raw methods on
}

// MessagesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessagesTransactorRaw struct {
	Contract *MessagesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessages creates a new instance of Messages, bound to a specific deployed contract.
func NewMessages(address common.Address, backend bind.ContractBackend) (*Messages, error) {
	contract, err := bindMessages(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Messages{MessagesCaller: MessagesCaller{contract: contract}, MessagesTransactor: MessagesTransactor{contract: contract}, MessagesFilterer: MessagesFilterer{contract: contract}}, nil
}

// NewMessagesCaller creates a new read-only instance of Messages, bound to a specific deployed contract.
func NewMessagesCaller(address common.Address, caller bind.ContractCaller) (*MessagesCaller, error) {
	contract, err := bindMessages(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesCaller{contract: contract}, nil
}

// NewMessagesTransactor creates a new write-only instance of Messages, bound to a specific deployed contract.
func NewMessagesTransactor(address common.Address, transactor bind.ContractTransactor) (*MessagesTransactor, error) {
	contract, err := bindMessages(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesTransactor{contract: contract}, nil
}

// NewMessagesFilterer creates a new log filterer instance of Messages, bound to a specific deployed contract.
func NewMessagesFilterer(address common.Address, filterer bind.ContractFilterer) (*MessagesFilterer, error) {
	contract, err := bindMessages(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessagesFilterer{contract: contract}, nil
}

// bindMessages binds a generic wrapper to an already deployed contract.
func bindMessages(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messages *MessagesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Messages.Contract.MessagesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messages *MessagesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messages.Contract.MessagesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messages *MessagesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messages.Contract.MessagesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messages *MessagesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Messages.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messages *MessagesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messages.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messages *MessagesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messages.Contract.contract.Transact(opts, method, params...)
}

// OutboxMetaData contains all meta data concerning the Outbox contract.
var OutboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"AlreadySpent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeCallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxProofLength\",\"type\":\"uint256\"}],\"name\":\"MerkleProofTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"NotRollup\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxIndex\",\"type\":\"uint256\"}],\"name\":\"PathNotMinimal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proofLength\",\"type\":\"uint256\"}],\"name\":\"ProofTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SimulationOnlyEntrypoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"UnknownRoot\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"zero\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionIndex\",\"type\":\"uint256\"}],\"name\":\"OutBoxTransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"}],\"name\":\"SendRootUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OUTBOX_VERSION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"calculateItemHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"item\",\"type\":\"bytes32\"}],\"name\":\"calculateMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1Block\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeTransactionSimulation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"isSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1BatchNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1EthBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1OutputId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ToL1Timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"roots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spent\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2BlockHash\",\"type\":\"bytes32\"}],\"name\":\"updateSendRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b506080516113a6610030600039600061050301526113a66000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80639f0c04bf116100925780639f0c04bf146101aa578063a04cee60146101bd578063ae6dead7146101d0578063b0f30537146101f0578063c4d66de8146101f8578063c75184df1461020b578063cb23bcb51461022b578063d5b5cc231461023e578063e78cea921461025e57600080fd5b80627436d3146100fa57806308635a95146101205780631198527114610135578063288e5b101461013c578063465477901461014f5780635a129efe1461015757806372f2a8c71461017a57806380648b02146101825780638515bc6a146101a2575b600080fd5b61010d610108366004610cc2565b610271565b6040519081526020015b60405180910390f35b61013361012e366004610de2565b6102ae565b005b600061010d565b61013361014a366004610ed6565b610321565b61010d61035c565b61016a610165366004610f71565b61038b565b6040519015158152602001610117565b61010d6103a8565b61018a6103c5565b6040516001600160a01b039091168152602001610117565b61010d6103e6565b61010d6101b8366004610f8a565b61040e565b6101336101cb366004611018565b610453565b61010d6101de366004610f71565b60036020526000908152604090205481565b61010d6104d7565b61013361020636600461103a565b6104f8565b610213600281565b6040516001600160801b039091168152602001610117565b60005461018a906001600160a01b031681565b61010d61024c366004610f71565b60026020526000908152604090205481565b60015461018a906001600160a01b031681565b60006102a684848460405160200161028b91815260200190565b604051602081830303815290604052805190602001206106b8565b949350505050565b60006102c0898989898989898961040e565b90506103028c8c808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508e92508591506107759050565b6103138a8a8a8a8a8a8a8a8a61087e565b505050505050505050505050565b331561034057604051630e13b69d60e01b815260040160405180910390fd5b61035189898989898989898961087e565b505050505050505050565b6004546000906001600160801b039081169081141561037d57600091505090565b6001600160801b0316919050565b600080600061039984610b68565b92509250506102a68282610ba5565b6006546000906000198114156103c057506000919050565b919050565b6007546000906001600160a01b03908116908114156103c057600091505090565b6004546000906001600160801b03600160801b90910481169081141561037d57600091505090565b6000888888888888888860405160200161042f98979695949392919061105e565b60405160208183030381529060405280519060200120905098975050505050505050565b6000546001600160a01b0316331461049857600054604051630e4cf1bf60e21b81523360048201526001600160a01b0390911660248201526044015b60405180910390fd5b60008281526003602052604080822083905551829184917fb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f67489190a35050565b6005546000906001600160801b039081169081141561037d57600091505090565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156105865760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b606482015260840161048f565b6001546001600160a01b0316156105b057604051633bcd329760e21b815260040160405180910390fd5b6040805160a0810182526001600160801b038082526020808301829052828401829052600019606084018190526001600160a01b0360809094018490526004818155600580546001600160801b031916909417909355600655600780546001600160a01b03199081168517909155600180549487169490911684179055835163cb23bcb560e01b81529351929363cb23bcb593818401939091829003018186803b15801561065d57600080fd5b505afa158015610671573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061069591906110b7565b600080546001600160a01b0319166001600160a01b039290921691909117905550565b82516000906101008111156106eb57604051637ed6198f60e11b815260048101829052610100602482015260440161048f565b8260005b8281101561076b57600087828151811061070b5761070b6110d4565b602002602001015190506002876107229190611100565b61073a5782600052806020526040600020925061074a565b8060005282602052604060002092505b61075560028861112a565b96505080806107639061113e565b9150506106ef565b5095945050505050565b61010083511061079d57825160405163ab6a068360e01b815260040161048f91815260200190565b82516107aa90600261123d565b82106107e15781835160026107bf919061123d565b604051630b8a724b60e01b81526004810192909252602482015260440161048f565b60006107ee848484610271565b600081815260036020526040902054909150610820576040516310e61af960e31b81526004810182905260240161048f565b600080600061082e86610b68565b92509250925061083e8282610ba5565b1561085f57604051639715b8d360e01b81526004810187905260240161048f565b600092835260026020526040909220600190911b909117905550505050565b6000886001600160a01b0316886001600160a01b03167f20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab189648c6040516108c591815260200190565b60405180910390a4600060046040518060a00160405290816000820160009054906101000a90046001600160801b03166001600160801b03166001600160801b031681526020016000820160109054906101000a90046001600160801b03166001600160801b03166001600160801b031681526020016001820160009054906101000a90046001600160801b03166001600160801b03166001600160801b03168152602001600282015481526020016003820160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152505090506040518060a00160405280886001600160801b03168152602001876001600160801b03168152602001866001600160801b031681526020018b60001b81526020018a6001600160a01b0316815250600460008201518160000160006101000a8154816001600160801b0302191690836001600160801b0316021790555060208201518160000160106101000a8154816001600160801b0302191690836001600160801b0316021790555060408201518160010160006101000a8154816001600160801b0302191690836001600160801b031602179055506060820151816002015560808201518160030160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550905050610af6888585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610bb492505050565b805160208201516001600160801b03918216600160801b91831691909102176004556040820151600580546001600160801b03191691909216179055606081015160065560800151600780546001600160a01b0319166001600160a01b03909216919091179055505050505050505050565b6000808080610b7860ff8661112a565b90506000610b8760ff87611100565b60008381526002602052604090205492979096509194509092505050565b80821c60011615155b92915050565b600154604051639e5d4c4960e01b815260009182916001600160a01b0390911690639e5d4c4990610bed90889088908890600401611279565b600060405180830381600087803b158015610c0757600080fd5b505af1158015610c1b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c4391908101906112c2565b9150915081610c7557805115610c5c5780518082602001fd5b604051631bb7daad60e11b815260040160405180910390fd5b5050505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715610cba57610cba610c7c565b604052919050565b600080600060608486031215610cd757600080fd5b83356001600160401b0380821115610cee57600080fd5b818601915086601f830112610d0257600080fd5b8135602082821115610d1657610d16610c7c565b8160051b9250610d27818401610c92565b828152928401810192818101908a851115610d4157600080fd5b948201945b84861015610d5f57853582529482019490820190610d46565b9a918901359950506040909701359695505050505050565b6001600160a01b0381168114610d8c57600080fd5b50565b80356103c081610d77565b60008083601f840112610dac57600080fd5b5081356001600160401b03811115610dc357600080fd5b602083019150836020828501011115610ddb57600080fd5b9250929050565b60008060008060008060008060008060006101208c8e031215610e0457600080fd5b8b356001600160401b0380821115610e1b57600080fd5b818e0191508e601f830112610e2f57600080fd5b813581811115610e3e57600080fd5b8f60208260051b8501011115610e5357600080fd5b60208381019e50909c508e01359a50610e6e60408f01610d8f565b9950610e7c60608f01610d8f565b985060808e0135975060a08e0135965060c08e0135955060e08e013594506101008e0135915080821115610eaf57600080fd5b50610ebc8e828f01610d9a565b915080935050809150509295989b509295989b9093969950565b60008060008060008060008060006101008a8c031215610ef557600080fd5b8935985060208a0135610f0781610d77565b975060408a0135610f1781610d77565b965060608a0135955060808a0135945060a08a0135935060c08a0135925060e08a01356001600160401b03811115610f4e57600080fd5b610f5a8c828d01610d9a565b915080935050809150509295985092959850929598565b600060208284031215610f8357600080fd5b5035919050565b60008060008060008060008060e0898b031215610fa657600080fd5b8835610fb181610d77565b97506020890135610fc181610d77565b965060408901359550606089013594506080890135935060a0890135925060c08901356001600160401b03811115610ff857600080fd5b6110048b828c01610d9a565b999c989b5096995094979396929594505050565b6000806040838503121561102b57600080fd5b50508035926020909101359150565b60006020828403121561104c57600080fd5b813561105781610d77565b9392505050565b60006bffffffffffffffffffffffff19808b60601b168352808a60601b16601484015250876028830152866048830152856068830152846088830152828460a8840137506000910160a801908152979650505050505050565b6000602082840312156110c957600080fd5b815161105781610d77565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601260045260246000fd5b60008261110f5761110f6110ea565b500690565b634e487b7160e01b600052601160045260246000fd5b600082611139576111396110ea565b500490565b600060001982141561115257611152611114565b5060010190565b600181815b8085111561119457816000190482111561117a5761117a611114565b8085161561118757918102915b93841c939080029061115e565b509250929050565b6000826111ab57506001610bae565b816111b857506000610bae565b81600181146111ce57600281146111d8576111f4565b6001915050610bae565b60ff8411156111e9576111e9611114565b50506001821b610bae565b5060208310610133831016604e8410600b8410161715611217575081810a610bae565b6112218383611159565b806000190482111561123557611235611114565b029392505050565b6000611057838361119c565b60005b8381101561126457818101518382015260200161124c565b83811115611273576000848401525b50505050565b60018060a01b038416815282602082015260606040820152600082518060608401526112ac816080850160208701611249565b601f01601f191691909101608001949350505050565b600080604083850312156112d557600080fd5b825180151581146112e557600080fd5b60208401519092506001600160401b038082111561130257600080fd5b818501915085601f83011261131657600080fd5b81518181111561132857611328610c7c565b61133b601f8201601f1916602001610c92565b915080825286602082850101111561135257600080fd5b611363816020840160208601611249565b508092505050925092905056fea2646970667358221220082257a77f6298c774ad3cbd71d830ba15da2dc6cbd281d196234b1ad8c0ea2464736f6c63430008090033",
}

// OutboxABI is the input ABI used to generate the binding from.
// Deprecated: Use OutboxMetaData.ABI instead.
var OutboxABI = OutboxMetaData.ABI

// OutboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OutboxMetaData.Bin instead.
var OutboxBin = OutboxMetaData.Bin

// DeployOutbox deploys a new Ethereum contract, binding an instance of Outbox to it.
func DeployOutbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Outbox, error) {
	parsed, err := OutboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OutboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Outbox{OutboxCaller: OutboxCaller{contract: contract}, OutboxTransactor: OutboxTransactor{contract: contract}, OutboxFilterer: OutboxFilterer{contract: contract}}, nil
}

// Outbox is an auto generated Go binding around an Ethereum contract.
type Outbox struct {
	OutboxCaller     // Read-only binding to the contract
	OutboxTransactor // Write-only binding to the contract
	OutboxFilterer   // Log filterer for contract events
}

// OutboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutboxSession struct {
	Contract     *Outbox           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutboxCallerSession struct {
	Contract *OutboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OutboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutboxTransactorSession struct {
	Contract     *OutboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutboxRaw struct {
	Contract *Outbox // Generic contract binding to access the raw methods on
}

// OutboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutboxCallerRaw struct {
	Contract *OutboxCaller // Generic read-only contract binding to access the raw methods on
}

// OutboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutboxTransactorRaw struct {
	Contract *OutboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutbox creates a new instance of Outbox, bound to a specific deployed contract.
func NewOutbox(address common.Address, backend bind.ContractBackend) (*Outbox, error) {
	contract, err := bindOutbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Outbox{OutboxCaller: OutboxCaller{contract: contract}, OutboxTransactor: OutboxTransactor{contract: contract}, OutboxFilterer: OutboxFilterer{contract: contract}}, nil
}

// NewOutboxCaller creates a new read-only instance of Outbox, bound to a specific deployed contract.
func NewOutboxCaller(address common.Address, caller bind.ContractCaller) (*OutboxCaller, error) {
	contract, err := bindOutbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxCaller{contract: contract}, nil
}

// NewOutboxTransactor creates a new write-only instance of Outbox, bound to a specific deployed contract.
func NewOutboxTransactor(address common.Address, transactor bind.ContractTransactor) (*OutboxTransactor, error) {
	contract, err := bindOutbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutboxTransactor{contract: contract}, nil
}

// NewOutboxFilterer creates a new log filterer instance of Outbox, bound to a specific deployed contract.
func NewOutboxFilterer(address common.Address, filterer bind.ContractFilterer) (*OutboxFilterer, error) {
	contract, err := bindOutbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutboxFilterer{contract: contract}, nil
}

// bindOutbox binds a generic wrapper to an already deployed contract.
func bindOutbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OutboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outbox *OutboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Outbox.Contract.OutboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outbox *OutboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outbox.Contract.OutboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outbox *OutboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outbox.Contract.OutboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Outbox *OutboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Outbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Outbox *OutboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Outbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Outbox *OutboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Outbox.Contract.contract.Transact(opts, method, params...)
}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_Outbox *OutboxCaller) OUTBOXVERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "OUTBOX_VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_Outbox *OutboxSession) OUTBOXVERSION() (*big.Int, error) {
	return _Outbox.Contract.OUTBOXVERSION(&_Outbox.CallOpts)
}

// OUTBOXVERSION is a free data retrieval call binding the contract method 0xc75184df.
//
// Solidity: function OUTBOX_VERSION() view returns(uint128)
func (_Outbox *OutboxCallerSession) OUTBOXVERSION() (*big.Int, error) {
	return _Outbox.Contract.OUTBOXVERSION(&_Outbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Outbox *OutboxCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Outbox *OutboxSession) Bridge() (common.Address, error) {
	return _Outbox.Contract.Bridge(&_Outbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Outbox *OutboxCallerSession) Bridge() (common.Address, error) {
	return _Outbox.Contract.Bridge(&_Outbox.CallOpts)
}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_Outbox *OutboxCaller) CalculateItemHash(opts *bind.CallOpts, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "calculateItemHash", l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_Outbox *OutboxSession) CalculateItemHash(l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	return _Outbox.Contract.CalculateItemHash(&_Outbox.CallOpts, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// CalculateItemHash is a free data retrieval call binding the contract method 0x9f0c04bf.
//
// Solidity: function calculateItemHash(address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) pure returns(bytes32)
func (_Outbox *OutboxCallerSession) CalculateItemHash(l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) ([32]byte, error) {
	return _Outbox.Contract.CalculateItemHash(&_Outbox.CallOpts, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_Outbox *OutboxCaller) CalculateMerkleRoot(opts *bind.CallOpts, proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "calculateMerkleRoot", proof, path, item)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_Outbox *OutboxSession) CalculateMerkleRoot(proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	return _Outbox.Contract.CalculateMerkleRoot(&_Outbox.CallOpts, proof, path, item)
}

// CalculateMerkleRoot is a free data retrieval call binding the contract method 0x007436d3.
//
// Solidity: function calculateMerkleRoot(bytes32[] proof, uint256 path, bytes32 item) pure returns(bytes32)
func (_Outbox *OutboxCallerSession) CalculateMerkleRoot(proof [][32]byte, path *big.Int, item [32]byte) ([32]byte, error) {
	return _Outbox.Contract.CalculateMerkleRoot(&_Outbox.CallOpts, proof, path, item)
}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 index) view returns(bool)
func (_Outbox *OutboxCaller) IsSpent(opts *bind.CallOpts, index *big.Int) (bool, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "isSpent", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 index) view returns(bool)
func (_Outbox *OutboxSession) IsSpent(index *big.Int) (bool, error) {
	return _Outbox.Contract.IsSpent(&_Outbox.CallOpts, index)
}

// IsSpent is a free data retrieval call binding the contract method 0x5a129efe.
//
// Solidity: function isSpent(uint256 index) view returns(bool)
func (_Outbox *OutboxCallerSession) IsSpent(index *big.Int) (bool, error) {
	return _Outbox.Contract.IsSpent(&_Outbox.CallOpts, index)
}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() pure returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1BatchNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1BatchNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() pure returns(uint256)
func (_Outbox *OutboxSession) L2ToL1BatchNum() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1BatchNum(&_Outbox.CallOpts)
}

// L2ToL1BatchNum is a free data retrieval call binding the contract method 0x11985271.
//
// Solidity: function l2ToL1BatchNum() pure returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1BatchNum() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1BatchNum(&_Outbox.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1Block(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Block")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxSession) L2ToL1Block() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Block(&_Outbox.CallOpts)
}

// L2ToL1Block is a free data retrieval call binding the contract method 0x46547790.
//
// Solidity: function l2ToL1Block() view returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1Block() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Block(&_Outbox.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1EthBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1EthBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_Outbox *OutboxSession) L2ToL1EthBlock() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1EthBlock(&_Outbox.CallOpts)
}

// L2ToL1EthBlock is a free data retrieval call binding the contract method 0x8515bc6a.
//
// Solidity: function l2ToL1EthBlock() view returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1EthBlock() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1EthBlock(&_Outbox.CallOpts)
}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_Outbox *OutboxCaller) L2ToL1OutputId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1OutputId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_Outbox *OutboxSession) L2ToL1OutputId() ([32]byte, error) {
	return _Outbox.Contract.L2ToL1OutputId(&_Outbox.CallOpts)
}

// L2ToL1OutputId is a free data retrieval call binding the contract method 0x72f2a8c7.
//
// Solidity: function l2ToL1OutputId() view returns(bytes32)
func (_Outbox *OutboxCallerSession) L2ToL1OutputId() ([32]byte, error) {
	return _Outbox.Contract.L2ToL1OutputId(&_Outbox.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxCaller) L2ToL1Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxSession) L2ToL1Sender() (common.Address, error) {
	return _Outbox.Contract.L2ToL1Sender(&_Outbox.CallOpts)
}

// L2ToL1Sender is a free data retrieval call binding the contract method 0x80648b02.
//
// Solidity: function l2ToL1Sender() view returns(address)
func (_Outbox *OutboxCallerSession) L2ToL1Sender() (common.Address, error) {
	return _Outbox.Contract.L2ToL1Sender(&_Outbox.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxCaller) L2ToL1Timestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "l2ToL1Timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxSession) L2ToL1Timestamp() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Timestamp(&_Outbox.CallOpts)
}

// L2ToL1Timestamp is a free data retrieval call binding the contract method 0xb0f30537.
//
// Solidity: function l2ToL1Timestamp() view returns(uint256)
func (_Outbox *OutboxCallerSession) L2ToL1Timestamp() (*big.Int, error) {
	return _Outbox.Contract.L2ToL1Timestamp(&_Outbox.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_Outbox *OutboxCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_Outbox *OutboxSession) Rollup() (common.Address, error) {
	return _Outbox.Contract.Rollup(&_Outbox.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_Outbox *OutboxCallerSession) Rollup() (common.Address, error) {
	return _Outbox.Contract.Rollup(&_Outbox.CallOpts)
}

// Roots is a free data retrieval call binding the contract method 0xae6dead7.
//
// Solidity: function roots(bytes32 ) view returns(bytes32)
func (_Outbox *OutboxCaller) Roots(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "roots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Roots is a free data retrieval call binding the contract method 0xae6dead7.
//
// Solidity: function roots(bytes32 ) view returns(bytes32)
func (_Outbox *OutboxSession) Roots(arg0 [32]byte) ([32]byte, error) {
	return _Outbox.Contract.Roots(&_Outbox.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0xae6dead7.
//
// Solidity: function roots(bytes32 ) view returns(bytes32)
func (_Outbox *OutboxCallerSession) Roots(arg0 [32]byte) ([32]byte, error) {
	return _Outbox.Contract.Roots(&_Outbox.CallOpts, arg0)
}

// Spent is a free data retrieval call binding the contract method 0xd5b5cc23.
//
// Solidity: function spent(uint256 ) view returns(bytes32)
func (_Outbox *OutboxCaller) Spent(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Outbox.contract.Call(opts, &out, "spent", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Spent is a free data retrieval call binding the contract method 0xd5b5cc23.
//
// Solidity: function spent(uint256 ) view returns(bytes32)
func (_Outbox *OutboxSession) Spent(arg0 *big.Int) ([32]byte, error) {
	return _Outbox.Contract.Spent(&_Outbox.CallOpts, arg0)
}

// Spent is a free data retrieval call binding the contract method 0xd5b5cc23.
//
// Solidity: function spent(uint256 ) view returns(bytes32)
func (_Outbox *OutboxCallerSession) Spent(arg0 *big.Int) ([32]byte, error) {
	return _Outbox.Contract.Spent(&_Outbox.CallOpts, arg0)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_Outbox *OutboxTransactor) ExecuteTransaction(opts *bind.TransactOpts, proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "executeTransaction", proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_Outbox *OutboxSession) ExecuteTransaction(proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransaction(&_Outbox.TransactOpts, proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0x08635a95.
//
// Solidity: function executeTransaction(bytes32[] proof, uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_Outbox *OutboxTransactorSession) ExecuteTransaction(proof [][32]byte, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransaction(&_Outbox.TransactOpts, proof, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransactionSimulation is a paid mutator transaction binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_Outbox *OutboxTransactor) ExecuteTransactionSimulation(opts *bind.TransactOpts, index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "executeTransactionSimulation", index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransactionSimulation is a paid mutator transaction binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_Outbox *OutboxSession) ExecuteTransactionSimulation(index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransactionSimulation(&_Outbox.TransactOpts, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// ExecuteTransactionSimulation is a paid mutator transaction binding the contract method 0x288e5b10.
//
// Solidity: function executeTransactionSimulation(uint256 index, address l2Sender, address to, uint256 l2Block, uint256 l1Block, uint256 l2Timestamp, uint256 value, bytes data) returns()
func (_Outbox *OutboxTransactorSession) ExecuteTransactionSimulation(index *big.Int, l2Sender common.Address, to common.Address, l2Block *big.Int, l1Block *big.Int, l2Timestamp *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Outbox.Contract.ExecuteTransactionSimulation(&_Outbox.TransactOpts, index, l2Sender, to, l2Block, l1Block, l2Timestamp, value, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_Outbox *OutboxTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "initialize", _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_Outbox *OutboxSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _Outbox.Contract.Initialize(&_Outbox.TransactOpts, _bridge)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _bridge) returns()
func (_Outbox *OutboxTransactorSession) Initialize(_bridge common.Address) (*types.Transaction, error) {
	return _Outbox.Contract.Initialize(&_Outbox.TransactOpts, _bridge)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 root, bytes32 l2BlockHash) returns()
func (_Outbox *OutboxTransactor) UpdateSendRoot(opts *bind.TransactOpts, root [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _Outbox.contract.Transact(opts, "updateSendRoot", root, l2BlockHash)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 root, bytes32 l2BlockHash) returns()
func (_Outbox *OutboxSession) UpdateSendRoot(root [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _Outbox.Contract.UpdateSendRoot(&_Outbox.TransactOpts, root, l2BlockHash)
}

// UpdateSendRoot is a paid mutator transaction binding the contract method 0xa04cee60.
//
// Solidity: function updateSendRoot(bytes32 root, bytes32 l2BlockHash) returns()
func (_Outbox *OutboxTransactorSession) UpdateSendRoot(root [32]byte, l2BlockHash [32]byte) (*types.Transaction, error) {
	return _Outbox.Contract.UpdateSendRoot(&_Outbox.TransactOpts, root, l2BlockHash)
}

// OutboxOutBoxTransactionExecutedIterator is returned from FilterOutBoxTransactionExecuted and is used to iterate over the raw logs and unpacked data for OutBoxTransactionExecuted events raised by the Outbox contract.
type OutboxOutBoxTransactionExecutedIterator struct {
	Event *OutboxOutBoxTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *OutboxOutBoxTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutboxOutBoxTransactionExecuted)
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
		it.Event = new(OutboxOutBoxTransactionExecuted)
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
func (it *OutboxOutBoxTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutboxOutBoxTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutboxOutBoxTransactionExecuted represents a OutBoxTransactionExecuted event raised by the Outbox contract.
type OutboxOutBoxTransactionExecuted struct {
	To               common.Address
	L2Sender         common.Address
	Zero             *big.Int
	TransactionIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOutBoxTransactionExecuted is a free log retrieval operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_Outbox *OutboxFilterer) FilterOutBoxTransactionExecuted(opts *bind.FilterOpts, to []common.Address, l2Sender []common.Address, zero []*big.Int) (*OutboxOutBoxTransactionExecutedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var zeroRule []interface{}
	for _, zeroItem := range zero {
		zeroRule = append(zeroRule, zeroItem)
	}

	logs, sub, err := _Outbox.contract.FilterLogs(opts, "OutBoxTransactionExecuted", toRule, l2SenderRule, zeroRule)
	if err != nil {
		return nil, err
	}
	return &OutboxOutBoxTransactionExecutedIterator{contract: _Outbox.contract, event: "OutBoxTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchOutBoxTransactionExecuted is a free log subscription operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_Outbox *OutboxFilterer) WatchOutBoxTransactionExecuted(opts *bind.WatchOpts, sink chan<- *OutboxOutBoxTransactionExecuted, to []common.Address, l2Sender []common.Address, zero []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var zeroRule []interface{}
	for _, zeroItem := range zero {
		zeroRule = append(zeroRule, zeroItem)
	}

	logs, sub, err := _Outbox.contract.WatchLogs(opts, "OutBoxTransactionExecuted", toRule, l2SenderRule, zeroRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutboxOutBoxTransactionExecuted)
				if err := _Outbox.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
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

// ParseOutBoxTransactionExecuted is a log parse operation binding the contract event 0x20af7f3bbfe38132b8900ae295cd9c8d1914be7052d061a511f3f728dab18964.
//
// Solidity: event OutBoxTransactionExecuted(address indexed to, address indexed l2Sender, uint256 indexed zero, uint256 transactionIndex)
func (_Outbox *OutboxFilterer) ParseOutBoxTransactionExecuted(log types.Log) (*OutboxOutBoxTransactionExecuted, error) {
	event := new(OutboxOutBoxTransactionExecuted)
	if err := _Outbox.contract.UnpackLog(event, "OutBoxTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OutboxSendRootUpdatedIterator is returned from FilterSendRootUpdated and is used to iterate over the raw logs and unpacked data for SendRootUpdated events raised by the Outbox contract.
type OutboxSendRootUpdatedIterator struct {
	Event *OutboxSendRootUpdated // Event containing the contract specifics and raw log

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
func (it *OutboxSendRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutboxSendRootUpdated)
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
		it.Event = new(OutboxSendRootUpdated)
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
func (it *OutboxSendRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutboxSendRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutboxSendRootUpdated represents a SendRootUpdated event raised by the Outbox contract.
type OutboxSendRootUpdated struct {
	BlockHash  [32]byte
	OutputRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSendRootUpdated is a free log retrieval operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed blockHash, bytes32 indexed outputRoot)
func (_Outbox *OutboxFilterer) FilterSendRootUpdated(opts *bind.FilterOpts, blockHash [][32]byte, outputRoot [][32]byte) (*OutboxSendRootUpdatedIterator, error) {

	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}

	logs, sub, err := _Outbox.contract.FilterLogs(opts, "SendRootUpdated", blockHashRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return &OutboxSendRootUpdatedIterator{contract: _Outbox.contract, event: "SendRootUpdated", logs: logs, sub: sub}, nil
}

// WatchSendRootUpdated is a free log subscription operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed blockHash, bytes32 indexed outputRoot)
func (_Outbox *OutboxFilterer) WatchSendRootUpdated(opts *bind.WatchOpts, sink chan<- *OutboxSendRootUpdated, blockHash [][32]byte, outputRoot [][32]byte) (event.Subscription, error) {

	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}
	var outputRootRule []interface{}
	for _, outputRootItem := range outputRoot {
		outputRootRule = append(outputRootRule, outputRootItem)
	}

	logs, sub, err := _Outbox.contract.WatchLogs(opts, "SendRootUpdated", blockHashRule, outputRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutboxSendRootUpdated)
				if err := _Outbox.contract.UnpackLog(event, "SendRootUpdated", log); err != nil {
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

// ParseSendRootUpdated is a log parse operation binding the contract event 0xb4df3847300f076a369cd76d2314b470a1194d9e8a6bb97f1860aee88a5f6748.
//
// Solidity: event SendRootUpdated(bytes32 indexed blockHash, bytes32 indexed outputRoot)
func (_Outbox *OutboxFilterer) ParseSendRootUpdated(log types.Log) (*OutboxSendRootUpdated, error) {
	event := new(OutboxSendRootUpdated)
	if err := _Outbox.contract.UnpackLog(event, "SendRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxMetaData contains all meta data concerning the SequencerInbox contract.
var SequencerInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"AlreadyValidDASKeyset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"BadSequencerNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataNotAuthenticated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDataLength\",\"type\":\"uint256\"}],\"name\":\"DataTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayedBackwards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayedTooFar\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceIncludeBlockTooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceIncludeTimeTooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HadZeroInit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectMessagePreimage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"NoSuchKeyset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBatchPoster\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOrigin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keysetHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidateKeyset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"OwnerFunctionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchSequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SequencerBatchData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchSequenceNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxBlockNumber\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structISequencerInbox.TimeBounds\",\"name\":\"timeBounds\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"enumISequencerInbox.BatchDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"}],\"name\":\"SequencerBatchDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keysetHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"keysetBytes\",\"type\":\"bytes\"}],\"name\":\"SetValidKeyset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DATA_AUTHENTICATED_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HEADER_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"}],\"name\":\"addSequencerL2Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"}],\"name\":\"addSequencerL2BatchFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dasKeySetInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidKeyset\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"creationBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint64[2]\",\"name\":\"l1BlockAndTime\",\"type\":\"uint64[2]\"},{\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"forceInclusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"getKeysetCreationBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"inboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation_\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"invalidateKeysetHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBatchPoster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"isValidKeysetHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTimeVariation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isBatchPoster_\",\"type\":\"bool\"}],\"name\":\"setIsBatchPoster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation_\",\"type\":\"tuple\"}],\"name\":\"setMaxTimeVariation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"keysetBytes\",\"type\":\"bytes\"}],\"name\":\"setValidKeyset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDelayedMessagesRead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b50608051611dc361003060003960006103e80152611dc36000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c80637fa3a40e116100ad578063d9dd67ab11610071578063d9dd67ab146102ca578063e5a358c8146102dd578063e78cea9214610301578063ebea461d14610314578063f19815781461034a57600080fd5b80637fa3a40e1461025d5780638442086014610266578063b31761f814610279578063cb23bcb51461028c578063d1ce8da8146102b757600080fd5b806361a93f8f116100f457806361a93f8f146101af5780636e7df3e7146101c25780636f12b0c9146101d5578063715ea34b146101e857806371c3e6fe1461023a57600080fd5b806306f13056146101315780631637be481461014c5780631f7a92b21461017f578063258f04951461019457806327957a49146101a7575b600080fd5b61013961035d565b6040519081526020015b60405180910390f35b61016f61015a3660046117db565b60009081526008602052604090205460ff1690565b6040519015158152602001610143565b61019261018d36600461180c565b6103dd565b005b6101396101a23660046117db565b610595565b610139602881565b6101926101bd366004611895565b6105fe565b6101926101d036600461190d565b6107dd565b6101926101e3366004611895565b6108e3565b61021b6101f63660046117db565b60086020526000908152604090205460ff81169061010090046001600160401b031682565b6040805192151583526001600160401b03909116602083015201610143565b61016f610248366004611946565b60036020526000908152604090205460ff1681565b61013960005481565b6101926102743660046117db565b6109fd565b61019261028736600461196a565b610b50565b60025461029f906001600160a01b031681565b6040516001600160a01b039091168152602001610143565b6101926102c53660046119dd565b610c50565b6101396102d83660046117db565b610e98565b6102e8600160fe1b81565b6040516001600160f81b03199091168152602001610143565b60015461029f906001600160a01b031681565b60045460055460065460075461032a9392919084565b604080519485526020850193909352918301526060820152608001610143565b610192610358366004611a1e565b610f1b565b600154604080516221048360e21b815290516000926001600160a01b0316916284120c916004808301926020929190829003018186803b1580156103a057600080fd5b505afa1580156103b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d89190611a8e565b905090565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156104705760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b60648201526084015b60405180910390fd5b6001546001600160a01b03161561049a57604051633bcd329760e21b815260040160405180910390fd5b6001600160a01b0382166104c157604051631ad0f74360e01b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0384169081179091556040805163cb23bcb560e01b8152905163cb23bcb591600480820192602092909190829003018186803b15801561051557600080fd5b505afa158015610529573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061054d9190611aa7565b600280546001600160a01b0319166001600160a01b03929092169190911790558035600490815560208201356005556040820135600655606082013560075581905b50505050565b600081815260086020908152604080832081518083019092525460ff81161515825261010090046001600160401b0316918101829052906105eb5760405162f20c5d60e01b815260048101849052602401610467565b602001516001600160401b031692915050565b8060005a3360009081526003602052604090205490915060ff1615801561063057506002546001600160a01b03163314155b1561064e57604051632dd9fc9760e01b815260040160405180910390fd5b60008061065c88888861124e565b91509150600080600080610672868b60006113c1565b93509350935093508c84146106a45760405163ac7411c960e01b815260048101859052602481018e9052604401610467565b80838e600080516020611d4e833981519152858e8a60016040516106cb9493929190611ac4565b60405180910390a48c7ffe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c208d8d604051610705929190611b38565b60405180910390a2505050506001600160a01b0384161591506107d49050576000333214156107315750365b826001600160a01b031663e3db8a49335a61074c9086611b7d565b6040516001600160e01b031960e085901b1681526001600160a01b039092166004830152602482015260448101849052606401602060405180830381600087803b15801561079957600080fd5b505af11580156107ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d19190611b94565b50505b50505050505050565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561082b57600080fd5b505afa15801561083f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108639190611aa7565b6001600160a01b0316336001600160a01b0316146108a557600254604051631194af8760e11b81526104679133916001600160a01b0390911690600401611bb1565b6001600160a01b038216600090815260036020526040808220805460ff191684151517905551600191600080516020611d6e83398151915291a25050565b8060005a90503332146109095760405163feb3d07160e01b815260040160405180910390fd5b3360009081526003602052604090205460ff1661093957604051632dd9fc9760e01b815260040160405180910390fd5b60008061094788888861124e565b9092509050600080808061095c868b8d6113c1565b93509350935093508c841461098e5760405163ac7411c960e01b815260048101859052602481018e9052604401610467565b80838e600080516020611d4e833981519152856000548a60006040516109b79493929190611ac4565b60405180910390a4505050506001600160a01b0384161591506107d4905057600033321415610731575036826001600160a01b031663e3db8a49335a61074c9086611b7d565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a4b57600080fd5b505afa158015610a5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a839190611aa7565b6001600160a01b0316336001600160a01b031614610ac557600254604051631194af8760e11b81526104679133916001600160a01b0390911690600401611bb1565b60008181526008602052604090205460ff16610af65760405162f20c5d60e01b815260048101829052602401610467565b600081815260086020526040808220805460ff191690555182917f5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a91a2604051600390600080516020611d6e83398151915290600090a250565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610b9e57600080fd5b505afa158015610bb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bd69190611aa7565b6001600160a01b0316336001600160a01b031614610c1857600254604051631194af8760e11b81526104679133916001600160a01b0390911690600401611bb1565b80516004556020810151600555604080820151600655606082015160075551600090600080516020611d6e833981519152908290a250565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610c9e57600080fd5b505afa158015610cb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cd69190611aa7565b6001600160a01b0316336001600160a01b031614610d1857600254604051631194af8760e11b81526104679133916001600160a01b0390911690600401611bb1565b60008282604051610d2a929190611bcb565b604051908190038120607f60f91b6020830152602182015260410160408051601f1981840301815291905280516020909101209050600160ff1b8118620100008310610dae5760405162461bcd60e51b81526020600482015260136024820152726b657973657420697320746f6f206c6172676560681b6044820152606401610467565b60008181526008602052604090205460ff1615610de157604051637d17eeed60e11b815260048101829052602401610467565b60408051808201825260018152436001600160401b0390811660208084019182526000868152600890915284902092518354915168ffffffffffffffffff1990921690151568ffffffffffffffff0019161761010091909216021790555181907fabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef65572290610e709087908790611b38565b60405180910390a2604051600290600080516020611d6e83398151915290600090a250505050565b6001546040516316bf557960e01b8152600481018390526000916001600160a01b0316906316bf55799060240160206040518083038186803b158015610edd57600080fd5b505afa158015610ef1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f159190611a8e565b92915050565b6000548611610f3d57604051633eb9f37d60e11b815260040160405180910390fd5b6000610fed8684610f516020890189611bf1565b610f6160408a0160208b01611bf1565b610f6c60018d611b7d565b6040805160f89690961b6001600160f81b03191660208088019190915260609590951b6001600160601b031916602187015260c093841b6001600160c01b031990811660358801529290931b909116603d85015260458401526065830188905260858084018790528151808503909101815260a59093019052815191012090565b60045490915043906110026020880188611bf1565b6001600160401b03166110159190611c1a565b106110335760405163ad3515d960e01b815260040160405180910390fd5b60065442906110486040880160208901611bf1565b6001600160401b031661105b9190611c1a565b106110795760405163c76d17e560e01b815260040160405180910390fd5b60006001881115611111576001546001600160a01b031663d5719dc26110a060028b611b7d565b6040518263ffffffff1660e01b81526004016110be91815260200190565b60206040518083038186803b1580156110d657600080fd5b505afa1580156110ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061110e9190611a8e565b90505b60408051602080820184905281830185905282518083038401815260609092019092528051910120600180546001600160a01b03169063d5719dc290611157908c611b7d565b6040518263ffffffff1660e01b815260040161117591815260200190565b60206040518083038186803b15801561118d57600080fd5b505afa1580156111a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111c59190611a8e565b146111e3576040516313947fd760e01b815260040160405180910390fd5b6000806111ef8a61164d565b91509150600080600080611205868f60006113c1565b9350935093509350808385600080516020611d4e833981519152856000548a60026040516112369493929190611ac4565b60405180910390a45050505050505050505050505050565b60006112586117b4565b84846000611267826028611c1a565b90506201cccc81111561129957604051634634691b60e01b8152600481018290526201cccc6024820152604401610467565b81158015906112d15750600160fe1b8084846000816112ba576112ba611bdb565b9050013560f81c60f81b166001600160f81b031916145b156112ef57604051631f97007f60e01b815260040160405180910390fd5b6021821080159061131d57508282600081811061130e5761130e611bdb565b90910135600160ff1b16151590505b15611371576000611332602160018587611c32565b61133b91611c5c565b60008181526008602052604090205490915060ff1661136f5760405162f20c5d60e01b815260048101829052602401610467565b505b60008061137d88611679565b915091506000828b8b60405160200161139893929190611ca6565b60408051808303601f1901815291905280516020909101209b919a509098505050505050505050565b6000806000806000548610156113ea57604051633eb9f37d60e11b815260040160405180910390fd5b600160009054906101000a90046001600160a01b03166001600160a01b031663eca067ad6040518163ffffffff1660e01b815260040160206040518083038186803b15801561143857600080fd5b505afa15801561144c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114709190611a8e565b8611156114905760405163925f8bd360e01b815260040160405180910390fd5b60015460405163422def4160e11b815260048101899052602481018890526001600160a01b039091169063845bde8290604401608060405180830381600087803b1580156114dd57600080fd5b505af11580156114f1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115159190611cce565b60008a9055929650909450925090508415611644576040805142602082015233606081901b6001600160601b03191692820192909252605481018990526074810186905248609482015260009060b40160408051808303601f190181529082905260015481516020830120637a88b10760e01b84526001600160a01b0386811660048601526024850191909152919350600092911690637a88b10790604401602060405180830381600087803b1580156115ce57600080fd5b505af11580156115e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116069190611a8e565b9050807fff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b836040516116389190611d04565b60405180910390a25050505b93509350935093565b60006116576117b4565b60008061166385611679565b8151602090920191909120969095509350505050565b60606116836117b4565b600061168d611720565b905060008160000151826020015183604001518460600151886040516020016116f595949392919060c095861b6001600160c01b0319908116825294861b8516600882015292851b8416601084015290841b8316601883015290921b16602082015260280190565b6040516020818303038152906040529050602881511461171757611717611d37565b94909350915050565b6117286117b4565b6117306117b4565b600654421115611753576006546117479042611b7d565b6001600160401b031681525b6007546117609042611c1a565b6001600160401b03166020820152600454431115611794576004546117859043611b7d565b6001600160401b031660408201525b6005546117a19043611c1a565b6001600160401b03166060820152919050565b60408051608081018252600080825260208201819052918101829052606081019190915290565b6000602082840312156117ed57600080fd5b5035919050565b6001600160a01b038116811461180957600080fd5b50565b60008082840360a081121561182057600080fd5b833561182b816117f4565b92506080601f198201121561183f57600080fd5b506020830190509250929050565b60008083601f84011261185f57600080fd5b5081356001600160401b0381111561187657600080fd5b60208301915083602082850101111561188e57600080fd5b9250929050565b6000806000806000608086880312156118ad57600080fd5b8535945060208601356001600160401b038111156118ca57600080fd5b6118d68882890161184d565b9095509350506040860135915060608601356118f1816117f4565b809150509295509295909350565b801515811461180957600080fd5b6000806040838503121561192057600080fd5b823561192b816117f4565b9150602083013561193b816118ff565b809150509250929050565b60006020828403121561195857600080fd5b8135611963816117f4565b9392505050565b60006080828403121561197c57600080fd5b604051608081018181106001600160401b03821117156119ac57634e487b7160e01b600052604160045260246000fd5b8060405250823581526020830135602082015260408301356040820152606083013560608201528091505092915050565b600080602083850312156119f057600080fd5b82356001600160401b03811115611a0657600080fd5b611a128582860161184d565b90969095509350505050565b60008060008060008060e08789031215611a3757600080fd5b86359550602087013560ff81168114611a4f57600080fd5b94506080870188811115611a6257600080fd5b60408801945035925060a0870135611a79816117f4565b8092505060c087013590509295509295509295565b600060208284031215611aa057600080fd5b5051919050565b600060208284031215611ab957600080fd5b8151611963816117f4565b600060e0820190508582528460208301526001600160401b038085511660408401528060208601511660608401528060408601511660808401528060608601511660a08401525060038310611b2957634e487b7160e01b600052602160045260246000fd5b8260c083015295945050505050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b634e487b7160e01b600052601160045260246000fd5b600082821015611b8f57611b8f611b67565b500390565b600060208284031215611ba657600080fd5b8151611963816118ff565b6001600160a01b0392831681529116602082015260400190565b8183823760009101908152919050565b634e487b7160e01b600052603260045260246000fd5b600060208284031215611c0357600080fd5b81356001600160401b038116811461196357600080fd5b60008219821115611c2d57611c2d611b67565b500190565b60008085851115611c4257600080fd5b83861115611c4f57600080fd5b5050820193919092039150565b80356020831015610f1557600019602084900360031b1b1692915050565b60005b83811015611c95578181015183820152602001611c7d565b8381111561058f5750506000910152565b60008451611cb8818460208901611c7a565b8201838582376000930192835250909392505050565b60008060008060808587031215611ce457600080fd5b505082516020840151604085015160609095015191969095509092509050565b6020815260008251806020840152611d23816040850160208701611c7a565b601f01601f19169190910160400192915050565b634e487b7160e01b600052600160045260246000fdfe7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7ea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456ea264697066735822122014a287f4e696c6319965c52043e93516cf7a1f7f97cb523167fc12e69064108264736f6c63430008090033",
}

// SequencerInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerInboxMetaData.ABI instead.
var SequencerInboxABI = SequencerInboxMetaData.ABI

// SequencerInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SequencerInboxMetaData.Bin instead.
var SequencerInboxBin = SequencerInboxMetaData.Bin

// DeploySequencerInbox deploys a new Ethereum contract, binding an instance of SequencerInbox to it.
func DeploySequencerInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SequencerInbox, error) {
	parsed, err := SequencerInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SequencerInboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SequencerInbox{SequencerInboxCaller: SequencerInboxCaller{contract: contract}, SequencerInboxTransactor: SequencerInboxTransactor{contract: contract}, SequencerInboxFilterer: SequencerInboxFilterer{contract: contract}}, nil
}

// SequencerInbox is an auto generated Go binding around an Ethereum contract.
type SequencerInbox struct {
	SequencerInboxCaller     // Read-only binding to the contract
	SequencerInboxTransactor // Write-only binding to the contract
	SequencerInboxFilterer   // Log filterer for contract events
}

// SequencerInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerInboxSession struct {
	Contract     *SequencerInbox   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SequencerInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerInboxCallerSession struct {
	Contract *SequencerInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SequencerInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerInboxTransactorSession struct {
	Contract     *SequencerInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SequencerInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerInboxRaw struct {
	Contract *SequencerInbox // Generic contract binding to access the raw methods on
}

// SequencerInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerInboxCallerRaw struct {
	Contract *SequencerInboxCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerInboxTransactorRaw struct {
	Contract *SequencerInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencerInbox creates a new instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInbox(address common.Address, backend bind.ContractBackend) (*SequencerInbox, error) {
	contract, err := bindSequencerInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SequencerInbox{SequencerInboxCaller: SequencerInboxCaller{contract: contract}, SequencerInboxTransactor: SequencerInboxTransactor{contract: contract}, SequencerInboxFilterer: SequencerInboxFilterer{contract: contract}}, nil
}

// NewSequencerInboxCaller creates a new read-only instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInboxCaller(address common.Address, caller bind.ContractCaller) (*SequencerInboxCaller, error) {
	contract, err := bindSequencerInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxCaller{contract: contract}, nil
}

// NewSequencerInboxTransactor creates a new write-only instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerInboxTransactor, error) {
	contract, err := bindSequencerInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxTransactor{contract: contract}, nil
}

// NewSequencerInboxFilterer creates a new log filterer instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerInboxFilterer, error) {
	contract, err := bindSequencerInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxFilterer{contract: contract}, nil
}

// bindSequencerInbox binds a generic wrapper to an already deployed contract.
func bindSequencerInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SequencerInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInbox *SequencerInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInbox.Contract.SequencerInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInbox *SequencerInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SequencerInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInbox *SequencerInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SequencerInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInbox *SequencerInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInbox *SequencerInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInbox *SequencerInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInbox.Contract.contract.Transact(opts, method, params...)
}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInbox *SequencerInboxCaller) DATAAUTHENTICATEDFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "DATA_AUTHENTICATED_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInbox *SequencerInboxSession) DATAAUTHENTICATEDFLAG() ([1]byte, error) {
	return _SequencerInbox.Contract.DATAAUTHENTICATEDFLAG(&_SequencerInbox.CallOpts)
}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInbox *SequencerInboxCallerSession) DATAAUTHENTICATEDFLAG() ([1]byte, error) {
	return _SequencerInbox.Contract.DATAAUTHENTICATEDFLAG(&_SequencerInbox.CallOpts)
}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInbox *SequencerInboxCaller) HEADERLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "HEADER_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInbox *SequencerInboxSession) HEADERLENGTH() (*big.Int, error) {
	return _SequencerInbox.Contract.HEADERLENGTH(&_SequencerInbox.CallOpts)
}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInbox *SequencerInboxCallerSession) HEADERLENGTH() (*big.Int, error) {
	return _SequencerInbox.Contract.HEADERLENGTH(&_SequencerInbox.CallOpts)
}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInbox *SequencerInboxCaller) BatchCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "batchCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInbox *SequencerInboxSession) BatchCount() (*big.Int, error) {
	return _SequencerInbox.Contract.BatchCount(&_SequencerInbox.CallOpts)
}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInbox *SequencerInboxCallerSession) BatchCount() (*big.Int, error) {
	return _SequencerInbox.Contract.BatchCount(&_SequencerInbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInbox *SequencerInboxCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInbox *SequencerInboxSession) Bridge() (common.Address, error) {
	return _SequencerInbox.Contract.Bridge(&_SequencerInbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInbox *SequencerInboxCallerSession) Bridge() (common.Address, error) {
	return _SequencerInbox.Contract.Bridge(&_SequencerInbox.CallOpts)
}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(bool isValidKeyset, uint64 creationBlock)
func (_SequencerInbox *SequencerInboxCaller) DasKeySetInfo(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "dasKeySetInfo", arg0)

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
func (_SequencerInbox *SequencerInboxSession) DasKeySetInfo(arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	return _SequencerInbox.Contract.DasKeySetInfo(&_SequencerInbox.CallOpts, arg0)
}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(bool isValidKeyset, uint64 creationBlock)
func (_SequencerInbox *SequencerInboxCallerSession) DasKeySetInfo(arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	return _SequencerInbox.Contract.DasKeySetInfo(&_SequencerInbox.CallOpts, arg0)
}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInbox *SequencerInboxCaller) GetKeysetCreationBlock(opts *bind.CallOpts, ksHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "getKeysetCreationBlock", ksHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInbox *SequencerInboxSession) GetKeysetCreationBlock(ksHash [32]byte) (*big.Int, error) {
	return _SequencerInbox.Contract.GetKeysetCreationBlock(&_SequencerInbox.CallOpts, ksHash)
}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInbox *SequencerInboxCallerSession) GetKeysetCreationBlock(ksHash [32]byte) (*big.Int, error) {
	return _SequencerInbox.Contract.GetKeysetCreationBlock(&_SequencerInbox.CallOpts, ksHash)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInbox *SequencerInboxCaller) InboxAccs(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "inboxAccs", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInbox *SequencerInboxSession) InboxAccs(index *big.Int) ([32]byte, error) {
	return _SequencerInbox.Contract.InboxAccs(&_SequencerInbox.CallOpts, index)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInbox *SequencerInboxCallerSession) InboxAccs(index *big.Int) ([32]byte, error) {
	return _SequencerInbox.Contract.InboxAccs(&_SequencerInbox.CallOpts, index)
}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInbox *SequencerInboxCaller) IsBatchPoster(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "isBatchPoster", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInbox *SequencerInboxSession) IsBatchPoster(arg0 common.Address) (bool, error) {
	return _SequencerInbox.Contract.IsBatchPoster(&_SequencerInbox.CallOpts, arg0)
}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInbox *SequencerInboxCallerSession) IsBatchPoster(arg0 common.Address) (bool, error) {
	return _SequencerInbox.Contract.IsBatchPoster(&_SequencerInbox.CallOpts, arg0)
}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInbox *SequencerInboxCaller) IsValidKeysetHash(opts *bind.CallOpts, ksHash [32]byte) (bool, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "isValidKeysetHash", ksHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInbox *SequencerInboxSession) IsValidKeysetHash(ksHash [32]byte) (bool, error) {
	return _SequencerInbox.Contract.IsValidKeysetHash(&_SequencerInbox.CallOpts, ksHash)
}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInbox *SequencerInboxCallerSession) IsValidKeysetHash(ksHash [32]byte) (bool, error) {
	return _SequencerInbox.Contract.IsValidKeysetHash(&_SequencerInbox.CallOpts, ksHash)
}

// MaxTimeVariation is a free data retrieval call binding the contract method 0xebea461d.
//
// Solidity: function maxTimeVariation() view returns(uint256 delayBlocks, uint256 futureBlocks, uint256 delaySeconds, uint256 futureSeconds)
func (_SequencerInbox *SequencerInboxCaller) MaxTimeVariation(opts *bind.CallOpts) (struct {
	DelayBlocks   *big.Int
	FutureBlocks  *big.Int
	DelaySeconds  *big.Int
	FutureSeconds *big.Int
}, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "maxTimeVariation")

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
func (_SequencerInbox *SequencerInboxSession) MaxTimeVariation() (struct {
	DelayBlocks   *big.Int
	FutureBlocks  *big.Int
	DelaySeconds  *big.Int
	FutureSeconds *big.Int
}, error) {
	return _SequencerInbox.Contract.MaxTimeVariation(&_SequencerInbox.CallOpts)
}

// MaxTimeVariation is a free data retrieval call binding the contract method 0xebea461d.
//
// Solidity: function maxTimeVariation() view returns(uint256 delayBlocks, uint256 futureBlocks, uint256 delaySeconds, uint256 futureSeconds)
func (_SequencerInbox *SequencerInboxCallerSession) MaxTimeVariation() (struct {
	DelayBlocks   *big.Int
	FutureBlocks  *big.Int
	DelaySeconds  *big.Int
	FutureSeconds *big.Int
}, error) {
	return _SequencerInbox.Contract.MaxTimeVariation(&_SequencerInbox.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInbox *SequencerInboxCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInbox *SequencerInboxSession) Rollup() (common.Address, error) {
	return _SequencerInbox.Contract.Rollup(&_SequencerInbox.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInbox *SequencerInboxCallerSession) Rollup() (common.Address, error) {
	return _SequencerInbox.Contract.Rollup(&_SequencerInbox.CallOpts)
}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInbox *SequencerInboxCaller) TotalDelayedMessagesRead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "totalDelayedMessagesRead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInbox *SequencerInboxSession) TotalDelayedMessagesRead() (*big.Int, error) {
	return _SequencerInbox.Contract.TotalDelayedMessagesRead(&_SequencerInbox.CallOpts)
}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInbox *SequencerInboxCallerSession) TotalDelayedMessagesRead() (*big.Int, error) {
	return _SequencerInbox.Contract.TotalDelayedMessagesRead(&_SequencerInbox.CallOpts)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x61a93f8f.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInbox *SequencerInboxTransactor) AddSequencerL2Batch(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "addSequencerL2Batch", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x61a93f8f.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInbox *SequencerInboxSession) AddSequencerL2Batch(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2Batch(&_SequencerInbox.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0x61a93f8f.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) AddSequencerL2Batch(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2Batch(&_SequencerInbox.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInbox *SequencerInboxTransactor) AddSequencerL2BatchFromOrigin(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "addSequencerL2BatchFromOrigin", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInbox *SequencerInboxSession) AddSequencerL2BatchFromOrigin(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2BatchFromOrigin(&_SequencerInbox.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) AddSequencerL2BatchFromOrigin(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.AddSequencerL2BatchFromOrigin(&_SequencerInbox.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInbox *SequencerInboxTransactor) ForceInclusion(opts *bind.TransactOpts, _totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "forceInclusion", _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInbox *SequencerInboxSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.ForceInclusion(&_SequencerInbox.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.ForceInclusion(&_SequencerInbox.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f7a92b2.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInbox *SequencerInboxTransactor) Initialize(opts *bind.TransactOpts, bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "initialize", bridge_, maxTimeVariation_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f7a92b2.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInbox *SequencerInboxSession) Initialize(bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInbox.Contract.Initialize(&_SequencerInbox.TransactOpts, bridge_, maxTimeVariation_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f7a92b2.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) Initialize(bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInbox.Contract.Initialize(&_SequencerInbox.TransactOpts, bridge_, maxTimeVariation_)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInbox *SequencerInboxTransactor) InvalidateKeysetHash(opts *bind.TransactOpts, ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "invalidateKeysetHash", ksHash)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInbox *SequencerInboxSession) InvalidateKeysetHash(ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.InvalidateKeysetHash(&_SequencerInbox.TransactOpts, ksHash)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) InvalidateKeysetHash(ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.InvalidateKeysetHash(&_SequencerInbox.TransactOpts, ksHash)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInbox *SequencerInboxTransactor) SetIsBatchPoster(opts *bind.TransactOpts, addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "setIsBatchPoster", addr, isBatchPoster_)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInbox *SequencerInboxSession) SetIsBatchPoster(addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SetIsBatchPoster(&_SequencerInbox.TransactOpts, addr, isBatchPoster_)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) SetIsBatchPoster(addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SetIsBatchPoster(&_SequencerInbox.TransactOpts, addr, isBatchPoster_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInbox *SequencerInboxTransactor) SetMaxTimeVariation(opts *bind.TransactOpts, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "setMaxTimeVariation", maxTimeVariation_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInbox *SequencerInboxSession) SetMaxTimeVariation(maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SetMaxTimeVariation(&_SequencerInbox.TransactOpts, maxTimeVariation_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) SetMaxTimeVariation(maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SetMaxTimeVariation(&_SequencerInbox.TransactOpts, maxTimeVariation_)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInbox *SequencerInboxTransactor) SetValidKeyset(opts *bind.TransactOpts, keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "setValidKeyset", keysetBytes)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInbox *SequencerInboxSession) SetValidKeyset(keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SetValidKeyset(&_SequencerInbox.TransactOpts, keysetBytes)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) SetValidKeyset(keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SetValidKeyset(&_SequencerInbox.TransactOpts, keysetBytes)
}

// SequencerInboxInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the SequencerInbox contract.
type SequencerInboxInboxMessageDeliveredIterator struct {
	Event *SequencerInboxInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *SequencerInboxInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxInboxMessageDelivered)
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
		it.Event = new(SequencerInboxInboxMessageDelivered)
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
func (it *SequencerInboxInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxInboxMessageDelivered represents a InboxMessageDelivered event raised by the SequencerInbox contract.
type SequencerInboxInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_SequencerInbox *SequencerInboxFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*SequencerInboxInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxInboxMessageDeliveredIterator{contract: _SequencerInbox.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_SequencerInbox *SequencerInboxFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *SequencerInboxInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxInboxMessageDelivered)
				if err := _SequencerInbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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
func (_SequencerInbox *SequencerInboxFilterer) ParseInboxMessageDelivered(log types.Log) (*SequencerInboxInboxMessageDelivered, error) {
	event := new(SequencerInboxInboxMessageDelivered)
	if err := _SequencerInbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the SequencerInbox contract.
type SequencerInboxInboxMessageDeliveredFromOriginIterator struct {
	Event *SequencerInboxInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *SequencerInboxInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxInboxMessageDeliveredFromOrigin)
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
		it.Event = new(SequencerInboxInboxMessageDeliveredFromOrigin)
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
func (it *SequencerInboxInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the SequencerInbox contract.
type SequencerInboxInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_SequencerInbox *SequencerInboxFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*SequencerInboxInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxInboxMessageDeliveredFromOriginIterator{contract: _SequencerInbox.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_SequencerInbox *SequencerInboxFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *SequencerInboxInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxInboxMessageDeliveredFromOrigin)
				if err := _SequencerInbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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
func (_SequencerInbox *SequencerInboxFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*SequencerInboxInboxMessageDeliveredFromOrigin, error) {
	event := new(SequencerInboxInboxMessageDeliveredFromOrigin)
	if err := _SequencerInbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxInvalidateKeysetIterator is returned from FilterInvalidateKeyset and is used to iterate over the raw logs and unpacked data for InvalidateKeyset events raised by the SequencerInbox contract.
type SequencerInboxInvalidateKeysetIterator struct {
	Event *SequencerInboxInvalidateKeyset // Event containing the contract specifics and raw log

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
func (it *SequencerInboxInvalidateKeysetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxInvalidateKeyset)
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
		it.Event = new(SequencerInboxInvalidateKeyset)
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
func (it *SequencerInboxInvalidateKeysetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxInvalidateKeysetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxInvalidateKeyset represents a InvalidateKeyset event raised by the SequencerInbox contract.
type SequencerInboxInvalidateKeyset struct {
	KeysetHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvalidateKeyset is a free log retrieval operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_SequencerInbox *SequencerInboxFilterer) FilterInvalidateKeyset(opts *bind.FilterOpts, keysetHash [][32]byte) (*SequencerInboxInvalidateKeysetIterator, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "InvalidateKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxInvalidateKeysetIterator{contract: _SequencerInbox.contract, event: "InvalidateKeyset", logs: logs, sub: sub}, nil
}

// WatchInvalidateKeyset is a free log subscription operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_SequencerInbox *SequencerInboxFilterer) WatchInvalidateKeyset(opts *bind.WatchOpts, sink chan<- *SequencerInboxInvalidateKeyset, keysetHash [][32]byte) (event.Subscription, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "InvalidateKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxInvalidateKeyset)
				if err := _SequencerInbox.contract.UnpackLog(event, "InvalidateKeyset", log); err != nil {
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
func (_SequencerInbox *SequencerInboxFilterer) ParseInvalidateKeyset(log types.Log) (*SequencerInboxInvalidateKeyset, error) {
	event := new(SequencerInboxInvalidateKeyset)
	if err := _SequencerInbox.contract.UnpackLog(event, "InvalidateKeyset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxOwnerFunctionCalledIterator is returned from FilterOwnerFunctionCalled and is used to iterate over the raw logs and unpacked data for OwnerFunctionCalled events raised by the SequencerInbox contract.
type SequencerInboxOwnerFunctionCalledIterator struct {
	Event *SequencerInboxOwnerFunctionCalled // Event containing the contract specifics and raw log

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
func (it *SequencerInboxOwnerFunctionCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxOwnerFunctionCalled)
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
		it.Event = new(SequencerInboxOwnerFunctionCalled)
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
func (it *SequencerInboxOwnerFunctionCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxOwnerFunctionCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxOwnerFunctionCalled represents a OwnerFunctionCalled event raised by the SequencerInbox contract.
type SequencerInboxOwnerFunctionCalled struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOwnerFunctionCalled is a free log retrieval operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_SequencerInbox *SequencerInboxFilterer) FilterOwnerFunctionCalled(opts *bind.FilterOpts, id []*big.Int) (*SequencerInboxOwnerFunctionCalledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "OwnerFunctionCalled", idRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxOwnerFunctionCalledIterator{contract: _SequencerInbox.contract, event: "OwnerFunctionCalled", logs: logs, sub: sub}, nil
}

// WatchOwnerFunctionCalled is a free log subscription operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_SequencerInbox *SequencerInboxFilterer) WatchOwnerFunctionCalled(opts *bind.WatchOpts, sink chan<- *SequencerInboxOwnerFunctionCalled, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "OwnerFunctionCalled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxOwnerFunctionCalled)
				if err := _SequencerInbox.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
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
func (_SequencerInbox *SequencerInboxFilterer) ParseOwnerFunctionCalled(log types.Log) (*SequencerInboxOwnerFunctionCalled, error) {
	event := new(SequencerInboxOwnerFunctionCalled)
	if err := _SequencerInbox.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxSequencerBatchDataIterator is returned from FilterSequencerBatchData and is used to iterate over the raw logs and unpacked data for SequencerBatchData events raised by the SequencerInbox contract.
type SequencerInboxSequencerBatchDataIterator struct {
	Event *SequencerInboxSequencerBatchData // Event containing the contract specifics and raw log

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
func (it *SequencerInboxSequencerBatchDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxSequencerBatchData)
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
		it.Event = new(SequencerInboxSequencerBatchData)
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
func (it *SequencerInboxSequencerBatchDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxSequencerBatchDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxSequencerBatchData represents a SequencerBatchData event raised by the SequencerInbox contract.
type SequencerInboxSequencerBatchData struct {
	BatchSequenceNumber *big.Int
	Data                []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchData is a free log retrieval operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_SequencerInbox *SequencerInboxFilterer) FilterSequencerBatchData(opts *bind.FilterOpts, batchSequenceNumber []*big.Int) (*SequencerInboxSequencerBatchDataIterator, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "SequencerBatchData", batchSequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxSequencerBatchDataIterator{contract: _SequencerInbox.contract, event: "SequencerBatchData", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchData is a free log subscription operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_SequencerInbox *SequencerInboxFilterer) WatchSequencerBatchData(opts *bind.WatchOpts, sink chan<- *SequencerInboxSequencerBatchData, batchSequenceNumber []*big.Int) (event.Subscription, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "SequencerBatchData", batchSequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxSequencerBatchData)
				if err := _SequencerInbox.contract.UnpackLog(event, "SequencerBatchData", log); err != nil {
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
func (_SequencerInbox *SequencerInboxFilterer) ParseSequencerBatchData(log types.Log) (*SequencerInboxSequencerBatchData, error) {
	event := new(SequencerInboxSequencerBatchData)
	if err := _SequencerInbox.contract.UnpackLog(event, "SequencerBatchData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxSequencerBatchDeliveredIterator is returned from FilterSequencerBatchDelivered and is used to iterate over the raw logs and unpacked data for SequencerBatchDelivered events raised by the SequencerInbox contract.
type SequencerInboxSequencerBatchDeliveredIterator struct {
	Event *SequencerInboxSequencerBatchDelivered // Event containing the contract specifics and raw log

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
func (it *SequencerInboxSequencerBatchDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxSequencerBatchDelivered)
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
		it.Event = new(SequencerInboxSequencerBatchDelivered)
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
func (it *SequencerInboxSequencerBatchDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxSequencerBatchDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxSequencerBatchDelivered represents a SequencerBatchDelivered event raised by the SequencerInbox contract.
type SequencerInboxSequencerBatchDelivered struct {
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
func (_SequencerInbox *SequencerInboxFilterer) FilterSequencerBatchDelivered(opts *bind.FilterOpts, batchSequenceNumber []*big.Int, beforeAcc [][32]byte, afterAcc [][32]byte) (*SequencerInboxSequencerBatchDeliveredIterator, error) {

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

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "SequencerBatchDelivered", batchSequenceNumberRule, beforeAccRule, afterAccRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxSequencerBatchDeliveredIterator{contract: _SequencerInbox.contract, event: "SequencerBatchDelivered", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchDelivered is a free log subscription operation binding the contract event 0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed batchSequenceNumber, bytes32 indexed beforeAcc, bytes32 indexed afterAcc, bytes32 delayedAcc, uint256 afterDelayedMessagesRead, (uint64,uint64,uint64,uint64) timeBounds, uint8 dataLocation)
func (_SequencerInbox *SequencerInboxFilterer) WatchSequencerBatchDelivered(opts *bind.WatchOpts, sink chan<- *SequencerInboxSequencerBatchDelivered, batchSequenceNumber []*big.Int, beforeAcc [][32]byte, afterAcc [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "SequencerBatchDelivered", batchSequenceNumberRule, beforeAccRule, afterAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxSequencerBatchDelivered)
				if err := _SequencerInbox.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
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
func (_SequencerInbox *SequencerInboxFilterer) ParseSequencerBatchDelivered(log types.Log) (*SequencerInboxSequencerBatchDelivered, error) {
	event := new(SequencerInboxSequencerBatchDelivered)
	if err := _SequencerInbox.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxSetValidKeysetIterator is returned from FilterSetValidKeyset and is used to iterate over the raw logs and unpacked data for SetValidKeyset events raised by the SequencerInbox contract.
type SequencerInboxSetValidKeysetIterator struct {
	Event *SequencerInboxSetValidKeyset // Event containing the contract specifics and raw log

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
func (it *SequencerInboxSetValidKeysetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxSetValidKeyset)
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
		it.Event = new(SequencerInboxSetValidKeyset)
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
func (it *SequencerInboxSetValidKeysetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxSetValidKeysetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxSetValidKeyset represents a SetValidKeyset event raised by the SequencerInbox contract.
type SequencerInboxSetValidKeyset struct {
	KeysetHash  [32]byte
	KeysetBytes []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetValidKeyset is a free log retrieval operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_SequencerInbox *SequencerInboxFilterer) FilterSetValidKeyset(opts *bind.FilterOpts, keysetHash [][32]byte) (*SequencerInboxSetValidKeysetIterator, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "SetValidKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxSetValidKeysetIterator{contract: _SequencerInbox.contract, event: "SetValidKeyset", logs: logs, sub: sub}, nil
}

// WatchSetValidKeyset is a free log subscription operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_SequencerInbox *SequencerInboxFilterer) WatchSetValidKeyset(opts *bind.WatchOpts, sink chan<- *SequencerInboxSetValidKeyset, keysetHash [][32]byte) (event.Subscription, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "SetValidKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxSetValidKeyset)
				if err := _SequencerInbox.contract.UnpackLog(event, "SetValidKeyset", log); err != nil {
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
func (_SequencerInbox *SequencerInboxFilterer) ParseSetValidKeyset(log types.Log) (*SequencerInboxSetValidKeyset, error) {
	event := new(SequencerInboxSetValidKeyset)
	if err := _SequencerInbox.contract.UnpackLog(event, "SetValidKeyset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
