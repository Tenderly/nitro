// Copyright 2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package das

import (
	"encoding/binary"

	"github.com/tenderly/nitro/go-ethereum/common"
	"github.com/tenderly/nitro/go-ethereum/crypto"

	"github.com/tenderly/nitro/das/dastree"
	"github.com/tenderly/nitro/util/signature"
)

var uniquifyingPrefix = []byte("Arbitrum Nitro DAS API Store:")

func applyDasSigner(signer signature.DataSignerFunc, data []byte, extraFields ...uint64) ([]byte, error) {
	return signer(dasStoreHash(data, extraFields...))
}

func DasRecoverSigner(data []byte, sig []byte, extraFields ...uint64) (common.Address, error) {
	pk, err := crypto.SigToPub(dasStoreHash(data, extraFields...), sig)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*pk), nil
}

func dasStoreHash(data []byte, extraFields ...uint64) []byte {
	var buf []byte

	for _, field := range extraFields {
		buf = binary.BigEndian.AppendUint64(buf, field)
	}

	return dastree.HashBytes(uniquifyingPrefix, buf, data)
}
