// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package server_arb

/*
#cgo CFLAGS: -g -Wall -I../../target/include/
#include "arbitrator.h"

ResolvedPreimage preimageResolverC(size_t context, uint8_t preimageType, const uint8_t* hash);
*/
import "C"
import (
	"context"
	"errors"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/tenderly/nitro/arbutil"
	"github.com/tenderly/nitro/go-ethereum/common"
	"github.com/tenderly/nitro/go-ethereum/log"
	"github.com/tenderly/nitro/util/containers"
	"github.com/tenderly/nitro/validator"
)

type MachineInterface interface {
	CloneMachineInterface() MachineInterface
	GetStepCount() uint64
	IsRunning() bool
	ValidForStep(uint64) bool
	Status() uint8
	Step(context.Context, uint64) error
	Hash() common.Hash
	GetGlobalState() validator.GoGlobalState
	ProveNextStep() []byte
	Freeze()
	Destroy()
}

// ArbitratorMachine holds an arbitrator machine pointer, and manages its lifetime
type ArbitratorMachine struct {
	mutex     sync.Mutex // needed because go finalizers don't synchronize (meaning they aren't thread safe)
	contextId *int64 // has a finalizer attached to remove the preimage resolver from the global map
	frozen    bool   // does not allow anything that changes machine state, not cloned with the machine
}

// Assert that ArbitratorMachine implements MachineInterface
var _ MachineInterface = (*ArbitratorMachine)(nil)

var preimageResolvers containers.SyncMap[int64, GoPreimageResolver]
var lastPreimageResolverId int64 // atomic

// Any future calls to this machine will result in a panic
func (m *ArbitratorMachine) Destroy() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.contextId = nil
}

func freeContextId(context *int64) {
	preimageResolvers.Delete(*context)
}

func LoadSimpleMachine(wasm string, libraries []string) (*ArbitratorMachine, error) {
	return nil, nil
}

func (m *ArbitratorMachine) Freeze() {
	m.frozen = true
}

// Even if origin is frozen - clone is not
func (m *ArbitratorMachine) Clone() *ArbitratorMachine {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return nil
}

func (m *ArbitratorMachine) CloneMachineInterface() MachineInterface {
	return m.Clone()
}

func (m *ArbitratorMachine) SetGlobalState(globalState validator.GoGlobalState) error {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.frozen {
		return errors.New("machine frozen")
	}
	return nil
}

func (m *ArbitratorMachine) GetGlobalState() validator.GoGlobalState {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return validator.GoGlobalState{}
}

func (m *ArbitratorMachine) GetStepCount() uint64 {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return 0
}

func (m *ArbitratorMachine) IsRunning() bool {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return false
}

func (m *ArbitratorMachine) IsErrored() bool {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return false
}

func (m *ArbitratorMachine) Status() uint8 {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return 0
}

func (m *ArbitratorMachine) ValidForStep(requestedStep uint64) bool {
	haveStep := m.GetStepCount()
	if haveStep > requestedStep {
		return false
	} else if haveStep == requestedStep {
		return true
	} else { // haveStep < requestedStep
		// if the machine is halted, its state persists for future steps
		return !m.IsRunning()
	}
}

func (m *ArbitratorMachine) Step(ctx context.Context, count uint64) error {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.frozen {
		return errors.New("machine frozen")
	}

	return ctx.Err()
}

func (m *ArbitratorMachine) StepUntilHostIo(ctx context.Context) error {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.frozen {
		return errors.New("machine frozen")
	}

	return ctx.Err()
}

func (m *ArbitratorMachine) Hash() (hash common.Hash) {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return
}

func (m *ArbitratorMachine) GetModuleRoot() (hash common.Hash) {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return
}
func (m *ArbitratorMachine) ProveNextStep() []byte {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return nil
}

func (m *ArbitratorMachine) SerializeState(path string) error {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return nil
}

func (m *ArbitratorMachine) DeserializeAndReplaceState(path string) error {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.frozen {
		return errors.New("machine frozen")
	}

	return nil
}

func (m *ArbitratorMachine) AddSequencerInboxMessage(index uint64, data []byte) error {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.frozen {
		return errors.New("machine frozen")
	}
	return nil
}

func (m *ArbitratorMachine) AddDelayedInboxMessage(index uint64, data []byte) error {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.frozen {
		return errors.New("machine frozen")
	}
	return nil
}

type GoPreimageResolver = func(arbutil.PreimageType, common.Hash) ([]byte, error)

//export preimageResolver
func preimageResolver(context C.size_t, ty C.uint8_t, ptr unsafe.Pointer) C.ResolvedPreimage {
	var hash common.Hash
	input := (*[1 << 30]byte)(ptr)[:32]
	copy(hash[:], input)
	resolver, ok := preimageResolvers.Load(int64(context))
	if !ok {
		return C.ResolvedPreimage{
			len: -1,
		}
	}
	preimage, err := resolver(arbutil.PreimageType(ty), hash)
	if err != nil {
		log.Error("preimage resolution failed", "err", err)
		return C.ResolvedPreimage{
			len: -1,
		}
	}
	return C.ResolvedPreimage{
		ptr: (*C.uint8_t)(C.CBytes(preimage)),
		len: (C.ptrdiff_t)(len(preimage)),
	}
}

func (m *ArbitratorMachine) SetPreimageResolver(resolver GoPreimageResolver) error {
	defer runtime.KeepAlive(m)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if m.frozen {
		return errors.New("machine frozen")
	}
	id := atomic.AddInt64(&lastPreimageResolverId, 1)
	preimageResolvers.Store(id, resolver)
	m.contextId = &id
	runtime.SetFinalizer(m.contextId, freeContextId)
	return nil
}
