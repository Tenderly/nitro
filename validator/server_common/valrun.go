package server_common

import (
	"github.com/tenderly/nitro/go-ethereum/common"
	"github.com/tenderly/nitro/util/containers"
	"github.com/tenderly/nitro/validator"
)

type ValRun struct {
	containers.PromiseInterface[validator.GoGlobalState]
	root common.Hash
}

func (r *ValRun) WasmModuleRoot() common.Hash {
	return r.root
}

func NewValRun(promise containers.PromiseInterface[validator.GoGlobalState], root common.Hash) *ValRun {
	return &ValRun{
		PromiseInterface: promise,
		root:             root,
	}
}
