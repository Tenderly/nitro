// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package server_arb

import (
	"context"
	"errors"
	"github.com/tenderly/nitro/go-ethereum/common"
	"github.com/tenderly/nitro/go-ethereum/log"
	"github.com/tenderly/nitro/validator/server_common"
	"os"
	"path/filepath"
)

func createArbMachine(ctx context.Context, locator *server_common.MachineLocator, config *ArbitratorMachineConfig, moduleRoot common.Hash) (*arbMachines, error) {
	binPath := filepath.Join(locator.GetMachinePath(moduleRoot), config.WavmBinaryPath)
	log.Info("creating nitro machine", "binpath", binPath)

	result := &arbMachines{}
	result.zeroStep.Freeze()

	// We try to store/load state before first host_io to a file.
	// We will chicken out of that if something fails, but still try to calculate the machine
	statePath := filepath.Join(locator.GetMachinePath(moduleRoot), config.UntilHostIoStatePath)
	_, err := os.Stat(statePath)
	if err == nil {
		log.Info("found cached machine until host io state", "moduleRoot", moduleRoot)

		if err != nil {
			// Safe as if DeserializeAndReplaceState returns an error it will not have mutated the machine
			log.Warn("failed to load machine until host io state; will reexecute", "err", err)
		} else {
			result.hostIo.Freeze()
			return result, nil
		}
	} else if errors.Is(err, os.ErrNotExist) {
		log.Info("didn't find cached machine until host io state", "path", statePath)
	} else {
		log.Warn("error checking if machine until host io state is cached", "path", statePath, "err", err)
	}

	result.hostIo.Freeze()
	return result, nil
}
