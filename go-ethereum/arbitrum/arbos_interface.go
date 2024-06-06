package arbitrum

import (
	"context"

	"github.com/tenderly/nitro/go-ethereum/arbitrum_types"
	"github.com/tenderly/nitro/go-ethereum/core"
	"github.com/tenderly/nitro/go-ethereum/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction, options *arbitrum_types.ConditionalOptions) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}
