package arbos

import (
	"github.com/tenderly/nitro/go-ethereum/arbitrum_types"
	"github.com/tenderly/nitro/go-ethereum/common"
	"github.com/tenderly/nitro/go-ethereum/core"
	"github.com/tenderly/nitro/go-ethereum/core/state"
	"github.com/tenderly/nitro/go-ethereum/core/types"
	"github.com/tenderly/nitro/go-ethereum/params"
	"github.com/tenderly/nitro/arbos/arbosState"
)

// extraPreTxFilter should be modified by chain operators to enforce additional pre-transaction validity rules
func extraPreTxFilter(
	chainConfig *params.ChainConfig,
	currentBlockHeader *types.Header,
	statedb *state.StateDB,
	state *arbosState.ArbosState,
	tx *types.Transaction,
	options *arbitrum_types.ConditionalOptions,
	sender common.Address,
	l1Info *L1Info,
) error {
	// TODO: implement additional pre-transaction checks
	return nil
}

// extraPostTxFilter should be modified by chain operators to enforce additional post-transaction validity rules
func extraPostTxFilter(
	chainConfig *params.ChainConfig,
	currentBlockHeader *types.Header,
	statedb *state.StateDB,
	state *arbosState.ArbosState,
	tx *types.Transaction,
	options *arbitrum_types.ConditionalOptions,
	sender common.Address,
	l1Info *L1Info,
	result *core.ExecutionResult,
) error {
	// TODO: implement additional post-transaction checks
	return nil
}
