package arbitrum

import (
	"context"

	"github.com/tenderly/nitro/go-ethereum/common/hexutil"
	"github.com/tenderly/nitro/go-ethereum/internal/ethapi"
	"github.com/tenderly/nitro/go-ethereum/rpc"
)

type TransactionArgs = ethapi.TransactionArgs

func EstimateGas(ctx context.Context, b ethapi.Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, gasCap uint64) (hexutil.Uint64, error) {
	return ethapi.DoEstimateGas(ctx, b, args, blockNrOrHash, gasCap)
}
