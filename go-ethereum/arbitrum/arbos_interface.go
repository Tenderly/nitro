package arbitrum

import (
	"context"

	"github.com/tenderly/nitro/go-ethereum/core"
	"github.com/tenderly/nitro/go-ethereum/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}
